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
workflow in `treporting/.github`, which takes the tool from here. That workflow is still open.

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
uplink. The code names them in `@implements SR-nn`, a test in `@verifies SR-nn`. The tests are on
the levels MT and SwIT.

**Why**: the first tool showed it. TOOL-03 holds five behaviours that can each be tested on their
own; by ISO29148 each is a requirement of its own (singular). Its review noted it, and
each test of reviewcheck verified one of the five while all named only TOOL-03 - which of them a
test verifies could not be told. A tool requirement says what a tool does for our work; a
software requirement says how it behaves.

**Rejected**: a tool requirement leading straight to code and tests - the first decision here. It
held a tool requirement as concrete as a software requirement, which TOOL-03 proved wrong.

## Who changed a requirement

**Decision**: `modified_by` holds the original as git records it: the author of the last commit,
not a merge, that changed the requirement, and its co-authors - `Claude Opus 5.5`, or `Mathias Fischer`.
Nothing is derived from it and kept beside it; there is no `mmodified_by`.

**Why**: git records author and co-authors by itself; a record made by a computer is the evidence an
audit trail asks for (REG-05), and the original stays original (ARC-R01, REG-10).

**Rejected**:

- only the one who wrote it, `claude-code` - a value already interpreted would stand in the field of
  the original, and the author of the commit would be lost
- a field `created_by` or `modified_by` written by hand in the YAML - visible in the file, but a
  claim rather than a record, it can drift from git unseen, and it would be a second source of the
  same truth

## Who the author of a requirement or a document is

**Decision**: the authors of a requirement or a document are the names in its `modified_by` - the
author and the co-authors of the last commit, not a merge, that changed it, whatever names stand
there. The reviewer is none of them (PROC-15; TOOL-03, TOOL-11). Who writes the text is the author of the commit:
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
- the check left to the approver by hand - TOOL-03 asks for it
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

## reviewcheck

Checks every review result in `reviews/` against the rules of formats.md (TOOL-03, TOOL-11; SR-01
to SR-08 and SR-16 in [software-requirements.yaml](software-requirements.yaml)): every criterion of its
criteria file once, the verdict passed only when every criterion passed, a note to every failed
criterion, no value but passed or failed, the revision reviewed exists, the criteria revision changed
the criteria file, and the reviewer is none of the authors of the requirement or document. It
reports the findings of a result by criterion, in the order of its criteria file, the same in every run.

```bash
cd tools
go run ./review/cmd/reviewcheck -root ..        # every result in reviews/
go test ./...                            # MT
go test -tags integration ./...          # SwIT, on a history of git made in the test
```

It exits with 1 where a result breaks a rule. It depends on `gopkg.in/yaml.v3` (MIT and Apache-2.0).

## reviewskeleton

Writes the skeleton of a review result into `reviews/` (TOOL-04; SR-11 to SR-15 in
[software-requirements.yaml](software-requirements.yaml)): the name from the time and the scope,
the criteria file of the level and the last commit that changed it, the commit `HEAD` as the
revision of each, and every criterion and the verdict as `todo`. It writes nothing where a target is
not found, or where the targets belong to more than one level.

```bash
cd tools
go run ./review/cmd/reviewskeleton -root .. -reviewer "Mathias Fischer" -scope harness harness.md
go run ./review/cmd/reviewskeleton -root .. -reviewer "Mathias Fischer" -scope sr-11-15 SR-11 SR-12 SR-13
```

The skeleton names its schema in the field `schema` (SR-17). It prints the path of the file it
wrote. The code is the package `review/skeleton`, in the component `review`.

## schemacheck

Checks every YAML file of our own data under version control against the schema it names in its
field `schema` (TOOL-05; SR-18 and SR-19 in [software-requirements.yaml](software-requirements.yaml)),
and reports each place where one does not match, with the file and the path within it. A YAML in
`.github/` has a form GitHub owns, and is not checked. The CI runs it on every pull request (SR-20).

```bash
cd tools
go run ./schema/cmd/schemacheck -root ..
```

A finding reads `requirements/tool-requirements.yaml /requirements/6/status: value must be one of
'draft', 'review', 'approved', 'obsolete'`. The code is the component `schema`: the package
`schema/check` and the command `schema/cmd/schemacheck`. Its dependencies are in devops.md.

## rulesetcheck

Compares the ruleset of main on GitHub with `.github/rulesets/main.json` (TOOL-13; SR-21 and SR-22 in
[software-requirements.yaml](software-requirements.yaml)), and reports each difference in its
target, enforcement, conditions and rules. What GitHub sets itself is left out, and `bypass_actors`,
which cannot be read without the rights of an admin; that main has no bypass, its demonstration
shows. A ruleset that cannot be read, or no ruleset of the name in the file, is a finding, not a
pass. The CI runs it on every pull request (SR-23), with the token of the workflow, read only.

```bash
cd tools
GITHUB_TOKEN=$(gh auth token) go run ./ruleset/cmd/rulesetcheck -root ..
```

The code is the component `ruleset`: the package `ruleset/check` and the command
`ruleset/cmd/rulesetcheck`. Its tests stand in for GitHub with a server in the test process
(`net/http/httptest`), so they run without a network.
