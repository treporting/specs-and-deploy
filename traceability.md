# Traceability

How a link between two levels stays valid when the parent changes. This document records the
problem, the options weighed, and the decisions taken. The questions still open are marked as such.
The changes to the rules, the schemas and the tools follow in changes of their own. It is written
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
| f) | Only a change of meaning counts | PROC-14 worded like PROC-17 | editorial changes set off nothing | eases the problem, does not solve it; who decides "meaning"? |
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

**Decision**: a link names the full version of its parent: `target_ids: [SYS-09@3.1.2]`.

**Why**: the link shows the exact state the child was checked against.

**Rejected**: only x on the link, `SYS-09@3`. It names what the child depends on, but not the state
the child saw.

**Decision**: the status changes the version only where the schema of the kind names it for x.
Otherwise a change of the status changes nothing.

**Why**: a status that changed z would change the version at every approval.

**Rejected**: the status as one of the other attributes, counted in z.

**Decision**: a change of the text and of the status at the same time counts once. A commit raises
x by one at most.

**Why**: PROC-14 sets a changed text back to `draft`. Where the status counts for x, one change
would otherwise raise x twice.

**Rejected**: to count each changed attribute. One change would then raise x by as many as it
touched.

**Decision**: each part only counts up. None of them is set back to 0 when another rises:
`3.1.2` becomes `4.1.2`.

**Why**: each part counts, over the whole life of the object, how often its kind of change
happened.

**Rejected**: to set y and z back to 0 when x rises, as SemVer does.

**Decision**: every object that exists today starts at `1.1.1`. Every link that exists today names
`@1.1.1`.

**Why**: an approved object is in its first version. The three parts count independently, so each
of them is in its first version too: text, links and other attributes all have their first state.

**Rejected**:

- `1.0.0`, as SemVer starts: it would say the links and the other attributes have no state yet,
  though they have
- `0.0.0`: it would say nothing has a state yet, though the object is approved

**Decision**: a new object starts at `1.1.1` when it is written, as a draft.

**Why**: the draft is its first version. The status changes the version nowhere, so not here
either. No child can name the draft as uplink at its own approval, because an approval needs
approved uplinks.

**Rejected**: `0.0.0` while the object is a draft, and `1.1.1` at its first approval. It would make
an exception to the rule of the status.

**Table cascade** - how far a change reaches

| What happens | UN-PM-01 | SYS-09 | SR-01 |
|---|---|---|---|
| Start | `2.1.1` | `3.2.1`, link `UN-PM-01@2.1.1` | link `SYS-09@3.2.1` |
| The text of UN-PM-01 changes | `3.1.1` | its link names an older x: **suspect** | clean |
| SYS-09 still fits; only its link is updated | | `3.3.1`, link `UN-PM-01@3.1.1` | clean, only y of SYS-09 changed |
| SYS-09 no longer fits; its text changes | | `4.3.1`, `draft` | its link names an older x: **suspect** |
| A remark on SYS-09 changes | | `4.3.2` | clean, only z of SYS-09 changed |

The cascade reaches only as far as meaning really changes. A child that still fits changes y only,
and its own children stay clean.

## Where a suspect link blocks

**Decision**: a suspect link is reported on every pull request. It blocks the baseline whose scope
holds the child of the link.

**Table baselines** - which suspect links block which baseline

| Baseline | Levels | Blocked, for example, by |
|---|---|---|
| Release | h1, h2, n1, r1 to r3 | SYS-09 → UN-PM-01, when UN-PM-01 changed its text |
| Library | r4 to r7, source code (sc1) | SR-01 → SYS-09, when SYS-09 changed its text |

So a changed SYS does not block the release of its own level. It blocks the next library of every
subsystem whose SRs name it, because there the links have to be updated.

**Why**: a pull request that blocked would bring the cascade back. A baseline must not hold a link
that no one has checked.

**Rejected**: to block the pull request; the cascade would be back. To block nothing; a release or a
library could hold a link that no one checked.

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

**Why**:

- x is defined to cover everything a child depends on: the text, and the attributes the schema of
  the kind names for x. An attribute that matters to the children belongs in x.
- our files keep uplinks only. A child writes its parent in `target_ids` (formats.md, section
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

Code and tests name their requirement in an annotation, not in `target_ids`.

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

`target_ids` counts for y. `status` counts for nothing unless the list names it. Every other field
counts for z.

**Why**:

- x can be read at one place, at a glance
- written out, x is visible in the schema itself; no one has to know a default
- JSON Schema 2020-12 ignores a keyword it does not know. schemacheck can check that every name in
  the list is a property of the kind.

**Rejected**: a mark on each attribute, `"safety": { …, "versionPart": "x" }`. Then x is found only
by searching every field.

## Open questions

1. **The rule changes:** PROC-13 "at its approval"; PROC-14 for the parts of the version; a rule
   for suspect links.
2. **The tools:** a check that a change raises the right part of the version, a check for suspect
   links and annotations, and a command that updates a link or an annotation to the current
   version of its parent. annotationcheck has to read `@x.y.z`.
3. **TOOL-09:** its software requirements wait until the rules stand. SR-28 would read "a
   requirement that becomes approved while an uplink of it is not approved".
