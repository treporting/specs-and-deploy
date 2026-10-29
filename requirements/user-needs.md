# User needs (r1)

What each role needs from treporting, without naming a solution. The system requirements (r2)
are derived from these.

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
`safety/hazard-and-risk-analysis.md`; the analysis of the tool's failures in the tool qualification traces to these needs.

`UN_PSF_01` to `UN_PSF_06`: in [user-needs.yaml](user-needs.yaml).

## Open

- **Safety of the product**: the tool qualification, whose analysis of the tool's failures traces
  to these needs.
