// Package check checks that every id an annotation names is approved in the base commit of the
// pull request (TOOL-06).
package check

import (
	"fmt"
	"go/scanner"
	"go/token"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/treporting/specs-and-deploy/tools/review"
	"gopkg.in/yaml.v3"
)

// Git is what the check asks of the history of the repo.
type Git interface {
	// Show answers the content of path in the commit rev.
	Show(rev, path string) ([]byte, error)
}

// Annotation is one @implements or @verifies, and the ids it names.
type Annotation struct {
	File string
	Line int
	Kind string // implements or verifies
	IDs  []string
}

// Finding is one id an annotation names that is missing or not approved in the base commit.
type Finding struct {
	Annotation Annotation
	ID         string
	Message    string
}

func (f Finding) String() string {
	return fmt.Sprintf("%s:%d @%s %s: %s", f.Annotation.File, f.Annotation.Line, f.Annotation.Kind, f.ID, f.Message)
}

// annotation is an @implements or @verifies, and the ids it names, separated by commas.
var annotation = regexp.MustCompile(`@(implements|verifies)\s+([A-Z][A-Z0-9]*-[A-Z0-9-]*[0-9](?:\s*,\s*[A-Z][A-Z0-9]*-[A-Z0-9-]*[0-9])*)`)

// notCode tells the files git knows that are no code by PROC-05 - neither built nor run - so that
// an example of an annotation in them is none: documents, schemas and images. Our data, the YAML
// that names a schema in its field schema, is told by its content (isData).
var notCode = []func(file string) bool{
	func(f string) bool { return strings.HasSuffix(f, ".md") },
	func(f string) bool { return strings.HasPrefix(f, "schemas/") && strings.HasSuffix(f, ".schema.json") },
	func(f string) bool {
		for _, ext := range []string{".png", ".jpg", ".jpeg", ".gif", ".svg", ".pdf"} {
			if strings.HasSuffix(strings.ToLower(f), ext) {
				return true
			}
		}
		return false
	},
}

// IsCode says whether a file git knows is code by PROC-05: every file, but documents, schemas,
// images, and our data.
//
// @implements SR-24
func IsCode(file string, content []byte) bool {
	for _, no := range notCode {
		if no(file) {
			return false
		}
	}
	return !isData(file, content)
}

// isData says whether a file is our data: a YAML that names its schema in its field schema.
func isData(file string, content []byte) bool {
	if !strings.HasSuffix(file, ".yaml") && !strings.HasSuffix(file, ".yml") {
		return false
	}
	var doc struct {
		Schema string `yaml:"schema"`
	}
	return yaml.Unmarshal(content, &doc) == nil && doc.Schema != ""
}

// Find answers every annotation in the files, relative to root, that are code. An annotation
// stands in a comment (formats.md, section Annotations): in Go the scanner of Go tells comments
// from strings; in another file a comment is a line that begins with //, #, /* or *.
//
// @implements SR-24
func Find(root string, files []string) ([]Annotation, error) {

	var found []Annotation
	for _, file := range files {
		data, err := os.ReadFile(filepath.Join(root, file))
		if err != nil {
			return nil, err
		}
		if !IsCode(file, data) {
			continue
		}
		for _, c := range comments(file, data) {
			for _, m := range annotation.FindAllStringSubmatch(c.text, -1) {
				var ids []string
				for _, id := range strings.Split(m[2], ",") {
					ids = append(ids, strings.TrimSpace(id))
				}
				found = append(found, Annotation{file, c.line, m[1], ids})
			}
		}
	}
	return found, nil
}

// comment is the text of one line of a comment, and the number of that line.
type comment struct {
	line int
	text string
}

// comments answers the lines of the comments of a file.
func comments(file string, data []byte) []comment {

	var out []comment
	if strings.HasSuffix(file, ".go") {
		fset := token.NewFileSet()
		f := fset.AddFile(file, -1, len(data))
		var s scanner.Scanner
		s.Init(f, data, nil, scanner.ScanComments)
		for {
			pos, tok, lit := s.Scan()
			if tok == token.EOF {
				break
			}
			if tok == token.COMMENT {
				first := fset.Position(pos).Line
				for i, l := range strings.Split(lit, "\n") {
					out = append(out, comment{first + i, l})
				}
			}
		}
		return out
	}
	for i, l := range strings.Split(string(data), "\n") {
		l = strings.TrimSpace(l)
		for _, mark := range []string{"//", "#", "/*", "*"} {
			if strings.HasPrefix(l, mark) {
				out = append(out, comment{i + 1, l})
				break
			}
		}
	}
	return out
}

// Check answers each id the annotations name that is not in the requirements of the base commit,
// or not approved there.
//
// @implements SR-24, SR-25
func Check(git Git, base string, annotations []Annotation) []Finding {

	var findings []Finding
	for _, a := range annotations {
		for _, id := range a.IDs {
			if message := status(git, base, id); message != "" {
				findings = append(findings, Finding{a, id, message})
			}
		}
	}
	return findings
}

// status answers why id does not count in the base commit, or "" where it is approved there.
func status(git Git, base, id string) string {

	file, key := review.FileOf(id)
	if file == "" || key == "" {
		return "no file of requirements holds such an id"
	}
	data, err := git.Show(base, file)
	if err != nil {
		return fmt.Sprintf("%s is not in the base commit %.7s", file, base)
	}
	object := review.ObjectIn(data, key)
	if object == nil {
		return fmt.Sprintf("not in %s in the base commit %.7s", file, base)
	}
	var o struct {
		Status string `yaml:"status"`
	}
	if err := yaml.Unmarshal(object, &o); err != nil {
		return err.Error()
	}
	if o.Status != "approved" {
		return fmt.Sprintf("%s in the base commit %.7s, not approved", o.Status, base)
	}
	return ""
}
