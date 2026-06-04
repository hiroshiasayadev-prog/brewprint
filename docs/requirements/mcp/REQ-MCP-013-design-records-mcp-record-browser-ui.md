# REQ-MCP-013: Design Records MCP record browser UI

- **id**: REQ-MCP-013
- **status**: captured
- **date**: 2026-06-02
- **source_refs**:
- **work_items**:

## Requirement

work item / task 数の増加に伴い、VSCode でファイルを直接探す方法が煩雑になっている。
Design Records MCP の read 側ツールを活用した record browser UI を提供し、
record の一覧・検索・詳細確認をファイルシステムを介さずに行えるようにする必要がある。

## Evidence

- record 数増加とともに VSCode でのファイル直接参照が煩雑になっている
- Design Records MCP はすでに `list_records` / `get_record` / `get_records` の read ツールを持つ

## Required Outcome

- record 種別（decision / requirement / work_item / task 等）を横断して一覧・検索できる UI が存在する
- ファイルシステムを直接触ることなく record の詳細を確認できる

## Explicitly Excluded Scope

- authoring / write 操作の UI 化はこの REQ のスコープ外
- 外部サービスとの連携

## Boundary

- Design Records MCP の read 側 API を使った UI レイヤーとして位置づける
- record 間 relation の navigation はこの REQ では扱わない
