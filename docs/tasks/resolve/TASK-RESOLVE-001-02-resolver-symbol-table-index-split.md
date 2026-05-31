# TASK-RESOLVE-001-02: resolver symbol table / index を public node と file-private sub node で分離する

- **id**: TASK-RESOLVE-001-02
- **status**: todo
- **date**: 2026-05-31
- **work_item**: WORK-RESOLVE-001
- **source_requirement**: REQ-RESOLVE-001
- **estimate**: 1d-2d
- **depends_on**:
  - TASK-RESOLVE-001-01
- **outputs**:
  - resolver / symbol table implementation update
  - file-local duplicate validation
  - local bare ID resolution preservation

## Goal

resolver の symbol table / index を修正し、main/public node は project-wide QualifiedID で、file-private sub node は file-local identity で扱う。

## Work

- `internal/resolve/symbols.go` の `addNode` 周辺を確認し、main node と sub node の登録経路を分離する。
- main/public node のみ `NodesByQID` の project-wide duplicate 判定対象にする。
- sub node は `NodesByFile[fileID]` 等の file-local index へ登録し、別 file 間の同名 local ID を許容する。
- 同一 file 内の sub node local ID 重複は diagnostic として検出する。
- duplicate 判定によって file-local index 登録が失敗し、`unresolved_flow_task` へ連鎖する挙動を解消する。
- `internal/resolve/builder.go` / `internal/resolve/names.go` / `internal/resolve/flow.go` など、sub node の QID / ID / local resolution に関わる箇所を必要最小限で更新する。
- ADR-078 の MCP synthetic ID 方針に反する public QID 化を増やさない。
- `list_objects(include_private)` / ObjectRef schema migration は本 task で扱わない。

## Done condition

- sub node が project-wide `duplicate_node` の対象にならない。
- 同一 file 内 sub node duplicate は検出される。
- 別 file 間の同名 sub node local ID は許容される。
- flow / reads / writes 等の file-local resolution が維持される。
- MCP public contract migration を混ぜずに resolver bugfix として閉じられる。

## Verification

- 対象コードを読み戻し、main/public node と file-private sub node の登録経路を確認する。
- 後続 `TASK-RESOLVE-001-03` で追加する regression tests が通ることを確認する。
- UC-002 validate / render の duplicate task QID / unresolved flow task issue が解消することを確認する。

## Evidence

未実施。
