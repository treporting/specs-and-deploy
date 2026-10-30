# Reviews

The reviews of the requirements and documents of treporting: the criteria of every level, and the result of
every review. See the section Reviews of [formats.md](../formats.md) and the loop for every change
of [harness.md](../harness.md). It is written for whoever makes or checks a review, and for the
auditor who looks for its evidence.

**Table criteria** - the criteria file of each level

| File | What |
|---|---|
| not set yet | the criteria of the hazards and threats (h1) |
| not set yet | the criteria of the controls (h2) |
| [criteria-n1.yaml](criteria-n1.yaml) | the criteria of the regulatory requirements (n1): how a duty was written down, not the duty |
| [criteria-r1.yaml](criteria-r1.yaml) | the criteria of the user needs (r1) |
| [criteria-r2.yaml](criteria-r2.yaml) | the criteria of the system requirements (r2) |
| [criteria-r3.yaml](criteria-r3.yaml) | the criteria of the elements of the architecture (r3) |
| [criteria-r5.yaml](criteria-r5.yaml) | the criteria of the software requirements (r5), of the subsystems and of the tools |
| not set yet | the criteria of the software architecture of a subsystem (r6) |
| not set yet | the criteria of the software design of a subsystem (r7) |
| [criteria-tool.yaml](criteria-tool.yaml) | the criteria of the tool requirements |
| not set yet | the criteria of the rules of our work (rules) |
| [criteria-documents.yaml](criteria-documents.yaml) | the criteria of the documents, after ISO15289 |

The results lie in [results/](results/), one file per review.

The format of a review and its result is in the section Reviews of [formats.md](../formats.md).
`reviewskeleton` writes the skeleton of a result, and `reviewcheck` checks every result, also in
the CI on every pull request - see [tools/README.md](../tools/README.md).

## The criteria of n1

**Decision**: a regulatory requirement is reviewed only for what a finding can change: how the duty
was written down - its source, its text, where it applies, who keeps it, how it is verified, its
form. Six criteria, in [criteria-n1.yaml](criteria-n1.yaml).

**Why**: the duty itself is set by the law or the standard; we cannot change it. What we write is
its source and its text in own words (regulatory-requirements.md), and there an error can be made
when it is written down - RC_01 for the source, RC_02 for the text.

**Rejected**: the criteria of r1 taken over - unambiguous, singular, complete, consistent, necessary,
feasible. They judge the duty, and a finding on it would change nothing we keep.

## The change history of a document

**Decision**: the change history of a document (formats.md, section Change history) names no review
result. Its row is written with the change itself, before the review, so that the review covers it.

**Why**: a row added after the review would change the document after it was reviewed - the content
approved would not be the content reviewed, and its hash would not match. The review result names
the document and the revision it reviewed: the link runs from the review to the document.

**Rejected**:

- a column Review with the file of the review result - the review result names the document
  already
- the row written by the commit of the decision - it changes the document after its review

## Open

- `treporting = 1_tools` - **Criteria**: not set yet.

## Change history

| Version | Date | Change | Reason | By |
|---|---|---|---|---|
| 1.1.1 | 2026-10-28 | - | #272 | Claude Opus 5.5 and Mathias Fischer, authors; Mathias Fischer, reviewer |
| 2.1.1 | 2026-10-29 | Reviews | #307 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 3.1.1 | 2026-10-29 | The criteria of n1 | #302 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 4.1.1 | 2026-10-29 | Open | #322 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 5.1.1 | 2026-10-30 | Reviews; Open | #322 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 6.1.1 | 2026-10-30 | The change history of a document; Change history | #323 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
