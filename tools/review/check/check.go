// Package check checks the results of reviews against the rules of formats.md (TOOL_03, TOOL_11).
package check

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/treporting/specs-and-deploy/tools/review"
	"gopkg.in/yaml.v3"
)

// Review is one file of review results.
type Review struct {
	Level            string   `yaml:"level"`
	CreatedBy        string   `yaml:"created_by"`
	CreatedAt        string   `yaml:"created_at"`
	Criteria         string   `yaml:"criteria"`
	CriteriaRevision string   `yaml:"criteria_revision"`
	Results          []Result `yaml:"results"`
}

// Result is the review of one requirement or document.
type Result struct {
	ObjectKey string               `yaml:"object_key"`
	TargetIDs []string             `yaml:"target_ids"`
	Revision  string               `yaml:"revision"`
	Criteria  map[string]Judgement `yaml:"criteria"`
	Verdict   string               `yaml:"verdict"`
}

// Judgement is the result of one criterion: `passed`, or a block of result and note.
type Judgement struct {
	Result string `yaml:"result"`
	Note   string `yaml:"note"`
}

// UnmarshalYAML reads a judgement written as a word or as a block.
func (j *Judgement) UnmarshalYAML(node *yaml.Node) error {
	if node.Kind == yaml.ScalarNode {
		j.Result = node.Value
		return nil
	}
	type plain Judgement
	return node.Decode((*plain)(j))
}

// Finding is one rule a review result breaks.
type Finding struct {
	File    string
	Result  string
	Message string
}

func (f Finding) String() string {
	return fmt.Sprintf("%s %s: %s", f.File, f.Result, f.Message)
}

// Git is what the check asks of the history of the repo.
type Git interface {
	// CommitExists says whether rev names a commit.
	CommitExists(rev string) bool
	// Show answers the content of path in the commit rev.
	Show(rev, path string) ([]byte, error)
	// FileCommits answers the commits up to rev that changed path, the newest first.
	FileCommits(rev, path string) ([]string, error)
	// Parent answers the first parent of a commit, and false for the first commit.
	Parent(commit string) (string, bool)
	// IsMerge says whether a commit has more than one parent.
	IsMerge(commit string) bool
	// Authors answers the author of a commit and the names of its co-authors.
	Authors(commit string) (author string, coAuthors []string, err error)
}

// Check reads the review result in path, below root, and answers every rule it breaks.
//
// @implements SR_01@1.1.1, SR_02@1.1.1, SR_03@1.1.1, SR_04@1.1.1, SR_05@1.1.1, SR_06@1.1.1, SR_07@1.1.1, SR_08@1.1.1, SR_16@1.1.1
func Check(root, path string, git Git) ([]Finding, error) {

	data, err := os.ReadFile(filepath.Join(root, path))
	if err != nil {
		return nil, err
	}
	var r Review
	if err := yaml.Unmarshal(data, &r); err != nil {
		return nil, fmt.Errorf("%s: %w", path, err)
	}

	criteria, err := criteriaKeys(git, r.CriteriaRevision, r.Criteria)
	if err != nil {
		return []Finding{{path, "", err.Error()}}, nil
	}

	var findings []Finding
	for _, result := range r.Results {
		for _, message := range checkResult(r, result, criteria, git) {
			findings = append(findings, Finding{path, result.ObjectKey, message})
		}
	}
	return findings, nil
}

// criteriaKeys answers the ids of the criteria in the criteria file, in the revision the review used.
// The revision is a commit that changed the file, so that the criteria the review followed can be
// traced.
//
// @implements SR_07@1.1.1
func criteriaKeys(git Git, revision, file string) ([]string, error) {

	if !git.CommitExists(revision) {
		return nil, fmt.Errorf("the criteria revision %s is no commit", revision)
	}
	commits, err := git.FileCommits(revision, file)
	if err != nil {
		return nil, err
	}
	if len(commits) == 0 || !strings.HasPrefix(commits[0], revision) {
		return nil, fmt.Errorf("the criteria revision %s did not change %s", revision, file)
	}
	data, err := git.Show(revision, file)
	if err != nil {
		return nil, fmt.Errorf("the criteria %s are not in %s", file, revision)
	}
	var c struct {
		Criteria []struct {
			ObjectKey string `yaml:"object_key"`
		} `yaml:"criteria"`
	}
	if err := yaml.Unmarshal(data, &c); err != nil {
		return nil, fmt.Errorf("%s in %s: %w", file, revision, err)
	}
	keys := make([]string, 0, len(c.Criteria))
	for _, k := range c.Criteria {
		keys = append(keys, k.ObjectKey)
	}
	return keys, nil
}

// checkResult answers the rules one result breaks.
//
// @implements SR_01@1.1.1, SR_02@1.1.1, SR_03@1.1.1, SR_04@1.1.1, SR_05@1.1.1, SR_06@1.1.1, SR_08@1.1.1, SR_16@1.1.1
func checkResult(r Review, result Result, criteria []string, git Git) []string {

	var messages []string

	//	every criterion once - a map has every key once, so what is left to
	//	ask is whether one is missing or one is not a criterion of the file;
	//	walked by criterion in the order of the file, those not in it by id
	known := map[string]bool{}
	for _, key := range criteria {
		known[key] = true
	}
	var foreign []string
	for key := range result.Criteria {
		if !known[key] {
			foreign = append(foreign, key)
		}
	}
	sort.Strings(foreign)

	allPassed := true
	for _, key := range append(append([]string(nil), criteria...), foreign...) {
		j, ok := result.Criteria[key]
		if !ok {
			messages = append(messages, "criterion "+key+" is missing")
			continue
		}
		if !known[key] {
			messages = append(messages, "criterion "+key+" is not in "+r.Criteria)
		}
		switch j.Result {
		case "passed":
		case "failed":
			allPassed = false
			if strings.TrimSpace(j.Note) == "" {
				messages = append(messages, "criterion "+key+" failed with no note")
			}
		default:
			allPassed = false
			messages = append(messages, "criterion "+key+" is neither passed nor failed")
		}
	}

	switch {
	case result.Verdict == "passed" && !allPassed:
		messages = append(messages, "the verdict is passed, but not every criterion passed")
	case result.Verdict == "failed" && allPassed:
		messages = append(messages, "the verdict is failed, but every criterion passed")
	case result.Verdict != "passed" && result.Verdict != "failed":
		messages = append(messages, "the verdict is neither passed nor failed")
	}

	if !git.CommitExists(result.Revision) {
		return append(messages, "the revision "+result.Revision+" is no commit")
	}

	for _, target := range result.TargetIDs {
		authors, err := authorsOf(git, result.Revision, target)
		if err != nil {
			messages = append(messages, err.Error())
			continue
		}
		for _, name := range authors {
			if name == r.CreatedBy {
				messages = append(messages, "the reviewer "+r.CreatedBy+" is an author of "+target)
				break
			}
		}
	}
	return messages
}

// authorsOf answers the authors of a requirement or document in a revision: the author and the
// co-authors of the last commit, not a merge, that changed it, whatever names stand there. A merge
// joins what others wrote; the one who merged approves, and is no author.
//
// @implements SR_05@1.1.1, SR_08@1.1.1
func authorsOf(git Git, revision, target string) ([]string, error) {

	file, key := review.FileOf(target)
	if file == "" {
		return nil, fmt.Errorf("no file holds %s", target)
	}

	commits, err := git.FileCommits(revision, file)
	if err != nil {
		return nil, err
	}
	for _, commit := range commits {
		if git.IsMerge(commit) {
			continue
		}
		now := objectAt(git, commit, file, key)
		if now == nil {
			continue
		}
		before := []byte(nil)
		if parent, ok := git.Parent(commit); ok {
			before = objectAt(git, parent, file, key)
		}
		if bytes.Equal(now, before) {
			continue
		}
		author, coAuthors, err := git.Authors(commit)
		if err != nil {
			return nil, err
		}
		return append([]string{author}, coAuthors...), nil
	}
	return nil, fmt.Errorf("%s is not in %s in %s", target, file, revision)
}

// objectAt answers the object with key in file in a commit, as YAML, or the whole file where key
// is empty; nil where it is not there.
func objectAt(git Git, commit, file, key string) []byte {

	data, err := git.Show(commit, file)
	if err != nil {
		return nil
	}
	return review.ObjectIn(data, key)
}
