# specs-and-deploy

The specifications of treporting, and its integration, deployment and system tests.

User needs, requirements, architecture and the way of working of
[treporting](https://github.com/treporting) - a tool that reports traceability and maturity
from the systems a project keeps its work in, and only reads them. This page is written for
whoever comes to the repo first, and points to where each thing stands.

## Files

The data of every level is in YAML, what explains it in Markdown beside it - see
[formats.md](formats.md).

**Table files** - the files of the repo, what each holds, and its ids

| File | What it holds | Ids |
|---|---|---|
| [harness.md](harness.md) | how we work: the process we follow (ISO12207, ISO29119), our levels, the rules of our work, the loops for every change and every release | `PROC_nn` |
| [tools/](tools/) | the tools of our work, in Go; where, why and how they are built; their software requirements in [software-requirements.yaml](tools/software-requirements.yaml) | `SR_nn` |
| [documents.yaml](documents.yaml) | the type, level, version and status of every document | - |
| [devops.md](devops.md) | the environment of our work: the tools, and the rule that tools are code | - |
| [definitions.md](definitions.md) | the terms: system, item, content, work product, release, library, baseline, traceability | - |
| [architecture/attributes.md](architecture/attributes.md) | every attribute of the model: meaning, origin, original, intermediate and mapped values, how it is carried down | - |
| [formats.md](formats.md) | how requirements, documents, reviews, annotations and tests are written down, and what the CI checks; the schemas in [schemas/](schemas/) | - |
| [standards.md](standards.md) | the process, safety and security standards per domain | - |
| [requirements/user-needs.yaml](requirements/user-needs.yaml), [.md](requirements/user-needs.md) | what each role needs, without a solution (r1) | `UN_<role>_nn` |
| [regulatory/](regulatory/), [regulatory-requirements.md](regulatory/regulatory-requirements.md) | what standards and laws ask, one file per module | `REG_<module>_nn` |
| [requirements/system-requirements.yaml](requirements/system-requirements.yaml), [.md](requirements/system-requirements.md) | what the system has to do, seen as one box (r2); each section is a feature | `SYS_nn` |
| [requirements/tool-requirements.yaml](requirements/tool-requirements.yaml), [.md](requirements/tool-requirements.md) | what the tools of our work have to do; tools are code | `TOOL_nn` |
| [architecture/system-architecture.yaml](architecture/system-architecture.yaml), [.md](architecture/system-architecture.md) | the subsystems, their interfaces, the rationale and the decisions, and which subsystem realizes which requirement (r3) | `ARC_nn` |
| [reviews/](reviews/) | the criteria of the reviews of requirements and documents, and their results | `RC`, `RR` |
| [safety/hazards.yaml](safety/hazards.yaml), [hazard-and-risk-analysis.md](safety/hazard-and-risk-analysis.md) | the hazards treporting can cause (h1) and the measures that control them (h2) | `HAZ_nn`, `HAZC_nn` |

The software requirements (r5, `SR_nn`) live in the repo of their subsystem, next to the code that
realizes them; an id is unique together with that repo and its document (formats.md, section Ids).

## How the ids hang together

```
HAZ, THR (h1)  →  HAZC, THRC (h2)  →  UN (r1) ──┐
                                      REG (n1) ─┴─→  SYS (r2)  →  ARC (r3)  →  SR (r5)  →  code, tests

THRC of a threat to our work  →  PROC, a rule of the harness
```

Every need and every regulatory requirement that treporting realizes has at least one system
requirement. Every requirement below the top names its uplinks, or gives its rationale where it has
none, and is allocated to what realizes it on the level below.

## Way of working

Every change has a ticket in the [project treporting](https://github.com/orgs/treporting/projects/3)
and a branch of its own, and reaches `main` only through a pull request - see the rules in
[harness.md](harness.md).

## Licence

[Apache-2.0](LICENSE).

## Open

`<empty>`

## Change history

| Version | Date | Change | Reason | By |
|---|---|---|---|---|
| 1.1.1 | 2026-10-28 | - | #272 | Claude Opus 5.5 and Mathias Fischer, authors; Mathias Fischer, reviewer |
| 2.1.1 | 2026-10-28 | Files | #40 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 3.1.1 | 2026-10-28 | Files; How the ids hang together | #291 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 4.1.1 | 2026-10-29 | Files | #303 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 5.1.1 | 2026-10-29 | How the ids hang together; Licence; Open | #322 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 6.1.1 | 2026-10-30 | Open | #322 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 7.1.1 | 2026-10-30 | Change history | #323 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
