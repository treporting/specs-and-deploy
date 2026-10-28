// Command linkcheck makes the link check (TOOL-16): it compares the version each link, entry of
// the allocation or of controls, and annotation names with the version its object has.
//
//	linkcheck [-root dir]
//
// It reads the files as they stand in the working tree. It prints every one whose object has
// another x than it names, and exits with 0 all the same: it reports, and a suspect link blocks
// the baseline, not the pull request (PROC-20). It exits with 2 where it cannot read the repo.
package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	annotation "github.com/treporting/specs-and-deploy/tools/annotation/check"
	"github.com/treporting/specs-and-deploy/tools/link/check"
)

// @implements SR-32@1.1.1, SR-33@1.1.1
func main() {

	root := flag.String("root", ".", "the root of the repo")
	flag.Parse()

	read := func(file string) ([]byte, error) {
		return os.ReadFile(filepath.Join(*root, file))
	}
	links, comparedLinks, err := check.Links(read)
	if err != nil {
		fail(err)
	}

	out, err := exec.Command("git", "-C", *root, "ls-files").Output()
	if err != nil {
		fail(err)
	}
	kinds, err := annotation.ReadKinds(filepath.Join(*root, "tools", "annotation", "kinds.yaml"))
	if err != nil {
		fail(err)
	}
	//	a file of unknown kind is annotationcheck's to report
	annotations, _, err := annotation.Find(*root, strings.Fields(string(out)), kinds)
	if err != nil {
		fail(err)
	}
	ids, comparedIDs := check.Annotations(read, annotations)

	for _, f := range append(links, ids...) {
		fmt.Println(f)
	}
	fmt.Printf("%d links and %d ids of annotations compared, %d suspect\n", comparedLinks, comparedIDs, len(links)+len(ids))
}

func fail(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(2)
}
