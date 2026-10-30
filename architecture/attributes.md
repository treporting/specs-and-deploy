# Attributes of the model

This is the contract of the information model, ARC_21: the common language of converter, storage
and reporter over the interfaces ARC_12, ARC_13 and ARC_18.

Every attribute of the model, type by type: what it means, where it comes from, whether it has an
intermediate (`I`) or a mapped (`M`) value, how it is carried down the trace, and which system
requirements use it. The names are those of the gel schema of the model and of the YAML
([formats.md](../formats.md)). The terms are those of
[definitions.md](../definitions.md).

The prototype is the reference: what is not changed on purpose is taken over as it is. Every
entry says how it stands to the prototype - `same`, `renamed`, `changed` or `new` - and why where
it differs; the deviations are listed at the end.

Where it comes from:

- **author** - written in the source, in our YAML by the author
- **context** - where the object lies, given by the extractor
- **derived** - made by the import, never written in the source

The prefixes:

- no prefix - the **original**, as the source writes it; never changed
- `M` - the **mapped** value: unified, translated or adapted to the model by the configuration of
  the customer. Where the original is empty, it is the mapped value of the parent - priority,
  safety, platform and launch are often set on a parent only
- `I` - an **intermediate** form, only where the original cannot be processed as it is and has to
  be transformed first. Only the status has one
- `C` - a **computed** value, made by the reporter from other objects, where one is needed: the
  compliance state of a clause, `cstatus`

## All types

Every type has the four attributes of creation and modification. The objects in a document -
Requirement, Testcase, Testresult, Hazard, Normative - have `module_prefix` as well.

### module_prefix

| | |
|---|---|
| Meaning | an abbreviation for the path and name of the document an object stands in, unique within its RProject or LProject; together with `object_id` it forms the id, and links use it |
| Where | on the objects in a document; on the document itself (Artifact) as well |
| Comes from | context (the module) |
| See | object_id and module_prefix under Requirement |
| Prototype | changed - on every object in a document; in the prototype not on every one |

### category

| | |
|---|---|
| Meaning | the kind of an object or a document; which values it takes depends on the type |
| Comes from | author or context, as the prototype sets it |
| Prototype | same - though it looks inconsistent, it is kept as it is |

Overview per type, as the prototype sets it; the terms are those of definitions.md - Document, Tests, Hazard and measure, Normative, Project, Baseline:

| Type | category |
|---|---|
| Requirement, Testcase, Testresult | empty on the object - the document it stands in has the category |
| Artifact - a document | `Requirement`, `Testspecification`, `Testprocedure`, `Testreport` - from the level of its objects |
| Hazard | `Hazard` |
| Normative | `Normative` |
| Component | `rproject`, `lproject`, `component`, `unit` |
| Baseline | `release`, `library` |

### created_at, created_by, modified_at, modified_by

| | |
|---|---|
| Meaning | when and by whom an object was created in its source, and when and by whom it was last modified there |
| Value | a time with its zone; the name of a person as the source gives it |
| Comes from | the source, as it records it - DOORS, for instance, keeps *Created On*, *Created By*, *Last Modified On*, *Last Modified By*; for a file in GIT the commit gives them |
| Note | not the time of the import: that is in the import record. The names are personal data, kept only for the purposes set (UN_ALL_04, SYS_44) |
| Used by | REG_FDADI_01, data attributable to whom it comes from; SYS_16, what changed between two baselines |
| Prototype | new - REG_FDADI_01: data attributable to whom it comes from |

## Requirement

### platform_launch

| | |
|---|---|
| Meaning | the platforms a requirement holds for, and per platform the launches it is planned for |
| Value | one line per platform, its launches after `=` separated by commas, each launch its number and its type: `platform-1 = 1_type1, 2_type1, 2_type2`, next line `platform-2 = 1_type1, 2_type1, 2_type2`; spaces around `=` and after a comma are allowed |
| Comes from | author |
| Variants | `mplatform_launch` - mapped, or where the original is empty the union of the `mplatform_launch` of the parents |
| Carried down | yes. A requirement with several parents carries the **union** of their platforms and launches - it holds wherever any of them holds (unlike the safety class, where the highest is carried) |
| Used by | SYS_09 filter over the whole trace; SYS_11 carrying down. The selection of a release offers the launches that release has |
| Version | a new version of an object where its launches change, not with every release |
| Prototype | same - the prototype writes the value taken from the parents into `mplatform_launch` |

### status

| | |
|---|---|
| Meaning | the state of an object in its source - in review, released, open |
| Example | `status = "status_platform-1 = in Review_Vendor` / `status_platform-2 = open"` |
| Variants | `istatus` - the intermediate form: `platform-1 = in Review` / `platform-2 = open`; `mstatus` - mapped: `review` when the filter selects platform-1. Where `mstatus` is not stored, it is made for the filter set |
| Note | the only attribute with an intermediate form |
| Prototype | same |

### object_id and module_prefix

| | |
|---|---|
| Meaning | `object_id` is the number of an object within its module - a document, kept as folder and name; it is unique only within that module. `module_prefix` is an abbreviation that makes the module unique within its RProject or LProject |
| Example | the LProjects converter-edgedb, reporter-edgedb and viewer-edgedb each hold a document *SoftwareRequirements* with the prefix `SR`; its requirements are numbered within it. A link to one of them is written with the prefix: `target_ids = SR_12` |
| Comes from | context (the module), the author (the number) |
| Prototype | same |

### changes

| | |
|---|---|
| Meaning | the names of the attributes that changed against the previous version of the object |
| Use | to show the difference between two versions efficiently - SYS_16 |
| Comes from | derived, by comparing with the previous version |
| Prototype | same |

### verification_method

| | |
|---|---|
| Meaning | how a requirement is verified, after ISO/IEC/IEEE 29148 |
| Value | a list of one or more of `inspection`, `analysis`, `demonstration`, `test` - after ISO/IEC/IEEE 29148 - and `review`, our own: verified by the review against the criteria of its level |
| Where | on every requirement, on every level |
| Used by | SYS_73, SYS_76 |
| Comes from | author |
| Variants | `mverification_method` mapped - the converter maps the values of a customer to those of 29148: *review* to `inspection`, *design* to `analysis`. A customer may keep it in an attribute called coverage, which is a different thing here: the converter has to take care |
| Prototype | renamed - from `coverage`, which had two meanings; values after ISO/IEC/IEEE 29148 |

### coverage

| | |
|---|---|
| Meaning | how far the code is covered by tests: the code coverage level, such as `function` or `statement` coverage |
| Comes from | derived, from the coverage reports of the test runs - never written by an author |
| Used by | the chain of maturity: one step per code coverage level of source and object code |
| Prototype | changed - the code coverage level now; the prototype kept the verification method in it |

### signal_ids, parameter_ids, issue_ids

| | |
|---|---|
| Meaning | the original ends of links to signals, parameters and issues, as the source writes them |
| Note | for treporting itself, `issue_ids` would hold the link to the tickets - where it cannot be derived the other way round |
| Prototype | same |

## Testcase

A test case and a test procedure are objects of the same type, `Testcase`, told apart by their
level. A test case is the unit a verdict belongs to - the results hang on it; a test procedure is
how tests are carried out, a script being an automated one.

### level

| | |
|---|---|
| Meaning | the key of the level, of two parts: the requirement level it verifies and the kind - `tc` a test case, `tp` a test procedure: `r5tc`, `r5tp` |
| Our test levels | MT `r7tc`, SwIT `r6tc`, SwVal `r5tc` and `r5tp`, SysIT `r3tc`, SysQT `r2tc`, Val `r1tc` |
| Comes from | context - for a test in code, the place of its file |
| Prototype | same |

### object_key

| | |
|---|---|
| Meaning | the id of the test case in its source |
| For a test in code | the package and the name of the test function - a test file is the test procedure, a test function with `@verifies` the test case |
| Prototype | same |

### stream

| | |
|---|---|
| Meaning | the variant the test is run in - a hardware, a mode of operation (see definitions.md) |
| Comes from | author |
| Used by | SYS_09, filter |
| Prototype | same |

### mfeature, mfunction, msafety

| | |
|---|---|
| Meaning | the feature, function and safety class of the requirement the test case verifies |
| Comes from | derived - always from the requirement; a test case has none of its own, so there is no original, and neither has a test result |
| Prototype | renamed - from `feature`, `function` and `safety`, so that the name says the value is derived |

### priority

| | |
|---|---|
| Meaning | the priority of a test case - a test manager may prioritize test cases of their own |
| Comes from | author, often empty |
| Variants | `mpriority` - the mapped own priority, or where it is empty the priority of the requirement it verifies |
| Prototype | same |

### component

| | |
|---|---|
| Meaning | the component a test case belongs to - for the tests per component: the unit tests and the integration tests of a component |
| Comes from | derived - from the requirement the test case verifies; it also places the file of the test beside its unit or under its component. Empty for the tests of a subsystem and of the system |
| Note | the name stays `component`, though derived |
| Prototype | same |

### category, unit

| | |
|---|---|
| Original | empty - a test case carries neither |
| Prototype | same |

### status, source_ids, target_ids

| | |
|---|---|
| Meaning | as for a requirement |
| Prototype | same |

### Links

| | |
|---|---|
| Meaning | `testedBy` from a requirement to its test cases and test procedures; `testFor` back to the requirement; `resultOf` to the results of a test case; `assignedTo` to a component |
| Prototype | same - one link `testedBy` for test cases and test procedures; no link of its own, such as `specifiedBy`, and none between a test procedure and its test cases |

## Testresult

A test result is one run of a test case: its verdict, in a stream.

### level, stream, category

| | |
|---|---|
| Meaning | `level` the key `r<n>tr` - a result of a test case of level r<n>; `stream` the variant the run took place in; `category` empty |
| Prototype | same |

### status

| | |
|---|---|
| Meaning | the verdict of the run |
| Comes from | the source as it writes it - for a test in Go, `pass`, `fail` or `skip` of `go test -json` |
| Variants | `mstatus` - mapped: `passed`, `failed`; only `passed` counts as reached in the maturity (SYS_47) |
| Prototype | same in the model; its generator leaves the original empty - see the deviations of the generator |

### mfeature, mfunction, msafety, mpriority

| | |
|---|---|
| Meaning | the feature, function, safety class and priority of the test case the result came out of |
| Comes from | derived - always from the test case; a result has none of its own, so there is no original |
| Prototype | renamed - from `feature`, `function`, `safety`; changed - the prototype also copies the original `priority` of the test case |

### component

| | |
|---|---|
| Meaning | the component of the test case the result came out of, for the tests per component |
| Comes from | derived, from the test case |
| Prototype | new - a test result has no component in the prototype |

### Links

| | |
|---|---|
| Meaning | `resultFor` to the test case the result came out of |
| Prototype | same |

## Hazard

A hazard (h1) and the measure that controls it (h2) are objects of the same type, `Hazard`, told
apart by their level. A threat and its control follow the same shape.

### level

| | |
|---|---|
| Meaning | `h1` a hazard, `h2` a measure that controls one |
| Prototype | same |

### safety

| | |
|---|---|
| Meaning | the safety class |
| h1 | the original - the result of the hazard analysis, set by its author; `msafety` mapped from it |
| h2 | `msafety` only - always the class of its hazard; no original |
| Note | the safety class starts here and is carried down to the requirements and tests |
| Prototype | changed - the prototype copies the class of the hazard into the original of the measure |

### mfeature, mpriority, mplatform_launch

| | |
|---|---|
| Meaning | the feature, priority, platforms and launches a hazard or a measure answers for |
| Comes from | derived, upwards: a measure from the requirements it controls, a hazard from its measures; no original |
| Prototype | changed - the prototype writes the feature into the original |

### status

| | |
|---|---|
| Meaning | the state of the hazard or measure in its log - open, controlled, closed |
| Comes from | author - its own state, not derived from below |
| Variants | `mstatus` - mapped from the original |
| Prototype | changed - the prototype takes `mstatus` from the requirements below |

### function

| | |
|---|---|
| Meaning | the function a hazard belongs to |
| Comes from | author |
| Prototype | same |

### Links

| | |
|---|---|
| Meaning | `parentOf` from a hazard to its measures; `controls` from a measure to the requirements (r1) it asks for |
| Prototype | same |

## Normative

A clause of a standard or a law (n1). The module is the standard, the object the clause.

### level, category

| | |
|---|---|
| Meaning | `level` is `n1`; `category` is `Normative` |
| Prototype | same |

### module_id, object_key, object_text

| | |
|---|---|
| Meaning | the standard the clause belongs to - shown as *standard*; the key of the clause; its text |
| For our own data | one YAML file per standard, so that the module is the standard |
| Prototype | same |

### status

| | |
|---|---|
| Meaning | the state of the clause in its source - applicable, tailored |
| Comes from | author - its own, not derived from the requirements that address it |
| Variants | `mstatus` - mapped from the original; `cstatus` - computed by the reporter where needed: open, addressed or verified, from the requirements that address the clause and their results |
| Prototype | changed - the prototype takes `mstatus` from the first requirement that addresses the clause |

### applies_to, not_realized

| | |
|---|---|
| Meaning | `applies_to` the domain a clause applies to; `not_realized` why treporting does not realize it, where it does not |
| Prototype | new - for the regulatory requirements of treporting itself |

### Links

| | |
|---|---|
| Meaning | `addressedBy` from a clause to the requirements of any level that address it |
| Prototype | same |

## ReviewCriteria, ReviewResult

The review of a requirement against criteria, kept like a test case and its result.

| | |
|---|---|
| ReviewCriteria (RC) | the criteria of a level - after ISO/IEC/IEEE 29148: necessary, unambiguous, complete, consistent, verifiable, feasible, traceable |
| ReviewResult (RR) | one review: who, when, the revision of the requirement reviewed, per criterion passed or failed, the verdict |
| Links | `reviewedBy` from a requirement to the criteria of its level; `resultOf` from the criteria to the results |
| Note | apart from the types of the tests, though of the same shape: easier to understand, and tests and reviews keep tables and indexes of their own |
| Used by | SYS_74, SYS_75 |
| Prototype | new - the prototype reads review results from spreadsheets and shows them, with no type of its own |

## Project, System, Baseline, Artifact, Component

The types of the structure: projects, the systems they keep their work in, baselines, documents and
the architecture. They are no objects in a document, and have `name` and `text` instead of
`object_key` and `object_text`.

| Type | Attributes | category | Links | Prototype |
|---|---|---|---|---|
| Project | name, text, status, mstatus | `rproject`, `lproject` | `parentOf`, `childOf` - the project tree; `storedAt` a system; `consistOf` baselines | same |
| System | name, text, status, mstatus | `DOORS`, `GIT`, `SVN`, `DMS`, `JFROG`, ... | `assignedTo` a project; `consistOf` documents | same |
| Baseline | name, level, text, start_at, end_at, status, mstatus | `release`, `library` | `parentOf`, `childOf`; `assignedTo` a project; `consistOf` documents | same |
| Artifact - a document | system, module_id, revision, name, level, stream, platform, module_path, module_prefix, status, mstatus | the kind of document | - | renamed - `prefix` to `module_prefix`, the same thing |
| Component | name, level, stream, text, status, mstatus | `rproject`, `lproject`, `component`, `unit` | `assignedFrom` the requirements allocated to it; `parentOf`, `childOf` - the design tree | same |

A system is file-based or attribute-based (see the system architecture); that follows from its
category, by the configuration, and needs no attribute of its own.

## Deviations from the prototype

| What | How | Why |
|---|---|---|
| feature, function and safety of a requirement, taken from the parent where the own original is empty | changed: they go into the mapped value (`M`); the prototype writes them into the original - for priority and platform and launch it uses the mapped value already | the original stays original (ARCR_01, REG_FDADI_01), and it can be shown that an object was not classified by itself |
| `verification_method` | renamed from `coverage` | `coverage` had two meanings; ISO/IEC/IEEE 29148 |
| `coverage` | changed: the code coverage level | as above |
| `module_prefix` | changed: on every object in a document | an id is unique only within its document, and links name it |
| `created_at`, `created_by`, `modified_at`, `modified_by` | new, on every type | REG_FDADI_01 |
| `prefix` of a document | renamed to `module_prefix` | the same thing as on its objects |
| ReviewCriteria, ReviewResult, the link `reviewedBy` | new | the review of every requirement before it is approved |
| `feature`, `function`, `safety` of a test case | renamed to `mfeature`, `mfunction`, `msafety`, with no original | they are always derived from the requirement; the name says so |
| `feature`, `function`, `safety` of a test result | renamed to `mfeature`, `mfunction`, `msafety`, with no original | always derived from the test case |
| `priority` of a test result | changed: no original, `mpriority` only | a run cannot be prioritized of its own |
| `component` of a test result | new, derived from the test case | the tests per component |
| `safety` of a measure (h2) | changed: `msafety` only, from its hazard | a measure has no class of its own |
| `feature`, `priority`, `platform_launch` of a hazard or measure | changed: `mfeature`, `mpriority`, `mplatform_launch` only, derived upwards | they answer for what lies below |
| `status` of a hazard or measure | changed: its own original, `mstatus` mapped from it - not taken from below | a hazard log keeps its own state |
| `status` of a clause | changed: its own original, `mstatus` mapped from it - not taken from below; `cstatus` computed where needed | the compliance state is computed, not kept |
| `applies_to`, `not_realized` of a clause | new | the regulatory requirements of treporting itself |

## Deviations of the generator

Here the generator of the prototype differs from the model. Each point is still to be done when
the generator is taken over.

| What | The prototype | The model asks |
|---|---|---|
| `status` of a test result | writes `mstatus` only, the original stays empty | the original as a source would write it - `pass`, `fail`, `skip` - and `mstatus` mapped from it |
| `priority`, `status`, `platform_launch` of a hazard or measure | writes the originals from the mapped values (`settleHazard`) | the original as a source writes it, or empty |
| `status` of a clause | takes `mstatus` from the first requirement that addresses it, and writes the original from it | its own original, `mstatus` mapped from it |

## Open

- `treporting = 1_demo` - **The value `all` of platform_launch**: not set yet. The converter
  would resolve it into the platforms, by the configuration of the customer.
