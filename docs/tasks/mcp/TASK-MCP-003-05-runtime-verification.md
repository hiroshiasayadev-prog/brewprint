# TASK-MCP-003-05: Workflow artifact MCP support の tests と runtime verification を実施する

- **id**: TASK-MCP-003-05
- **status**: done
- **date**: 2026-05-26
- **work_item**: WORK-MCP-003
- **source_requirement**: REQ-MCP-003
- **estimate**: 0.5d
- **depends_on**:
  - TASK-MCP-003-04
- **outputs**:
  - automated test result evidence
  - stdio MCP runtime verification evidence
  - public contract と runtime behavior の照合結果

## Goal

採用・実装された workflow artifact MCP support が public contract と一致し、実際の workflow ID-as-ref chain を runtime で扱えることを確認する。

## Work

- TASK-MCP-003-04 で追加・更新された automated tests の結果を確認する。
- 採用された surface に応じて `REQ-MCP-003` / `WORK-MCP-003` / `TASK-MCP-003-*` を対象とする runtime query / resolve / validation を確認する。
- Unsupported または MVP 外とした入力・diagnostic が contract 境界を越えて実装されていないことを確認する。
- Verification 結果を close review に渡す。

## Done condition

- 採用された public contract の正常系と重要な境界系について、test と runtime evidence が記録されている。
- Spec / implementation / runtime behavior に未解消の矛盾がない、または blocker として明示されている。
- TASK-MCP-003-06 が close review を開始できる。

## Verification

- Test command と結果、runtime request / response の確認対象を evidence に記録する。
- ID-as-ref と physical path 非対応の境界を採用 contract に応じて検証する。

## Evidence

- 2026-05-27: Runtime verification に着手した。対象は `TASK-MCP-003-04` で実装済みの workflow artifact public record surface、workflow ID-as-ref resolver、declared workflow relation validation、physical path / workflow semantic prefix unsupported boundary に限定する。
- 2026-05-27: Automated test verification: `go test ./internal/designrecords ./internal/designrecordsmcp` passed。続けて `go test ./...` passed。
- 2026-05-27: Runtime command: repository root で `go run ./cmd/design-records-mcp -root .` を stdio JSON-RPC runtime として起動し、`initialize` / `notifications/initialized` / `tools/list` / `tools/call` を実行した。Process exit code は 0、stderr 出力なし。
- `tools/list`: `list_records` / `get_record` / `get_records` / `resolve_reference` / `validate_records` が exposed。`list_records.kind` と `validate_records.kind` enum は `requirement` / `work_item` / `task` を含むことを runtime schema response で確認した。
- `list_records`: `kind=requirement,id=REQ-MCP-003` は `REQ-MCP-003` を `kind=requirement` として返し、`requirement.work_items` に `WORK-MCP-003` を含んだ。`kind=work_item,id=WORK-MCP-003` は `source_requirement=REQ-MCP-003` と `TASK-MCP-003-01`〜`TASK-MCP-003-06` を返した。`kind=task,id=TASK-MCP-003-04` は `work_item=WORK-MCP-003` / `source_requirement=REQ-MCP-003` を返した。
- `get_record`: `REQ-MCP-003` / `WORK-MCP-003` / `TASK-MCP-003-04` はすべて found。各 record は kind-specific detail object を持ち、`WORK-MCP-003.status=verification_pending`、`TASK-MCP-003-04.status=done` が current artifact と一致した。`include_body=false` では `record.body` は返らないことを確認した。
- `get_record(include_body=true)`: `TASK-MCP-003-04` で `record.body` が返り、body prefix が `# TASK-MCP-003-04: Workflow artifact MCP support を実装する` であることを確認した。
- `get_records`: `ADR-092` / `REQ-MCP-003` / `WORK-MCP-003` / `TASK-MCP-003-04` は `retrieval_status=found`。`TASK-MCP-003-99` は item-level `retrieval_status=not_found` と `record_not_found` diagnostic を返し、tool execution error にはならなかった。重複指定した `TASK-MCP-003-04` は first occurrence のみ返り、top-level `duplicate_requested_id_ignored` info diagnostic は `first_index=3`, `duplicate_indexes=[5]` を返した。
- `resolve_reference`: `REQ-MCP-003` / `WORK-MCP-003` / `TASK-MCP-003-04` はすべて `status=resolved`, `ref_kind=record_id`, `target.target_type=record` を返し、`target.record_kind` はそれぞれ `requirement` / `work_item` / `task`、path と record ID は current artifact と一致した。
- `resolve_reference` missing / unsupported boundary: `TASK-MCP-999-99` は `status=unresolved`, `ref_kind=record_id`, `target=null`, `unresolved_reference` diagnostic。Physical path `docs/tasks/mcp/TASK-MCP-003-04-workflow-artifact-mcp-implementation.md` と semantic prefix `task:TASK-MCP-003-04` はいずれも `status=unsupported`, `ref_kind=unsupported`, `target=null`, `unsupported_reference` info diagnostic を返した。
- `validate_records`: `kind=requirement` と `kind=task` は diagnostic なし。`kind=work_item` と filter なし `{}` は `WORK-MCP-001.tasks` の `M19` に対する `invalid_workflow_relation_target` を 1 件返したが、`REQ-MCP-003` / `WORK-MCP-003` / `TASK-MCP-003-*` に対する `unresolved_workflow_relation` / `invalid_workflow_relation_target` / `workflow_relation_mismatch` / `workflow_source_requirement_mismatch` は出ていない。これは既存別 artifact 由来の observed diagnostic として TASK-MCP-003-06 の close review に渡す。
- Investigation metadata boundary: current repository の `docs/investigations/**/INV-*.md` metadata block には `REQ-*` / `WORK-*` / `TASK-*` を参照する runtime fixture は存在しない。Runtime fixture を追加・改変せず、Pass 2 automated tests により `REQ-*` / `WORK-*` support と `TASK-*` unsupported boundary が固定済みであることを確認対象とした。
- 結論: Workflow artifact MCP support の public query surface、workflow ID-as-ref resolver、declared workflow relation validation、physical path / semantic prefix unsupported boundary は実 runtime で採用 contract と一致した。Implementation / spec / runtime behavior の未解消 blocker は観測されていないため、本 task を `done` とする。
