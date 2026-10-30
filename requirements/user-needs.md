# User needs (r1)

What each role needs from treporting, without naming a solution. The needs are the reason for the
tool itself. The document is written for everyone, above all for whoever derives the system
requirements (r2) from them. The roles are those treporting assumes in an organization that may use
it - not the roles of our own work, which harness.md sets in its section Roles.

The needs themselves - ids `UN_<role>_nn` - are in [user-needs.yaml](user-needs.yaml), in the format
of [formats.md](../formats.md); the field `function` of a need is the short form of its role.

## Roles

**Table roles** - the roles treporting assumes in an organization, and the short form each is
known by in the field `function`

| Short form | Role | Description |
|---|---|---|
| PM | Project Manager | Plans and steers the development of a product - its scope, schedule and resources. |
| RM | Release Manager | Decides whether a release is ready to be delivered, and answers for that decision. |
| CM | Configuration Manager | Identifies the configuration of a product, controls the changes to it and records its baselines. |
| QA | Quality Assurance | Checks, independently of the project, whether its processes and work products follow what was agreed. |
| SAF | Safety | Plans and coordinates the safety activities of a project, and answers for its hazards being controlled. |
| ASR | Assessor | Judges independently, from outside the project, whether it meets the standards chosen - an assessor or an auditor. |
| DEV | Developer | Turns requirements into design and code. |
| TST | Tester | Specifies and runs the tests, and records their results. |
| OPS | Operator | Installs, runs and maintains treporting in an organization, and keeps its data. |
| CFG | Configurator | Sets treporting up for the systems of a customer. |
| ALL | All roles | What every role above needs. |
| PSF | Safety of the product | What every role relies on when it decides on what treporting shows - see the section below. |

## Safety of the product - PSF

What every role relies on when it decides on what treporting shows. treporting controls
nothing, but a gap it hides can lead to a wrong release. The hazards behind them are in
`safety/hazard-and-risk-analysis.md`. If treporting gets a tool qualification, its analysis of the
tool's failures traces to these needs.

## Open

`<empty>`

## Change history

| Version | Date | Change | Reason | By |
|---|---|---|---|---|
| 1.1.1 | 2026-10-30 | - | #325 | Claude Opus 5.5 and Mathias Fischer, authors; Mathias Fischer, reviewer |
| 2.1.1 | 2026-10-30 | Safety of the product - PSF | #338 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
