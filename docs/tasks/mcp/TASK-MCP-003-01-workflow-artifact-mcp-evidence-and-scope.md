# TASK-MCP-003-01: Workflow artifact MCP support の evidence と MVP scope 候補を整理する

- **id**: TASK-MCP-003-01
- **status**: done
- **date**: 2026-05-26
- **work_item**: WORK-MCP-003
- **source_requirement**: REQ-MCP-003
- **estimate**: 0.5d
- **depends_on**:
- **outputs**:
  - workflow artifact MCP support の現行 gap と利用 evidence summary
  - 最小 public contract 判断へ渡す scope 候補と保留論点
  - investigation artifact 起票要否の判断

## Goal

ADR-091 で確定した `REQ-* -> WORK-* -> TASK-*` 運用を MCP 経由で扱うために、現行 Design Records MCP contract の不足と二回目 dogfooding で必要になる最小 capability 候補を整理し、次の設計判断 task が開始可能な状態にする。

## Work

- ADR-091、REQ-MCP-003、WORK-MCP-002 の dogfooding 結果を確認する。
- 現行 Design Records MCP の対象 kind、`get_record(s)`、`resolve_reference`、`validate_records` の workflow artifact に対する境界を確認する。
- `REQ-*` / `WORK-*` / `TASK-*` の公開 surface と relation validation に関する MVP 候補を整理する。
- orphan diagnostics と task status 由来 progress projection を初期 MVP に含めるべきか、後続判断へ送るべきかを論点化する。
- 独立 investigation artifact が必要か判断する。

## Done condition

- 現行 MCP contract で workflow artifact chain を辿れない箇所が根拠付きで整理されている。
- public contract 判断で扱う candidate scope と、初期判断から除外し得る拡張論点が整理されている。
- investigation artifact の要否が記録されている。
- TASK-MCP-003-02 が開始可能である。

## Verification

- ADR-091 の ID-as-ref / physical path 非対応 / task status ownership と矛盾しないことを確認する。
- `spec:project-artifact-model` が定義する investigation から requirement / work item への relation を判断対象から落としていないことを確認する。
- 現行 spec で未対応の capability を、既に利用可能であるかのように扱っていないことを確認する。
- 本 task が contract を確定せず、判断対象を TASK-MCP-003-02 へ分離していることを確認する。

## Evidence

### 確認した運用前提

- `ADR-091` は、active workflow artifact を requirement / work item / task の三層とし、従来 milestone の到達点・作業フロー・task graph・close 条件の責務を work item が所有する方針を accepted として確定している。
- `ADR-091` は、workflow artifact 間の canonical relation を `REQ-*` / `WORK-*` / `TASK-*` ID-as-ref に限定し、physical path および `req:` / `work:` / `task:` semantic prefix を導入しないことを確定している。
- `WORK-MCP-002` は、evidence、判断、spec、implementation、runtime verification、close review を六つの短期 task に分割する初回 dogfooding を完了し、Mermaid flow と work item status による管理が有効だったことを記録している。
- `REQ-MCP-003` は、この運用チェーンを MCP で query / resolve / validate できない不足を捕捉し、workflow artifact support の最小 scope を後続判断へ送っている。
- `spec:project-artifact-model` は change flow において、investigation から ADR への判断候補、requirement への要求候補、spec への更新候補、work item への後続候補という relation を既に定義している。したがって `REQ-*` / `WORK-*` support の判断は、workflow 内部 chain のみでなく、investigation から workflow artifact を参照する既存概念 relation の operationalization にも影響する。

### Evidence 補正（2026-05-27）

- 初回整理では、workflow artifact 内部の relation validation を主な追加論点として記録した一方、`spec:project-artifact-model` が既に定義している investigation から requirement / work item への relation を判断対象として明示できていなかった。
- `docs/spec/concepts/project-artifact-model/index.md` と現行 investigation / Design Records MCP contract を追加確認した結果、investigation も relation-bearing artifact であり、比較軸は relation の有無ではなく、relation の意味・強制度・validation boundary であると補正する。
- この補正は TASK-MCP-003-01 の evidence / scope 整理を更新するものであり、public contract の採用判断自体は引き続き TASK-MCP-003-02 が所有する。

### 現行 Design Records MCP contract で確認した gap

- `SPEC-design-records-mcp-overview` / `schema` / `tools` は、現行 public record kind を `decision` / `spec` / `investigation` とし、requirement / work item / task を index / query / validation 対象として公開していない。
- `get_records` は exact record ID lookup のみを行う contract であり、現行 index に含まれない `REQ-*` / `WORK-*` / `TASK-*` は成功対象にならず、指定しても item-level `not_found` となる。
- `resolve_reference` は active `spec:` と `ADR-*` / `SPEC-*` / `INV-*` のみを supported input とし、`REQ-*` / `WORK-*` および physical path を `unsupported` として扱う現行 contract である。`TASK-*` も supported workflow ID として定義されていない。
- `validate_records` は design record metadata と investigation canonical reference を対象とし、`requirement.work_items`、`work_item.source_requirement` / `tasks`、`task.work_item` / `source_requirement` / `depends_on` の relation integrity を検証しない。
- 現行 contract には、workflow artifact の orphan diagnostics および task status から work item progress を導出する projection がない。
- `spec:project-artifact-model` は investigation から requirement / work item への relation を定義しているが、現行 resolver / validation contract が扱う record ID-as-ref は `ADR-*` / `SPEC-*` / `INV-*` に限られる。このため、investigation metadata に `REQ-*` / `WORK-*` を canonical reference として記載し、concept model 上の relation を MCP 経由で解決・検証する contract は未定義である。

### TASK-MCP-003-02 に渡す最小 public contract 候補

次 task では、少なくとも以下を一つの coherent MVP boundary として判断する必要がある。

- `requirement` / `work_item` / `task` を既存 record/query surface に追加し、`get_record` / `get_records` の成功対象に含めるか。
- `resolve_reference` の supported record ID-as-ref に `REQ-*` / `WORK-*` / `TASK-*` を追加するか。
- Workflow relation の最低限の整合性検査として、`requirement.work_items`、`work_item.source_requirement` / `tasks`、`task.work_item` / `source_requirement` / `depends_on` を validation 対象に含めるか。
- `spec:project-artifact-model` が既に示す investigation -> requirement / work item relation を operationalize するため、investigation の `follow_up_candidates` / `follow_up_results`、必要なら `source_refs` が `REQ-*` / `WORK-*` / `TASK-*` ID-as-ref を扱える範囲をどこまで含めるか。
- 上記を existing investigation integration と同じ record-oriented extension として扱うか、workflow 専用 query surface を新設するか。

### 初期 MVP から分離し得る論点

- Orphan diagnostics は、relation の参照解決と逆向き網羅性・未接続状態の運用診断を混在させるため、最小 relation validation と分離して後続へ送る選択肢がある。
- Task status からの progress projection は、ADR-091 が手動 checkbox を正本にしない方針を示す一方、MCP public response と aggregation semantics を新たに必要とするため、record / resolve / relation validation の最小 support とは分離して後続へ送る選択肢がある。
- Investigation の `trigger` / optional `related_*` は、現行 contract でも resolve / validation rule が未確定である。Workflow ID-as-ref を追加する場合でも、この field 群まで同時に operationalize するかは最小 scope と分離して判断できる。
- これらは本 task で除外を確定せず、TASK-MCP-003-02 が MVP 採否と ADR 要否を判断する対象とする。

### Investigation artifact 起票要否

- 独立した investigation artifact は現時点では起票しない。
- 理由は、必要性の根拠が accepted な ADR-087 / ADR-091、`spec:project-artifact-model`、完了済み WORK-MCP-002、REQ-MCP-003、および現行 MCP spec の明示的な未対応境界から直接整理でき、比較調査や外部不確実性を保存する別 artifact がなくても判断 task を開始できるためである。
- `spec:project-artifact-model` の relation を初回 evidence で明示できていなかった点は本 task 内で補正可能であり、それ自体を理由に investigation を新設しない。
- TASK-MCP-003-02 で複数の incompatible public surface 案を深く比較する必要が生じた場合は、その時点で investigation の追加要否を再判断する。

### 次工程

- 本 evidence により、`TASK-MCP-003-02` は workflow artifact MCP support の最小 public contract、investigation から workflow artifact への canonical relation operationalization、ADR 起票要否、orphan diagnostics / progress projection の MVP 内外を判断できる状態になった。
