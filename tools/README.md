# tools

The tools of our work: what checks and supports the requirements, reviews and documents of
treporting. Tools are code (see [devops.md](../devops.md)): every tool realizes an approved tool
requirement of [tool-requirements.yaml](../requirements/tool-requirements.yaml).

## Where the tools live

**Decision**: in `tools/` of this repo, specs-and-deploy.

**Why**: the tools check above all the files of this repo - requirements, reviews, schemas - so what
is checked and what checks lie together. The repo is meant for more than documents already: the
containers and the system tests come here as well. The checks of the code repos run as a reusable
workflow in `treporting/.github`, which takes the tool from here.

**Rejected**:

- a repo of its own - a clean cut, but the CI of every repo would first have to fetch the tool, and
  it is one repo more
- reporter-annotations - it is meant as a tool beside the product, but a check of reviews has
  nothing to do with annotations

## Language

**Decision**: Go.

**Why**: the same language as the subsystems, and its toolchain stands in devops.md already - no
new language in the chain of our tools, which THR-02 asks to keep small. A tool is one program
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
own; by ISO/IEC/IEEE 29148 each is a requirement of its own (singular). Its review noted it, and
each test of reviewcheck verified one of the five while all named only TOOL-03 - which of them a
test verifies could not be told. A tool requirement says what a tool does for our work; a
software requirement says how it behaves.

**Rejected**: a tool requirement leading straight to code and tests - the first decision here. It
held a tool requirement as concrete as a software requirement, which TOOL-03 proved wrong.

## Who changed a requirement

**Decision**: `modified_by` holds the original as git records it: the author of the last commit that
changed the requirement, and its co-authors - `Claude Opus 5.5`, or `Mathias Fischer`.
Nothing is derived from it and kept beside it; there is no `mmodified_by`.

**Why**: git records author and co-authors by itself; a record made by a computer is the evidence an
audit trail asks for (REG-05), and the original stays original (ARC-R01, REG-10).

**Rejected**:

- only the one who wrote it, `claude-code` - a value already interpreted would stand in the field of
  the original, and the author of the commit would be lost
- a field `created_by` or `modified_by` written by hand in the YAML - visible in the file, but a
  claim rather than a record, it can drift from git unseen, and it would be a second source of the
  same truth

## Who the author of a requirement is

**Decision**: the authors of a requirement are the names in its `modified_by` - the author and the
co-authors of the last commit that changed it, whatever names stand there. The reviewer of a
requirement is none of them (PROC-15, TOOL-03). Who writes the text is the author of the commit:
Claude Code commits with `--author="Claude Opus 5.5 <noreply@anthropic.com>"`; the Product Owner
stays its committer.

```bash
git commit --author="Claude Opus 5.5 <noreply@anthropic.com>" -m "..."
```

**Why**: git tells author and committer apart by itself. With the writer as author, the check needs
no rule about names: it holds for any assistant or person, and needs no change when the assistant
changes. A commit the Product Owner writes alone names them, and they may not review it - as it
should be.

**Rejected**:

- the co-author whose name begins with `Claude`, else the author of the commit - the first decision;
  it names a product, so another assistant, such as Copilot, would need a change of the rule
- the authors without the committer, where none is left the committer - it keeps the history valid,
  but it is a rule about git again, not the record of git
- the check left to the approver by hand - TOOL-03 asks for it

Every review result made before this decision reviewed commits with the Product Owner as author. It
breaks the check, so its requirements are reviewed again and the old result is removed.

## reviewcheck

Checks every review result in `reviews/` against the rules of formats.md (TOOL-03): every criterion
of its criteria file once, the verdict passed only when every criterion passed, a note to every
failed criterion, the revision reviewed exists, and the reviewer is none of its authors.

```bash
cd tools
go run ./cmd/reviewcheck -root ..        # every result in reviews/
go test ./...                            # MT
go test -tags integration ./...          # SwIT, on a history of git made in the test
```

It exits with 1 where a result breaks a rule. It depends on `gopkg.in/yaml.v3` (MIT and Apache-2.0).
