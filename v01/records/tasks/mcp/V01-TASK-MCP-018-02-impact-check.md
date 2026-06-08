# V01-TASK-MCP-018-02: Assess strict create contract test and runtime impact

- **id**: V01-TASK-MCP-018-02
- **status**: done
- **date**: 2026-06-03
- **work_item**: V01-WORK-MCP-018
- **source_requirement**: V01-REQ-MCP-019
- **estimate**: 0.5d
- **depends_on**:
  - V01-TASK-MCP-018-01
- **outputs**:
  - Strict fields-required create contract impact inventory
  - Test and runtime smoke update plan
  - Decision on whether validate_records failure is in-scope or split out

## Goal

`propose_record_create` を strict contract に切り替える前に、既存 test / runtime smoke / fixture への影響範囲を洗い出し、修正方針を確定する。

Strict contract は以下を前提とする。

- `fields` は schema-level required とする。
- `fields + body` は section-only body source として valid とする。
- `fields + body_cache_id` は `fields + body` の retry form として valid とする。
- `body`-only / `body_cache_id`-only full-record create は invalid とする。
- `body + body_cache_id` と `fields + body + body_cache_id` は invalid のままとする。
- invalid request でも submitted `body` が string として受け取れている場合は、新しい `body_cache` を返して本文を保護する。

## Work

- `V01-TASK-MCP-018-01` review result の affected tests を再確認する。
- `fields` schema required 化で破壊される MCP schema tests / runtime smoke を一覧化する。
- `body`-only / `body_cache_id`-only invalid 化で破壊される unit / integration tests を一覧化する。
- invalid request with submitted body が body_cache を返すべき cases を整理する。
- `TestToolsCallSuccess/validate_records` failure が V01-WORK-MCP-018 の blocker か、V01-WORK-MCP-016 側の別対応に分離すべきか判断する。
- V01-TASK-MCP-018-03 / 04 / 05 に渡す修正順序を整理する。

## Done condition

- Strict contract 変更で落ちる既存 test / smoke の一覧がある。
- 各 test / smoke をどう修正するかが説明されている。
- `validate_records` failure の扱いが in-scope / out-of-scope / blocker のいずれかに分類されている。
- V01-TASK-MCP-018-03 の spec / guidance 更新と V01-TASK-MCP-018-04 の実装に渡せる判断材料が揃っている。

## Verification

- `rg` / `go test` などで影響対象を確認する。
- 調査結果は Evidence に、変更対象 file / test name / expected update を含めて記録する。

## Evidence
- Impact check completed by Claude Code review for `V01-TASK-MCP-018-02`.
- Verdict: OK to proceed to `V01-TASK-MCP-018-03`.
- `go test ./internal/designrecords ./internal/designrecordsmcp -count=1 -v` passed under the current contract.
- Strict contract summary confirmed:
  - `fields` is schema-level required.
  - `fields`, `fields + body`, and `fields + body_cache_id` are the valid create forms.
  - `body`-only / `body_cache_id`-only create become invalid.
  - `body + body_cache_id` and `fields + body + body_cache_id` remain invalid.
  - invalid request with submitted string `body` should return a new `body_cache` to preserve the submitted content.
- Test impact identified 5 breaking tests, all assigned to `V01-TASK-MCP-018-04`:
  - `TestToolsProposeRecordCreateSchemaFieldsOptional` should invert the required-field assertion.
  - `TestAuthoringCreateInputContractNormalization` legacy body-only subcase should expect `proposal_created:false`, `invalid_request`, and body_cache preservation.
  - `TestAuthoringCreateInputContractNormalization` cachedLegacy subcase should expect invalid create without submitted body preservation.
  - `TestAuthoringCreateInputContractNormalization` fieldsAndCache subcase should change from early `invalid_request` to cache lookup behavior.
  - `TestBodyCacheReturnClassification/no_cache_fields_plus_body_cache_id_create` should change from early `invalid_request` to cache lookup behavior.
- Additional tests recommended for `V01-TASK-MCP-018-04`:
  - successful `fields + body_cache_id` retry flow with valid cached section body.
  - body-only / body_cache_id-only invalid behavior, including body preservation for submitted `body`.
- Spec impact confirmed for `V01-TASK-MCP-018-03`:
  - update `docs/spec/design-records-mcp/tools.md` body source rules / input combinations / error handling / legacy compatibility language.
  - update `docs/spec/design-records-mcp/schema.md` body cache model wording.
  - authoring guidance files do not require updates because they already recommend fields-based MCP create.
- Implementation impact assigned to `V01-TASK-MCP-018-04`:
  - add `fields` to `propose_record_create` required schema in `internal/designrecordsmcp/tools.go`.
  - remove `fields + body_cache_id` early rejection in `internal/designrecords/authoring.go`.
  - reject `fields == nil` while preserving submitted `body` in body_cache.
  - simplify `prepareCreate` by removing the legacy body-only branch.
  - no new diagnostic code is required.
- `TestToolsCallSuccess/validate_records` failure was classified as not an issue for this work item: it currently passes and is not a blocker for `V01-WORK-MCP-018`.
