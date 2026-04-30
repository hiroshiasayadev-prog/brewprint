# Milestone 11: MCP diagram element query を拡張する

- **status**: open
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

- [ ] **flow wiring queryを設計する**
  - DAG上のflow step / param wiringをMCPから辿れるようにする
  - `get_references` に含めるか、`inspect(task)` / `inspect(file)` / `get_reference_tree` / `analyze_impact` の材料に留めるかを決める
  - relation kind候補: `flow_step` / `flow_param` / `flow_branch_case` / `flow_foreach_over`
  - 既存方針「MCP v1ではflow wiringをget_referencesに返さない」を変更する場合は、小ADRまたはspec更新で扱う
  - `inspect(task).members.flow.entries` のdraft schemaと整合させる
  - DAG rendererのview modelとQueryServiceの責務境界を崩さない
