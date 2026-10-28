// Package check checks that every approved requirement names approved uplinks only (TOOL-09).
package check

import (
	"fmt"
	"strings"

	"github.com/treporting/specs-and-deploy/tools/review"
	"gopkg.in/yaml.v3"
)

// Read answers the content of a file of the repo, by its path from the root.
type Read func(file string) ([]byte, error)

// Finding is one uplink of an approved requirement that no file holds, or that is not approved.
type Finding struct {
	ID      string // the approved requirement
	Uplink  string // the uplink as it names it, with its version
	Message string
}

func (f Finding) String() string {
	return f.ID + " names " + f.Uplink + ": " + f.Message
}

// Check reads every file of requirements and answers each uplink in the source_ids of an approved
// requirement that no file holds, or that is not approved. It answers too how many approved
// requirements it checked. A file of requirements that is not there holds none.
//
// @implements SR-28@1.1.1
func Check(read Read, list review.List) ([]Finding, int, error) {

	files := map[string][]byte{}
	load := func(file string) []byte {
		if data, ok := files[file]; ok {
			return data
		}
		data, err := read(file)
		if err != nil {
			data = nil
		}
		files[file] = data
		return data
	}

	var findings []Finding
	approved := 0
	for _, file := range review.Files(list) {
		for _, object := range objects(load(file)) {
			if object.Status != "approved" {
				continue
			}
			approved++
			for _, uplink := range object.SourceIDs {
				if message := status(load, uplink); message != "" {
					findings = append(findings, Finding{object.Key, uplink, message})
				}
			}
		}
	}
	return findings, approved, nil
}

type object struct {
	Key       string   `yaml:"object_key"`
	Status    string   `yaml:"status"`
	SourceIDs []string `yaml:"source_ids"`
}

// objects answers the objects of every section of a file of requirements; none where it cannot be
// read.
func objects(data []byte) []object {

	var doc yaml.Node
	if yaml.Unmarshal(data, &doc) != nil || len(doc.Content) == 0 || doc.Content[0].Kind != yaml.MappingNode {
		return nil
	}
	//	the sections in the order of the file: keys and values alternate
	sections := doc.Content[0].Content
	var out []object
	for i := 1; i < len(sections); i += 2 {
		var list []object
		if sections[i].Decode(&list) == nil {
			for _, o := range list {
				if o.Key != "" {
					out = append(out, o)
				}
			}
		}
	}
	return out
}

// status answers why an uplink does not count, or "" where it is approved. The uplink names its
// version after @; the version is not compared here (PROC-20).
func status(load func(string) []byte, uplink string) string {

	id, _, _ := strings.Cut(uplink, "@")
	file, key := review.FileOf(id)
	if file == "" || key == "" {
		return "no file of requirements holds such an id"
	}
	data := load(file)
	if data == nil {
		return fmt.Sprintf("%s cannot be read", file)
	}
	found := review.ObjectIn(data, key)
	if found == nil {
		return fmt.Sprintf("not in %s", file)
	}
	var o object
	if err := yaml.Unmarshal(found, &o); err != nil {
		return err.Error()
	}
	if o.Status != "approved" {
		if o.Status == "" {
			return "it has no status"
		}
		return o.Status + ", not approved"
	}
	return ""
}
