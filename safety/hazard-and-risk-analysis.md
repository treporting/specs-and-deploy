# Hazard and risk analysis

The hazards treporting can cause, by the method of IEC 61508-1, 7.4. treporting controls nothing; it is a
tool that people rely on when they decide on a release. Its hazards are therefore wrong
information that leads to a wrong decision - a product released with a gap that treporting did
not show.

The hazards are on level h1, with the ids `HAZ_nn`; the threats of security share that level
with the ids `THR_nn`. Each hazard is controlled by a measure on level h2, with the ids
`HAZC_nn` - the prefix of the hazard and C for control -, and each measure asks for the user needs (r1) that realize it - mostly those of the
section Safety of the product (PSF) in `requirements/user-needs.md`. The tool
qualification package builds its analysis of tool errors on this file.

## Hazards

`HAZ_01` to `HAZ_07`: in [hazards.yaml](hazards.yaml).

## Measures

`HAZC_01` to `HAZC_07`: in [hazards.yaml](hazards.yaml).

The detection itself is validated with faults built into data of known content (SYS_27).

## Risk

The severity of each hazard is that of the product a user releases on its information, which
treporting does not know. The risk is therefore kept low by detection rather than rated: every
hazard is either prevented by the requirements above, or made visible to the user, and the
detection itself is validated with faults built into data of known content (SYS_27).
