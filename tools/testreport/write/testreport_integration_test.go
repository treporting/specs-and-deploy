//go:build integration

package write

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// testreportIn builds testreport and runs it with args, for the test level r5tc of -dir tools, in a
// repo of git, with junit-go.xml and junit-vitest.xml given from outside the repo; and answers the
// repo and the names in tools/testreports/.
func testreportIn(t *testing.T, args ...string) (string, []string) {
	t.Helper()
	bin := filepath.Join(t.TempDir(), "testreport")
	if out, err := exec.Command("go", "build", "-o", bin, "../cmd/testreport").CombinedOutput(); err != nil {
		t.Fatalf("go build: %v\n%s", err, out)
	}
	repo := t.TempDir()
	for name, content := range map[string]string{
		"documents.yaml":                "documents:\n",
		"tools/testcases/software.yaml": "testcases:\n",
	} {
		if err := os.MkdirAll(filepath.Join(repo, filepath.Dir(name)), 0o750); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(repo, name), []byte(content), 0o600); err != nil {
			t.Fatal(err)
		}
	}
	for _, git := range [][]string{{"init", "-q"}, {"add", "."}, {"-c", "user.name=t", "-c", "user.email=t@t", "commit", "-q", "-m", "a"}} {
		if out, err := exec.Command("git", append([]string{"-C", repo}, git...)...).CombinedOutput(); err != nil {
			t.Fatalf("git %v: %v\n%s", git, err, out)
		}
	}
	runs := t.TempDir()
	goXML, vitestXML := filepath.Join(runs, "junit-go.xml"), filepath.Join(runs, "junit-vitest.xml")
	if err := os.WriteFile(goXML, []byte(goJUnit), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(vitestXML, []byte(vitestJUnit), 0o600); err != nil {
		t.Fatal(err)
	}
	all := append([]string{"-root", repo, "-dir", "tools", "-level", "r5tc",
		"-junit", goXML, "-made-by", "gotestsum --junitfile junit-go.xml -- ./...",
		"-junit", vitestXML, "-made-by", "vitest run --reporter=junit"}, args...)
	if out, err := exec.Command(bin, all...).CombinedOutput(); err != nil {
		t.Fatalf("testreport: %v\n%s", err, out)
	}
	entries, err := os.ReadDir(filepath.Join(repo, "tools", "testreports"))
	if err != nil {
		t.Fatal(err)
	}
	var names []string
	for _, e := range entries {
		names = append(names, e.Name())
	}
	return repo, names
}

// @implements TC_139@2.2.1
func TestCommandWithoutStoreCopiesNothing(t *testing.T) {
	_, names := testreportIn(t)
	if len(names) != 1 || names[0] != "software.md" {
		t.Errorf("tools/testreports/ holds %v", names)
	}
}

// @implements TC_154@1.1.1, TC_155@1.1.1
func TestCommandWithStoreCopiesEachJUnit(t *testing.T) {
	repo, names := testreportIn(t, "-store")
	if len(names) != 4 {
		t.Fatalf("tools/testreports/ holds %v", names)
	}
	// <ts>_r5tc-junit-go.xml, <ts>_r5tc-junit-vitest.xml, <ts>_r5tc.yaml, software.md
	ts, ok := strings.CutSuffix(names[0], "_r5tc-junit-go.xml")
	if !ok || len(ts) != len("20261105T100000") || names[1] != ts+"_r5tc-junit-vitest.xml" || names[2] != ts+"_r5tc.yaml" || names[3] != "software.md" {
		t.Fatalf("tools/testreports/ holds %v", names)
	}
	copied, err := os.ReadFile(filepath.Join(repo, "tools", "testreports", names[0]))
	if err != nil || string(copied) != goJUnit {
		t.Errorf("the copy of junit-go.xml: %v", err)
	}
	md, err := os.ReadFile(filepath.Join(repo, "tools", "testreports", "software.md"))
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{
		"| junit-go.xml | `gotestsum --junitfile junit-go.xml -- ./...` | ./tools/testreports/" + names[0] + " |",
		"| junit-vitest.xml | `vitest run --reporter=junit` | ./tools/testreports/" + names[1] + " |",
	} {
		if !strings.Contains(string(md), want) {
			t.Errorf("Original data has no %q:\n%s", want, md)
		}
	}
	record, err := os.ReadFile(filepath.Join(repo, "tools", "testreports", names[2]))
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{
		"document: tools/testreports/software.md\n",
		"- file: " + names[0] + "\n    made_by: gotestsum --junitfile junit-go.xml -- ./...\n",
		"- file: " + names[1] + "\n    made_by: vitest run --reporter=junit\n",
	} {
		if !strings.Contains(string(record), want) {
			t.Errorf("%s has no %q:\n%s", names[2], want, record)
		}
	}
}
