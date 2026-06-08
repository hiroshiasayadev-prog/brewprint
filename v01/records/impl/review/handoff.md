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

### V01-ADR-050 spec-first転換に伴う方針更新（2026-04-29）

source層レビュー（Sonnet担当）でADR間の整合性問題（F-S-001 等）が表面化したことをきっかけに、V01-ADR-050（proposed）でドキュメント運用方針を**spec-first**に転換した。
これに伴い、本レビューの基準も以下のように変わる。

- **specが「現行仕様の唯一の正」**。implが見るべきは spec
- **ADRは「設計判断の根拠記録」**。現行仕様の記述は今後 spec に移管していく
- 既存49 ADRは漸進移行ルールに従い、レビューで触れたタイミングで spec に移管する

レビューの軸は2つに増える：

1. **spec ↔ impl 整合性**: 現行仕様（spec）と実装の差分
2. **ADR → spec 移行発掘**: ADRに書かれているが spec に書かれていない決定の発掘と移行作業

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

- ADR/specで決めた型分割（例: V01-ADR-001 node type splitting, V01-ADR-007 asset/store boundary）が実際のpackage/typeに反映されているか
- ADR/specで決めた境界（例: V01-ADR-047 query layer boundary, V01-ADR-048 indexes, V01-ADR-049 reference vocabulary）が守られているか
- specで定義した構造（nodes.md / edges.md のフィールド体系）が型に反映されているか

### B. 命名整合

- ADR/specで決めた語彙（例: V01-ADR-049 `references` 統一）がコード上の名前と一致しているか
- ADR/specで言及されている概念（QualifiedID / FileID / TransitionID など）が型として存在するか
- 概念をリネームしたが ADR/spec に反映されていない例がないか

### C. 機能整合

- specで定義したMCP responseの形が query の返却型と一致するか
- specで定義した diagnostic code が実装で出力されているか
- ADR/specで決めた index 群（V01-ADR-048）が `semantic.Project` に揃っているか
- ADR/specで決めた render rule が renderer の出力で観測できるか

### D. ADR ↔ spec ↔ impl の三者整合（V01-ADR-050以降の追加軸）

- ADR本文に「現行仕様」が書かれているがspecに反映されていないものがないか
- spec / impl が一致しているが、その根拠ADRが見えなくなっているものがないか
- 旧ADRが superseded されずに残っていないか
- spec gap として一時的に残した未決事項が忘れられていないか

---

## 4. レビュー計画（6会話）

各会話は独立に動くよう、必要なADR/spec/implだけ読めばよい構成。

| # | レビュー会話 | 対象 impl | 突き合わせ ADR | 突き合わせ spec |
|---|---|---|---|---|
| 1 | source層 | `internal/source/` (3) | V01-ADR-002, 030 | overview.md（→ project-layout.md / file-types.md 新設候補） |
| 2 | rawyaml層 | `internal/rawyaml/` (12) | V01-ADR-001, 006, 011, 030 | nodes.md, edges.md |
| 3 | semantic層 | `internal/semantic/` (17) | V01-ADR-001, 006, 007, 011, 014, 018, 019, 021, 024, 026, 028, 029, 042, 048 | nodes.md, edges.md |
| 4 | resolve層 | `internal/resolve/` (12) | V01-ADR-003, 015, 020, 023, 025, 027, 031, 033, 034, 035, 040, 048 | overview.md, diagnostics.md（→ naming.md 新設候補） |
| 5 | render層 | `internal/render/` (17) | V01-ADR-004, 005, 008, 009, 022, 024, 026, 028, 029, 035, 036, 037, 038, 039, 041, 042, 043, 044, 045, 046 | views/全部（→ project-layout.md 新設候補） |
| 6 | query+mcp層 | `internal/query/` (7) + `internal/mcp/` (5) + `cmd/brewprint/` (2) | V01-ADR-047, 048, 049 | mcp.md, diagnostics.md |

CLI（cmd/brewprint）は query+mcp層の会話で扱う（mcp起動 / validate / render コマンド構成）。

新設candidate spec（漸進移行で作成予定）:
- `project-layout.md` — yaml/, renders/, render_index.yaml の構造（V01-ADR-043, 045由来）
- `file-types.md` — `as:`, `nodes:`, FileKind分類（V01-ADR-030, 039由来）
- `naming.md` — QualifiedID, name resolution, module nesting, FK解決（V01-ADR-002, 003, 027, 031, 033由来）

### 会話の進め方テンプレート

各レビュー会話の冒頭で：

1. `docs/doc-policy.md` を読む（spec-first方針を含む最新版）
2. このファイル（`docs/impl/review/handoff.md`）を読む
3. 該当する **対象impl** を全ファイル読む
4. 該当する **spec** を読む（spec-firstなのでまずこちら）
5. 必要に応じて **ADR** を読む（根拠を辿るため、または spec に未反映の決定を発掘するため）
6. 整合性チェックの軸 A/B/C/D に沿って所見を出す
7. 発見事項を `docs/impl/review/findings-{layer}.md` に記録
8. 必要に応じて ADR の修正・新規起票・spec修正・spec新設・ADR→spec移行を提案

`{layer}` は `source` / `rawyaml` / `semantic` / `resolve` / `render` / `query-mcp` のいずれか。

---

## 5. レビューで見つかったら起こすアクション

レビュー中に発生しうる対応を分類して、findings-{layer}.md にラベル付きで残す。

- **[doc-only]** ADR/specに書き漏れがあるだけ。コードはOK。doc追記のみ。
- **[ADR-update]** 実装が ADR の決定を更新している。ADR を修正 or supersedeする新ADR起票。
- **[ADR-new]** 新しい設計判断が暗黙に含まれている。新ADR起票。
- **[code-fix]** ADR/spec が正で、implが逸脱している。コード修正。
- **[question]** 判断つかず、ユーザー確認が必要。
- **[spec-gap]** ADR/specにも書かれず、実装にも反映されていない論点。残課題として記録。
- **[spec-migrate]** V01-ADR-050以降の追加。ADRに書かれた仕様記述をspecに移管する作業。

---

## 6. 進捗トラッカー

| # | 層 | status | 担当 | findings file | 主なアクション |
|---|---|---|---|---|---|
| 1 | source | done | Sonnet (review) + Opus (spec移行) | [findings-source.md](./findings-source.md) | F-S-001〜008 整理完了。V01-ADR-051/052 起票、V01-ADR-002/030/043/045 漸進移行、spec 3本新設 |
| 2 | rawyaml | not_started | - | - | - |
| 3 | semantic | not_started | - | - | - |
| 4 | resolve | not_started | - | - | naming.md 肉付け（actor global / FK 解決 / cross-edge）も担当範囲 |
| 5 | render | not_started | - | - | - |
| 6 | query-mcp | not_started | - | - | - |

各会話で開始時に `not_started -> in_progress`、完了時に `in_progress -> review_done`、findings記録後に `review_done -> done`、findings fileのリンクを記入。

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

V01-ADR-001〜049 全部 accepted、V01-ADR-050（spec-first転換）proposed。`docs/adr/` を参照。

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

新設予定（V01-ADR-050漸進移行で作成）:
- `project-layout.md`
- `file-types.md`
- `naming.md`

---

## 9. レビュアー候補

各層レビューを誰が担当するか。

- **Sonnet 4.6**: 機械的な突き合わせに強い。source / rawyaml / 一部render（DAG/ER/API/state）に推奨
- **Opus 4.7（自分）**: ADR間の関係性や暗黙の前提を読む推論力で価値が出る。semantic / resolve / sequence / wireframe / query-mcp に推奨
- **Opus 4.6**: 試したい場合に1層で比較
- **チャッピー**: 実装担当のため除外（書いた本人のレビューは見落としを起こす）

ただし source層レビュー（Sonnet）の結果から、Sonnet は「ADR間の整合性問題の指摘」までは強いが「実装の暗黙前提の発掘」「対応案の提示」がやや弱い傾向が見えた。
複雑な層は Opus 4.7、機械的な層は Sonnet という振り分けが現実的。

---

## 10. 次にやること

1. V01-ADR-050（spec-first転換）が accepted になったら、source層レビューで発見した F-S-001〜F-S-008 を spec-first 方針で改めて整理
2. その整理に基づき `findings-source.md` を作成
3. F-S-001 等のうち spec 移行が必要なものは、`project-layout.md` / `file-types.md` の新設と合わせて作業
4. source層が完了したら rawyaml層レビューに進む

---

## 11. このファイルの更新履歴

- 2026-04-29 (1): 初版。レビュー計画 + チェック軸 A/B/C/D + 進捗トラッカー作成
- 2026-04-29 (2): V01-ADR-050（spec-first転換）に伴う更新。レビュー軸を2軸化、漸進移行ルール反映、source層レビュー結果を進捗表に反映、レビュアー候補セクションを追加
