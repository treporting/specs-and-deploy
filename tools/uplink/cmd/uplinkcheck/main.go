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

	"github.com/treporting/specs-and-deploy/tools/uplink/check"
)

// @implements SR-28
func main() {

	root := flag.String("root", ".", "the root of the repo")
	flag.Parse()

	findings, approved, err := check.Check(func(file string) ([]byte, error) {
		return os.ReadFile(filepath.Join(*root, file))
	})
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

func fail(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(2)
}
