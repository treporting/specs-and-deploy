package review

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// fakeGit is a history of one commit and its parent, kept in memory.
type fakeGit struct {
	files     map[string]map[string]string // commit -> path -> content
	parent    map[string]string
	changed   map[string][]string // path -> the commits that changed it, the newest first
	author    string
	coAuthors []string
}

func (f fakeGit) CommitExists(rev string) bool { _, ok := f.files[rev]; return ok }

func (f fakeGit) Show(rev, path string) ([]byte, error) {
	if c, ok := f.files[rev][path]; ok {
		return []byte(c), nil
	}
	return nil, os.ErrNotExist
}

func (f fakeGit) FileCommits(rev, path string) ([]string, error) {
	if commits, ok := f.changed[path]; ok {
		return commits, nil
	}
	return []string{rev}, nil
}

func (f fakeGit) Parent(commit string) (string, bool) {
	p, ok := f.parent[commit]
	return p, ok
}

func (f fakeGit) Authors(string) (string, []string, error) { return f.author, f.coAuthors, nil }

const criteria = `level: tool
criteria:
- object_key: RC-01
- object_key: RC-02
`

const requirements = `level: tool
requirements:
- object_key: TOOL-03
  object_text: check every review result
`

func history(author string, coAuthors ...string) fakeGit {
	return fakeGit{
		files: map[string]map[string]string{
			"c1": {"reviews/criteria-tool.yaml": criteria, "requirements/tool-requirements.yaml": requirements},
		},
		parent:    map[string]string{},
		author:    author,
		coAuthors: coAuthors,
	}
}

// check writes a review result into a temporary root and checks it.
func check(t *testing.T, result string, git fakeGit) []string {
	t.Helper()
	root := t.TempDir()
	os.MkdirAll(filepath.Join(root, "reviews"), 0o755)
	review := `level: review
created_by: Mathias Fischer
created_at: '2026-10-25T13:30:00+01:00'
criteria: reviews/criteria-tool.yaml
criteria_revision: 'c1'
results:
` + result
	os.WriteFile(filepath.Join(root, "reviews", "r.yaml"), []byte(review), 0o644)
	findings, err := Check(root, "reviews/r.yaml", git)
	if err != nil {
		t.Fatal(err)
	}
	var messages []string
	for _, f := range findings {
		messages = append(messages, f.Message)
	}
	return messages
}

func expect(t *testing.T, messages []string, want string) {
	t.Helper()
	for _, m := range messages {
		if strings.Contains(m, want) {
			return
		}
	}
	t.Errorf("no finding %q in %q", want, messages)
}

// @verifies SR-01, SR-02, SR-03, SR-04, SR-05
func TestAGoodResultHasNoFinding(t *testing.T) {
	got := check(t, `- object_key: RR-01
  target_ids: [TOOL-03]
  revision: 'c1'
  criteria: {RC-01: passed, RC-02: passed}
  verdict: passed
`, history("Claude Opus 5.5", "Claude Opus 5.5"))
	if len(got) != 0 {
		t.Errorf("want no finding, got %q", got)
	}
}

// @verifies SR-01
func TestEveryCriterionOnce(t *testing.T) {
	got := check(t, `- object_key: RR-01
  target_ids: [TOOL-03]
  revision: 'c1'
  criteria: {RC-01: passed, RC-09: passed}
  verdict: passed
`, history("Claude Opus 5.5", "Claude Opus 5.5"))
	expect(t, got, "criterion RC-02 is missing")
	expect(t, got, "criterion RC-09 is not in")
}

// @verifies SR-02
func TestTheVerdictFollowsTheCriteria(t *testing.T) {
	got := check(t, `- object_key: RR-01
  target_ids: [TOOL-03]
  revision: 'c1'
  criteria: {RC-01: passed, RC-02: {result: failed, note: unclear}}
  verdict: passed
- object_key: RR-02
  target_ids: [TOOL-03]
  revision: 'c1'
  criteria: {RC-01: passed, RC-02: passed}
  verdict: failed
`, history("Claude Opus 5.5", "Claude Opus 5.5"))
	expect(t, got, "the verdict is passed, but not every criterion passed")
	expect(t, got, "the verdict is failed, but every criterion passed")
}

// @verifies SR-03
func TestAFailedCriterionHasANote(t *testing.T) {
	got := check(t, `- object_key: RR-01
  target_ids: [TOOL-03]
  revision: 'c1'
  criteria: {RC-01: passed, RC-02: failed}
  verdict: failed
`, history("Claude Opus 5.5", "Claude Opus 5.5"))
	expect(t, got, "criterion RC-02 failed with no note")
}

// @verifies SR-04
func TestTheRevisionExists(t *testing.T) {
	got := check(t, `- object_key: RR-01
  target_ids: [TOOL-03]
  revision: 'nope'
  criteria: {RC-01: passed, RC-02: passed}
  verdict: passed
`, history("Claude Opus 5.5", "Claude Opus 5.5"))
	expect(t, got, "the revision nope is no commit")
}

// @verifies SR-05
func TestTheReviewerIsNotTheAuthor(t *testing.T) {
	result := `- object_key: RR-01
  target_ids: [TOOL-03]
  revision: 'c1'
  criteria: {RC-01: passed, RC-02: passed}
  verdict: passed
`
	//	written by the reviewer alone
	expect(t, check(t, result, history("Mathias Fischer")), "the reviewer Mathias Fischer is an author of TOOL-03")

	//	written by another, with the reviewer as co-author
	expect(t, check(t, result, history("Copilot", "Mathias Fischer")), "the reviewer Mathias Fischer is an author of TOOL-03")

	//	written by another, with no co-author: whatever name stands there
	if got := check(t, result, history("Copilot")); len(got) != 0 {
		t.Errorf("want no finding, got %q", got)
	}
}

// @verifies SR-05
func TestASoftwareRequirementOfTheToolsHasAnAuthor(t *testing.T) {
	git := history("Mathias Fischer")
	git.files["c1"]["tools/software-requirements.yaml"] = `level: r5
requirements:
- object_key: SR-01
  object_text: report a criterion that is missing
`
	got := check(t, `- object_key: RR-01
  target_ids: [SR-01]
  revision: 'c1'
  criteria: {RC-01: passed, RC-02: passed}
  verdict: passed
`, git)
	expect(t, got, "the reviewer Mathias Fischer is an author of SR-01")
}

// @verifies SR-06
func TestAValueOtherThanPassedOrFailed(t *testing.T) {
	got := check(t, `- object_key: RR-01
  target_ids: [TOOL-03]
  revision: 'c1'
  criteria: {RC-01: passed, RC-02: pased}
  verdict: passed
- object_key: RR-02
  target_ids: [TOOL-03]
  revision: 'c1'
  criteria: {RC-01: passed, RC-02: passed}
  verdict: pased
`, history("Claude Opus 5.5", "Claude Opus 5.5"))
	expect(t, got, "criterion RC-02 is neither passed nor failed")
	//	counted as not passed
	expect(t, got, "the verdict is passed, but not every criterion passed")
	expect(t, got, "the verdict is neither passed nor failed")
}

// @verifies SR-07
func TestTheCriteriaRevisionChangedTheCriteria(t *testing.T) {
	git := history("Claude Opus 5.5", "Claude Opus 5.5")
	git.changed = map[string][]string{"reviews/criteria-tool.yaml": {"c0"}}
	got := check(t, `- object_key: RR-01
  target_ids: [TOOL-03]
  revision: 'c1'
  criteria: {RC-01: passed, RC-02: passed}
  verdict: passed
`, git)
	expect(t, got, "the criteria revision c1 did not change reviews/criteria-tool.yaml")
}

// @verifies SR-08
func TestTheReviewerIsNotTheAuthorOfADocument(t *testing.T) {
	result := `- object_key: RR-01
  target_ids: [formats.md]
  revision: 'c1'
  criteria: {RC-01: passed, RC-02: passed}
  verdict: passed
`
	git := history("Mathias Fischer")
	git.files["c1"]["formats.md"] = "# formats\n"
	expect(t, check(t, result, git), "the reviewer Mathias Fischer is an author of formats.md")

	git = history("Claude Opus 5.5")
	git.files["c1"]["formats.md"] = "# formats\n"
	if got := check(t, result, git); len(got) != 0 {
		t.Errorf("want no finding, got %q", got)
	}
}
