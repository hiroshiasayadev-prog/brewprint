# Milestone 0: 実装境界を固定する

- **status**: closed
- **scope**: Go semantic model / QueryService boundary
- **source**: migrated from docs/TASKS.md
- **last_updated**: 2026-04-30

---

## Tasks

- [x] **V01-ADR-047: Go semantic model / query layer boundary を起票**
  - Raw YAML structs / ResolvedProject / QueryService / Renderer の責務境界を決める
  - renderer と MCP wrapper が Raw YAML structs を直接読まない方針を明文化する
  - validation / name resolution / derived model build の責務範囲を決める
  - renderer は ResolvedProject または view-specific view model を読む方針を明文化する
  - 具体的なview model型はDAG vertical slice実装中に固める

- [x] **V01-ADR-048: ResolvedProject index strategy を起票**
  - reverse lookup index を ResolvedProject build 時に構築する方針を決める
  - MCP tool / renderer は都度Raw構造を走査せず、ResolvedProjectのindexを読む
  - 初期index候補を決める
    - referencesBySource
    - referencesByTarget
    - tasksReadingStore
    - tasksWritingStore
    - transitionsByStateEventGuard
    - actionsByTask
    - scenariosByID

- [x] **V01-ADR-049: MCP / QueryService の reference 語彙統一 を起票**
  - 外部MCP tool名を `get_references` に統一する
  - 内部QueryService method名を `GetReferences` に統一する
  - `get_deps` / `GetDeps` は採用しない
  - MCP responseの中心語彙を `references` とする

- [x] **docs/spec/mcp/overview.md / schema.md / tools/*.md の元仕様を作る**
  - `get_signature` / `get_references` / `inspect` / `list_endpoints` の input / output を決める
  - QualifiedID / FileID / TransitionID / private sub node synthetic ID の指定形式を決める
  - ObjectRef / TransitionRef / AssetRef / Diagnostic / Reference の共通schemaを決める
  - not found 時の error 形式を決める
  - references は v1 では direct のみとする
  - inspect の粒度と reverse lookup の返却範囲を決める
  - まだ MCP server transport は実装しない
