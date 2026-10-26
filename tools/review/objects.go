// Package review is the component review: what its tools share - the history of git, and how an
// id or a document is found in its file. The check is in review/check, the skeleton in
// review/skeleton, their commands in review/cmd.
package review

import (
	"strings"

	"gopkg.in/yaml.v3"
)

// requirementFiles tell by the prefix of an id the file that holds it.
var requirementFiles = []struct{ prefix, file string }{
	{"UN-", "requirements/user-needs.yaml"},
	{"REG-", "requirements/regulatory-requirements.yaml"},
	{"SYS-", "requirements/system-requirements.yaml"},
	{"TOOL-", "requirements/tool-requirements.yaml"},
	{"SR-", "tools/software-requirements.yaml"},
	{"ARC-", "architecture/system-architecture.yaml"},
	{"HAZC-", "safety/hazards.yaml"},
	{"HAZ-", "safety/hazards.yaml"},
}

// lists are the sections of a YAML that hold objects with an object_key.
var lists = []string{"requirements", "elements", "hazards", "controls"}

// FileOf answers the file that holds a target, and the key to find it there; a document is its
// own file and has no key.
func FileOf(target string) (file, key string) {

	if strings.HasSuffix(target, ".md") {
		return target, ""
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
