package check

import (
	"os"
	"reflect"
	"sort"
	"testing"

	"github.com/treporting/specs-and-deploy/tools/review"
)

// files holds the files of requirements in memory.
type files map[string]string

func (f files) read(file string) ([]byte, error) {
	if c, ok := f[file]; ok {
		return []byte(c), nil
	}
	return nil, os.ErrNotExist
}

// index makes the index of the files, in the order of their names.
func (f files) index() *review.Index {
	var names []string
	for file := range f {
		names = append(names, file)
	}
	sort.Strings(names)
	return review.NewIndex(f.read, names)
}

func messages(t *testing.T, f files) ([]string, int) {
	t.Helper()
	findings, approved, err := Check(f.read, f.index())
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
type: requirements
level: tool
requirements:
- object_key: TOOL_01
  status: approved
- object_key: TOOL_02
  status: draft
`

const rules = `schema: schemas/rules.schema.json
type: rules
level: rules
rules:
- object_key: PROC_13
  status: approved
`

// @verifies SR_28@1.1.1
func TestApprovedUplinksPass(t *testing.T) {
	got, approved := messages(t, files{
		"requirements/tool-requirements.yaml": tools,
		"rules.yaml":                          rules,
		"tools/software-requirements.yaml": `type: requirements
requirements:
- object_key: SR_01
  source_ids: [TOOL_01@1.1.1, PROC_13@2.1.1]
  status: approved
`,
	})
	if len(got) != 0 {
		t.Errorf("findings %q, want none", got)
	}
	if approved != 3 {
		t.Errorf("approved %d, want 3: TOOL_01, PROC_13, SR_01", approved)
	}
}

// @verifies SR_28@1.1.1
func TestUplinkNotApprovedIsAFinding(t *testing.T) {
	got, _ := messages(t, files{
		"requirements/tool-requirements.yaml": tools,
		"tools/software-requirements.yaml": `type: requirements
requirements:
- object_key: SR_01
  source_ids: [TOOL_02@1.1.1, TOOL_07@1.1.1]
  status: approved
`,
	})
	want := []string{
		"SR_01 names TOOL_02@1.1.1: draft, not approved",
		"SR_01 names TOOL_07@1.1.1: no file of requirements holds such an id",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("findings %q, want %q", got, want)
	}
}

// A requirement that is not approved may name what it likes: only the state of approved ones is
// checked.
//
// @verifies SR_28@1.1.1
func TestDraftIsNotChecked(t *testing.T) {
	got, approved := messages(t, files{
		"requirements/tool-requirements.yaml": tools,
		"tools/software-requirements.yaml": `type: requirements
requirements:
- object_key: SR_01
  source_ids: [TOOL_02@1.1.1]
  status: review
`,
	})
	if len(got) != 0 || approved != 1 {
		t.Errorf("findings %q, approved %d; want none and 1", got, approved)
	}
}

// A file whose head names no type that holds requirements holds none, wherever it lies.
//
// @verifies SR_28@1.1.1
func TestAFileIsReadByItsType(t *testing.T) {
	got, approved := messages(t, files{
		"requirements/tool-requirements.yaml": `requirements:
- object_key: TOOL_01
  status: approved
`,
		"reviews/criteria-tool.yaml": `type: criteria
criteria:
- object_key: RC_01
`,
		"elsewhere/more.yaml": `type: requirements
requirements:
- object_key: TOOL_09
  status: approved
`,
		"tools/software-requirements.yaml": `type: requirements
requirements:
- object_key: SR_01
  source_ids: [TOOL_01@1.1.1, TOOL_09@1.1.1, RC_01@1.1.1]
  status: approved
`,
	})
	want := []string{
		"SR_01 names TOOL_01@1.1.1: no file of requirements holds such an id",
		"SR_01 names RC_01@1.1.1: no file of requirements holds such an id",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("findings %q, want %q", got, want)
	}
	if approved != 2 {
		t.Errorf("approved %d, want 2: TOOL_09, SR_01", approved)
	}
}

// A regulatory requirement is found in the file of its module.
//
// @verifies SR_28@1.1.1
func TestARegulatoryRequirementIsFoundInItsModule(t *testing.T) {
	got, approved := messages(t, files{
		"regulatory/ISO26262-8.yaml": `type: normative
requirements:
- object_key: REG_ISO26262-8_01
  status: approved
`,
		"tools/software-requirements.yaml": `type: requirements
requirements:
- object_key: SR_01
  source_ids: [REG_ISO26262-8_01@1.1.1, REG_EU2016-679_01@1.1.1]
  status: approved
`,
	})
	want := []string{"SR_01 names REG_EU2016-679_01@1.1.1: no file of requirements holds such an id"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("findings %q, want %q", got, want)
	}
	if approved != 2 {
		t.Errorf("approved %d, want 2", approved)
	}
}
