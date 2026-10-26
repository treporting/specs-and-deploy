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
- REG-01 to REG-04 are requirements on the tool from the domains of its users, not our process.

### Levels

**Table levels** - the levels of the work, and the clause of each

| Level | Here | Reference |
|---|---|---|
| h1 | hazards (HAZ) and threats (THR), told apart by their prefix - a trace starts here or at r1 | the method of IEC61508-1, clause 7.4; IEC62443-3-2 |
| h2 | the measure that controls a hazard (HAZC) or a threat (THRC), and asks for r1 | - |
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

## Roles

**Table roles** - who does what

| Role | Who | Task |
|---|---|---|
| Product Owner | the owner of treporting | stands for all roles of the user needs; confirms that a need is correct; approves |
| Author | Claude Code or the Product Owner | writes requirements, architecture and code |
| Reviewer | a person who is not the author: the Product Owner where Claude Code was the author, a second person where the Product Owner was | checks against the criteria; never the author (PROC-15) |
| Approver | the Product Owner | sets `approved` (PROC-13) |

Claude Code may prepare a review: go through the criteria and note findings. A person is
responsible for its result.

The criterion *correct* of a user need is always confirmed by the Product Owner, the only one who
can speak for the roles - also where they are the author. It is the one exception to PROC-15.

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

The ids are `PROC-nn`. They are kept apart from the ids of the requirements, so that keeping
to the process does not count into the maturity of the product.

**Table rules** - the rules of the work

| Id | Rule |
|---|---|
| PROC-01 | The language for tickets, code, code comments and documentation is English only. |
| PROC-02 | Every change has a ticket in the org project treporting and a branch `<ticket number>-<slug>`. |
| PROC-03 | A change reaches `main` only through a pull request; `main` is protected. |
| PROC-04 | A pull request passes the full gate before it is merged to `main`. |
| PROC-05 | Every change of code realizes or verifies an approved requirement and names it in `@implements` or `@verifies`. Code is everything that is built or run - also configurations and container files. A schema is no code: it is the contract of the requirement it refines, and names it in `@refines`. |
| PROC-06 | A baseline is a signed tag: in this repo (specs-and-deploy) it is a release, in any other repo a library. |
| PROC-07 | Every release and every library writes its data into the reference edgedb. |
| PROC-08 | Every release has an SBOM, and the licences of its components are checked against Apache-2.0 (REG-15, REG-18). |
| PROC-09 | A reported vulnerability is assessed and fixed without delay; one that is actively exploited is notified within 24 hours (REG-15, REG-16). |
| PROC-10 | Every release gets security fixes over the support period it states (REG-17). |
| PROC-11 | A document names no ticket; what is open, it states as open. The ticket names the document, not the other way round. |
| PROC-12 | A level is merged before the level below names it: user needs and regulatory requirements, system requirements, architecture, software requirements, code. |
| PROC-13 | Whatever has review criteria of its kind in `reviews/` - today requirements and documents - is approved only when it passed its review against those criteria, and whatever it names as uplink is approved. |
| PROC-14 | A change of the content of whatever is approved - the text of a requirement, a document as a whole - sets it back to `draft`. |
| PROC-15 | The reviewer of whatever is reviewed is not its author. |
| PROC-16 | A pull request closes one ticket and changes one concern; a ticket that touches data, rules and explanations at once is split into several pull requests. |
| PROC-17 | Where a rule of the harness changes its meaning, every requirement that names it as uplink goes back to `review`. |
| PROC-18 | A decision is written down with its reason and the alternatives rejected. |
| PROC-19 | Every test belongs to exactly one test level, read from the place of its file. |

**Decision**: a rule of the harness is approved when it is merged into `main`. It names no criteria
and passes no review of its own. So a rule can be the uplink of an approved requirement (PROC-13).

**Why**:

- the harness is the central safeguard: it sets how everything else is checked
- it is improved iteratively: rules arise and change from the work
- the merge is an approval with evidence: a rule reaches `main` only by ticket and pull request
  (PROC-02, PROC-03), one concern each (PROC-16); the Product Owner merges, and git keeps who and
  when

**Rejected**: to review the rules of the harness against criteria, as requirements are. The rules
that set how a review is made would then need a review made by those rules - they go round in a
circle; and a formal review of every change would slow the very means by which the process gets
better.

This overcomes the chicken-and-egg problem: the rules of review stand first, and every requirement
is reviewed by them.

## Loop for every change

1. A ticket is the start of every change (PROC-02).
2. A branch is created for the ticket.
3. The change is realized, in commits.
4. If the change touches what has review criteria of its kind (PROC-13), it needs a review. The
   author sets it to `review`. The change is then finished.
5. The reviewer checks it against the criteria of its kind. The reviewer is not the author (PROC-15).
6. The review result goes into `reviews/`.
7. A finding can be resolved on the branch.
8. A finding that cannot be resolved on the branch fails the review. It becomes one or more tickets
   of its own. Each goes to step 2 again.
9. A finding that shows a missing rule stops the review. The rule is changed in a change of its
   own. Then the review goes on, against the new rule.
10. The status of the change follows the review result: `approved` if passed, `draft` if failed.
11. A pull request takes the branch to `main` (PROC-03). The CI checks it (PROC-04).
12. The Product Owner merges. The merge is the approval.
13. The branch is deleted automatically. Its commits stay in `main`.

For code, step 3 means:

- the requirement it realizes is approved (PROC-05)
- it compiles, the linter passes, the unit tests pass
- update documentation
- the quick gate passes before a commit

The CI runs `reviewcheck` on every pull request; a review result that breaks a rule blocks the
merge (TOOL-12, see devops.md).

## Loop for every release or library

1. compile
2. run linter
3. run unit-testing in local
4. run the tests of the levels SwVal, SysIT, SysQT and Val (see table test levels)
5. report on
   - test coverage
   - structural debt
6. create the SBOM and check the licences of the components (PROC-08)
7. provide feedback
8. create a signed tag as approval (PROC-06)
9. write the data of the release or library into a separate edgedb, kept as reference
   data - the real data of treporting itself, used instead of synthetic data to show that
   it works in practice (PROC-07)
