# Milestone 10: MCP project exploration / view inspect を拡張する

- **status**: open
- **scope**: MCP / QueryService
- **source**: migrated from docs/TASKS.md
- **last_updated**: 2026-04-30

---

## Tasks

- [ ] **`list_objects` を実装する**
  - ADR-054 / `docs/spec/mcp.md` の設計対話coverageに準拠
  - project内の主要semantic objectを一覧・絞り込みできるようにする
  - object kind / module / file / label / source を返す
  - node / view / transition / field を対象に含める
  - asset / private sub node はM11で直接selector対応するまで対象外または任意扱いにする
  - MCP tool inputSchema / server wrapper testを追加する
  - QueryService unit testでUC-001の一覧を固定する
  - 実装後に `docs/spec/mcp.md` の tool仕様を追記する

- [ ] **`inspect(file)` を実装する**
  - FileID selectorで YAML file単位の定義内容を返す
  - node fileでは main node / sub nodes / flow summary / diagnostics を返す
  - state fileでは states / events / transitions / wireframe presence を返す
  - view fileでは view kind / id / target files or modules を返す
  - `render_index.yaml` は group summary / uncovered module warning 等を返す候補として扱う
  - `get_references(file)` の既存 `state_file` partial対応と整合させる
  - MCP wrapper test / QueryService testで代表fileを固定する
  - 実装後に `docs/spec/mcp.md` の selector support matrixを更新する

- [ ] **`inspect(view: api_table)` を実装する**
  - API Table view objectを直接inspectできるようにする
  - `http_root_path` / modules / include_submodules / collected endpoints / computed routes を返す
  - `list_endpoints` との役割分担を明確にする
  - excluded endpoint候補や収集対象0件sectionの扱いは実装時に判断し、必要ならspecへ追記する
  - UC-001の API Table view でQueryService / MCP wrapper testを追加する

- [ ] **`inspect(view: er_diagram)` を実装する**
  - ER Diagram view objectを直接inspectできるようにする
  - 対象modules / included stores / included models / FK relations / excluded refs summary を返す
  - default module単位ERと view YAMLによる横断ERの扱いを整理する
  - UC-001の ER view でQueryService / MCP wrapper testを追加する
  - 実装後に `docs/spec/mcp.md` の view inspect仕様を追記する
