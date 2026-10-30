# tools

The tools of our work: what checks and supports the requirements, reviews and documents of
treporting. Tools are code (see [devops.md](../devops.md)): every tool realizes an approved tool
requirement of [tool-requirements.yaml](../requirements/tool-requirements.yaml). It is written for
whoever builds, runs or reviews the tools.

## Where the tools live

**Decision**: in `tools/` of this repo, specs-and-deploy.

**Why**: the tools check above all the files of this repo - requirements, reviews, schemas - so what
is checked and what checks lie together. The repo is meant for more than documents already: the
containers and the system tests come here as well. The checks of the code repos run as a reusable
workflow in `treporting/.github`, which takes the tool from here (see Open).

**Rejected**:

- a repo of its own - a clean cut, but the CI of every repo would first have to fetch the tool, and
  it is one repo more
- reporter-annotations - it is meant as a tool beside the product, but a check of reviews has
  nothing to do with annotations

## Language

**Decision**: Go.

**Why**: the same language as the subsystems, and its toolchain stands in devops.md already - no
new language in the chain of our tools, which the threat of an outdated toolchain asks to keep
small. A tool is one program
with no dependency at run time, which the CI runs as it is. It is tested like the code of the
product - a checker with a fault would let faults pass unseen.

**Rejected**:

- Python - quickly written, but a new tool in the chain, with packages of its own, not in
  devops.md, and a setup of its own in the CI
- shell - no dependency, but too error-prone for reading YAML and schemas

## The chain of a tool

**Decision**: a tool requirement leads to software requirements, and those to code and tests - as
for a subsystem. The software requirements of the tools are in
[software-requirements.yaml](software-requirements.yaml), each naming its tool requirement as
uplink. The code names them in `@implements SR_nn`, a test in `@verifies SR_nn`. The tests are on
the levels MT and SwIT.

**Why**: the first tool showed it. TOOL_03 holds five behaviours that can each be tested on their
own; by ISO29148 each is a requirement of its own (singular). Its review noted it, and
each test of reviewcheck verified one of the five while all named only TOOL_03 - which of them a
test verifies could not be told. A tool requirement says what a tool does for our work; a
software requirement says how it behaves.

**Rejected**: a tool requirement leading straight to code and tests - the first decision here. It
held a tool requirement as concrete as a software requirement, which TOOL_03 proved wrong.

## Who changed a requirement

**Decision**: `modified_by` holds the original as git records it: the author of the last commit,
not a merge, that changed the requirement, and its co-authors - `Claude Opus 5.5`, or `Mathias Fischer`.
Nothing is derived from it and kept beside it; there is no `mmodified_by`.

**Why**: git records author and co-authors by itself; a record made by a computer is the evidence an
audit trail asks for (REG_CFR21-11_01), and the original stays original (ARCR_01, REG_FDADI_01).

**Rejected**:

- only the one who wrote it, `claude-code` - a value already interpreted would stand in the field of
  the original, and the author of the commit would be lost
- a field `created_by` or `modified_by` written by hand in the YAML - visible in the file, but a
  claim rather than a record, it can drift from git unseen, and it would be a second source of the
  same truth

## Who the author of a requirement or a document is

**Decision**: the authors of a requirement or a document are the names in its `modified_by` - the
author and the co-authors of the last commit, not a merge, that changed it, whatever names stand
there. The reviewer is none of them (PROC_15; TOOL_03, TOOL_11). Who writes the text is the author of the commit:
Claude Code commits with `--author="Claude Opus 5.5 <noreply@anthropic.com>"`; the Product Owner
stays its committer.

```bash
git commit --author="Claude Opus 5.5 <noreply@anthropic.com>" -m "..."
```

**Why**: git tells author and committer apart by itself. With the writer as author, the check needs
no rule about names: it holds for any assistant or person, and needs no change when the assistant
changes. A commit the Product Owner writes alone names them, and they may not review it - as it
should be.

A merge commit is no authorship: the one who merges approves, and writes nothing. Where two pull
requests change one file from an older `main`, the merge commit equals neither of its parents, and
git counts it as a change of the file - with the one who merged as its author. The merge commit
stays as the record of the approval; it is only not read as authorship. The known limit: a conflict
resolved by hand in a merge is a change no author is found for.

**Rejected**:

- the co-author whose name begins with `Claude`, else the author of the commit - the first decision;
  it names a product, so another assistant, such as Copilot, would need a change of the rule
- the authors without the committer, where none is left the committer - it keeps the history valid,
  but it is a rule about git again, not the record of git
- the check left to the approver by hand - TOOL_03 asks for it
- rebase merges, so that `main` has no merge commits - git would no longer record who merged, the
  evidence of the approval

Every review result made before this decision reviewed commits with the Product Owner as author. It
breaks the check, so its requirements are reviewed again and the old result is removed.

## The component review

**Decision**: everything of the component `review` lies below `tools/review`:

```
tools/review/                the component: what its tools share - git, and finding an id in its file
  check/                     the module review/check: reviewcheck
  skeleton/                  the module review/skeleton: reviewskeleton
  cmd/reviewcheck/           the command
  cmd/reviewskeleton/        the command
```

**Why**: the hierarchy of the component can be read from the folders. `check` and `skeleton`
depend on `review`, not on each other. A command keeps the name of its tool: Go names a program by
the last part of its path. An SR names the component, `review`; the module is set by the code.

**Rejected**:

- the commands in `tools/cmd` and the packages beside them - the component would be spread over
  the tree
- a folder `lib/` - in Go every package without `main` is a library already
- commands named `check` and `skeleton` - too general a name for a program
- the module as the component of an SR - which package realizes a behaviour is a decision of the
  design, on the levels r6 and r7

**Decision**: a tool finds the files of requirements by their `type`. The component `review` reads
the head of every YAML git knows in the working tree, and keeps those whose `type` holds
requirements: `requirements`, `normative`, `architecture`, `hazards` and `rules` (formats.md, table
types). From the `object_key` in them it knows the file of every id. In an old revision a tool reads
an id in the file the working tree names for it.

**Why**: a file says itself what it holds (formats.md, section The head of a file). A new level, a
new file or a changed prefix needs no change of code.

**Rejected**:

- the id tells its file, by a fixed list of prefixes and the rule `REG_<module>_nn` →
  `regulatory/<module>.yaml`, as before - every new level or prefix is a change of code, and the
  list stood a second time beside the table files of the levels in formats.md
- a file `levels.yaml` - see formats.md, section Documents

## reviewcheck

Checks every review result in `reviews/results/` against the rules of formats.md (TOOL_03, TOOL_11; SR_01
to SR_08 and SR_16 in [software-requirements.yaml](software-requirements.yaml)): every criterion of its
criteria file once, the verdict passed only when every criterion passed, a note to every failed
criterion, no value but passed or failed, the revision reviewed exists, the criteria revision changed
the criteria file, and the reviewer is none of the authors of the requirement or document. It
reports the findings of a result by criterion, in the order of its criteria file, the same in every run.

```bash
cd tools
go run ./review/cmd/reviewcheck -root ..        # every result in reviews/results/
go test ./...                            # MT
go test -tags integration ./...          # SwIT, on a history of git made in the test
```

It exits with 1 where a result breaks a rule. It depends on `gopkg.in/yaml.v3` (MIT and Apache-2.0).

## reviewskeleton

Writes the skeleton of a review result into `reviews/results/` (TOOL_04; SR_11 to SR_15 in
[software-requirements.yaml](software-requirements.yaml)): the name from the time and the scope,
the criteria file of the level and the last commit that changed it, the commit `HEAD` as the
revision of each, and every criterion and the verdict as `todo`. It writes nothing where a target is
not found, or where the targets belong to more than one level.

```bash
cd tools
go run ./review/cmd/reviewskeleton -root .. -reviewer "Mathias Fischer" -scope harness harness.md
go run ./review/cmd/reviewskeleton -root .. -reviewer "Mathias Fischer" -scope sr-11-15 SR_11 SR_12 SR_13
```

In its head the skeleton names its schema, the type `review` and the level of its criteria file
(SR_17). It prints the path of the file it
wrote. The code is the package `review/skeleton`, in the component `review`.

## schemacheck

Checks every YAML file of our own data under version control against the schema it names in its
field `schema` (TOOL_05; SR_18 and SR_19 in [software-requirements.yaml](software-requirements.yaml)),
and reports each place where one does not match, with the file and the path within it. A YAML in
`.github/` has a form GitHub owns, and is not checked. The CI runs it on every pull request (SR_20).

```bash
cd tools
go run ./schema/cmd/schemacheck -root ..
```

A finding reads `requirements/tool-requirements.yaml /requirements/6/status: value must be one of
'draft', 'review', 'approved', 'obsolete'`. The code is the component `schema`: the package
`schema/check` and the command `schema/cmd/schemacheck`. Its dependencies are in devops.md.

## rulesetcheck

Compares the ruleset of main on GitHub with `.github/rulesets/main.json` (TOOL_13; SR_21 and SR_22 in
[software-requirements.yaml](software-requirements.yaml)), and reports each difference in its
target, enforcement, conditions and rules. What GitHub sets itself is left out, and `bypass_actors`,
which cannot be read without the rights of an admin; that main has no bypass, its demonstration
shows. A ruleset that cannot be read, or no ruleset of the name in the file, is a finding, not a
pass. The CI runs it on every pull request (SR_23), with the token of the workflow, read only.

```bash
cd tools
GITHUB_TOKEN=$(gh auth token) go run ./ruleset/cmd/rulesetcheck -root ..
```

The code is the component `ruleset`: the package `ruleset/check` and the command
`ruleset/cmd/rulesetcheck`. Its tests stand in for GitHub with a server in the test process
(`net/http/httptest`), so they run without a network.

## annotationcheck

Checks that every id named in `@implements`, `@verifies` or `@refines` is in the requirements of the
base commit of the pull request, and approved there (TOOL_06; SR_24, SR_25 and SR_35 in
[software-requirements.yaml](software-requirements.yaml)), and that it names its version, as
`ID@x.y.z` (TOOL_20, SR_30). The id is read apart from its version and checked as before. A requirement approved in the same pull
request does not count: it has to be through before the code (PROC_05, PROC_12).

It reads a file by its kind, from the two lists in [kinds.yaml](annotation/kinds.yaml): the
whitelist of the kinds that are code, the blacklist of those that are no code. A file of a kind on
neither list is a finding (SR_27), until its kind is set; a YAML that names a schema is our data and
is not read. In a file it reads, an annotation stands first in a line of comment: after spaces, a
comment mark of one to three signs other than a quote, then `@implements` or `@verifies` - in every
language, with no list of languages. Why both lists: formats.md, "Where an annotation is read".
`@refines` it reads in the `$comment` of each schema in [schemas/](../schemas/) (SR_35).

**Decision**: a line of comment is told by its start alone, with no knowledge of the strings of a
language. Its known limit: a line of a string over lines - a raw string of Go, a template of
JavaScript - that begins with a comment mark and an annotation is read as one. A test writes such
content with `\n`, not over lines.

**Why**: the limit does the gate no harm. With an id that is approved, no breach gets through; with
one that is not, the false alarm is loud, and the remedy is a line in a test. It arises in practice
only in tests that need annotations as their content.

**Rejected**:

- a lexer for each kind, set in `kinds.yaml` - comment marks, strings, escapes - a large effort,
  with traps of its own (escapes that differ, a `'` that is no string, strings with no quotes), for
  an error that does the gate no harm
- a parser for each language - right where an error harms: in reporter-annotations, which counts an
  annotation as a realization; not for the gate
- `go/scanner` for Go alone - exact for Go, but a switch by language again

The CI runs it on every pull request with the base commit GitHub gives (SR_26, SR_31).

```bash
cd tools
go run ./annotation/cmd/annotationcheck -root ..                 # base: where HEAD branched from main
go run ./annotation/cmd/annotationcheck -root .. -base <commit>
```

The code is the component `annotation`: the package `annotation/check` and the command
`annotation/cmd/annotationcheck`.

## uplinkcheck

Checks that every approved requirement names approved uplinks only (TOOL_09; SR_28 and SR_29 in
[software-requirements.yaml](software-requirements.yaml)). It reads every file of requirements as it
stands - the levels, the tools, the rules - and reports each uplink in `source_ids` of an approved
requirement that no file holds, or that is not approved. The version after `@` is not compared here;
that is PROC_20.

It checks the state, not a change: on `main` a requirement stays approved until its change passes
its review, and a change whose review failed does not reach `main` (harness.md, step 10). So an
approved requirement never names a draft there, unless someone merges one.

The CI runs it on every pull request (SR_29).

```bash
cd tools
go run ./uplink/cmd/uplinkcheck -root ..
```

The code is the component `uplink`: the package `uplink/check` and the command
`uplink/cmd/uplinkcheck`. It finds an id in its file through the component `review`.

## linkcheck

Makes the link check (TOOL_16; SR_32 to SR_34 in [software-requirements.yaml](software-requirements.yaml)),
as definitions.md sets it out in the section Link check and impact assessment. It compares the
version each entry names with the version its object has, and reports each whose object has
another x (traceability.md, section When a link is suspect):

- the links in `source_ids`, the entries of `allocated` and of `controls` in every file of
  requirements (SR_32)
- the ids of `@implements`, `@verifies` and `@refines`, found as annotationcheck finds them (SR_33)

It reads the files as they stand. An entry whose object no file holds, or that names no version, is
not compared: uplinkcheck, schemacheck and annotationcheck report it. It reports and does not
decide whether the child still fits; that is the impact assessment.

The CI runs it on every pull request in a job of its own (SR_34). The job does not fail and is no
required check: a suspect link blocks the baseline, not the pull request (PROC_20).

```bash
cd tools
go run ./link/cmd/linkcheck -root ..
```

The code is the component `link`: the package `link/check` and the command `link/cmd/linkcheck`. It
finds an id in its file through the component `review`, and the annotations through the component
`annotation`.

## Open

- `treporting = 1_tools` - **Checks of the code repos**: not set up yet. They are to run as a
  reusable workflow in `treporting/.github`, which takes the tools from specs-and-deploy.
