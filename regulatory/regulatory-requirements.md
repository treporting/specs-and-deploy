# Regulatory requirements

What standards and laws ask of treporting. They stand beside the user needs (r1): they do not come
from a role, and they hold whether a user asks for them or not. The system requirements (r2) are
derived from both, and every regulatory requirement that treporting realizes has at least one system requirement; one it does not realize says why.

The requirements themselves - ids `REG_<module>_nn` - are in this directory, one file per module:
a standard or its part, a law, a guidance, or the licences of the components. A file is named by
the short form of [standards.md](../standards.md), in the format of [formats.md](../formats.md),
section Regulatory requirements. The sections below group them by topic. Which system requirements
realize one is read from their uplinks; one that treporting does not realize says why in
`not_realized`.

Each line says in own words what the clause asks; the clause itself is authoritative. Which of them
apply to a project is decided by the organization that uses treporting, by its domain.

## Tool qualification

treporting controls nothing, but it can fail to show a gap - a tool that people rely on when they
release.

`REG_ISO26262-8_01`, `REG_IEC61508-3_01`, `REG_DO178C_01`, `REG_DO330_01`, `REG_ISO13485_01`: in
[ISO26262-8.yaml](ISO26262-8.yaml), [IEC61508-3.yaml](IEC61508-3.yaml), [DO178C.yaml](DO178C.yaml), [DO330.yaml](DO330.yaml), [ISO13485.yaml](ISO13485.yaml).

## Electronic records

`REG_CFR21-11_01` to `REG_CFR21-11_05`, `REG_FDADI_01`: in [CFR21-11.yaml](CFR21-11.yaml), [FDADI.yaml](FDADI.yaml).

## Personal data and legal notice

`REG_EU2016-679_01`, `REG_EU2016-679_02`, `REG_BetrVG_01`, `REG_DDG_01`: in
[EU2016-679.yaml](EU2016-679.yaml), [BetrVG.yaml](BetrVG.yaml), [DDG.yaml](DDG.yaml).

## Safety classes

`REG_ISO26262-9_01`, `REG_ARP4754B_01`, `REG_DO178C_02`, `REG_IEC61508-1_01`, `REG_IEC61508-2_01`,
`REG_IEC61508-4_01`, `REG_IEC62304_01`: in [ISO26262-9.yaml](ISO26262-9.yaml), [ARP4754B.yaml](ARP4754B.yaml), [DO178C.yaml](DO178C.yaml), [IEC61508-1.yaml](IEC61508-1.yaml),
[IEC61508-2.yaml](IEC61508-2.yaml), [IEC61508-4.yaml](IEC61508-4.yaml), [IEC62304.yaml](IEC62304.yaml).

## Products on the market

treporting is open source, with the possibility of licensed features, and offered commercially. Whether and from
when the Cyber Resilience Act applies to it is to be confirmed by legal advice.

`REG_EU2024-2847_01` to `REG_EU2024-2847_04`, `REG_Licences_01`: in [EU2024-2847.yaml](EU2024-2847.yaml), [Licences.yaml](Licences.yaml).
