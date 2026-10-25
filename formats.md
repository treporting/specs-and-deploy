# Formats

How requirements, architecture, risks, annotations and tests are written down, so that people
can review them in a pull request and a program can read them - the CI check and treporting
itself, in the reference edgedb.

## Requirements: YAML for the data, Markdown for the explanations

Every level is kept as data in a YAML file; text that explains - reasons, examples, the chain of
maturity, introductions - stays in Markdown beside it and refers to the YAML.

| Level | Data | Explanation | Repo |
|---|---|---|---|
| h1, h2 | `safety/hazards.yaml` (HAZ, HAZC); threats (THR, THRC), still to come | `safety/hazard-and-risk-analysis.md` | specs-and-deploy |
| n1 | `requirements/regulatory-requirements.yaml` (REG) | `requirements/regulatory-requirements.md` | specs-and-deploy |
| r1 | `requirements/user-needs.yaml` (UN) | `requirements/user-needs.md` | specs-and-deploy |
| r2 | `requirements/system-requirements.yaml` (SYS) | `requirements/system-requirements.md` | specs-and-deploy |
| r3, r4 | `architecture/system-architecture.yaml` (ARC, with the allocation) | `architecture/system-architecture.md` | specs-and-deploy |
| r5 | `requirements/software-requirements.yaml` (SR) | `requirements/software-requirements.md` | the repo of the subsystem |

A readable Markdown view of the YAML is generated on request only.

### Fields are the properties of the gel schema

The model of the prototype (cm-mono) is the reference: what is not changed on purpose is taken over
as it is; every deviation is recorded in [attributes.md](architecture/attributes.md), with its reason.

The fields of the YAML carry the names of the properties of the gel schema of the model
in storage-edgedb, so that our own data needs no configuration to be read. The YAML
holds only what the author writes:

| Kind of field | Examples | Where it comes from | In the YAML |
|---|---|---|---|
| content | `object_key`, `object_text`, `level`, `function`, `feature`, `component`, `priority`, `safety`, `status`, the uplink | the author | yes |
| context | `organization`, `project`, `baseline`, `library`, `system`, `module_id`, `revision` | the extractor: repo, file, commit, tag | no |
| derived | `istatus`, `mstatus`, `mpriority`, `msafety`, `mplatform_launch`, `mfeature`, `mfunction`, `parentOf`, `assignedTo`, `unit` | the import | never |

- **function** is the role of a user need (PM, RM, ..., PSF); a requirement below takes it from its
  parent into `mfunction` where its own is empty.
- **feature** is given on r2; r5 takes it from its parent into `mfeature`.
- **priority** is `later` where a requirement is kept for later; empty otherwise.
- **status** of every requirement: `draft`, `review`, `approved`, `obsolete`; a requirement starts
  as `draft` and is approved only after its review (PROC-13).
- **verification_method** of every requirement, after ISO/IEC/IEEE 29148: `inspection`,
  `analysis`, `demonstration`, `test`.
- **platform_launch** of every user need and system requirement: `treporting = 2_selfhosted,
  1_demo`, or narrower where a requirement holds for one of them only.
- **the links**: an object writes the ends of its own links in `target_ids` - in our files, the
  parents it is derived from (the uplink). As in the gel schema, `source_ids` and `target_ids`
  keep the original ends of the links without saying which is parent or child; the import makes
  the link `parentOf` from them, a derived value.

### Example

```yaml
# requirements/system-requirements.yaml
level: r2
requirements:
  - object_key: SYS-58
    feature: Licensing
    object_text: >-
      check, on every request to a licensed feature, the role and the rights of the user and
      the licence of the installation, and refuse the request with the reason when one of
      them is missing
    target_ids: [UN-OPS-10]
    priority: later
```

### Schemas

One JSON Schema per kind of file, in [schemas/](schemas/). They are drafts until the gel schema
stands; then they become a part of it, and a CI check in storage-edgedb compares both.

## Reviews

A review checks a requirement against the criteria of its level. Criteria and result are kept
like a test case and its result, in `reviews/`:

```
reviews/
  criteria-r2.yaml               the criteria of a level, after ISO/IEC/IEEE 29148 (RC)
  2026-10-25-sys-trace.yaml      a review: who, when, the revision of every requirement
                                 reviewed, per requirement and criterion passed or failed (RR)
```

A review counts for the revision it reviewed only (SYS-75). See the loop for every review in
[harness.md](harness.md).

## Breakdown

| Level | treporting | Example | Set by |
|---|---|---|---|
| system | treporting | - | r2 |
| subsystem | a repo, ARC-01 to ARC-06 | reporter-edgedb | r3 |
| component | a group of packages with one responsibility | `trace`, `maturity` | the field `component` of an SR |
| module | a Go package or JS module | `trace/gaps` | the code |
| unit | a source file | `trace/gaps/mark.go` | the annotation, never by hand |
| code function or method | not a function as a group of requirements | `markGaps()` | the place of the annotation |

## Ids

An id is unique within its document only: its prefix is the `module_prefix` of the document, its
number the `object_id` within it. The document is unique within its context - the project, and
for a subsystem the repo. A reference names the id; the context is known where it is written - an
annotation names software requirements of its own repo only.

| Ids | What | Level |
|---|---|---|
| `HAZ-nn`, `THR-nn` | hazards, threats | h1 |
| `HAZC-nn`, `THRC-nn` | the controls of a hazard, of a threat | h2 |
| `REG-nn` | regulatory requirements | n1 |
| `UN-<role>-nn` | user needs | r1 |
| `SYS-nn` | system requirements | r2 |
| `ARC-nn` | subsystems, interfaces | r3 |
| `ARC-Rnn`, `ARC-Dnn` | rationale and decisions of the architecture; a requirement a decision made possible names it as uplink | r3 |
| `SR-nn` | software requirements of a subsystem - unique together with their context, the repo of the subsystem (LProject) and the document | r5 |
| *open* | software architecture (r6) and software design (r7) of a subsystem; their prefixes are still to be set | r6, r7 |
| `PROC-nn` | rules of our work | - |

### Software requirements

- a software requirement states **one behaviour** of its subsystem that a test can verify - "the
  trace marks a missing link as gap"
- it names no route and no view: where a behaviour shows - a route, its contract - is set by the
  architecture, at most
- it names as a rule **one** system requirement as uplink; several only where the behaviour
  serves each of them
- every subsystem repo is an element of its own in the architecture, so that the CI can check
  the allocation (r4): ARC-06 *extractor-SYSTEM* describes the kind; extractor-git, a second
  reporter or a second viewer each get an element of their own when their repo is made

## Annotations

The same in Go and JavaScript, in the comment right above the function, method, component or
test, with no blank line between:

```go
// @implements SR-02, SR-03
func markGaps(trace *Trace) { ... }

// @verifies SR-02
func TestMarkGapsNeverCovered(t *testing.T) { ... }
```

```js
/** @implements SR-05 */
export function StatusBand(props) { ... }

// @verifies SR-05
it('shows the hidden filters', () => { ... })
```

- one annotation per line, several ids separated by commas; never over two lines
- `@implements` names software requirements of its own repo only
- `@verifies` names software requirements of its own repo in a software test, system
  requirements in a system test

## Test levels

The level of a test follows from where its file lies (see the test levels in
[harness.md](harness.md)):

| Level | Where | Repo |
|---|---|---|
| MT | `*_test.go` in the package, `*.test.js` | the subsystem |
| SwIT | `*_test.go` with the build tag `integration`, `*.integration.test.js` | the subsystem |
| SwVal | `subsystemtest/` | the subsystem |
| SysIT | `systemtest/integration/` | specs-and-deploy |
| SysQT | `systemtest/system/` | specs-and-deploy |
| Val | `systemtest/acceptance/` | specs-and-deploy |

The containers are built in `deploy/` of this repo from the parts the subsystems deliver - the
reporter and the viewer in one image. A test result is linked to its test case by package and
function name, from the JSON output of `go test` and of vitest.

## CI check

On every pull request of a code repo:

1. the YAML of the requirements matches its schema
2. every id in `@implements` and `@verifies` exists in the YAML of the **base commit** of the
   pull request - a requirement added in the same pull request does not count (PROC-05); a SYS
   named by a system test must exist in treporting on `main`
3. the number of annotations the parser found equals the count of `git grep -c`
4. the uplink of every SR: the SYS it names exists in treporting on `main` and is allocated to
   this subsystem in the architecture (r4)
5. every test file belongs to exactly one test level

In this repo: the schemas, the references both ways, and every SYS allocated.

The checking program lives in extractor-content, beside the parser of the annotations; the
workflow is a reusable workflow in `treporting/.github`.
