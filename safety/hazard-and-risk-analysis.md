# Hazard and risk analysis

The hazards treporting can cause, by the method of IEC 61508-1, 7.4. treporting controls nothing; it is a
tool that people rely on when they decide on a release. Its hazards are therefore wrong
information that leads to a wrong decision - a product released with a gap that treporting did
not show.

The hazards are on level h1, with the ids `HAZ_nn`; the threats of security share that level
with the ids `THR_nn`. Each hazard is controlled by a measure on level h2, with the ids
`HAZC_nn` - the prefix of the hazard and C for control -, and each measure asks for the user needs (r1) that realize it - mostly those of the
section Safety of the product (PSF) in `requirements/user-needs.md`. The hazards and measures are
in [hazards.yaml](hazards.yaml), in the format of [formats.md](../formats.md). If treporting gets a
tool qualification, its analysis of tool errors builds on this file. It is written for whoever
derives the user needs from the measures, and for an assessor.

## Risk

The severity of each hazard is that of the product a user releases on its information, which
treporting does not know. The risk is therefore kept low by detection rather than rated: every
hazard is either prevented by its measure, or made visible to the user, and the
detection itself is validated with faults built into data of known content (SYS_27).

## Open

- `treporting = 3_tools` - **Tool qualification**: the tool qualification package is not made yet.
  Its analysis of tool errors builds on this file and shows whether this analysis is complete.

## Change history

| Version | Date | Change | Reason | By |
|---|---|---|---|---|
| 1.1.1 | 2026-10-30 | - | #329 | Claude Opus 5.5 and Mathias Fischer, authors; Mathias Fischer, reviewer |
