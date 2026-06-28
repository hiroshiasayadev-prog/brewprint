# DRMCP-TASK-MCP-006-04: Define source-location and exceptional path-exposure contract

- **id**: DRMCP-TASK-MCP-006-04
- **status**: done
- **date**: 2026-06-28
- **work_item**: DRMCP-WORK-MCP-006
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 2d
- **depends_on**:
  - DRMCP-TASK-MCP-006-03
- **outputs**:
  - spec:drmcp.design_records_mcp.schema.diagnostics
  - spec:drmcp.design_records_mcp.tools.validate_records
  - spec:drmcp.design_records_mcp.tools.list_records
  - spec:drmcp.design_records_mcp.tools.get_records
  - spec:drmcp.design_records_mcp.tools.resolve_reference
  - spec:drmcp.design_records_mcp.schema.authoring_transaction_schema
  - spec:drmcp.design_records_mcp.tools.authoring_transaction_model
  - spec:drmcp.design_records_mcp.tools.propose_record_create
  - spec:drmcp.design_records_mcp.tools.accept_proposed_write
  - DRMCP-WORK-MCP-006

## Goal

Define the concrete source-location association and narrow exceptional physical-path exposure contract.

Enable writer and repair agents to identify repairable files without adding paths to normal successful list, retrieval, or resolver projections.

## Work

- Consume the accepted T01 authority baseline, T02 validation-execution contract, and T03 diagnostic representation.
- Confirm the exact `DRMCP-TASK-MCP-006-*` inventory before creating this Task.
- Inventory current-root, legacy-root, source-provenance, authoring-target, diff, and files-written path fields.
- Separate diagnostic location, explicit patch output, files-written confirmation, and privileged debug or emergency inspection purposes.
- Define one concrete `location` object for current and legacy source-backed findings.
- Define portable path representation, root identification, separator normalization, and Windows path handling.
- Define stable location identity and the T03-required deterministic location sort key.
- Define operation-specific location and physical-path exposure without changing W004 or W005 normal successful projections.
- Define handling when a source-backed diagnostic cannot construct the required repairable location.
- Present and accept one design decision at a time.
- Record each accepted decision immediately in this Task Evidence.
- Keep normative specifications unchanged until every T04 decision is accepted.
- Reflect only the accepted contract into directly affected normative specifications.
- Leave final overview, responsibility, and closure synchronization to T05 unless a direct T04 contract contradiction requires an earlier edit.
- Run scoped validation and changed-file whitespace checks after normative reflection.
- Run independent review before changing this Task to `done`.

## Done condition

- D01 through D07 are accepted and recorded separately.
- The `location` object identifies every repairable current or legacy source required by T03.
- Current and legacy locations use a deterministic and machine-readable representation.
- Repository-relative, root-relative, and absolute path behavior is explicit for every affected surface.
- Separator normalization and Windows drive, UNC, and case handling are explicit.
- Stable location identity and sort-key behavior satisfy T03 deterministic ordering and duplicate suppression.
- `validate_records`, proposal-local validation, authoring diagnostics, conflicts, read-operation warnings, patch output, `files_written`, and debug or emergency surfaces have explicit exposure rules.
- Missing required source location cannot silently degrade to an opaque source token.
- Normal successful `list_records`, `get_records`, and `resolve_reference` projections remain path-free.
- W004 warning triggers, ordering, deduplication, partial success, wrappers, and successful projections remain unchanged.
- W005 resolver order, statuses, fallback eligibility, and successful target projection remain unchanged.
- T02 request, execution, wrapper, and validation-subject boundaries remain unchanged.
- T03 category, severity, ordering, duplicate-suppression, and authoring-trigger boundaries remain unchanged.
- Changed normative specifications pass scoped strict validation.
- Tracked and untracked changed files pass the applicable whitespace checks.
- Independent review reports no blocking, major, or minor findings before this Task is marked `done`.

## Verification

- Compare path-exposure ownership with `DRMCP-ADR-MCP-001` and `DRMCP-REQ-MCP-001`.
- Compare current-root and legacy-root identity with `namespace-scanning.md`.
- Compare retained current source provenance with `schema/record-model.md`, `schema/record-source.md`, and `schema/discovery.md`.
- Compare diagnostic requirements with the accepted T03 D05, D08, D09, and D11 decisions.
- Compare normal list and exact-retrieval path hiding with final W004 contracts.
- Compare successful resolver target hiding with final W005 contracts.
- Compare authoring target, diff, proposal retrieval, and files-written path fields with the authoring transaction contracts.
- Confirm that no new debug or emergency operation is introduced without separate authority.
- Confirm that no normative specification changes before D01 through D07 are accepted.
- Run the strict spec validator against only changed normative specifications.
- Run `git diff --check` for tracked T04 files.
- Run `git diff --no-index --check` for this Task while it remains untracked.
- Run independent review before changing status to `done`.

## Evidence

### Exact Task inventory

The exact directory `drmcp/records/tasks/mcp/` was listed on 2026-06-28.

Existing `DRMCP-TASK-MCP-006-*` records were:

- `DRMCP-TASK-MCP-006-01-establish-validation-diagnostics-and-path-exposure-correction-baseline.md`;
- `DRMCP-TASK-MCP-006-02-define-current-repository-and-relation-validation-execution-contract.md`;
- `DRMCP-TASK-MCP-006-03-define-machine-readable-diagnostic-representation-and-semantic-invalidity-mapping.md`.

No `DRMCP-TASK-MCP-006-04` record existed.
The next Task is therefore `DRMCP-TASK-MCP-006-04`.

### Authority matrix

| concern | authority or accepted input | T04 treatment |
|---|---|---|
| Semantic and response ownership split | `DRMCP-ADR-MCP-001` | Preserve canonical identity responses and allow paths only on explicit diagnostic, patch, debug, or emergency surfaces. |
| Required path boundary | `DRMCP-REQ-MCP-001` | Define narrow operational exceptions without extending normal read projections. |
| Current source and root state | `DRMCP-WORK-MCP-003` | Consume configured current roots, app association, retained source provenance, containment, and conflict state. |
| Normal list and exact retrieval | `DRMCP-WORK-MCP-004` | Preserve warning triggers, ordering, partial success, successful projections, and normal path hiding. |
| Resolver and configured legacy lookup | `DRMCP-WORK-MCP-005` | Preserve resolver order, public statuses, legacy lookup states, and successful path-free targets. |
| Validation execution | `DRMCP-TASK-MCP-006-02` | Preserve selectors, subjects, lookup, failure boundaries, wrapper, and `ok` semantics. |
| Diagnostic representation | `DRMCP-TASK-MCP-006-03` | Consume the accepted `location` requirement, conflict aggregation, ordering, and duplicate-suppression handoff. |
| Concrete location and exceptional exposure | T04 and `schema/diagnostics.md` | Own object fields, path representation, root identification, normalization, stable identity, and surface policy. |
| Authoring transaction behavior | Authoring operation contracts | Reuse existing target, diff, and files-written outcomes; change only path representation or exposure when the accepted T04 contract directly requires it. |
| Final cross-spec synchronization | T05 | Leave navigation, broad responsibility summaries, final review, and W006 closure to T05. |

### T03 inherited location requirements

T04 preserves these accepted requirements:

- source-backed repository validation diagnostics expose a machine-readable repair location;
- identity-less, parse-failed, and unreadable current sources expose a location when their source is known;
- source-backed metadata, section, identity, and relation diagnostics expose a location;
- every current identity conflict member exposes its own location;
- every legacy lookup conflict candidate exposes its own location;
- proposal-local validation exposes a location when a repairable file target exists;
- persisted-file authoring diagnostics may expose a location under the T04 operation policy;
- malformed request items, source-less unresolved requested refs, unknown proposal IDs, expired proposal or body-cache IDs, and lifecycle-only conditions normally omit location;
- an opaque internal source token is not a sufficient repair target.

### Existing root and source information

| state | existing information |
|---|---|
| Repository | One configured `repository_root`. |
| Current root | Explicit `app_namespace` and repository-relative `records_root`; normalized root must equal `<app_namespace>/records`. |
| Legacy root | Repository-relative `records_root`; no app namespace, archive ID, or explicit stable root ID exists. |
| Current source | `app_namespace`, record `kind`, and repository-relative source `path`. |
| Current source boundary | Symlinked sources and directories are excluded; canonicalized candidates must remain within the configured source tree. |
| Legacy source boundary | Symlink, junction, reparse-point, and other alias sources are excluded; canonicalized candidates must remain within the configured legacy root and `repository_root`. |

A configured legacy-root list position is not currently defined as a stable identity.
T04 must not silently treat traversal order or configuration order as a portable source identity.

### Existing authoring path outputs

| surface | existing path output |
|---|---|
| Proposal `target` | Repository-relative `path`; transparency output only. |
| `diff.files[]` | Per-file `path` plus create or modify change kind. |
| Patch-mode `diff.text` | Unified diff headers containing repository paths. |
| `get_proposed_write` | Always returns patch-equivalent proposal detail with full `diff.text`. |
| `accept_proposed_write.files_written[]` | Per-file `path` for files actually written. |

These fields are authoring transaction outputs, not diagnostic-envelope fields.
T03 intentionally left them unchanged.
T04 may define their path representation and exposure but does not redesign proposal lifecycle, diff modes, write eligibility, or write behavior.

### Diagnostic location versus patch output

| surface | purpose |
|---|---|
| Diagnostic `location` | Identify the source file that contains or participates in a finding. It must support deterministic association, sorting, deduplication, and repair targeting. |
| Explicit patch output | Preview the concrete files and line changes a retained proposal would write. It may contain multiple file paths and unified diff headers. |
| `files_written` | Confirm which repository files were modified by an accepted proposal. |
| Debug or emergency inspection | Permit privileged physical inspection only when an explicit non-normal surface requires machine-local detail. |

A patch path is not automatically the diagnostic-location schema.
A diagnostic location is not automatically authorization for patch, normal read, or privileged absolute-path exposure.

### Debug and emergency surface state

The accepted ADR, Requirement, overview, MVP scope, responsibility boundary, and tool catalog permit narrow debug or emergency path exposure in principle.

No dedicated public debug or emergency operation, request, response, privilege model, or absolute-path field is currently defined.
T04 therefore defines an exposure boundary only.
It does not create a new debug tool or imply that normal tools have a hidden debug mode.

### Candidate changed-file manifest

Unconditional normative candidate:

- `drmcp/records/spec/design-records-mcp/schema/diagnostics.md`.

Conditional direct-contract candidates:

- `drmcp/records/spec/design-records-mcp/schema/record-model.md`;
- `drmcp/records/spec/design-records-mcp/namespace-scanning.md`;
- `drmcp/records/spec/design-records-mcp/tools/validate-records.md`;
- `drmcp/records/spec/design-records-mcp/tools/propose-record-create.md`;
- `drmcp/records/spec/design-records-mcp/tools/propose-record-update.md`;
- `drmcp/records/spec/design-records-mcp/tools/authoring-transaction-model.md`;
- `drmcp/records/spec/design-records-mcp/tools/get-proposed-write.md`;
- `drmcp/records/spec/design-records-mcp/tools/accept-proposed-write.md`;
- `drmcp/records/spec/design-records-mcp/schema/authoring-transaction-schema.md`.

A conditional file changes only when an accepted T04 decision requires a concrete field, representation, or direct operation rule.
Mere relevance does not justify an edit.

### Recheck-only manifest

The following files remain recheck-only unless an accepted T04 decision reveals a direct contradiction:

- `drmcp/records/spec/design-records-mcp/schema/discovery.md`;
- `drmcp/records/spec/design-records-mcp/schema/record-source.md`;
- `drmcp/records/spec/design-records-mcp/tools/list-records.md`;
- `drmcp/records/spec/design-records-mcp/tools/get-records.md`;
- `drmcp/records/spec/design-records-mcp/tools/resolve-reference.md`;
- `drmcp/records/spec/design-records-mcp/resolver.md`.

The following broader summaries remain T05 synchronization candidates:

- `drmcp/records/spec/design-records-mcp/overview.md`;
- `drmcp/records/spec/design-records-mcp/tools/overview.md`;
- `drmcp/records/spec/design-records-mcp/mvp-scope.md`;
- `drmcp/records/spec/design-records-mcp/responsibility-boundary.md`.

### T05 handoff

T05 retains:

- final cross-spec pointer and navigation synchronization;
- final responsibility-boundary and MVP summary synchronization;
- confirmation that W007 disposition scope remains untouched;
- final complete changed-file manifest;
- final scoped validation and whitespace evidence assessment;
- independent final review and finding correction;
- T04 and W006 closure synchronization.

### Decision register

| decision | question | status | accepted result |
|---|---|---|---|
| D01 | Which fields identify one current or legacy source location, and do both source families use one object shape? | accepted | Use one shared `location` object for current and legacy sources. Required identity fields are `source_scope`, `records_root`, and `path`. `source_scope` is `current` or `legacy`. Current locations additionally require `app_namespace`. Stable location identity is `(source_scope, records_root, path)`; `app_namespace` is contextual and does not add another identity component. |
| D02 | Which portable path forms, separator rules, and Windows path rules apply to diagnostic, patch, and written-file outputs? | accepted | Use normalized repository-relative paths for normal exposed path values. `records_root` and `path` use `/`, prohibit absolute, drive-qualified, UNC, URI, empty-segment, `.`, and `..` forms, preserve repository spelling without case folding, and remain canonically contained by `repository_root`; `path` must also remain under its declared `records_root`. Absolute physical paths remain D06-only. |
| D03 | Which location identity and stable sort key satisfy T03 ordering and duplicate suppression? | accepted | Reuse the D01 identity tuple `(source_scope, records_root, path)`. The stable sort key is `(source_scope_rank, records_root, path)` with `current = 0` and `legacy = 1`. Compare D02-normalized strings by locale-independent UTF-8 bytewise ascending order without case folding. Configuration order, scan order, discovery order, and absolute paths are never identity or sort inputs. |
| D04 | Which diagnostic and read-operation surfaces expose `location`, and which normal surfaces must omit it? | accepted | Expose `location` only for source-backed repair findings. `validate_records`, proposal-local validation, and applicable authoring diagnostics require direct location when one current source is the repair target. Conflict sources expose location only inside `members` or `candidates`. `list_records.missing_compact_field` requires the returned record source location. `get_records` and `resolve_reference` expose conflict-member or conflict-candidate locations, and expose direct location for a unique unreadable legacy source. Malformed, unsupported, duplicate, unresolved-without-source, disabled lookup, lifecycle-only, and all successful projections omit location. |
| D05 | Which path representation applies to explicit patch output, proposal targets, and `files_written`? | accepted | Keep existing scalar authoring path fields and apply the D02 normalized repository-relative grammar to `target.path`, `diff.files[].path`, unified diff path operands, and `files_written[].path`. Git `a/` and `b/` side prefixes and `/dev/null` are diff syntax, not repository path values. Matching proposal, diff-summary, patch, and write-confirmation entries use the same repository-relative path. |
| D06 | Which privileged debug or emergency surface may expose an absolute physical path, under what restrictions? | accepted | No current operation may expose an absolute physical path. A future separately contracted, host-enabled privileged debug or emergency operation may expose `physical_path` as a distinct field, never by replacing portable `location.path`. The operation must define request, response, privilege, scope, and failure behavior; existing operations gain no hidden debug flag. Absolute paths never affect identity, ordering, or duplicate suppression. |
| D07 | What happens when a source-backed diagnostic cannot construct the required repairable location? | accepted | Fail closed. A source-less condition may omit location, but a source-backed entry whose required direct, member, or candidate location cannot be constructed is not returned partially or with an opaque or absolute substitute. The operation fails before emitting a normal response or beginning a write. No affected item, conflict member, or candidate may be silently dropped. A post-write invariant failure must not misreport an actual write as `written: false`. |

D01 through D07 were accepted on 2026-06-28.
All T04 design decisions are accepted.
Normative reflection may now begin.

### D01 accepted shared location identity

Current and legacy source-backed findings use one shared `location` object.

Common required fields:

- `source_scope`;
- `records_root`;
- `path`.

`source_scope` values:

- `current` for a source under one configured current root;
- `legacy` for a source under one configured legacy root.

Current-only required context:

- `app_namespace`.

Stable location identity is:

```text
(source_scope, records_root, path)
```

D01 rules:

- `records_root` identifies the configured source root through its repository-relative configuration value.
- `path` identifies one source beneath that configured root; its exact path basis and separator normalization remain D02-owned.
- `app_namespace` is required for current locations because current-root configuration explicitly binds an app namespace to the root.
- `app_namespace` does not add another location identity component because one valid current root already has a one-to-one app association.
- Legacy locations do not invent an app namespace, archive namespace, configuration index, or traversal-order identity.
- Current and legacy conflict members use the same object shape and differ through `source_scope` and current-only `app_namespace`.
- A configured root that cannot be identified through this shared object is handled under D07 rather than represented by an opaque token.
- Absolute machine-local path exposure remains D06-owned.

### D02 accepted portable path representation

Normal exposed path values use normalized repository-relative form.

The same portable lexical grammar applies to diagnostic locations and to authoring path values whose concrete surface treatment remains D05-owned.

`records_root` and `path` rules:

- both are relative to the configured `repository_root`;
- `/` is the only exposed separator;
- a leading `/`, trailing `/`, duplicate separator, empty segment, `.` segment, and `..` segment are prohibited;
- Windows drive-qualified paths, UNC paths, device paths, and other absolute physical forms are prohibited;
- URI and `file://` forms are prohibited;
- repository path spelling is preserved; output does not lowercase, uppercase, Unicode-normalize, or otherwise case-fold path text;
- canonical filesystem resolution and containment checks remain internal execution behavior;
- a valid exposed `path` must canonically remain under `repository_root` and its declared `records_root`;
- a valid exposed `records_root` must canonically remain under `repository_root`;
- aliases, symlinks, junctions, reparse points, or other paths that escape the accepted source boundary cannot be represented as a valid normal location;
- machine-local absolute paths remain outside normal path values and are governed only by D06.

Example current location:

```json
{
  "source_scope": "current",
  "app_namespace": "drmcp",
  "records_root": "drmcp/records",
  "path": "drmcp/records/tasks/mcp/DRMCP-TASK-MCP-006-04-define-source-location-and-exceptional-path-exposure-contract.md"
}
```

Example legacy location:

```json
{
  "source_scope": "legacy",
  "records_root": "v01/records",
  "path": "v01/records/tasks/mcp/V01-TASK-MCP-003-01-example.md"
}
```

D02 defines the shared portable grammar.
D05 still decides how proposal targets, structured diff file entries, unified diff headers, and `files_written` apply that grammar.

### D03 accepted location identity and stable sort key

Location identity reuses the D01 tuple:

```text
(source_scope, records_root, path)
```

The stable sort key is:

```text
(source_scope_rank, records_root, path)
```

`source_scope_rank` is fixed as:

- `current = 0`;
- `legacy = 1`.

Comparison rules:

- `records_root` and `path` use their D02-normalized repository-relative strings;
- strings compare in locale-independent UTF-8 bytewise ascending order;
- no case folding, locale collation, natural-number sorting, or host-filesystem collation applies;
- `app_namespace` remains required current-location context but is not an additional location identity or location sort component;
- configured-root order, scan order, discovery order, conflict discovery order, and machine-local absolute paths are not identity or sort inputs.

Application rules:

- a `source` subject keeps the T03 order `app_namespace`, then `record_kind`, then this location sort key;
- `current_identity.members` sort by this location sort key;
- `legacy_lookup.candidates` sort by this location sort key;
- location-specific diagnostic duplicate suppression uses exact equality of the complete location identity tuple;
- T03 subject-type, category, field, target, conflict, and operation-specific ordering remain unchanged.

### D04 accepted operation-specific location exposure

`location` is exposed only when a normal diagnostic or warning identifies a known repository source that can be inspected or repaired.

Validation surfaces:

- `validate_records` diagnostics require direct `location` when one known current source is the repair target;
- proposal-local `validation.diagnostics` requires direct `location` when the candidate has a deterministic current repository target, including a create target that is not yet materialized;
- persisted-file authoring diagnostics may expose direct `location` when the current repository file is the cause or repair target;
- source-less configuration, request-shape, proposal-lifecycle, body-cache-lifecycle, and operation-lifecycle conditions omit location.

Conflict placement:

- `current_identity_conflict` exposes every conflicting source only through `conflict.members[].location`;
- read-operation current-conflict diagnostics use the same member-location placement;
- `legacy_lookup_conflict` exposes every legacy candidate only through `conflict.candidates[].location`;
- repository validation of a current relation with conflicting legacy targets additionally exposes the referring current source through the diagnostic's direct `location`;
- conflict member or candidate locations are not duplicated at the diagnostic top level.

Read-operation policy:

| operation state | location exposure |
|---|---|
| `list_records` `missing_compact_field` | Direct location of the returned current record source is required. |
| Successful `list_records`, `get_records`, or `resolve_reference` projection | Prohibited. |
| Malformed or unsupported requested ref | Prohibited. |
| Unresolved current or legacy ref with no source | Prohibited. |
| Current identity conflict | Each conflict member location is required; no direct diagnostic location. |
| Legacy lookup disabled | Prohibited. |
| Unique unreadable legacy source | Direct legacy source location is required on the warning or diagnostic. |
| Legacy lookup conflict | Each conflict candidate location is required; no direct diagnostic location. |
| Duplicate requested ref | Prohibited. |

Placement rules:

- direct `location` identifies the one source that contains or causes the finding;
- conflict-source locations appear only in their typed conflict collection;
- paths never move into `subject`, `field`, `value`, `target`, a successful record, or a successful resolver target;
- W004 warning triggers, request ordering, duplicate handling, partial success, wrappers, and successful projections remain unchanged;
- W005 resolver statuses, lookup order, fallback eligibility, wrappers, and successful target projection remain unchanged.

### D05 accepted authoring path representation

Existing authoring transaction path fields remain scalar transparency, diff, and write-confirmation outputs.
They are not replaced by the diagnostic `location` object.

The D02 normalized repository-relative grammar applies to:

- proposal `target.path`;
- every `diff.files[].path`;
- repository path operands embedded in `diff.text`;
- every `accept_proposed_write.files_written[].path`.

Surface rules:

- `target.path` identifies the resolved current repository destination, including a create target that does not yet exist;
- `diff.files[].path` exactly matches the repository-relative path of each proposed changed file;
- matching `target.path`, `diff.files[].path`, unified diff operands, and `files_written[].path` use the same normalized repository-relative spelling;
- `files_written` lists only files actually modified by the accept call and is empty when `written: false`;
- post-write validation failure does not erase paths of files that were actually written;
- authoring requests continue to use record identity rather than physical path as the primary target input.

Unified diff rules:

- modify patches use Git-style `diff --git a/<path> b/<path>`, `--- a/<path>`, and `+++ b/<path>` forms;
- create patches may use `--- /dev/null` and `+++ b/<path>`;
- `a/` and `b/` are Git side prefixes and are not part of the repository-relative path value;
- `/dev/null` is a unified-diff sentinel and is not a repository or physical source path;
- host absolute paths and backslash-separated Windows paths are prohibited in patch text;
- `get_proposed_write` returns the same patch-equivalent representation.

`diff_mode` exposure remains:

| mode or operation | path surfaces |
|---|---|
| `none` | `target.path`; diff body is intentionally omitted. |
| `summary` | `target.path` and `diff.files[].path`. |
| `patch` | `target.path`, `diff.files[].path`, and `diff.text`. |
| `get_proposed_write` | Patch-equivalent path surfaces. |
| successful `accept_proposed_write` | `files_written[].path` for files actually written. |

D05 does not change proposal lifecycle, validation scope, accept eligibility, write behavior, or operation wrappers.

### D06 accepted privileged absolute-path boundary

No current DRMCP operation may expose an absolute physical path.

The prohibition applies to:

- `list_records`;
- `get_records`;
- `resolve_reference`;
- `validate_records`;
- authoring proposal responses;
- proposal-local validation;
- authoring operation diagnostics;
- `diff`;
- `files_written`;
- the shared warning and diagnostic envelope.

A future absolute-path exposure requires a separately specified privileged debug or emergency operation.
T04 does not create that operation and does not add a hidden debug flag to an existing operation.

A future privileged contract must define at least:

- explicit host-side enablement;
- request and response shapes;
- caller privilege boundary;
- configured-root or discovered-source scope;
- exposure purpose;
- failure and redaction behavior.

Future representation rules:

- portable `location` remains unchanged and uses the D01 through D03 contract;
- absolute host detail appears only in a separate field such as `physical_path`;
- Windows drive-qualified, UNC, device, or host-native path syntax is allowed only in that privileged field;
- caller-supplied arbitrary strings are not echoed as physical paths;
- values originate only from configured roots, discovered candidates, or canonicalization results within the operation's declared scope;
- the surface is not a general-purpose filesystem read interface;
- `physical_path` never participates in semantic identity, deterministic ordering, or duplicate suppression;
- when portable `location` is constructible, `physical_path` supplements it rather than replacing it.

Root escape, symlink, junction, reparse-point, or canonicalization investigation may expose a rejected candidate's physical path only when the future privileged operation explicitly authorizes that evidence.

### D07 accepted missing-required-location failure boundary

A source-less condition may omit `location` exactly where D04 permits omission.
A source-backed warning or diagnostic that requires a direct, conflict-member, or conflict-candidate location is valid only when every required location can be constructed under D01 through D03.

Fail-closed rules:

- do not emit the source-backed entry without its required location;
- do not silently drop the affected item, conflict member, or legacy candidate;
- do not return a partial location object;
- do not substitute an opaque source token;
- do not fall back to an absolute physical path;
- do not weaken or remap the original category or severity to continue;
- do not emit an incomplete conflict collection;
- treat the inability to construct the required machine-readable response as operation execution failure rather than repository semantic invalidity.

Operation behavior:

| operation | failure behavior |
|---|---|
| `validate_records` | Return no normal `{ ok, scope, summary, diagnostics }` wrapper and do not continue with partial validation output. |
| `list_records`, `get_records`, `resolve_reference` | Fail the operation when a required warning, member, candidate, or unreadable-source location cannot be constructed. This is not W004 per-item partial success or a W005 unresolved status. |
| Proposal creation | Create no retained proposal and write no file; use the proposal-preparation failure boundary. |
| `accept_proposed_write` before writing | Return `written: false` with `files_written: []` and begin no write when an affected target or required source-backed diagnostic location cannot be constructed. |
| Conflict aggregation | If any required member or candidate location is unavailable, return no incomplete conflict and fail the operation. |

A separate shared diagnostic for “location unavailable” is not introduced because it would itself fail to provide the required repair target.

If an implementation violates this invariant after repository files have already been modified, it must not report `written: false` or erase the actual `files_written` state.
The exact fatal transport or implementation-failure representation remains owned by the applicable operation and runtime boundary; T04 does not invent a post-write rollback or false success model.

### Normative reflection

The accepted D01 through D07 contract was reflected after every decision was accepted.

Direct normative changes:

| file | reflected contract |
|---|---|
| `drmcp/records/spec/design-records-mcp/schema/diagnostics.md` | Concrete current and legacy `location` shape, portable path grammar, semantic identity, stable sort key, operation exposure matrix, missing-location fail-closed behavior, and absolute-path boundary. |
| `drmcp/records/spec/design-records-mcp/tools/validate-records.md` | Required diagnostic-location construction is part of trustworthy execution; inability to construct one prevents the normal validation wrapper. |
| `drmcp/records/spec/design-records-mcp/tools/list-records.md` | A missing-compact-field warning whose required source location cannot be constructed causes operation execution failure rather than a partial listing response. |
| `drmcp/records/spec/design-records-mcp/tools/get-records.md` | Required warning, conflict-member, conflict-candidate, or unreadable-source location failure is outside W004 partial success and prevents a normal response. |
| `drmcp/records/spec/design-records-mcp/tools/resolve-reference.md` | Required conflict-member, conflict-candidate, or unreadable-source location failure prevents a normal resolver response and does not become an unresolved status. |
| `drmcp/records/spec/design-records-mcp/schema/authoring-transaction-schema.md` | Normalized repository-relative target path grammar and proposal-local required-location fail-closed behavior. |
| `drmcp/records/spec/design-records-mcp/tools/authoring-transaction-model.md` | Shared target, diff, patch, and written-file path representation plus proposal and pre-write required-location checks. |
| `drmcp/records/spec/design-records-mcp/tools/propose-record-create.md` | Create-patch example now uses Git-style repository-relative `diff --git`, `/dev/null`, and `b/<path>` syntax. |
| `drmcp/records/spec/design-records-mcp/tools/accept-proposed-write.md` | Normalized repository-relative `files_written` paths, pre-write required-location failure, and truthful post-write state. |

The read-operation changes are narrow D07 response-construction corrections.
They do not change W004 warning triggers, request order, ordered deduplication, partial-success outcomes when representation succeeds, wrappers, or successful record projections.
They do not change W005 grammar order, fallback eligibility, public resolver statuses, or successful target projection.

No new execution-failure identifier was invented for `list_records`, `get_records`, or `resolve_reference`.
Those contracts define the fail-closed boundary and leave exact runtime error representation to the response and implementation boundary.

### Conditional and recheck disposition

| file | disposition |
|---|---|
| `schema/record-model.md` | Rechecked; no change. Existing current source provenance already supplies `app_namespace`, kind, and repository-relative path. |
| `namespace-scanning.md` | Rechecked; no change. Existing current and legacy root configuration already supplies repository-relative `records_root` and containment behavior. |
| `schema/discovery.md` | Rechecked; no change. Discovery and invalid-source retention ownership remain W003. |
| `schema/record-source.md` | Rechecked; no change. Source body and heading representation do not own exposed location shape. |
| `resolver.md` | Rechecked; no change. Resolver orchestration and public status mapping remain W005-owned. |
| `tools/propose-record-update.md` | Rechecked; no change. Its existing Git-style modify patch and shared authoring-model pointer already satisfy D05 and D07. |
| `tools/get-proposed-write.md` | Rechecked; no change. It already returns the shared patch-equivalent proposal representation. |
| `overview.md`, `tools/overview.md`, `mvp-scope.md`, `responsibility-boundary.md` | Deferred to T05 final synchronization. |

No dedicated debug or emergency operation was introduced.
No current operation exposes `physical_path` or another absolute host path.

### Changed-file manifest

T04 files changed in this phase:

- `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-006-04-define-source-location-and-exceptional-path-exposure-contract.md`;
- `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-006-validation-diagnostics-and-path-exposure-contract-realignment.md`;
- `drmcp/records/spec/design-records-mcp/schema/diagnostics.md`;
- `drmcp/records/spec/design-records-mcp/tools/validate-records.md`;
- `drmcp/records/spec/design-records-mcp/tools/list-records.md`;
- `drmcp/records/spec/design-records-mcp/tools/get-records.md`;
- `drmcp/records/spec/design-records-mcp/tools/resolve-reference.md`;
- `drmcp/records/spec/design-records-mcp/schema/authoring-transaction-schema.md`;
- `drmcp/records/spec/design-records-mcp/tools/authoring-transaction-model.md`;
- `drmcp/records/spec/design-records-mcp/tools/propose-record-create.md`;
- `drmcp/records/spec/design-records-mcp/tools/accept-proposed-write.md`.

### Independent review prompt

```text
C:\Users\imved\projects\brewprint

`DRMCP-TASK-MCP-006-04`の独立reviewを行う。

ファイルは変更しないこと。

DRMCPは現在利用できない。
repositoryの読み込みにはfilesystem MCPを使用すること。
sandboxへrepositoryを複製しないこと。
repository-local commandを実行できない場合、実行したと捏造しないこと。

## 最初に読む

- `prompt_chappy.md`
- `product/records/spec/design-records/authoring-standards/task-authoring.md`
- `product/records/spec/design-records/authoring-standards/work-item-authoring.md`
- `product/records/spec/design-records/authoring-standards/spec-authoring.md`
- `product/records/spec/design-records/authoring-standards/writing-standard.md`
- `product/records/spec/design-records/authoring-standards/agent-authoring-policy.md`

## Accepted baseline

- `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-006-01-establish-validation-diagnostics-and-path-exposure-correction-baseline.md`
- `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-006-02-define-current-repository-and-relation-validation-execution-contract.md`
- `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-006-03-define-machine-readable-diagnostic-representation-and-semantic-invalidity-mapping.md`
- `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-006-validation-diagnostics-and-path-exposure-contract-realignment.md`
- `drmcp/records/adr/mcp/DRMCP-ADR-MCP-001-design-records-mcp-contract-baseline-and-realignment.md`
- `drmcp/records/requirements/mcp/DRMCP-REQ-MCP-001-multi-root-multi-namespace-mcp-tool-contract.md`
- `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-003-current-discovery-and-active-index-contract-realignment.md`
- `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-004-query-and-exact-retrieval-contract-realignment.md`
- `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-005-resolver-and-configured-legacy-fallback-contract-realignment.md`

## Review targets

- `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-006-04-define-source-location-and-exceptional-path-exposure-contract.md`
- `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-006-validation-diagnostics-and-path-exposure-contract-realignment.md`
- `drmcp/records/spec/design-records-mcp/schema/diagnostics.md`
- `drmcp/records/spec/design-records-mcp/tools/validate-records.md`
- `drmcp/records/spec/design-records-mcp/tools/list-records.md`
- `drmcp/records/spec/design-records-mcp/tools/get-records.md`
- `drmcp/records/spec/design-records-mcp/tools/resolve-reference.md`
- `drmcp/records/spec/design-records-mcp/schema/authoring-transaction-schema.md`
- `drmcp/records/spec/design-records-mcp/tools/authoring-transaction-model.md`
- `drmcp/records/spec/design-records-mcp/tools/propose-record-create.md`
- `drmcp/records/spec/design-records-mcp/tools/accept-proposed-write.md`

## Review scope

Review only T04 source-location and exceptional path-exposure behavior.
Do not redesign T02 validation selectors, subjects, relation lookup, wrapper, or `ok` semantics.
Do not reopen T03 category, severity, subject, target, occurrence, conflict, ordering, duplicate-suppression, or semantic-authority decisions.
Do not reopen W003 discovery and source-retention behavior, W004 warning triggers and successful projections, W005 resolver order and statuses, or authoring proposal lifecycle.
Do not perform T05 overview or final Work Item closure synchronization.

Confirm:

- D01 through D07 are explicit and internally consistent;
- current and legacy locations use one object with `source_scope`, `records_root`, `path`, and current-only `app_namespace`;
- location identity and sort key are deterministic and independent of scan or configuration order;
- repository-relative paths use `/`, reject absolute, drive, UNC, URI, `.`, `..`, and alias escape forms, and preserve repository spelling;
- direct versus conflict-member and conflict-candidate location placement is unambiguous;
- normal successful list, retrieval, and resolver projections remain path-free;
- unreadable unique legacy sources and conflicts expose the required repair locations;
- missing required location fails closed without partial entries, omitted conflict sources, opaque tokens, or absolute fallback;
- read-operation fail-closed behavior does not redefine W004 partial success or W005 unresolved status;
- proposal targets, diff summaries, unified patches, and `files_written` use matching repository-relative paths;
- Git `a/`, `b/`, and `/dev/null` syntax is distinguished from repository path values;
- no current operation exposes an absolute physical path or gains a hidden debug flag;
- no new debug or emergency operation is invented;
- conditional and recheck-only file dispositions are justified;
- the changed-file manifest is complete;
- changed normative specs pass scoped strict validation and all T04 files pass applicable whitespace checks.

Repository-local commands, when available:

`python -X utf8 product/src/tools/validate_spec.py <changed normative spec files> --strict --no-color`

`git diff --check -- <tracked T04 files>`

`git diff --no-index --check -- NUL drmcp/records/tasks/mcp/DRMCP-TASK-MCP-006-04-define-source-location-and-exceptional-path-exposure-contract.md`

Do not infer a clean working tree.
Do not use `git add .`.

出力形式:

1. Verdict: PASS / NEEDS REVISION
2. Blocking findings
3. Major findings
4. Minor findings
5. Advisories
6. D01-D07 assessment
7. Location shape and path portability assessment
8. Ordering and duplicate-suppression assessment
9. Operation exposure and fail-closed assessment
10. W003-W005 non-regression assessment
11. Authoring path assessment
12. Conditional-file and changed-file assessment
13. Validation evidence assessment
14. T04 closure readiness
```

### Independent review finding and correction

Initial independent review verdict: `NEEDS REVISION`.

`F-MIN-01` identified that Task metadata `outputs` listed only `spec:drmcp.design_records_mcp.schema.diagnostics` and `DRMCP-WORK-MCP-006` even though T04 modified nine normative specifications.
The normative contract, D01 through D07, body changed-file manifest, and validation evidence were otherwise assessed as PASS-equivalent.

Correction applied on 2026-06-28:

- added `spec:drmcp.design_records_mcp.tools.validate_records`;
- added `spec:drmcp.design_records_mcp.tools.list_records`;
- added `spec:drmcp.design_records_mcp.tools.get_records`;
- added `spec:drmcp.design_records_mcp.tools.resolve_reference`;
- added `spec:drmcp.design_records_mcp.schema.authoring_transaction_schema`;
- added `spec:drmcp.design_records_mcp.tools.authoring_transaction_model`;
- added `spec:drmcp.design_records_mcp.tools.propose_record_create`;
- added `spec:drmcp.design_records_mcp.tools.accept_proposed_write`.

The corrected metadata now lists every normative specification in the T04 changed-file manifest plus the W006 Work Item.
No normative specification changed for this correction.

Post-`F-MIN-01` correction whitespace verification was executed externally on 2026-06-28 before the first limited re-review:

- the tracked W006 Work Item check returned `tracked_exit=0`;
- the untracked T04 Task check returned `untracked_exit=1`;
- neither check reported a whitespace error;
- no exit code `2` or greater occurred;
- exit code `1` for the untracked Task is expected because the new file differs from `NUL`.

The first limited re-review closed `F-MIN-01` but returned `NEEDS REVISION` with `F-MIN-02` because these post-correction results had not yet been recorded in the Task and synchronized to W006.

`F-MIN-02` correction applied on 2026-06-28:

- recorded the post-`F-MIN-01` `tracked_exit=0` and `untracked_exit=1` results in this Task;
- synchronized the same evidence to W006;
- removed the stale statement that corrected whitespace verification remained pending;
- changed no normative specification.

The second limited re-review closed `F-MIN-02` but reported `F-MIN-03` because it required final whitespace results for the current Task and W006 bytes to be embedded back into those same files.

That requirement is self-invalidating: writing the final results into either checked file changes the bytes that were checked.
T04 therefore uses this final verification boundary:

- Task and W006 evidence are finalized first;
- whitespace checks run after the last evidence edit;
- their raw command output is supplied directly to the final limited re-review;
- the final results are not written back into either checked file;
- any later edit requires a new external final check.

No normative specification changed for this closure correction.

### Current verification state

- Required instruction and authoring standards: read.
- Exact W006 Task inventory: complete.
- Authority and upstream records: read.
- Required source, root, diagnostic, validation, read-operation, resolver, and authoring contracts: read.
- T04 Task creation: complete.
- W006 T04 linkage and start Evidence: synchronized.
- D01 through D07: accepted.
- Normative specification reflection: complete.
- Conditional and recheck-only disposition: complete.
- Static cross-contract consistency recheck: complete.
- External scoped strict validation executed on 2026-06-28 for all nine changed normative specifications: `[strict]  All 9 file(s) OK.`
- External tracked-file whitespace verification executed for the nine changed normative specifications and W006 Work Item; `git diff --check` reported no whitespace error.
- LF-to-CRLF messages for the tracked files are non-blocking working-copy conversion notices and do not indicate `git diff --check` failure.
- External untracked Task whitespace verification used `git diff --no-index --check -- NUL <T04 Task>` and returned exit code `1` with no whitespace error.
- Exit code `1` is expected because the new Task differs from `NUL`; no exit code `2` or greater occurred.
- Targeted status confirmed exactly nine modified normative specifications, the modified W006 Work Item, and the untracked T04 Task listed in the changed-file manifest.
- Initial independent review verdict: `NEEDS REVISION` with one minor finding, `F-MIN-01`.
- `F-MIN-01`: closed by the first limited re-review.
- First limited re-review verdict: `NEEDS REVISION` with one new minor finding, `F-MIN-02`.
- Post-`F-MIN-01` correction whitespace evidence: recorded and synchronized (`tracked_exit=0`, `untracked_exit=1`, no whitespace error, no exit code `2` or greater).
- `F-MIN-02`: closed by the second limited re-review.
- Second limited re-review verdict: `NEEDS REVISION` with `F-MIN-03` concerning final-evidence self-reference.
- Final verification boundary: current Task and W006 bytes are checked externally after this last edit; results are supplied to review and are intentionally not written back into the checked files.
- Pre-closure final external whitespace verification passed with `tracked_exit=0` and `untracked_exit=1`; no whitespace error and no exit code `2` or greater were reported.
- Final limited independent re-review verdict: `PASS`.
- `F-MIN-01`: closed.
- `F-MIN-02`: closed.
- `F-MIN-03`: closed.
- No blocking, major, minor, or advisory findings remain.
- Final review accepted the external-evidence boundary and confirmed T04 closure readiness.
- Task status changed to `done` on 2026-06-28.
- Because closure synchronization changes the checked Task and W006 bytes, one post-closure external whitespace check must run after the final closure edits.
- That post-closure result is supplied externally and is intentionally not written back into either checked file.
- Repository-wide clean status: not inferred.
