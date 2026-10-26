# Tool requirements

The requirements on the tools of our work. Tools are code (see [devops.md](../devops.md)): every
tool realizes an approved requirement of this file. It is written for whoever builds a tool or
reviews its requirement, and for the auditor who asks how the tools are controlled.

The requirements themselves - ids `TOOL-nn` - are in [tool-requirements.yaml](tool-requirements.yaml),
in the format of [formats.md](../formats.md). A tool requirement names as uplink the rule of the
harness (`PROC-nn`) or the regulatory requirement (`REG-nn`) it serves. It is reviewed and approved
like every other requirement (PROC-13). Each section is a feature, the field `feature`.

## Documents to PDF

`TOOL-01` to `TOOL-02`: in [tool-requirements.yaml](tool-requirements.yaml).

A document handed to an auditor shows what it is, which revision, and its status (REG-10).

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

Where it keeps it: in the `documents` section of the YAML beside it, or in `documents.yaml`, with
the values of a requirement (formats.md, section Documents).

## Reviews

`TOOL-03` to `TOOL-04`, and `TOOL-11`: in [tool-requirements.yaml](tool-requirements.yaml).

The tools that create and check a review. `reviewcheck` checks every review result (TOOL-03,
TOOL-11) and runs on every pull request (TOOL-12). `reviewskeleton` writes the skeleton of a review
result (TOOL-04).

## CI check

`TOOL-05` to `TOOL-10`, and `TOOL-12`: in [tool-requirements.yaml](tool-requirements.yaml).

The checks run on every pull request - one tool requirement per check.
