// Command uplinkcheck checks that every approved requirement names approved uplinks only (TOOL-09).
//
//	uplinkcheck [-root dir]
//
// It reads the files of requirements as they stand in the working tree. It prints every uplink of
// an approved requirement that is missing or not approved, and exits with 1 where there is one.
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/treporting/specs-and-deploy/tools/review"
	"github.com/treporting/specs-and-deploy/tools/uplink/check"
)

// @implements SR-28@1.1.1
func main() {

	root := flag.String("root", ".", "the root of the repo")
	flag.Parse()

	findings, approved, err := check.Check(func(file string) ([]byte, error) {
		return os.ReadFile(filepath.Join(*root, file))
	}, list(*root))
	if err != nil {
		fail(err)
	}
	for _, f := range findings {
		fmt.Println(f)
	}
	fmt.Printf("%d approved requirements checked, %d findings\n", approved, len(findings))
	if len(findings) > 0 {
		os.Exit(1)
	}
}

// list answers the names of the files in a directory below root.
func list(root string) review.List {
	return func(dir string) ([]string, error) {
		entries, err := os.ReadDir(filepath.Join(root, dir))
		if err != nil {
			return nil, err
		}
		var names []string
		for _, e := range entries {
			if !e.IsDir() {
				names = append(names, e.Name())
			}
		}
		return names, nil
	}
}

func fail(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(2)
}
