# Standards

The standards of the domains treporting is used in: the process a project follows, the safety
standard, and the analysis that finds what is to be controlled. A company process has to match
the standardized one (see Process in [definitions.md](definitions.md)).

## Short forms

Every document names a standard by its short form: its first publisher and its number, with no
space between them. A part keeps its hyphen: IEC 61508-1 is IEC61508-1. Only the table below names
a standard by its number and title.

**Table short forms** - every standard named in the documents

| Short form | Standard | Remark |
|---|---|---|
| ISO12207 | ISO/IEC/IEEE 12207 Systems and software engineering - Software life cycle processes | |
| ISO15026 | ISO/IEC/IEEE 15026 Systems and software engineering - Systems and software assurance | |
| ISO15289 | ISO/IEC/IEEE 15289 Systems and software engineering - Content of life-cycle information items (documentation) | |
| ISO29119 | ISO/IEC/IEEE 29119 Software and systems engineering - Software testing | |
| ISO29148 | ISO/IEC/IEEE 29148 Systems and software engineering - Life cycle processes - Requirements engineering | |
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
| IEC62443 | IEC 62443 Security for industrial automation and control systems | |
| EN50126 | EN 50126 Railway applications - The specification and demonstration of Reliability, Availability, Maintainability and Safety (RAMS) | |
| EN50716 | EN 50716 Railway applications - Requirements for software development | |
| EN50128 | EN 50128 Railway applications - Communication, signalling and processing systems - Software for railway control and protection systems | replaced by EN50716 |
| EN50657 | EN 50657 Railway applications - Rolling stock applications - Software on board rolling stock | replaced by EN50716 |
| EN50129 | EN 50129 Railway applications - Communication, signalling and processing systems - Safety related electronic systems for signalling | |
| CLC50701 | CLC/TS 50701 Railway applications - Cybersecurity | |
| ASPICE | Automotive SPICE Process Reference and Assessment Model | |
| ARP4754B | ARP4754B Guidelines for Development of Civil Aircraft and Systems | |
| ARP4761A | ARP4761A Guidelines for Conducting the Safety Assessment Process on Civil Aircraft, Systems, and Equipment | |
| DO178C | DO-178C Software Considerations in Airborne Systems and Equipment Certification | |
| DO254 | DO-254 Design Assurance Guidance for Airborne Electronic Hardware | |
| DO326A | DO-326A Airworthiness Security Process Specification | |
| DO356A | DO-356A Airworthiness Security Methods and Considerations | |
| CMMI | CMMI Capability Maturity Model Integration | |

## Per domain

**Table domains** - the standards of each domain

| Domain | Process | Safety | Safety analysis | Cybersecurity | Security analysis |
|---|---|---|---|---|---|
| automotive | ASPICE, ISO330xx | ISO26262 | HARA - hazard analysis and risk assessment, ISO26262-3 | ISO21434 | TARA - threat analysis and risk assessment, ISO21434 |
| airborne | ARP4754B (system), DO178C (software), DO254 (hardware) | ARP4761A | FHA - functional hazard assessment, then PSSA and SSA, ARP4761A | DO326A | security risk assessment, DO326A and DO356A |
| railway | EN50126-1 (RAMS), EN50716 (software), EN50129 | EN50126-1, EN50126-2 | hazard log, with hazard analysis and risk assessment, EN50126-1 | CLC50701 | threat and risk assessment, CLC50701 |
| medical | ISO13485 (quality management), IEC62304 (software), IEC60601-1 (electrical equipment) | ISO14971 | risk management file, ISO14971 | IEC81001-5-1 | threat model, IEC81001-5-1 |
| industrial | IEC61508 | IEC61508 | H&RA - hazard and risk analysis, IEC61508-1, clause 7.4 | IEC62443 | threat and risk assessment, IEC62443-3-2 |
| software | ISO12207, ISO29119 (test) | ISO15026 (assurance) | hazard and risk analysis, by the method of IEC61508-1, clause 7.4 | IEC62443-4-1 | threat and risk assessment, IEC62443-3-2 |
