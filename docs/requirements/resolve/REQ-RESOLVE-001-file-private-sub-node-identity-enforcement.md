# REQ-RESOLVE-001: file-private sub node identity enforcement が必要

- **id**: REQ-RESOLVE-001
- **status**: captured
- **date**: 2026-05-31
- **source_refs**:
  - ADR-058
  - ADR-078
- **work_items**:
  - WORK-RESOLVE-001

## 要求

brewprint resolver は、file-private sub node を project-wide public QualifiedID の一意性制約に混入させず、file-local identity として扱う必要がある。

特に task file 内の private sub task は、別ファイルに同じ local ID が存在しても `duplicate_node` として扱わず、同一ファイル内の flow / reads / writes / params / returns.source 等から file-local に解決できる必要がある。

## 発見根拠

UC-002 full validate / render において、以下の repeated task IDs に起因する既存 failure が確認された。

- `build_response`
- `query_service`
- `validate_request`

Codex investigation では、`go run ./cmd/brewprint validate --yaml-root docs\uc\002-brewprint-self-hosting\yaml` が `42 error(s), 0 warning(s)` で失敗し、内訳は `duplicate_node` 21件と `unresolved_flow_task` 21件であると確認された。

repeated IDs は各 file 内では1回ずつであり、別 file 間の private sub task ID 重複である。したがって、UC-002 YAML fixture の同一 file 内 duplicate ではない。

直接原因は `internal/resolve/symbols.go` の `addNode` が sub task を main/public task と同じ `NodesByQID` / project-wide duplicate 判定に載せている点である。duplicate 判定後に return するため、該当 sub task が `NodesByFile[fileID]` にも登録されず、後続の flow step 解決が `unresolved_flow_task` に連鎖している。

## 期待する状態

- main node は project-wide public QualifiedID を持ち、従来通り project-wide unique として扱われる。
- file-private sub node は public QualifiedID を持たず、project-wide `duplicate_node` 判定の対象外となる。
- sub node local ID は同一 file 内で一意である。
- 別 file に存在する同名 sub node local ID は許容される。
- flow / reads / writes 等の bare ID 解決は、同一 file 内 sub node を優先し、見つからない場合のみ同一 module の main node へフォールバックする。
- duplicate によって file-local index 登録が失敗し、別 diagnostic に連鎖する状態を解消する。

## 関連する設計判断

ADR-058 は、sub node を QualifiedID 一意性制約から除外し、file-private として扱う判断を accepted として保持している。

ADR-078 は、MCP query layer 上で private / generated object を安定識別する synthetic ID を `<semantic-anchor-id>#<local-id>` とする方針を accepted として保持している。ただし、この requirement は MCP public contract migration そのものを所有しない。

## 明示的に除外する範囲

以下は本 requirement の直接範囲外とする。

- M15 / `v1.1.0-spec` の再オープン
- UC-002 YAML の private sub task ID を project-wide unique に rename する運用回避
- `list_objects(include_private)` の追加・変更
- ObjectRef の `anchor` / `visibility` schema migration
- private helper model exposure schema
- transition / state machine identity policy

MCP private object exposure の未反映・未確定部分は、必要に応じて別 requirement として扱う。

## Boundary

- 本 requirement は、file-private sub node identity に関する期待状態と、現行 resolver behavior との乖離を捕捉する。
- 本 requirement は、具体的な task graph、実装順序、修正ファイル一覧、テストケース詳細、verification evidence を所有しない。
- 後続 work item は、ADR-058 の spec 反映、resolver symbol table / index 修正、regression tests、UC-002 validate / render verification を分解して追跡する。
- M14a legacy record と問題領域は重なるが、新形式 artifact として追跡する場合は本 requirement を起点にする。
