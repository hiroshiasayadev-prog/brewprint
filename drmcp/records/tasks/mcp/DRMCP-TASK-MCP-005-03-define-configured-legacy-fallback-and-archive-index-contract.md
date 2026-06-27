# DRMCP-TASK-MCP-005-03: Define configured legacy fallback and archive-index contract

- **id**: DRMCP-TASK-MCP-005-03
- **status**: done
- **date**: 2026-06-27
- **work_item**: DRMCP-WORK-MCP-005
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 2d
- **depends_on**:
  - DRMCP-TASK-MCP-005-02
- **outputs**:
  - spec:drmcp.design_records_mcp.namespace_scanning
  - spec:drmcp.design_records_mcp.resolver
  - spec:drmcp.design_records_mcp.tools.resolve_reference
  - spec:drmcp.design_records_mcp.tools.get_records
  - DRMCP-WORK-MCP-005

## Goal

Define the configured legacy-fallback contract after the accepted current-first resolver stage.

Specify legacy-root configuration, a minimal filename-keyed exact lookup map, duplicate handling, source-first exact retrieval, final legacy outcomes, and a minimal non-path resolver target.

## Work

- Confirm the T01 and T02 authority, operation, and ownership boundaries.
- Define each `legacy_roots` entry and the meaning of missing or empty configuration.
- Define relative-path resolution, current-root separation, legacy-root duplicates, and legacy-root overlap behavior.
- Define missing, unreadable, or otherwise unusable legacy-root behavior without weakening mandatory current-root validation.
- Consume the accepted legacy-family set from Brewprint compatibility authority without local extension, while defining DRMCP-owned lexical parser mapping for those families.
- Define filename-derived issued legacy identity without adding legacy exceptions to current metadata grammar.
- Define a minimal exact, case-sensitive issued-ID-to-source lookup map separate from the active index.
- Define duplicate issued-ID conflict behavior without selecting a filesystem-order winner.
- Define source-first exact retrieval that preserves readable body access without current-model metadata normalization.
- Define final resolver outcomes for disabled fallback, missing target, duplicate conflict, unreadable source, and successful resolution.
- Define the minimal successful legacy target discriminator, issued identity, and kind without physical paths or parsed source fields.
- Synchronize only the normative contracts required by the accepted decisions.
- Record why each candidate normative file changed or remained unchanged.
- Run scoped strict validation for the changed normative specs.
- Prepare an independent review prompt without changing this Task to `done`.

## Done condition

- `legacy_roots` entry fields and path-resolution rules are explicit.
- Missing and empty legacy-root configuration have one defined fallback-disabled meaning.
- Current and legacy roots remain separate and non-overlapping.
- Duplicate and overlapping legacy-root entries have deterministic configuration behavior.
- Missing, unreadable, or unusable legacy-root behavior is explicit.
- The accepted legacy-family set is consumed only from Brewprint compatibility authority; DRMCP defines only the lexical parser mapping required for recognition and filename extraction.
- Legacy lookup does not weaken current source or metadata grammar.
- The separate minimal issued-ID-to-source lookup map is explicit.
- Family-specific legacy filename grammar yields one deterministic issued ID and suffix boundary.
- Legacy source enumeration does not traverse or adopt symlinks, junctions, reparse points, or sources outside the configured legacy boundary.
- Readable uniquely identified legacy sources remain retrievable even when optional parsed fields are missing or malformed.
- Duplicate issued legacy IDs never select one source by traversal order.
- Exact legacy lookup is case-sensitive and keyed by the issued legacy ID.
- Disabled, missing, duplicate, and unreadable accepted legacy cases map to `unresolved`; successful lookup maps to `resolved`.
- `unsupported` remains limited to inputs accepted by neither current nor accepted legacy grammar.
- Successful legacy targets preserve the issued legacy ID and expose only `target_type`, `ref`, and `kind`.
- `get_records` classifies accepted legacy-family overlap as legacy before current exact families and performs one lookup only.
- `get_records` and `resolve_reference` remain separate operations.
- W003 current discovery and active-index contracts remain unchanged.
- W004 request, ordering, partial-success, and common response-wrapper contracts remain unchanged; only the W005-owned minimal legacy source projection is synchronized.
- W006 diagnostic, warning, severity, message, source-location, validation, and exceptional path representation remain delegated.
- Every changed normative spec passes scoped strict validation.
- Independent review reports no blocking, major, or minor findings before this Task is marked `done`.

## Verification

- Compare accepted legacy families with `spec:product.brewprint.compatibility.legacy_id_compatibility`.
- Compare current-first ordering and fallback gating with `DRMCP-ADR-MCP-001`, `DRMCP-REQ-MCP-001`, and T02.
- Compare current and legacy index separation with final W003 contracts.
- Compare exact retrieval separation with final W004 contracts.
- Confirm that current metadata grammar contains no legacy exception.
- Confirm that duplicate issued IDs produce no selected winner.
- Confirm that legacy filename grammar deterministically separates issued ID from optional slug for every accepted family.
- Confirm that legacy enumeration does not traverse or adopt filesystem aliases or canonically escaped sources.
- Confirm that `get_records` classifies accepted legacy-family matches before current exact families without resolver invocation or fallback.
- Confirm that normal successful legacy targets contain no physical path, source location, provenance, index path, or resolver trace.
- Confirm that W006-owned diagnostic representation remains absent from changed W005 contracts.
- Run the scoped validator against only changed normative specs.
- Run an independent review before changing status to `done`.

## Evidence

### Accepted upstream baseline

- T01 authority, affected-file, and ownership baseline: complete.
- T02 current-first resolver contract: complete.
- Current grammar and active-index lookup run before legacy grammar evaluation.
- Accepted legacy-family grammar is independent of `legacy_roots` availability.
- `unsupported` is limited to inputs accepted by neither current nor accepted legacy grammar.
- Current successful target projection and exact-retrieval separation are closed upstream.
- Diagnostic and physical-path representation remains delegated to W006.

### Accepted decision D01: `legacy_roots` entry shape

Each `legacy_roots` entry contains one required field:

```yaml
legacy_roots:
  - records_root: v01/records
```

- `records_root` is a repository-relative path resolved from configured `repository_root`.
- An entry does not declare `app_namespace`, namespace prefix, root namespace, or archive identity.
- The canonicalized filesystem root is internal source provenance and configuration identity only.
- The issued legacy ID is the archive-record identity and archive-index key.
- Root identity does not qualify, rewrite, or partition the issued legacy ID.
- Physical root paths do not appear in normal successful resolver targets.

Reason: Accepted V01 IDs already provide globally comparable issued identity. Adding an archive or app namespace would create a second identity axis without changing exact lookup semantics.

### Accepted decision D02: legacy-root validity and failure scope

- Missing `legacy_roots` and an explicit empty list are equivalent valid current-only configurations.
- Either form disables legacy fallback and does not affect active-index startup.
- When one or more legacy roots are configured, every entry is mandatory.
- A configured legacy root must exist, be a readable directory, remain inside `repository_root` after canonical resolution, and be unique after canonicalization.
- Legacy roots must not equal, contain, or be contained by another configured legacy root.
- Legacy roots must not equal, contain, or be contained by a configured current root.
- Link or alias resolution that creates equality or overlap is invalid.
- Any invalid configured legacy root fails DRMCP startup and prevents creation of a partial legacy archive index.
- DRMCP does not drop only the invalid root and continue with remaining legacy roots.
- A valid legacy root may contain zero accepted legacy sources.
- An invalid individual legacy source does not invalidate its root. Source validity and root validity are separate concerns.

Reason: Explicit configuration is an operator assertion. Partial archive construction would turn configuration failure into a misleading not-found result.

### Accepted decision D03: legacy source candidate and issued-ID authority

- DRMCP scans configured legacy roots recursively for regular Markdown files.
- A file is an accepted legacy candidate only when its filename begins with one exact accepted legacy-family ID.
- Accepted filename forms are `<issued-id>.md` and `<issued-id>-<slug>.md`.
- The filename-derived issued legacy ID is the archive-record identity and archive-index key.
- Identity comparison is exact and case-sensitive.
- Record `kind` is derived from the accepted issued-ID family.
- H1 ID and metadata `id`, when present, are consistency values only.
- A missing or malformed H1 does not remove addressability when the filename yields one unique accepted issued ID.
- An H1 or metadata ID mismatch does not create an alias, replacement identity, or second index key.
- Title and lifecycle status are parsed when available. Missing or invalid values remain missing or invalid.
- Legacy parsing uses a dedicated legacy parser boundary. It does not invoke or weaken the current parser or current metadata grammar.
- Files whose names do not yield an accepted issued legacy ID are outside the archive index.
- DRMCP parser mapping fixes exact issued-ID forms as `V01-ADR-<SEQ3>`, `V01-INV-<DOMAIN>-<SEQ3>`, `V01-REQ-<DOMAIN>-<SEQ3>`, `V01-WORK-<DOMAIN>-<SEQ3>`, and `V01-TASK-<DOMAIN>-<WORK_SEQ3>-<TASK_SEQ2>`.
- `<DOMAIN>` matches `[A-Z][A-Z0-9]*`; the sequence widths are exact and ASCII-digit-only.
- An optional filename slug begins only after the complete family-specific issued ID, is non-empty, and matches `[a-z0-9][a-z0-9-]*`.
- Symlinked, junction, reparse-point, and other filesystem-alias files are not candidates; alias directories are not traversed.
- Every candidate must remain canonically inside its configured legacy root and `repository_root` and outside current roots.
- A candidate-level boundary violation excludes only that source and does not invalidate the configured root or unrelated entries.

Reason: Filename identity preserves exact retrieval of malformed historical sources while keeping legacy exceptions out of the current H1-authoritative model.

### Superseded decision D04: dedicated legacy archive schema

The initial D04 selected a dedicated normalized legacy archive schema. This was withdrawn before normative reflection because it exceeded the required compatibility surface.

### Accepted decision D04R: minimal legacy lookup model

- Do not create a dedicated legacy archive schema.
- Do not normalize legacy records into the current record model.
- Do not add legacy records to `list_records`.
- Build only a separate exact lookup map from filename-derived issued legacy ID to source file.
- Preserve the configured legacy roots as read-only inputs.
- `get_records` may return available H1, status, headings, and complete source Markdown, but compatibility does not require full current metadata normalization.
- `resolve_reference` requires only exact existence lookup and a minimal non-path legacy target.
- Legacy sources remain outside current repository-wide validation and authoring.
- Duplicate issued IDs produce no selected winner.

Reason: The required compatibility surface is exact retrieval and fallback reference resolution, not current-model parity for archived records.

### Accepted decision D05: source-first exact legacy retrieval

- Exact legacy retrieval succeeds when the source file can be read, even when optional Markdown structure or metadata cannot be parsed completely.
- `ref` preserves the filename-derived issued legacy ID.
- `kind` is derived from the accepted legacy ID family.
- When `include_body` is `true`, `body` contains the complete source Markdown verbatim.
- `title`, `status`, and headings are returned only when they can be parsed from the source.
- Missing or malformed optional parsed values do not block body retrieval.
- DRMCP does not repair, default, normalize, or validate archived metadata as a condition of retrieval.
- When the indexed source file cannot be read, that ref does not produce a successful record.
- Warning and diagnostic representation for unreadable sources remains owned by `DRMCP-WORK-MCP-006`.

Reason: Legacy compatibility guarantees exact access from an issued ID to retained source, not current-model metadata conformance.

### Accepted decision D06: final legacy outcome mapping

- Return `resolved` only when one readable legacy source exists for the exact accepted issued legacy ID.
- Return `unresolved` when legacy fallback is disabled because `legacy_roots` is missing or empty.
- Return `unresolved` when the configured legacy lookup map contains no source for the accepted issued legacy ID.
- Return `unresolved` when duplicate sources prevent selection of one source for the issued legacy ID.
- Return `unresolved` when the selected indexed source cannot be read.
- Return `unsupported` only when neither current canonical grammar nor accepted legacy-family grammar accepts the input.
- Do not add public resolver statuses such as `disabled`, `unavailable`, or `conflicted`.
- W006-owned warning and diagnostic representation distinguishes disabled fallback, missing target, duplicate conflict, and unreadable source where required.

Reason: Existing public statuses are sufficient. Operational distinctions belong in diagnostics rather than a larger resolver state vocabulary.

### Accepted decision D07: minimal successful legacy target

A successful legacy resolver target contains exactly:

```json
{
  "target_type": "legacy_sequential_record",
  "ref": "V01-REQ-MCP-001",
  "kind": "requirement"
}
```

- `target_type` is fixed to `legacy_sequential_record`.
- `ref` preserves the issued legacy ID exactly.
- `kind` is derived from the accepted legacy ID family.
- The target does not include `title`, `status`, body, headings, metadata, physical path, source location, provenance, index state, or resolver trace.
- Callers use `get_records` with the returned `ref` when archived source content is required.

Reason: Resolver success proves one exact archived target exists. Source projection remains the responsibility of exact retrieval.

### Open decision register

No semantic design decision remains. T03 must now record changed-file and unchanged-file disposition during normative reflection.

### Normative reflection disposition

Changed:

- `spec:drmcp.design_records_mcp.namespace_scanning`: added explicit `legacy_roots` fields, disabled semantics, strict root validation, family-specific filename grammar, filesystem-alias containment rules, exact lookup-map construction, duplicate handling, and active-index/listing exclusion.
- `spec:drmcp.design_records_mcp.resolver`: added configured legacy lookup outcomes and the minimal successful target boundary.
- `spec:drmcp.design_records_mcp.tools.resolve_reference`: added final three-status mapping and the exact three-field legacy target.
- `spec:drmcp.design_records_mcp.tools.get_records`: added legacy-first exact classification for overlapping grammar and source-first minimal legacy retrieval while preserving the W004 request, ordering, partial-success, and common wrapper contract.

Unchanged after recheck:

- `schema/record-model.md`: current active-index model only; legacy compatibility does not use it.
- `schema/record-source.md`: current source model only; no legacy normalization contract is added.
- `schema/fields.md`: current normalized field vocabulary only; legacy retrieval requires only derived `kind` plus optional parsed title/status.
- `schema/metadata-grammar.md`: remains current-only with no legacy exception.
- `schema/overview.md`: no dedicated legacy schema was created, so no parent Topics synchronization is required.
- `schema/discovery.md` and `schema/id-normalization.md`: current discovery and identity remain unchanged.
- `tools/list-records.md`: legacy records remain excluded from normal listing.
- tool overview, MVP scope, and responsibility boundary: no T03-specific pointer correction is required; broader recheck remains T04-owned.

### Scoped validation command

Run from the repository root:

```powershell
python -X utf8 product/src/tools/validate_spec.py `
  drmcp/records/spec/design-records-mcp/namespace-scanning.md `
  drmcp/records/spec/design-records-mcp/resolver.md `
  drmcp/records/spec/design-records-mcp/tools/resolve-reference.md `
  drmcp/records/spec/design-records-mcp/tools/get-records.md `
  --strict `
  --no-color
```

Initial execution on 2026-06-27 from the repository root returned:

```text
[strict]  All 4 file(s) OK.
```

This PASS predates the independent-review corrections below.
The same scoped command must be rerun after correction.

### Initial independent review and correction disposition

Initial verdict: `NEEDS REVISION`.

- `F-MAJ-01`: corrected. `get_records` now classifies accepted legacy-family matches before current exact families, queries only the legacy lookup map, and performs no second lookup or resolver invocation.
- `F-MAJ-02`: corrected. `namespace-scanning.md` now defines exact family-specific issued-ID grammar, token widths, domain grammar, slug grammar, and deterministic suffix boundaries.
- `F-MIN-01`: corrected. Legacy enumeration now rejects filesystem-alias files, does not traverse alias directories, requires canonical containment, and excludes only the violating source.

Finding closure remains subject to independent re-review.

### Post-correction scoped validation

The same four-file scoped command was rerun on 2026-06-27 after the review corrections.

Result:

```text
[strict]  All 4 file(s) OK.
```

Post-correction scoped validation verdict: `PASS`.

### Independent re-review result

Final verdict: `PASS`.

Previous-finding disposition:

- `F-MAJ-01`: `CLOSED`.
- `F-MAJ-02`: `CLOSED`.
- `F-MIN-01`: `CLOSED`.

The re-review reported:

- blocking findings: none;
- major findings: none;
- minor findings: none;
- advisories: none;
- regression findings: none;
- ownership-boundary violations: none;
- T03 closure readiness: `ready`.

The reviewer confirmed that:

- `get_records` uses deterministic legacy-first exact classification for overlap inputs while preserving W004 ordering, partial-success, common-wrapper, and body-inclusion contracts;
- the PRODUCT-owned accepted legacy-family set and DRMCP-owned lexical parser mapping remain separate authorities;
- the issued-ID and optional-slug boundary is deterministic for every accepted family;
- legacy enumeration excludes filesystem aliases and canonically escaped sources without invalidating unrelated entries;
- current-first resolver behavior, three-status outcome mapping, minimal legacy target projection, source-first body retrieval, and W003/W004/W006 ownership boundaries remain intact.

The reviewer did not rerun the validator because repo-local command execution was unavailable and assessed the recorded post-correction PASS by static inspection.

### Final state

- Normative DRMCP spec changes: four files listed above.
- Initial scoped validation before review corrections: `PASS` — `[strict]  All 4 file(s) OK.`
- Post-correction scoped validation: `PASS` — `[strict]  All 4 file(s) OK.`
- Initial independent review: `NEEDS REVISION`.
- Independent re-review: `PASS`.
- All findings are closed.
- T03 closure readiness: `ready`.
- T03 status changed to `done` on 2026-06-27.
