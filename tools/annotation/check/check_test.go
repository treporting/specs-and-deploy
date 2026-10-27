package check

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// fakeGit holds the files of the base commit in memory.
type fakeGit map[string]string

func (f fakeGit) Show(rev, path string) ([]byte, error) {
	if c, ok := f[path]; ok {
		return []byte(c), nil
	}
	return nil, os.ErrNotExist
}

const requirements = `schema: schemas/software-requirements.schema.json
level: r5
requirements:
- object_key: SR-01
  status: approved
- object_key: SR-02
  status: draft
`

var base = fakeGit{"tools/software-requirements.yaml": requirements}

func write(t *testing.T, files map[string]string) string {
	t.Helper()
	dir := t.TempDir()
	for path, content := range files {
		os.MkdirAll(filepath.Join(dir, filepath.Dir(path)), 0o755)
		if err := os.WriteFile(filepath.Join(dir, path), []byte(content), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	return dir
}

func messages(findings []Finding) []string {
	var out []string
	for _, f := range findings {
		out = append(out, f.String())
	}
	return out
}

// @verifies SR-24
func TestOnlyCodeIsRead(t *testing.T) {
	files := map[string]string{
		"a.go":                  "// @implements SR-01\nfunc a() {}\n",
		"b.js":                  "/** @verifies SR-01 */\n",
		"Dockerfile":            "# @implements SR-01\nFROM scratch\n",
		"doc.md":                "// @implements SR-99\n",
		"schemas/x.schema.json": `{"$comment": "@implements SR-99"}`,
		"data.yaml":             "schema: schemas/x.schema.json\nnote: '@implements SR-99'\n",
		"picture.png":           "@implements SR-99",
		"config.yaml":           "# @implements SR-01\nkey: value\n",
	}
	dir := write(t, files)
	var names []string
	for f := range files {
		names = append(names, f)
	}
	found, err := Find(dir, names)
	if err != nil {
		t.Fatal(err)
	}
	var got []string
	for _, a := range found {
		got = append(got, a.File)
	}
	for _, want := range []string{"a.go", "b.js", "Dockerfile", "config.yaml"} {
		if !strings.Contains(strings.Join(got, " "), want) {
			t.Errorf("want %s read as code, got %q", want, got)
		}
	}
	if len(got) != 4 {
		t.Errorf("want only the four files of code, got %q", got)
	}
}

// @verifies SR-24
func TestAnIdNotInTheBaseCommit(t *testing.T) {
	got := messages(Check(base, "b1", []Annotation{{"a.go", 3, "implements", []string{"SR-01", "SR-07"}}}))
	if len(got) != 1 || got[0] != "a.go:3 @implements SR-07: not in tools/software-requirements.yaml in the base commit b1" {
		t.Errorf("got %q", got)
	}
}

// @verifies SR-25
func TestAnIdNotApprovedInTheBaseCommit(t *testing.T) {
	got := messages(Check(base, "b1", []Annotation{{"a_test.go", 9, "verifies", []string{"SR-02"}}}))
	if len(got) != 1 || got[0] != "a_test.go:9 @verifies SR-02: draft in the base commit b1, not approved" {
		t.Errorf("got %q", got)
	}
}

// @verifies SR-24
func TestAStringIsNoAnnotation(t *testing.T) {
	dir := write(t, map[string]string{"a_test.go": "package a\n\nconst s = \"// @implements SR-99\"\n\n// @verifies SR-01\nfunc TestA() {}\n"})
	found, err := Find(dir, []string{"a_test.go"})
	if err != nil {
		t.Fatal(err)
	}
	if len(found) != 1 || found[0].IDs[0] != "SR-01" || found[0].Line != 5 {
		t.Errorf("want only the comment on line 5, got %+v", found)
	}
}

// @verifies SR-24
func TestSeveralIdsInOneAnnotation(t *testing.T) {
	dir := write(t, map[string]string{"a.go": "// @implements SR-01, SR-02,SR-03\n"})
	found, err := Find(dir, []string{"a.go"})
	if err != nil {
		t.Fatal(err)
	}
	if len(found) != 1 || strings.Join(found[0].IDs, " ") != "SR-01 SR-02 SR-03" {
		t.Errorf("got %+v", found)
	}
}

func writeFile(dir, path, content string) error {
	os.MkdirAll(filepath.Join(dir, filepath.Dir(path)), 0o755)
	return os.WriteFile(filepath.Join(dir, path), []byte(content), 0o644)
}
