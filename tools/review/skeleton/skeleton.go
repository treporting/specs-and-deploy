// Package skeleton writes the skeleton of a review result (TOOL-04): what the reviewer would
// otherwise copy by hand. Every judgement is todo, so an unfilled skeleton fails reviewcheck.
package skeleton

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/treporting/specs-and-deploy/tools/review"
	"gopkg.in/yaml.v3"
)

// Git is what the skeleton asks of the history of the repo.
type Git interface {
	// Show answers the content of path in the commit rev.
	Show(rev, path string) ([]byte, error)
	// FileCommits answers the commits up to rev that changed path, the newest first.
	FileCommits(rev, path string) ([]string, error)
	// Short answers the short hash of rev.
	Short(rev string) (string, error)
}

// Request is what a skeleton is made for.
type Request struct {
	Reviewer string
	Scope    string
	Targets  []string // ids of requirements, or paths of documents
	Now      time.Time
}

// Make writes the skeleton of a review result for req into reviews/ below root, made on the commit
// HEAD, and answers its path relative to root. Where a target is not found, or the targets belong
// to more than one level, it writes nothing.
//
// @implements SR-11, SR-12, SR-13, SR-14, SR-15
func Make(root string, req Request, git Git) (string, error) {

	head, err := git.Short("HEAD")
	if err != nil {
		return "", err
	}

	level, err := levelOf(git, req.Targets)
	if err != nil {
		return "", err
	}

	criteriaFile, criteriaRevision, keys, err := criteriaOf(git, level)
	if err != nil {
		return "", err
	}

	path := filepath.Join("reviews", req.Now.Format("20060102T150405")+"_"+req.Scope+".yaml")
	if _, err := os.Stat(filepath.Join(root, path)); err == nil {
		return "", fmt.Errorf("%s is there already", path)
	}

	var b strings.Builder
	fmt.Fprintf(&b, "# Review of %s against the criteria of the level %s.\n", strings.Join(req.Targets, ", "), level)
	fmt.Fprintf(&b, "# Skeleton made by reviewskeleton: every todo is for the reviewer to decide.\n")
	fmt.Fprintf(&b, "# Format: ../formats.md, section Reviews.\n")
	fmt.Fprintf(&b, "level: review\n")
	fmt.Fprintf(&b, "created_by: %s\n", req.Reviewer)
	fmt.Fprintf(&b, "created_at: '%s'\n", req.Now.Format("2006-01-02T15:04:05-07:00"))
	fmt.Fprintf(&b, "criteria: %s\n", criteriaFile)
	fmt.Fprintf(&b, "criteria_revision: '%s'\n", criteriaRevision)
	fmt.Fprintf(&b, "results:\n")
	for i, target := range req.Targets {
		fmt.Fprintf(&b, "- object_key: RR-%02d\n", i+1)
		fmt.Fprintf(&b, "  target_ids: [%s]\n", target)
		fmt.Fprintf(&b, "  revision: '%s'\n", head)
		fmt.Fprintf(&b, "  criteria:\n")
		for _, key := range keys {
			fmt.Fprintf(&b, "    %s: todo\n", key)
		}
		fmt.Fprintf(&b, "  verdict: todo\n")
	}

	if err := os.WriteFile(filepath.Join(root, path), []byte(b.String()), 0o644); err != nil {
		return "", err
	}
	return path, nil
}

// levelOf answers the one level of the targets in HEAD: the field level of the YAML that holds a
// requirement, and document for a document.
//
// @implements SR-15
func levelOf(git Git, targets []string) (string, error) {

	if len(targets) == 0 {
		return "", fmt.Errorf("no id or document given")
	}
	level := ""
	for _, target := range targets {
		file, key := review.FileOf(target)
		if file == "" {
			return "", fmt.Errorf("no file holds %s", target)
		}
		data, err := git.Show("HEAD", file)
		if err != nil || review.ObjectIn(data, key) == nil {
			return "", fmt.Errorf("%s is not in %s", target, file)
		}
		l := "document"
		if key != "" {
			var doc struct {
				Level string `yaml:"level"`
			}
			if err := yaml.Unmarshal(data, &doc); err != nil {
				return "", fmt.Errorf("%s: %w", file, err)
			}
			l = doc.Level
		}
		if level != "" && l != level {
			return "", fmt.Errorf("%s is of the level %s, the others of %s - one review, one level", target, l, level)
		}
		level = l
	}
	return level, nil
}

// criteriaOf answers the criteria file of a level, the last commit that changed it, and the ids of
// its criteria in that commit.
//
// @implements SR-12
func criteriaOf(git Git, level string) (file, revision string, keys []string, err error) {

	name := "criteria-" + level + ".yaml"
	if level == "document" {
		name = "criteria-documents.yaml"
	}
	file = "reviews/" + name

	commits, err := git.FileCommits("HEAD", file)
	if err != nil || len(commits) == 0 {
		return "", "", nil, fmt.Errorf("no commit holds %s", file)
	}
	if revision, err = git.Short(commits[0]); err != nil {
		return "", "", nil, err
	}
	data, err := git.Show(revision, file)
	if err != nil {
		return "", "", nil, err
	}
	var c struct {
		Criteria []struct {
			ObjectKey string `yaml:"object_key"`
		} `yaml:"criteria"`
	}
	if err := yaml.Unmarshal(data, &c); err != nil {
		return "", "", nil, fmt.Errorf("%s: %w", file, err)
	}
	for _, k := range c.Criteria {
		keys = append(keys, k.ObjectKey)
	}
	return file, revision, keys, nil
}
