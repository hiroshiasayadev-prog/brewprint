# TASK-DATA-001-01: ADR-067 enum minimum の acceptance と実行境界を確定する

- **id**: TASK-DATA-001-01
- **status**: done
- **date**: 2026-05-29
- **work_item**: WORK-DATA-001
- **source_requirement**: REQ-DATA-001
- **estimate**: 0.5d
- **depends_on**:
- **outputs**:
  - ADR-067 の enum minimum acceptance 判断と status / evidence 更新
  - WORK-DATA-001 の F1 execution boundary 確認記録

## Goal

`REQ-DATA-001` が採用した F1 boundary を実行可能にするため、proposed の `ADR-067` について、M15 critical path に含める enum minimum のみを accepted として確定し、後続 spec / implementation / YAML migration が拡張 scope を暗黙に抱えない状態にする。

## Work

- `ADR-067` の現行本文と `REQ-DATA-001` / `WORK-DATA-001` の F1 boundary を照合する。
- ADR-067 §7 の初期移行対象が、以下の 3 enum model と初期 field migration に閉じていることを確認する。
  - `mcp_object_type`
  - `mcp_diagnostic_severity`
  - `reference_tree_direction`
  - `object_selector.object`
  - `object_ref.object`
  - `diagnostic.severity`
  - `get_reference_tree_request.direction`
  - `get_reference_tree_response.direction`
- ADR-067 を acceptance に進めるために必要な status / Evidence / boundary 記述の更新を行う。
- `get_references.direction`、`reference.direction`、object-dependent `kind`、`impact_severity`、`impact_fixability`、helper model、tagged union、DAG TypeRef hint を本 task の acceptance boundary に追加しない。
- ADR-074 は enum render の必須前提ではなく、v1.1 では DAG TypeRef hint を後続へ送る境界であることを確認する。

## Done condition

- ADR-067 の enum minimum が、`REQ-DATA-001` / `WORK-DATA-001` の F1 boundary と矛盾せず accepted として扱える状態になっている。
- Initial enum migration scope と明示的 follow-up scope が混同されていない。
- `TASK-DATA-001-04` が ADR-067 の判断未確定を理由に停止しない状態になっている。

## Verification

- ADR-067 の status / initial migration scope / M15 への影響記述を読み戻す。
- `REQ-DATA-001` / `WORK-DATA-001` の含有範囲・除外範囲と差分がないことを照合する。
- ADR-070 / ADR-073 / ADR-074 / ADR-078〜080 が enum minimum acceptance の必須 dependency として追加されていないことを確認する。

## Evidence

- `ADR-067` の決定本文を `REQ-DATA-001` / `WORK-DATA-001` の F1 boundary と照合した。
- `ADR-067` の enum capability 自体、初期 3 enum model、および初期 5 field migration は F1 boundary と一致することを確認した。
- `ADR-067` を accepted に更新し、3 enum model 定義追加と 5 field migration を同一実行単位として明記した。
- `get_references.direction`、`reference.direction`、object-dependent `kind`、`impact_severity`、`impact_fixability`、helper model、tagged union、DAG TypeRef hint、MCP / state identity、notes retreat 完全解消を M15 blocker に追加しない境界を明記した。
- Design Records MCP validation: `ok: true` を確認した。info diagnostic は既存 `INV-DATA-001` / `INV-DATA-002` の follow-up candidate 表記に関するものであり、本 task の acceptance boundary を block しない。
