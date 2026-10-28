# System requirements (r2)

What treporting has to do, seen as one box - viewer, reporter and storage together. A requirement
names as uplinks the user needs (r1), the regulatory requirements and the decisions of the
architecture (`ARCD_nn`) it comes from. Where none stands above it - a requirement derived from the
view of the system - it gives its rationale instead (ISO/IEC/IEEE 29148, 5.2.8). Every need and
every regulatory requirement that treporting realizes has at least one requirement. How the
system does it is left to the architecture (r3).

The requirements themselves - ids `SYS_nn` - are in [system-requirements.yaml](system-requirements.yaml),
in the format of [formats.md](../formats.md); this file holds what explains them.

Each section is a feature, named by a short keyword - the field `feature` of a requirement. A
requirement carries the functions of the needs it comes from - the roles of `user-needs.md`.

## Trace

`SYS_01` to `SYS_08`: in [system-requirements.yaml](system-requirements.yaml).

## Filter

`SYS_09` to `SYS_11`: in [system-requirements.yaml](system-requirements.yaml).

## Hazards

`SYS_12`: in [system-requirements.yaml](system-requirements.yaml).

## Standards

`SYS_13` to `SYS_14`: in [system-requirements.yaml](system-requirements.yaml).

## Baselines

`SYS_15` to `SYS_17`: in [system-requirements.yaml](system-requirements.yaml).

## Provenance

`SYS_18` to `SYS_22`: in [system-requirements.yaml](system-requirements.yaml).

## Integrity

treporting controls nothing, but it can fail to show a gap (tool impact TI2 in ISO 26262-8,
clause 11). Whether it has to be qualified depends on how surely its own errors are detected: with
high confidence (TD1) the tool confidence level is TCL1 and no qualification is needed. IEC 61508
(class T2) and DO-330 argue alike. The requirements of this feature are what makes errors of an
import detectable.

The check has to be independent of the reading: a converter that counts what it reads shares every
error of the reading with its count. So the source is counted by a way of its own, and the storage
is counted after the import, not by the converter. Four levels together give the evidence:

| Level | What | Finds |
|---|---|---|
| A. reconciliation at every import | count at four points - in the source by a query of its own (the object count of a DOORS module, the total of a codebeamer query, a `git grep` for `@implements`), read, refused with a reason, stored; source = read + refused, read = stored; per kind of object, per type of link, per state | missing objects, missing links, states counted wrong |
| B. content check | a fingerprint of the fields that matter, per object, on the side of the source and on the side of the storage | attributes swapped, links to the wrong object, an unknown value turned into zero or empty |
| C. known-answer test per release | data sets of known content, with faults built in - missing objects, missing links, unknown values, spellings that differ; every fault has to be reported | that the checks find what they are for - the evidence for TD1 |
| D. reference edgedb | treporting reads its own repos; the expected counts come independently from git | real data instead of generated |

Equal counts do not prove the content right - that is level B. That the checks find anything at all
is shown by level C only.

`SYS_23` to `SYS_27`: in [system-requirements.yaml](system-requirements.yaml).

## Errors

`SYS_28` to `SYS_30`: in [system-requirements.yaml](system-requirements.yaml).

## Export

`SYS_31` to `SYS_32`: in [system-requirements.yaml](system-requirements.yaml).

## Usability

`SYS_33` to `SYS_34`: in [system-requirements.yaml](system-requirements.yaml).

## Operation

`SYS_35` to `SYS_38`: in [system-requirements.yaml](system-requirements.yaml).

## Audit

`SYS_39` to `SYS_44`: in [system-requirements.yaml](system-requirements.yaml).

## Maturity

The maturity of a requirement is how far down the longest chain of the V-model it has got: the
share of the steps of the chain it has reached. The number of steps is not fixed: each process
reference model has its own chain, and source code and object code have a step for each code
coverage level coded, such as sc1 function and sc2 statement coverage, or oc1 object and oc2 assembler
coverage. The table shows the chain of ASPICE; with one code coverage level each it has 20 steps, so
each step is 5 %. A step whose level is empty is skipped and the others count more,
so that 100 % can always be reached.

| Order | Object | ASPICE | Group |
|---|---|---|---|
| 1 | stakeholder requirement | SYS.1 | RG-SYS-R |
| 2 | system requirement | SYS.2 | RG-SYS-R |
| 3 | system architecture | SYS.3 | RG-SYS-R |
| 4 | software requirement | SWE.1 | RG-SW-R |
| 5 | software architecture | SWE.2 | RG-SW-R |
| 6 | detailed design | SWE.3 | RG-SW-R |
| 7 | source code, one step per code coverage level (sc1, sc2, ...) | SWE.3 | RG-SC |
| 8 | object code, one step per code coverage level (oc1, oc2, ...) | build | RG-OC |
| 9 | test case, test result of the unit test | SWE.4 | RG-SW-TC, RG-SW-TR |
| 10 | test case, test result of the software integration test | SWE.5 | RG-SW-TC, RG-SW-TR |
| 11 | test case, test result of the software qualification test | SWE.6 | RG-SW-TC, RG-SW-TR |
| 12 | test case, test result of the system integration test | SYS.4 | RG-SYS-TC, RG-SYS-TR |
| 13 | test case, test result of the system qualification test | SYS.5 | RG-SYS-TC, RG-SYS-TR |
| 14 | test case, test result of the validation | VAL.1 | RG-SYS-TC, RG-SYS-TR |

`SYS_45` to `SYS_50`: in [system-requirements.yaml](system-requirements.yaml).

## Performance

The non-functional requirements. Their figures are measured first on the prototype with generated
data, then set as targets; until then they read *open*. A requirement of another feature that
depends on a figure refers to the requirement here.

The sizes of a data set:

| Size | Requirements | Links | Test results | Baselines |
|---|---|---|---|---|
| small | open | open | open | open |
| medium | open | open | open | open |
| large | open | open | open | open |

`SYS_51` to `SYS_56`: in [system-requirements.yaml](system-requirements.yaml).

## Constraints

`SYS_57`: in [system-requirements.yaml](system-requirements.yaml).

## Licensing

Priority **later**: treporting has no licensed feature yet - everything is under Apache-2.0.
These requirements keep the way open for features developed for customers. A feature
meant to be licensed must not be published under Apache-2.0 before, or it stays free for good.

`SYS_58` to `SYS_59`: in [system-requirements.yaml](system-requirements.yaml).

## Configuration

`SYS_60` to `SYS_62`: in [system-requirements.yaml](system-requirements.yaml).

## Delivery

`SYS_63` to `SYS_67`: in [system-requirements.yaml](system-requirements.yaml).

## Derived values

`SYS_68` to `SYS_69`: in [system-requirements.yaml](system-requirements.yaml).

## Plausibility

`SYS_70` to `SYS_72`: in [system-requirements.yaml](system-requirements.yaml).

## Verification

`SYS_73` to `SYS_76`: in [system-requirements.yaml](system-requirements.yaml).

Every requirement has a status - `draft`, `review`, `approved`, `obsolete` - and a verification
method after ISO/IEC/IEEE 29148. The review of a requirement checks the requirement itself
against criteria; the verification method says how its realization is shown to meet it.

## Open

- **Figures**: the measured and the target figures in the tables of [Performance](#performance).
