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
| h1, h2 | `safety/hazards.yaml` (HAZ, HAZC); threats (THR, THRC): see Open in harness.md | `safety/hazard-and-risk-analysis.md` | specs-and-deploy-product |
| n1 | `regulatory/<module>.yaml` (REG), one file per module | `regulatory/regulatory-requirements.md` | specs-and-deploy |
| r1 | `requirements/user-needs.yaml` (UN) | `requirements/user-needs.md` | specs-and-deploy-product |
| r2 | `requirements/system-requirements.yaml` (SYS) | `requirements/system-requirements.md` | specs-and-deploy-product |
| r3, r4 | `architecture/system-architecture.yaml` (ARC, with the allocation) | `architecture/system-architecture.md` | specs-and-deploy-product |
| r2 | `requirements/tool-requirements.yaml` (TOOL) | `requirements/tool-requirements.md` | specs-and-deploy |
| r5 | `requirements/software-requirements.yaml` (SR) | `requirements/software-requirements.md` | the repo of the subsystem; in tools `software-requirements.yaml`, explained in its README.md |
| r1tc, r2tc, r3tc | `testcases/<test level>.yaml` (TC), one file per test level; see Test cases | the explanation of their requirements | specs-and-deploy-product |
| r5tc, r6tc, r7tc | `testcases/<test level>.yaml` (TC), one file per test level; see Test cases | the explanation of their requirements | the repo of the subsystem, or tools |
| rules | `rules.yaml` (PROC) | `harness.md` | specs-and-deploy |

A readable Markdown view of the YAML is generated on request only - save the table of the rules
(see Rules).

### Levels

**Decision**: `levels.yaml` at the root of the repo lists every level: its key, the parts that use
it - the product, our process and tools, or both -, its name, the abbreviation of its name, what
it holds and the clause it follows; a test level also the data its tests run with; a level of our
process and tools also the level of the product it stands for, such as r1 for rules. Each repo
specs-and-deploy-* has its own `levels.yaml`; tools and the repos of the subsystems use the one of
the repo above them (#637). A tool that names a level takes the name from there.

**Why**: the levels were named only in the tables of harness.md, which no tool reads, and each
schema lists its keys apart, without a name. One file serves every tool, and the tables of
harness.md can be built from it.

**Rejected**:

- a fixed mapping in the code of a tool - a new or renamed level needs a change of code
- the names in each schema, as the title of its keys - each schema knows only its own levels

### Rules

The rules of our work are requirements under another name (traceability.md). They are kept as
data in `rules.yaml`, in the root beside harness.md, at a level of their own, `rules`, apart from
the requirements of the product. Each has the fields of a requirement; its `status` is `approved`
on `main`, and its `verification_method` is empty: a rule passes no review of its own, it is
approved by its merge (harness.md).

**Decision**: harness.md shows the rules in a table generated from `rules.yaml`, between two
marks, whenever `rules.yaml` changes. The table names the hash of the content of `rules.yaml`, in
the mark and in its title, such as `rules.yaml@1a2b3c4`. Whoever checks runs `git hash-object rules.yaml`;
the same hash says the table fits the data.

**Why**: the export to PDF can insert only Markdown, not YAML, so a table that is not in
harness.md is in no PDF of it. The hash of the content is known before the commit, so the table
and the data change in the same commit.

**Rejected**:

- the table in rules.yaml only: it would be in no PDF of harness.md
- a tool that writes the table: too much for one table
- the hash of the commit that last changed `rules.yaml`: it is known only after that commit, so a
  second commit would follow, and in between the table would not fit the data

**Decision**: harness.md is named in `documents.yaml`, not in `rules.yaml`.

**Why**: the hash covers all of `rules.yaml`. Were the status of harness.md in it, the decision of
a review of harness.md would change the hash, the table, and so harness.md again - without end.

**Rejected**: the hash over the rules alone. `git hash-object` could no longer check it, and a tool
would be needed.

### Regulatory requirements

A regulatory requirement is a clause of a standard, a law, a guidance or a licence (n1). The model
keeps it as an object in a module, and the module is where it comes from (attributes.md, section
Normative).

**Decision**: one file per module, `regulatory/<module>.yaml`. The module is a standard, or a part
of it; a law; a guidance; or the licences of the components. Its name is the short form of
standards.md, such as `regulatory/ISO26262-8.yaml`.

**Why**: the standards and laws are given from outside; the requirements are written by us. So they
lie in a directory of their own, apart from `requirements/`. The numbers of the clauses start anew
in each part of a standard: clause 11 stands in ISO26262-8 and in ISO26262-6. A file per part keeps
a clause unique.

**Rejected**:

- one file for all: the module would be the file, not the standard
- one file per standard, such as `ISO26262.yaml`: clause 11 would be ambiguous in it
- `requirements/regulatory/`: what is given and what we write would lie in one directory

**Decision**: the id is `REG_<module>_nn`, numbered from 01 in each file, such as
`REG_ISO26262-8_01`. A requirement that comes from more than one module is split, one per module.

**Why**: an id is unique within its document only (section Ids); the module makes it unique in the
project. `REG` shows the level, as `UN` does for the user needs. The module holds hyphens itself,
so the underscore separates its parts.

**Rejected**:

- `<module>_nn`, without `REG`: the id would not tell its level
- a requirement with more than one module: it would belong in more than one file

**Decision**: `object_text` begins with the clause, as its source names it, then ` - `, such as
`'clause 11 - classify a software tool …'` or `'Art. 13(1) - …'`. The field `source` goes.
`category` is `Normative`. The topics, such as tool qualification, are the sections of
regulatory-requirements.md.

**Why**: the file names the module, so of `source` only the clause is left. A clause is found in its
source by the word it carries there - clause, Art., §, Annex. In the model, `category` is the kind
of an object, as `Hazard` is for a hazard.

**Rejected**:

- a field `clause`: the gel schema has no such property (section Fields are the properties of the
  gel schema), and the clause fits in `object_text`
- the clause as a number alone, `'11: …'`: Art. 14 and Annex I of a law would both be numbers
- the topic in `category`: it would mean another thing for this kind than for every other

**Decision**: regulatory-requirements.md lies in `regulatory/`, and is named in `documents.yaml`.

**Why**: it explains every module, not one of them; `documents.yaml` holds every document (see
Documents).

**Rejected**: the section `documents` of one of the modules - which one would be chosen at will.

### Fields are the properties of the gel schema

The model of the prototype is the reference: what is not changed on purpose is taken over as it
is; every deviation is recorded in [attributes.md](https://github.com/treporting/specs-and-deploy-product/blob/main/architecture/attributes.md) of the product, with its reason. So
a format taken over from the prototype has the prototype as its reason, and no alternative of its
own: the prototype was examined, and where it could be better, the deviation stands there.

The fields of the YAML carry the names of the properties of the gel schema of the model
in the storage, so that our own data needs no configuration to be read. The YAML
holds only what the author writes:

**Table kinds of field** - what the author writes, and what comes from elsewhere

| Kind of field | Examples | Where it comes from | In the YAML |
|---|---|---|---|
| content | `object_key`, `object_version`, `object_text`, `level`, `function`, `feature`, `component`, `priority`, `safety`, `status`, the uplink | the author | yes |
| context | `organization`, `project`, `baseline`, `library`, `system`, `module_id`, `revision` | the extractor: repo, file, commit, tag | no |
| derived | `istatus`, `mstatus`, `mpriority`, `msafety`, `mplatform_launch`, `mfeature`, `mfunction`, `parentOf`, `assignedTo`, `unit` | the import | never |

- **function** is the class of a user need, the question it answers (STATUS, TRACE, ..., USE;
  user-needs.md); a requirement below takes it from its parent into `mfunction` where its own is
  empty.
- **feature** is given on r2; r5 takes it from its parent into `mfeature`.
- **object_version** of every requirement, element and decision of the architecture, hazard,
  control and document: `x.y.z`, written in quotes - x the text and the attributes its schema
  names in `attributes_for_version_x`, y the links, z every other attribute. It is set at the
  approval; a new object starts at `'1.1.1'` (traceability.md).
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
  as `draft` and is approved only after its review (PROC_13).
- **verification_method** of every requirement, a list - a requirement may be verified in more
  than one way, `[inspection, test]`. The values of ISO29148: `inspection`, `analysis`,
  `demonstration`, `test`; and of our own: `review` - verified by the review against the criteria
  of its level (PROC_13), for instance a document.
- **platform_launch** of every user need and system requirement: `treporting = 2_selfhosted,
  1_demo`, or narrower where a requirement holds for one of them only. Of every tool requirement: the
  launch of the tools it is planned for, such as `treporting = 1_tools`.
- **the links**: an object writes the parents it is derived from, its uplinks, in `source_ids` -
  its sources. `target_ids` is what an object points at: a review result points at what it
  reviewed, a verification at what it verified. As in the gel schema, `source_ids` and
  `target_ids` keep the original ends of the links; the import makes the link `parentOf` from
  them, a derived value: the sources of an object are its parents. Every uplink names the version
  of its parent, `UN_OPS_10@1.1.1` (PROC_20); so does an entry of `allocated`. Rejected: the
  uplink in `target_ids`, as it stood before - read so, parent and child turn round.
- **controls** of a control of a hazard: the system requirements that realize it, with their
  versions, `controls: [SYS_02@1.1.1]` - a down-link between two trees (traceability.md, section Links
  between trees).

### Example

```yaml
# requirements/system-requirements.yaml
schema: schemas/system-requirements.schema.json
type: requirements
level: r2
requirements:
  - object_key: SYS_58
    object_version: '1.1.1'
    feature: Licensing
    object_text: >-
      check, on every request to a licensed feature, the role and the rights of the user and
      the licence of the installation, and refuse the request with the reason when one of
      them is missing
    source_ids: [UN_OPS_10@1.1.1]
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
`$comment`: `@refines ARC_22@2.1.1`. The schema of a kind with versions names the attributes for x in
the list `attributes_for_version_x`, always written out (traceability.md). The schemas are the
contract of ARC_22 and are reviewed with it.
They are drafts until the gel schema stands; then they become a part of it, and a CI check in
the storage compares both.

**Why**: which schema belongs to a file stands where the file is read, and a tool needs no list of
its own (TOOL_05). A schema gives the form of a requirement and realizes none, so it refines it.

**Rejected**:

- the schema in a comment of the YAML - no field, so neither a schema nor a tool can rely on it
- a table of patterns of paths in the tool - a new kind of file that no one enters stays unchecked
- `@implements` in a schema - a schema realizes no requirement (PROC_05)

### The head of a file

**Decision**: the head of every YAML of our own data has three fields:

- `schema` - its form (see Schemas)
- `type` - what it holds, as table types sets it
- `level` - the level it belongs to. A review result names the level of its criteria, a
  verification the level of its targets.

A Markdown file has no head. Its entry in documents.yaml names its type and level (see Documents).

**Why**: a tool for requirements management knows what each of its modules holds. Our files take
its place, so each says what it holds itself. The level of a review result is the level it was
reviewed against.

**Rejected**:

- telling what a file holds by where it lies, or by the attributes its schema has - the file does
  not say it, a tool has to infer it
- `level` alone - the requirements of r2 and their criteria both have `level: r2`
- `level: review` and `level: verification` - they name the kind of the record, not its level

**Table types** - what a file holds, and its level

| type | Holds | Level | Files |
|---|---|---|---|
| requirements | requirements | r1, r2, r5 | `requirements/*.yaml`; in tools `software-requirements.yaml` |
| testcases | test cases | r1tc to r7tc | `testcases/*.yaml` |
| architecture | the system architecture | r3 | `architecture/system-architecture.yaml` |
| normative | regulatory requirements | n1 | `regulatory/<module>.yaml` |
| hazards | hazards and their controls | h1 | `safety/hazards.yaml` |
| threats | threats and their controls: see Open in harness.md | h1 | - |
| rules | the rules of our work | rules | `rules.yaml` |
| criteria | review criteria | the level they are for, or document | `reviews/criteria-*.yaml` |
| review | review results | the level of their criteria | `reviews/<time>_<scope>.yaml` |
| verification | verification results | the level of their targets | `verifications/<time>_<scope>.yaml` |
| documents | every Markdown file, with its type, level and version | - | `documents.yaml` |
| kinds | the kinds of files, for annotationcheck | - | `kinds.yaml`; in tools `annotation/kinds.yaml` |
| levels | the levels of the work, each with its name | - | `levels.yaml` |
| explanation | the explanation of a level, in Markdown | its level | see table files of the levels |
| document | a document without a level, in Markdown | - | such as formats.md |

### Documents

**Decision**: `documents.yaml` in the root of the repo holds every Markdown file, by its path from
the root. The description of a verification is not in it: it is a part of its record (see
Verifications). Each entry names:

- `type` and `level`, as table types sets them - `explanation` with the level it explains, or
  `document` without a level
- a version and a status of its own - the same values as a requirement

A YAML of a level names the Markdown files that explain it in its field `documents`, by their path,
without a version.

A document is reviewed like a requirement (`review` as its verification method), except a document
a tool writes (see below). An exported document shows its version, its revision and its status
(TOOL_02). Its `object_hash` is the content approved: `git hash-object` of its file, 7 characters,
set at the approval with the version, `''` before the first. It is the one attribute for x of a document (traceability.md, section The version
of a document):

```yaml
# documents.yaml
schema: schemas/documents.schema.json
type: documents
documents:
- name: requirements/tool-requirements.md
  type: explanation
  level: r2
  object_version: '2.1.1'
  object_hash: 'a1b2c3d'
  status: approved
  verification_method: [review]
- name: formats.md
  type: document
  object_version: '8.1.1'
  ...

# requirements/tool-requirements.yaml
schema: schemas/tool-requirements.schema.json
type: requirements
level: r2
documents: [requirements/tool-requirements.md]
requirements:
- ...
```

**Why**: one place holds the version of every document. A tool finds it there, whatever the
document explains.

**Decision**: a document a tool writes is a record: the tool writes it complete, without a review
(PROC_13). Its verification method is `generated`. The tool sets its entry in `documents.yaml` on
every change of its content: x of its version raised by one - `1.1.1` the first time -, its
`object_hash`, and the status `completed`, as a decided review result has (section Reviews). Today
these are `sbom.md` of a subsystem (SR_40) and its test reports (section Test report).

**Decision**: a generated document is written from original data. The original data is kept only
on demand, not with every document.

**Table original data** - of each generated document

| Document | Original data |
|---|---|
| `sbom.md` | the output of `go list` for each go.mod, each `package-lock.json`, and the list of the Alpine packages of each image |
| test report | the `junit.xml` of each run of the tests, of `go test` and of Vitest |

**Why**:

- it is a report of what a tool found, not text a person wrote. The tests of the tool and the
  checks that read it - such as the check of the licences (TOOL_22) - check what it holds; a
  review could only compare it with the same output once more
- everyone can make it alone, without a reviewer
- it keeps the one form of a document: a version, a hash, a status and a change history

**Rejected**:

- a review of every generated document - it adds nothing to what the tool and its checks show,
  and it needs a second person for every baseline
- no entry in `documents.yaml`, with the commit as its version - a second form of a version, beside
  the `x.y.z` of every other document (traceability.md, section The version of a document)

**Rejected**:

- the version of an explanation in the YAML of its level, as before - versions of documents in
  eight places; harness.md and regulatory-requirements.md have no such YAML of their own
- a file `levels.yaml` with the files and the explanations of each level - nothing fails without
  it, and it would hold a second time what the heads of the files and documents.yaml say

### Open points

A document lists what is open (PROC_11) in one section `## Open`, right before `## Change history`.
An open point can also come from a review. Each point is rated by the launches that need its answer,
in the form of `platform_launch`. Where a point names objects, these are the launches of those
objects. A document with no open point has the section all the same, with `<empty>`.

```markdown
## Open

- `treporting = 2_tools` - **Maintenance**: not set how the state of specs-and-deploy of an older
  release is made and kept. A fix for an older release needs it, from its tag, not `main`.
```

### Change history

A document has a section `## Change history`, its last, after `## Open`. It logs every approved version
since its inception (REG_ISO29119-3_01), one row each:

| Version | Date | Change | Reason | By |
|---|---|---|---|---|

- Change: the sections or table revised
- Reason: the ticket
- By: the author and the reviewer, with their roles

The row is written with the change itself, so that its review covers it (reviews/README.md, section
The change history of a document).
A document a tool writes names the tool with its version in By, and no reviewer.

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

A review checks what is reviewed against the criteria of its kind (PROC_13). Criteria and result
are kept like a test case and its result: the criteria in `reviews/`, the results apart from them in
`reviews/results/`, so that the few criteria are not lost among the many results:

```
reviews/
  criteria-r2.yaml                 the criteria of a level, after ISO29148 (RC)
  results/
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

A skeleton of a review result is made by `reviewskeleton` (see README.md of tools). It sets each
criterion and the verdict to `todo`, never to a value that could be right: the reviewer decides each.
`reviewcheck` reports every `todo`, so an unfilled skeleton cannot be merged.

A review result changes until its merge, never after - save a field its form adds that changes no
judgement, such as `schema`; the commit that adds it says so. A finding resolved on the branch is kept in a
note of its criterion: "found in this review and resolved in it: ...". The revision is then the
commit with the fix.

**Decision**: a review result is a record. It carries a version and a status of its own, apart
from its verdict:

- `object_version`: `1.1.1`. A completed result does not change, so it keeps it.
- `status`: `open` while the result is made - the skeleton and the proposal; `completed` once the
  reviewer has decided. A completed result may have the verdict `failed`.
- `created_at`, the date recorded: when the skeleton was made. `completed_at`, the date of record:
  when the reviewer decided. It is set with `completed`.

**Why**: ISO15288 asks that each information item is identified - its version - and confirmed
valid - its status (REG_ISO15288_02). ISO15289 asks a record for the date of record, the date
recorded and its status. In practice a review is open while it runs, and closed when the reviewer
has decided; approved and rejected belong to what is reviewed, not to the review.

**Rejected**:

- the statuses of a document, `draft` to `obsolete`: they say whether a specification may be used.
  A record is not approved by a review of its own; it is completed.
- `closed` in place of `completed`: it presumes that something was open. A record a tool writes
  complete, such as sbom.md, never was.
- one date for both: the decision often comes hours after the proposal.
- the date of the decision from git only: the commit of the decision is not in the record.

On the branch, the commits of a review result are its course: the proposal, the decision, a finding
resolved. What reaches `main` by the merge is the decision.

- the commit of a proposal says so - "Proposed review result of …" - leaves the result `open`, and
  sets the status of what is reviewed to `review`; its header names it a proposal of its author
- once the reviewer has decided, a commit of its own - "Review result of … decided: passed" or
  "… decided: failed" - sets the result `completed` with `completed_at`, sets the status of what is
  reviewed, and names the decision in the header
- where the reviewer merges without a word, the merge is the decision; the status is then set
  before the merge

A file of results is one review, with one result for everything it reviewed:

```yaml
# reviews/results/20261026T140000_r1-pm.yaml
schema: schemas/review-result.schema.json
type: review
level: r1
object_version: '1.1.1'
status: completed
created_by: the reviewer
created_at: '2026-10-26T14:00:00+01:00'
completed_at: '2026-10-26T15:30:00+01:00'
criteria: reviews/criteria-r1.yaml
criteria_revision: 'd5922d1'
results:
- object_key: RR_01
  target_ids: [UN_PM_01]
  revision: '7215683'
  criteria:
    RC_01: passed
    RC_02: passed
    RC_03:
      result: failed
      note: '"how far" can be read as a share or as a level - say which'
  verdict: failed
```

- the verdict is `passed` only when every criterion passed (PROC_13)
- every criterion of the criteria file appears once
- a failed criterion is a block of `result` and `note`, so that the author knows what to change; a
  passed one may have a note as well, else it is just `passed`
- `revision` is the commit of what is reviewed: the review counts for it only (SYS_75)
- the reviewer is not the author of what is reviewed (PROC_15)
- a revision and a time are written in quotes: unquoted, YAML reads `7215683` as a number
- a note with `: ` in it is written in quotes: unquoted, YAML reads a new key

The schemas are `schemas/review-criteria.schema.json` and `schemas/review-result.schema.json`;
what they cannot say - every criterion once, the reviewer not the author - TOOL_03 checks. See the
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
type: verification
level: r5
method: demonstration
target_ids: [SR_20]
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

- `created_by` ran it, `confirmed_by` confirmed it, and the two are not the same (PROC_15)
- `revision` is the commit it ran on; `setting` the commit that last changed a setting outside the
  repo it depends on, such as the ruleset of `main` in `.github/rulesets/main.json`
- `observed` holds the messages themselves, not only a link: the logs of GitHub Actions are
  deleted after 90 days
- the Markdown has the headings Requirement, Configuration, Steps, Expected, Observed and Verdict -
  the questions an auditor asks of a verification
- a screenshot is committed with its longer side at most 1792 pixels, half of a Retina screen: it
  stays readable, and a commit is kept for good. On a Mac, before the commit:
  `sips -Z 1792 verifications/<name>.jpg`

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
| subsystem | a repo, ARC_01 to ARC_06; or the tools in specs-and-deploy, which are not part of the product | reporter-trace, tools | r3 |
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
| `HAZ_nn`, `THR_nn` | hazards, threats | h1 |
| `HAZC_nn`, `THRC_nn` | the controls of a hazard, of a threat | h2 |
| `REG_<module>_nn` | regulatory requirements; `<module>` is the short form of standards.md | n1 |
| `UN_<class>_nn` | user needs | r1 |
| `SYS_nn` | system requirements | r2 |
| `ARC_nn` | subsystems, interfaces | r3 |
| `ARCR_nn`, `ARCD_nn` | rationale and decisions of the architecture; a requirement a decision made possible names it as uplink | r3 |
| `SR_nn` | software requirements of a subsystem - unique together with their context, the repo of the subsystem (LProject) and the document | r5 |
| `TC_nn` | test cases of a test level - unique together with their context, the repo and the file | r1tc to r7tc |
| `SA_nn` | software architecture of a subsystem | r6 |
| `SAR_nn`, `SAD_nn` | rationale and decisions of the software architecture of a subsystem | r6 |
| `SD_nn` | software design of a subsystem | r7 |
| `PROC_nn` | rules of our work | rules |
| `TOOL_nn` | requirements on the tools of our work; tools are code (devops.md) | tool |

**Decision**: a key is any number of parts joined by underscores, the last part the number:
`SYS_09`, `UN_PSF_01`, `REG_ISO26262-8_01`, `x_y_z_00`. The keys of criteria and review results take
this form too: `RC_nn`, `RR_nn`.

**Why**: a part may hold a hyphen itself, as the module of a regulatory requirement does,
`REG_ISO26262-8_01`. The underscore separates the parts without doubt, and every key has one form.

**Rejected**:

- the hyphen for every key but those of the regulatory requirements - two forms of keys would
  stand side by side
- `RC-nn` and `RR-nn` kept, as no keys of requirements - two forms of keys would stand side by side

**Decision**: the number `nn` of a key has two or more digits, from 01 on: `SR_09`, `SR_99`,
`SR_100`. Every schema takes this form.

**Why**: a level can hold more than 99 objects, as the test cases of the tools do. Two digits at
least keep every key of today.

**Rejected**:

- three digits for every key, such as `SR_001` - every key of today would change

**Decision**: the review results and the verifications name the keys as they are now. A review
result whose target got a new key names as its revision the commit that changed the key; one whose
criteria got new keys names as its criteria revision the commit that changed them. What a record
quotes from a log stays as the log printed it.

**Why**: every key in the repo changes, the records too. A tool finds the target of a review result
and its criteria in the revisions it names; a revision before the change does not hold the new key. A quote of a log
that changed would no longer be what the tool printed.

**Rejected**: the records as they were written, with the old keys - two forms of keys would stand
side by side.

### Software requirements

- a software requirement states **one behaviour** of its subsystem that a test can verify - "the
  trace marks a missing link as gap"
- it names no route and no view: where a behaviour shows - a route, its contract - is set by the
  architecture, at most
- it names as a rule **one** system requirement as uplink; several only where the behaviour
  serves each of them
- every subsystem repo is an element of its own in the architecture, so that the CI can check
  the allocation (r4): ARC_06 *extractor-SYSTEM* describes the kind; extractor-git, a second
  reporter or a second viewer each get an element of their own when their repo is made

## Annotations

The same in Go and JavaScript, in the comment right above the function, method, component or
test, with no blank line between:

```go
// @implements SR_02@1.1.1, SR_03@2.1.1
func markGaps(trace *Trace) { ... }

// @verifies SR_02@1.1.1
func TestMarkGapsNeverCovered(t *testing.T) { ... }
```

```js
/** @implements SR_05@1.2.1 */
export function StatusBand(props) { ... }

// @verifies SR_05@1.2.1
it('shows the hidden filters', () => { ... })
```

- one annotation per line, several ids separated by commas; never over two lines
- each id names the version of its requirement, as `ID@x.y.z`, as a link does (TOOL_20;
  traceability.md, section The right side of the V)
- `@implements` names software requirements of its own repo only - in tools as well
- `@verifies` names software requirements of its own repo in a software test, system
  requirements in a system test
- `@refines` names the requirement a schema gives the form of, in its `$comment`, where the
  requirement lies in the repo of the schema or above it:
  `"$comment": "@refines ARC_22@2.1.1"`
- a file whose form another owns carries no annotation (see Files whose form another owns)

### Where an annotation is read

**Decision**: annotationcheck tells by the kind of a file whether to read it, from two lists in
the file of kinds of the repo, which annotationcheck is given with `-kinds`: a **whitelist** of the kinds that are code, read for annotations,
and a **blacklist** of the kinds that are no code, never read. A file of a kind on neither list is a
finding (SR_27) until its kind is set. A YAML that names a schema is our data and is not read,
though `*.yaml` is on the whitelist. In a file that is read, an annotation stands first in a line of
comment: after spaces, a comment mark of one to three signs other than a quote, then `@implements`
or `@verifies`. `@refines` is read apart from the lists: in the `$comment` of each JSON in
schemas/, though `*.json` is on the blacklist (SR_35).

**Why**: for a risk, the question is which error goes unseen, and which does more harm.
annotationcheck is a gate: it is to stop code that names a requirement not approved (PROC_05).

**Table failures of one list** - what can go wrong with one list alone

| One list alone | What can go wrong | Seen? | Harm to the gate |
|---|---|---|---|
| blacklist - an unknown kind is code | an example in a `.txt` is read, with an id missing or not approved | yes, a finding that stops the pull request | small: the false alarm is loud, and the remedy is an entry in the list |
| blacklist - an unknown kind is code | an example in a `.txt` is read, with an id approved | no | none for the gate: no breach gets through |
| whitelist - an unknown kind is no code | a file of a new language is not read | no | the breach the gate is to stop gets through |

For the gate alone, the blacklist is the safer: its errors are loud, or harmless for the gate. The
error of the whitelist is silent, and hits the purpose of the tool. That is fail-safe: in doubt,
check, and stop once too often rather than let through once unseen. With both lists and a finding
for every unknown kind, every error is loud, both ways - the least risk, at the cost of SR_27 and an
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

The level of a test follows from where its file lies (see table product in
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
reporter and the viewer in one image. A test result is linked to its test function by package and
function name, from the JSON output of `go test` and of vitest.

### Test cases

A test case says what is verified and how (traceability.md, section Tests). Every level of
requirements has a test level that verifies it (harness.md, table maturity), and beside it three
levels of its own: the test cases `r<n>tc`, the test procedures `r<n>tp` and the test results
`r<n>tr`. The test cases of a test level lie in one file, its test specification, in the directory
`testcases/` of the repo that holds the requirements:

**Table test cases** - where the test cases of each level lie

| Requirements | Test level | Test cases | File | Repo |
|---|---|---|---|---|
| r1 | Val | r1tc | `testcases/acceptance.yaml` | specs-and-deploy |
| r2 | SysQT | r2tc | `testcases/system.yaml` | specs-and-deploy |
| r3 | SysIT | r3tc | `testcases/system-integration.yaml` | specs-and-deploy |
| r5 | SwVal | r5tc | `testcases/software.yaml` | the subsystem |
| r6 | SwIT | r6tc | `testcases/software-integration.yaml` | the subsystem |
| r7 | MT | r7tc | `testcases/unit.yaml` | the subsystem |

The tools keep theirs in `testcases/` of the repo tools, beside `software-requirements.yaml`.
Today only the tools have test cases, at r5tc.

```yaml
- object_key: TC_01
  object_version: '1.1.1'
  object_text: 'a completed document whose object_hash is not the hash of its file is reported'
  design_technique: equivalence partitioning
  target_ids: [SR_79@2.1.2]
  status: draft
  verification_method: [review]
```

- **object_text** states one case: what is given and what is expected. A positive and a negative
  case are two test cases.
- **design_technique** names how the case was derived, one of the techniques of
  REG_ISO29119-4_01 to _03, such as `equivalence partitioning` or `boundary value analysis`.
- **target_ids** names the requirements it verifies, with their version. A test case points to
  its requirement; the requirement does not name its test cases. Only a requirement whose
  verification method holds `test` has test cases.
- **status** and **verification_method** as of a requirement: a test case is approved after its
  review against the criteria of its level, such as `reviews/criteria-r5tc.yaml`.

**Decision**: the test cases lie in `testcases/`, one file per test level, apart from the
requirements and from the test procedures.

**Why**: the right side of the V has a place of its own, as `requirements/` is the place of the
left. One file per test level is the test specification of that level in the sense of
ISO/IEC/IEEE 29119. Where testing is let as a work package, the customer writes the test cases and
the contractor the procedures; apart, each can be delivered alone.

**Rejected**:

- beside the requirements of their level, such as `requirements/system-test-cases.yaml` - the left
  and the right side of the V would mix, and the architecture and the design would hold test
  cases too
- beside the test procedures, such as `subsystemtest/test-cases.yaml` - the unit and the
  integration tests have no one place, their procedures lie in the packages; and case and
  procedure could not be delivered apart

**Decision**: `design_technique`, and not `technique` or `method`.

**Why**: it is the term of ISO/IEC/IEEE 29119-4, clause 5, without the word test, which the test
case says itself. It keeps apart how a test case was derived from how a requirement is verified,
the verification method.

**Rejected**:

- `method` - it would be read as the verification method
- `technique`, `test_technique` - they do not say that it is about the design of the case

### Test report

**Decision**: the result of a run of the tests of a test level is its test report,
`testreports/<test level>.md` beside `testcases/`, such as `testreports/software.md` for r5tc. A
tool writes it from the `junit.xml` of the run, so it is a record (section Documents).
It is made at every baseline, and may be made at any time between.

**Table sections of a test report** - in their order

| Section | What it holds |
|---|---|
| (head) | the title and what the document is, as in every document |
| Test environment | the baseline (its label), the commit, the date, the operating system, the runner, and each runtime and test runner the tests ran with, with its version |
| Original data | every file of original data the report is written from: the command that made it, and where it is stored, as `./<folder>/<file>`, or `-` |
| Test specification | every test case: its requirement, its key, its text and its status |
| Test implementation | every test procedure: its requirement, its test case, its name and its file |
| Test execution | every test procedure that ran: its requirement, its test case, its name and its verdict |
| Deviations | every test result that has a verdict, and the verdict is not passed |
| Open | every test case and every test procedure without a test result (section Open points) |
| Change history | as in every document a tool writes (section Change history) |

The three tables follow the links of a test (traceability.md, section Tests). The requirement is
the first column, where there is one: of a test case its `target_ids`, of a test procedure without
a test case its `@verifies`. The column of the test case is empty for a test procedure that
implements none.

```markdown
## Original data

| File | Made by | Stored At |
|---|---|---|
| junit-go.xml | `gotestsum --junitfile junit-go.xml -- -tags integration ./...` | - |

## Test specification

| Requirement | Test case | Text | Status |
|---|---|---|---|
| SR_03@1.1.3 | TC_08@1.1.1 | a failed criterion without a note: reviewcheck reports the criterion | approved |

## Test implementation

| Requirement | Test case | Test procedure | File |
|---|---|---|---|
| SR_03@1.1.3 | TC_08@1.1.1 | TestAFailedCriterionHasANote | review/check/check_test.go |
| SR_24@3.1.1 | | TestOnlyFirstInALineOfComment | annotation/check/check_test.go |

## Test execution

| Requirement | Test case | Test procedure | Verdict |
|---|---|---|---|
| SR_03@1.1.3 | TC_08@1.1.1 | TestAFailedCriterionHasANote | passed |
```

A row of the change history names the label of the baseline in Change and its ticket in Reason. A
report made between two baselines leaves both empty.

**Why**:

- it keeps the result as long as the baseline; the log of the CI does not (traceability.md,
  section Tests, E1)
- one table per part of a test shows each part, also where another is missing: a test case
  without a procedure, a procedure without a test case

**Rejected**: one table of the test cases with their verdicts - a test procedure without a test
case, and a test case without a result, would have no row.

## Open

- `treporting = 1_tools` - **The test level of the tools**: the tests of the tools lie as
  `*_test.go` in their packages, the places of MT and SwIT, but they verify software requirements,
  which SwVal verifies. Not set where they belong.

## Change history

| Version | Date | Change | Reason | By |
|---|---|---|---|---|
| 1.1.1 | 2026-10-28 | - | #272 | Claude Opus 5.5 and Mathias Fischer, authors; Mathias Fischer, reviewer |
| 2.1.1 | 2026-10-28 | Fields are the properties of the gel schema; Documents | #272 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 3.1.1 | 2026-10-28 | Annotations | #279 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 4.1.1 | 2026-10-28 | Requirements: YAML for the data, Markdown for the explanations; Regulatory requirements; Ids | #40 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 5.1.1 | 2026-10-28 | Regulatory requirements; Fields are the properties of the gel schema; Example; Schemas; Documents; Reviews; Verifications; Breakdown; Ids; Software requirements; Annotations; Where an annotation is read; CI check | #291 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 6.1.1 | 2026-10-29 | Schemas; Annotations; Where an annotation is read | #283 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 7.1.1 | 2026-10-29 | Example; The head of a file; Documents; Reviews; Verifications | #304 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 8.1.1 | 2026-10-29 | Reviews | #307 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 9.1.1 | 2026-10-29 | Regulatory requirements; The head of a file; Documents | #303 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 10.1.1 | 2026-10-29 | Reviews; Ids | #302 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 11.1.1 | 2026-10-29 | Fields are the properties of the gel schema | #277 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 12.1.1 | 2026-10-29 | Requirements: YAML for the data, Markdown for the explanations; The head of a file; Ids; Open | #322 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 13.1.1 | 2026-10-30 | Requirements: YAML for the data, Markdown for the explanations; The head of a file; Open points; Ids; Open | #322 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 14.1.1 | 2026-10-30 | Open points; Change history | #323 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 15.1.1 | 2026-10-30 | Requirements: YAML for the data, Markdown for the explanations | #338 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 16.1.1 | 2026-10-30 | Fields are the properties of the gel schema; Schemas; Breakdown | #334 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 17.1.1 | 2026-10-30 | CI check | #320 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 18.1.1 | 2026-11-01 | Documents; Change history | #343 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 19.1.1 | 2026-11-01 | Open | #373 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 20.1.1 | 2026-11-01 | Verifications | #382 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 21.1.1 | 2026-11-02 | Reviews: version, status and date of record of a review result | #153 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 22.1.1 | 2026-11-03 | Documents: a generated document is a record with the status completed | #423 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 23.1.1 | 2026-11-03 | Test cases: their levels, files and fields; Ids: TC_nn | #437 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 24.1.1 | 2026-11-03 | Table types: the type testcases; Open | #437 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 25.1.1 | 2026-11-03 | Ids: the number of a key has two or more digits | #455 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 26.1.1 | 2026-11-04 | Test report; Documents: the test report is a record | #438 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 27.1.1 | 2026-11-04 | Documents: the original data of a generated document; Test report: written from junit.xml | #438 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 28.1.1 | 2026-11-04 | Test report: the section Original data | #438 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 29.1.1 | 2026-11-04 | Test report: the runtimes and test runners in Test environment | #438 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 30.1.1 | 2026-11-04 | Test report: one per test level | #438 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 31.1.1 | 2026-11-05 | Levels: levels.yaml; Table types: the type levels | #535 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 32.1.1 | 2026-11-05 | Test levels: the table levels of harness.md | #544 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 33.1.1 | 2026-11-05 | Levels: the parts that use a level; Test levels: the table product of harness.md | #551 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 34.1.1 | 2026-11-11 | Table kinds of field: function is the class of a user need; Table ids: UN_<class>_nn | #590 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 35.1.1 | 2026-11-12 | Table files of the levels and table types: the three repos; the tool requirements on r2 | #641, #628 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
