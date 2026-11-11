# User needs (r1)

What is needed from treporting, without naming a solution. The needs are the reason for the tool
itself. The document is written for everyone, above all for whoever derives the system
requirements (r2) from them.

The needs themselves - ids `UN_<class>_nn` - are in [user-needs.yaml](user-needs.yaml), in the
format of [formats.md](../formats.md); the field `function` of a need names its class.

## Classes

**Decision**: a user need is grouped by the question it answers, its class - not by the role that
has it.

**Why**: roles differ from organization to organization. A need written for a role fits an
organization that has that role, and is missed by one that splits it or names it otherwise. The
question a need answers is the same in every organization.

**Rejected**: grouping by role (version 2.1.1 of this document) - one need belonged to several
roles, and the roles were those of one kind of organization. Grouping by activity (#576) -
activities differ from organization to organization as well.

Two kinds of class: what treporting answers about the data of a customer, and what treporting
itself has to be - the qualities of the product, after ISO/IEC 25010.

**Table classes** - each class, the field `function` of its needs, and the question they answer

| Class | Question |
|---|---|
| *What treporting answers* | |
| STATUS | How far has it got, what is open? |
| TRACE | What is linked to what? |
| COMPLIANCE | Was it done as the standard and the process demand? |
| CONFIG | What belongs to which baseline, what changed? |
| SCOPE | How are the data narrowed? |
| SHARE | How does a result reach others? |
| *What treporting has to be* | |
| EVIDENCE | Where does a figure come from, and can it be repeated? |
| SAFETY | Does what is shown never mislead? |
| SECURITY | Are the data and the access protected? |
| OPERATE | Does treporting run, and can it be maintained? |
| SETUP | Does it fit the systems of a customer? |
| USE | Is it understood without being told? |

Each user need is in one class. A need names no role: it says what is needed, not who needs it.

## SAFETY and SECURITY

**SAFETY** - what is relied on when deciding on what treporting shows. treporting controls
nothing, but a gap it hides can lead to a wrong release. The hazards behind these needs are in
`safety/hazard-and-risk-analysis.md`. If treporting gets a tool qualification, its analysis of the
tool's failures traces to these needs.

**SECURITY** - the protection of the data and of the access to them. If treporting gets a threat
analysis, its threats share level h1 with the hazards (`safety/hazard-and-risk-analysis.md`).

## Terms

- **record**: what treporting keeps as evidence - the data read in, their history, and who did
  what and when. A record can name a person; what it keeps of a person is what UN_SECURITY_06
  limits. The log that treporting writes while it runs is not a record (UN_OPERATE_06).

## Open

`<empty>`

## Change history

| Version | Date | Change | Reason | By |
|---|---|---|---|---|
| 1.1.1 | 2026-10-30 | - | #325 | Claude Opus 5.5 and Mathias Fischer, authors; Mathias Fischer, reviewer |
| 2.1.1 | 2026-10-30 | Safety of the product - PSF | #338 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 3.1.1 | 2026-11-11 | Classes instead of roles; SAFETY and SECURITY instead of PSF; Terms | #590, #602 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
