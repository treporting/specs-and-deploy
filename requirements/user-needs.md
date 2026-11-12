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
roles, and the roles were those of one kind of organization. Grouping by activity - activities
differ from organization to organization as well.

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

- **record**: in the user needs, what treporting keeps as evidence - the data read in, their history, and who did
  what and when. A record can name a person; what it keeps of a person is what UN_SECURITY_06
  limits. The log that treporting writes while it runs is not a record (UN_OPERATE_06).
- **finding**: something found wrong in a work product, such as a requirement, a test or code. Someone
  evaluates it: they decide whether it is a defect and what is to be done. The evaluation is written
  down somewhere. There are three kinds:
  - **of the customer**: made in the customer's process, by a review, an assessment or a tool such as a
    static analysis. It is kept in the customer's tools. treporting shows it and its evaluation, but
    makes and evaluates none (UN_COMPLIANCE_02, UN_COMPLIANCE_03).
  - **of treporting**: treporting finds it in the data of the customer, for instance a rule of
    plausibility that does not hold (UN_SAFETY_03).
  - **of our own process**: a review of treporting finds it. It is not meant in the user needs.
- **report**: a document made of views of the data, to be read without the system. There are three
  kinds:
  - **of treporting**: treporting makes it of the views chosen, from the data of the customer
    (UN_SHARE_01, UN_SHARE_04, UN_SHARE_06).
  - **of the customer**: made in the customer's own process or tools. It is not meant in the user
    needs.
  - **of our own process**: made by our tools about treporting, such as a test report. It is not
    meant in the user needs.

  A report is not a record. A report is made from the data at one moment; a record is what
  treporting keeps as evidence.

## Open

`<empty>`

## Change history

| Version | Date | Change | Reason | By |
|---|---|---|---|---|
| 1.1.1 | 2026-10-30 | - | #325 | Claude Opus 5.5 and Mathias Fischer, authors; Mathias Fischer, reviewer |
| 2.1.1 | 2026-10-30 | Safety of the product - PSF | #338 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 3.1.1 | 2026-11-11 | Classes instead of roles; SAFETY and SECURITY instead of PSF; Terms | #590, #602 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 4.1.1 | 2026-11-12 | Terms: finding | #614 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 5.1.1 | 2026-11-12 | Terms: report | #617, #611 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
