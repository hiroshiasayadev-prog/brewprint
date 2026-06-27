# DRMCP-TASK-MCP-005-05: Validate, review, and close resolver and legacy-fallback contract

- **id**: DRMCP-TASK-MCP-005-05
- **status**: done
- **date**: 2026-06-28
- **work_item**: DRMCP-WORK-MCP-005
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 1d
- **depends_on**:
  - DRMCP-TASK-MCP-005-04
- **outputs**:
  - spec:drmcp.design_records_mcp.namespace_scanning
  - spec:drmcp.design_records_mcp.resolver
  - spec:drmcp.design_records_mcp.tools.resolve_reference
  - spec:drmcp.design_records_mcp.tools.get_records
  - DRMCP-WORK-MCP-005
  - DRMCP-TASK-MCP-001-05

## Goal

Validate the final W005 resolver and configured legacy-fallback contract.

Run an independent final review and close W005 only after validation passes and no blocking, major, or minor finding remains.

## Work

- Confirm the final normative changed-file manifest from T01 through T04.
- Recheck the unchanged root, tool, schema, diagnostics, validation, and authoring-transaction candidates.
- Confirm the current-first resolver contract without redesigning accepted T01 through T04 decisions.
- Confirm the operation-specific `get_records` exact-classification contract.
- Confirm configured legacy-root validation, filename-derived issued identity, exact lookup-map construction, and source-first legacy retrieval.
- Confirm the rejected-input matrix and lexical current-spec classification.
- Confirm W003, W004, W006, and PRODUCT ownership boundaries.
- Run scoped strict validation against the four final normative specs.
- Record the exact validator result only after repository-local execution.
- Prepare and run an independent final review after validator PASS.
- Record every previous finding disposition and any new finding.
- Apply only the minimum required correction when a finding is reported.
- Rerun scoped validation and independent review after any normative correction.
- Set this Task, W005, and `DRMCP-TASK-MCP-001-05` to `done` only after closure conditions are satisfied.
- Synchronize `DRMCP-WORK-MCP-001` Evidence only when the accepted child-gate evidence requires it.

This Task does not implement runtime behavior, create fixtures, design W006 diagnostics, modify authoring transactions, or reopen PRODUCT compatibility authority.

## Done condition

- The final normative changed-file manifest contains exactly the four W005 specs recorded in Evidence.
- Every unchanged candidate is rechecked and has no new contradiction requiring a W005 correction.
- Current canonical grammar and active-index lookup run before resolver legacy eligibility.
- A resolved current target stops resolver legacy evaluation and lookup.
- Accepted legacy resolver inputs query only the configured legacy lookup map after the current stage remains unresolved.
- Accepted current or legacy resolver lookup failure returns `unresolved`.
- Only an input accepted by neither current nor accepted legacy grammar returns `unsupported`.
- `get_records` retains legacy-first exact classification, one lookup scope, and no resolver invocation.
- Exact ordered deduplication, first-occurrence ordering, partial success, successful-record-only response, top-level warnings, and request-wide `include_body` remain unchanged.
- Missing or empty `legacy_roots` disables fallback without affecting current-only operation.
- Every configured legacy root remains mandatory and no partial legacy-root acceptance occurs.
- Filename-derived issued identity, exact case-sensitive lookup, duplicate conflict behavior, filesystem-alias exclusion, and source-first retrieval remain explicit.
- Successful legacy resolver targets expose only `target_type`, `ref`, and `kind` and preserve the issued legacy ID.
- Rejected inputs are not repaired, inferred, redirected, or looked up through aliases, headings, sections, paths, fixtures, or obsolete prefixes.
- Every current-grammar-matching `spec:` value is classified lexically as current and queries only the active index.
- W003, W004, W006, and PRODUCT ownership boundaries remain non-overlapping.
- The four final normative specs pass scoped strict validation.
- Independent final review returns `PASS`.
- Independent final review reports no blocking, major, or minor finding.
- Every previous and newly reported finding is `CLOSED`.
- This Task, W005, and `DRMCP-TASK-MCP-001-05` contain the accepted final evidence before status changes to `done`.

## Verification

- Compare the final contract with `DRMCP-ADR-MCP-001`, `DRMCP-REQ-MCP-001`, and `DRMCP-INV-MCP-002`.
- Compare accepted current and legacy input semantics with PRODUCT identity, traceability, and compatibility authorities.
- Compare current discovery, parsing, identity, addressability, and active-index assumptions with final W003 contracts.
- Compare exact retrieval request, ordering, deduplication, partial success, wrapper, and body behavior with final W004 contracts.
- Compare warning, diagnostic, validation, source-location, and exceptional path claims with the W006 ownership boundary.
- Run the scoped validator against only the four final normative specs.
- Run an independent final review after validator PASS.
- Confirm that no completion status is recorded before both checks pass.

## Evidence

### Upstream completion baseline

- `DRMCP-TASK-MCP-005-01` is `done`.
  - Authority, affected-file, operation-split, and ownership baseline accepted.
  - Final independent re-review verdict: `PASS`.
  - Findings `F-MAJ-01`, `F-MAJ-02`, `F-MIN-01`, `F-MAJ-03`, and `F-MIN-02` are closed.
- `DRMCP-TASK-MCP-005-02` is `done`.
  - Current-first grammar and active-index resolution contract accepted.
  - Post-correction validator result: `[strict]  All 2 file(s) OK.`
  - Final independent re-review verdict: `PASS`.
  - Finding `F-MAJ-01` is closed.
- `DRMCP-TASK-MCP-005-03` is `done`.
  - Configured legacy roots, lexical issued-ID mapping, minimal exact lookup map, duplicate behavior, and source-first retrieval accepted.
  - Post-correction validator result: `[strict]  All 4 file(s) OK.`
  - Final independent re-review verdict: `PASS`.
  - Findings `F-MAJ-01`, `F-MAJ-02`, and `F-MIN-01` are closed.
- `DRMCP-TASK-MCP-005-04` is `done`.
  - Rejected-input behavior and lexical current-spec classification accepted.
  - Post-correction validator result: `[strict]  All 4 file(s) OK.`
  - Final independent re-review verdict: `PASS`.
  - Finding `F-MAJ-01` is closed.
  - Advisories A-01 and A-02 are non-blocking and remain outside T05 scope unless a new contradiction is found.

### Final normative changed-file manifest

T01 through T04 produce one final W005 normative scope:

| spec | final W005 responsibility |
|---|---|
| `spec:drmcp.design_records_mcp.namespace_scanning` | `legacy_roots`, strict root validation, lexical issued-ID mapping, enumeration boundary, and separate exact lookup-map construction. |
| `spec:drmcp.design_records_mcp.resolver` | Current-first resolver orchestration, configured legacy fallback, rejected-input boundary, operation separation, and successful target boundaries. |
| `spec:drmcp.design_records_mcp.tools.resolve_reference` | Public request, three-status outcome, current and legacy target projection, and operation-specific rejected-input behavior. |
| `spec:drmcp.design_records_mcp.tools.get_records` | Legacy-first exact classification, one lookup scope, rejected string-item behavior, and minimal legacy source projection while preserving W004 batch semantics. |

No fifth normative W005 file is required by the accepted T01 through T04 decisions.

### Rechecked and unchanged manifest

The following candidates remain unchanged unless final review finds a new contradiction:

- root overview;
- tool overview;
- MVP scope;
- responsibility boundary;
- `tools/list-records.md`;
- current discovery, record-model, record-source, fields, metadata-grammar, ID-normalization, and schema-overview specs;
- diagnostics and validation specs;
- authoring transaction specs.

Reasons:

- Root and tool navigation already point to the W003 through W006 operation split.
- W004 owns compact listing, exact retrieval request and response behavior, normal path hiding, and common successful-record projection.
- W005 uses a minimal separate legacy lookup map and does not add a normalized legacy schema or current record-model extension.
- W006 owns warning and diagnostic schema, validation execution, source locations, and exceptional path exposure.
- Authoring transaction behavior is outside `DRMCP-REQ-MCP-001` W005 scope.
- A-01 and A-02 do not create a resolver or configured-fallback contradiction.

### Final resolver contract baseline

- Evaluate current canonical grammar first.
- Query the active index first for a current input.
- Stop after one resolved current target.
- Evaluate accepted legacy grammar only when the current stage has no resolved target.
- Query only the configured legacy lookup map after an accepted legacy match.
- Return `unresolved` for accepted current or legacy lookup failure.
- Return `unsupported` only when neither current nor accepted legacy grammar accepts the input.
- Never rewrite a current input into a legacy input.
- Do not perform alias lookup, repair, path inference, section lookup, or heading lookup.

Current spec classification is lexical.
Every string matching current spec grammar is a current input regardless of semantic origin.
An exact active-index miss remains `unresolved`.

### Final exact-retrieval contract baseline

- Evaluate exact accepted legacy grammar before current exact families.
- Query only the configured legacy lookup map for a legacy overlap input.
- Query only the active index for a current spec input.
- Query only the active index for a current sequential input.
- Do not invoke `resolve_reference`.
- Do not run a second lookup after failure.
- Deduplicate by exact string equality.
- Preserve first-occurrence request order.
- Preserve partial success.
- Return successful records only in `records`.
- Return operation warnings at the top level.
- Apply `include_body` to the entire request.

### Final configured legacy lookup baseline

- Missing and empty `legacy_roots` are equivalent and disable fallback.
- Every configured legacy root is mandatory.
- One invalid configured legacy root fails startup and prevents partial lookup-map construction.
- Current and legacy roots remain separate and non-overlapping.
- Issued legacy identity is derived from exact family-specific filename grammar.
- Exact lookup is case-sensitive.
- Duplicate issued IDs have no selected winner.
- Filesystem aliases are not candidates and alias directories are not traversed.
- Legacy sources are not normalized into the current record model.
- No dedicated legacy schema is introduced.
- Legacy sources remain excluded from normal listing, current repository validation, and authoring targets.
- A successful legacy resolver target contains only `target_type`, `ref`, and `kind`.
- A successful legacy retrieval preserves the issued ID and permits readable-source retrieval without current metadata conformance.

### Final rejected-input baseline

The following remain rejected when they do not independently match a current canonical grammar or accepted legacy issued-ID grammar:

- `V01-SPEC-*`;
- app-prefixless sequential IDs;
- physical paths;
- fuzzy, partial, or inferred refs;
- direct `yaml:` inputs;
- `fixture:` inputs;
- `internal-design:` inputs;
- `coverage:` inputs;
- `COV-*` inputs;
- legacy YAML-only alias spellings;
- metadata-only alias spellings;
- section-like `spec:` values that fail current spec grammar;
- values requiring whitespace, case, prefix, domain, or sequence repair;
- empty strings.

Resolver behavior is `unsupported` for these string values.
`get_records` treats them as per-item malformed or unsupported warning triggers, not request-shape failures.

### Ownership confirmation

| owner | retained responsibility |
|---|---|
| W003 | Current discovery, parsing, canonical identity, addressability, duplicate handling, and active-index construction. |
| W004 | `get_records` request shape, ordering, exact deduplication, partial success, common wrapper, body inclusion, normal successful-record projection, and normal path hiding. |
| W005 | Resolver order and outcomes, configured legacy roots and lookup map, legacy lexical parser application, minimal legacy projections, and operation-specific rejected-input behavior. |
| W006 | Warning and diagnostic schema, category, severity, message, shared fields, ref association, source location, validation execution, and exceptional path representation. |
| PRODUCT | Current canonical grammar and identity, accepted legacy-family semantics, and canonical traceability semantics. |

### Final scoped validation

Repository-local command:

```powershell
python -X utf8 product/src/tools/validate_spec.py `
  drmcp/records/spec/design-records-mcp/namespace-scanning.md `
  drmcp/records/spec/design-records-mcp/resolver.md `
  drmcp/records/spec/design-records-mcp/tools/resolve-reference.md `
  drmcp/records/spec/design-records-mcp/tools/get-records.md `
  --strict `
  --no-color
```

User-reported execution result on 2026-06-28:

```text
[strict]  All 4 file(s) OK.
```

The recorded result covered exactly the final W005 normative changed-file manifest before the final-review correction below.

### Independent final review finding

Independent final review verdict on 2026-06-28: `NEEDS REVISION`.

`F-MIN-FINAL-01` remains open pending validation and limited re-review.

Finding:

- `resolver.md`, `tools/resolve-reference.md`, and `tools/get-records.md` retained `date: 2026-06-27` after the substantive lexical current-spec classification correction performed on 2026-06-28.
- The stale dates conflicted with the spec-authoring requirement to record the latest substantive contract-change date.
- The required minimal correction is metadata-only and does not change contract behavior or expand the four-file normative manifest.

Correction applied on 2026-06-28:

- `resolver.md`: `date` changed from `2026-06-27` to `2026-06-28`.
- `tools/resolve-reference.md`: `date` changed from `2026-06-27` to `2026-06-28`.
- `tools/get-records.md`: `date` changed from `2026-06-27` to `2026-06-28`.
- `namespace-scanning.md` required no correction.
- No contract body changed.

The previous final validator PASS predates this normative metadata correction and is superseded for closure purposes.

Post-correction repository-local validation was rerun on 2026-06-28 with the same four-file scope.

```text
[strict]  All 4 file(s) OK.
```

Post-correction scoped validation verdict: `PASS`.

Repository-local whitespace verification was also run against the three corrected specs, this Task, and W005.
`git diff --check` reported no whitespace error.
It emitted LF-to-CRLF working-copy conversion warnings for four files; these are non-blocking line-ending notices and not `git diff --check` failures.

A limited independent re-review must now confirm the three corrected dates, the post-correction validator result, and `F-MIN-FINAL-01` closure.

### Final review and closure state

- Initial independent final review verdict: `NEEDS REVISION`.
- Reported finding: `F-MIN-FINAL-01`.
- The finding was corrected by updating the three affected spec dates to `2026-06-28` without changing contract body text.
- Post-correction final scoped validation: `PASS` — `[strict]  All 4 file(s) OK.`
- `git diff --check`: no whitespace error; LF-to-CRLF warnings only.
- Limited independent re-review verdict: `PASS`.
- `F-MIN-FINAL-01`: `CLOSED`.
- Blocking findings: none.
- Major findings: none.
- Minor findings: none.
- Advisories A-01 and A-02 remain non-blocking and do not reopen W005 scope.
- Final normative manifest remains `namespace-scanning.md`, `resolver.md`, `tools/resolve-reference.md`, and `tools/get-records.md`.
- T05 closure readiness: `READY`.
- W005 closure readiness: `READY`.
- `DRMCP-TASK-MCP-001-05` closure readiness: `READY`.
- T05 status changed to `done` on 2026-06-28 after accepted limited re-review PASS.
