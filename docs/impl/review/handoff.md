# Impl Review Handoff

- **status**: in-progress
- **last_updated**: 2026-04-29
- **purpose**: ADR/spec ↔ impl 整合性レビューを複数会話に分割して進めるための親メモ

---

## 1. 背景

implがM1スコープ（UC-001のDAG render垂直スライス）から大きく成長した。
現在は M0〜M9 + Post-M9-5 まで完了しており、`internal/` 配下の Go ファイルは約77本。
1会話で全体をレビューするのはcontext的に不可能なため、**層ごとに別会話で分割レビュー**する。

レビューの主目的は **ADR / spec と impl の整合性確認**。
コード品質や設計クリティーク自体は副次的で、「ADR/specの決定が実装に反映されているか」「実装中に変えたが ADR/spec に反映していない決定がないか」を中心に見る。

---

## 2. impl 全体俯瞰

### 規模

```
cmd/brewprint/             2 ファイル
internal/
├── source/                3 ファイル
├── rawyaml/              12 ファイル
├── semantic/             17 ファイル
├── resolve/              12 ファイル
├── render/
│   ├── dag/              4 ファイル
│   ├── api/              2 ファイル
│   ├── er/               2 ファイル
│   ├── sequence/         2 ファイル
│   ├── state/            2 ファイル
│   ├── wireframe/        2 ファイル
│   ├── placement/        3 ファイル
│   └── project/          2 ファイル
├── query/                7 ファイル
├── mcp/                  5 ファイル
└── testutil/golden/      1 ファイル
合計: 約77ファイル
```

### CLI構成

`cmd/brewprint/main.go` に3コマンド：

- `brewprint mcp --yaml-root <path>` — MCP server (stdio JSON-RPC)
- `brewprint validate --yaml-root <path> [--format text|json]` — diagnostics validation
- `brewprint render --yaml-root <path> --out <path> [--clean]` — 一括render

### 依存方向（M9 summary より）

```text
source  -> rawyaml
resolve -> rawyaml, semantic
query   -> semantic
mcp     -> query
renderer -> semantic
```

禁止事項：
- `internal/mcp -> rawyaml` 直参照
- mcp内で YAML load / resolve / renderer 呼び出し
- renderer内で Raw YAML structs を直接読むこと
- renderer内で name resolution / semantic validation を再実装すること

---

## 3. 整合性チェックの軸

各層レビューで以下の4軸を共通的に確認する。

### A. 構造的整合

- ADRで決めた型分割（例: ADR-001 node type splitting, ADR-007 asset/store boundary）が実際のpackage/typeに反映されているか
- ADRで決めた境界（例: ADR-047 query layer boundary, ADR-048 indexes, ADR-049 reference vocabulary）が守られているか
- specで定義した構造（nodes.md / edges.md のフィールド体系）が型に反映されているか

### B. 命名整合

- ADRで決めた語彙（例: ADR-049 `references` 統一）がコード上の名前と一致しているか
- ADRで言及されている概念（QualifiedID / FileID / TransitionID など）が型として存在するか
- 概念をリネームしたが ADR に反映されていない例がないか

### C. 機能整合

- specで定義したMCP responseの形が query の返却型と一致するか
- specで定義した diagnostic code が実装で出力されているか
- ADRで決めた index 群（ADR-048）が `semantic.Project` に揃っているか
- ADRで決めた render rule が renderer の出力で観測できるか

### D. ADRからの逸脱

- 「実装中に変えたが ADRに反映していない」決定がないか
- 旧ADRが superseded されずに残っていないか
- spec gap として一時的に残した未決事項が忘れられていないか

---

## 4. レビュー計画（6会話）

各会話は独立に動くよう、必要なADR/spec/implだけ読めばよい構成。

| # | レビュー会話 | 対象 impl | 突き合わせ ADR | 突き合わせ spec |
|---|---|---|---|---|
| 1 | source層 | `internal/source/` (3) | ADR-002, 030 | overview.md |
| 2 | rawyaml層 | `internal/rawyaml/` (12) | ADR-001, 006, 011, 030 | nodes.md, edges.md |
| 3 | semantic層 | `internal/semantic/` (17) | ADR-001, 006, 007, 011, 014, 018, 019, 021, 024, 026, 028, 029, 042, 048 | nodes.md, edges.md |
| 4 | resolve層 | `internal/resolve/` (12) | ADR-003, 015, 020, 023, 025, 027, 031, 033, 034, 035, 040, 048 | overview.md, diagnostics.md |
| 5 | render層 | `internal/render/` (17) | ADR-004, 005, 008, 009, 022, 024, 026, 028, 029, 035, 036, 037, 038, 039, 041, 042, 043, 044, 045, 046 | views/全部 |
| 6 | query+mcp層 | `internal/query/` (7) + `internal/mcp/` (5) + `cmd/brewprint/` (2) | ADR-047, 048, 049 | mcp.md, diagnostics.md |

CLI（cmd/brewprint）は query+mcp層の会話で扱う（mcp起動 / validate / render コマンド構成）。

### 会話の進め方テンプレート

各レビュー会話の冒頭で：

1. `docs/doc-policy.md` を読む
2. このファイル（`docs/impl/review/handoff.md`）を読む
3. 該当する **対象impl** を全ファイル読む
4. 該当する **ADR / spec** を必要に応じて読む（ADRタイトル + Decisionだけでもよい）
5. 整合性チェックの軸 A/B/C/D に沿って所見を出す
6. 発見事項を `docs/impl/review/findings-{layer}.md` に記録
7. 必要に応じて ADR の修正・新規起票・spec修正を提案

`{layer}` は `source` / `rawyaml` / `semantic` / `resolve` / `render` / `query-mcp` のいずれか。

---

## 5. レビューで見つかったら起こすアクション

レビュー中に発生しうる対応を分類して、findings-{layer}.md にラベル付きで残す。

- **[doc-only]** ADRやspecに書き漏れがあるだけ。コードはOK。doc追記のみ。
- **[ADR-update]** 実装が ADR の決定を更新している。ADR を修正 or supersedeする新ADR起票。
- **[ADR-new]** 新しい設計判断が暗黙に含まれている。新ADR起票。
- **[code-fix]** ADR/spec が正で、implが逸脱している。コード修正。
- **[question]** 判断つかず、ユーザー確認が必要。
- **[spec-gap]** ADR/specにも書かれず、実装にも反映されていない論点。残課題として記録。

---

## 6. 進捗トラッカー

| # | 層 | status | 担当会話 | findings file | 主なADR更新 |
|---|---|---|---|---|---|
| 1 | source | not_started | - | - | - |
| 2 | rawyaml | not_started | - | - | - |
| 3 | semantic | not_started | - | - | - |
| 4 | resolve | not_started | - | - | - |
| 5 | render | not_started | - | - | - |
| 6 | query-mcp | not_started | - | - | - |

各会話で開始時に `not_started -> in_progress`、完了時に `in_progress -> done`、findings fileのリンクを記入。

---

## 7. 既知のreview対象外スコープ

- **テスト網羅性**: 各 `*_test.go` の網羅性レビューはこの整合性レビューでは扱わない。fixtureが実態と合っているかは見るが、test caseの過不足は別作業。
- **golden fixtureの妥当性**: UC-001 renders/ の中身が ADR どおりかは render層レビューで部分的に見るが、UC-001 を網羅的に見直す作業は別。
- **Goコード品質（命名、エラー処理、idiom）**: 副次的に見る。ただし「ADR命名と一致しない型名」は B 軸で扱う。
- **Performance / 並行性**: スコープ外。

---

## 8. 補助情報

### 直近の実装summary

- `docs/impl/go-m7-summary.md`
- `docs/impl/go-m8-summary.md`
- `docs/impl/go-m9-summary.md`

これらは「何を実装したか」「どのADRに準拠したか」を時系列で説明している。
特に M9 summary は QueryService / reverse lookup index / scenario / transition の対応関係を詳しく書いている。

### ADR一覧（タイトルベース）

ADR-001〜049 全部 accepted。`docs/adr/` を参照。

### spec一覧

```
docs/spec/
  overview.md
  nodes.md
  edges.md
  diagnostics.md
  mcp.md
  views/
    dag.md
    er.md
    state-diagram.md
    sequence-diagram.md
    api-table.md
    wireframe.md
    wireframe.css
    wireframe.preview.css
```

---

## 9. 次にやること

1. 別会話を立ち上げて、source層レビューから開始する
2. 進捗トラッカーを更新していく
3. 6層分のレビューが終わったら、このファイル自体に総括セクションを追加する

---

## 10. このファイルの更新履歴

- 2026-04-29: 初版。レビュー計画 + チェック軸 A/B/C/D + 進捗トラッカー作成。
