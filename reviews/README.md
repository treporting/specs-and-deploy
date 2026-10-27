# Reviews

The reviews of the requirements and documents of treporting: the criteria of every level, and the result of
every review. See the section Reviews of [formats.md](../formats.md) and the loop for every change
of [harness.md](../harness.md). It is written for whoever makes or checks a review, and for the
auditor who looks for its evidence.

**Table criteria** - the criteria file of each level

| File | What |
|---|---|
| [criteria-n1.yaml](criteria-n1.yaml) | the criteria of the regulatory requirements (n1): how a duty was written down, not the duty |
| [criteria-r1.yaml](criteria-r1.yaml) | the criteria of the user needs (r1) |
| [criteria-r2.yaml](criteria-r2.yaml) | the criteria of the system requirements (r2) |
| [criteria-r3.yaml](criteria-r3.yaml) | the criteria of the elements of the architecture (r3) |
| [criteria-r5.yaml](criteria-r5.yaml) | the criteria of the software requirements (r5), of the subsystems and of the tools |
| [criteria-tool.yaml](criteria-tool.yaml) | the criteria of the tool requirements |
| [criteria-documents.yaml](criteria-documents.yaml) | the criteria of the documents, after ISO15289 |

The format of a review and its result is in the section Reviews of [formats.md](../formats.md).
`reviewskeleton` writes the skeleton of a result, and `reviewcheck` checks every result, also in
the CI on every pull request - see [tools/README.md](../tools/README.md).

## The criteria of n1

**Decision**: a regulatory requirement is reviewed only for what a finding can change: how the duty
was written down - its source, its text, where it applies, who keeps it, how it is verified, its
form. Six criteria, in [criteria-n1.yaml](criteria-n1.yaml).

**Why**: the duty itself is set by the law or the standard; we cannot change it. What we write is
its source and its text in own words (regulatory-requirements.md), and there an error can be made
when it is written down - RC-01 for the source, RC-02 for the text.

**Rejected**: the criteria of r1 taken over - unambiguous, singular, complete, consistent, necessary,
feasible. They judge the duty, and a finding on it would change nothing we keep.

*Open:* the criteria of the other levels.
