// Package review checks the results of reviews against the rules of formats.md (TOOL-03).
package review

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"

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
	// Authors answers the author of a commit and the names of its co-authors.
	Authors(commit string) (author string, coAuthors []string, err error)
}

// requirementFiles tell by the prefix of an id the file that holds it.
var requirementFiles = []struct{ prefix, file string }{
	{"UN-", "requirements/user-needs.yaml"},
	{"REG-", "requirements/regulatory-requirements.yaml"},
	{"SYS-", "requirements/system-requirements.yaml"},
	{"TOOL-", "requirements/tool-requirements.yaml"},
	{"SR-", "tools/software-requirements.yaml"},
	{"ARC-", "architecture/system-architecture.yaml"},
	{"HAZC-", "safety/hazards.yaml"},
	{"HAZ-", "safety/hazards.yaml"},
}

// lists are the sections of a YAML that hold objects with an object_key.
var lists = []string{"requirements", "elements", "hazards", "controls"}

// Check reads the review result in path, below root, and answers every rule it breaks.
//
// @implements SR-01, SR-02, SR-03, SR-04, SR-05, SR-06, SR-07, SR-08
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
// @implements SR-07
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
// @implements SR-01, SR-02, SR-03, SR-04, SR-05, SR-06, SR-08
func checkResult(r Review, result Result, criteria []string, git Git) []string {

	var messages []string

	//	every criterion once - a map has every key once, so what is left to
	//	ask is whether one is missing or one is not a criterion of the file
	known := map[string]bool{}
	for _, key := range criteria {
		known[key] = true
		if _, ok := result.Criteria[key]; !ok {
			messages = append(messages, "criterion "+key+" is missing")
		}
	}
	allPassed := true
	for key, j := range result.Criteria {
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
// co-authors of the last commit that changed it, whatever names stand there.
//
// @implements SR-05, SR-08
func authorsOf(git Git, revision, target string) ([]string, error) {

	file, key := fileOf(target)
	if file == "" {
		return nil, fmt.Errorf("no file holds %s", target)
	}

	commits, err := git.FileCommits(revision, file)
	if err != nil {
		return nil, err
	}
	for _, commit := range commits {
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

// fileOf answers the file that holds a target, and the key to find it there; a document is its
// own file and has no key.
func fileOf(target string) (file, key string) {

	if strings.HasSuffix(target, ".md") {
		return target, ""
	}
	for _, rf := range requirementFiles {
		if strings.HasPrefix(target, rf.prefix) {
			return rf.file, target
		}
	}
	return "", ""
}

// objectAt answers the object with key in file in a commit, as YAML, or the whole file where key
// is empty; nil where it is not there.
func objectAt(git Git, commit, file, key string) []byte {

	data, err := git.Show(commit, file)
	if err != nil {
		return nil
	}
	if key == "" {
		return data
	}
	var doc map[string]any
	if yaml.Unmarshal(data, &doc) != nil {
		return nil
	}
	for _, list := range lists {
		items, _ := doc[list].([]any)
		for _, item := range items {
			object, _ := item.(map[string]any)
			if object["object_key"] == key {
				out, _ := yaml.Marshal(object)
				return out
			}
		}
	}
	return nil
}
