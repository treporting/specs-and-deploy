package skeleton

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

// fakeGit is a history of two commits kept in memory: c1 changed the criteria, c2 is HEAD.
type fakeGit struct {
	files map[string]string // path -> content, the same in both commits
}

func (f fakeGit) Show(rev, path string) ([]byte, error) {
	if c, ok := f.files[path]; ok {
		return []byte(c), nil
	}
	return nil, os.ErrNotExist
}

func (f fakeGit) FileCommits(rev, path string) ([]string, error) {
	if strings.HasPrefix(path, "reviews/criteria-") {
		return []string{"c1full"}, nil
	}
	return []string{"c2full"}, nil
}

func (f fakeGit) Short(rev string) (string, error) {
	if rev == "HEAD" {
		return "c2", nil
	}
	return strings.TrimSuffix(rev, "full"), nil
}

func history() fakeGit {
	return fakeGit{files: map[string]string{
		"reviews/criteria-tool.yaml":      "level: tool\ncriteria:\n- object_key: RC-01\n- object_key: RC-02\n",
		"reviews/criteria-documents.yaml": "level: document\ncriteria:\n- object_key: RC-01\n",
		"requirements/tool-requirements.yaml": `level: tool
requirements:
- object_key: TOOL-03
  object_text: check every review result
- object_key: TOOL-04
  object_text: create the skeleton
`,
		"harness.md": "# Harness\n",
	}}
}

var now = time.Date(2026, 10, 26, 14, 30, 5, 0, time.FixedZone("CET", 3600))

// makeSkeleton writes a skeleton into a temporary root and answers its path and content.
func makeSkeleton(t *testing.T, targets ...string) (string, string, error) {
	t.Helper()
	root := t.TempDir()
	os.MkdirAll(filepath.Join(root, "reviews"), 0o755)
	path, err := Make(root, Request{Reviewer: "Mathias Fischer", Scope: "tool-03-04", Targets: targets, Now: now}, history())
	if err != nil {
		entries, _ := os.ReadDir(filepath.Join(root, "reviews"))
		if len(entries) != 0 {
			t.Errorf("want nothing written, got %d files", len(entries))
		}
		return "", "", err
	}
	data, _ := os.ReadFile(filepath.Join(root, path))
	return path, string(data), nil
}

func contains(t *testing.T, content string, want ...string) {
	t.Helper()
	for _, w := range want {
		if !strings.Contains(content, w) {
			t.Errorf("no %q in\n%s", w, content)
		}
	}
}

// @verifies SR-11@1.1.1
func TestTheNameTheReviewerAndTheTime(t *testing.T) {
	path, content, err := makeSkeleton(t, "TOOL-03")
	if err != nil {
		t.Fatal(err)
	}
	if path != filepath.Join("reviews", "20261026T143005_tool-03-04.yaml") {
		t.Errorf("path %s", path)
	}
	contains(t, content, "created_by: Mathias Fischer\n", "created_at: '2026-10-26T14:30:05+01:00'\n")
}

// @verifies SR-12@1.1.1
func TestTheCriteriaOfTheLevel(t *testing.T) {
	_, content, err := makeSkeleton(t, "TOOL-03")
	if err != nil {
		t.Fatal(err)
	}
	contains(t, content, "criteria: reviews/criteria-tool.yaml\n", "criteria_revision: 'c1'\n")

	_, content, err = makeSkeleton(t, "harness.md")
	if err != nil {
		t.Fatal(err)
	}
	contains(t, content, "criteria: reviews/criteria-documents.yaml\n")
}

// @verifies SR-13@1.1.1
func TestTheRevisionIsHead(t *testing.T) {
	_, content, err := makeSkeleton(t, "TOOL-03")
	if err != nil {
		t.Fatal(err)
	}
	contains(t, content, "  revision: 'c2'\n")
}

// @verifies SR-14@1.1.1
func TestEveryCriterionAndTheVerdictAreTodo(t *testing.T) {
	_, content, err := makeSkeleton(t, "TOOL-03", "TOOL-04")
	if err != nil {
		t.Fatal(err)
	}
	contains(t, content, "- object_key: RR-01\n  target_ids: [TOOL-03]\n", "- object_key: RR-02\n  target_ids: [TOOL-04]\n")
	if n := strings.Count(content, ": todo\n"); n != 6 {
		t.Errorf("want 6 todo, two criteria and a verdict for each of two; got %d", n)
	}
	if strings.Contains(content, "passed") || strings.Contains(content, "failed") {
		t.Errorf("a skeleton holds no judgement:\n%s", content)
	}
}

// @verifies SR-15@1.1.1
func TestNothingWhereATargetIsMissingOrTheLevelsDiffer(t *testing.T) {
	if _, _, err := makeSkeleton(t, "TOOL-99"); err == nil || !strings.Contains(err.Error(), "TOOL-99 is not in") {
		t.Errorf("want TOOL-99 not found, got %v", err)
	}
	if _, _, err := makeSkeleton(t, "nothing.md"); err == nil {
		t.Errorf("want nothing.md not found")
	}
	if _, _, err := makeSkeleton(t, "TOOL-03", "harness.md"); err == nil || !strings.Contains(err.Error(), "one review, one level") {
		t.Errorf("want two levels refused, got %v", err)
	}
}

// @verifies SR-17@1.1.1
func TestTheSkeletonNamesItsSchema(t *testing.T) {
	_, content, err := makeSkeleton(t, "TOOL-03")
	if err != nil {
		t.Fatal(err)
	}
	contains(t, content, "schema: schemas/review-result.schema.json\nlevel: review\n")
}
