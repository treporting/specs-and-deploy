//go:build integration

package approval

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/treporting/specs-and-deploy/tools/review"
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
	if err := os.MkdirAll(filepath.Join(root, filepath.Dir(path)), 0o750); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, path), []byte(content), 0o600); err != nil {
		t.Fatal(err)
	}
}

const schema = `{"properties": {"requirements": {"items": {"attributes_for_version_x": ["object_text"]}}}}`

// tool writes a file of tool requirements: each key with its text and status.
func tool(keys, texts []string, status string) string {
	var b strings.Builder
	b.WriteString("schema: schemas/tool.schema.json\ntype: requirements\nlevel: tool\nrequirements:\n")
	for i, key := range keys {
		b.WriteString("- object_key: " + key + "\n  object_version: '1.1.1'\n  object_text: " + texts[i] + "\n  status: " + status + "\n")
	}
	return b.String()
}

func documentsYAML(hash, status string) string {
	return "type: documents\ndocuments:\n- name: a.md\n  object_version: '1.1.1'\n  object_hash: '" + hash +
		"'\n  status: " + status + "\n  verification_method: [review]\n"
}

func reviewResult(at, target, revision, verdict string) string {
	return "created_at: '" + at + "'\nresults:\n- object_key: RR_01\n  target_ids: [" + target +
		"]\n  revision: '" + revision + "'\n  verdict: " + verdict + "\n"
}

// A history of three commits: the objects as drafts; approved, with their reviews of the first
// commit; then TOOL_03 and a.md changed while approved.
func history(t *testing.T) (root, approved string) {

	root = t.TempDir()
	gitIn(t, root, "init", "-q")
	keys := []string{"TOOL_01", "TOOL_02", "TOOL_03", "TOOL_04"}
	texts := []string{"one", "two", "three", "four"}
	write(t, root, "schemas/tool.schema.json", schema)
	write(t, root, "reviews/criteria-tool.yaml", "level: tool\n")
	write(t, root, "reviews/criteria-documents.yaml", "level: document\n")
	write(t, root, "requirements/tool.yaml", tool(keys, texts, "draft"))
	write(t, root, "requirements/other.yaml", tool([]string{"TOOL_09"}, []string{"nine"}, "draft"))
	write(t, root, "a.md", "# A\n")
	write(t, root, "documents.yaml", documentsYAML("", "draft"))
	gitIn(t, root, "add", "-A")
	gitIn(t, root, "commit", "-q", "-m", "drafts")
	drafts := gitIn(t, root, "rev-parse", "--short", "HEAD")

	write(t, root, "requirements/tool.yaml", tool(keys, texts, "approved"))
	write(t, root, "requirements/other.yaml", tool([]string{"TOOL_09"}, []string{"nine"}, "approved"))
	write(t, root, "documents.yaml", documentsYAML(gitIn(t, root, "hash-object", "a.md")[:7], "approved"))
	write(t, root, "reviews/results/1.yaml", reviewResult("2026-11-02T10:00:00+01:00", "TOOL_01", drafts, "passed"))
	write(t, root, "reviews/results/2.yaml", reviewResult("2026-11-02T10:00:00+01:00", "TOOL_02", drafts, "passed"))
	write(t, root, "reviews/results/3.yaml", reviewResult("2026-11-02T11:00:00+01:00", "TOOL_02", drafts, "failed"))
	write(t, root, "reviews/results/4.yaml", reviewResult("2026-11-02T10:00:00+01:00", "TOOL_03", drafts, "passed"))
	write(t, root, "reviews/results/5.yaml", reviewResult("2026-11-02T10:00:00+01:00", "a.md", drafts, "passed"))
	gitIn(t, root, "add", "-A")
	gitIn(t, root, "commit", "-q", "-m", "approved")
	approved = gitIn(t, root, "rev-parse", "--short", "HEAD")

	write(t, root, "requirements/tool.yaml", tool(keys, []string{"one", "two", "three, changed", "four"}, "approved"))
	write(t, root, "a.md", "# A, changed\n")
	gitIn(t, root, "commit", "-q", "-am", "changed while approved")
	return root, approved
}

func messages(t *testing.T, root, base string) string {
	t.Helper()
	index, err := review.IndexAt(root)
	if err != nil {
		t.Fatal(err)
	}
	findings, err := Check(root, base, review.ExecGit{Root: root}, index)
	if err != nil {
		t.Fatal(err)
	}
	var lines []string
	for _, f := range findings {
		lines = append(lines, f.String())
	}
	return strings.Join(lines, "\n")
}

// Every approved object, and the findings each error gives.
//
// @verifies SR_75@1.1.1, SR_76@1.1.1, SR_77@1.1.1, SR_78@1.1.1, SR_79@1.1.1, SR_80@1.1.1
func TestEveryApprovedObject(t *testing.T) {

	root, _ := history(t)
	got := messages(t, root, "")

	for _, want := range []string{
		"requirements/tool.yaml: TOOL_02: the last review result that names it, reviews/results/3.yaml, has the verdict failed",
		"requirements/tool.yaml: TOOL_03: its content for x is not the one reviews/results/4.yaml reviewed at",
		"requirements/tool.yaml: TOOL_03: its attributes for x are not those of its approval with x 1",
		"requirements/tool.yaml: TOOL_04: approved, and no review result names it",
		"requirements/other.yaml: TOOL_09: approved, and no review result names it",
		"a.md: a.md: its content for x is not the one reviews/results/5.yaml reviewed at",
		"documents.yaml: a.md: its object_hash",
	} {
		if !strings.Contains(got, want) {
			t.Errorf("missing %q in\n%s", want, got)
		}
	}
	if strings.Contains(got, "TOOL_01") {
		t.Errorf("TOOL_01 has a passed review of its content, and is reported:\n%s", got)
	}
	if n := strings.Count(got, "\n") + 1; n != 7 {
		t.Errorf("%d findings, want 7:\n%s", n, got)
	}
}

// With a base commit, only the objects in the files changed against it.
//
// @verifies SR_80@1.1.1
func TestTheFilesChangedAgainstTheBase(t *testing.T) {

	root, approved := history(t)
	got := messages(t, root, approved)

	if strings.Contains(got, "TOOL_09") {
		t.Errorf("TOOL_09 is in a file the base commit and HEAD share, and is reported:\n%s", got)
	}
	for _, want := range []string{"TOOL_03", "TOOL_04", "a.md"} {
		if !strings.Contains(got, want) {
			t.Errorf("missing %s in\n%s", want, got)
		}
	}
}

// A passed result that is still open is a proposal, not an approval.
//
// @verifies SR_85@1.1.1
func TestAnOpenResultIsNoApproval(t *testing.T) {

	root, approved := history(t)
	write(t, root, "reviews/results/6.yaml", "status: open\n"+reviewResult("2026-11-02T12:00:00+01:00", "TOOL_01", approved, "passed"))
	gitIn(t, root, "add", "-A")
	gitIn(t, root, "commit", "-q", "-m", "an open review of TOOL_01")
	got := messages(t, root, "")

	want := "requirements/tool.yaml: TOOL_01: the last review result that names it, reviews/results/6.yaml, is open"
	if !strings.Contains(got, want) {
		t.Errorf("missing %q in\n%s", want, got)
	}
}
