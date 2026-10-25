# devops

The environment of our work: the tools we build treporting with, and how they are ruled.

## Tools are code

**Decision**: PROC-05 holds for the tools of our work as for the code of the product. Editor
settings, scripts and the CI are code. Every tool realizes an approved requirement. There is no
part where something is built without one.

**Why**: an auditor asks how the tools that build the product are controlled. With this rule the
answer is the same as for the product: by requirements, reviewed and approved before the change.

**Rejected**: to narrow PROC-05 to the code of the product, and rule the tools by PROC rules alone.
It is lighter, but it leaves a part without requirements.

**Consequence**: the requirements on our tools are kept in `requirements/tool-requirements.yaml`,
with the ids `TOOL-nn`. A tool requirement names as uplink the PROC rule or the regulatory
requirement it serves. It is reviewed and approved like every other requirement (PROC-13).

## Environments

| Tool | What for |
|---|---|
| a Mac | the development machine |
| VS Code | the editor, with extensions - see below |
| Claude Code | the AI assistant; its settings are kept one level above the repos (see harness.md) |
| GitHub | the repos of the organization treporting, its project, protected `main`, pull requests |
| CI | *Open*: the checks of formats.md, SBOM and licences |

### Extensions of VS Code

| Extension | What for |
|---|---|
| Markdown PDF (yzane) | the export of the documents to PDF |

## Documents to PDF

The documents are exported to PDF with Markdown PDF.

- It renders with the Chrome of the machine.
- Mermaid diagrams are drawn; the script of Mermaid is loaded from a CDN, the diagram stays on the
  machine.
- PlantUML diagrams are drawn by a local PlantUML server only, never the public one: a diagram must
  not leave the machine.
- An exported PDF is never committed.

*Open*: the settings of the export - header, footer, margins - and revision and status in header
and footer. Each waits for its approved tool requirement.

## Build and deployment

*Open*: the containers in `deploy/`, built from the parts the subsystems deliver; the demo.
