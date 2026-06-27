# DRMCP-TASK-MCP-005-02: Correct current-first grammar and active-index resolution contract

- **id**: DRMCP-TASK-MCP-005-02
- **status**: done
- **date**: 2026-06-27
- **work_item**: DRMCP-WORK-MCP-005
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 1d
- **depends_on**:
  - DRMCP-TASK-MCP-005-01
- **outputs**:
  - spec:drmcp.design_records_mcp.resolver
  - spec:drmcp.design_records_mcp.tools.resolve_reference
  - DRMCP-WORK-MCP-005

## Goal

Reflect the accepted current-first resolver sequence into the authoritative resolver contracts.

Define active-index lookup, current-resolution stopping, unresolved-only legacy eligibility, and current successful non-path target projection without deciding legacy archive mechanics.

## Work

- Replace stale resolver pointers to broad listing and retired single-record retrieval.
- Point current sequential grammar to the PRODUCT app-aware artifact-ID authority.
- Point current spec grammar to the PRODUCT path-derived document-level `spec:` authority.
- Remove front-matter `semantic_refs`, section-level spec refs, and V01-era current grammar assumptions.
- Define exact active-index lookup for accepted current inputs.
- Stop resolution immediately when one current target resolves.
- Enter legacy fallback eligibility only after the current stage remains unresolved.
- Preserve current-first ordering when current and accepted legacy grammar overlap.
- Keep `resolve_reference` separate from `get_records`.
- Define the current successful target discriminator, canonical identity, kind, title, and status fields.
- Exclude normal physical paths from successful targets.
- Delegate diagnostic, warning, source-location, validation, and exceptional path representation to W006.
- Leave `legacy_roots`, archive parsing, archive indexing, duplicate legacy IDs, root failures, and legacy target projection to T03.
- Run scoped strict validation for the two changed normative specs.
- Prepare an independent review prompt without closing this Task.

## Done condition

- `resolver.md` defines current canonical grammar evaluation before legacy grammar evaluation.
- Accepted current grammar points to complete app-aware sequential IDs and path-derived document-level `spec:` refs.
- Active-index lookup is explicit.
- A resolved current target stops all legacy evaluation and lookup.
- Legacy eligibility begins only after the current stage has no resolved target.
- Current and accepted legacy grammar overlap preserves active-index-first order.
- Current inputs are not rewritten as legacy inputs.
- `resolve_reference` and `get_records` remain separate operations.
- Current spec and current sequential targets have distinct discriminators.
- Current successful targets expose canonical identity, kind, title, and status without normal path fields.
- Diagnostic and path representation remain delegated to W006.
- T03-owned legacy mechanics and projection are not defined.
- Both changed normative specs pass scoped strict validation.
- Independent review reports no blocking, major, or minor findings before this Task is marked `done`.

## Verification

- Compare current sequential input grammar with `spec:product.design_records.namespace_model.artifact_id_grammar`.
- Compare current spec input grammar with `spec:product.design_records.spec_format.spec_id_as_ref`.
- Compare lookup sources and inactive section/front-matter refs with `spec:product.design_records.traceability.resolve_and_validation`.
- Compare the current-first sequence with `DRMCP-ADR-MCP-001` and `DRMCP-REQ-MCP-001`.
- Compare active-index and addressability behavior with final W003 contracts.
- Compare operation separation with final W004 `get_records` contract.
- Confirm that diagnostic and exceptional path details remain outside the changed specs.
- Run the scoped validator against only `resolver.md` and `tools/resolve-reference.md`.
- Run an independent review before changing status to `done`.

## Evidence

### Changed normative files

- `spec:drmcp.design_records_mcp.resolver`.
- `spec:drmcp.design_records_mcp.tools.resolve_reference`.

Both normative specs received substantive contract changes and use date `2026-06-27`.

### Accepted authority

| concern | authority or accepted input |
|---|---|
| Current sequential canonical grammar | `spec:product.design_records.namespace_model.artifact_id_grammar` |
| Current spec canonical ref | `spec:product.design_records.spec_format.spec_id_as_ref` |
| Canonical lookup sources and section-ref exclusion | `spec:product.design_records.traceability.resolve_and_validation` |
| Current identity and addressability | Final W003 contracts. |
| Exact retrieval boundary | Final W004 contracts. |
| Current-first and configured fallback order | `DRMCP-ADR-MCP-001` and `DRMCP-REQ-MCP-001`. |
| Accepted legacy families | `spec:product.brewprint.compatibility.legacy_id_compatibility` |
| Diagnostic and path representation | `DRMCP-WORK-MCP-006` |

### Current-first sequence

1. Evaluate the input exactly against current canonical grammar.
2. Query the active index when current grammar accepts the input.
3. Return one addressable current target and stop when current resolution succeeds.
4. Proceed to legacy fallback eligibility only when the current stage has no resolved target.
5. Preserve active-index-first order when one string can match current and accepted legacy grammar.
6. Never rewrite a current input into a legacy input.

The current stage is unresolved when no current grammar accepts the value or accepted current grammar produces no addressable active-index target.
`unsupported` applies only when neither current grammar nor accepted legacy-family grammar accepts the value.
Accepted legacy-family grammar remains accepted regardless of `legacy_roots` availability.
The exact legacy grammar reflection, archive lookup mechanics, and final status mapping for accepted legacy inputs remain T03-owned.

### Current successful target projection

| field | contract |
|---|---|
| `target_type` | Required. `current_spec` or `current_sequential_record`. |
| `ref` | Required resolved canonical current identity. |
| `kind` | Required current record kind. |
| `title` | Required string-or-null projection of the parsed title. |
| `status` | Required string-or-null projection of the parsed lifecycle status. |

Parsed invalid values remain unchanged.
`null` represents unavailable parsed title or status on an invalid but addressable current source.
No physical path, source location, source provenance, section anchor, index path, or resolver trace appears in the normal successful target.

### W004 boundary

- `get_records` remains the sole exact-retrieval operation.
- `resolve_reference` does not call `get_records`.
- `get_records` does not call `resolve_reference`.
- T02 does not change the `get_records` request, response, warning, heading, body, or successful-record projection.
- T02 does not add resolver fallback behavior to exact retrieval.

### W006 boundary

T02 does not define:

- diagnostic object shape;
- warning schema;
- category names;
- severity;
- messages;
- source-location fields;
- validation execution;
- exceptional path representation.

The changed resolver contracts define outcome conditions and successful non-path targets only.

### T03 remainder

T03 still owns:

- concrete `legacy_roots` entry fields;
- missing, unreadable, duplicate, or overlapping legacy-root behavior;
- legacy source parsing;
- archive-record construction;
- archive-index entry and duplicate-issued-ID behavior;
- normalized legacy field availability;
- exact accepted legacy grammar reflection;
- configured legacy-index lookup;
- issued legacy ID preservation in the legacy result;
- legacy successful-target projection;
- shared-schema extension versus dedicated legacy schema selection;
- conditional `schema/overview.md` synchronization;
- final public status mapping when accepted legacy grammar matches but fallback is disabled, unavailable, unresolved, or conflicted.

T02 defines only the gate into legacy fallback and does not pre-decide these mechanics.

### Independent review correction

- Initial independent review verdict: `NEEDS REVISION`.
- Blocking findings: none.
- Major finding `F-MAJ-01`: `tools/resolve-reference.md` coupled accepted legacy grammar with configured fallback availability and left accepted legacy inputs without a stable classification when fallback was disabled, unavailable, or conflicted.
- Correction applied on 2026-06-27:
  - `unsupported` is now limited to values accepted by neither current grammar nor accepted legacy-family grammar;
  - accepted legacy-family grammar is explicitly independent of `legacy_roots` configuration and usability;
  - accepted legacy grammar after an unresolved current stage proceeds to the T03-owned configured fallback contract;
  - final status mapping for disabled fallback, unavailable roots or index, unresolved legacy targets, duplicate issued IDs, and resolved legacy targets remains T03-owned;
  - `resolver.md` received the matching orchestration synchronization.
- F-MAJ-01 disposition: `CLOSED` by independent re-review.

### Static verification

- Changed files were read back after writing.
- The T02 Task has one H1 and all canonical Task H2 sections.
- Both normative specs have one H1, H1-adjacent metadata, and date `2026-06-27`.
- Targeted stale-claim inspection found only explicit exclusions or delegated-boundary references for front-matter refs, section refs, retired tools, diagnostics, and legacy mechanics.
- No positive current contract remains for front-matter lookup, section-level spec targets, V01-era current grammar, retired `get_record`, broad `list_records`, or normal path-bearing targets.

### Validation

Scoped validator command:

```powershell
python -X utf8 product/src/tools/validate_spec.py `
  drmcp/records/spec/design-records-mcp/resolver.md `
  drmcp/records/spec/design-records-mcp/tools/resolve-reference.md `
  --strict `
  --no-color
```

First execution result on 2026-06-27: `FAIL`.

```text
drmcp\records\spec\design-records-mcp\tools\resolve-reference.md
  ERROR  [MISSING_SECTION]  Required section '## Response' is missing for Contract (interface) spec.

[strict]  1 file(s) with issues: 1 error(s)
```

Correction applied: renamed `## Response outcomes` to the required canonical `## Response` heading without changing the response contract content.

Post-correction execution result on 2026-06-27: `PASS`.

```text
[strict]  All 2 file(s) OK.
```

### Review state

- Initial independent review: `NEEDS REVISION`.
- Initial blocking findings: none.
- Initial major findings: `F-MAJ-01`.
- Initial minor findings: none.
- F-MAJ-01 correction: applied.
- Scoped strict validation: `PASS` for both changed normative specs.
- Independent re-review verdict: `PASS`.
- F-MAJ-01: `CLOSED`.
- Re-review blocking findings: none.
- Re-review major findings: none.
- Re-review minor findings: none.
- Re-review advisories: none.
- Independent reviewer did not rerun the validator because repository-local command execution was unavailable; the reviewer confirmed the recorded first failure, heading correction, recorded PASS, and presence of canonical `## Response` by static inspection.
- T02 closure readiness: `READY`.
- T02 status changed to `done` on 2026-06-27.
