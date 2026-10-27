// Command annotationcheck checks that every id named in @implements or @verifies is approved in the
// base commit of the pull request (TOOL-06).
//
//	annotationcheck [-root dir] [-base rev]
//
// Without -base, the base commit is where HEAD branched from main. It prints every id that is
// missing or not approved there, and exits with 1 where there is one.
package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/treporting/specs-and-deploy/tools/annotation/check"
	"github.com/treporting/specs-and-deploy/tools/review"
)

// @implements SR-24, SR-25, SR-27
func main() {

	root := flag.String("root", ".", "the root of the repo")
	base := flag.String("base", "", "the base commit of the pull request; where HEAD branched from main, if empty")
	flag.Parse()

	if *base == "" {
		out, err := exec.Command("git", "-C", *root, "merge-base", "HEAD", "main").Output()
		if err != nil {
			fail(fmt.Errorf("git merge-base HEAD main: %w", err))
		}
		*base = strings.TrimSpace(string(out))
	}

	out, err := exec.Command("git", "-C", *root, "ls-files").Output()
	if err != nil {
		fail(err)
	}
	kinds, err := check.ReadKinds(filepath.Join(*root, "tools", "annotation", "kinds.yaml"))
	if err != nil {
		fail(err)
	}
	annotations, unknown, err := check.Find(*root, strings.Fields(string(out)), kinds)
	if err != nil {
		fail(err)
	}
	findings := append(unknown, check.Check(review.ExecGit{Root: *root}, *base, annotations)...)
	for _, f := range findings {
		fmt.Println(f)
	}
	ids := 0
	for _, a := range annotations {
		ids += len(a.IDs)
	}
	fmt.Printf("%d annotations, %d ids, checked against the base commit %.7s, %d findings\n", len(annotations), ids, *base, len(findings))
	if len(findings) > 0 {
		os.Exit(1)
	}
}

func fail(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(2)
}
