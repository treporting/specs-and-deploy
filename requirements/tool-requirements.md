# Tool requirements

The requirements on the tools of our work. Tools are code (see [devops.md](../devops.md)): every
tool realizes an approved requirement of this file. It is written for whoever builds a tool or
reviews its requirement, and for the auditor who asks how the tools are controlled.

The requirements themselves - ids `TOOL_nn` - are in [tool-requirements.yaml](tool-requirements.yaml),
in the format of [formats.md](../formats.md). A tool requirement names as uplink the rule of the
harness (`PROC_nn`) or the regulatory requirement (`REG_<module>_nn`) it serves. It is reviewed and approved
like every other requirement (PROC_13). Each section after Launches is a feature, the field `feature`.

## Launches

**Decision**: a tool requirement names the launch of the tools it is planned for in
`platform_launch`: `treporting = 1_tools`, `treporting = 2_tools`, `treporting = 3_tools`.

**Why**: a launch says when a requirement is to be available (definitions.md, section Platform,
launch and stream). It stays right when a release is made.

**Rejected**:

- `priority: 'later'` - it says "not now", not for which release
- a list in a document - a second place beside the requirement

## Documents to PDF

`TOOL_01` to `TOOL_02`: in [tool-requirements.yaml](tool-requirements.yaml).

A document handed to an auditor shows what it is, which version, which revision, and its status
(REG_FDADI_01).

**Which version and revision** an exported document shows:

**Decision**: its version, `x.y.z` (traceability.md, section The version of a document), and its
revision: the last commit that changed its file - the revision tells changes of the same item apart
(definitions.md). For example "devops.md · 2.1.1 · a1b2c3d · approved".

**Why**: the version shows which approved state lies before the auditor, and it matches the entry
in the YAML. A draft keeps its version while its content changes; only the revision tells two of
its exports apart.

**Rejected**:

- the version alone: two exports of a draft of different days would look the same
- the revision alone: the auditor would have to look up in git which approved state it is
- the commit the export ran on: two exports of the same content would show different revisions

**Which status** an exported document shows:

**Table status variants** - what the footer of an exported document could show

| | In the footer | From |
|---|---|---|
| a) a summary of the status of the requirements it holds | "70 draft, 4 review, 2 approved" | counted from the YAML |
| b) a status of the document computed from them | "approved" when all are approved, else "draft" | derived from a) |
| c) **a status of the document itself** | "draft", "released", ... | an attribute of its own - in the model a document (Artifact) has `status` and `mstatus` |

**Decision**: c). A document has a status of its own, as the model has it for a document, and it
works for every document - also for one that holds no requirements, such as devops.md or
harness.md, where a) and b) would give none.

Where it keeps it: in `documents.yaml`, with the values of a requirement (formats.md, section
Documents).

## Reviews

`TOOL_03` to `TOOL_04`, `TOOL_11`, `TOOL_15` and `TOOL_29`: in [tool-requirements.yaml](tool-requirements.yaml).

The tools that create and check a review. `reviewcheck` checks every review result (TOOL_03,
TOOL_11) and runs on every pull request (TOOL_12). `reviewskeleton` writes the skeleton of a review
result (TOOL_04). A requirement approved in a pull request, for the first time or with a new x,
needs a passed review result in that pull request (TOOL_15, PROC_13). So does a document with the
verification method `review` (TOOL_29); a document a tool writes, with `generated`, needs none.

## CI check

`TOOL_05` to `TOOL_10`, `TOOL_12`, `TOOL_13` and `TOOL_20`: in [tool-requirements.yaml](tool-requirements.yaml).

The checks run on every pull request - one tool requirement per check.

## Versions

`TOOL_14`, `TOOL_16` to `TOOL_19`: in [tool-requirements.yaml](tool-requirements.yaml).

The tools for the version of an object and the links that name it (traceability.md, PROC_14,
PROC_20). TOOL_14 checks that an approval sets the right version, against the base commit of the
pull request; TOOL_19 checks that an approved document is the file approved. TOOL_16 makes the link
check (definitions.md): it reports, and does not fail, as a suspect link blocks the baseline, not
the pull request. TOOL_17 is that block. TOOL_18 updates a link.

## sbom

`TOOL_21`: in [tool-requirements.yaml](tool-requirements.yaml).

A package of the SBOM is named by its PURL, such as `pkg:golang/gopkg.in/yaml.v3@v3.0.1`.

## licence check

`TOOL_22`: in [tool-requirements.yaml](tool-requirements.yaml).

## build steps

`TOOL_23` to `TOOL_26`: in [tool-requirements.yaml](tool-requirements.yaml).

## baseline

`TOOL_27` to `TOOL_28`: in [tool-requirements.yaml](tool-requirements.yaml).

The baseline of a repo is a change of its own: it adds sbom.md to `main`. So it needs a ticket and
a branch of that ticket (PROC_02).

**How a baseline is started**:

**Decision**: a ticket starts it. Whoever wants a baseline opens a ticket in the repo, with the
template Baseline, and names the label of the baseline in it. Whoever may label tickets of the repo
then sets the label `baseline` on it. The workflow runs:

1. It reads the label of the baseline from the ticket.
2. It creates the branch `<ticket number>-baseline-<label>`.
3. It runs the loop for every release or library on the head of `main`, up to the check of the
   licences (harness.md).
4. Where every step passes, it opens a pull request with sbom.md. The pull request names the ticket
   and the label it ran with.
5. It names the result in the ticket: the pull request, or the step that failed.

To start it again, the label `baseline` is taken off the ticket and set again.

**Why**: the ticket exists before the run, so there is nothing to check. The ticket is the request
itself, as step 1 of the loop for every change asks (harness.md).

**Rejected**:

- the number of the ticket as input of a workflow started by hand - the workflow has to check that
  the ticket exists, and a wrong number can still name another open ticket
- a ticket opened by the workflow - it would follow the request instead of starting it

**Table baseline faults** - what can make the data of a baseline inconsistent, and what prevents it

| Fault | Prevented by |
|---|---|
| a stranger opens a ticket in a public repo and starts a run | the template does not set the label `baseline`; only whoever may label tickets of the repo sets it (TOOL_27) |
| the label in the ticket changes after the run | the pull request names the label it ran with; the pull request counts, not the ticket (TOOL_27) |
| two runs at once, or the branch or the tag of the label exists already | the run fails (TOOL_27) |
| the ticket is in another repo than the baseline | the workflow makes the baseline of the repo of the ticket only (TOOL_27) |
| `main` changes between the run and the merge, so sbom.md does not describe what is tagged | the pull request fails where its packages are not those of its head (TOOL_28) |

The last fault is a threat as well: a package delivered but missing from sbom.md hides its known
vulnerabilities (REG_EU2024-2847_02).

## Open

`<empty>`

## Change history

| Version | Date | Change | Reason | By |
|---|---|---|---|---|
| 1.1.1 | 2026-10-28 | - | #272 | Claude Opus 5.5 and Mathias Fischer, authors; Mathias Fischer, reviewer |
| 2.1.1 | 2026-10-28 | Documents to PDF | #272 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 3.1.1 | 2026-10-28 | Reviews; CI check; Versions | #255 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 4.1.1 | 2026-10-28 | Documents to PDF | #40 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 5.1.1 | 2026-10-28 | Tool requirements; Documents to PDF; Reviews; CI check; Versions | #291 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 6.1.1 | 2026-10-29 | Documents to PDF | #303 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 7.1.1 | 2026-10-29 | Tool requirements; Launches; Versions | #277 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 8.1.1 | 2026-10-30 | Open; Change history | #323 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 9.1.1 | 2026-10-31 | Release | #343 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 10.1.1 | 2026-10-31 | Release | #343 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 11.1.1 | 2026-10-31 | sbom; licence check | #343 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 12.1.1 | 2026-10-31 | build steps | #342 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 13.1.1 | 2026-10-31 | baseline | #343 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 14.1.1 | 2026-11-01 | baseline | #343 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 15.1.1 | 2026-11-01 | Reviews | #369 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
