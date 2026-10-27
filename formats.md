# Formats

How requirements, architecture, risks, documents, reviews, annotations and tests are written down, so that people
can review them in a pull request and a program can read them - the CI check and treporting
itself, in the reference edgedb.

## Requirements: YAML for the data, Markdown for the explanations

Every level is kept as data in a YAML file; text that explains - reasons, examples, the chain of
maturity, introductions - stays in Markdown beside it and refers to the YAML.

**Table files of the levels** - where the data of each level and its explanation lie

| Level | Data | Explanation | Repo |
|---|---|---|---|
| h1, h2 | `safety/hazards.yaml` (HAZ, HAZC); threats (THR, THRC), still to come | `safety/hazard-and-risk-analysis.md` | specs-and-deploy |
| n1 | `requirements/regulatory-requirements.yaml` (REG) | `requirements/regulatory-requirements.md` | specs-and-deploy |
| r1 | `requirements/user-needs.yaml` (UN) | `requirements/user-needs.md` | specs-and-deploy |
| r2 | `requirements/system-requirements.yaml` (SYS) | `requirements/system-requirements.md` | specs-and-deploy |
| r3, r4 | `architecture/system-architecture.yaml` (ARC, with the allocation) | `architecture/system-architecture.md` | specs-and-deploy |
| tool | `requirements/tool-requirements.yaml` (TOOL) | `requirements/tool-requirements.md` | specs-and-deploy |
| r5 | `requirements/software-requirements.yaml` (SR) | `requirements/software-requirements.md` | the repo of the subsystem |

A readable Markdown view of the YAML is generated on request only.

### Fields are the properties of the gel schema

The model of the prototype is the reference: what is not changed on purpose is taken over as it
is; every deviation is recorded in [attributes.md](architecture/attributes.md), with its reason. So
a format taken over from the prototype has the prototype as its reason, and no alternative of its
own: the prototype was examined, and where it could be better, the deviation stands there.

The fields of the YAML carry the names of the properties of the gel schema of the model
in storage-edgedb, so that our own data needs no configuration to be read. The YAML
holds only what the author writes:

**Table kinds of field** - what the author writes, and what comes from elsewhere

| Kind of field | Examples | Where it comes from | In the YAML |
|---|---|---|---|
| content | `object_key`, `object_version`, `object_text`, `level`, `function`, `feature`, `component`, `priority`, `safety`, `status`, the uplink | the author | yes |
| context | `organization`, `project`, `baseline`, `library`, `system`, `module_id`, `revision` | the extractor: repo, file, commit, tag | no |
| derived | `istatus`, `mstatus`, `mpriority`, `msafety`, `mplatform_launch`, `mfeature`, `mfunction`, `parentOf`, `assignedTo`, `unit` | the import | never |

- **function** is the role of a user need (PM, RM, ..., PSF); a requirement below takes it from its
  parent into `mfunction` where its own is empty.
- **feature** is given on r2; r5 takes it from its parent into `mfeature`.
- **object_version** of every requirement, element and decision of the architecture, hazard and
  control: `x.y.z`, written in quotes - x the text and the attributes its schema names in
  `attributes_for_version_x`, y the links, z every other attribute. It is set at the approval;
  a new object starts at `'1.1.1'` (traceability.md).
- **priority** of every requirement, element of the architecture, hazard and control: a string,
  `'medium'` by default, `'later'` where it is kept for later. It stands in every object, also
  where it holds the default.
- **safety** of every requirement, element of the architecture and hazard: the safety class, a
  string, `''` by default. It stands in every object, also where it holds the default. A control
  has none: its class is always the one of its hazard (attributes.md).
- **rationale** of a requirement below the top that names no uplink - a system requirement, a
  software requirement, a tool requirement, a control of a hazard - or of an element of the
  architecture allocated no system requirement: why it stands, for instance derived from the view of
  its level (ISO29148, clause 5.2.8). The top - user needs, regulatory requirements, hazards -
  has no uplinks.
- **status** of every requirement: `draft`, `review`, `approved`, `obsolete`; a requirement starts
  as `draft` and is approved only after its review (PROC-13).
- **verification_method** of every requirement, a list - a requirement may be verified in more
  than one way, `[inspection, test]`. The values of ISO29148: `inspection`, `analysis`,
  `demonstration`, `test`; and of our own: `review` - verified by the review against the criteria
  of its level (PROC-13), for instance a document.
- **platform_launch** of every user need and system requirement: `treporting = 2_selfhosted,
  1_demo`, or narrower where a requirement holds for one of them only.
- **the links**: an object writes the parents it is derived from, its uplinks, in `source_ids` -
  its sources. `target_ids` is what an object points at: a review result points at what it
  reviewed, a verification at what it verified. As in the gel schema, `source_ids` and
  `target_ids` keep the original ends of the links; the import makes the link `parentOf` from
  them, a derived value: the sources of an object are its parents. Every uplink names the version
  of its parent, `UN-OPS-10@1.1.1` (PROC-20); so does an entry of `allocated`.
  Rejected: the uplink in `target_ids`, as it stood before - read so, parent and child turn
  round.

### Example

```yaml
# requirements/system-requirements.yaml
schema: schemas/system-requirements.schema.json
level: r2
requirements:
  - object_key: SYS-58
    object_version: '1.1.1'
    feature: Licensing
    object_text: >-
      check, on every request to a licensed feature, the role and the rights of the user and
      the licence of the installation, and refuse the request with the reason when one of
      them is missing
    source_ids: [UN-OPS-10@1.1.1]
    priority: 'later'
    safety: ''
```

### Files whose form another owns

A file whose form another owns - we only configure it, in a format another sets, such as a
workflow or the ruleset of GitHub in `.github/` - is not our data. It names no schema of ours and
carries no annotation. The requirement it serves names it.

### Schemas

**Decision**: one JSON Schema per kind of YAML file of our own data, in [schemas/](schemas/). A file
whose form another owns has none (see Files whose form another owns). Every YAML of our own data names its schema in the field `schema`, from the root of the repo, as a review result names its criteria:
`schema: schemas/tool-requirements.schema.json`. The schema holds that path as a fixed value, so a
file that names another schema fails. Every schema names the requirement it refines in its
`$comment`: `@refines ARC-22`. The schema of a kind with versions names the attributes for x in
the list `attributes_for_version_x`, always written out (traceability.md). The schemas are the
contract of ARC-22 and are reviewed with it.
They are drafts until the gel schema stands; then they become a part of it, and a CI check in
storage-edgedb compares both.

**Why**: which schema belongs to a file stands where the file is read, and a tool needs no list of
its own (TOOL-05). A schema gives the form of a requirement and realizes none, so it refines it.

**Rejected**:

- the schema in a comment of the YAML - no field, so neither a schema nor a tool can rely on it
- a table of patterns of paths in the tool - a new kind of file that no one enters stays unchecked
- `@implements` in a schema - a schema realizes no requirement (PROC-05)

### Documents

The YAML of a level names the documents beside it that explain it, each with a status of its own -
the same values as a requirement. A document is reviewed like a requirement (`review` as its
verification method), and an exported document shows its status (TOOL-02):

```yaml
schema: schemas/tool-requirements.schema.json
level: tool
documents:
- name: tool-requirements.md
  status: draft
  verification_method: [review]
requirements:
- ...
```

The documents without a YAML of requirements beside them - harness.md, devops.md, definitions.md,
standards.md, formats.md, the READMEs - are named in `documents.yaml` in the root of the repo, in
the same form, by their path.

### Standards, clauses and tables in a document

- A standard is named by its short form, as table short forms in standards.md sets it: ISO12207,
  not ISO/IEC/IEEE 12207. Only that table names it by its number and title.
- A clause names its standard: "ISO12207, clause 6.4.2", never "6.4.2" alone, also not in a
  column headed by the standard.
- Every table has a name, on a line of its own above it. A reference names the table by it - "see
  table levels", or "table levels in harness.md" from another document:

```markdown
**Table levels** - the levels of the work, and the clause of each

| Level | Here | Reference |
|---|---|---|
| r2 | system requirements (SYS) | ISO12207, clause 6.4.3 |
```

## Reviews

A review checks what is reviewed against the criteria of its kind (PROC-13). Criteria and result
are kept like a test case and its result, in `reviews/`:

```
reviews/
  criteria-r2.yaml               the criteria of a level, after ISO29148 (RC)
  20261025T140000_sys-trace.yaml a review: who, when, the revision of everything reviewed,
                                 for each and per criterion passed or failed (RR)
```

A file of results is named `<yyyymmddThhmmss>_<scope>.yaml`: the time of its `created_at`, in the
basic format of ISO8601, and what it reviewed - such as `20261026T140000_r1-pm.yaml`. Every review
has a name of its own, and the files sort by time.

A review is replaced only where it no longer holds as evidence - where a rule of its form changed;
the old file goes, and the commit message says why. Git keeps it. A review that failed is not
replaced: it stays, and the next review follows it in a file of its own. So the trail stays: the
finding, the change it caused, and the review that found it resolved.

A skeleton of a review result is made by `reviewskeleton` (see tools/README.md). It sets each
criterion and the verdict to `todo`, never to a value that could be right: the reviewer decides each.
`reviewcheck` reports every `todo`, so an unfilled skeleton cannot be merged.

A review result changes until its merge, never after - save a field its form adds that changes no
judgement, such as `schema`; the commit that adds it says so. A finding resolved on the branch is kept in a
note of its criterion: "found in this review and resolved in it: ...". The revision is then the
commit with the fix.

On the branch, the commits of a review result are its course: the proposal, the decision, a finding
resolved. What reaches `main` by the merge is the decision.

- the commit of a proposal says so - "Proposed review result of …" - and leaves the status at
  `review`; its header names it a proposal of its author
- once the reviewer has decided, a commit of its own - "Review result of … decided: passed" or
  "… decided: failed" - sets the status and names the decision in the header
- where the reviewer merges without a word, the merge is the decision; the status is then set
  before the merge

A file of results is one review, with one result for everything it reviewed:

```yaml
# reviews/20261026T140000_r1-pm.yaml
schema: schemas/review-result.schema.json
level: review
created_by: the reviewer
created_at: '2026-10-26T14:00:00+01:00'
criteria: reviews/criteria-r1.yaml
criteria_revision: 'd5922d1'
results:
- object_key: RR-01
  target_ids: [UN-PM-01]
  revision: '7215683'
  criteria:
    RC-01: passed
    RC-02: passed
    RC-03:
      result: failed
      note: '"how far" can be read as a share or as a level - say which'
  verdict: failed
```

- the verdict is `passed` only when every criterion passed (PROC-13)
- every criterion of the criteria file appears once
- a failed criterion is a block of `result` and `note`, so that the author knows what to change; a
  passed one may have a note as well, else it is just `passed`
- `revision` is the commit of what is reviewed: the review counts for it only (SYS-75)
- the reviewer is not the author of what is reviewed (PROC-15)
- a revision and a time are written in quotes: unquoted, YAML reads `7215683` as a number
- a note with `: ` in it is written in quotes: unquoted, YAML reads a new key

The schemas are `schemas/review-criteria.schema.json` and `schemas/review-result.schema.json`;
what they cannot say - every criterion once, the reviewer not the author - TOOL-03 checks. See the
loop for every change in [harness.md](harness.md).

## Verifications

**Decision**: a verification by inspection, analysis or demonstration leaves a record in
`verifications/`, named as a review result is: `<yyyymmddThhmmss>_<scope>`. The record is a YAML
with the schema `schemas/verification-result.schema.json`. Beside it stands a Markdown that
describes the verification, so that it can be repeated, and a screenshot where a log does not show
what happened.

```yaml
# verifications/20261027T100000_sr-20.yaml
schema: schemas/verification-result.schema.json
level: verification
method: demonstration
target_ids: [SR-20]
created_by: Claude Opus 5.5
confirmed_by: Mathias Fischer
created_at: '2026-10-27T10:00:00+01:00'
revision: 'a1b2c3d'
setting: 'e4f5a6b'
description: verifications/20261027T100000_sr-20.md
expected: the job schemacheck fails, and the merge is blocked
observed: |
  schemacheck  fail
  mergeStateStatus: BLOCKED
evidence:
- verifications/20261027T100000_sr-20.png
verdict: passed
```

- `created_by` ran it, `confirmed_by` confirmed it, and the two are not the same (PROC-15)
- `revision` is the commit it ran on; `setting` the commit that last changed a setting outside the
  repo it depends on, such as the ruleset of `main` in `.github/rulesets/main.json`
- `observed` holds the messages themselves, not only a link: the logs of GitHub Actions are
  deleted after 90 days
- the Markdown has the headings Requirement, Configuration, Steps, Expected, Observed and Verdict -
  the questions an auditor asks of a verification

A review and a test leave records of their own: a review in `reviews/`, a test in the output of its
run.

**Why**: an auditor asks which requirement a verification shows, what was done and expected, who
did it, when and on which revision, whether it lasts, whether it still holds, and whether it can be
repeated. A link to a pull request answers only the first.

**Rejected**:

- a link in devops.md - it answers one question of six, and the logs behind it go after 90 days
- a folder per method - the same form three times
- the records in `reviews/` - a review checks a requirement, a verification its realization

## Breakdown

**Table breakdown** - the levels of the breakdown, and what sets each

| Level | treporting | Example | Set by |
|---|---|---|---|
| system | treporting | - | r2 |
| subsystem | a repo, ARC-01 to ARC-06; or the tools in specs-and-deploy, which are not part of the product | reporter-edgedb, tools | r3 |
| component | a group of packages with one responsibility | `trace`, `maturity` | the field `component` of an SR |
| module | a Go package or JS module | `trace/gaps` | the code |
| unit | a source file | `trace/gaps/mark.go` | the annotation, never by hand |
| code function or method | not a function as a group of requirements | `markGaps()` | the place of the annotation |

## Ids

An id is unique within its document only: its prefix is the `module_prefix` of the document, its
number the `object_id` within it. The document is unique within its context - the project, and
for a subsystem the repo. A reference names the id; the context is known where it is written - an
annotation names software requirements of its own repo only.

**Table ids** - the prefixes of the ids, and their level

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
| `TOOL-nn` | requirements on the tools of our work; tools are code (devops.md) | tool |

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
- `@implements` names software requirements of its own repo only - in `tools/` as well
- `@verifies` names software requirements of its own repo in a software test, system
  requirements in a system test
- `@refines` names the requirement a schema gives the form of, in its `$comment`:
  `"$comment": "@refines ARC-22"`
- a file whose form another owns carries no annotation (see Files whose form another owns)

### Where an annotation is read

**Decision**: annotationcheck tells by the kind of a file whether to read it, from two lists in
`tools/annotation/kinds.yaml`: a **whitelist** of the kinds that are code, read for annotations,
and a **blacklist** of the kinds that are no code, never read. A file of a kind on neither list is a
finding (SR-27) until its kind is set. A YAML that names a schema is our data and is not read,
though `*.yaml` is on the whitelist. In a file that is read, an annotation stands first in a line of
comment: after spaces, a comment mark of one to three signs other than a quote, then `@implements`
or `@verifies`.

**Why**: for a risk, the question is which error goes unseen, and which does more harm.
annotationcheck is a gate: it is to stop code that names a requirement not approved (PROC-05).

**Table failures of one list** - what can go wrong with one list alone

| One list alone | What can go wrong | Seen? | Harm to the gate |
|---|---|---|---|
| blacklist - an unknown kind is code | an example in a `.txt` is read, with an id missing or not approved | yes, a finding that stops the pull request | small: the false alarm is loud, and the remedy is an entry in the list |
| blacklist - an unknown kind is code | an example in a `.txt` is read, with an id approved | no | none for the gate: no breach gets through |
| whitelist - an unknown kind is no code | a file of a new language is not read | no | the breach the gate is to stop gets through |

For the gate alone, the blacklist is the safer: its errors are loud, or harmless for the gate. The
error of the whitelist is silent, and hits the purpose of the tool. That is fail-safe: in doubt,
check, and stop once too often rather than let through once unseen. With both lists and a finding
for every unknown kind, every error is loud, both ways - the least risk, at the cost of SR-27 and an
entry for each new kind.

For counting realizations it is the other way round - the trace in the reporter, and
reporter-annotations: an example in a `.txt` with an approved id would count as a realization, a
false evidence, worse for a tool of traceability than a missing one. There the whitelist is the
safer. That is a question for those tools, not for annotationcheck.

**Rejected**:

- a whitelist alone - a file of a new language would go unchecked, unseen
- a blacklist alone - every new kind of text would be read as code, unseen
- a scanner of its own for the comments of each language - a table of languages to keep; the first
  place in a line of comment holds for every language

## Test levels

The level of a test follows from where its file lies (see table test levels in
[harness.md](harness.md)):

**Table places of the tests** - where the file of a test lies on each level

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

On every pull request of a code repo, each check a tool with a requirement of its own
(`requirements/tool-requirements.yaml`):

**Table CI checks** - each check, and the tool requirement it realizes

| Check | Tool requirement |
|---|---|
| every YAML file matches its schema | TOOL-05 |
| every id in `@implements` and `@verifies` exists, and is approved, in the YAML of the **base commit** of the pull request - a requirement added in the same pull request does not count (PROC-05); a SYS named by a system test exists in specs-and-deploy on `main` | TOOL-06 |
| the annotations counted a second way, apart from the parser - `git grep -c` | TOOL-07 |
| the uplink of every SR: the SYS it names exists in specs-and-deploy on `main` and is allocated to this subsystem in the architecture (r4) | TOOL-08 |
| an approved requirement names approved uplinks only | TOOL-09 |
| every test file belongs to exactly one test level | TOOL-10 |
| every review result keeps the rules of the section Reviews, and the tests of the tools pass | TOOL-12 |

In specs-and-deploy as well: the references both ways, and every SYS allocated.

TOOL-12 runs in `.github/workflows/tools.yml` of specs-and-deploy.

*Open*: the tools of TOOL-05 to TOOL-10 are not built yet; each waits for its approved tool
requirement. The workflow is to be a reusable workflow in `treporting/.github`.
