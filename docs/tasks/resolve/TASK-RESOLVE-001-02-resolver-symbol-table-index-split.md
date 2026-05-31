# TASK-RESOLVE-001-02: resolver symbol table / index を public node と file-private sub node で分離する

- **id**: TASK-RESOLVE-001-02
- **status**: done
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

- Implementation updated resolver / symbol table indexing so `NodesByQID` is limited to public/main nodes.
- File-private sub nodes are resolved by internal file-local identity (`<file>#<local>`), and registered for file-local lookup without project-wide public QID uniqueness.
- Cross-file private sub nodes with the same local ID are allowed.
- Same-file private sub node local ID duplication is reported as `duplicate_sub_node`.
- Bare ID resolution now preserves the required order: same-file private sub node/source first, then same-module main node fallback.
- Initial implementation review returned `Needs revision before commit` due to two identity-boundary issues:
  - private sub task / join return asset `ProducedBy` still used public-shaped QID, causing possible asset identity collision;
  - public-shaped private aliases were visible through `NodesByID` / kind indexes and could be misused as public selectors.
- Follow-up fix changed private sub task / join return asset `ProducedBy` to the node internal QID (`<file>#<local>`), so `AssetID` / `AssetObjectKey` do not collide across files with same private local ID and same return name.
- Follow-up fix stopped registering private sub node public-shaped aliases in `NodesByID` / `TasksByQID` / `BranchesByQID` / `ForksByQID` / `JoinsByQID`; public selector paths now use public-only lookup and do not return private nodes via public-looking QID aliases.
- Regression tests were added / updated for:
  - cross-file same private sub task local ID allowed;
  - same-file duplicate private sub task local ID emits `duplicate_sub_node`;
  - same-module duplicate main task emits `duplicate_node`;
  - local flow step resolves same-file private sub task first;
  - duplicate private sub task handling does not cascade into `unresolved_flow_task`;
  - cross-file same private sub task local ID with same `returns.name` does not collide in asset ID / object key;
  - public-shaped private alias is not registered in public lookup indexes;
  - full QID transition action does not validate through a private alias;
  - asset query paths do not first-hit the wrong file-private asset.
- Changed files reported by implementation handoff:
  - `internal/semantic/project.go`
  - `internal/resolve/builder.go`
  - `internal/resolve/symbols.go`
  - `internal/resolve/flow.go`
  - `internal/resolve/names.go`
  - `internal/resolve/references.go`
  - `internal/resolve/validation.go`
  - `internal/resolve/return_source_test.go`
  - `internal/resolve/type_ref_test.go`
  - `internal/resolve/subnode_scope_test.go`
  - `internal/query/service.go`
  - `internal/query/source.go`
  - `internal/query/service_test.go`
  - `internal/mcp/server_test.go`
  - `internal/render/dag/flow_renderer.go`
  - `internal/render/dag/renderer_test.go`
- Verification reported by implementation handoff:
  - `go test ./...` -> pass
  - `go run ./cmd/brewprint validate --yaml-root docs\uc\002-brewprint-self-hosting\yaml` -> ok
  - `go run ./cmd/brewprint render --yaml-root docs\uc\002-brewprint-self-hosting\yaml --out $env:TEMP\brewprint-uc002-render-review2 --clean` -> rendered 11 file(s)
- Scope constraints preserved:
  - M15 / `v1.1.0-spec` was not reopened.
  - MCP private object exposure / ObjectRef schema migration was not implemented.
  - spec files were not changed in this implementation task.
