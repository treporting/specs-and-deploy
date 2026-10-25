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

The versions as installed on 2026-10-25. A tool is listed when the work on treporting needs it.

| Tool | Version | What for |
|---|---|---|
| macOS | 26.4.1 | the development machine |
| VS Code | 1.142.0 | the editor, with extensions - see below |
| Claude Code | 2.1.310 | the AI assistant, as an extension of VS Code; its settings are kept one level above the repos (see harness.md) |
| git | 2.50.1 | version control |
| GitHub CLI (`gh`) | 2.101.0 | tickets, pull requests, the project |
| Go | 1.25.14 | the subsystems in Go; the CI builds with 1.25 as well |
| Node.js, npm | 24.21.0, 11.19.0 | the viewer and its tests |
| Docker Desktop | 4.17.0 - engine 20.10.23, Compose 2.15.1 | the containers, gel |
| Google Chrome | 155.0.8059.86 | renders the export of the documents to PDF |
| GitHub | - | the repos of the organization treporting, its project, protected `main`, pull requests |
| CI | - | *Open*: the checks of formats.md, SBOM and licences |

How to read a version: `sw_vers`, `git --version`, `gh --version`, `go version`, `node --version`,
`npm --version`, `docker --version`, `docker compose version`; VS Code and its extensions in
*Help > About* and the view *Extensions*.

### Extensions of VS Code

| Extension | Version | What for |
|---|---|---|
| Claude Code (Anthropic) | 2.1.310 | the AI assistant |
| Markdown PDF (yzane) | 2.2.0 | the export of the documents to PDF |
| Go (golang) | 0.58.0 | Go: language server, tests, debugging |
| Container Tools (Microsoft) | 2.5.3 | containers and compose files |
| Dev Containers (Microsoft) | 0.469.0 | development inside a container |
| GitHub Actions (GitHub) | 0.32.3 | the workflows of the CI |

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
