# ADR Authoring Guide

## Abstract

ADR を起票・レビュー・更新するときの実践ルールを定める。ADR は設計判断の履歴を所有し、現行仕様本文や作業 checklist を所有しない。

## Migration Note

This section is a maintainer note about phase-1 guide extraction history.
It is not an instruction to read the original file, and it is not part of the public authoring guidance retrieval contract.

Extracted from:

- `docs/adr-authoring-guide.md`

## Purpose

ADR は Architecture Decision Record であり、設計判断の記録である。

ADR には、その判断を理解するために必要な以下を書く。

- 背景
- 決定
- 理由
- 却下した代替案
- 影響範囲
- evidence

ADR は、現行仕様そのもの、作業 checklist、fixture migration 状態、実装引継ぎメモの代替文書ではない。

## Format

ADR file path:

```text
docs/adr/NNN-title.md
```

ADR body:

```markdown
# NNN: Title

- **status**: proposed / accepted / superseded
- **date**: YYYY-MM-DD
- **supersedes**: ADR-NNN
- **migrated_to_spec**: YYYY-MM-DD

> This ADR is a snapshot of the decision at the time it was written.
> Refer to spec for the current behavior.

## 背景

## 決定

## 理由

## 影響

## Evidence
- commit: <commit hash>
- impl commit: <implementation commit hash or tbd>
```

`id` は H1 の number から導かれる。ADR metadata は Design Records MCP によって読まれるため、field value に注釈を混ぜない。

MCP create note:
The Markdown above is the full artifact file shape.
When calling `propose_record_create` with `fields` plus `body`, put metadata in `fields` and make `body` content sections only, starting at `## 背景`.
Do not include the H1, metadata block, metadata `id`, or guessed server-resolved ID in that MCP `body`.

## Status

| status | meaning |
|---|---|
| `proposed` | 議論中・まだ覆りうる |
| `accepted` | 確定。変更する場合は新しい ADR で supersedes する |
| `superseded` | 旧ADR。新 ADR への置換済み |

## Evidence Rules

- `commit` / `impl commit` は git log から拾う。
- ADR 起票と実装反映が同 commit なら同一 hash でよい。
- artifact reference は可能な限り canonical ID-as-ref を使う。
- physical path は補助的な所在説明に限り、canonical reference として扱わない。
- URL / retrieved date のような長い参考情報は原則不要。

## Responsibility Boundary

ADR に書くもの:

- なぜその判断が必要になったか
- 何を決めたか
- なぜその形を選んだか
- どの代替案を却下したか
- どの spec / implementation / UC / task に影響するか
- どの証拠・実例・過去 ADR に基づくか

ADR に抱え込まないもの:

- 現在の完全な仕様本文
- 作業 checklist
- migration の細かい順序
- fixture の作業状態
- 実装手順の詳細
- 判断前の探索ログ、影響範囲調査、選択肢比較、未確定論点の蓄積

## Spec Boundary

ADR は起票時点の決定を記録する。

Spec は現在の正しい仕様を記録する。

ADR に仕様案や具体例を書いてよいが、それは起票時点の snapshot である。設計が accepted になり spec に反映されたら、現行仕様は spec を参照する。

## Investigation Boundary

ADR は設計判断とその理由を所有する。

Investigation は判断前の調査結果、根拠、影響範囲の仮説、未確定点、選択肢、後続 artifact 候補を所有する。

複雑な変更では、ADR に探索ログを抱え込まず investigation を参照する。

## Anti-patterns

- fixture shape をそのまま仕様判断にする。
- ADR に作業 checklist を置く。
- ADR に現行仕様本文を長く抱え込む。
- 実装手順を ADR に書きすぎる。
