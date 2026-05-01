# Milestone 12: MCP impact traversal / source assist を拡張する

- **status**: open
- **scope**: MCP / QueryService
- **source**: migrated from docs/TASKS.md
- **last_updated**: 2026-05-01

---

## Tasks

- [x] **`get_source` を実装する**
  - semantic objectに対応するYAML snippetを返す
  - Raw YAML AST全体公開ではなく、ResolvedProject objectのsource補助情報として扱う
  - selector / source range / fallback挙動を `docs/spec/mcp/tools/get-source.md` に定義する
  - source line/columnが未取得の場合の返却形式を決める
  - MCP wrapper test / QueryService testを追加する

- [x] **`get_reference_tree` または depth指定つきreference traversalを設計する**
  - direct referencesだけでは不足する変更影響範囲を辿れるようにする
  - ただし、このtaskは「参照グラフ traversal」までを扱い、変更種別ごとの解釈は `analyze_impact` に分ける
  - 別tool `get_reference_tree` にするか、`get_references` に `depth` inputを追加するかを比較する
  - ADR-055で `get_references` direct only維持 + `get_reference_tree` 別toolを採用済み
  - ADR-055で cycle detection / max depth / kind filter / direction の方針を確定済み
  - ADR-049のdirect reference方針はsupersedeせず、ADR-055で拡張する
  - `docs/spec/mcp/tools/get-reference-tree.md` へ正式仕様を反映済み
  - 実装タスクは次セッション以降で分割する

- [x] **`analyze_impact` を設計する**
  - 設計完了: ADR-056 で全体方針を確定させ、 `docs/spec/mcp/tools/analyze-impact.md` に仕様反映済み
  - 設計対話向けに、対象objectと変更種別から影響範囲を返す上位toolとして位置づけ
  - `get_references` は direct reference APIとして維持し、impact analysisを混ぜない
  - `get_reference_tree` / direct references / flow wiring / sequence step / render output (file粒度) を材料にする
  - input: `selector` + `change` (discriminated object) + `scope_modules` + `max_impacts`
  - output: `summary` / `impacts[]` / `coverage` (必須) / `assumptions` / `truncated` / `diagnostics`
  - severity と fixability を別軸として分け、 `mechanical` の必要条件5要件をspecで固定
  - `recommended_action` と `suggested_fixes[]` を二段で返す
  - flow wiring / sequence step / type signature identity は v1 coverage に含める
  - structural compatibility / semantic contract / render presentation details は v1 除外
  - 初期対象は field / model / task / transition / state / event / actor / store とし、 view系・asset・primitiveは unsupported selector
  - 実装タスクは次milestoneで管理
