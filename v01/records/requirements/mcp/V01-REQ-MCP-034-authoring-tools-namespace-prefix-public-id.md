# V01-REQ-MCP-034: authoring tools が namespace prefix 環境で正しい public ID を生成・解決すること

- **id**: V01-REQ-MCP-034
- **status**: accepted
- **date**: 2026-06-09
- **source_refs**:
  - V01-WORK-DRMCP-001
- **work_items**:
  - V01-WORK-MCP-031

## Requirement

Design Records MCP の authoring tools（`propose_record_create` / `propose_record_update`）は、namespace prefix が付与された records root（例: `v01/records` → prefix `V01-`）において正しく動作しなければならない。

- **R1**: `propose_record_create` で生成されるファイルの H1 / metadata `id` には namespace prefix を付与した public ID を使用すること
- **R2**: 次の REQ / WORK / TASK / ADR ID を算出する際、index 内の public ID から namespace prefix を除いた bare ID で比較を行い、重複・スキップを防ぐこと
- **R3**: `propose_record_create` / `propose_record_update` のレスポンス `target.domain` は bare ID から導出した正しい workflow domain を返すこと
- **R4**: `list_authoring_guides` / `get_authoring_guidance` は `<RecordsRoot>/guides/` を参照すること
- **R5**: `source_requirement` 等の relation target validation は namespace prefix 付き public ID を正しく受け付けること

## Required Outcome

- `propose_record_create` で作成されたファイルの `id` metadata が public ID（例: `V01-TASK-DRMCP-001-06`）である
- V01- prefix 付き records root に既存レコードがある状態で新規 ID を採番すると、重複しない次の連番が返る
- `list_authoring_guides` が `v01/records/guides/` のガイドを返す
- relation フィールドに `V01-REQ-MCP-033` 等の public ID を指定しても `invalid_target` にならない
- `go test ./drmcp/src/...` が全件 PASS
