# V01-TASK-DATA-011-01: Review request-side and generic container candidates and decide ownership

- **id**: V01-TASK-DATA-011-01
- **status**: done
- **date**: 2026-06-05
- **work_item**: V01-WORK-DATA-011
- **source_requirement**: V01-REQ-DATA-002
- **estimate**: 0.5d-1d
- **depends_on**:
- **outputs**:
  - N-002, N-004, N-007, N-008, N-012, N-016, N-018, TF-QUERY-RESULT の候補レビュー決定テーブル
  - 明示的 no-action アウトカム
  - V01-TASK-DATA-011-02 へのインプット

## Goal

V01-WORK-DATA-011 が所有する request-side / generic container バケットの各候補（N-002, N-004, N-007, N-008, N-012, N-016, N-018, TF-QUERY-RESULT）について、それぞれの所有権と処理方針を決定する。

このタスクは決定・計画のみを行う。UC-002 YAML 移行、fixture / golden 再生成、実装変更、新規 REQ / WORK 起票は行わない。

## Work

- V01-INV-DATA-002 の各候補の note 退避内容と現在の YAML フィールドを確認する。
- TF-QUERY-RESULT については、YAML ファイルおよびタスク資料上の実体を特定してから方針を決定する。
- 各候補に対して以下のいずれか一つの処理方針を選択する:
  - 公開モデルシェイプとして正式化が必要（public model shape needed）
  - ジェネリックコンテナ規約（`list<str>` 等）への移行が必要（generic container convention needed）
  - リクエストサイドヘルパーモデルとしての扱いが必要（request-side helper treatment needed）
  - 明示的 no-action（変更不要、または別オーナーが既に存在）
- 決定テーブルをこのタスクの Evidence セクションに記録する。
- V01-TASK-DATA-011-02（スペック / YAML クリーンアップスコープ決定）への入力を明確にする。

## Included Scope

- V01-INV-DATA-002 の N-002, N-004, N-007, N-008, N-012, N-016, N-018 の確認と方針決定。
- TF-QUERY-RESULT の実体特定と方針決定。
- 明示的 no-action 理由の記録。

## Excluded Scope

- UC-002 YAML 移行。
- fixture / golden 再生成。
- パーサー、レンダラー、バリデーター、MCP 実装変更。
- 新規 REQ / WORK / ADR の起票。
- V01-TASK-DATA-011-02 以降のタスク作成。
- V01-WORK-DATA-011 の `done` 遷移。
- 完了済み DATA work（V01-WORK-DATA-001 〜 V01-WORK-DATA-010）の再開。

## Done condition

- 全候補（N-002, N-004, N-007, N-008, N-012, N-016, N-018, TF-QUERY-RESULT）に対して処理方針が決定されている。
- 明示的 no-action アウトカムが記録されている。
- V01-TASK-DATA-011-02 へのインプットが使用可能な状態である。
- UC-002 YAML 移行、fixture 再生成、実装変更、新規 REQ / WORK 起票は行っていない。

## Verification

- このタスクと V01-WORK-DATA-011 のみが変更されていることを確認する。
- Design Records MCP validation をこのタスクと V01-WORK-DATA-011 に対して実行する。
- Design Records MCP 経由で両レコードを取得して確認する。

## Evidence

完了日: 2026-06-05

### Sources reviewed

- `docs/investigations/data/INV-DATA-002-uc002-notes-retreat-inventory-and-m15-release-boundary-input.md`
- `docs/tasks/data/TASK-DATA-009-01-reconcile-remaining-uc-002-notes-retreat-candidates.md`
- `docs/tasks/data/TASK-DATA-009-02-classify-remaining-uc-002-notes-retreat-successor-buckets.md`
- `docs/tasks/data/TASK-DATA-009-03-decide-remaining-uc-002-notes-retreat-successor-outcomes.md`
- `docs/requirements/data/REQ-DATA-003-private-helper-model-signature-exposure-boundary.md`
- `docs/work-items/data/WORK-DATA-011-uc-002-request-side-generic-container-cleanup.md`

### TF-QUERY-RESULT の実体確認

V01-TASK-DATA-009-01 のインベントリより:

> Equivalent non-N-id source from V01-TASK-DATA-003-04 / V01-TASK-DATA-006-01: eight UC-002 MCP task-file `query_service.returns.model:any` to `build_response.params[].model:any` patterns remain deferred by the DATA-004 / V01-REQ-DATA-003 private-helper params policy.

TF-QUERY-RESULT は、UC-002 の MCP タスクファイルにある `query_service.returns.model:any` → `build_response.params[].model:any` のクロスタスクデータフローパターン（8 件）を指す。V01-INV-DATA-002 の N-xxx 形式のモデルフィールドエントリではなく、タスクフロー（TF）レベルのパターンである。

V01-REQ-DATA-003 により `params[].model` に private helper 参照を使うと validation error になるため、このパターンは `any` で据え置かれてきた。

### 決定テーブル

| ID | YAML ファイル / フィールドパス | 現在 | note 退避の意味 | 処理方針 | 理由 |
|---|---|---|---|---|---|
| N-002 | `analyze_impact_request.yaml` / `scope_modules` | `any` | モジュールスコープフィルタの文字列配列 | generic container convention needed (`list<str>`) | 単純な文字列配列。V01-ADR-069 の原則に沿って `list<str>` で正式化できる |
| N-004 | `analyze_impact_response.yaml` / `summary` | `any` | by_severity / by_fixability / by_kind のカウント辞書 | public model shape needed | 意味的に異なる 3 フィールドを持つ構造体。`list<str>` では不十分で、名前付きヘルパーモデル（例: `AnalyzeImpactSummary`）が必要 |
| N-007 | `analyze_impact_response.yaml` / `assumptions` | `any` | ツール前提・制約の文字列配列 | generic container convention needed (`list<str>`) | 単純な文字列配列 |
| N-008 | `analyze_impact_response.yaml` / `truncated_reasons` | `any` | 切り捨て理由の文字列配列 | generic container convention needed (`list<str>`) | 単純な文字列配列 |
| N-012 | `get_reference_tree_request.yaml` / `kinds` | `any` | Reference.kind フィルタの文字列配列 | generic container convention needed (`list<str>`) | 文字列配列として正式化できる。要素 vocabulary（Reference.kind enum）は V01-WORK-DATA-012 の enum 整理スコープに委任する |
| N-016 | `get_reference_tree_response.yaml` / `truncated_reasons` | `any` | max-node / max-edge 切り捨て理由の文字列配列 | generic container convention needed (`list<str>`) | 単純な文字列配列 |
| N-018 | `get_references_request.yaml` / `kinds` | `any` | Reference.kind フィルタの文字列配列 | generic container convention needed (`list<str>`) | N-012 と同様。要素 vocabulary は V01-WORK-DATA-012 に委任 |
| TF-QUERY-RESULT | task-file `query_service.returns.model` → `build_response.params[].model`（8 件） | `any` | クロスタスク query result フローパターン | no-action | `params[].model` は V01-REQ-DATA-003 により private helper 参照が validation error。task-file レベルのクロスタスク契約問題であり、モデル YAML フィールドの型付けではない。完全解決には公開モデルが必要だが、それは別スコープ |

### 明示的 no-action アウトカム

- **TF-QUERY-RESULT**: `params[].model:any` の `any` は V01-REQ-DATA-003 による意図的な据え置き。`returns.model` 側のみ private helper を付けても cross-task contract pattern の半解決にとどまり有害。完全解決には公開モデルが必要だが、それは別スコープ。

### V01-TASK-DATA-011-02 へのインプット

スペック / YAML クリーンアップスコープ決定タスクへの入力:

- **generic container 対象（6 件）**: N-002, N-007, N-008, N-012, N-016, N-018 → `list<str>` への移行
- **public model shape 対象（1 件）**: N-004 → `AnalyzeImpactSummary` 相当の名前付きヘルパーモデル定義
- **no-action（1 件）**: TF-QUERY-RESULT → スコープ外

### Verification note

このタスクは決定・計画のみを行った。

UC-002 YAML 移行、fixture / golden 再生成、パーサー / レンダラー / バリデーター / MCP 実装変更、新規 REQ / WORK / ADR 起票は行っていない。

### Post-edit verification

Design Records MCP validation passed for:

- V01-TASK-DATA-011-01

Design Records MCP retrieval confirmed:

- V01-TASK-DATA-011-01 exists with status `done`.
- V01-WORK-DATA-011.tasks includes V01-TASK-DATA-011-01.

Working tree check confirmed this task touched only:

- `docs/tasks/data/TASK-DATA-011-01-review-request-side-and-generic-container-candidates-and-decide-ownership.md`
