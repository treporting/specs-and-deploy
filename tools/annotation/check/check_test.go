package check

import (
	"os"
	"path/filepath"
	"sort"
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

var kinds = Kinds{
	Whitelist: []string{"*.go", "*.js", "*.sh", "Dockerfile", "*.yaml"},
	Blacklist: []string{"*.md", "*.txt", "*.json", "*.png"},
}

func find(t *testing.T, files map[string]string) ([]Annotation, []Finding) {
	t.Helper()
	dir := write(t, files)
	var names []string
	for f := range files {
		names = append(names, f)
	}
	found, unknown, err := Find(dir, names, kinds)
	if err != nil {
		t.Fatal(err)
	}
	return found, unknown
}

func filesOf(found []Annotation) string {
	var got []string
	for _, a := range found {
		got = append(got, a.File)
	}
	sort.Strings(got)
	return strings.Join(got, " ")
}

// @verifies SR-24
func TestOnlyCodeIsRead(t *testing.T) {
	found, unknown := find(t, map[string]string{
		"a.go":                  "// @implements SR-01\nfunc a() {}\n",
		"b.js":                  "/** @verifies SR-01 */\n",
		"Dockerfile":            "# @implements SR-01\nFROM scratch\n",
		"run.sh":                "  # @implements SR-01\n",
		"config.yaml":           "# @implements SR-01\nkey: value\n",
		"doc.md":                "// @implements SR-99\n",
		"datei.txt":             "Find(root, files):\n    # @implements SR-99\n",
		"schemas/x.schema.json": `{"$comment": "@implements SR-99"}`,
		"data.yaml":             "schema: schemas/x.schema.json\n# @implements SR-99\n",
		"picture.png":           "# @implements SR-99",
	})
	if got := filesOf(found); got != "Dockerfile a.go b.js config.yaml run.sh" {
		t.Errorf("want the five files of code, got %q", got)
	}
	if len(unknown) != 0 {
		t.Errorf("want no unknown kind, got %q", messages(unknown))
	}
}

// @verifies SR-27
func TestAKindOnNeitherList(t *testing.T) {
	found, unknown := find(t, map[string]string{
		"query.sql": "-- @implements SR-99\n",
		"a.go":      "// @implements SR-01\n",
	})
	got := messages(unknown)
	if len(got) != 1 || got[0] != "query.sql: its kind is set neither as code nor as no code (tools/annotation/kinds.yaml)" {
		t.Errorf("want query.sql reported, got %q", got)
	}
	if filesOf(found) != "a.go" {
		t.Errorf("want only a.go read, got %q", filesOf(found))
	}
}

// @verifies SR-24
func TestOnlyFirstInALineOfComment(t *testing.T) {
	//	built with \n, not as a string over lines: a line of a raw string that begins with a mark of
	//	comment would be read as an annotation of this file itself
	found, _ := find(t, map[string]string{"a_test.go": "package a\n\n" +
		"const s = \"// @implements SR-99\"\n" +
		"var list = []string{\n\t\"// @implements SR-98\",\n}\n" +
		"func f() {} // @implements SR-97\n\n" +
		"// @verifies SR-01\n" +
		"func TestA() {}\n"})
	if len(found) != 1 || found[0].IDs[0] != "SR-01" || found[0].Line != 9 {
		t.Errorf("want only the comment on line 9, got %+v", found)
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
func TestSeveralIdsInOneAnnotation(t *testing.T) {
	found, _ := find(t, map[string]string{"a.go": "// @implements SR-01, SR-02,SR-03\n"})
	if len(found) != 1 || strings.Join(found[0].IDs, " ") != "SR-01 SR-02 SR-03" {
		t.Errorf("got %+v", found)
	}
}

func writeFile(dir, path, content string) error {
	os.MkdirAll(filepath.Join(dir, filepath.Dir(path)), 0o755)
	return os.WriteFile(filepath.Join(dir, path), []byte(content), 0o644)
}
