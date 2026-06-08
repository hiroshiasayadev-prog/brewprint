# V01-TASK-DATA-011-02: Decide spec and YAML cleanup scope for generic container and public model shape candidates

- **id**: V01-TASK-DATA-011-02
- **status**: done
- **date**: 2026-06-05
- **work_item**: V01-WORK-DATA-011
- **source_requirement**: V01-REQ-DATA-002
- **estimate**: 0.5d-1d
- **depends_on**:
  - V01-TASK-DATA-011-01
- **outputs**:
  - generic container 候補（6 件）の YAML クリーンアップスコープ決定
  - N-004 public model shape のスペックスコープ決定
  - V01-TASK-DATA-011-03 へのインプット

## Goal

V01-TASK-DATA-011-01 の決定テーブルに基づき、各候補の spec / YAML クリーンアップスコープを決定する。

このタスクは仕様決定のみを行う。YAML 移行、fixture / golden 再生成、実装変更は行わない。

## Work

- generic container 対象（N-002, N-007, N-008, N-012, N-016, N-018）について:
  - 対象 YAML ファイルの現在の `any` フィールドを確認する。
  - `list<str>` への移行が spec 上どのように表現されるべきか決定する。
  - 各フィールドに必要な変更範囲（フィールド型のみか、model note も更新するかなど）を特定する。
- public model shape 対象（N-004: `analyze_impact_response.summary`）について:
  - 現在の YAML を確認する。
  - `AnalyzeImpactSummary` 相当のヘルパーモデル定義に必要なフィールドと型を決定する。
  - 定義場所（既存ファイルへの追加 or 新規ヘルパーファイル）を決定する。
- スコープ決定をこのタスクの Evidence セクションに記録する。
- V01-TASK-DATA-011-03 への実装インプットを明確にする。

## Included Scope

- N-002, N-007, N-008, N-012, N-016, N-018 の `list<str>` 移行スコープ決定。
- N-004 の公開ヘルパーモデル定義スコープ決定。
- 変更対象ファイルリストと変更内容の特定。

## Excluded Scope

- YAML 移行の実施。
- fixture / golden 再生成。
- パーサー、レンダラー、バリデーター、MCP 実装変更。
- 新規 REQ / WORK / ADR の起票。
- V01-TASK-DATA-011-03 以降のタスク作成。
- V01-WORK-DATA-011 の `done` 遷移。
- TF-QUERY-RESULT（no-action 決定済み）の再検討。
- 完了済み DATA work（V01-WORK-DATA-001 〜 V01-WORK-DATA-010）の再開。

## Done condition

- generic container 対象 6 件すべてについてクリーンアップスコープが決定されている。
- N-004 の public model shape スコープが決定されている。
- V01-TASK-DATA-011-03 への実装インプットが使用可能な状態である。
- YAML 移行、fixture 再生成、実装変更は行っていない。

## Verification

- このタスクと V01-WORK-DATA-011 のみが変更されていることを確認する。
- Design Records MCP validation をこのタスクと V01-WORK-DATA-011 に対して実行する。
- Design Records MCP 経由で両レコードを取得して確認する。

## Evidence
完了日: 2026-06-05

### Sources reviewed

- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/analyze_impact_request.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/analyze_impact_response.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/get_reference_tree_request.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/get_reference_tree_response.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/get_references_request.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/string_list.yaml`
- `docs/tasks/data/TASK-DATA-011-01-review-request-side-and-generic-container-candidates-and-decide-ownership.md`

### string_list 型の確認

`string_list.yaml` で `kind: list, element: str` の共有 list model として定義済み。model note には "例: reference kind filter、truncated_reasons、coverage.analyzed / not_analyzed など" と明記されており、全 generic container 候補の移行先として適切。

### Generic container スコープ決定（6 件）

全候補を `any` から `string_list` へ移行する。変更は field type と field note のみで、model 構造への影響はない。

| 候補 | ファイル / フィールドパス | 変更内容 |
|---|---|---|
| N-002 | `analyze_impact_request.yaml` / `scope_modules` | `type: any` → `type: string_list`; field note から「専用list modelを作らずanyで暫定表現する」を除去 |
| N-007 | `analyze_impact_response.yaml` / `assumptions` | `type: any` → `type: string_list`; field note から「専用list modelを作らずanyで暫定表現する」を除去 |
| N-008 | `analyze_impact_response.yaml` / `truncated_reasons` | `type: any` → `type: string_list`; field note から「専用list modelを作らずanyで暫定表現する」を除去 |
| N-012 | `get_reference_tree_request.yaml` / `kinds` | `type: any` → `type: string_list`; field note から「専用list modelを作らずanyで暫定表現する」を除去 |
| N-016 | `get_reference_tree_response.yaml` / `truncated_reasons` | `type: any` → `type: string_list`; field note から「専用list modelを作らずanyで暫定表現する」を除去; model note の「`truncated_reasons` はstring list制約をv1 modelで厳密表現できないため any + note で保持する」も除去 |
| N-018 | `get_references_request.yaml` / `kinds` | `type: any` → `type: string_list`; field note から「専用string list modelを作らずanyで暫定表現する」を除去 |

### N-004 public model shape スコープ決定

`analyze_impact_response.yaml` に以下のヘルパーモデルを追加し、`summary.type` を `any` から `analyze_impact_summary` に変更する。

#### 新規ヘルパーモデル

**`analyze_impact_severity_counts`**（kind: struct）

| フィールド | 型 | 必須 | 説明 |
|---|---|---|---|
| `breaking` | `int` | 必須 | breaking severity の件数 |
| `warning` | `int` | 必須 | warning severity の件数 |
| `info` | `int` | 必須 | info severity の件数 |

**`analyze_impact_fixability_counts`**（kind: struct）

| フィールド | 型 | 必須 | 説明 |
|---|---|---|---|
| `mechanical` | `int` | 必須 | mechanical fixability の件数 |
| `suggested` | `int` | 必須 | suggested fixability の件数 |
| `manual_review` | `int` | 必須 | manual_review fixability の件数 |
| `unknown` | `int` | 必須 | unknown fixability の件数 |

**`analyze_impact_summary`**（kind: struct）

| フィールド | 型 | 必須 | 説明 |
|---|---|---|---|
| `by_severity` | `analyze_impact_severity_counts` | 必須 | severity 別集計 |
| `by_fixability` | `analyze_impact_fixability_counts` | 必須 | fixability 別集計 |
| `by_kind` | `any` | 必須 | impact kind 別集計。kind 語彙は実装裁量のため `any` を維持する |

#### `analyze_impact_response.summary` フィールド変更

- `type: any` → `type: analyze_impact_summary`
- field note: 「dict shapeを厳密表現しないためany」を除去し、`analyze_impact_summary` を参照する内容に更新

#### by_kind が `any` のまま残る理由

`by_severity` / `by_fixability` は固定語彙（3 値・4 値）のため具体型化できる。`by_kind` は impact kind 語彙が実装裁量であり、v1 model では dict key semantics を型として表現できないため `any` を維持する。

### V01-TASK-DATA-011-03 への実装インプット

V01-TASK-DATA-011-03（実装タスク）が行うべき変更:

1. `analyze_impact_request.yaml`: `scope_modules` → `string_list`
2. `analyze_impact_response.yaml`: `assumptions` → `string_list`、`truncated_reasons` → `string_list`、`summary` → `analyze_impact_summary`（3 ヘルパーモデルを同ファイルに追加）
3. `get_reference_tree_request.yaml`: `kinds` → `string_list`
4. `get_reference_tree_response.yaml`: `truncated_reasons` → `string_list`、model note 更新
5. `get_references_request.yaml`: `kinds` → `string_list`

合計 5 ファイル変更。新規ファイルなし。fixture / golden 再生成が必要かどうかは T3 で判断する。

### Verification note

このタスクは仕様決定のみを行った。

YAML 移行、fixture / golden 再生成、パーサー / レンダラー / バリデーター / MCP 実装変更、新規 REQ / WORK / ADR 起票は行っていない。

### Post-edit verification

Design Records MCP validation passed for:

- V01-TASK-DATA-011-02

Design Records MCP retrieval confirmed:

- V01-TASK-DATA-011-02 exists with status `done`.
- V01-WORK-DATA-011.tasks includes V01-TASK-DATA-011-01 and V01-TASK-DATA-011-02.

Working tree check confirmed this task touched only:

- `docs/tasks/data/TASK-DATA-011-02-decide-spec-and-yaml-cleanup-scope-for-generic-container-and-public-model-shape-candidates.md`
