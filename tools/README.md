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

**Decision**: a tool requirement leads straight to code and tests. The code names it in
`@implements TOOL-nn`, a test in `@verifies TOOL-nn`. The tests are on the levels MT and SwIT.

**Why**: a tool requirement is as concrete as a software requirement already. A software
requirement beside it would say the same sentence again.

**Rejected**: tool requirement, software requirement, code - the full chain of the product. It
would double every tool requirement. A tool that grows large can still get software requirements.

## Who changed a requirement

**Decision**: `modified_by` holds the original as git records it: the author of the last commit that
changed the requirement, and its co-authors - `Mathias Fischer, Co-Authored-By: Claude Opus 5.5`.
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

**Decision**: the author of a requirement is the co-author of the last commit that changed it whose
name begins with `Claude`, where there is one; else the author of the commit. The reviewer of a
requirement is not its author (PROC-15, TOOL-03).

**Why**: the commits are made under the account of the Product Owner, so git names them as author of
every commit, also where Claude Code wrote the text; the co-author tells who wrote it. The check
looks at the beginning of the name only - the trailer names the model, `Claude Opus 5.5`, and is not
to break when the model changes.

**Rejected**:

- the reviewer may not appear in `modified_by` at all - every review by the Product Owner would fail,
  as they are the author of every commit
- the check left to the approver by hand - TOOL-03 asks for it
