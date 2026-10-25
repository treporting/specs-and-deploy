# Tool requirements

The requirements on the tools of our work. Tools are code (see [devops.md](../devops.md)): every
tool realizes an approved requirement of this file.

The requirements themselves - ids `TOOL-nn` - are in [tool-requirements.yaml](tool-requirements.yaml),
in the format of [formats.md](../formats.md). A tool requirement names as uplink the rule of the
harness (`PROC-nn`) or the regulatory requirement (`REG-nn`) it serves. It is reviewed and approved
like every other requirement (PROC-13). Each section is a feature, the field `feature`.

## Documents to PDF

`TOOL-01` to `TOOL-02`: in [tool-requirements.yaml](tool-requirements.yaml).

A document handed to an auditor shows what it is, which revision, and its status (REG-10).

**Which status** an exported document shows:

| | In the footer | From |
|---|---|---|
| a) a summary of the status of the requirements it holds | "70 draft, 4 review, 2 approved" | counted from the YAML |
| b) a status of the document computed from them | "approved" when all are approved, else "draft" | derived from a) |
| c) **a status of the document itself** | "draft", "released", ... | an attribute of its own - in the model a document (Artifact) has `status` and `mstatus` |

**Decision**: c). A document has a status of its own, as the model has it for a document, and it
works for every document - also for one that holds no requirements, such as devops.md or
harness.md, where a) and b) would give none.

*Open*: where a document keeps its own status, and which values it takes.
