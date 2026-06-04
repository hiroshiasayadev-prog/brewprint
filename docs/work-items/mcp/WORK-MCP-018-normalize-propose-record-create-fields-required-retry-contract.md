# WORK-MCP-018: Normalize propose_record_create fields-required retry contract

- **id**: WORK-MCP-018
- **status**: done
- **date**: 2026-06-03
- **source_requirement**: REQ-MCP-019
- **impact_refs**:
  - SPEC-design-records-mcp-tools
  - ADR-093
  - REQ-MCP-014
  - REQ-MCP-015
  - WORK-MCP-014
  - WORK-MCP-012
- **tasks**:
  - TASK-MCP-018-01
  - TASK-MCP-018-02
  - TASK-MCP-018-03
  - TASK-MCP-018-04
  - TASK-MCP-018-05

## Goal

`REQ-MCP-019` を解消するため、`propose_record_create` の正規 create contract を fields-required / section-only body source 前提へ整理する。

特に `fields + body_cache_id` を `fields + body` failure 後の retry form として valid 化し、`body`-only / `body_cache_id`-only create を invalid にして、MCP が H1 / metadata / resolved ID の生成責任を持つ contract に統一する。

## Boundary

- `propose_record_create` の input contract と retry behavior を対象とする。
- `fields + body` の section-only create 方針は `WORK-MCP-014` の決定を前提とする。
- failed propose response の body_cache return behavior は `WORK-MCP-012` の close 済み成果を前提とし、再オープンしない。
- `propose_record_update` の body source contract は対象外とする。
- Design Records MCP authoring transaction model 全体の再設計は対象外とする。

## Impact Scope

- `SPEC-design-records-mcp-tools`: `propose_record_create` の preferred / valid / invalid input combinations、retry form、strict body source boundary を更新する。
- Authoring guidance: 新規 REQ / WORK / TASK / ADR 作成時の preferred create form を fields-required 前提に整理する。
- MCP tool schema / validation / rendering: `fields` schema required 化、`fields + body_cache_id` valid 化、legacy body-only create invalid 化、invalid request body_cache preservation を実装する。
- Regression tests: fields-only、fields + body、fields + body_cache_id、invalid legacy body-only boundary、invalid combinations、invalid request body_cache preservation を追加・更新する。
- Runtime smoke: stdio JSON-RPC path で strict create contract と retry path を確認する。

## Task flow

1. `TASK-MCP-018-01` で current contract、legacy mode、retry gap をレビューする。
2. `TASK-MCP-018-02` で strict fields-required contract の test / runtime 影響範囲を調査する。
3. `TASK-MCP-018-03` で spec / guidance を strict contract に更新する。
4. `TASK-MCP-018-04` で実装と regression tests を行う。
5. `TASK-MCP-018-05` で runtime smoke、validation、REQ / WORK close synchronization を行う。

## Task Candidates

- `TASK-MCP-018-01`: current contract / legacy mode / retry gap review
- `TASK-MCP-018-02`: strict fields-required create contract impact assessment
- `TASK-MCP-018-03`: spec and guidance update for strict fields-required create contract
- `TASK-MCP-018-04`: implementation and regression tests
- `TASK-MCP-018-05`: runtime smoke and close synchronization

## Completion Condition

- `propose_record_create` の preferred create mode が `fields` schema-required として spec / guidance / tool schema に反映されている。
- `fields + body_cache_id` が `fields + body` の retry form として valid に実装・検証されている。
- `body`-only / `body_cache_id`-only create が invalid として実装・検証されている。
- `body + body_cache_id` と `fields + body + body_cache_id` は invalid のまま維持されている。
- invalid request でも submitted `body` が string として受け取れている場合、新しい `body_cache` を返して本文を保護する behavior が実装・検証されている。
- Regression tests と runtime smoke が通っている。
- `REQ-MCP-019` の status / work_items / evidence が同期されている。

## Evidence

- `TASK-MCP-018-01` reviewed the current create contract and retry gap.
- `TASK-MCP-018-02` completed strict contract impact assessment and confirmed all breaking tests were scoped to implementation.
- `TASK-MCP-018-03` updated `docs/spec/design-records-mcp/tools.md` and `docs/spec/design-records-mcp/schema.md` for strict fields-required create contract.
- `TASK-MCP-018-04` implemented strict `propose_record_create` behavior and regression tests.
- `TASK-MCP-018-05` completed runtime smoke through actual Design Records MCP stdio JSON-RPC path.
- Runtime smoke result: PASS.
- Implementation verification: `go test ./... -count=1` passed across all packages.
- Strict contract verified:
  - `fields`, `fields + body`, and `fields + body_cache_id` are valid create forms.
  - `body`-only / `body_cache_id`-only create are invalid.
  - `body + body_cache_id` and `fields + body + body_cache_id` remain invalid.
  - invalid request with submitted string `body` preserves the body through returned `body_cache`.
