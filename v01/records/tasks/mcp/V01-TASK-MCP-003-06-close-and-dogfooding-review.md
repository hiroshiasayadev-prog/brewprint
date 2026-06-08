# V01-TASK-MCP-003-06: Workflow artifact MCP support の evidence / status / dogfooding review を反映する

- **id**: V01-TASK-MCP-003-06
- **status**: done
- **date**: 2026-05-26
- **work_item**: V01-WORK-MCP-003
- **source_requirement**: V01-REQ-MCP-003
- **estimate**: 0.5d
- **depends_on**:
  - V01-TASK-MCP-003-05
- **outputs**:
  - V01-REQ-MCP-003 / V01-WORK-MCP-003 の close status と evidence 更新
  - 二回目 workflow artifact dogfooding の所感と follow-up candidate

## Goal

Workflow artifact MCP support の判断・反映・検証結果を requirement と work item に同期し、新形式 workflow artifact 運用の二回目 dogfooding として close 判定と残課題を記録する。

## Work

- V01-TASK-MCP-003-02〜05 の判断・spec・implementation・verification evidence を照合する。
- 採用、延期、または不採用の結果を V01-REQ-MCP-003 と V01-WORK-MCP-003 に反映する。
- Work item completion condition が満たされたか判定する。
- Task graph、短期 task 粒度、status ownership、MCP support の運用上の所感と follow-up candidate を記録する。

## Done condition

- V01-REQ-MCP-003 と V01-WORK-MCP-003 に最終結果と evidence が反映されている。
- Completion condition の充足または未充足理由が記録されている。
- 二回目 dogfooding から得た follow-up candidate が存在する場合、後続 requirement / work 候補として失われず記録されている。

## Verification

- Work item に task status checkbox / status copy を追加していないことを確認する。
- Requirement / work item / task relation が ID-as-ref のままであることを確認する。
- Verification evidence なしに capability 完了を宣言していないことを確認する。

## Evidence

### Close review input

- 2026-05-27: `V01-ADR-092` は accepted であり、workflow artifact を既存 Design Records MCP の record-oriented surface に統合する boundary を採用済みである。採用範囲は `requirement` / `work_item` / `task` の public retrieval、`REQ-*` / `WORK-*` / `TASK-*` の direct resolver input、declared workflow relation integrity validation、investigation metadata での `REQ-*` / `WORK-*` canonical reference support である。
- `V01-ADR-092` は、investigation metadata からの `TASK-*` canonical reference、orphan diagnostics、progress projection、workflow traversal、dependency cycle / execution order projection、physical path relation support、`req:` / `work:` / `task:` semantic prefix を MVP 外としている。
- `V01-TASK-MCP-003-03` は done。Design Records MCP overview / schema / tools、project artifact model、traceability specs に V01-ADR-092 の public contract を反映済みである。Spec reflection review で見つかった empty `task.depends_on:` normalization と resolver response ownership の implementation-facing ambiguity は implementation 前に解消済みである。
- `V01-TASK-MCP-003-04` は done。Workflow record surface、workflow parser / discovery / retrieval、direct resolver support、investigation metadata boundary、declared workflow relation validation、diagnostic payload / filter behavior を実装し、`go test ./internal/designrecords ./internal/designrecordsmcp` と `go test ./...` が passing である。
- `V01-TASK-MCP-003-05` は done。`go run ./cmd/design-records-mcp -root .` の stdio JSON-RPC runtime で workflow record discovery / retrieval、mixed `get_records`、workflow ID-as-ref resolver、missing / unsupported boundary、declared relation validation を確認済みである。

### Completion condition assessment

- `V01-WORK-MCP-003` completion condition 1 は満たされた。Public contract adoption、spec update、implementation、automated tests、runtime verification、`V01-REQ-MCP-003` への decision result 反映、close evidence がすべて完了している。
- `V01-REQ-MCP-003 -> V01-WORK-MCP-003 -> TASK-MCP-003-*` の target chain には、workflow relation diagnostic の blocker はない。
- Runtime 全体 validation では、target chain 外の既存 artifact `V01-WORK-MCP-001.tasks = M19` が `invalid_workflow_relation_target` として観測された。これは V01-ADR-091 / V01-ADR-092 より前の legacy M-series task label を現在の `work_item.tasks` validator が `TASK-*` ID-as-ref として検査した migration gap であり、`V01-WORK-MCP-003` の capability completion blocker ではない。
- Repository-wide `validate_records({})` は、上記既知 diagnostic により完全 clean ではない。この事実は close evidence と follow-up candidate として保持し、今回 `V01-WORK-MCP-001` は直接書き換えない。

### Second dogfooding assessment

- `REQ -> WORK -> TASK` の三層構造は、evidence 整理、decision、spec update、implementation、runtime verification、close review の追跡に十分だった。Requirement が必要性、work item が横断フローと completion condition、task が短期作業と evidence を所有する分担は維持できた。
- Work item が Mermaid task graph と completion condition を所有する運用は有効だった。工程の順序と close 判定を work item に集約でき、task artifact に個別 evidence を閉じ込められた。
- Task 粒度は概ね適切だった。`V01-TASK-MCP-003-01` で evidence / scope、`02` で decision、`03` で spec、`04` で implementation、`05` で runtime verification、`06` で close review に分けたことで、implementation 前の ambiguity を spec reflection review で止められた。
- Task status の正本を各 task artifact に置き、work item に checkbox / copied status を追加しない境界は維持された。Progress は work item status と progress summary に留め、task list は relation metadata のまま残した。
- ID-as-ref と physical path 非対応の境界は、MCP support 実装後も維持された。Runtime verification では `REQ-*` / `WORK-*` / `TASK-*` direct resolver が成立し、physical path と `task:` prefix は unsupported として確認済みである。
- Investigation relation の見落としは `V01-TASK-MCP-003-01` の evidence 補正で吸収できた。独立 investigation を増やさず、`spec:project-artifact-model` の既存 relation を decision input へ戻した判断は妥当だった。
- Spec reflection review は、empty `task.depends_on:` と resolver response ownership のような implementation-facing ambiguity を検出する gate として有効だった。Implementation に入る前に parser contract と spec ownership を明確化できた。
- Validator 導入により `V01-WORK-MCP-001.tasks = M19` の legacy migration gap が可視化された。次の運用改善では、pre-ADR-091 work item が保持する M-series label を `TASK-*` relation へ移行するか、legacy compatibility diagnostic / migration guidance として別 follow-up で扱うべきである。

### Entry docs cleanup

- Spec reflection review の non-blocking follow-up だった entry docs stale wording を確認した。
- `docs/doc-policy.md` には、workflow artifact MCP support をまだ V01-REQ-MCP-003 の後続判断対象であり現行 tool で利用可能と仮定しない、という stale 表現が残っていたため、今回の close evidence と矛盾しない範囲で更新した。
- `docs/requirements/README.md` / `docs/work-items/README.md` には、MCP record kind 追加を後続 spec / ADR 対象とする stale 表現が残っていたため、Design Records MCP spec が parser / public schema を所有する説明へ更新した。
- `docs/investigations/README.md` には、investigation metadata の canonical record ID-as-ref が `ADR-*` / `SPEC-*` / `INV-*` のみに見える stale 表現が残っていたため、`REQ-*` / `WORK-*` support と `TASK-*` unsupported boundary を追記した。
- `docs/tasks/README.md` の status-derived projection 記述は、progress projection が引き続き MVP 外であるため stale ではない。

### Follow-up candidates

- Legacy migration / compatibility follow-up: `V01-WORK-MCP-001.tasks = M19` のような pre-ADR-091 work item relation metadata を、`TASK-*` へ移行するか、legacy compatibility diagnostic として扱うかを判断する。現時点で新 requirement / work item の即時起票は必須ではない。
- Broader entry docs cleanup follow-up: 今回は close evidence と直接矛盾する小さい stale wording のみ更新した。Authoring docs 全体の体系的な文言統一は別作業候補として残せるが、`V01-WORK-MCP-003` close blocker ではない。
