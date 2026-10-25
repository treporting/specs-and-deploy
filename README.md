# specs-and-deploy

The specifications of treporting, and its integration, deployment and system tests.

User needs, requirements, architecture and the way of working of
[treporting](https://github.com/treporting) - a tool that reports traceability and maturity
from the systems a project keeps its work in, and only reads them.

## Files

The data of every level is in YAML, what explains it in Markdown beside it - see
[formats.md](formats.md).

| File | What it holds | Ids |
|---|---|---|
| [harness.md](harness.md) | how we work: the process we follow (ISO/IEC/IEEE 12207, 29119), our levels, the rules of our work, the loops for every change and every release | `PROC-nn` |
| [tools/](tools/) | the tools of our work, in Go; where, why and how they are built | - |
| [documents.yaml](documents.yaml) | the status of the documents that have no YAML of requirements beside them | - |
| [devops.md](devops.md) | the environment of our work: the tools, and the rule that tools are code | - |
| [definitions.md](definitions.md) | the terms: system, item, content, work product, release, library, baseline, traceability | - |
| [architecture/attributes.md](architecture/attributes.md) | every attribute of the model: meaning, origin, original, intermediate and mapped values, how it is carried down | - |
| [formats.md](formats.md) | how requirements, annotations and tests are written down, and what the CI checks; the schemas in [schemas/](schemas/) | - |
| [standards.md](standards.md) | the process, safety and security standards per domain | - |
| [requirements/user-needs.yaml](requirements/user-needs.yaml), [.md](requirements/user-needs.md) | what each role needs, without a solution (r1) | `UN-<role>-nn` |
| [requirements/regulatory-requirements.yaml](requirements/regulatory-requirements.yaml), [.md](requirements/regulatory-requirements.md) | what standards and laws ask | `REG-nn` |
| [requirements/system-requirements.yaml](requirements/system-requirements.yaml), [.md](requirements/system-requirements.md) | what the system has to do, seen as one box (r2); each section is a feature | `SYS-nn` |
| [requirements/tool-requirements.yaml](requirements/tool-requirements.yaml), [.md](requirements/tool-requirements.md) | what the tools of our work have to do; tools are code | `TOOL-nn` |
| [architecture/system-architecture.yaml](architecture/system-architecture.yaml), [.md](architecture/system-architecture.md) | the subsystems, their interfaces, the rationale and the decisions, and which subsystem realizes which requirement (r3) | `ARC-nn` |
| [reviews/](reviews/) | the criteria of the reviews of requirements, and their results | `RC`, `RR` |
| [safety/hazards.yaml](safety/hazards.yaml), [hazard-and-risk-analysis.md](safety/hazard-and-risk-analysis.md) | the hazards treporting can cause (h1) and the measures that control them (h2) | `HAZ-nn`, `HAZC-nn` |

The software requirements (r5, `SR-<subsystem>-nn`) live in the repo of their subsystem, next to
the code that realizes them.

## How the ids hang together

```
HAZ, THR (h1)  →  HAZC, THRC (h2)  →  UN (r1) ─┐
                                   REG (n1) ─┼─→  SYS (r2)  →  ARC (r3)  →  SR (r5)  →  code, tests
                                              ┘
THRC of a threat to our work  →  PROC, a rule of the harness
```

*Open:* the threats (`THR-nn`) and their controls (`THRC-nn`) are not written down yet.

Every need and every regulatory requirement that treporting realizes has at least one system
requirement. Every requirement below the top names its uplinks, or gives its rationale where it has
none, and is allocated to what realizes it on the level below.

## Way of working

Every change has a ticket in the [project treporting](https://github.com/orgs/treporting/projects/3)
and a branch of its own, and reaches `main` only through a pull request - see the rules in
[harness.md](harness.md).

## Licence

[Apache-2.0](LICENSE).
