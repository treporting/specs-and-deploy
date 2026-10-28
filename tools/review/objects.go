// Package review is the component review: what its tools share - the history of git, and how an
// id or a document is found in its file. The check is in review/check, the skeleton in
// review/skeleton, their commands in review/cmd.
package review

import (
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

// requirementFiles tell by the prefix of an id the file that holds it.
var requirementFiles = []struct{ prefix, file string }{
	{"UN_", "requirements/user-needs.yaml"},
	{"SYS_", "requirements/system-requirements.yaml"},
	{"TOOL_", "requirements/tool-requirements.yaml"},
	{"SR_", "tools/software-requirements.yaml"},
	{"ARC_", "architecture/system-architecture.yaml"},
	{"ARCD_", "architecture/system-architecture.yaml"},
	{"ARCR_", "architecture/system-architecture.yaml"},
	{"HAZC_", "safety/hazards.yaml"},
	{"HAZ_", "safety/hazards.yaml"},
	{"PROC_", "rules.yaml"},
}

// regulatory is the directory of the regulatory requirements, one file per module: the id
// REG_<module>_nn is in regulatory/<module>.yaml (formats.md, section Regulatory requirements).
const regulatory = "regulatory"

// List answers the names of the files in a directory of the repo, by its path from the root.
type List func(dir string) ([]string, error)

// Files answers every file that holds requirements, each once: those of requirementFiles in their
// order, then the modules in regulatory/ that list finds. A directory list cannot read holds none.
func Files(list List) []string {

	var files []string
	seen := map[string]bool{}
	for _, rf := range requirementFiles {
		if !seen[rf.file] {
			seen[rf.file] = true
			files = append(files, rf.file)
		}
	}
	names, _ := list(regulatory)
	sort.Strings(names)
	for _, name := range names {
		if strings.HasSuffix(name, ".yaml") {
			files = append(files, regulatory+"/"+name)
		}
	}
	return files
}

// moduleOf answers the module a regulatory requirement names in its id, REG_<module>_nn, and
// false where the id has another form.
func moduleOf(id string) (string, bool) {

	rest, ok := strings.CutPrefix(id, "REG_")
	if !ok {
		return "", false
	}
	i := strings.LastIndex(rest, "_")
	if i <= 0 || i == len(rest)-1 || strings.Trim(rest[i+1:], "0123456789") != "" {
		return "", false
	}
	return rest[:i], true
}

// lists are the sections of a YAML that hold objects with an object_key.
var lists = []string{"requirements", "elements", "decisions", "hazards", "controls", "rules"}

// FileOf answers the file that holds a target, and the key to find it there; a document is its
// own file and has no key.
func FileOf(target string) (file, key string) {

	if strings.HasSuffix(target, ".md") {
		return target, ""
	}
	if module, ok := moduleOf(target); ok {
		return regulatory + "/" + module + ".yaml", target
	}
	for _, rf := range requirementFiles {
		if strings.HasPrefix(target, rf.prefix) {
			return rf.file, target
		}
	}
	return "", ""
}

// ObjectIn answers the object with key in the YAML data, as YAML, or the whole data where key is
// empty; nil where it is not there.
func ObjectIn(data []byte, key string) []byte {

	if key == "" {
		return data
	}
	var doc map[string]any
	if yaml.Unmarshal(data, &doc) != nil {
		return nil
	}
	for _, list := range lists {
		items, _ := doc[list].([]any)
		for _, item := range items {
			object, _ := item.(map[string]any)
			if object["object_key"] == key {
				out, _ := yaml.Marshal(object)
				return out
			}
		}
	}
	return nil
}
