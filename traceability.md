# Traceability

How a link between two levels stays valid when the parent changes. This document records the
problem, the options weighed, and the decisions taken. The questions still open are marked as such.
The rules that follow from it stand in harness.md: PROC-13, PROC-14 and PROC-20. The changes to
the schemas and the tools follow in changes of their own. It is written
for whoever changes a requirement or checks its trace, and for the auditor who looks for the
reasons.

## The problem: the cascade of PROC-13

PROC-13 says a requirement is approved only when "whatever it names as uplink is approved". Read as
a state that must hold at all times, every change above sets off a cascade below. TOOL-09
would check exactly that state, and so it showed the problem:

1. TOOL-03 fails its review and goes to `draft`.
2. SR-01 to SR-08 break PROC-13, because their uplink is no longer approved.
3. The code that names SR-01 to SR-08 fails annotationcheck, because the IDs are no longer
   approved.
4. A small finding on TOOL-03 blocks the whole tool.

The industry knows this trap under the name **suspect links**. DOORS has them built in.

## The options

**Table options** - the ways out that were weighed

| | Option | Core | Advantage | Drawback |
|---|---|---|---|---|
| a) | A status for the link: suspect | A change above marks the links below | common in the industry (DOORS), precise | a new state for links; a tool must find it |
| b) | PROC-13 at the moment of approval | An approval needs approved uplinks; later it does not | simple, no cascade | alone, it does not say what follows a change above |
| c) | Strict cascade | Children follow the status of their parents at once | PROC-13 as written | every small change blocks everything below |
| d) | Computed status | The own status stays; a computed status shows "approved, but an uplink is not" | uses what the model has: an original and a derived value | misses a parent that changed and was approved again |
| e) | Uplink with a revision | The child names the version of the parent it was approved against | precise, like a pinned dependency | every change above needs an update below |
| f) | Only a change of meaning counts | PROC-14 worded with "changes its meaning" | editorial changes set off nothing | eases the problem, does not solve it; who decides "meaning"? |
| g) | Consistency only at the baseline | Between baselines the state may be inconsistent | little load in daily work | gaps stay unnoticed for long |
| h) | A failed review never reaches `main` | Only the result is merged; `main` keeps the approved version | `main` is always consistent | against the rule "failed sets `draft` on `main`"; a successful change still raises the question whether the children fit |
| i) | Impact assessment in the same pull request | Whoever changes a parent checks every child in the same pull request | no state in between | large pull requests where a parent has many children |

**Decision**: b) + e) + i).

- **b)** An approval needs approved uplinks at the moment of approval. A later change of an uplink
  does not take the approval back. TOOL-09 checks this at the approval.
- **e)** Every link names the version of its parent. A link whose version is outdated is suspect.
- **i)** Whoever changes a parent updates the links of its children in the same pull request where
  that is cheap.

**Why**:

- b) ends the cascade: a parent in `draft` no longer breaks its approved children
- e) finds the dangerous case: a parent changes its meaning, is approved again, and a child no
  longer fits it
- updating the version on the link is the record of the impact assessment: the commit shows who
  checked the child against which version, and when

**Rejected**:

- c) blocks everything after every small change
- d) alone: it asks only whether the uplink is approved **now**. It misses a parent that changed
  and was approved again (Table d versus e).
- f) eases the problem but does not solve it
- g) alone: gaps stay unnoticed until the baseline
- h) keeps `main` consistent but does not answer whether the children fit a parent that changed
  with success; it would also need a) or e)
- a) as a state of its own on the link: e) shows the same in the version the link names, and needs
  no state that someone has to set and clear

**Table d versus e** - TOOL-03 changes its text and is approved again; does SR-01 still fit?

| What happens | d) computed status | e) version on the link |
|---|---|---|
| TOOL-03 fails, `draft` | found: "SR-01 approved, uplink not" | found: TOOL-03 changed |
| TOOL-03 is revised and approved again, with another text | **missed**: all approved, SR-01 never checked against the new text | found: SR-01 names an older version |
| TOOL-03 changes only editorially and is approved again | missed, which is right here | found: a short impact assessment is due |

## The version a link names

The first idea for e) named a git revision on the link, `TOOL-03@c3`, meaning the last commit that
changed the text of TOOL-03.

**Decision**: every object carries a version of its own, `x.y.z`. A link names the version of its
parent, not a git revision.

**Why**: a revision means a commit everywhere else. On the link it would mean something else: the
last commit that changed the text, not the last commit that changed the object. `TOOL-03@c3` does
not show this. Only a tool could tell when a new revision makes a child suspect and when it does
not. The meaning has to be visible in the data.

**Rejected**:

- a git revision on the link: ambiguous, as above
- a revision computed from git and the review results, with nothing written into the link: no one
  sees in the file which version a child was approved against
- a single number for the text: it says nothing about a change of an attribute such as safety
- a hash of the text, `SYS-09#a1b2`: it changes by itself, but it is hard to read, and it does not
  show which version is newer

## The parts of the version

**Decision**: the version is `x.y.z`.

- **x** changes when the text changes, or an attribute the schema of the kind names for x
- **y** changes when a link changes
- **z** changes when another attribute changes

Where the schema of a kind names no attribute for x, x covers the text and `safety`.

**Why**: a child depends on the text of its parent and on what the parent means for safety. It does
not depend on a remark or on the links of its parent. So each part of the version shows in the
data which kind of change happened.

The parts stand in the order of their weight. A change of a link changes what the object is derived
from, its place in the trace; a change of another attribute, such as a remark or `priority`, does
not.

**Rejected**: y for the other attributes and z for the links. The links would stand below a remark.

**Decision**: the version is set once, at the approval - the act that closes the review. Each part
rises by one where its kind changed since the last approved version; the others stay. While an
object is a draft, its version stays as it is.

**Why**: a child is checked against an approved parent. The working steps of a draft are no state
anyone checks against.

**Rejected**: to raise the version at every change. A draft would count its working steps, and a
child could name a state that was never approved.

**Decision**: only a change of x sets an approved object back to `draft`, and needs a review. A
change of y or z only - a link updated, a remark - leaves it `approved`. Its merge is its approval,
as for a rule, and raises y or z by one. The author makes the impact assessment - does the child
still fit the new version of its parent? - and the Product Owner confirms it by the merge.

**Why**: x is there to set apart a change of meaning. What does not change the meaning needs no
full review. So updating a link stays cheap, and the cascade stays away.

**Rejected**: every change sets back to `draft` and needs a full review. Every updated link would
be checked again against all criteria.

**Decision**: a link names the full version of its parent: `source_ids: [SYS-09@3.1.2]`.

**Why**: the link shows the exact state the child was checked against.

**Rejected**: only x on the link, `SYS-09@3`. It names what the child depends on, but not the state
the child saw.

**Decision**: the status changes the version only where the schema of the kind names it for x.
Otherwise a change of the status changes nothing.

**Why**: a status that changed z would change the version at every approval.

**Rejected**: the status as one of the other attributes, counted in z.

**Decision**: an approval raises each part by one at most, however many of its attributes
changed - the text and the status as well, where the status counts for x.

**Why**: each part counts how often its kind of change was approved, not how many attributes it
touched.

**Rejected**: to count each changed attribute. One approval would then raise x by as many as it
touched.

**Decision**: each part only counts up. None of them is set back to 0 when another rises:
`3.1.2` becomes `4.1.2`.

**Why**: each part counts, over the whole life of the object, how often its kind of change
happened.

**Rejected**: to set y and z back to 0 when x rises, as SemVer does.

**Decision**: every object that exists today starts at `1.1.1`. Every link that exists today names
`@1.1.1`.

**Why**: an object is in its first version from the moment it is written - a draft as well as an
approved one. The three parts count independently, so each of them is in its first version too:
text, links and other attributes all have their first state.

**Rejected**:

- `1.0.0`, as SemVer starts: it would say the links and the other attributes have no state yet,
  though they have
- `0.0.0`: it would say nothing has a state yet, though the object is approved

**Decision**: a new object carries `1.1.1` by default, as a draft. Its first approval keeps
`1.1.1`: there is no approved version before it to compare with.

**Why**: the draft is its first version. The status changes the version nowhere, so not here
either. No child can name the draft as uplink at its own approval, because an approval needs
approved uplinks.

**Rejected**: `0.0.0` while the object is a draft, and `1.1.1` at its first approval.

- it would make an exception to the rule of the status
- the objects that are drafts today would have to start at `0.0.0`, against the start of today's
  objects at `1.1.1`

**Table cascade** - how far a change reaches; every change ends with its approval

| What happens | UN-PM-01 | SYS-09 | SR-01 |
|---|---|---|---|
| Start | `2.1.1` | `3.2.1`, link `UN-PM-01@2.1.1` | link `SYS-09@3.2.1` |
| The text of UN-PM-01 changes, approved | `3.1.1` | its link names an older x: **suspect** | clean |
| SYS-09 still fits; only its link is updated, approved | | `3.3.1`, link `UN-PM-01@3.1.1` | clean, only y of SYS-09 changed |
| SYS-09 no longer fits; its text changes, approved | | `4.3.1` | its link names an older x: **suspect** |
| A remark on SYS-09 changes, approved | | `4.3.2` | clean, only z of SYS-09 changed |

The cascade reaches only as far as meaning really changes. A child that still fits changes y only,
and its own children stay clean.

## Where a suspect link blocks

A baseline is not the state of `main`. It is the set of versions it is made of (definitions.md,
section Baseline). A release holds its levels and the libraries it delivers; a library holds its
levels and its code.

**Decision**: a suspect link is reported on every pull request. It blocks the baseline that holds
both its ends. It is checked against the version of its parent in that baseline, not on `main`.

**Table baselines** - which baseline a suspect link blocks

| Link | Child in | Parent in | Blocked |
|---|---|---|---|
| SYS-09 → UN-PM-01 | release | release | the release |
| `@implements SR-01` in the code | library | library | the library: its code is not checked against its own SRs |
| SR-01 → SYS-09 | library | release | not the library; the release that delivers the library and holds SYS-09 in its new version |
| HAZC-07 `controls` UN-SAF-03 | release | release | the release |

**Table maintenance** - a fix for an older release, while `main` has moved on

| | specs-and-deploy | reporter |
|---|---|---|
| Release 1 | SYS-09 `3.1.2` | library 1.3, SR-01 → `SYS-09@3.1.2` |
| afterwards on `main` | SYS-09 `4.1.2`, for release 2 | |
| Release 1.1, a fix | SYS-09 `3.1.2`, the state of release 1 | library 1.3.1, SR-01 → `SYS-09@3.1.2`: clean |
| Release 2 | SYS-09 `4.1.2` | a library whose SR-01 names `SYS-09@4.x.x`; else blocked |

So a library can still be built on the software requirements approved so far, while the system
requirements change. Only a release that promises the new version waits until its libraries
follow.

**Why**:

- a pull request that blocked would bring the cascade back
- a baseline must not hold a link that no one has checked
- a contradiction exists only where both ends of a link come together. A library on the approved
  software requirements says nothing about a newer system requirement.
- checked against `main`, every fix of an existing library would be blocked

**Rejected**:

- to block the pull request: the cascade would be back
- to block nothing: a release could deliver a library that no one checked against its version
- to block the baseline that holds the child: a changed system requirement would block every
  library, also a fix of an existing one
- to check against `main`: a baseline is not the state of `main`

## When a link is suspect

The link names the full version, such as `SYS-09@3.1.2`. The parent may have moved on in y or z
only.

**Table suspect** - a link `SYS-09@3.1.2` against the version SYS-09 has now

| SYS-09 now | What changed | Suspect? |
|---|---|---|
| `3.2.2` | a link of SYS-09 | no |
| `3.1.3` | a remark | no |
| `4.1.2` | the text, or `safety` | yes |

**Decision**: only a different x makes a link suspect. A difference in y or z does not.

A parent in `draft` keeps its approved version. Its links turn suspect only when it is approved
with a new x.

**Why**:

- x is defined to cover everything a child depends on: the text, and the attributes the schema of
  the kind names for x. An attribute that matters to the children belongs in x.
- within a tree, our files keep uplinks only. A child writes its parent in `source_ids` (formats.md, section
  Fields are the properties of the gel schema). So the y of a parent counts the changes of its own
  uplinks, and no child depends on them. Updating a link changes the y of the child only, and so
  the cascade ends there.
- so anyone can read from the first number whether a link is suspect

**Rejected**:

- every difference makes a link suspect: z would mean the same as x, and every child would have
  to follow a change of `priority` of its parent. With y, the grandchildren would have to follow
  every update of a link of the child, and the cascade would run to the bottom.
- down-links, a parent that names its children: then a new child would change the parent, and
  every child would see it

## The right side of the V

Code and tests name their requirement in an annotation, not in `source_ids`.

**Decision**: an annotation names the version of its requirement, as a link does:
`// @implements SR-01@3.1.2`. It is suspect by the same rule: when x of the requirement differs.

**Why**:

- one rule for both sides of the V
- the version is visible in the file, as on a link
- updating the line is the record of the impact assessment: the commit shows who checked the code
  against which version, and when
- the cascade ends here: nothing lies below code and tests

**Rejected**: a tool that computes from git whether x of the requirement rose after the line was
last changed. No one would have to update anything, but no one would see in the file which version
the code was checked against - the same reason the git revision on a link was rejected.

## The version in the schema

**Decision**: every object carries its version in a field `object_version`, `x.y.z`. The schema of
its kind requires it.

**Why**: the fields fall into three groups.

**Table groups of fields** - which fields carry the prefix `object_`

| Group | Fields | What they say | Prefix |
|---|---|---|---|
| Place | `organization`, `project`, `system`, `module_id`, `baseline`, `library`, `revision` | where the object lies; all objects of a module share them | no |
| The object itself | `object_id`, `object_key`, `object_text`, `object_hid`, `object_heading` | who it is and what it says; every object has its own | `object_` |
| Attributes | `level`, `feature`, `status`, `safety`, … | properties to filter and compute by | no |

The version belongs to the object itself. Every object has its own, so it is no place. It counts
the changes of the attributes, so it is no attribute. With `object_key` it names one state of the
object, and that is what a link names: `SYS-09@3.1.2` is `object_key@object_version`.

**Rejected**: `version`. Without the prefix it would belong to the place, beside `revision` and
`baseline`: the version of the module or document the object lies in. Both fields could stand
side by side without contradicting each other - which shows that `version` alone leaves open whose
version it is.

**Decision**: the schema of every kind names the attributes for x in a list beside the properties
of the object, `attributes_for_version_x`. The list is always written out, also where it holds the
default.

```json
"items": {
  "type": "object",
  "attributes_for_version_x": ["object_text", "safety"],
  "properties": {
    "object_version": { "type": "string", "pattern": "^[0-9]+\\.[0-9]+\\.[0-9]+$" },
    …
  }
}
```

`source_ids` and `allocated` count for y. `status` counts for nothing unless the list names it. Every other field
counts for z.

**Why**:

- x can be read at one place, at a glance
- written out, x is visible in the schema itself; no one has to know a default
- JSON Schema 2020-12 ignores a keyword it does not know. schemacheck can check that every name in
  the list is a property of the kind.

**Rejected**: a mark on each attribute, `"safety": { …, "versionPart": "x" }`. Then x is found only
by searching every field.

## The allocation

The allocation (r4) is the field `allocated` of each subsystem in the architecture. The system
requirements allocated to a subsystem are its requirements, derived directly from r2
(system-architecture.md). So `allocated` is an uplink: the element stands below the system
requirements and names its parents. Only its name differs from `source_ids`.

**Decision**: `allocated` names the version of each system requirement, as a link does:
`allocated: [SYS-01@1.1.1, SYS-02@1.1.1, …]`. The same rules hold: a different x makes an entry
suspect, and updating it raises y of the element. Both ends lie in the release, so a suspect entry
blocks the release.

**Table allocation** - what only the allocation finds

| | Allocation | Software requirements |
|---|---|---|
| Before | SYS-09 `3.1.2` → viewer, reporter | reporter: SR-01 → `SYS-09@3.1.2` |
| SYS-09 becomes `4.1.2`: "… and exports the selection" | Does the export still belong to viewer and reporter, or does it need converter as well? | SR-01 turns suspect; the reporter checks its own part |

**Why**: the software requirements of a subsystem can tell only whether their own part still fits.
Whether another subsystem has to join, only the allocation shows.

**Rejected**: the allocation without versions. The case of the table would go unnoticed. The lists
are long, and versions make them harder to read; a view in the reporter can help, without changing
the data.

## Links between trees

A control of a hazard (h2) names the user needs that realize it, in `controls`. That link does not
run within the tree of the requirements: it joins two trees, the one of safety (hazard → control)
and the one of the requirements (user need → system requirement → …). The user needs stay the top
of their tree, with no uplink (attributes.md: `controls` from a measure to the requirements it asks
for).

**Decision**: `controls` is a down-link, from the control to its user needs, and names their
versions: `controls: [UN-PSF-01@1.1.1]`. It is suspect where x of a user need changed: the control
then checks whether it is still realized. A change of x of the control does not make its user
needs suspect; its review covers its list `controls`.

**Why**:

- the user needs stay the top of their tree; a link from each of them to a control would make
  them children of the tree of safety
- the control answers for being realized, so a change of its user needs is what it has to see
- the reason against down-links within a tree - a new child would change the parent - no longer
  holds: a new user need changes y of the control only, and no one sees y

**Rejected**: `controlledBy` on every user need, an uplink to the control.

**Decision**: PROC-13 asks approved uplinks only. A control may be approved while a user need it
names in `controls` is not. Whether a control is realized is checked on its own.

**Why**: the approval of a control is about what it asks for. That it is realized is another
question, and it has to be answered on its own.

**Rejected**: approved user needs at the approval of a control, as for uplinks.

**Decision**: PROC-12 follows the links, not a list of levels: an object is merged before another
names it.

**Why**: a down-link runs the other way. The user needs are merged before a control names them in
`controls`, though the control stands above them. A list of levels would have to say so for every
link between trees.

**Rejected**: a list of levels with the hazards added. It would hold for uplinks only.

## Links to rules

Requirements name rules of the harness as uplink, such as `source_ids: [PROC-12]`. Before this
decision, a rule had no version. Where a rule changed its meaning, a rule of its own, PROC-17, sent
every requirement that named it back to `review` - the strict cascade c).

**Decision**: a rule carries a version, `x.y.z`, as every requirement does. A link names it:
`source_ids: [PROC-12@1.1.1]`. The version is set at the merge, because the merge is the approval
of a rule. PROC-17 goes: a suspect link does its work.

**Why**: rules are requirements under another name. One rule holds for every uplink, and the
cascade ends for rules as well.

**Rejected**: rules without a version, and PROC-17 kept. Two ways would stand side by side, and
for rules the cascade would go on.

**Decision**: the rules become data, as the requirements are: a YAML file with a schema, each rule
with `object_key`, `object_text`, `object_version` and its uplinks in `source_ids`. harness.md
explains them and refers to it, as formats.md sets out for every level.

**Why**: rules are requirements under another name, so they are kept as requirements are. The tools
read them as any other requirement, with the same check of version and links.

**Rejected**: a column for the version in the table of harness.md. A tool would have to read a
Markdown table, and the rules would be kept apart from every other requirement.

## Open questions

1. **The tools:** a check that an approval sets the right version, a check for suspect
   links and annotations, and a command that updates a link or an annotation to the current
   version of its parent. annotationcheck has to read `@x.y.z`.
2. **TOOL-09:** its software requirements follow the rules as they stand now. SR-28 would read
   "a requirement that becomes approved while an uplink of it is not approved".
3. **Maintenance:** a fix for an older release needs the state of specs-and-deploy of that release,
   from its tag, not `main`. How such a state is made and kept is not set yet.
4. **The realization of a control:** how it is checked that a control is realized by its user
   needs, and when.
