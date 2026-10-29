# Standards

The standards of the domains treporting is used in: the process a project follows, the safety
standard, and the analysis that finds what is to be controlled. treporting takes them as the basis
of its development (see Process in [definitions.md](definitions.md)). It also sets the short form
by which every document names a standard. It is written for whoever writes or reviews a document
of treporting, and for the auditor who looks for a standard.

## Short forms

**Decision**: every document names a standard by its short form: its first publisher and its
number, with no space between them. A part keeps its hyphen: IEC 61508-1 is IEC61508-1. Only the
table below names a standard by its number and title.

**Why**: the full name, such as ISO/IEC/IEEE 12207, is long, and was written out each time. A
short form is the same in every document, and the table says what stands behind it.

**Rejected**:

- the full name throughout - too long, and found in the review of harness.md
- a short form with a hyphen, ISO-12207 - longer than needed, and the hyphen is taken by the part

**Table short forms** - every standard named in the documents

| Short form | Standard | Remark |
|---|---|---|
| ISO12207 | ISO/IEC/IEEE 12207 Systems and software engineering - Software life cycle processes | |
| ISO15026 | ISO/IEC/IEEE 15026 Systems and software engineering - Systems and software assurance | |
| ISO15289 | ISO/IEC/IEEE 15289 Systems and software engineering - Content of life-cycle information items (documentation) | |
| ISO29119 | ISO/IEC/IEEE 29119 Software and systems engineering - Software testing | |
| ISO29148 | ISO/IEC/IEEE 29148 Systems and software engineering - Life cycle processes - Requirements engineering | |
| ISO25010 | ISO/IEC 25010:2023 Systems and software engineering - Systems and software Quality Requirements and Evaluation (SQuaRE) - Product quality model | |
| ISO330xx | ISO/IEC 330xx Information technology - Process assessment | |
| ISO15504 | ISO/IEC 15504 Information technology - Process assessment | replaced by ISO330xx |
| ISO27001 | ISO/IEC 27001 Information security, cybersecurity and privacy protection - Information security management systems - Requirements | |
| ISO9001 | ISO 9001 Quality management systems - Requirements | |
| ISO8601 | ISO 8601 Date and time - Representations for information interchange | |
| ISO26262 | ISO 26262 Road vehicles - Functional safety | |
| ISO21434 | ISO/SAE 21434 Road vehicles - Cybersecurity engineering | |
| ISO13485 | ISO 13485 Medical devices - Quality management systems - Requirements for regulatory purposes | |
| ISO14971 | ISO 14971 Medical devices - Application of risk management to medical devices | |
| IEC62304 | IEC 62304 Medical device software - Software life cycle processes | |
| IEC60601-1 | IEC 60601-1 Medical electrical equipment - Part 1: General requirements for basic safety and essential performance | |
| IEC81001-5-1 | IEC 81001-5-1 Health software and health IT systems safety, effectiveness and security - Part 5-1: Security - Activities in the product life cycle | |
| IEC61508 | IEC 61508 Functional safety of electrical/electronic/programmable electronic safety-related systems | |
| IEC62443 | IEC 62443 Industrial communication networks - Network and system security | |
| EN50126 | EN 50126 Railway applications - The specification and demonstration of Reliability, Availability, Maintainability and Safety (RAMS) | |
| EN50716 | EN 50716 Railway applications - Requirements for software development | |
| EN50128 | EN 50128 Railway applications - Communication, signalling and processing systems - Software for railway control and protection systems | replaced by EN50716 |
| EN50657 | EN 50657 Railway applications - Rolling stock applications - Software on board rolling stock | replaced by EN50716 |
| EN50129 | EN 50129 Railway applications - Communication, signalling and processing systems - Safety related electronic systems for signalling | merges into IEC62425 |
| IEC62425 | DIN EN IEC 62425 Railway applications - Communication, signalling and processing systems - Safety related electronic systems for signalling | |
| CLC50701 | CLC/TS 50701 Railway applications - Cybersecurity | |
| ASPICE | Automotive SPICE Process Reference and Assessment Model | |
| ARP4754B | ARP4754B Guidelines for Development of Civil Aircraft and Systems | |
| ARP4761A | ARP4761A Guidelines for Conducting the Safety Assessment Process on Civil Aircraft, Systems, and Equipment | |
| DO178C | DO-178C Software Considerations in Airborne Systems and Equipment Certification | |
| DO330 | DO-330 Software Tool Qualification Considerations | |
| DO254 | DO-254 Design Assurance Guidance for Airborne Electronic Hardware | |
| DO326A | DO-326A Airworthiness Security Process Specification | |
| DO356A | DO-356A Airworthiness Security Methods and Considerations | |
| CMMI | CMMI Capability Maturity Model Integration | |

## Editions

**Decision**: the tables name a standard, a law or a guidance with its edition, as its publisher
designates it, such as ISO/IEC 25010:2023. A regulatory requirement names its clause in that
edition.

**Why**: the numbering of the clauses changes from one edition to the next. Without the edition,
a clause such as 6.1 does not say where it stands.

**Rejected**:

- a column Edition - each publisher gives its edition in another way: a year, a date, a letter
- the edition in the short form, ISO25010-2023 - the short form names a file and stands in an id; a
  new edition would rename both

## Laws, guidance and licences

A regulatory requirement comes from a standard, but also from a law, a guidance or a licence. Its
module is named by a short form as well (formats.md, section Regulatory requirements).

**Decision**: a law or a guidance with a number is named like a standard: its publisher and its
number, with no space between them, and a part with its hyphen. A slash in the number becomes a
hyphen. One with no number is named by its own short name. Only the table below names it in full.

**Why**: the same rule holds for every module, standard or law. The short form names a file and
stands in an id, where a slash would separate directories.

**Rejected**:

- the common names, GDPR and CRA: the rule of the standards gives the publisher and the number
- VO(EU)2016/679: VO is German, and the slash separates directories
- 21CFR11: the only short form to begin with its number, its part without a hyphen

**Table other short forms** - every law, guidance and licence named in the documents

| Short form | Source | Remark |
|---|---|---|
| CFR21-11 | 21 CFR Part 11 Electronic Records; Electronic Signatures | |
| EU2016-679 | Regulation (EU) 2016/679 General Data Protection Regulation | |
| EU2024-2847 | Regulation (EU) 2024/2847 Cyber Resilience Act | |
| DDG | Digitale-Dienste-Gesetz | its own short name |
| BetrVG | Betriebsverfassungsgesetz | its own short name |
| FDADI | FDA Data Integrity and Compliance With Drug CGMP - Guidance for Industry (2018) | no number |
| ATAM | CMU/SEI-2000-TR-004 ATAM: Method for Architecture Evaluation | its own short name |
| LASR | Team-Driven Architecture Evaluation with LASR, Stefan Toth and Stefan Zörner, ICSA 2026 | its own short name, no number |
| Licences | the licences of the components a release contains, such as Apache-2.0 | no publisher |

## Per domain

**Table domains** - the standards of each domain, at a glance

| Domain | Process | Safety | Security |
|---|---|---|---|
| automotive | ASPICE, ISO330xx | HARA ISO26262 | TARA ISO21434 |
| airborne | ARP4754B, DO178C, DO254 | FHA, SSA ARP4761A | SRA DO326A |
| railway | EN50126-1, EN50716, EN50129 | HL EN50126-1, EN50126-2 | TRA CLC50701 |
| medical | ISO13485, IEC62304, IEC60601-1 | RMF ISO14971 | TM IEC81001-5-1 |
| industrial | IEC61508 | H&RA IEC61508 | TRA IEC62443 |
| software | ISO12207, ISO29119 | H&RA ISO15026 | TRA IEC62443 |

- FHA - functional hazard assessment according to ARP4761A
- H&RA - hazard and risk analysis according to IEC61508-1, clause 7.4
- HARA - hazard analysis and risk assessment according to ISO26262-3
- HL - hazard log, with hazard analysis and risk assessment, according to EN50126-1
- RMF - risk management file according to ISO14971
- SRA - security risk assessment according to DO326A and DO356A
- SSA - system safety assessment according to ARP4761A
- TARA - threat analysis and risk assessment according to ISO21434
- TM - threat model according to IEC81001-5-1
- TRA - threat and risk assessment according to CLC50701 in railway, and IEC62443-3-2 in industrial
  and software
