# User needs (r1)

What each role needs from treporting, without naming a solution. The needs are the reason for the
tool itself. The document is written for everyone, above all for whoever derives the system
requirements (r2) from them. The roles are those treporting assumes in an organization that may use
it - not the roles of our own work, which harness.md sets in its section Roles.

The needs themselves - ids `UN_<role>_nn` - are in [user-needs.yaml](user-needs.yaml), in the format
of [formats.md](../formats.md); each section is a role, the field `function` of a need.

## Project Manager - PM

Plans and steers the development of a product - its scope, schedule and resources.

## Release Manager - RM

Decides whether a release is ready to be delivered, and answers for that decision.

## Configuration Manager - CM

Identifies the configuration of a product, controls the changes to it and records its baselines.

## Quality Assurance - QA

Checks, independently of the project, whether its processes and work products follow what was
agreed.

## Safety - SAF

Plans and coordinates the safety activities of a project, and answers for its hazards being
controlled.

## Assessor - ASR

Judges independently, from outside the project, whether it meets the standards chosen - an
assessor or an auditor.

## Developer - DEV

Turns requirements into design and code.

## Tester - TST

Specifies and runs the tests, and records their results.

## Operator - OPS

Installs, runs and maintains treporting in an organization, and keeps its data.

## Configurator - CFG

Whoever sets treporting up for the systems of a customer.

## All roles - ALL

What every role above needs.

## Safety of the product - PSF

What every role relies on when it decides on what treporting shows. treporting controls
nothing, but a gap it hides can lead to a wrong release. The hazards they control are in
`safety/hazard-and-risk-analysis.md`. If treporting gets a tool qualification, its analysis of the
tool's failures traces to these needs.

## Open

`<empty>`

## Change history

| Version | Date | Change | Reason | By |
|---|---|---|---|---|
| 1.1.1 | 2026-10-30 | - | #325 | Claude Opus 5.5 and Mathias Fischer, authors; Mathias Fischer, reviewer |
