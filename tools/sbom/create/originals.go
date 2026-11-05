package create

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

// Original is one file of original data sbom.md is written from.
type Original struct {
	File     string // its name without the time of the run, such as golist-tools.txt
	MadeBy   string // the command that made it
	StoredAt string // where it is stored, as ./sbom/<file>, or -
	Data     []byte
}

// Folder is the folder of sbom.md and of the original data stored with it, from the root of the repo.
const Folder = "sbom"

// GoListOriginal answers the output of GoList in dir, a directory of the repo in root.
//
// @implements SR_101@1.1.1, SR_102@1.1.1
func GoListOriginal(root, dir, list string) Original {
	return Original{
		File:     "golist-" + slug(root, dir) + ".txt",
		MadeBy:   in(root, dir) + "go list -deps -f '" + moduleFormat + "' $(" + commandsCommand + ")",
		StoredAt: "-",
		Data:     []byte(list),
	}
}

// LockOriginal answers the package-lock.json of dir, a directory of the repo in root.
//
// @implements SR_101@1.1.1, SR_102@1.1.1
func LockOriginal(root, dir string, lock []byte) Original {
	return Original{
		File:     slug(root, dir) + "-package-lock.json",
		MadeBy:   in(root, dir) + "npm install",
		StoredAt: "-",
		Data:     lock,
	}
}

// ApkOriginal answers the /lib/apk/db/installed of an image.
//
// @implements SR_101@1.1.1, SR_102@1.1.1
func ApkOriginal(image string, installed []byte) Original {
	return Original{
		File:     "apk-" + dashes(image) + ".txt",
		MadeBy:   "docker run --rm --entrypoint cat " + image + " /lib/apk/db/installed",
		StoredAt: "-",
		Data:     installed,
	}
}

// Store writes each original into the folder sbom/ of the repo in root as <ts>_<file>, and answers
// them with where each is stored.
//
// @implements SR_102@1.1.1
func Store(root, ts string, originals []Original) ([]Original, error) {
	if err := os.MkdirAll(filepath.Join(root, Folder), 0o750); err != nil {
		return nil, err
	}
	stored := make([]Original, 0, len(originals))
	for _, o := range originals {
		name := ts + "_" + o.File
		if err := os.WriteFile(filepath.Join(root, Folder, name), o.Data, 0o600); err != nil {
			return nil, err
		}
		o.StoredAt = "./" + Folder + "/" + name
		stored = append(stored, o)
	}
	return stored, nil
}

// OriginalYAML answers sbom/<ts>.yaml: every original stored, with the command that made it.
//
// @implements SR_102@1.1.1
func OriginalYAML(label, commit, createdAt string, originals []Original) ([]byte, error) {
	type file struct {
		File   string `yaml:"file"`
		MadeBy string `yaml:"made_by"`
	}
	record := struct {
		Schema    string `yaml:"schema"`
		Type      string `yaml:"type"`
		Document  string `yaml:"document"`
		Label     string `yaml:"label"`
		Commit    string `yaml:"commit"`
		CreatedAt string `yaml:"created_at"`
		Files     []file `yaml:"files"`
	}{"schemas/original-data.schema.json", "original-data", Name, label, commit, createdAt, nil}
	for _, o := range originals {
		record.Files = append(record.Files, file{File: strings.TrimPrefix(o.StoredAt, "./"+Folder+"/"), MadeBy: o.MadeBy})
	}
	var b bytes.Buffer
	b.WriteString("# The original data " + Name + " is written from (formats.md, section Documents).\n")
	enc := yaml.NewEncoder(&b)
	enc.SetIndent(2)
	if err := enc.Encode(record); err != nil {
		return nil, err
	}
	return b.Bytes(), enc.Close()
}

// slug answers dir from root for a file name: root for the root itself, with / and : as -.
func slug(root, dir string) string {
	rel, err := filepath.Rel(root, dir)
	if err != nil || rel == "." {
		return "root"
	}
	return dashes(filepath.ToSlash(rel))
}

// dashes answers name with / and : as -.
func dashes(name string) string {
	return strings.NewReplacer("/", "-", ":", "-").Replace(name)
}

// in answers the cd into dir that a command needs, and nothing for the root.
func in(root, dir string) string {
	rel, err := filepath.Rel(root, dir)
	if err != nil || rel == "." {
		return ""
	}
	return "cd " + filepath.ToSlash(rel) + " && "
}
