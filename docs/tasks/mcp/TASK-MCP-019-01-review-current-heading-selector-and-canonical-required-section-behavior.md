# TASK-MCP-019-01: Review current heading selector and canonical required-section behavior

- **id**: TASK-MCP-019-01
- **status**: done
- **date**: 2026-06-05
- **work_item**: WORK-MCP-019
- **source_requirement**: REQ-MCP-021
- **estimate**: 0.5d-1d
- **depends_on**:
- **outputs**:
  - Current behavior inventory for named section update, heading selector, and required-section validation
  - Decision input for spec contract and implementation boundary
  - Follow-up task list confirmation for WORK-MCP-019

## Goal

REQ-MCP-021 の実装前に、Design Records MCP authoring update path と workflow artifact validation path における heading handling の current behavior / contract gap を確認し、case-only required-section heading canonicalization をどこで扱うべきか決める。

## Work

- `SPEC-design-records-mcp-tools` の `propose_record_update` / named section replacement / diagnostics contract を確認する。
- `ADR-093` の authoring transaction model boundary を確認する。
- `internal/designrecords` の authoring update, section selector, workflow artifact validation, required section handling を調査する。
- `internal/designrecordsmcp` の tool call test / contract test で既存期待値を確認する。
- case-only mismatch を proposal-based repair として扱う場合の候補実装層を整理する。
- ambiguous heading match と non-case heading mismatch の扱いを分離して記録する。

## Done condition

- case-only heading mismatch の current behavior が説明できる。
- canonical required section headings の定義元または判定元が特定されている。
- 修正対象層が authoring proposal generation / validation / selector matching のどれかに分類されている。
- spec update task, regression test task, implementation task の入力が明確になっている。
- 直接 filesystem edit なしで実施できた範囲と、Codex 等への委譲が必要な範囲が分離されている。

## Verification

- 読んだ spec / ADR / implementation / test file を Evidence に列挙する。
- 既存挙動について、必要なら Codex に `go test` / targeted grep / reproduction を依頼する prompt を作成する。
- この task と WORK-MCP-019 の relation が Design Records MCP validation で通ることを確認する。

## Evidence

Investigation completed.

Result:

- Proceed with a mixed boundary.
- Scope is limited to validation-required workflow section headings.
- User-defined optional headings remain outside this requirement.
- Exact selector behavior remains the default.
- Case-only fallback is allowed only after exact match fails and only for canonical required headings.
- A successful fallback must produce a proposal diff that changes the actual heading to the canonical heading text.
- Ambiguous case-insensitive matches must fail with candidate headings.
- Non-case differences remain failures under existing rules.
- Validation should stay strict and may add repair diagnostics for case-only required-heading mismatches.

Current behavior confirmed:

- Case-only mismatch does not match today.
- Validation reports the canonical required heading as missing.
- Existing candidate-heading support exists on selector failure.
- Existing package tests passed for `internal/designrecords` and `internal/designrecordsmcp`.

Next step: create and execute `TASK-MCP-019-02` for the tools spec contract update.
