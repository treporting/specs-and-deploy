# Standards

The standards of the domains treporting is used in: the process a project follows, the safety
standard, and the analysis that finds what is to be controlled. A company process has to match
the standardized one (see Process in [definitions.md](definitions.md)).

## Short forms

Every document names a standard by its short form: the first of its publishers, a hyphen, and its
number with its part. A standard whose name has no publisher in front keeps its name, such as
DO-178C. Only the table below names the standards in full; a part follows the rule, so IEC 61508-1
is IEC-61508-1.

**Table short forms** - every standard named in the documents, in full and short

| Standard | Short form | What |
|---|---|---|
| ISO/IEC/IEEE 12207 | ISO-12207 | software life cycle processes |
| ISO/IEC/IEEE 15026 | ISO-15026 | systems and software assurance |
| ISO/IEC/IEEE 15289 | ISO-15289 | content of life-cycle information items |
| ISO/IEC/IEEE 29119 | ISO-29119 | software testing |
| ISO/IEC/IEEE 29148 | ISO-29148 | requirements engineering |
| ISO/IEC 330xx | ISO-330xx | process assessment |
| ISO/IEC 15504 | ISO-15504 | process assessment, replaced by ISO-330xx |
| ISO/IEC 27001 | ISO-27001 | information security management |
| ISO 9001 | ISO-9001 | quality management |
| ISO 8601 | ISO-8601 | dates and times |
| ISO 26262 | ISO-26262 | road vehicles, functional safety |
| ISO/SAE 21434 | ISO-21434 | road vehicles, cybersecurity |
| ISO 13485 | ISO-13485 | medical devices, quality management |
| ISO 14971 | ISO-14971 | medical devices, risk management |
| IEC 62304 | IEC-62304 | medical device software, life cycle |
| IEC 60601-1 | IEC-60601-1 | medical electrical equipment |
| IEC 81001-5-1 | IEC-81001-5-1 | health software, security |
| IEC 61508 | IEC-61508 | functional safety of electrical, electronic and programmable systems |
| IEC 62443 | IEC-62443 | industrial automation, security |
| EN 50126 | EN-50126 | railway, RAMS |
| EN 50716 | EN-50716 | railway, software |
| EN 50128 | EN-50128 | railway, software, replaced by EN-50716 |
| EN 50657 | EN-50657 | railway rolling stock, software, replaced by EN-50716 |
| EN 50129 | EN-50129 | railway, safety related electronic systems for signalling |
| CLC/TS 50701 | CLC-50701 | railway, cybersecurity |
| Automotive SPICE | Automotive SPICE | process assessment, automotive |
| ARP4754B | ARP4754B | civil aircraft and systems, development |
| ARP4761A | ARP4761A | civil aircraft and systems, safety assessment |
| DO-178C | DO-178C | airborne systems, software |
| DO-254 | DO-254 | airborne electronic hardware |
| DO-326A | DO-326A | airworthiness security, process |
| DO-356A | DO-356A | airworthiness security, methods |
| CMMI | CMMI | process improvement |

## Per domain

**Table domains** - the standards of each domain

| Domain | Process | Safety | Safety analysis | Cybersecurity | Security analysis |
|---|---|---|---|---|---|
| automotive | Automotive SPICE, ISO-330xx (formerly ISO-15504) | ISO-26262 | HARA - hazard analysis and risk assessment, ISO-26262-3 | ISO-21434 | TARA - threat analysis and risk assessment, ISO-21434 |
| airborne | ARP4754B (system), DO-178C (software), DO-254 (hardware) | ARP4761A | FHA - functional hazard assessment, then PSSA and SSA, ARP4761A | DO-326A | security risk assessment, DO-326A and DO-356A |
| railway | EN-50126-1 (RAMS), EN-50716 (software, replaces EN-50128 and EN-50657), EN-50129 | EN-50126-1, EN-50126-2 | hazard log, with hazard analysis and risk assessment, EN-50126-1 | CLC-50701 | threat and risk assessment, CLC-50701 |
| medical | ISO-13485 (quality management), IEC-62304 (software), IEC-60601-1 (electrical equipment) | ISO-14971 | risk management file, ISO-14971 | IEC-81001-5-1 | threat model, IEC-81001-5-1 |
| industrial | IEC-61508 | IEC-61508 | H&RA - hazard and risk analysis, IEC-61508-1, clause 7.4 | IEC-62443 | threat and risk assessment, IEC-62443-3-2 |
| software | ISO-12207, ISO-29119 (test) | ISO-15026 (assurance) | hazard and risk analysis, by the method of IEC-61508-1, clause 7.4 | IEC-62443-4-1 | threat and risk assessment, IEC-62443-3-2 |

## Across domains

**Table kinds** - the standards that hold in every domain

| Kind | Standards |
|---|---|
| process assessment and improvement | ISO-330xx (Automotive SPICE), CMMI |
| quality management | ISO-9001 |
| information security | ISO-27001 |
| software life cycle | ISO-12207 |
