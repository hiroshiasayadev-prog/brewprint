# TASK-MCP-002-06: Evidence / status と dogfooding 所感を反映する

- **id**: TASK-MCP-002-06
- **status**: done
- **date**: 2026-05-26
- **work_item**: WORK-MCP-002
- **source_requirement**: REQ-MCP-002
- **estimate**: 0.5d
- **depends_on**:
  - TASK-MCP-002-05
- **outputs**:
  - REQ-MCP-002 / WORK-MCP-002 の完了状態反映
  - workflow artifact 運用の dogfooding 所感
  - 必要な場合、task 粒度・format・旧 milestone 役割の work item 統合を判断する後続 ADR 候補

## Goal

Batch retrieval capability の結果を requirement / work item に反映し、新しい小粒度 task と flowchart 形式が実運用に適していたかを記録する。

## Work

- Capability の採用結果と verification evidence を REQ-MCP-002 / WORK-MCP-002 に反映する。
- 関連 ADR / spec / implementation evidence の更新漏れがないことを確認する。
- 各 task の実際の重さ、分割の過不足、blocker 表現、Mermaid task flow の有用性を短く記録する。
- 問題がなければ、task 粒度判定・work item format・旧 milestone 役割を work item が引き取る方針を formalize する ADR 起票候補を整理する。

## Done condition

- Requirement と work item の status / evidence が完了結果と矛盾しない。
- 新形式の継続採用を判断できる dogfooding 所感が残っている。
- 正式な workflow artifact ADR を起票すべきか判断できる状態になっている。

## Verification

- Close は implementation 完了だけでなく、runtime verification と artifact 更新の完了後に行う。
- 新しい workflow policy は、この task 内で暗黙に確定せず、必要なら ADR として起票する。

## Evidence

### Capability close result

- `REQ-MCP-002` は `accepted` のまま維持する。Requirement status は採用判断を表し、implementation / runtime verification の完了は `WORK-MCP-002` の close evidence で追跡する運用に従う。
- `ADR-090`、`docs/spec/design-records-mcp/{overview,schema,tools}.md`、`internal/designrecords` / `internal/designrecordsmcp` の implementation と tests、`TASK-MCP-002-05` の runtime evidence が揃ったため、`WORK-MCP-002` は done へ進められる。
- Runtime verification では `get_records` の正常系、item-level partial result、duplicate requested ID の info diagnostic、`include_body: true` の raw body path が stdio JSON-RPC 経由で成立した。

### Dogfooding 所感

- 小粒度 task の分割は有効だった。利用価値確認（T1）、採用判断と ADR（T2）、spec contract 固定（T3）、implementation（T4）、runtime verification（T5）、close review（T6）が分かれたことで、ADR-090 review で発見された schema の未確定点を実装前に解消できた。
- `decision_pending` / `design_spec_pending` / `implementation_pending` / `verification_pending` の blocker 表現は、次に何を確定すべきかを追いやすかった。
- Mermaid task flow は依存順と close 条件を共有するには有用だった。一方、この dogfooding 時点では workflow artifact が MCP 取得対象外であり、task ID から task record を取得できないため追加探索が発生した。これは physical path を relation として保持する理由ではなく、`REQ-MCP-003` が扱う ID-as-ref query / resolve support の evidence とする。
- Implementation と Go test の反復は、文書編集用の MCP 経路より Codex 実行経路へ任せる方が速いという運用上の所感を得た。設計判断・spec・status/evidence 同期は本対話経路、コード変更・format・test loop は Codex 側、という役割分担は今後の dogfooding 候補である。
- ChatGPT 側に露出する MCP tool registry/schema が新 build 後も `get_records` を表示せず、stdio runtime では新 tool が利用可能であった。これは本 capability の contract / implementation defect としては扱わないが、MCP 再登録または capability schema refresh の確認事項として残す。

### Follow-up candidate

- Close 時点では、workflow artifact の task 粒度、status 遷移、Mermaid flow、旧 milestone 役割の work item 統合、implementation/test を別実行経路へ委譲する運用について、追加事例後に ADR 化を判断する候補として残した。
- その後のレビューと明示判断により、task 粒度、status ownership、Mermaid flow、旧 milestone 役割を work item が引き取る境界、workflow relation の ID-as-ref / physical path 非対応は `ADR-091` として accepted になった。`milestone` を新しい artifact layer、metadata field、または relation として導入しない。
- Implementation/test の別実行経路への委譲は ADR-091 の確定範囲には含めず、引き続き運用上の候補として扱う。
- MCP tool registration / schema refresh の挙動が再登録後にも再現する場合は、capability 実装とは独立した運用 gap / connector behavior の requirement 候補として捕捉する。
