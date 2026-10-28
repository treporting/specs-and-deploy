package check

import (
	"os"
	"reflect"
	"strings"
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

// list answers the names of the files directly in dir.
func (f files) list(dir string) ([]string, error) {
	var names []string
	for file := range f {
		if name, ok := strings.CutPrefix(file, dir+"/"); ok && !strings.Contains(name, "/") {
			names = append(names, name)
		}
	}
	return names, nil
}

func messages(t *testing.T, f files) ([]string, int) {
	t.Helper()
	findings, approved, err := Check(f.read, f.list)
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
- object_key: TOOL_01
  status: approved
- object_key: TOOL_02
  status: draft
`

const rules = `schema: schemas/rules.schema.json
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
		"tools/software-requirements.yaml": `requirements:
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
		"tools/software-requirements.yaml": `requirements:
- object_key: SR_01
  source_ids: [TOOL_02@1.1.1, TOOL_07@1.1.1, NOPE-01@1.1.1]
  status: approved
`,
	})
	want := []string{
		"SR_01 names TOOL_02@1.1.1: draft, not approved",
		"SR_01 names TOOL_07@1.1.1: not in requirements/tool-requirements.yaml",
		"SR_01 names NOPE-01@1.1.1: no file of requirements holds such an id",
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
		"tools/software-requirements.yaml": `requirements:
- object_key: SR_01
  source_ids: [TOOL_02@1.1.1]
  status: review
`,
	})
	if len(got) != 0 || approved != 1 {
		t.Errorf("findings %q, approved %d; want none and 1", got, approved)
	}
}

// @verifies SR_28@1.1.1
func TestFileNotThere(t *testing.T) {
	got, _ := messages(t, files{
		"tools/software-requirements.yaml": `requirements:
- object_key: SR_01
  source_ids: [TOOL_01@1.1.1]
  status: approved
`,
	})
	want := []string{"SR_01 names TOOL_01@1.1.1: requirements/tool-requirements.yaml cannot be read"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("findings %q, want %q", got, want)
	}
}

// A regulatory requirement is found in the file of its module, which regulatory/ holds.
//
// @verifies SR_28@1.1.1
func TestARegulatoryRequirementIsFoundInItsModule(t *testing.T) {
	got, approved := messages(t, files{
		"regulatory/ISO26262-8.yaml": `requirements:
- object_key: REG_ISO26262-8_01
  status: approved
`,
		"tools/software-requirements.yaml": `requirements:
- object_key: SR_01
  source_ids: [REG_ISO26262-8_01@1.1.1, REG_EU2016-679_01@1.1.1]
  status: approved
`,
	})
	want := []string{"SR_01 names REG_EU2016-679_01@1.1.1: regulatory/EU2016-679.yaml cannot be read"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("findings %q, want %q", got, want)
	}
	if approved != 2 {
		t.Errorf("approved %d, want 2", approved)
	}
}
