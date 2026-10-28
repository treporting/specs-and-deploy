package check

import (
	"os"
	"reflect"
	"testing"
)

// files holds the files of requirements in memory.
type files map[string]string

func (f files) read(file string) ([]byte, error) {
	if c, ok := f[file]; ok {
		return []byte(c), nil
	}
	return nil, os.ErrNotExist
}

func messages(t *testing.T, f files) ([]string, int) {
	t.Helper()
	findings, approved, err := Check(f.read)
	if err != nil {
		t.Fatal(err)
	}
	var out []string
	for _, finding := range findings {
		out = append(out, finding.String())
	}
	return out, approved
}

const tools = `schema: schemas/tool-requirements.schema.json
level: tool
requirements:
- object_key: TOOL-01
  status: approved
- object_key: TOOL-02
  status: draft
`

const rules = `schema: schemas/rules.schema.json
level: rules
rules:
- object_key: PROC-13
  status: approved
`

// @verifies SR-28
func TestApprovedUplinksPass(t *testing.T) {
	got, approved := messages(t, files{
		"requirements/tool-requirements.yaml": tools,
		"rules.yaml":                          rules,
		"tools/software-requirements.yaml": `requirements:
- object_key: SR-01
  source_ids: [TOOL-01@1.1.1, PROC-13@2.1.1]
  status: approved
`,
	})
	if len(got) != 0 {
		t.Errorf("findings %q, want none", got)
	}
	if approved != 3 {
		t.Errorf("approved %d, want 3: TOOL-01, PROC-13, SR-01", approved)
	}
}

// @verifies SR-28
func TestUplinkNotApprovedIsAFinding(t *testing.T) {
	got, _ := messages(t, files{
		"requirements/tool-requirements.yaml": tools,
		"tools/software-requirements.yaml": `requirements:
- object_key: SR-01
  source_ids: [TOOL-02@1.1.1, TOOL-07@1.1.1, NOPE-01@1.1.1]
  status: approved
`,
	})
	want := []string{
		"SR-01 names TOOL-02@1.1.1: draft, not approved",
		"SR-01 names TOOL-07@1.1.1: not in requirements/tool-requirements.yaml",
		"SR-01 names NOPE-01@1.1.1: no file of requirements holds such an id",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("findings %q, want %q", got, want)
	}
}

// A requirement that is not approved may name what it likes: only the state of approved ones is
// checked.
//
// @verifies SR-28
func TestDraftIsNotChecked(t *testing.T) {
	got, approved := messages(t, files{
		"requirements/tool-requirements.yaml": tools,
		"tools/software-requirements.yaml": `requirements:
- object_key: SR-01
  source_ids: [TOOL-02@1.1.1]
  status: review
`,
	})
	if len(got) != 0 || approved != 1 {
		t.Errorf("findings %q, approved %d; want none and 1", got, approved)
	}
}

// @verifies SR-28
func TestFileNotThere(t *testing.T) {
	got, _ := messages(t, files{
		"tools/software-requirements.yaml": `requirements:
- object_key: SR-01
  source_ids: [TOOL-01@1.1.1]
  status: approved
`,
	})
	want := []string{"SR-01 names TOOL-01@1.1.1: requirements/tool-requirements.yaml cannot be read"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("findings %q, want %q", got, want)
	}
}
