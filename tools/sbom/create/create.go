// Package create creates the SBOM of a subsystem as sbom.md (TOOL_21): every package it delivers,
// its own and third-party, with its name, version, licence and PURL.
package create

import (
	"bufio"
	"sort"
	"strings"
)

// Unknown is the licence of a package whose licence is not recognized.
const Unknown = "unknown"

// Package is one row of the table of sbom.md.
type Package struct {
	PURL    string
	Name    string
	Version string
	Licence string
	Named   string // the licence as the package names it itself, where it does
}

// SBOM is what sbom.md is written from.
type SBOM struct {
	Repo     string    // the repo of the subsystem, such as storages
	Packages []Package // every package it delivers, its own among them
	History  []string  // the rows of the Change history, the new one last
}

// Own answers the package of the subsystem itself: the label of its baseline as its version, the
// licence of the LICENSE file in the root of its repo, and its PURL on GitHub.
//
// @implements SR_46@1.1.1
func Own(repo, label, licenceText string) Package {
	return Package{
		PURL:    "pkg:github/treporting/" + repo + "@" + label,
		Name:    repo,
		Version: label,
		Licence: Licence(licenceText),
	}
}

// Row answers the row of the Change history that a run adds: the version of sbom.md after the run;
// the day it runs; the label of the baseline; the ticket, as #<number>; the tool with its version.
//
// @implements SR_40@3.3.1
func Row(version, date, label, ticket, by string) string {
	return "| " + cell(version) + " | " + date + " | " + cell(label) + " | #" + cell(ticket) + " | " + cell(by) + " |"
}

// History answers the rows of the Change history of an sbom.md as it stands, without its head, so
// that a run keeps them. A file that is not there has none.
//
// @implements SR_40@3.3.1
func History(sbomMD []byte) []string {
	var rows []string
	in := false
	scanner := bufio.NewScanner(strings.NewReader(string(sbomMD)))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		switch {
		case strings.HasPrefix(line, "## "):
			in = line == "## Change history"
		case in && strings.HasPrefix(line, "| ") && !strings.HasPrefix(line, "| Version |"):
			rows = append(rows, line)
		}
	}
	return rows
}

// Markdown answers sbom.md: the table of the packages sorted by PURL, then the sections Open and
// Change history.
//
// @implements SR_41@1.2.1
// @implements SR_42@1.2.1
// @implements SR_45@1.1.1
func Markdown(s SBOM) string {
	packages := append([]Package(nil), s.Packages...)
	sort.Slice(packages, func(i, j int) bool { return packages[i].PURL < packages[j].PURL })

	var b strings.Builder
	b.WriteString("# SBOM\n\n")
	b.WriteString("The packages " + s.Repo + " delivers, its own and third-party, each with its version, its\n")
	b.WriteString("licence and its PURL. It is written by the tool sbom (TOOL_21) and not edited by hand.\n\n")
	b.WriteString("**Table packages** - every package " + s.Repo + " delivers\n\n")
	b.WriteString("| PURL | Name | Version | Licence |\n|---|---|---|---|\n")
	var unknown []Package
	for _, p := range packages {
		b.WriteString("| " + cell(p.PURL) + " | " + cell(p.Name) + " | " + cell(p.Version) + " | " + cell(p.Licence) + " |\n")
		if p.Licence == Unknown {
			unknown = append(unknown, p)
		}
	}
	b.WriteString("\n## Open\n\n")
	if len(unknown) == 0 {
		b.WriteString("`<empty>`\n")
	}
	for _, p := range unknown {
		b.WriteString("- `" + p.PURL + "` - the licence is unknown")
		if p.Named != "" {
			b.WriteString("; the package names it \"" + p.Named + "\"")
		}
		b.WriteString("\n")
	}
	b.WriteString("\n## Change history\n\n")
	b.WriteString("| Version | Date | Change | Reason | By |\n|---|---|---|---|---|\n")
	for _, row := range s.History {
		b.WriteString(row + "\n")
	}
	return b.String()
}

// cell keeps a value from breaking the row of a table.
func cell(value string) string {
	return strings.ReplaceAll(value, "|", `\|`)
}
