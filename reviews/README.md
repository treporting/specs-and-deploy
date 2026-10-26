# Reviews

The reviews of the requirements and documents of treporting: the criteria of every level, and the result of
every review. See the section Reviews of [formats.md](../formats.md) and the loop for every change
of [harness.md](../harness.md). It is written for whoever makes or checks a review, and for the
auditor who looks for its evidence.

**Table criteria** - the criteria file of each level

| File | What |
|---|---|
| [criteria-r1.yaml](criteria-r1.yaml) | the criteria of the user needs (r1) |
| [criteria-r2.yaml](criteria-r2.yaml) | the criteria of the system requirements (r2) |
| [criteria-r5.yaml](criteria-r5.yaml) | the criteria of the software requirements (r5), of the subsystems and of the tools |
| [criteria-tool.yaml](criteria-tool.yaml) | the criteria of the tool requirements |
| [criteria-documents.yaml](criteria-documents.yaml) | the criteria of the documents, after ISO15289 |

The format of a review and its result is in the section Reviews of [formats.md](../formats.md).
`reviewskeleton` writes the skeleton of a result, and `reviewcheck` checks every result, also in
the CI on every pull request - see [tools/README.md](../tools/README.md).

*Open:* the criteria of the other levels.
