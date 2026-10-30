# User needs (r1)

What each role needs from treporting, without naming a solution. The needs are the reason for the
tool itself. The document is written for everyone, above all for whoever derives the system
requirements (r2) from them. The roles are those treporting assumes in an organization that may use
it - not the roles of our own work, which harness.md sets in its section Roles.

The needs themselves - ids `UN_<role>_nn` - are in [user-needs.yaml](user-needs.yaml), in the format
of [formats.md](../formats.md); each section is a role, the field `function` of a need.

## Project Manager - PM

`UN_PM_01` to `UN_PM_03`: in [user-needs.yaml](user-needs.yaml).

## Release Manager - RM

`UN_RM_01` to `UN_RM_05`: in [user-needs.yaml](user-needs.yaml).

## Configuration Manager - CM

`UN_CM_01` to `UN_CM_04`: in [user-needs.yaml](user-needs.yaml).

## Quality Assurance - QA

`UN_QA_01` to `UN_QA_06`: in [user-needs.yaml](user-needs.yaml).

## Safety - SAF

`UN_SAF_01` to `UN_SAF_03`: in [user-needs.yaml](user-needs.yaml).

## Assessor - ASR

`UN_ASR_01` to `UN_ASR_07`: in [user-needs.yaml](user-needs.yaml).

## Developer - DEV

`UN_DEV_01`: in [user-needs.yaml](user-needs.yaml).

## Tester - TST

`UN_TST_01` to `UN_TST_02`: in [user-needs.yaml](user-needs.yaml).

## Operator - OPS

`UN_OPS_01` to `UN_OPS_13`: in [user-needs.yaml](user-needs.yaml).

## Configurator - CFG

Whoever sets treporting up for the systems of a customer.

`UN_CFG_01` to `UN_CFG_03`: in [user-needs.yaml](user-needs.yaml).

## All roles - ALL

`UN_ALL_01` to `UN_ALL_04`: in [user-needs.yaml](user-needs.yaml).

## Safety of the product - PSF

What every role relies on when it decides on what treporting shows. treporting controls
nothing, but a gap it hides can lead to a wrong release. The hazards they control are in
`safety/hazard-and-risk-analysis.md`. If treporting gets a tool qualification, its analysis of the
tool's failures traces to these needs.

`UN_PSF_01` to `UN_PSF_06`: in [user-needs.yaml](user-needs.yaml).

## Open

`<empty>`

## Change history

| Version | Date | Change | Reason | By |
|---|---|---|---|---|
| 1.1.1 | 2026-10-30 | - | #325 | Claude Opus 5.5 and Mathias Fischer, authors; Mathias Fischer, reviewer |
