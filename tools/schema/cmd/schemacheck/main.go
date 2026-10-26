// Command schemacheck checks every YAML file under version control against the schema it names in
// its field schema (TOOL-05).
//
//	schemacheck [-root dir]
//
// It prints every place where a file does not match, and every file that names no schema or a
// schema that is not there, and exits with 1 where there is one.
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/treporting/specs-and-deploy/tools/schema/check"
)

// @implements SR-18, SR-19
func main() {

	root := flag.String("root", ".", "the root of the repo")
	flag.Parse()

	files, err := check.Files(*root)
	if err != nil {
		fail(err)
	}
	findings, err := check.Check(*root, files)
	if err != nil {
		fail(err)
	}
	for _, f := range findings {
		fmt.Println(f)
	}
	fmt.Printf("%d YAML files checked, %d findings\n", len(files), len(findings))
	if len(findings) > 0 {
		os.Exit(1)
	}
}

func fail(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(2)
}
