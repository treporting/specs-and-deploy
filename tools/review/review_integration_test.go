//go:build integration

package review

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// gitIn runs git in dir and answers its output.
func gitIn(t *testing.T, dir string, args ...string) string {
	t.Helper()
	cmd := exec.Command("git", append([]string{"-C", dir}, args...)...)
	cmd.Env = append(os.Environ(), "GIT_AUTHOR_NAME=Mathias Fischer", "GIT_AUTHOR_EMAIL=m@example.org",
		"GIT_COMMITTER_NAME=Mathias Fischer", "GIT_COMMITTER_EMAIL=m@example.org")
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("git %v: %v\n%s", args, err, out)
	}
	return strings.TrimSpace(string(out))
}

func write(t *testing.T, root, path, content string) {
	t.Helper()
	os.MkdirAll(filepath.Join(root, filepath.Dir(path)), 0o755)
	if err := os.WriteFile(filepath.Join(root, path), []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
}

// The check on a real history: the requirements are written by Claude Code as author, then another
// requirement of the same file is changed by the reviewer alone - which must not make the reviewer
// an author of the first.
//
// @verifies TOOL-03
func TestTheCheckOnAHistoryOfGit(t *testing.T) {

	root := t.TempDir()
	gitIn(t, root, "init", "-q")
	write(t, root, "reviews/criteria-tool.yaml", criteria)
	write(t, root, "requirements/tool-requirements.yaml", requirements+`- object_key: TOOL-04
  object_text: first
`)
	gitIn(t, root, "add", "-A")
	gitIn(t, root, "commit", "-q", "--author=Claude Opus 5.5 <noreply@anthropic.com>", "-m", "requirements")
	first := gitIn(t, root, "rev-parse", "--short", "HEAD")

	write(t, root, "requirements/tool-requirements.yaml", requirements+`- object_key: TOOL-04
  object_text: changed by the reviewer alone
`)
	gitIn(t, root, "commit", "-q", "-am", "TOOL-04 changed")
	second := gitIn(t, root, "rev-parse", "--short", "HEAD")

	write(t, root, "reviews/r.yaml", `level: review
created_by: Mathias Fischer
created_at: '2026-10-25T13:30:00+01:00'
criteria: reviews/criteria-tool.yaml
criteria_revision: '`+first+`'
results:
- object_key: RR-01
  target_ids: [TOOL-03]
  revision: '`+second+`'
  criteria: {RC-01: passed, RC-02: passed}
  verdict: passed
- object_key: RR-02
  target_ids: [TOOL-04]
  revision: '`+second+`'
  criteria: {RC-01: passed, RC-02: passed}
  verdict: passed
`)

	findings, err := Check(root, "reviews/r.yaml", ExecGit{Root: root})
	if err != nil {
		t.Fatal(err)
	}
	if len(findings) != 1 || findings[0].Result != "RR-02" ||
		!strings.Contains(findings[0].Message, "is an author of TOOL-04") {
		t.Errorf("want one finding, that the reviewer is an author of TOOL-04; got %v", findings)
	}
}
