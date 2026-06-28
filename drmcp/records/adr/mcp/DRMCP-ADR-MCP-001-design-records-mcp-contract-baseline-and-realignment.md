# DRMCP-ADR-MCP-001: Design Records MCP contract baseline and realignment

- **status**: accepted
- **date**: 2026-06-25
- **depends_on**:
  - PRODUCT-ADR-SPEC-001
- **supersedes**:
  - V01-ADR-076
  - V01-ADR-077
  - V01-ADR-087
  - V01-ADR-088
  - V01-ADR-090
  - V01-ADR-092
  - V01-ADR-093
- **migrated_to_spec**: null

## Context

`DRMCP-INV-MCP-002` found 26 issues across the active DRMCP contract set.
Seventeen issues have major severity.

The current contracts mix three incompatible baselines:

- an early read-only MCP trial;
- a later authoring transaction trial;
- current PRODUCT-owned Design Records semantics.

Passing DRMCP tests prove consistency with legacy behavior.
They do not prove alignment with current Design Records contracts.

The largest conflicts concern spec format, authority ownership, phase naming, legacy compatibility, authoring guidance, and operation support.
A new implementation baseline is required before broad DRMCP contract or code changes begin.

## Decision

### Contract authority

DRMCP consumes Design Records semantics and does not redefine them.

| concern | authority |
|---|---|
| App-independent artifact semantics and responsibilities | PRODUCT Design Records specs. |
| Authoring standards and canonical section rules | PRODUCT Design Records specs. The portable package is their distribution. |
| App and domain namespace semantics | PRODUCT Design Records specs. |
| Canonical ID and reference grammar | PRODUCT Design Records specs. |
| Repository layout semantics | PRODUCT Design Records specs. |
| Traceability semantics | PRODUCT Design Records specs. |
| Canonical spec format | PRODUCT Design Records specs. |
| Current Brewprint profile and compatibility facts | Brewprint profile specs. |
| Parsing, discovery, indexing, and normalization | DRMCP specs. |
| Query, resolution, validation, and diagnostics | DRMCP specs. |
| MCP request and response contracts | DRMCP specs. |
| Proposal lifecycle and write transactions | DRMCP specs. |

DRMCP contracts cite semantic authorities instead of copying their normative rules.
DRMCP may define parser mappings and diagnostic representations for those rules.

DRMCP overview specs become navigation-first entry points.
Shared schema specs own shared structures.
Each operation contract owns its request, response, and error behavior.
An overview cannot override a schema or operation contract.

The current read-oriented record surface covers decision, spec, investigation, requirement, work item, and task.
These kinds share discovery, retrieval, resolution, and validation where corrected operation contracts support them.
Kind-specific semantics remain owned by PRODUCT contracts.

### Historical phase model

The previous phases are historical trials:

| phase label | meaning |
|---|---|
| `v0.x read-only trial` | Initial indexing, query, resolution, and validation work. |
| `v0.x authoring trial` | Proposal, cache, diff, validation, and accept-only write work. |
| Current contract realignment | Replacement baseline aligned with current PRODUCT semantics. |

The `v0.x` labels classify historical development.
They do not claim that repository tags or releases used exact `v0.x` semantic versions.

The current delivery sequence is:

1. Current-format discovery, indexing, query, resolution, and validation.
2. ADR, requirement, work-item, and task authoring.
3. Spec authoring.
4. Investigation authoring.

A later phase may depend on an earlier phase.
Later authoring support does not redefine the current-format read baseline.

### Current spec format

DRMCP supports the current spec format as the only active spec source format.

Current specs use:

- H1-adjacent visible metadata;
- path-derived canonical `spec:` refs;
- current PRODUCT spec kinds and section rules;
- no YAML front matter as a metadata source.

Legacy YAML-front-matter specs are not migrated into the active spec tree.
They are not indexed as active specs.
`V01-SPEC-*` is not a supported compatibility reference family.

Source-format compatibility and issued-ID compatibility are separate concerns.
Rejecting legacy spec sources does not rename issued sequential artifacts.

The existing `v01/` tree remains a read-only archive.
This decision does not require moving that tree under another archive directory.

### Legacy sequential-ID fallback

Legacy compatibility is configuration-gated.
DRMCP does not discover legacy archive roots automatically.

A configuration may declare one or more `legacy_roots`.
Without that declaration, legacy fallback is disabled.

Resolution follows this order:

1. Parse and resolve the input using current canonical grammar.
2. If unresolved, test the input against an accepted legacy grammar.
3. Query the legacy archive index only after an exact legacy grammar match.
4. Return unresolved or unsupported when neither grammar accepts the input.

Accepted fallback families are:

- `V01-ADR-*`;
- `V01-INV-*`;
- `V01-REQ-*`;
- `V01-WORK-*`;
- `V01-TASK-*`.

The fallback excludes:

- `V01-SPEC-*`;
- app-prefixless bare IDs;
- physical-path inference;
- fuzzy normalization;
- legacy YAML semantic-ref aliases.

The current and legacy indexes remain separate.
Legacy records are read-only.

| operation | legacy archive behavior |
|---|---|
| Normal listing | Excluded. |
| Exact retrieval | Allowed after an exact accepted legacy ID match. |
| Reference resolution | Allowed only as fallback. |
| Active-record relation validation | A current record may resolve an accepted legacy target. |
| Repository-wide current validation | Legacy archive records are excluded. |
| Create or update | Prohibited. |

Legacy resolution preserves the issued legacy ID.
It does not translate the ID into a current app-aware ID.

### Spec create identity and placement

Physical paths are implementation details in the normal authoring interface.
A caller supplies a logical spec create selector.

| selector form | requested node | persisted canonical ref |
|---|---|---|
| `spec:<segments>` | Leaf spec | `spec:<segments>` |
| `spec:<segments>.index` | Topic index spec | `spec:<segments>` |

The `.index` suffix is a create selector.
It is not stored in the canonical ref.
Without the `.index` suffix, the selector always targets a leaf spec.
DRMCP does not infer leaf or topic placement from repository state.

Example:

| create selector | internal placement | persisted canonical ref |
|---|---|---|
| `spec:drmcp.design_records_mcp.tools.index` | `design-records-mcp/tools/index.md` | `spec:drmcp.design_records_mcp.tools` |
| `spec:drmcp.design_records_mcp.tools.resolve_reference` | `design-records-mcp/tools/resolve-reference.md` | `spec:drmcp.design_records_mcp.tools.resolve_reference` |

DRMCP derives the physical path from the logical selector and app namespace.
The persisted `id` must still match the path-derived canonical ref.

A leaf and topic index cannot share the same canonical ref.
DRMCP rejects both conflict directions:

- creating `<topic>.md` when `<topic>/index.md` exists;
- creating `<topic>/index.md` when `<topic>.md` exists.

A topic index may coexist with child specs below that topic.
Create proposals never overwrite an existing spec path or canonical ref.

Normal request and response contracts use canonical identities.
Normal listing, retrieval, and proposal-summary responses do not expose physical paths.
Explicit patch, diagnostic, or debug output may include a path when required.

The PRODUCT spec authoring interface requires a coordinated update.
The update replaces author-supplied physical paths with logical create selectors.

### Portable authoring standards

Brewprint and DRMCP must support use outside this repository.
DRMCP therefore consumes a portable, self-contained standards package.

The package has these properties:

- a fixed app namespace independent of the host repository;
- an explicit package root and entrypoint selected by configuration;
- enough authoring mappings to perform supported create, update, and validation operations;
- no dependency on a host repository's `product` namespace;
- no Brewprint-specific app registry or V01 compatibility facts;
- no DRMCP request or response definitions;
- no BPDSL-specific rules.

The package includes the minimum transitive authoring contract required at runtime.
Examples include identity mapping, metadata shape, section rules, lifecycle gates, and logical placement mapping.

The package is a distribution of PRODUCT-owned semantics.
It does not become a second semantic owner.

Relative references, namespace remapping, and package-local alias syntax are deferred.
The package uses its fixed namespace for internal canonical refs.
The package namespace identifier is deferred to the requirement that defines the portable package contract.

Authoring guidance resolves from the standards package.
Legacy `records/guides/*.md` files are not a canonical guidance source.
DRMCP does not maintain a separate legacy guide parser as the current guidance model.
Any retained guidance tools act as projections over indexed standards-package specs.

### Tool disposition

`suggest_next_record` is removed.
It has no compatibility surface.

Namespace-aware `new` placeholders replace sequential-ID suggestion for supported sequential artifacts.
Spec creation uses the logical selector contract instead of numeric allocation.

The following authoring transaction mechanics remain valid design inputs:

- propose then accept;
- proposal retention and staleness checks;
- body cache;
- diff modes;
- affected-record validation;
- accept-only filesystem writes;
- exact named-section replacement;
- machine-readable diagnostics.

Retained mechanics must be rewritten where they depend on legacy spec format, path exposure, or duplicated PRODUCT semantics.

Exact batch retrieval remains an exact lookup operation.
It does not silently invoke reference normalization or fuzzy resolution.

### Follow-up sequence

The realignment follows this order:

1. Revise `DRMCP-REQ-MCP-001` and `DRMCP-REQ-MCP-002` against this ADR.
2. Create one coordinating DRMCP Work Item for contract and implementation realignment.
3. Update affected PRODUCT authoring and Brewprint compatibility contracts.
4. Coordinate validation-policy owner-pointer changes before replacing existing validation Work Items.
5. Correct DRMCP specs in small ownership-focused batches.
6. Add current-format and legacy-fallback fixtures.
7. Implement only against corrected and reviewable contracts.
8. Run an independent contract and implementation-plan review.

`DRMCP-WORK-SPEC-001` and `DRMCP-WORK-SPEC-002` remain disposition candidates.
They may be superseded or absorbed only with matching PRODUCT validation-policy pointer updates.

## Rationale

A single current-format baseline removes incompatible parser and authoring assumptions.
A separate legacy index keeps historical compatibility out of normal query and validation behavior.

Exact legacy grammar matching prevents accidental normalization.
Configuration gating prevents unrelated repositories from inheriting Brewprint history.

Logical spec selectors preserve path-derived identity without exposing repository placement as the public authoring interface.
The `.index` selector also resolves the leaf-versus-topic ambiguity before file creation.

A self-contained standards package makes authoring portable.
The package also avoids copying only one standards directory while leaving broken transitive references.

Phased authoring support reduces contract and implementation risk.
Spec and investigation authoring each require distinct mappings beyond sequential workflow artifacts.

Removing `suggest_next_record` avoids a second sequence-allocation contract.
The proposal transaction model already provides the correct allocation boundary.

`V01-ADR-090` and `V01-ADR-093` are superseded as decision authorities.
Their exact batch-retrieval boundary and proposal-transaction mechanics remain retained design inputs under this baseline.

## Rejected alternatives

| alternative | rejection reason |
|---|---|
| Keep the read-only and authoring trials under one MVP label | The contracts assign incompatible write boundaries to the same phase. |
| Treat exact historical phase numbers as releases | No matching repository release evidence was found. |
| Keep YAML-front-matter specs as active compatibility records | The old source model would remain in parser, validator, authoring, and tests. |
| Migrate legacy specs into the active tree | The legacy specs are not current contracts. |
| Keep `V01-SPEC-*` resolution | The alias would preserve obsolete spec identity and source-format behavior. |
| Merge current and legacy records into one index | Legacy records would leak into normal listing, validation, and authoring. |
| Auto-discover `v01/` or other legacy roots | Portability requires legacy behavior to be explicit and host-controlled. |
| Accept app-prefixless bare IDs | No current authority accepts them as canonical or compatibility IDs. |
| Accept physical paths as primary authoring inputs | Physical placement is not canonical identity and should remain hidden. |
| Infer leaf or topic placement from repository state | The same new canonical ref would remain ambiguous before creation. |
| Allow both `<topic>.md` and `<topic>/index.md` | Both paths derive the same canonical ref and confuse authors and tools. |
| Copy only the current authoring-standards directory | Current authoring rules depend on sibling semantic contracts. |
| Add relative refs or namespace remapping now | A fixed package namespace satisfies the first portable distribution. |
| Retain legacy guide files as canonical guidance | Legacy guides are neither current standards nor portable package contracts. |
| Retain `suggest_next_record` as a convenience tool | `new` placeholders already own sequential allocation. |
| Deliver all authoring kinds in one phase | Spec and investigation support add separate parser and lifecycle risks. |

## Consequences

- The active DRMCP spec tree requires broad correction before implementation work resumes.
- Current spec parsing and authoring code cannot be treated as the target contract.
- Legacy YAML spec fixtures and `V01-SPEC-*` tests must be removed from the current runtime baseline.
- Accepted V01 sequential IDs remain readable only through configured archive fallback.
- `spec:product.brewprint.compatibility.legacy_id_compatibility` must remove the `V01-SPEC-*` accepted-family entry and its compatibility-only spec identity rule before legacy fallback is implemented.
- PRODUCT spec authoring must define logical create selectors and leaf/topic collision rules.
- A portable standards package must be created before portable authoring is complete.
- Normal DRMCP responses become less path-oriented.
- Patch and diagnostic surfaces may still reveal paths when operationally necessary.
- Existing proposal mechanics can be retained after semantic and response cleanup.
- `DRMCP-REQ-MCP-001` and `DRMCP-REQ-MCP-002` remain separate but coordinated requirements.
- No implementation task may use the current DRMCP contracts as-is.

## Evidence

- `DRMCP-INV-MCP-002`: 30-file contract audit, 26 findings, and implementation drift analysis.
- `PRODUCT-ADR-SPEC-001`: Accepted semantic ownership boundary.
- `spec:product.design_records.authoring_standards`: Current authoring authority.
- `spec:product.design_records.namespace_model`: Current namespace and sequential-ID authority.
- `spec:product.design_records.repository_layout`: Current placement authority.
- `spec:product.design_records.spec_format`: Current spec format and identity authority.
- `spec:product.brewprint.compatibility`: Parent Brewprint compatibility authority.
- `spec:product.brewprint.compatibility.legacy_id_compatibility`: Compatibility authority that must drop the `V01-SPEC-*` family and compatibility-only spec identity rule.
- `DRMCP-REQ-MCP-001`: Multi-root and multi-namespace query and resolution gaps.
- `DRMCP-REQ-MCP-002`: Namespace-aware authoring transaction requirements.
- V01 ADRs listed in `supersedes`: Historical DRMCP trial decisions replaced by this baseline.
