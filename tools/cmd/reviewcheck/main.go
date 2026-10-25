// Command reviewcheck checks the review results of a repo against the rules of formats.md (TOOL-03).
//
//	reviewcheck [-root dir] [file ...]
//
// With no file, it checks every result in reviews/ - every YAML but the criteria. It prints every
// rule a result breaks and exits with 1 where there is one.
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/treporting/specs-and-deploy/tools/review"
)

// @implements SR-01, SR-02, SR-03, SR-04, SR-05, SR-06, SR-07, SR-08
func main() {

	root := flag.String("root", ".", "the root of the repo")
	flag.Parse()

	files := flag.Args()
	if len(files) == 0 {
		found, err := filepath.Glob(filepath.Join(*root, "reviews", "*.yaml"))
		if err != nil {
			fail(err)
		}
		for _, f := range found {
			if !strings.HasPrefix(filepath.Base(f), "criteria-") {
				rel, _ := filepath.Rel(*root, f)
				files = append(files, rel)
			}
		}
	}

	git := review.ExecGit{Root: *root}
	broken := 0
	for _, f := range files {
		findings, err := review.Check(*root, f, git)
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
