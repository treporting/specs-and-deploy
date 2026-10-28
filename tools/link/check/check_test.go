package check

import (
	"os"
	"reflect"
	"strings"
	"testing"

	annotation "github.com/treporting/specs-and-deploy/tools/annotation/check"
)

// files holds the files of requirements in memory.
type files map[string]string

func (f files) read(file string) ([]byte, error) {
	if c, ok := f[file]; ok {
		return []byte(c), nil
	}
	return nil, os.ErrNotExist
}

// links makes the link check on the files.
func (f files) links() ([]Finding, int, error) {
	return Links(f.read, f.list)
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

func lines(findings []Finding) []string {
	var out []string
	for _, f := range findings {
		out = append(out, f.String())
	}
	return out
}

// TOOL_09 has moved on in x, TOOL_01 in y and z only.
const tools = `schema: schemas/tool-requirements.schema.json
level: tool
requirements:
- object_key: TOOL_01
  object_version: '1.2.3'
- object_key: TOOL_09
  object_version: '4.1.1'
`

// @verifies SR_32@1.1.1
func TestOnlyAnotherXIsReported(t *testing.T) {
	got, compared, err := files{
		"requirements/tool-requirements.yaml": tools,
		"tools/software-requirements.yaml": `requirements:
- object_key: SR_01
  source_ids: [TOOL_01@1.1.1, TOOL_09@3.1.1]
`,
	}.links()
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"SR_01 source_ids names TOOL_09@3.1.1: its object has 4.1.1"}
	if !reflect.DeepEqual(lines(got), want) {
		t.Errorf("findings %q, want %q", lines(got), want)
	}
	if compared != 2 {
		t.Errorf("compared %d, want 2", compared)
	}
}

// @verifies SR_32@1.1.1
func TestTheAllocationAndControlsAreCompared(t *testing.T) {
	got, compared, err := files{
		"requirements/system-requirements.yaml": `requirements:
- object_key: SYS_09
  object_version: '4.1.2'
`,
		"requirements/user-needs.yaml": `requirements:
- object_key: UN_PSF_01
  object_version: '2.1.1'
`,
		"architecture/system-architecture.yaml": `elements:
- object_key: ARC_01
  allocated: [SYS_09@3.1.2]
`,
		"safety/hazards.yaml": `controls:
- object_key: HAZC_01
  controls: [UN_PSF_01@1.1.1]
`,
	}.links()
	if err != nil {
		t.Fatal(err)
	}
	want := []string{
		"ARC_01 allocated names SYS_09@3.1.2: its object has 4.1.2",
		"HAZC_01 controls names UN_PSF_01@1.1.1: its object has 2.1.1",
	}
	if !reflect.DeepEqual(lines(got), want) {
		t.Errorf("findings %q, want %q", lines(got), want)
	}
	if compared != 2 {
		t.Errorf("compared %d, want 2", compared)
	}
}

// A link to an object no file holds is not compared; uplinkcheck reports it.
//
// @verifies SR_32@1.1.1
func TestAMissingObjectIsNotCompared(t *testing.T) {
	got, compared, err := files{
		"requirements/tool-requirements.yaml": tools,
		"tools/software-requirements.yaml": `requirements:
- object_key: SR_01
  source_ids: [TOOL_99@1.1.1, TOOL_09]
`,
	}.links()
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 0 || compared != 0 {
		t.Errorf("findings %q, compared %d, want none", lines(got), compared)
	}
}

// @verifies SR_33@1.1.1
func TestAnAnnotationWithAnotherXIsReported(t *testing.T) {
	got, compared := Annotations(files{"requirements/tool-requirements.yaml": tools}.read, []annotation.Annotation{
		{File: "a.go", Line: 3, Kind: "implements", IDs: []string{"TOOL_01", "TOOL_09"}, Versions: []string{"1.1.1", "3.1.1"}},
		{File: "a_test.go", Line: 7, Kind: "verifies", IDs: []string{"TOOL_09"}, Versions: []string{"4.1.1"}},
		{File: "b.go", Line: 1, Kind: "implements", IDs: []string{"TOOL_09"}, Versions: []string{""}},
	})
	want := []string{"a.go:3 @implements names TOOL_09@3.1.1: its object has 4.1.1"}
	if !reflect.DeepEqual(lines(got), want) {
		t.Errorf("findings %q, want %q", lines(got), want)
	}
	if compared != 3 {
		t.Errorf("compared %d, want 3", compared)
	}
}

// A link to a regulatory requirement is compared with the version in the file of its module.
//
// @verifies SR_32@1.1.1
func TestALinkToAModuleIsCompared(t *testing.T) {
	got, compared, err := files{
		"regulatory/DO178C.yaml": `requirements:
- object_key: REG_DO178C_02
  object_version: '2.1.1'
`,
		"requirements/system-requirements.yaml": `requirements:
- object_key: SYS_09
  source_ids: [REG_DO178C_02@1.1.1]
`,
	}.links()
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"SYS_09 source_ids names REG_DO178C_02@1.1.1: its object has 2.1.1"}
	if !reflect.DeepEqual(lines(got), want) {
		t.Errorf("findings %q, want %q", lines(got), want)
	}
	if compared != 1 {
		t.Errorf("compared %d, want 1", compared)
	}
}
