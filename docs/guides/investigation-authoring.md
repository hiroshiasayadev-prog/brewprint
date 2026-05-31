# Investigation Authoring Guide

## Abstract

Investigation artifact を起票・更新するときの実践ルールを定める。Investigation は複雑な変更における調査結果、根拠、影響範囲、未確定点、選択肢、後続 artifact 候補を所有する。

## Migration Note

Maintainer note only. This is not an instruction to read another file, and it is not part of the public guide retrieval contract.

Extracted from:

- `docs/investigations/README.md`

## Purpose

Investigation は、複雑な変更における調査結果、根拠、影響範囲、未確定点、選択肢、後続 artifact 候補を保存する artifact である。

Investigation は、requirement / work item / task / ADR / spec / internal design / coverage / 別 investigation の起票・更新前に必ず必要な gate ではない。

## ID and Layout

Investigation ID:

```text
INV-<DOMAIN>-NNN
```

File path:

```text
docs/investigations/<domain>/INV-<DOMAIN>-NNN-<slug>.md
```

`DOMAIN` は uppercase の短い domain label とする。

`NNN` は domain ごとに 001 から始まる3桁ゼロ埋め連番とする。

Investigation ID は、ADR number、requirement ID、work item ID、task ID、旧 M-series label、coverage mapping ID などとは結合しない。

## Format

```markdown
# INV-<DOMAIN>-NNN: <title>

- **status**: investigating / concluded / superseded
- **date**: YYYY-MM-DD
- **trigger**: <起点 artifact または起票理由>
- **scope**: <短い調査スコープ>
- **non_scope**: <短い非スコープ>
- **source_refs**:
  - <artifact ID-as-ref or semantic ref>
- **follow_up_candidates**:
  - <candidate artifact ID-as-ref / semantic ref / human-readable candidate, or なし>
- **follow_up_results**:
  - <artifact ID-as-ref or semantic ref>

## 調査スコープ

## 非スコープ

## 背景

## 調査したもの

## 調査項目ごとの確認結果

## 横断的な観測事実

## 後続判断に渡す候補

## 推奨案

## 後続 artifact 候補

## 未確定点
```

## Required Metadata

- `status`
- `date`
- `trigger`
- `scope`
- `non_scope`
- `source_refs`
- `follow_up_candidates`

## Optional Metadata

- `supersedes`
- `related_requirements`
- `related_work_items`
- `related_adrs`
- `related_specs`
- `related_internal_design`
- `related_coverage`
- `follow_up_results`

Optional metadata は、該当する情報がある場合のみ書く。空 field を義務付けない。

## Status

| status | meaning |
|---|---|
| `investigating` | 調査中 |
| `concluded` | 調査結果がまとまり、後続判断に渡せる状態 |
| `superseded` | 後続 investigation または別 artifact により置き換えられた状態 |

`concluded` は、後続 artifact の採用判断や実装完了を意味しない。

`proposed` は ADR status と混同しやすいため、investigation status には使わない。

## Responsibility Boundary

Investigation が所有するもの:

- 調査結果
- 根拠
- 影響範囲の仮説
- 未確定点
- 選択肢
- 後続 artifact 候補
- 調査者の推奨案

Investigation が所有しないもの:

- 決定
- 現行仕様
- 要求そのもの
- 横断進捗
- 完了状態
- 具体的な作業手順

## Recommendation Boundary

Investigation は決定を所有しない。

ただし、調査者の見立てとして `推奨案` を書いてよい。その場合は「〜と考えられる」「〜が妥当と見られる」のように、判断ではなく調査結果に基づく候補であることが分かる表現にする。

採用判断は後続 ADR / README / doc-policy / task file 等に委ねる。

## Canonical Reference Boundary

`source_refs` / `follow_up_results` / artifact reference として記載された `follow_up_candidates` には、MVP の canonical reference として record ID-as-ref または active `spec:` semantic ref を使う。

MVP の record ID-as-ref:

- `ADR-*`
- `SPEC-*`
- `INV-*`
- `REQ-*`
- `WORK-*`

`TASK-*` は workflow artifact 間 relation と direct resolver input では support されるが、investigation metadata の canonical reference には含めない。

Physical path は canonical reference として使わない。

## Split Boundary

Investigation の調査中に別領域の調査が必要になった場合、別 investigation を起票してよい。

ただし、別 investigation の起票は必須ではない。軽微な追加確認は元 investigation 内に留めてよい。
