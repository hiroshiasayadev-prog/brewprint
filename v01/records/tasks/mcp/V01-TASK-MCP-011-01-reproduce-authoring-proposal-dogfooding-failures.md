# V01-TASK-MCP-011-01: Reproduce authoring proposal dogfooding failures

- **id**: V01-TASK-MCP-011-01
- **status**: done
- **date**: 2026-06-02
- **work_item**: V01-WORK-MCP-011
- **source_requirement**: V01-REQ-MCP-011
- **estimate**: 0.5d
- **depends_on**:
- **outputs**:
  - Reproduction notes for authoring proposal input normalization failures
  - Reproduction notes for proposal validation scope pollution
  - Cause hypothesis and recommended split for V01-TASK-MCP-011-02

## Goal

Reproduce and classify the observed Design Records MCP authoring proposal failures before changing the spec or implementation.

This task should produce enough evidence for the follow-up spec update and implementation fix tasks to proceed without guessing.

## Work

- Run targeted repository status and diff checks to confirm the current working tree before investigation.
- Inspect the authoring transaction implementation path for `propose_record_create` and related validation calls.
- Reproduce or attempt to reproduce the observed dogfooding failures:
  - `body` and `fields` supplied together causing undocumented precedence or ignored fields;
  - `fields.id` being required even when top-level `id` is supplied;
  - lowercase `domain: mcp` failing against uppercase ID domain `MCP`;
  - propose-time validation returning unrelated repository diagnostics, especially diagnostics not reproduced by direct `validate_records` over the expected affected records.
- Compare propose-time diagnostics with direct `validate_records` behavior for the same kind and affected ID range where possible.
- Record whether each issue is a contract gap, implementation bug, stale documentation, or expected behavior with missing spec text.
- Do not modify implementation or spec files in this task unless the user explicitly expands the scope.

## Done condition

- Each observed failure has a reproduction result: reproduced / not reproduced / blocked, with command or JSON-RPC input and relevant output summarized.
- The likely implementation entrypoints for each failure are identified.
- The difference between propose-time validation and direct `validate_records` is characterized, or explicitly marked unresolved with the next command needed.
- A recommended scope for `V01-TASK-MCP-011-02` is provided.

## Verification

Expected checks include:

```powershell
git status --short
rg "propose_record_create|ProposeRecordCreate|validate" internal/designrecordsmcp internal/designrecords
```

If runtime reproduction is feasible, run the Design Records MCP over the repository root and call `propose_record_create` / `validate_records` via JSON-RPC or the available test harness.

No write operation should be accepted during this task. Proposal creation is allowed only if it remains no-write.

## Evidence

2026-06-02 Codex investigation reproduced all four observed authoring proposal failures.

- Finding 1: `body` plus `fields` is reproduced. When `body` is present, file content is taken from `body`; structured `fields` do not drive rendered content. Implementation entrypoint: `authoring.go` line 512. Classification: spec gap.
- Finding 2: missing `fields.id` is reproduced even when top-level `id` is present. Implementation entrypoint: `authoring.go` lines 1499 and 1511. Classification: implementation bug plus contract gap.
- Finding 3: `domain: "mcp"` with an uppercase ID domain such as `V01-REQ-MCP-992` is reproduced as a case-sensitive mismatch. Implementation entrypoint: `authoring.go` lines 653 and 664. Classification: spec gap.
- Finding 4: propose-time validation scope pollution is reproduced. `propose_record_create(V01-REQ-MCP-990)` emitted unresolved `spec:*` diagnostics for `V01-INV-DOCS-002` through `V01-INV-DOCS-006`, while direct `validate_records(kind=investigation)` and direct full `validate_records` did not reproduce the same diagnostics. Implementation entrypoint: `authoring.go` lines 447 and 1023. Classification: implementation bug plus missing spec.

Root cause hypothesis for Finding 4: propose-time validation runs against a hypothetical full repository index via `ValidateRecords(ctx, hyp, ValidateRecordsRequest{})`, and the hypothetical index does not inherit existing `SemanticRefSources`. As a result, non-record spec semantic refs disappear during proposal validation.

Recommended next step: proceed to `V01-TASK-MCP-011-02` and specify the public MCP contract for `body` / `fields`, top-level `id` / `fields.id`, domain case normalization, and proposal affected-record validation scope.

Additional dogfooding evidence: attempting to create `V01-TASK-MCP-011-02` through Design Records MCP produced proposal `pw_000008`, but `accept_proposed_write` returned `written:false` with `invalid_request: proposal validation has error diagnostics`. The proposal diagnostics included unrelated `V01-INV-DOCS-002` through `V01-INV-DOCS-006` unresolved `spec:*` references and unrelated `TASK-MCP-005-*` metadata errors. This confirms that proposal accept is also blocked by repository-wide or otherwise unrelated diagnostics, matching Finding 4 and strengthening `V01-REQ-MCP-012`.
