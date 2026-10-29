// Command reviewcheck checks the review results of a repo against the rules of formats.md (TOOL_03).
//
//	reviewcheck [-root dir] [file ...]
//
// With no file, it checks every result in reviews/results/. It prints every
// rule a result breaks and exits with 1 where there is one.
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/treporting/specs-and-deploy/tools/review"
	"github.com/treporting/specs-and-deploy/tools/review/check"
)

// @implements SR_01@1.1.1, SR_02@1.1.1, SR_03@1.1.1, SR_04@1.1.1, SR_05@1.1.1, SR_06@1.1.1, SR_07@1.1.1, SR_08@1.1.1
func main() {

	root := flag.String("root", ".", "the root of the repo")
	flag.Parse()

	files := flag.Args()
	if len(files) == 0 {
		found, err := filepath.Glob(filepath.Join(*root, "reviews", "results", "*.yaml"))
		if err != nil {
			fail(err)
		}
		for _, f := range found {
			rel, _ := filepath.Rel(*root, f)
			files = append(files, rel)
		}
	}

	git := review.ExecGit{Root: *root}
	broken := 0
	for _, f := range files {
		findings, err := check.Check(*root, f, git)
		if err != nil {
			fail(err)
		}
		for _, finding := range findings {
			fmt.Println(finding)
		}
		broken += len(findings)
	}
	fmt.Printf("%d review results checked, %d findings\n", len(files), broken)
	if broken > 0 {
		os.Exit(1)
	}
}

func fail(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(2)
}
