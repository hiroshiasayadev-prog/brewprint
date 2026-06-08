# V01-TASK-DATA-011-03: Migrate generic container any fields and add public model shapes to UC-002 YAML

- **id**: V01-TASK-DATA-011-03
- **status**: done
- **date**: 2026-06-05
- **work_item**: V01-WORK-DATA-011
- **source_requirement**: V01-REQ-DATA-002
- **estimate**: 0.5d
- **depends_on**:
  - V01-TASK-DATA-011-02
- **outputs**:
  - 5ファイルの any フィールドを string_list またはヘルパーモデル型に移行
  - analyze_impact_response.yaml への 3 ヘルパーモデル追加
  - V01-TASK-DATA-011-04 へのインプット

## Goal

V01-TASK-DATA-011-02 のスコープ決定に基づき、UC-002 YAML の `any` フィールドを `string_list` またはヘルパーモデル型に移行する。

このタスクは YAML の変更のみを行う。fixture / golden 再生成、パーサー・レンダラー・バリデーター・MCP 実装変更は行わない。

## Work

- `analyze_impact_request.yaml`: `scope_modules` を `any` → `string_list` に変更、field note 更新
- `analyze_impact_response.yaml`:
  - `assumptions` を `any` → `string_list` に変更、field note 更新
  - `truncated_reasons` を `any` → `string_list` に変更、field note 更新
  - `summary` を `any` → `analyze_impact_summary` に変更、field note 更新
  - `analyze_impact_severity_counts`、`analyze_impact_fixability_counts`、`analyze_impact_summary` の 3 ヘルパーモデルを同ファイルに追加
- `get_reference_tree_request.yaml`: `kinds` を `any` → `string_list` に変更、field note 更新
- `get_reference_tree_response.yaml`: `truncated_reasons` を `any` → `string_list` に変更、field note 更新、model note 更新
- `get_references_request.yaml`: `kinds` を `any` → `string_list` に変更、field note 更新

## Included Scope

- 対象 5 ファイルの YAML フィールド型変更および field note 更新。
- `analyze_impact_response.yaml` への 3 ヘルパーモデル追加。

## Excluded Scope

- fixture / golden 再生成。
- パーサー、レンダラー、バリデーター、MCP 実装変更。
- 新規 REQ / WORK / ADR 起票。
- V01-WORK-DATA-011 の `done` 遷移（V01-TASK-DATA-011-04 担当）。
- TF-QUERY-RESULT（no-action 決定済み）の再検討。

## Done condition

- 対象 5 ファイルすべての変更が完了している。
- 3 ヘルパーモデルが `analyze_impact_response.yaml` に追加されている。
- fixture / golden 再生成、実装変更は行っていない。

## Verification

- このタスクと V01-WORK-DATA-011 のみが変更されていることを確認する。
- Design Records MCP validation をこのタスクと V01-WORK-DATA-011 に対して実行する。
- Design Records MCP 経由で両レコードを取得して確認する。

## Evidence
完了日: 2026-06-05

### YAML 変更サマリ

| ファイル | フィールド | 変更内容 |
|---|---|---|
| `analyze_impact_request.yaml` | `scope_modules` | `any` → `string_list`、field note 更新 |
| `analyze_impact_response.yaml` | `summary` | `any` → `analyze_impact_summary`、field note 更新 |
| `analyze_impact_response.yaml` | `assumptions` | `any` → `string_list`、field note 更新 |
| `analyze_impact_response.yaml` | `truncated_reasons` | `any` → `string_list`、field note 更新 |
| `analyze_impact_response.yaml` | （新規）`analyze_impact_severity_counts` | kind: struct、breaking / warning / info フィールド |
| `analyze_impact_response.yaml` | （新規）`analyze_impact_fixability_counts` | kind: struct、mechanical / suggested / manual_review / unknown フィールド |
| `analyze_impact_response.yaml` | （新規）`analyze_impact_summary` | kind: struct、by_severity / by_fixability / by_kind フィールド |
| `get_reference_tree_request.yaml` | `kinds` | `any` → `string_list`、field note 更新 |
| `get_reference_tree_response.yaml` | `truncated_reasons` | `any` → `string_list`、field note 更新、model note 更新 |
| `get_references_request.yaml` | `kinds` | `any` → `string_list`、field note 更新 |

### Verification note

YAML 移行のみを行った。fixture / golden 再生成、パーサー・レンダラー・バリデーター・MCP 実装変更は行っていない。

### Post-edit verification

Design Records MCP validation passed for:

- V01-TASK-DATA-011-03
- V01-WORK-DATA-011

Working tree check confirmed this task touched only:

- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/analyze_impact_request.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/analyze_impact_response.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/get_reference_tree_request.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/get_reference_tree_response.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/get_references_request.yaml`
- `docs/work-items/data/WORK-DATA-011-uc-002-request-side-generic-container-cleanup.md`
- `docs/tasks/data/TASK-DATA-011-03-migrate-generic-container-any-fields-and-add-public-model-shapes-to-uc-002-yaml.md`
