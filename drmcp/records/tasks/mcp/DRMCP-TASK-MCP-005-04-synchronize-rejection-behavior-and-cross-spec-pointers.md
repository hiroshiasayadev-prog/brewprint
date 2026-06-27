# DRMCP-TASK-MCP-005-04: Synchronize rejection behavior and cross-spec pointers

- **id**: DRMCP-TASK-MCP-005-04
- **status**: done
- **date**: 2026-06-27
- **work_item**: DRMCP-WORK-MCP-005
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 1d
- **depends_on**:
  - DRMCP-TASK-MCP-005-03
- **outputs**:
  - spec:drmcp.design_records_mcp.namespace_scanning
  - spec:drmcp.design_records_mcp.resolver
  - spec:drmcp.design_records_mcp.tools.resolve_reference
  - spec:drmcp.design_records_mcp.tools.get_records
  - DRMCP-WORK-MCP-005

## Goal

Synchronize rejected-input behavior across resolver and exact-retrieval contracts.

Keep current-first resolver behavior, legacy-first exact-retrieval classification, configured legacy lookup, and W003-W006 ownership boundaries unchanged.

## Work

- Confirm the completed T01-T03 authority and accepted decisions.
- Inventory candidate specs before normative edits.
- Record changed and unchanged dispositions for every candidate.
- Define one rejected-input matrix for resolver and exact retrieval.
- Distinguish string-item failure from request-shape failure.
- Preserve resolver current-first orchestration and `get_records` operation-specific classification.
- Synchronize legacy lexical-mapping pointers without extending the PRODUCT-owned accepted family set.
- Remove provisional T04 wording from normative contracts.
- Keep warning and diagnostic taxonomy delegated to W006.
- Run scoped strict validation against changed normative specs.
- Prepare an independent review prompt after validator PASS.

## Done condition

- Rejected current, legacy, alias, path, and repair-dependent inputs have explicit operation-specific behavior.
- `V01-SPEC-*` remains distinct from accepted V01 sequential families.
- Current `spec:` input classification uses lexical grammar only; section targets, headings, aliases, and physical paths do not become lookup surfaces.
- Resolver string inputs accepted by neither current nor accepted legacy grammar return `unsupported`.
- Accepted current or legacy resolver inputs with no selectable target remain `unresolved`.
- `get_records` rejected string items remain per-item warning triggers and do not become request-shape errors.
- `get_records` retains legacy-first exact classification, one lookup scope, and no resolver invocation.
- Positive current and accepted legacy contracts remain unchanged.
- Overview, MVP, responsibility, listing, and schema candidates change only when a pointer is stale.
- W003 discovery, W004 request and response, and W006 diagnostics remain outside this Task.
- Every changed normative spec passes scoped strict validation.
- Independent review reports no blocking, major, or minor findings before this Task is marked `done`.

## Verification

- Compare current sequential and spec inputs with PRODUCT identity authorities.
- Compare accepted legacy families with Brewprint compatibility authority.
- Compare resolver order and outcomes with T02 and T03.
- Compare `get_records` classification and partial-success behavior with final W004 and T03.
- Confirm each rejected input is neither repaired nor redirected to a second lookup scope.
- Confirm string-item rejection is not promoted to `invalid_request`.
- Confirm overview and ownership pointers do not cite retired `get_record` or stale listing behavior.
- Confirm no archive schema, normalized legacy record, or legacy spec-ID contract is introduced.
- Run the scoped strict validator against only changed normative specs.
- Run an independent review before changing status to `done`.

## Evidence

### Baseline

- `DRMCP-TASK-MCP-005-01`, `DRMCP-TASK-MCP-005-02`, and `DRMCP-TASK-MCP-005-03` are complete.
- T03 post-correction scoped validation passed: `[strict]  All 4 file(s) OK.`
- T03 independent re-review verdict was `PASS`.
- T03 findings `F-MAJ-01`, `F-MAJ-02`, and `F-MIN-01` are closed.
- No T03 blocking, major, minor, advisory, regression, or ownership-boundary finding remains.
- T04 does not reopen current-first resolution, configured legacy fallback, the minimal legacy lookup model, or exact-retrieval classification.

### Positive contract preserved

| input | `resolve_reference` | `get_records` |
|---|---|---|
| Complete current app-aware sequential ID | Current grammar first; exact active-index lookup. | Current sequential classification; active index only. |
| Current path-derived document-level `spec:` ref | Current grammar first; exact active-index lookup. | Current spec classification; active index only. |
| Exact accepted V01 sequential ID | Evaluated only after the current stage remains unresolved; configured legacy lookup map only. | Classified as legacy before current exact families; configured legacy lookup map only. |

One string may match both current and accepted legacy grammar.
Resolver behavior remains active-index-first with unresolved-only fallback.
`get_records` behavior remains legacy-first exact classification with one lookup scope.

### Rejected-input matrix

Every row assumes the request shape is valid and the operation receives a string item exactly as supplied.

| input class | `resolve_reference` outcome | `get_records` item behavior | request-shape error | validation boundary |
|---|---|---|---|---|
| `V01-SPEC-*` | `unsupported` | Malformed or unsupported item warning trigger. | No. | No active spec compatibility target. |
| App-prefixless sequential ID | `unsupported` | Malformed or unsupported item warning trigger. | No. | Noncanonical when present in declared current relations. |
| Physical path | `unsupported` | Malformed or unsupported item warning trigger. | No. | May be reported as a noncanonical declared relation value; not a read lookup key. |
| Fuzzy, partial, or completed-by-inference reference | `unsupported` | Malformed or unsupported item warning trigger. | No. | No repair or inference during validation lookup. |
| Legacy YAML-only alias spelling that fails current canonical grammar | `unsupported` | Malformed or unsupported item warning trigger. | No. | No alias-registry lookup or repair. |
| Direct `yaml:` input | `unsupported` | Malformed or unsupported item warning trigger. | No. | No active lookup class. |
| `fixture:` input | `unsupported` | Malformed or unsupported item warning trigger. | No. | Fixture lookup is outside current operations. |
| `internal-design:` input | `unsupported` | Malformed or unsupported item warning trigger. | No. | No active lookup class. |
| `coverage:` input | `unsupported` | Malformed or unsupported item warning trigger. | No. | No active lookup class. |
| `COV-*` input | `unsupported` | Malformed or unsupported item warning trigger. | No. | No active lookup class. |
| Section-like `spec:` value that does not match current spec grammar | `unsupported` | Malformed or unsupported item warning trigger. | No. | No repair, section-target lookup, heading lookup, or alias lookup. |
| Metadata-only alias spelling that fails current canonical grammar | `unsupported` | Malformed or unsupported item warning trigger. | No. | A referring metadata value does not register a target alias. |
| Value requiring whitespace, case, prefix, domain, or sequence repair | `unsupported` | Malformed or unsupported item warning trigger. | No. | Exact canonical lookup only; no repair. |
| Empty string | `unsupported` | Malformed item warning trigger. | No. | Not a canonical input. |

Current spec classification is lexical and does not use semantic origin.
A `spec:` string matching current spec grammar queries only the active index, including a string previously used as a section alias.
A missing current target remains resolver `unresolved` and a `get_records` unresolved-current warning trigger.
Front-matter alias data, section metadata, and headings are not consulted for classification or lookup.

The matrix defines operation outcomes and warning triggers only.
W006 owns warning categories, severity, messages, shared fields, source locations, and exceptional path representation.

### Candidate-file disposition before normative edits

| candidate | current observation | required disposition |
|---|---|---|
| `overview.md` | Current/legacy source separation and Topics pointers are current. | Unchanged unless a later normative edit creates a new navigation need. |
| `namespace-scanning.md` | Accepted-family lexical mapping is correct; operation pointers are incomplete. | Add pointer-only clarification from lexical mapping to resolver and exact-retrieval contracts. |
| `resolver.md` | Positive orchestration is correct; the complete rejected-input boundary is not summarized. | Add a rejected-input boundary and operation pointer without changing outcomes. |
| `mvp-scope.md` | W004/W005/W006 operation split is current. | Unchanged. |
| `responsibility-boundary.md` | W003-W006 ownership is current and non-overlapping. | Unchanged. |
| `tools/overview.md` | Navigation-first catalog and owner pointers are current. | Unchanged. |
| `tools/resolve-reference.md` | Positive contract is correct; rejected inputs are listed under provisional T04 wording. | Replace provisional wording with the final operation-specific rejection contract. |
| `tools/get-records.md` | Exactness rules are correct; rejected string-item classes are not enumerated completely. | Add explicit rejected-item behavior without changing W004 request, ordering, wrapper, or warning ownership. |
| `tools/list-records.md` | Explicitly excludes exact lookup, resolver, spec listing, and legacy listing. | Unchanged. |
| `schema/discovery.md` | Current-only source and no-legacy-alias boundaries are correct. | Unchanged. |
| `schema/id-normalization.md` | Current identity and no-repair boundary are correct; legacy remains separate. | Unchanged. |
| `schema/record-model.md` | Current active-index model only; no legacy normalization. | Unchanged. |
| `schema/record-source.md` | Current source model only; public projection delegated to W004. | Unchanged. |

### Ownership confirmation

- W003 remains the owner of current discovery, parsing, identity, addressability, and active-index construction.
- W004 remains the owner of `list_records`, `get_records` request shape, ordering, deduplication, partial success, response wrapper, and body inclusion.
- W005 owns resolver behavior, configured legacy lookup, legacy lexical parser application, and rejected-input operation behavior.
- W006 remains the owner of warning and diagnostic schema, category, severity, message, shared fields, source location, validation execution, and exceptional path exposure.

### Normative reflection

Changed:

- `spec:drmcp.design_records_mcp.namespace_scanning`: Clarified the exact lexical-mapping boundary and added operation pointers.
- `spec:drmcp.design_records_mcp.resolver`: Added the complete rejected-input boundary without changing orchestration or outcomes.
- `spec:drmcp.design_records_mcp.tools.resolve_reference`: Replaced provisional T04 wording with the final rejection table.
- `spec:drmcp.design_records_mcp.tools.get_records`: Added the rejected string-item table without changing W004 batch behavior.

Unchanged after recheck:

- `spec:drmcp.design_records_mcp.overview`;
- `spec:drmcp.design_records_mcp.mvp_scope`;
- `spec:drmcp.design_records_mcp.responsibility_boundary`;
- `spec:drmcp.design_records_mcp.tools.overview`;
- `spec:drmcp.design_records_mcp.tools.list_records`;
- `spec:drmcp.design_records_mcp.schema.discovery`;
- `spec:drmcp.design_records_mcp.schema.id_normalization`;
- `spec:drmcp.design_records_mcp.schema.record_model`;
- `spec:drmcp.design_records_mcp.schema.record_source`.

Each unchanged candidate already carries the current operation or ownership pointer.
No new navigation child, legacy schema, normalized legacy record, or listing behavior was introduced.

### Static verification

- The four changed normative specs were read back after editing.
- Resolver current-first order and configured legacy fallback remain unchanged.
- `get_records` retains legacy-first exact classification for accepted legacy overlap inputs.
- Accepted V01 sequential families remain positive inputs; `V01-SPEC-*` remains excluded.
- Current `spec:` inputs are classified by lexical grammar; an exact lookup miss remains unresolved.
- Section targets, headings, aliases, and physical paths are not separate lookup surfaces.
- No rejected input is repaired or redirected to a second lookup scope.
- No rejected string item is promoted to a request-shape error.
- No W006-owned diagnostic representation was added.

### Scoped validation

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

Pre-review execution result on 2026-06-28:

```text
[strict]  All 4 file(s) OK.
```

The pre-review scoped strict validation verdict was `PASS`.
This result predates the F-MAJ-01 semantic-classification correction.

Post-correction execution result on 2026-06-28:

```text
[strict]  All 4 file(s) OK.
```

Post-correction scoped strict validation verdict: `PASS`.

### Review state

- Initial independent review verdict: `NEEDS REVISION`.
- `F-MAJ-01` identified semantic-origin classification of section-level `spec:` strings as incompatible with lexical current-grammar classification.
- Correction applied: every current-grammar-matching `spec:` string queries only the active index; an exact miss remains unresolved.
- Front-matter alias registries, section metadata, headings, and section targets remain outside lookup.
- Pre-correction scoped strict validation: `PASS` — `[strict]  All 4 file(s) OK.`
- Post-correction scoped strict validation: `PASS` — `[strict]  All 4 file(s) OK.`
- Independent re-review verdict: `PASS`.
- `F-MAJ-01`: `CLOSED`.
- Re-review reported no blocking, major, or minor findings.
- Advisories A-01 and A-02 remain non-blocking and require no T04 correction.
- Lexical-versus-semantic classification, resolver outcomes, `get_records` item classification, positive contracts, ownership boundaries, validation evidence, and regression assessment were all accepted.
- Closure readiness: `READY`.
- T04 status changed to `done` on 2026-06-28.
