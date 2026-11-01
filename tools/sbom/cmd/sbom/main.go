// Command sbom creates the SBOM of a subsystem as sbom.md in the root of its repo (TOOL_21).
//
//	sbom -label x.y.z -ticket number [-root dir] [-repo name] [-image name]...
//	sbom -check [-root dir] [-repo name] [-image name]...
//
// It runs in the repo checked out at the baseline the label names. It lists the package of the
// subsystem itself; the Go modules of its commands and the standard library of Go, where a directory
// of the repo holds a go.mod; the packages of each package-lock.json of the repo; and the Alpine
// packages of each image named. It lists each package once. It keeps the rows of the Change history of an sbom.md that is there,
// and adds one. It sets the entry of sbom.md in documents.yaml of the repo: the version of the new
// row, the hash of the file written, approved and generated (PROC_13).
//
// With -check it writes nothing: it compares the packages with the table of sbom.md, and exits with
// 1 where they differ (TOOL_28).
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
// @implements SR_59@1.1.1
func main() {

	root := flag.String("root", ".", "the root of the repo of the subsystem")
	label := flag.String("label", "", "the label of the baseline, such as 1.2.0")
	ticket := flag.String("ticket", "", "the number of the ticket of the baseline, such as 343")
	repo := flag.String("repo", "", "the name of the repo; the name of the directory of root where empty")
	check := flag.Bool("check", false, "compare the packages with the table of sbom.md, and write nothing")
	var named images
	flag.Var(&named, "image", "an image of the subsystem, built at the baseline; may be given more than once")
	flag.Parse()

	if *repo == "" {
		abs, err := filepath.Abs(*root)
		if err != nil {
			fail(err)
		}
		*repo = filepath.Base(abs)
	}
	path := filepath.Join(*root, create.Name)

	if *check {
		checkSbom(path, *root, *repo, named)
		return
	}

	if *label == "" {
		fail(fmt.Errorf("-label is missing"))
	}
	if *ticket == "" {
		fail(fmt.Errorf("-ticket is missing"))
	}
	documentsPath := filepath.Join(*root, "documents.yaml")
	documents, err := os.ReadFile(documentsPath)
	if err != nil {
		fail(err)
	}
	next, err := create.Next(create.Version(documents))
	if err != nil {
		fail(err)
	}

	packages := delivered(*root, *repo, *label, named)

	old, _ := os.ReadFile(path) // an sbom.md that is not there has no rows to keep
	history := append(create.History(old), create.Row(next, time.Now().Format("2006-01-02"), *label, *ticket, "sbom "+version()))

	md := []byte(create.Markdown(create.SBOM{Repo: *repo, Packages: packages, History: history}))
	if err := os.WriteFile(path, md, 0o600); err != nil {
		fail(err)
	}
	if err := os.WriteFile(documentsPath, create.Entry(documents, next, create.Hash(md)), 0o600); err != nil {
		fail(err)
	}
	fmt.Printf("%s written: %d packages, version %s\n", path, len(packages), next)
}

// checkSbom compares the packages the subsystem delivers with the table of its sbom.md, with the
// label of its last row as the version of its own package. It writes nothing, and exits with 1
// naming each package that differs.
//
// @implements SR_66@1.1.1
func checkSbom(path, root, repo string, named []string) {
	md, err := os.ReadFile(path)
	if err != nil {
		fail(err)
	}
	label := create.Label(md)
	if label == "" {
		fail(fmt.Errorf("%s: no row in its Change history, so no label", path))
	}
	differ := create.Differ(delivered(root, repo, label, named), create.Table(md))
	for _, line := range differ {
		fmt.Println(line)
	}
	if len(differ) > 0 {
		fmt.Printf("%s does not list the packages delivered: %d differ\n", path, len(differ))
		os.Exit(1)
	}
	fmt.Printf("%s lists the packages delivered\n", path)
}

// delivered answers every package the subsystem delivers, each once: its own, with the label as its
// version; the Go modules of its commands and the standard library of Go, where a directory of the
// repo holds a go.mod; the packages of each package-lock.json of the repo; and the Alpine packages of
// each image.
func delivered(root, repo, label string, named []string) []create.Package {
	packages := []create.Package{create.Own(repo, label, create.LicenceIn(root))}

	dirs, err := create.DirsWith(root, "go.mod")
	if err != nil {
		fail(err)
	}
	var lists, withCommands []string
	for _, dir := range dirs {
		list, err := create.GoList(dir)
		if err != nil {
			fail(err)
		}
		if list != "" {
			lists = append(lists, list)
			withCommands = append(withCommands, dir)
		}
	}
	if len(lists) > 0 {
		packages = append(packages, create.GoModules(strings.Join(lists, "\n"), create.LicenceIn)...)
		version, err := create.GoEnv(withCommands[0], "GOVERSION")
		if err != nil {
			fail(err)
		}
		goroot, err := create.GoEnv(withCommands[0], "GOROOT")
		if err != nil {
			fail(err)
		}
		licence := create.LicenceIn(goroot)
		if licence == "" {
			licence = create.LicenceIn(filepath.Dir(goroot)) // Homebrew keeps it beside GOROOT
		}
		packages = append(packages, create.Stdlib(version, licence))
	}

	locks, err := create.DirsWith(root, "package-lock.json")
	if err != nil {
		fail(err)
	}
	for _, dir := range locks {
		lock, err := os.ReadFile(filepath.Join(dir, "package-lock.json"))
		if err != nil {
			fail(err)
		}
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
	return create.Once(packages)
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

func fail(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(2)
}
