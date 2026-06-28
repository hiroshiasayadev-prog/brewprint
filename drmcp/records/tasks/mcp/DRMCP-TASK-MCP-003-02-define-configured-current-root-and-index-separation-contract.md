# DRMCP-TASK-MCP-003-02: Define configured current-root and index-separation contract

- **id**: DRMCP-TASK-MCP-003-02
- **status**: done
- **date**: 2026-06-27
- **work_item**: DRMCP-WORK-MCP-003
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 1d
- **depends_on**:
  - DRMCP-TASK-MCP-003-01
- **outputs**:
  - spec:drmcp.design_records_mcp.namespace_scanning

## Goal

Define the authoritative DRMCP contract for configured current roots, app association, active-index scope, and current/legacy index separation.

Remove repository-wide root auto-discovery and `V01-`-centric namespace derivation without entering current spec parsing, query, resolver, validation, fixture, or implementation scope.

## Work

- Define the configuration input for one or more current records roots.
- Define how each configured current root is associated with one app namespace.
- Prohibit repository-wide `*/records/` auto-discovery.
- Define invalid current-root, duplicate-root, and overlapping current/legacy root conditions at the discovery-contract level.
- Define the active read index over current records from configured current roots, while preserving invalid but uniquely addressable sources for later validation and repair.
- Define the optional legacy index as a separate structure built only from configured legacy roots.
- Exclude legacy roots and records from active-index construction.
- Define deterministic duplicate canonical-identity conflict handling across current roots without selecting a winner by filesystem order.
- Keep exact diagnostic codes and response fields delegated to `DRMCP-WORK-MCP-006`.
- Rewrite `spec:drmcp.design_records_mcp.namespace_scanning` against the accepted PRODUCT and DRMCP authorities.

This Task does not define current spec metadata parsing or path-derived spec identity details.
Those contracts belong to `DRMCP-TASK-MCP-003-03`.

## Done condition

- Current roots are configuration-only inputs.
- No active contract auto-discovers `*/records/` directories.
- Each current root has one unambiguous app association.
- Active and optional legacy indexes are separate structures and operational scopes.
- Legacy roots do not contribute records to the active index.
- Duplicate canonical identity across current roots makes only the conflicted identity unavailable for normal ID-based retrieval, without filesystem-order selection.
- Root configuration failure conditions are identified without fixing W006 diagnostic taxonomy.
- Query, retrieval, resolver, validation execution, fixtures, and implementation remain excluded.
- The corrected namespace-scanning contract cites PRODUCT authorities instead of restating their semantics.

## Verification

- Compare the updated contract against `DRMCP-ADR-MCP-001`, `DRMCP-REQ-MCP-001`, and `PRODUCT-WORK-SPEC-014`.
- Confirm removal of automatic `*/records/` discovery and `V01-`-specific examples from normative text.
- Confirm duplicate canonical identity is isolated as an entry conflict and does not make unrelated current records unreadable.
- Confirm that current spec source parsing remains delegated to T03.
- Confirm that diagnostic codes, tool response shape, fixtures, and implementation are not introduced.

## Evidence

- Upstream baseline: `DRMCP-TASK-MCP-003-01`.
- Compatibility gate: `PRODUCT-WORK-SPEC-014`, completed 2026-06-27.
- Target contract: `spec:drmcp.design_records_mcp.namespace_scanning`.
- Accepted decision: each configured current root is paired explicitly with one `app_namespace` value.
- Rejected alternative: derive the app namespace from the parent directory of `records_root`.
- Rejected alternative: auto-discover repository folders and suppress selected trees through folder-local `.ignore` files.
- Reason: physical placement and ignore-file distribution must not become namespace or active-index authority.
- Accepted decision: one `app_namespace` may be associated with exactly one configured current root.
- A repeated `app_namespace` entry is a configuration error; the runtime does not merge multiple current roots for one app.
- Reason: allowing multiple roots per app would introduce spec-tree partitioning, sequential-ID search scope, and authoring-target selection contracts outside this Work Item.
- Accepted decision: configuration declares one `repository_root`; every `records_root` is resolved as a path relative to that root.
- Relative `records_root` values are not resolved from the process working directory, and host-specific absolute paths are not the portable configuration form.
- Reason: one explicit resolution base keeps configuration reproducible and portable across runtime launch locations.
- Accepted decision: each current-root entry must satisfy `records_root == <app_namespace>/records` after repository-relative path normalization.
- A mismatched pair is a configuration error; DRMCP does not accept arbitrary alternate current-record locations.
- This is an integrity check between two explicit configuration values, not namespace inference from physical placement.
- Reason: enforcing the PRODUCT-owned placement rule prevents silent misassociation and hard-to-diagnose partial indexing.
- Accepted decision: every configured current root is mandatory.
- If any configured current root is missing, unreadable, or otherwise invalid, active-index construction fails as a whole.
- DRMCP does not omit the invalid root and continue with a partial active index.
- Reason: configuration correctness is the operator's responsibility, and partial startup would make missing-root failures indistinguishable from unresolved records.
- Accepted decision: `current_roots` must contain at least one entry.
- Missing or empty `current_roots` is a configuration error; DRMCP does not support legacy-only startup.
- Reason: current-format operation is the runtime baseline, while configured legacy roots are only an optional fallback layer.
- Accepted decision: after path normalization and canonical filesystem resolution, no configured current root and legacy root may be identical or stand in an ancestor/descendant relationship.
- Any overlap or containment between current and legacy roots is a configuration error.
- Reason: current and legacy indexes are separate operational scopes, and overlapping trees would make leakage prevention depend on fragile scan filters.
- Accepted decision: a valid configured current root may contain zero discoverable current records.
- Such a root contributes an empty app-scoped portion to the active index and does not cause startup failure.
- Root validity and record count are separate concerns: missing, unreadable, or misplaced roots fail configuration, while a valid empty root is allowed.
- Reason: records-only app namespaces may be introduced before their first record, and emptiness does not make the explicit configuration false.
- Changed file: `drmcp/records/spec/design-records-mcp/namespace-scanning.md`.
- Static verification confirmed:
  - no repository-wide `*/records/` auto-discovery remains;
  - app association is explicit and is not derived from directory names;
  - current roots are strict configuration inputs and invalid roots fail the complete active-index build;
  - valid empty current roots are allowed;
  - active and optional legacy indexes remain separate;
  - current/legacy root overlap is rejected;
  - current spec parsing remains delegated to T03;
  - exact diagnostic identifiers, tool response shape, fixtures, and implementation remain excluded.
- T03 clarification applied on 2026-06-27:
  - duplicate canonical identity is an index-entry conflict, not a configured-root failure;
  - the conflicted identity is unavailable for normal ID-based retrieval;
  - all conflicting sources remain validation inputs with provenance;
  - unaffected current records remain readable;
  - filesystem traversal order never selects a winner.
- Scoped validator execution remains part of T05 final synchronization.
