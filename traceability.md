# Traceability

How a link between two levels stays valid when the parent changes. This document records the
discussion of ticket #249: the problem, the options weighed, and the decisions taken. The questions
still open are marked as such. The changes to the rules, the schemas and the tools follow in
tickets of their own.

## The problem: the cascade of PROC-13

PROC-13 says a requirement is approved only when "whatever it names as uplink is approved". Read as
a state that must hold at all times, every change above sets off a cascade below. TOOL-09 (#247)
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

DOORS and ReqIF keep a version per object too, apart from the revision of the store.

## The parts of the version

**Decision**: the version is `x.y.z`.

- **x** changes when the text changes, or an attribute the schema of the kind names for x
- **y** changes when another attribute changes
- **z** changes when a link changes

Where the schema of a kind names no attribute for x, x covers the text and `safety`.

**Why**: a child depends on the text of its parent and on what the parent means for safety. It does
not depend on a remark or on the links of its parent. So each part of the version shows in the
data which kind of change happened.

**Decision**: a link names the full version of its parent: `target_ids: [SYS-09@3.0.2]`.

**Why**: the link shows the exact state the child was checked against.

**Decision**: the status changes the version only where the schema of the kind names it for x.
Otherwise a change of the status changes nothing.

**Why**: a status that changed y would change the version at every approval.

**Decision**: a change of the text and of the status at the same time counts once. A commit raises
x by one at most.

**Why**: PROC-14 sets a changed text back to `draft`. Where the status counts for x, one change
would otherwise raise x twice.

**Decision**: each part only counts up. None of them is set back to 0 when another rises:
`3.0.2` becomes `4.0.2`.

**Why**: each part counts, over the whole life of the object, how often its kind of change
happened.

**Rejected**: to set y and z back to 0 when x rises, as SemVer does.

**Decision**: every object that exists today starts at `1.1.1`. Every link that exists today names
`@1.1.1`.

**Table cascade** - how far a change reaches

| What happens | UN-PM-01 | SYS-09 | SR-01 |
|---|---|---|---|
| Start | `2.0.0` | `3.0.1`, link `UN-PM-01@2.0.0` | link `SYS-09@3.0.1` |
| The text of UN-PM-01 changes | `3.0.0` | its link names an older x: **suspect** | clean |
| SYS-09 still fits; only its link is updated | | `3.0.2`, link `UN-PM-01@3.0.0` | clean, only z of SYS-09 changed |
| SYS-09 no longer fits; its text changes | | `4.0.2`, `draft` | its link names an older x: **suspect** |
| A remark on SYS-09 changes | | `4.1.2` | clean, only y of SYS-09 changed |

The cascade reaches only as far as meaning really changes. A child that still fits changes z only,
and its own children stay clean.

## Where a suspect link blocks

A suspect link is reported on every pull request. It blocks the baseline whose scope holds the
child of the link.

**Table baselines** - which suspect links block which baseline

| Baseline | Levels | Blocked, for example, by |
|---|---|---|
| Release | h1, h2, n1, r1 to r3 | SYS-09 → UN-PM-01, when UN-PM-01 changed its text |
| Library | r4 to r7, source code (sc1) | SR-01 → SYS-09, when SYS-09 changed its text |

So a changed SYS does not block the release of its own level. It blocks the next library of every
subsystem whose SRs name it, because there the links have to be updated.

**Why**: a pull request that blocked would bring the cascade back. A baseline must not hold a link
that no one has checked.

## When a link is suspect

**Decision**: only a different x makes a link suspect. A difference in y or z does not.

**Why**:

- x is defined to cover everything a child depends on: the text, and the attributes the schema of
  the kind names for x. An attribute that matters to the children belongs in x.
- our files keep uplinks only. A child writes its parent in `target_ids` (formats.md, section
  Fields). So the z of a parent counts the changes of its own uplinks, and no child depends on
  them. Updating a link changes the z of the child only, and so the cascade ends there.
- so anyone can read from the first number whether a link is suspect

**Rejected**:

- every difference makes a link suspect: y would mean the same as x, and every child would have
  to follow a change of `priority` of its parent
- down-links, a parent that names its children: then a new child would change the parent, and
  every child would see it

## Open questions

1. **The right side of the V.** Code and tests name their requirement in an annotation,
   `@implements SR-01`, not in `target_ids`. Either the annotation names the version too,
   `@implements SR-01@3.0.2`, or a tool computes from git whether the text of the SR changed after
   the line with the annotation.
2. **How the schema of a kind names the attributes for x.**
3. **The rule changes:** PROC-13 "at its approval"; PROC-14 for the parts of the version; a rule
   for suspect links.
4. **The tools:** a check that a change raises the right part of the version, a check for suspect
   links, and a command that updates a link to the current version of its parent.
5. **TOOL-09** (#248): SR-28 would read "a requirement that becomes approved while an uplink of it
   is not approved". #248 waits until the rules stand.
