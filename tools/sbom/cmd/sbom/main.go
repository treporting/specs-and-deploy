// Command sbom creates the SBOM of a subsystem as sbom.md in the root of its repo (TOOL_21).
//
//	sbom -label x.y.z [-root dir] [-repo name] [-image name]...
//
// It runs in the repo checked out at the baseline the label names. It lists the package of the
// subsystem itself; the Go modules of its commands and the standard library of Go, where the root
// holds a go.mod; the packages of package-lock.json, where the root holds one; and the Alpine
// packages of each image named. It keeps the rows of the Change history of an sbom.md that is there,
// and adds one.
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime/debug"
	"strings"
	"time"

	"github.com/treporting/specs-and-deploy/tools/sbom/create"
)

// images collects the flag -image, which may be given more than once.
type images []string

func (i *images) String() string     { return strings.Join(*i, ",") }
func (i *images) Set(v string) error { *i = append(*i, v); return nil }

// @implements SR_39@1.2.1
func main() {

	root := flag.String("root", ".", "the root of the repo of the subsystem")
	label := flag.String("label", "", "the label of the baseline, such as 1.2.0")
	repo := flag.String("repo", "", "the name of the repo; the name of the directory of root where empty")
	var named images
	flag.Var(&named, "image", "an image of the subsystem, built at the baseline; may be given more than once")
	flag.Parse()

	if *label == "" {
		fail(fmt.Errorf("-label is missing"))
	}
	if *repo == "" {
		abs, err := filepath.Abs(*root)
		if err != nil {
			fail(err)
		}
		*repo = filepath.Base(abs)
	}

	packages := []create.Package{create.Own(*repo, *label, create.LicenceIn(*root))}

	if exists(filepath.Join(*root, "go.mod")) {
		list, err := create.GoList(*root)
		if err != nil {
			fail(err)
		}
		if list != "" {
			packages = append(packages, create.GoModules(list, create.LicenceIn)...)
			version, err := create.GoEnv(*root, "GOVERSION")
			if err != nil {
				fail(err)
			}
			goroot, err := create.GoEnv(*root, "GOROOT")
			if err != nil {
				fail(err)
			}
			licence := create.LicenceIn(goroot)
			if licence == "" {
				licence = create.LicenceIn(filepath.Dir(goroot)) // Homebrew keeps it beside GOROOT
			}
			packages = append(packages, create.Stdlib(version, licence))
		}
	}

	if lock, err := os.ReadFile(filepath.Join(*root, "package-lock.json")); err == nil {
		npm, err := create.NpmPackages(lock)
		if err != nil {
			fail(err)
		}
		packages = append(packages, npm...)
	}

	for _, image := range named {
		installed, err := create.ApkInstalled(image)
		if err != nil {
			fail(err)
		}
		packages = append(packages, create.ApkPackages(installed)...)
	}

	path := filepath.Join(*root, "sbom.md")
	old, _ := os.ReadFile(path) // an sbom.md that is not there has no rows to keep
	history := append(create.History(old), create.Row(time.Now().Format("2006-01-02"), *label, "sbom "+version()))

	if err := os.WriteFile(path, []byte(create.Markdown(create.SBOM{Repo: *repo, Packages: packages, History: history})), 0o600); err != nil {
		fail(err)
	}
	fmt.Printf("%s written: %d packages\n", path, len(packages))
}

// version answers the version of this tool: the commit of specs-and-deploy it was built from, as Go
// records it in the build, with -dirty where the working tree had changes; (devel) where Go recorded
// none, as with go run.
func version() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "(devel)"
	}
	var revision, modified string
	for _, s := range info.Settings {
		switch s.Key {
		case "vcs.revision":
			revision = s.Value
		case "vcs.modified":
			modified = s.Value
		}
	}
	if revision == "" {
		return info.Main.Version
	}
	if len(revision) > 7 {
		revision = revision[:7]
	}
	if modified == "true" {
		revision += "-dirty"
	}
	return revision
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func fail(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(2)
}
