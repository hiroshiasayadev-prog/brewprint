# 085: investigation artifact boundary

- **status**: accepted
- **date**: 2026-05-18
- **depends_on**: ADR-050, ADR-068, ADR-081, ADR-083, ADR-084
- **supersedes**:
- **migrated_to_spec**:

> このADRは起票時点での決定を記録したスナップショットである。
> 現在の仕様は spec を参照すること。

## 背景

ADR-081 により `docs/requirements/` が導入され、要求・不足・要望・spec gap 候補を捕捉する layer が定義された。
ADR-083 により、requirements / work-items / tasks / internal design / coverage / YAML の責務境界が整理された。
ADR-084 により、semantic trace MVP の対象 artifact と scope が整理された。

しかし、M18 semantic traceability foundation の検討中に、requirement から work item / task へ落とす前段階として、影響範囲調査、根拠整理、問題点の洗い出し、設計判断候補の列挙を保存する artifact が不足していることが分かった。

requirements は「何が必要か」を所有するが、調査過程や選択肢比較を長く持つ場所ではない。
work item は layer 別の横断進捗と完了状態を所有するが、なぜその影響範囲になったのかを説明する調査レポートではない。
tasks は具体的な作業手順・チェックリストを所有するが、task 化の根拠を所有しない。
ADR は設計判断の理由を記録するが、判断前の探索ログや未確定論点の集積場所ではない。

このままでは、複雑な変更において「なぜこの task が必要なのか」「どの artifact が影響対象なのか」「どの設計判断が未確定なのか」が会話ログや一時メモに散り、後続作業者が再調査を強いられる。

特に traceability / internal design / coverage / Design Records MCP のようなメタ設計では、仕様そのものを作る前に既存 ADR、milestone、spec、実装状態を調査し、どの設計判断が必要かを整理する必要がある。
その調査結果を docs として積み重ねられる置き場が必要である。

## 決定

brewprint project layout に、調査レポートを置く artifact layer として `docs/investigations/` を導入する。

investigation は、requirement / bug / design gap / milestone planning / observed issue を起点に、以下を記録する artifact とする。

- 何を調べたか
- 調査で分かった問題点
- 影響しうる artifact / layer
- 未確定点
- 設計判断が必要な論点
- 採りうる選択肢
- 後続の ADR / requirement / work item / task 候補

investigation は決定を所有しない。
現行仕様を所有しない。
要求そのものを所有しない。
横断進捗や完了状態を所有しない。
具体的な作業手順を所有しない。

investigation は、requirement / work item / task / ADR / spec / internal design / coverage / 別 investigation だけでは、根拠・影響範囲・判断候補を安全に保存できない場合に起票する。

investigation は、requirement / work item / task / ADR / spec / internal design / coverage / 別 investigation の起票・更新前に必ず必要な gate ではない。
複雑な変更で調査結果、根拠、影響範囲、判断候補を後続 artifact へ保存する必要がある場合にのみ使う。

investigation は、後続 artifact の起点になりうる。
ただし investigation は、後続 artifact の採用判断、現行仕様、作業進捗、完了状態を所有しない。

investigation が記録する「影響範囲」は、調査上の仮説・根拠・不確実性・選択肢を含む分析結果である。
work item が所有する layer 別の処理状態・完了管理・横断進捗の代替ではない。

`docs/investigations/` の front matter、本文フォーマット、status 語彙、ID 規則、MCP index 方針、lifecycle は本ADRでは確定しない。
初回 investigation `INV-DOCS-001` で調査し、必要に応じて後続 spec / ADR / doc-policy 更新で確定する。

ADR-085 の根拠となった初回調査は、`docs/investigations/` 導入後に `INV-DOCS-001` として retroactive に保存してよい。
これは bootstrap exception であり、以後の investigation 起票ルールを無限再帰させるものではない。

## 理由

### なぜ investigation layer が必要か

複雑な変更では、requirement を見つけた時点でただちに task へ落とすと、影響範囲や設計判断候補が失われる。

一方で、調査内容を requirement に入れると requirement が分析文書化する。
work item に入れると work item が調査レポート化する。
task に入れると task が根拠説明まで背負う。
ADR に入れると ADR が探索ログ化する。

そのため、判断前の調査結果を保存する artifact として investigation を分離する。

### なぜ詳細仕様を本ADRで確定しないか

investigation artifact は、実際に初回 investigation を書いてみないと、必要な front matter、status、ID、本文構成、参照関係の粒度が分からない。

本ADRで詳細 schema まで固定すると、実例に基づかない過剰設計になる危険がある。
そのため、本ADRでは artifact layer の導入と責務境界のみを決め、詳細は `INV-DOCS-001` に委ねる。

### なぜすべての task に investigation を要求しないか

investigation を常に必須にすると、軽微な修正や明白な作業まで儀式化し、運用負荷が増える。

investigation は、根拠・影響範囲・判断候補が requirement / work item / task だけでは安全に保存できない場合に使う。
単純な typo 修正、既に影響範囲が明確な実装作業、既存 task の細分化には必須ではない。

## 却下した代替案

### 代替案A: requirement に調査内容を含める

- 利点: 起点と調査結果を1ファイルに置ける
- 欠点: requirement が「何が必要か」ではなく分析文書になり、責務が膨らむ

→ 却下。requirement は要求・不足・要望を所有し、調査レポートは investigation が所有する。

### 代替案B: work item に調査内容を含める

- 利点: 横断進捗と影響範囲を同じ場所で扱える
- 欠点: work item が調査レポート化し、layer 別進捗と根拠整理が混ざる

→ 却下。work item は処理状態と完了管理を所有し、調査根拠は investigation が所有する。

### 代替案C: task file に調査結果を書く

- 利点: 作業手順と根拠を近くに置ける
- 欠点: task が具体作業ではなく判断根拠まで背負い、再利用しにくくなる

→ 却下。task は具体的な作業手順・チェックリストを所有する。

### 代替案D: ADR に調査ログを含める

- 利点: 設計判断と調査経緯をまとめられる
- 欠点: ADR が判断前の探索ログまで所有し、決定記録として重くなる

→ 却下。ADR は決定と理由を所有する。未確定論点や影響範囲調査は investigation が所有する。

## 影響

### docs への影響

`docs/investigations/` を新設する。

初回 investigation として `INV-DOCS-001` を作成し、以下を調査する。

- investigation front matter の暫定形
- required / optional field
- status 語彙
- ID 規則
- 調査 scope の明示方法
- 起点 artifact と後続 artifact の記録方法
- requirement / work item / ADR / task との参照方法
- どこまでを metadata に置き、どこからを Markdown 本文に置くか
- 将来的に Design Records MCP または別 MCP index 対象にするか

### ADR-083 への影響

ADR-083 の artifact placement decision rule は、investigation layer を含む形で後続更新が必要になる。
特に、requirements と work-items の間に「影響範囲調査・根拠整理・選択肢比較」を置く分岐を追加する必要がある。

work item が所有する「影響範囲」は、layer 別の処理状態・完了管理を指すものとして明確化する必要がある。
investigation が所有する「影響範囲」は、調査・根拠・不確実性・選択肢を含む分析結果である。

### ADR-081 / ADR-084 への影響

requirements から work item / task / ADR へ進む前段階として、investigation を参照できるようになる。

semantic trace MVP 自体の仕様化に先立ち、M18 の問題点、既存ADRとの関係、設計判断候補を investigation として積み重ねられる。

### docs/doc-policy.md / adr-authoring-guide.md への影響

`docs/doc-policy.md` と `docs/adr-authoring-guide.md` は、ADR-081〜085 および traceability spec の進捗を踏まえて更新する必要がある。

ただし、investigation の詳細 format が `INV-DOCS-001` で調査されるまでは、doc-policy への詳細反映は行わない。

### Design Records MCP への影響

MVP直後の必須変更ではない。

将来的に investigation を index / query / validate 対象にする場合は、`kind: investigation` または別 MCP interface を検討する必要がある。
ただし、本ADRでは MCP contract を確定しない。

## Evidence

- commit: 31bfe14
- impl commit: tbd
- 参考: ADR-081 project requirements layer と semantic traceability, ADR-083 project artifact boundary と YAML as primary implementation source, ADR-084 semantic trace MVP scope と artifact boundary, M18 semantic traceability foundation 検討中の会話
