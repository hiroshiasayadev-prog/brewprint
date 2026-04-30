# Milestone 12: MCP impact traversal / source assist を拡張する

- **status**: open
- **scope**: MCP / QueryService
- **source**: migrated from docs/TASKS.md
- **last_updated**: 2026-04-30

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

- [ ] **`analyze_impact` を設計する**
  - 設計変更相談向けに、対象objectと変更種別から影響範囲を返す上位toolとして扱う
  - `get_references` は direct reference APIとして維持し、impact analysisを混ぜない
  - `get_reference_tree` / direct references / flow wiring / render output mapping を材料にする
  - input候補:
    - `selector`
    - `change`: `rename` / `remove` / `change_type` / `change_contract` など
  - output候補:
    - `impacts[]`
    - `kind`
    - `severity`
    - `object` または `file`
    - `reason`
    - `via`
    - `recommended_action`
  - 初期対象は field / model / task / transition 程度に絞る
  - flow wiring と render output impact は段階的に追加する
  - 実装前に `docs/spec/mcp/overview.md` へtool仕様案を追記する
