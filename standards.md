# Standards

The standards of the domains treporting is used in: the process a project follows, the safety
standard, and the analysis that finds what is to be controlled. A company process has to match
the standardized one (see Process in [definitions.md](definitions.md)).

## Per domain

**Table domains** - the standards of each domain

| Domain | Process | Safety | Safety analysis | Cybersecurity | Security analysis |
|---|---|---|---|---|---|
| automotive | Automotive SPICE, ISO/IEC 330xx (formerly ISO/IEC 15504) | ISO 26262 | HARA - hazard analysis and risk assessment, ISO 26262-3 | ISO/SAE 21434 | TARA - threat analysis and risk assessment, ISO/SAE 21434 |
| airborne | ARP4754B (system), DO-178C (software), DO-254 (hardware) | ARP4761A | FHA - functional hazard assessment, then PSSA and SSA, ARP4761A | DO-326A | security risk assessment, DO-326A and DO-356A |
| railway | EN 50126-1 (RAMS), EN 50716 (software, replaces EN 50128 and EN 50657), EN 50129 | EN 50126-1, EN 50126-2 | hazard log, with hazard analysis and risk assessment, EN 50126-1 | CLC/TS 50701 | threat and risk assessment, CLC/TS 50701 |
| medical | ISO 13485 (quality management), IEC 62304 (software), IEC 60601-1 (electrical equipment) | ISO 14971 | risk management file, ISO 14971 | IEC 81001-5-1 | threat model, IEC 81001-5-1 |
| industrial | IEC 61508 | IEC 61508 | H&RA - hazard and risk analysis, IEC 61508-1 7.4 | IEC 62443 | threat and risk assessment, IEC 62443-3-2 |
| software | ISO/IEC/IEEE 12207, ISO/IEC/IEEE 29119 (test) | ISO/IEC/IEEE 15026 (assurance) | hazard and risk analysis, by the method of IEC 61508-1 7.4 | IEC 62443-4-1 | threat and risk assessment, IEC 62443-3-2 |

## Across domains

**Table kinds** - the standards that hold in every domain

| Kind | Standards |
|---|---|
| process assessment and improvement | ISO/IEC 330xx (Automotive SPICE), CMMI |
| quality management | ISO 9001 |
| information security | ISO/IEC 27001 |
| software life cycle | ISO/IEC/IEEE 12207 |

## Short forms

Every other document names a standard by its short form: the first of its publishers, a hyphen,
and its number with its part. A standard whose name has no publisher in front keeps its name -
DO-178C, ARP4754B, Automotive SPICE, CMMI.

**Table short forms** - the standards with more than one publisher, in full and short

| Standard | Short form |
|---|---|
| ISO/IEC/IEEE 12207 - software life cycle processes | ISO-12207 |
| ISO/IEC/IEEE 15026 - assurance | ISO-15026 |
| ISO/IEC/IEEE 15289 - content of life-cycle information items | ISO-15289 |
| ISO/IEC/IEEE 29119 - software testing | ISO-29119 |
| ISO/IEC/IEEE 29148 - requirements engineering | ISO-29148 |
| ISO/IEC 330xx - process assessment | ISO-330xx |
| ISO/IEC 15504 - process assessment, replaced by ISO/IEC 330xx | ISO-15504 |
| ISO/IEC 27001 - information security | ISO-27001 |
| ISO/SAE 21434 - road vehicles, cybersecurity | ISO-21434 |
| CLC/TS 50701 - railway, cybersecurity | CLC-50701 |

A standard of one publisher only takes the hyphen: IEC 61508-1 is IEC-61508-1, EN 50716 is
EN-50716, ISO 26262-3 is ISO-26262-3.
