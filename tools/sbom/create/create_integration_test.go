//go:build integration

package create

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// The module of the tools is the subsystem here: its commands link yaml.v3, jsonschema and x/text.
const toolsRoot = "../.."

// @verifies SR_36@2.1.1
func TestGoListOfTheTools(t *testing.T) {
	list, err := GoList(toolsRoot)
	if err != nil {
		t.Fatal(err)
	}
	modules := map[string]Package{}
	for _, p := range GoModules(list, LicenceIn) {
		modules[p.Name] = p
	}
	if p := modules["gopkg.in/yaml.v3"]; p.PURL != "pkg:golang/gopkg.in/yaml.v3@v3.0.1" || p.Licence != "Apache-2.0 AND MIT" {
		t.Errorf("yaml.v3: %+v", p)
	}
	if p := modules["golang.org/x/text"]; p.Licence != "BSD-3-Clause" {
		t.Errorf("x/text: %+v", p)
	}
	if _, ok := modules["github.com/treporting/specs-and-deploy/tools"]; ok {
		t.Error("the module of the subsystem is listed")
	}
}

// @verifies SR_44@1.1.1
func TestGoEnvNamesTheVersionOfGo(t *testing.T) {
	version, err := GoEnv(toolsRoot, "GOVERSION")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(Stdlib(version, "").PURL, "pkg:golang/stdlib@go1.") {
		t.Errorf("version %q", version)
	}
}

// @verifies SR_39@1.2.1
// @verifies SR_40@2.2.1
func TestCommandWritesSbomMDInTheRootAndKeepsTheHistory(t *testing.T) {
	bin := filepath.Join(t.TempDir(), "sbom")
	if out, err := exec.Command("go", "build", "-o", bin, "../cmd/sbom").CombinedOutput(); err != nil {
		t.Fatalf("go build: %v\n%s", err, out)
	}
	repo := filepath.Join(t.TempDir(), "viewer-trace")
	os.MkdirAll(repo, 0o755)
	os.WriteFile(filepath.Join(repo, "LICENSE"), []byte(apacheText), 0o644)
	os.WriteFile(filepath.Join(repo, "package-lock.json"), []byte(lock), 0o644)

	for _, label := range []string{"1.0.0", "1.1.0"} {
		if out, err := exec.Command(bin, "-root", repo, "-label", label).CombinedOutput(); err != nil {
			t.Fatalf("sbom: %v\n%s", err, out)
		}
	}
	md, err := os.ReadFile(filepath.Join(repo, "sbom.md"))
	if err != nil {
		t.Fatal(err)
	}
	rows := History(md)
	if len(rows) != 2 || !strings.Contains(rows[0], " | 1.0.0 | - | sbom ") || !strings.Contains(rows[1], " | 1.1.0 | - | sbom ") {
		t.Errorf("rows %q", rows)
	}
	if !strings.Contains(string(md), "| pkg:github/treporting/viewer-trace@1.1.0 | viewer-trace | 1.1.0 | Apache-2.0 |") {
		t.Errorf("own package missing:\n%s", md)
	}
}
