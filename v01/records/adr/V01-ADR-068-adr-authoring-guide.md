# V01-ADR-068: ADR authoring guide の標準化

- **status**: accepted
- **date**: 2026-05-09

> このADRは起票時点での決定を記録したスナップショットである。
> 現在の docs 運用方針は doc-policy.md および adr-authoring-guide.md を参照すること。

## 背景

brewprint は ADR と spec を分離して運用している。
V01-ADR-050 および `docs/doc-policy.md` により、責務は以下のように整理されている。

- spec: 現行仕様の唯一の正
- ADR: 設計判断の根拠記録
- task file: 具体作業・順序・完了条件の管理
- UC docs / fixture: 実例・検証材料・gap 発見ログ

一方で、ADR の執筆実務では、判断根拠として実装・UC・fixture・一時的な YAML shape を参照する場面が増えている。
特に M15 Phase C では、UC-002 self-hosting の実例から enum / discriminated object / inline struct / TypeRef depth 方針などの設計判断を抽出している。

V01-ADR-067 起票時には、以下の問題が顕在化した。

- UC-002 の現行 YAML shape を根拠として enum model の必要性を説明した
- しかし、その YAML は enum 導入後の migration によってすぐ更新されうる
- 実例をそのまま書くと、将来読まれたときに「この YAML shape が仕様固定だった」と誤読される
- 一方で、実例を消すと「なぜ enum が必要なのか」という設計判断の根拠が弱くなる
- さらに、後続 task list を ADR 本文に入れるか task file に外すかの判断が議論となった

この問題は V01-ADR-067 固有ではない。
今後も、実例・実装・fixture から設計判断を抽出する ADR が増える見込みである。

したがって、ADR の書き方に関する実践的な標準を `doc-policy.md` とは別に整備する。

## 決定

### 1. `docs/adr-authoring-guide.md` を新設する

ADR 執筆時の実践ルールをまとめる文書として、`docs/adr-authoring-guide.md` を新設する。

この guide は、ADR のフォーマットそのものではなく、ADR を書くときの判断・責務境界・アンチパターンを扱う。
また、brewprint 言語仕様ではなく docs 運用補助文書であるため、`docs/spec/` ではなく `docs/` 直下に置く。

対象読者:

- ADR を起票・レビューする人間
- ADR / spec / task / UC docs を編集する AI assistant
- 過去ADRを読んで後続設計判断をする reviewer

### 2. `doc-policy.md` から authoring guide を参照する

`docs/doc-policy.md` は引き続きドキュメント運用方針の上位文書とする。

`doc-policy.md` には、以下を残す。

- docs 全体の上位運用方針
- spec / ADR の基本責務境界
- ADR フォーマット
- Evidence の基本的な書き方
- セッション開始時に読むべき文書としての役割

`docs/adr-authoring-guide.md` には、以下を置く。

- 実例由来ADRの書き方
- ADR / spec / task / UC docs の責務境界
- 観測事実と設計判断の分離
- アンチパターン
- 文例集
- reference example

ADR / spec / task / UC docs の4分割表は authoring guide を一次情報とする。
`doc-policy.md` には表を転記せず、ADR 執筆時は authoring guide も参照する旨を短く追記する。

### 3. authoring guide に「実例は evidence、仕様ではない」原則を入れる

実例・実装・fixture から設計判断を抽出する ADR では、以下を区別する。

- 具体例は起票時点の evidence であり、恒久仕様ではない
- ADR が判断根拠として扱うのは、具体 shape そのものではなく、そこから抽出された設計上の性質である
- fixture migration の範囲・順序は ADR ではなく task file / UC task file で管理する

具体的な推奨文例は authoring guide 本文で示し、guide の更新として改善できるものとする。

### 4. authoring guide に「観測事実」と「設計判断」の分離を入れる

ADR では、実例から見つかった観測事実と、それに基づく設計判断を分けて書く。

この分離により、fixture の現行 shape が変わっても ADR の判断根拠が腐りにくくなる。

具体的な良い例 / 悪い例は authoring guide 本文で示し、guide の更新として改善できるものとする。

### 5. authoring guide に ADR / spec / task / UC docs の責務境界を入れる

ADR は Architecture Decision Record であり、設計判断の記録である。
ADR 本文は、その判断を理解するために必要な背景・決定・理由・却下案・影響範囲に集中する。

authoring guide では、文書間の関係を「所有」「参照」「影響範囲」の3概念で整理する。

- 所有: その文書が一次情報として保持し、更新責任を持つ内容
- 参照: 他文書が一次情報として持つ内容を、必要に応じてリンク・言及すること
- 影響範囲: その決定により更新・実装・検証が必要になる領域

ADR は、spec / task file / UC docs / fixture / impl notes など、他責務 artifact の詳細を所有しない。
必要な場合は、影響範囲と参照先を示すに留める。

| 文書 | 所有するもの | 参照のみ示すもの |
|---|---|---|
| ADR | 設計判断、背景、理由、却下案、影響範囲 | 現行仕様本文、具体task、fixture migration詳細、実装手順 |
| spec | 現在の正しい仕様 | ADRの判断背景（由来注記） |
| task file | 作業項目、順序、完了条件、実装・migration checklist | 設計判断の長い理由、現行仕様本文 |
| UC docs / fixture | 実例、coverage、gap発見ログ、migration状態 | 汎用仕様そのもの、設計判断の最終根拠 |

この境界により、ADR は時間が経っても「なぜその判断をしたか」を保持し、現行仕様・作業状態・実例の変化はそれぞれの責務文書で更新できる。

### 6. authoring guide に reference example セクションを設ける

`docs/adr-authoring-guide.md` には、実例由来ADRの書き方を確認できる reference example セクションを設ける。

V01-ADR-067 は、実例由来の設計判断を ADR として整理する過程で、以下の学びを生んだ。

- UC-002 の実例は enum model 導入の重要な evidence だった
- しかし、実例の YAML shape は migration によって変わりうる
- そのため、V01-ADR-067 には「判断根拠としての UC-002 実例の扱い」を明記した
- 後続 task list は ADR から外し、M15 / UC-002 task file で追跡する方針にした

そのため、V01-ADR-067 は authoring guide 初版で参照しうる初期 example の一つとして扱う。
将来より適切な例が増えた場合、guide 側で reference example を追加・差し替えできる構造にする。

## 理由

### なぜ doc-policy.md に全部入れないか

`doc-policy.md` はドキュメント運用方針の上位文書であり、既に広範囲の運用ルールを含んでいる。
ADR 執筆技術の詳細まで詰め込むと、doc-policy が肥大化し、起動時に読むべき最低限の運用方針と実践的な執筆ノウハウが混ざる。

ADR authoring guide を分離することで、以下が得られる。

- doc-policy は上位方針として保てる
- ADR 起票・レビュー時に参照しやすい
- 実例・アンチパターン・文例を増やしやすい
- 将来、AI assistant 向け skill / playbook として再利用しやすい

### なぜ今標準化するか

M15 Phase C では、UC-002 self-hosting の実例から設計判断を抽出する ADR が続く可能性が高い。

候補:

- enum model
- discriminated object
- inline struct
- TypeRef depth / lint 方針
- optional / required 制約
- recursive struct / union list / arbitrary JSON object の扱い

これらはいずれも、実例・fixture・実装制約から必要性が見える一方で、現行実例をそのまま仕様固定すると陳腐化しやすい。

V01-ADR-067 の経験を標準化することで、以後の ADR で同じ混乱を繰り返しにくくなる。

### なぜ ADR として起票するか

ADR authoring guide の新設は、単なる文書追加ではなく、ADR / spec / task / UC docs の責務境界に影響する運用上の設計判断である。

また、doc-policy.md との関係を明示し、将来の AI assistant / reviewer が「なぜ別文書があるのか」を理解できるようにする必要がある。
そのため ADR として判断を記録する。

### なぜ ADR と他文書の責務境界を明記するか

ADR は設計判断の根拠記録であり、作業管理文書でも現行仕様本文でもない。
ADR 本文に他文書の責務に属する詳細を抱え込みすぎると、実装状況・仕様・fixture の変化により ADR が stale に見える。

一方で、ADR の「影響」には、どの spec / implementation / UC に作業が波及するかを書く必要がある。
したがって、ADR には影響範囲と参照先を書き、詳細は各責務の文書に置く、という境界を明文化する。

### なぜ抽象 artifact model / 命名 / MCP化を本ADRで扱わないか

本ADRは brewprint 固有の docs 運用として authoring guide を整備するものである。
ただし、「所有」「参照」「影響範囲」という責務原則は、将来の抽象 artifact model や MCP 化と整合しうる抽象軸として採用する。

一方で、DecisionRecord / Spec / TaskTracker 等の一般化された artifact 型名、他プロジェクト適用、MCP schema 化は本ADRでは扱わない。
これらは別ADRで検討する。

## 却下した代替案

### 代替案A: 何も標準化しない

- 利点: 文書追加が不要
- 欠点: V01-ADR-067 と同じ混乱が今後も起きる。実例を仕様固定のように書く ADR、task list を抱え込む ADR が増える

→ 却下。

### 代替案B: `doc-policy.md` に ADR 執筆詳細をすべて追記する

- 利点: 参照先が1つで済む
- 欠点: doc-policy が肥大化する。運用方針と執筆ノウハウが混ざる。文例・アンチパターンを増やしづらい

→ 却下。doc-policy には参照だけを置き、詳細は authoring guide に分離する。

### 代替案C: `docs/spec/adr.md` として spec 配下に置く

- 利点: 仕様文書として扱える
- 欠点: ADR authoring guide は brewprint 言語仕様ではなく、docs 運用の補助文書である。spec-first の「現行仕様の唯一の正」と混ざる

→ 却下。`docs/` 直下の運用文書とする。

### 代替案D: AI assistant の外部 skill としてだけ管理する

- 利点: AI作業に直接効く
- 欠点: repo 内に判断根拠が残らない。人間 reviewer や別環境の assistant が参照できない

→ 却下。repo 内文書として管理し、必要に応じて外部 skill / prompt へ派生させる。

## 影響

### docs への影響

- `docs/adr-authoring-guide.md` を新設する
- `docs/doc-policy.md` に authoring guide への参照を追記する

### ADR への影響

今後の ADR 起票時は、必要に応じて authoring guide に従う。
特に以下のケースでは参照する。

- 実例・fixture・実装から設計判断を抽出する ADR
- 後続 migration が発生する ADR
- spec / task / UC docs の責務境界が曖昧になりやすい ADR
- 却下案や影響範囲が長くなりやすい ADR

既存 ADR を一括修正することはしない。
V01-ADR-050 / doc-policy.md の漸進移行方針に従い、触れたタイミングで必要に応じて改善する。

### spec への影響

なし。

本ADRは brewprint 言語仕様や MCP schema を変更しない。

### M15 への影響

M15 Phase C では、V01-ADR-067 以降の設計判断で authoring guide を参照する。
特に、UC-002 実例から設計上の性質を抽出し、fixture migration task と ADR 判断を分離する際に使う。

### 実装への影響

なし。

## Evidence

- commit: 693e3c0
- impl commit: 該当なし
- 参考: V01-ADR-050 spec-first documentation policy, V01-ADR-067 enum model 導入時のレビュー経験
