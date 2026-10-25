# Harness

## Setup

The settings of Claude Code for all repos of treporting are kept in `~/git/treporting/.claude/`,
one level above the repos.

## Process

treporting is a pure software project and is developed following the process of the domain
*software* (see [standards.md](standards.md)):

- **ISO/IEC/IEEE 12207** is the life cycle, **ISO/IEC/IEEE 29119** gives the test levels.
- 12207 knows no safety and no security, so companion standards are added:
  - **safety analysis**: the hazard and risk analysis by the method of IEC 61508-1, 7.4, in
    [safety/hazard-and-risk-analysis.md](safety/hazard-and-risk-analysis.md) (HAZ);
  - **security**: IEC 62443-4-1 for the secure development, IEC 62443-3-2 for the analysis of the
    threats (THR);
  - **assurance**: ISO/IEC/IEEE 15026 is proposed for the case why the tool can be relied on, in
    the tool qualification.
- **project risks** are handled by the risk management process of 12207, 6.3.4.
- REG-01 to REG-04 are requirements on the tool from the domains of its users, not our process.

### Levels

| Level | Here | ISO/IEC/IEEE 12207 |
|---|---|---|
| h1 | hazards (HAZ) and threats (THR), told apart by their prefix - a trace starts here or at r1 | - (method of IEC 61508-1, 7.4; IEC 62443-3-2) |
| h2 | the measure that controls a hazard (HAZC) or a threat (THRC), and asks for r1 | - |
| n1 | the clauses of the standards and laws to be met: regulatory requirements (REG) | 6.4.2 stakeholder needs and requirements definition |
| r1 | user needs (UN) | 6.4.2 stakeholder needs and requirements definition |
| r2 | system requirements (SYS) | 6.4.3 system requirements definition |
| r3 | system architecture (ARC) | 6.4.4 architecture definition |
| r4 | the system requirements allocated to the subsystems - derived directly from r2 by the allocation in the system architecture, with no ids of their own | 6.4.4 architecture definition |
| r5 | software requirements (SR) per subsystem | 6.4.3 software requirements definition |
| r6 | software architecture of a subsystem: its components | 6.4.4 architecture definition |
| r7 | software design of a subsystem: its modules and units; it names how a requirement is verified (`verification_method`) | 6.4.5 design definition |

In the level view of the reporter, RISK+SYS+SW, h1 and h2 are RISK, r1 to r3 with the system
tests SYS, r5 to r7 with the code and the software tests SW. n1 is read by the compliance
view, not by a trace.

### Test levels

| Level | ISO/IEC/IEEE 29119 | What | Data |
|---|---|---|---|
| MT | unit test | a function or a package | made in the test |
| SwIT | integration test | a subsystem, through its API and its model | synthetic, in memory |
| SwVal | subsystem test | a subsystem against its SR, through its API and its model | from gel |
| SysIT | system integration test | the containers built and working together | - |
| SysQT | system test | the system against SYS | the containers with demo data |
| Val | acceptance test | the system against the user needs | the containers with reference data |

A software test names SR in `@verifies`, a system test SYS.

### Maturity of our own data

The data of treporting itself (the reference edgedb) is measured by this chain:

| Order | Object | ISO/IEC/IEEE 12207, 29119 | Group |
|---|---|---|---|
| 1 | stakeholder requirement | 6.4.2 | RG-SYS-R |
| 2 | system requirement | 6.4.3 | RG-SYS-R |
| 3 | system architecture | 6.4.4 | RG-SYS-R |
| 4 | software requirement | 6.4.3 | RG-SW-R |
| 5 | software architecture - the components | 6.4.4 | RG-SW-R |
| 6 | software design - the modules | 6.4.5 | RG-SW-R |
| 7 | source code, one step per code coverage level | 6.4.7 | RG-SC |
| 8 | object code, one step per code coverage level | 6.4.7 | RG-OC |
| 9 | test case, test result of the unit test (MT) | 29119 | RG-SW-TC, RG-SW-TR |
| 10 | test case, test result of the integration test (SwIT) | 29119 | RG-SW-TC, RG-SW-TR |
| 11 | test case, test result of the subsystem test (SwVal) | 29119 | RG-SW-TC, RG-SW-TR |
| 12 | test case, test result of the system integration test (SysIT) | 29119 | RG-SYS-TC, RG-SYS-TR |
| 13 | test case, test result of the system test (SysQT) | 29119 | RG-SYS-TC, RG-SYS-TR |
| 14 | test case, test result of the acceptance test (Val) | 29119 | RG-SYS-TC, RG-SYS-TR |

With one code coverage level each, the chain has 20 steps.

## Rules

The ids are `PROC-nn`. They are kept apart from the ids of the requirements, so that keeping
to the process does not count into the maturity of the product.

| Id | Rule |
|---|---|
| PROC-01 | The language for tickets, code, code comments and documentation is English only. |
| PROC-02 | Every change has a ticket in the org project treporting and a branch `<ticket number>-<slug>`. |
| PROC-03 | A change reaches `main` only through a pull request; `main` is protected. |
| PROC-04 | A pull request passes the full gate before it is merged to `main`. |
| PROC-05 | Every change of code realizes or verifies an approved requirement and names it in `@implements` or `@verifies`. Code is everything that is built or run - also schemas, configurations and container files. |
| PROC-06 | A baseline is a signed tag: in this repo (specs-and-deploy) it is a release, in any other repo a library. |
| PROC-07 | Every release and every library writes its data into the reference edgedb. |
| PROC-08 | Every release has an SBOM, and the licences of its components are checked against Apache-2.0 (REG-15, REG-18). |
| PROC-09 | A reported vulnerability is assessed and fixed without delay; one that is actively exploited is notified within 24 hours (REG-15, REG-16). |
| PROC-10 | Every release gets security fixes over the support period it states (REG-17). |
| PROC-11 | A document names no ticket; what is open, it states as open. The ticket names the document, not the other way round. |
| PROC-12 | A level is merged before the level below names it: user needs and regulatory requirements, system requirements, architecture, software requirements, code. |
| PROC-13 | A requirement is approved only when it passed its review against the criteria of its level, and every requirement it names as uplink is approved. |
| PROC-14 | A change of the text of an approved requirement sets it back to `draft`. |
| PROC-15 | The reviewer of a requirement is not its author. |

## Loop for every code change

1. the requirement the change realizes is approved (PROC-05)
2. compile
3. run linter
4. run unit-testing in local
5. update documentation (swagger, *.md)
6. gate checks - the quick gate before a commit, the full gate before a merge to main (PROC-04)
7. commit, pull request, merge (PROC-03)
8. provide feedback

## Loop for every review

1. set the requirements to be reviewed to `review`, by a pull request
2. the reviewer - not the author (PROC-15) - checks them against the criteria of their level
3. write the result into a file of its own in `reviews/`: who, when, the revision of every
   requirement reviewed, per requirement and criterion passed or failed, and the verdict - one
   result may cover many requirements
4. pull request for the result - it checks the form of the evidence, not the requirements
5. pull request that sets the passed requirements to `approved` (PROC-13); the failed ones go
   back to `draft`

## Loop for every release or library

1. compile
2. run linter
3. run unit-testing in local
4. run spec-testing on test environment
5. report on
   - test coverage
   - structural debt
   - requirements
6. create the SBOM and check the licences of the components (PROC-08)
7. provide feedback
8. create a signed tag as approval (PROC-06)
9. write the data of the release or library into a separate edgedb, kept as reference
   data - the real data of treporting itself, used instead of synthetic data to show that
   it works in practice (PROC-07)
