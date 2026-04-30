# Milestone 11: MCP diagram element query を拡張する

- **status**: closed
- **scope**: MCP / QueryService
- **source**: migrated from docs/TASKS.md
- **last_updated**: 2026-04-30

---

## Tasks

- [x] **implicit asset selectorを実装する**
  - DAG上のasset nodeを直接queryできるようにする
  - selector形式は `<producer>#<name>` のstable synthetic IDを採用
  - asset signatureで name / producer / model / scope_file を返す
  - asset referencesで producer / consumer task への関係を返す
  - `produces_asset` との整合を保つ
  - UC-001のDAG assetでQueryService / MCP wrapper testを追加済み
  - `docs/spec/mcp.md` の AssetRef / selector support matrixを更新済み

- [x] **private sub node selectorを実装する**
  - file-local task / branch / fork / join を直接queryできるようにする
  - selectorは `<file-id>#<local-id>` または `file` + `local_id` を使う
  - get_signature / get_references / inspect に対応
  - main task inspect内の `members.sub_tasks` と同じObjectRef表現に揃える
  - UC-001の checkout sub task / fork 相当でQueryService / MCP wrapper testを追加済み

- [x] **flow wiring queryを設計する**
  - DAG上のflow step / param wiringは、MCP v1の `get_references` には含めない
  - flow wiringは `inspect(task).members.flow.entries` のflow inspect用schemaに閉じる
  - `flow_step` / `flow_param` / `flow_branch_case` / `flow_foreach_over` は `Reference.kind` ではなく、flow inspect用の語彙として扱う
  - 将来の `get_reference_tree` / `analyze_impact` では traversal 材料として利用可能
  - `inspect(task).members.flow.entries` の最小schemaを `docs/spec/mcp.md` に追記済み
  - DAG rendererのview modelとQueryServiceの責務境界は維持する
  - 既存方針を覆していないため、新規ADRは不要
