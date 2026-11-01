# devops

The environment of our work: the tools we build treporting with, and how they are ruled. It is
written for whoever sets up or changes that environment, and for the auditor who checks it.

## Tools are code

**Decision**: PROC_05 holds for the tools of our work as for the code of the product. Editor
settings, scripts and the CI are code. Every tool realizes an approved requirement. There is no
part where something is built without one.

**Why**: an auditor asks how the tools that build the product are controlled. With this rule the
answer is the same as for the product: by requirements, reviewed and approved before the change.

**Rejected**: to narrow PROC_05 to the code of the product, and rule the tools by PROC rules alone.
It is lighter, but it leaves a part without requirements.

**Consequence**: the requirements on our tools are kept in `requirements/tool-requirements.yaml`,
with the ids `TOOL_nn`. A tool requirement names as uplink the PROC rule or the regulatory
requirement it serves. It is reviewed and approved like every other requirement (PROC_13). Where the tools live,
in which language, and their chain from requirement to code: [tools/README.md](tools/README.md).

## Environments

The versions as installed on 2026-10-25. A tool is listed when the work on treporting needs it.

**Table environments** - the tools of our work, and their versions

| Tool | Version | What for |
|---|---|---|
| macOS | 26.4.1 | the development machine |
| VS Code | 1.142.0 | the editor, with extensions - see below |
| Claude Code | 2.1.310 | the AI assistant, as an extension of VS Code; its settings are kept one level above the repos (see harness.md) |
| git | 2.50.1 | version control |
| GitHub CLI (`gh`) | 2.101.0 | tickets, pull requests, the project |
| Go | 1.25.14 | the subsystems in Go; the CI builds with 1.25 as well |
| Node.js, npm | 24.21.0, 11.19.0 | the viewer and its tests |
| Docker Desktop | 4.17.0 - engine 20.10.23, Compose 2.15.1 | the containers, gel |
| Google Chrome | 155.0.8059.86 | renders the export of the documents to PDF |
| GitHub | - | the repos of the organization treporting, its project, protected `main`, pull requests |
| GitHub Actions | - | the CI: the gate of the tools, see Continuous integration |

How to read a version: `sw_vers`, `git --version`, `gh --version`, `go version`, `node --version`,
`npm --version`, `docker --version`, `docker compose version`; VS Code and its extensions in
*Help > About* and the view *Extensions*.

### Extensions of VS Code

**Table extensions** - the extensions of VS Code, and their versions

| Extension | Version | What for |
|---|---|---|
| Claude Code (Anthropic) | 2.1.310 | the AI assistant |
| Markdown PDF (yzane) | 2.2.0 | the export of the documents to PDF |
| Go (golang) | 0.58.0 | Go: language server, tests, debugging |
| Container Tools (Microsoft) | 2.5.3 | containers and compose files |
| Dev Containers (Microsoft) | 0.469.0 | development inside a container |
| GitHub Actions (GitHub) | 0.32.3 | the workflows of the CI |

### Dependencies of the tools

**Table dependencies** - the modules the tools in `tools/` are built with

| Module | Version | Licence | What for |
|---|---|---|---|
| `gopkg.in/yaml.v3` | 3.0.1 | MIT, Apache-2.0 | reading YAML |
| `github.com/santhosh-tekuri/jsonschema/v6` | 6.0.3 | Apache-2.0 | checking a YAML against its JSON Schema (TOOL_05) |
| `golang.org/x/text` | 0.14.0 | BSD-3-Clause | the messages of jsonschema; of the Go team |

`github.com/dlclark/regexp2` stands in `go.sum` as part of the module graph of jsonschema; it is not
built into the tools (`go list -deps`).

## Documents to PDF

**Decision**: the documents are exported to PDF with the extension Markdown PDF (yzane) of VS Code.

**Why**: it is the most widely used extension for it. Header and footer are set in its settings, as
HTML templates with placeholders for title and page. PlantUML goes to a server whose address is
set, so a local server keeps a diagram on the machine.

**Rejected**:

- Markdown Preview Enhanced - it draws PlantUML without a server, from `plantuml.jar`, but needs
  Java, and its header and footer in the PDF could not be confirmed
- Pretty Markdown PDF - a fork of Markdown PDF that writes the date in the form of ISO8601, used by
  fewer than the original

How it works:

- It renders with the Chrome of the machine.
- Mermaid diagrams are drawn; the script of Mermaid is loaded from a CDN, the diagram stays on the
  machine.
- PlantUML diagrams are drawn by a local PlantUML server only, never the public one: a diagram must
  not leave the machine.
- An exported PDF is never committed.

## Merging

**Decision**: every repo of treporting merges a pull request by a merge commit only; squash and
rebase are switched off in the settings of the repo. GitHub deletes the branch after the merge;
its commits stay in `main`. A new repo gets the same settings when it is made:

```bash
gh api -X PATCH repos/treporting/<repo> \
  -F allow_merge_commit=true -F allow_squash_merge=false -F allow_rebase_merge=false \
  -F delete_branch_on_merge=true
```

**Why**: the commits of the pull request reach `main` as they are, with their author. The authors of
a requirement or a document are the author and the co-authors of the last commit, not a merge, that
changed it (tools/README.md), and the revision a review names stays a commit on `main`.

**Rejected**:

- squash - it makes the one who opened the pull request the author of the one commit on `main`, so
  the Product Owner would be the author of every requirement, and could review none
- rebase - it keeps the author, but writes every commit again with a new hash; a revision named
  before the merge is not on `main`

## Continuous integration

**Decision**: `.github/workflows/tools.yml` runs on every pull request to `main`: the job
`reviewcheck` (SR_09), the job `tests` (SR_10), the job `schemacheck` (SR_20), the job
`rulesetcheck` (SR_23), the job `annotationcheck` (SR_26, SR_31) and the job `uplinkcheck`
(SR_29). Each is a required check in the ruleset `main is protected`, which has no bypass - a pull request whose
job failed cannot be merged, by anyone. The job `linkcheck` (SR_34) runs beside them. It is no
required check. It reports a suspect link and does not fail: a suspect link blocks the baseline, not
the pull request (PROC_20). The job `build` calls `.github/workflows/build.yml`, the build steps of
every repo (SR_51 to SR_58), on the tools: its jobs `build / build`, `build / lint` and
`build / test` are required checks; `build / report` reports and is not. The ruleset also allows the merge commit only, as the settings of the repo do.

The ruleset is kept in the repo, in `.github/rulesets/main.json` - beside the workflows, as its form
is GitHub's - and the job `rulesetcheck` checks
on every pull request that GitHub has the same (TOOL_13). The file is the source: to change the
ruleset, change the file by pull request and review. As `rulesetcheck` is required, that pull
request cannot be merged until GitHub has the change. So once it is green but for `rulesetcheck`,
the file is sent whole; `rulesetcheck` then turns green, and the pull request is merged. Where it is
not merged, the file of `main` is sent back:

```bash
gh api repos/treporting/specs-and-deploy/rulesets --jq '.[] | {id, name}'
gh api -X PUT repos/treporting/specs-and-deploy/rulesets/<id> --input .github/rulesets/main.json
```

Until a change is sent or sent back, `rulesetcheck` fails on every pull request - so it is not
forgotten. For the minutes between the sending and the merge, GitHub holds the ruleset of the pull
request, not that of `main`. A required
check is named by its job, with `integration_id` 15368 - GitHub Actions. GitHub knows the name of a
check only after it has run once, so a new job runs green before it becomes required.

**Why**: a run by hand leaves no record, and depends on the one who runs it; a run of the CI is
recorded with the pull request, by a computer (REG_CFR21-11_01).

**Rejected**:

- `reviewcheck` only where `reviews/` or `tools/` change - a required check that a filter of paths
  skips stays pending and blocks the merge; to save a minute, the workflow would need logic of its
  own
- a bypass for the admin - a check that someone can step over is a check by hand again
- `linkcheck` as a required check - it never fails, so it would block no merge; it would only make
  each merge wait until it has run

Each required check without a test is verified by a demonstration, recorded in `verifications/` as
formats.md, section Verifications, sets.

## Baseline

The baseline of a repo is prepared by `.github/workflows/baseline.yml` of specs-and-deploy
(TOOL_27, SR_61 to SR_65), started by the label `baseline` on a ticket of the repo. The run itself is
in tools/README.md, section baseline. Where the run passes, it opens a pull request with sbom.md. A
person merges it and sets the signed tag with the command in its description (PROC_06). Where a pull
request changes sbom.md, the job `build / build` checks that it lists the packages of its head
(TOOL_28).

It writes with a token of the GitHub App `treporting-baseline`, not with `GITHUB_TOKEN`: a pull
request opened with `GITHUB_TOKEN` starts no workflow, so its required checks would never run.

**Set up the App, once for the org** - by an owner of the org:

1. Org settings, Developer settings, GitHub Apps, New GitHub App:
   - name `treporting-baseline`; homepage `https://github.com/treporting/specs-and-deploy`
   - Webhook: not active
   - Repository permissions: Contents, Issues and Pull requests read and write; Metadata read
   - Where can this GitHub App be installed: only on this account
2. In the App: note its App ID; Generate a private key - a `.pem` file is downloaded.
3. Install App, on the org, for the repos that make a baseline.
4. Hand the ID and the key to the workflows, and delete the file:

```bash
gh variable set BASELINE_APP_ID --org treporting --visibility all --body <App ID>
gh secret set BASELINE_APP_PRIVATE_KEY --org treporting --visibility all < <file>.pem
rm <file>.pem
```

**Set up a repo for its baseline**:

1. Copy `.github/workflows/ticket-baseline.yml` and `.github/ISSUE_TEMPLATE/baseline.yml` of
   specs-and-deploy; in the caller, set `go-version-file` and name the images in `images`.
2. Create the label:

```bash
gh label create baseline --repo treporting/<repo> --color 0E8A16 \
  --description "starts the baseline of this repo"
```

3. Add the repo to the installation of the App (step 3 above).

**Set up the key that signs a tag, once per person** - git signs with an SSH key:

```bash
git config --global gpg.format ssh
git config --global user.signingkey ~/.ssh/<key>.pub
```

Add the same public key on GitHub a second time, as signing key: Settings, SSH and GPG keys, New SSH
key, Key type Signing Key. GitHub then shows the tag verified. To check a tag: `git cat-file -p
<label>` ends with `-----END SSH SIGNATURE-----`; on GitHub, Code, Tags, the tag shows Verified.

## CI checks

The checks of a pull request, each with a tool requirement of its own. Which of them the CI runs
today is set in the section Continuous integration.

**Table CI checks** - each check, and the tool requirement it realizes

| Check | Tool requirement |
|---|---|
| every YAML file matches its schema | TOOL_05 |
| every id in `@implements`, `@verifies` and `@refines` exists, and is approved, in the base commit of the pull request - a requirement added in the same pull request does not count (PROC_05) | TOOL_06 |
| the annotations counted a second way, apart from the parser | TOOL_07 |
| the uplink of every SR: the SYS it names exists in specs-and-deploy on `main` and is allocated to the subsystem of the SR (r4) | TOOL_08 |
| an approved requirement names approved uplinks only | TOOL_09 |
| every test file belongs to exactly one test level | TOOL_10 |
| every review result keeps the rules of the section Reviews of formats.md, and the tests of the tools pass | TOOL_12 |
| the ruleset of `main` on GitHub has the rules of `.github/rulesets/main.json` | TOOL_13 |
| every `@implements`, `@verifies` and `@refines` names the version of its id, `ID@x.y.z` | TOOL_20 |

## Open

`<empty>`

## Change history

| Version | Date | Change | Reason | By |
|---|---|---|---|---|
| 1.1.1 | 2026-10-28 | - | #272 | Claude Opus 5.5 and Mathias Fischer, authors; Mathias Fischer, reviewer |
| 2.1.1 | 2026-10-28 | Continuous integration | #279 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 3.1.1 | 2026-10-28 | Continuous integration | #280 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 4.1.1 | 2026-10-28 | Continuous integration | #40 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 5.1.1 | 2026-10-28 | Tools are code; Dependencies of the tools; Continuous integration | #291 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 6.1.1 | 2026-10-29 | Environments; Documents to PDF; Build and deployment; Open | #322 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 7.1.1 | 2026-10-30 | Continuous integration; Open | #322 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 8.1.1 | 2026-10-30 | Change history | #323 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 9.1.1 | 2026-10-30 | Continuous integration; CI checks | #320 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 10.1.1 | 2026-10-31 | Continuous integration | #342 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 11.1.1 | 2026-11-01 | Baseline; Open | #343 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
| 12.1.1 | 2026-11-01 | Baseline | #386 | Claude Opus 5.5, author; Mathias Fischer, reviewer |
