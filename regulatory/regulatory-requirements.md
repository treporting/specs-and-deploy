# Regulatory requirements

What standards and laws ask of treporting. They stand beside the user needs (r1): they do not come
from a role, and they hold whether a user asks for them or not. The system requirements (r2) are
derived from both, and every regulatory requirement that treporting realizes has at least one system requirement; one it does not realize says why.

The requirements themselves - ids `REG_<module>_nn` - are in this directory, one file per module:
a standard or its part, a law, a guidance, or the licences of the components. A file is named by
the short form of [standards.md](../standards.md), in the format of [formats.md](../formats.md),
section Regulatory requirements. Table modules names each file and its topic. Which system requirements
realize one is read from their uplinks; one that treporting does not realize says why in
`not_realized`.

Each line says in own words what the clause asks; the clause itself is authoritative. Which of them
apply to a project is decided by the organization that uses treporting, by its domain.

## Modules

**Table modules** - each file of regulatory requirements, and its topic

| File | Topic |
|---|---|
| [ISO26262-8.yaml](ISO26262-8.yaml) | Tool qualification |
| [IEC61508-3.yaml](IEC61508-3.yaml) | Tool qualification |
| [DO178C.yaml](DO178C.yaml) | Tool qualification, Safety classes |
| [DO330.yaml](DO330.yaml) | Tool qualification |
| [ISO13485.yaml](ISO13485.yaml) | Tool qualification |
| [CFR21-11.yaml](CFR21-11.yaml) | Electronic records |
| [FDADI.yaml](FDADI.yaml) | Electronic records |
| [EU2016-679.yaml](EU2016-679.yaml) | Personal data and legal notice |
| [BetrVG.yaml](BetrVG.yaml) | Personal data and legal notice |
| [DDG.yaml](DDG.yaml) | Personal data and legal notice |
| [ISO26262-9.yaml](ISO26262-9.yaml) | Safety classes |
| [ARP4754B.yaml](ARP4754B.yaml) | Safety classes |
| [IEC61508-1.yaml](IEC61508-1.yaml) | Safety classes |
| [IEC61508-2.yaml](IEC61508-2.yaml) | Safety classes |
| [IEC61508-4.yaml](IEC61508-4.yaml) | Safety classes |
| [IEC62304.yaml](IEC62304.yaml) | Safety classes |
| [EU2024-2847.yaml](EU2024-2847.yaml) | Products on the market |
| [Licences.yaml](Licences.yaml) | Products on the market |
| [ISO29148.yaml](ISO29148.yaml) | Requirements engineering |
| [ISO25010.yaml](ISO25010.yaml) | Product quality |

## Tool qualification

treporting controls nothing, but it can fail to show a gap - a tool that people rely on when they
release.

## Products on the market

treporting is open source, with the possibility of licensed features, and offered commercially. Whether and from
when the Cyber Resilience Act applies to it is to be confirmed by legal advice.

## Requirements engineering

The characteristics every requirement of treporting is reviewed against. treporting does not realize
them; our work keeps them, in the reviews against the criteria of each level (PROC_13). A criterion
names the one it comes from in its field `source`.

## Product quality

The characteristics of the product quality model of ISO25010. treporting does not realize them as
requirements of its own; our work keeps them in its decisions. Where a means was applied to gain
more information on them, the decision states what it found (PROC_18).

## Open

- `treporting = 2_selfhosted` - **Products on the market**: not known whether, and from when, the
  Cyber Resilience Act applies to treporting. Legal advice is to confirm it.
