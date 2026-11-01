# Harness

How treporting is developed: the process, its levels, the roles, the rules of the work and the loops
for every change. It is written for everyone who works on treporting, and for the auditor who checks
how it was made.

## Setup

The settings of Claude Code for all repos of treporting are kept in `~/git/treporting/.claude/`,
one level above the repos.

## Process

treporting is a pure software project and is developed following the process of the domain
*software* (see [standards.md](standards.md)):

- **ISO12207** is the life cycle, **ISO29119** gives the test levels.
- ISO12207 knows no safety and no security, so companion standards are added:
  - **safety analysis**: the hazard and risk analysis by the method of IEC61508-1, clause 7.4, in
    [safety/hazard-and-risk-analysis.md](safety/hazard-and-risk-analysis.md) (HAZ);
  - **security**: IEC62443-4-1 for the secure development, IEC62443-3-2 for the analysis of the
    threats (THR);
  - **assurance**: ISO15026 is proposed for the case why the tool can be relied on, in
    the tool qualification.
- **project risks** are handled by the risk management process of ISO12207, clause 6.3.4.
- REG_ISO26262-8_01, REG_IEC61508-3_01, REG_DO178C_01, REG_DO330_01 and REG_ISO13485_01 are
  requirements on the tool from the domains of its users, not our process.

### Levels

**Table levels** - the levels of the work, and the clause of each

| Level | Here | Reference |
|---|---|---|
| h1 | hazards (HAZ) and threats (THR), told apart by their prefix - a trace starts here or at r1 | the method of IEC61508-1, clause 7.4; IEC62443-3-2 |
| h2 | the measure that controls a hazard (HAZC) or a threat (THRC), and asks for r2 | - |
| n1 | the clauses of the standards and laws to be met: regulatory requirements (REG) | ISO12207, clause 6.4.2 - stakeholder needs and requirements definition |
| r1 | user needs (UN) | ISO12207, clause 6.4.2 - stakeholder needs and requirements definition |
| r2 | system requirements (SYS) | ISO12207, clause 6.4.3 - system requirements definition |
| r3 | system architecture (ARC) | ISO12207, clause 6.4.4 - architecture definition |
| r4 | the system requirements allocated to the subsystems - derived directly from r2 by the allocation in the system architecture, with no ids of their own | ISO12207, clause 6.4.4 - architecture definition |
| r5 | software requirements (SR) per subsystem | ISO12207, clause 6.4.3 - software requirements definition |
| r6 | software architecture of a subsystem: its components | ISO12207, clause 6.4.4 - architecture definition |
| r7 | software design of a subsystem: its modules and units | ISO12207, clause 6.4.5 - design definition |

In the level view of the reporter, RISK+SYS+SW, h1 and h2 are RISK, r1 to r3 with the system
tests SYS, r5 to r7 with the code and the software tests SW. n1 is read by the compliance
view, not by a trace.

### Test levels

**Table test levels** - the levels of the tests, and the term of each in ISO29119

| Level | Term in ISO29119 | What | Data |
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

**Table maturity** - the steps of the chain, the clause of each, and its group

| Order | Object | Reference | Group |
|---|---|---|---|
| 1 | stakeholder requirement | ISO12207, clause 6.4.2 | RG-SYS-R |
| 2 | system requirement | ISO12207, clause 6.4.3 | RG-SYS-R |
| 3 | system architecture | ISO12207, clause 6.4.4 | RG-SYS-R |
| 4 | software requirement | ISO12207, clause 6.4.3 | RG-SW-R |
| 5 | software architecture - the components | ISO12207, clause 6.4.4 | RG-SW-R |
| 6 | software design - the modules | ISO12207, clause 6.4.5 | RG-SW-R |
| 7 | source code, one step per code coverage level | ISO12207, clause 6.4.7 | RG-SC |
| 8 | object code, one step per code coverage level | ISO12207, clause 6.4.7 | RG-OC |
| 9 | test case, test result of the unit test (MT) | ISO29119 | RG-SW-TC, RG-SW-TR |
| 10 | test case, test result of the integration test (SwIT) | ISO29119 | RG-SW-TC, RG-SW-TR |
| 11 | test case, test result of the subsystem test (SwVal) | ISO29119 | RG-SW-TC, RG-SW-TR |
| 12 | test case, test result of the system integration test (SysIT) | ISO29119 | RG-SYS-TC, RG-SYS-TR |
| 13 | test case, test result of the system test (SysQT) | ISO29119 | RG-SYS-TC, RG-SYS-TR |
| 14 | test case, test result of the acceptance test (Val) | ISO29119 | RG-SYS-TC, RG-SYS-TR |

With one code coverage level each, the chain has 20 steps.

### Platform and launches

The platform is treporting, this organization on GitHub. The launches are the deliveries: of the
tools our work needs, of the demo environment, and of a selfhosted environment. An object (requirement
etc.) names the launches it is planned for in `platform_launch`.

**Table launches** - the launches of each platform, and the release and the library of each

| Platform | Launch | Release | Library |
|---|---|---|---|
| treporting | 1_tools | specs-and-deploy 1.0.0 | – |
| treporting | 1_demo | specs-and-deploy 2.0.0 | reporter-trace 1.0.0<br>viewer-trace 1.0.0<br>… |
| treporting | 2_selfhosted | | |
| treporting | 2_tools | | |
| treporting | 3_tools | | |
| treporting | 4_selfhosted | | |

The values in the table are an example only, made up to show the form (see Open).

## Roles

**Table roles** - who does what

| Role | Who | Task |
|---|---|---|
| Product Owner | the owner of treporting | stands for all roles of the user needs; confirms that a need is correct; approves |
| Author | Claude Code or the Product Owner | writes requirements, architecture and code |
| Reviewer | a person who is not the author: the Product Owner where Claude Code was the author, a second person where the Product Owner was | checks against the criteria; never the author (PROC_15) |
| Approver | the Product Owner | sets `approved` (PROC_13) |

Claude Code may prepare a review: go through the criteria and note findings. A person is
responsible for its result.

The criterion *correct* of a user need is always confirmed by the Product Owner, the only one who
can speak for the roles - also where they are the author. It is the one exception to PROC_15.

**Decision**: the roles are tailored to a small team. Claude Code and the Product Owner write; the
Product Owner holds every other role. Claude Code is never the reviewer: it prepares a review, and a
person decides it.
What the Product Owner writes alone needs a second person as reviewer, or it stays `draft`.

**Why**: a review is a judgement, and a person answers for it. A tailoring is allowed where it is
reasoned and written down, as here.

**Rejected**:

- Claude Code as reviewer where the Product Owner wrote - the judgement would come from a tool, and
  no person would answer for it
- a person of their own for every role - the team is too small

## Rules

The ids are `PROC_nn`. They are kept apart from the ids of the requirements, so that keeping
to the process does not count into the maturity of the product. The rules are kept as data in
[rules.yaml](rules.yaml); the table below is generated from it whenever it changes, and names the
hash of its content, `git hash-object rules.yaml`.

<!-- generated from rules.yaml@e826f63 - do not edit -->
**Table rules** - the rules of the work, generated from rules.yaml@e826f63

| Id | Version | Rule |
|---|---|---|
| PROC_01 | 1.1.1 | The language for tickets, code, code comments and documentation is English only. |
| PROC_02 | 1.1.1 | Every change has a ticket in the org project treporting and a branch `<ticket number>-<slug>`. |
| PROC_03 | 1.1.1 | A change reaches `main` only through a pull request; `main` is protected. |
| PROC_04 | 1.1.1 | A pull request passes the full gate before it is merged to `main`. |
| PROC_05 | 1.1.1 | Every change of code realizes or verifies an approved requirement and names it in `@implements` or `@verifies`. Code is everything that is built or run - also configurations and container files. A schema is no code: it is the contract of the requirement it refines, and names it in `@refines`. |
| PROC_06 | 1.1.1 | A baseline is a signed tag: in this repo (specs-and-deploy) it is a release, in any other repo a library. |
| PROC_07 | 1.1.1 | Every release and every library writes its data into the reference edgedb. |
| PROC_08 | 2.2.1 | Every release has an SBOM, and the licences of its components are checked against Apache-2.0 (REG_EU2024-2847_02, REG_Licences_01). |
| PROC_09 | 2.2.1 | A reported vulnerability is assessed and fixed without delay; one that is actively exploited is notified within 24 hours (REG_EU2024-2847_02, REG_EU2024-2847_03). |
| PROC_10 | 2.2.1 | Every release gets security fixes over the support period it states (REG_EU2024-2847_04). |
| PROC_11 | 2.1.1 | A document names no ticket, except in its change history; what is open, it states as open. The ticket names the document, not the other way round. |
| PROC_12 | 2.1.1 | An object is merged before another names it. |
| PROC_13 | 2.1.1 | Whatever has review criteria of its kind in `reviews/` is approved only when it passed its review against those criteria, and whatever it names as uplink is approved at that moment. A later change of an uplink does not take the approval back. A document a tool writes, with the verification method `generated`, is approved as the tool writes it, without a review. See traceability.md. |
| PROC_14 | 1.1.1 | A change of the content of whatever is approved that changes x of its version - the text of a requirement or an attribute its kind names for x, a document as a whole - sets it back to `draft`. A change of y or z only - a link, another attribute - leaves it `approved`; its merge approves it. See traceability.md. |
| PROC_15 | 1.1.1 | The reviewer of whatever is reviewed is not its author. |
| PROC_16 | 1.1.1 | A pull request closes one ticket and changes one concern. Group things which belong together so they do not contradict. |
| PROC_18 | 2.1.1 | A decision is written down with:<br>- its reason<br>- the alternatives rejected<br>- what a means found, where one was applied to gain more information according to ISO25010:<br>&nbsp;&nbsp;- scenarios of change and growth, after ATAM<br>&nbsp;&nbsp;- one source for each information<br>&nbsp;&nbsp;- a comparison with models<br>&nbsp;&nbsp;- a pre-mortem, after LASR |
| PROC_19 | 1.1.1 | Every test belongs to exactly one test level, read from the place of its file. |
| PROC_20 | 2.1.1 | Every link names the version of the linked object, as `ID@x.y.z`. A link is suspect where the linked object has another x than the link names. A suspect link is reported on every pull request. It blocks a baseline that holds both its ends, compared with the versions in that baseline, not on `main`. See traceability.md. |
<!-- end of generated -->

**Decision**: a rule of the harness is approved when it is merged into `main`. It names no criteria
and passes no review of its own. So a rule can be the uplink of an approved requirement (PROC_13).

**Why**:

- the harness is the central safeguard: it sets how everything else is checked
- it is improved iteratively: rules arise and change from the work
- the merge is an approval with evidence: a rule reaches `main` only by ticket and pull request
  (PROC_02, PROC_03), one concern each (PROC_16); the Product Owner merges, and git keeps who and
  when

**Rejected**: to review the rules of the harness against criteria, as requirements are. The rules
that set how a review is made would then need a review made by those rules - they go round in a
circle; and a formal review of every change would slow the very means by which the process gets
better.

This overcomes the chicken-and-egg problem: the rules of review stand first, and every requirement
is reviewed by them.

## Loop for every change

1. A ticket is the start of every change (PROC_02).
2. A branch is created for the ticket.
3. The change is realized, in commits.
4. If the change touches what has review criteria of its kind (PROC_13) - something new, or a
   change of x of its version (PROC_14) - it needs a review. The author sets it to `review`. The
   change is then finished. A change of y or z only needs no review: the merge approves it.
5. The reviewer checks it against the criteria of its kind. The reviewer is not the author (PROC_15).
6. The review result goes into `reviews/results/`.
7. A finding can be resolved on the branch.
8. A finding that cannot be resolved on the branch fails the review. It becomes one or more tickets
   of its own. Each goes to step 2 again.
9. A finding that shows a missing rule stops the review. The rule is changed in a change of its
   own. Then the review goes on, against the new rule.
10. The status of the change follows the review result: `approved` if passed, `draft` if failed -
    save a change of what is approved: where its review failed, the change is rejected. It is taken
    back on the branch, and only the review result and the tickets reach `main`. The approved
    version stands (traceability.md).
11. A pull request takes the branch to `main` (PROC_03). The CI checks it, with the jobs of devops.md
    (PROC_04).
12. The Product Owner merges. The merge is the approval.
13. The branch is deleted automatically. Its commits stay in `main`.

For code, step 3 means:

- the requirement it realizes is approved (PROC_05)
- it compiles, the linter passes, the unit tests pass
- update documentation
- the quick gate passes before a commit

## Loop for every release or library

1. compile
2. run linter
3. run unit-testing in local
4. run full set of tests
5. report on
   - test coverage
   - structural debt
6. create the SBOM (PROC_08)
7. check the licences of the components (PROC_08)
8. create a signed tag as approval (PROC_06)

*Note*: write the data of the release or library into a separate edgedb, kept as reference data
as soon as possible (PROC_07)

## Open

- `treporting = 1_demo` - **Threats**: not decided whether treporting gets a threat analysis. One
  threat is known: a package delivered but missing from sbom.md hides its known vulnerabilities
  (REG_EU2024-2847_02); TOOL_28 controls it.
- `treporting = 3_tools` - **Process**: not decided whether treporting gets a tool qualification
  (ISO15026).
- `treporting = 1_tools` - **Platform and launches**: not decided whether the reading of table
  launches holds, and with which values - or whether its columns stay empty until the tags exist.

## Change history

| Version | Date | Change | Reason | By |
|---|---|---|---|---|
| 1.1.1 | 2026-10-28 | - | #272 | Claude Opus 5.5 and Mathias Fischer, authors; Mathias Fischer, reviewer |
| 2.1.1 | 2026-10-28 | Process; Rules | #40 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 3.1.1 | 2026-10-28 | Roles; Rules; Loop for every change; Loop for every release or library | #291 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 4.1.1 | 2026-10-29 | Loop for every change | #307 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 5.1.1 | 2026-10-29 | Rules | #309 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 6.1.1 | 2026-10-29 | Platform and launches | #277 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 7.1.1 | 2026-10-29 | Platform and launches; Loop for every release or library; Open | #322 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 8.1.1 | 2026-10-30 | Open | #322 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 9.1.1 | 2026-10-30 | Rules; Change history | #323 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 10.1.1 | 2026-10-30 | Process | #338 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 11.1.1 | 2026-10-30 | Platform and launches | #334 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 12.1.1 | 2026-10-30 | Platform and launches; Loop for every change; Loop for every release or library | #320 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 13.1.1 | 2026-11-01 | Rules | #343 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 14.1.1 | 2026-11-01 | Open | #343 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
