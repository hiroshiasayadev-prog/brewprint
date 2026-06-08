# V01-ADR-089: LLM-first 設計原則と intent transfer 境界

- **status**: accepted
- **date**: 2026-05-24
- **depends_on**: V01-ADR-081, V01-ADR-083, V01-ADR-087, V01-ADR-088
- **supersedes**:
- **migrated_to_spec**:

> このADRは起票時点での決定を記録したスナップショットである。
> 現在の仕様は spec を参照すること。

## 背景

brewprint は、brewprint DSL YAML により対象 system / design model の意味構造を記述し、人間向けには Mermaid 等の derived view、LLM 向けには MCP 経由の structured query interface を提供する project である。

既存の `docs/doc-policy.md` は、spec-first、artifact ごとの ownership boundary、MCP 経由のファイル操作、canonical reference、Design Records MCP による探索・resolve / validation といった運用方針を既に定めている。V01-ADR-081 / V01-ADR-083 / V01-ADR-087 / V01-ADR-088 も、requirements、spec、investigation、internal design、YAML、semantic trace、MCP の責務を分離し、MVP を canonical reference resolution foundation に限定した。

しかし、これらの判断を束ねる理由、すなわち「なぜ brewprint は artifact と interface をこの粒度で構造化するのか」は、project-wide の設計原則として明示されていなかった。

本 ADR の起点となった V01-INV-DOCS-006 では、OpenFastTrace (OFT) / Doorstop のような既存 tool を評価する際、機能適合性や artifact ownership だけでなく、LLM が semantic に読み取り・検証し、必要に応じて安全に操作できる形へ接続する総コストを評価しなければ、brewprint の目的に照らした比較にならないことが整理された。

本判断の前提となる問題認識は以下の二点である。

1. **LLM の会話文脈は project state の永続的な source of truth にならない。**
   別 session や別 agent が作業する際、以前の暗黙文脈を前提に一貫した判断を期待できない。継続すべき意図、判断、現行 contract、参照関係は、artifact と machine-readable interface に外部化する必要がある。
2. **人間から渡される意図は、初期状態では曖昧であり得る。**
   ユーザーが毎回すべての ownership、前提、影響範囲、完了条件を精密に指定することを前提にしない。LLM が意図を適切な artifact に分解し、根拠を確認し、確定した contract と未確定の要求・調査・判断候補を混同しないための構造が必要である。

このため、spec-first、artifact ownership、canonical reference、MCP query interface、および段階的な validation foundation は、付随的なドキュメント事務ではなく、LLM と人間が設計を協働するための product-level constraint である。

## 決定

### 1. brewprint は LLM-first を設計原則として採用する

brewprint は、**人間が意図と最終判断を担い、LLM が構造化 artifact の起草・探索・検証・許可された操作を担うことを主たる協働モデルとする**。

ここで LLM-first は、人間が artifact を閲覧・レビュー・修正してはならないことを意味しない。Human review と最終意思決定は引き続き必要である。一方、project の interface、metadata、ownership、validation は、LLM が別 session でも正確に文脈を再構成し、誤った layer に情報を書き込まず、必要な根拠を辿れることを主要要件として設計する。

### 2. Artifact ownership と spec-first は LLM-first の中核制約として維持する

LLM-first における intent transfer は、会話上の意図を即座に implementation へ変換することではない。意図を、責務の異なる artifact を通じて検証可能な形へ精密化することである。

起票時点での主要 ownership は以下の通りである。

| artifact / mechanism | LLM-first における役割 |
| --- | --- |
| `docs/requirements/` | 要求・不足・要望・spec gap 候補を、現行仕様と混同せず捕捉する |
| `docs/investigations/` | 根拠、不確実性、選択肢、後続判断候補を保存する |
| `docs/adr/` | 人間が採用した設計判断と理由を保存する |
| `docs/spec/` | 現在有効な contract の唯一の正を保持する |
| brewprint DSL YAML | 対象 design model の primary implementation source を保持する |
| semantic ref / ID-as-ref | path 変更に依存せず LLM が概念と record を辿る identity を提供する |
| MCP / Design Records MCP | LLM が構造化された query、resolve、validation を行う interface を提供する |

この責務分離により、LLM は曖昧な intent、調査結果、採用判断、現行 contract、implementation source を同一の情報として扱わずに済む。

### 3. 現行 MVP の read-first boundary は維持し、まず dogfooding で十分性を検証する

LLM-first の採用は、直ちに MCP writer tool、controlled mutation workflow、追加 provenance metadata、または新たな intent artifact layer を導入する判断ではない。

V01-ADR-087 / V01-ADR-088 および M19 が扱う canonical read / resolve / validation foundation は、以下を確認するための意図的な first step と位置付ける。

- 現行 front matter、semantic ref、canonical reference、artifact ownership が、LLM 協働で誤読を防ぎ、必要な根拠を辿るために十分か
- LLM が design record / investigation / spec を探索・検証する際に、どの metadata、query、diagnostic が不足するか
- 実運用で、手作業の反復、曖昧な承認境界、修正困難な診断、または安全な write の必要性が具体的に発生するか

Writer tool、controlled write、追加の ambiguity / provenance / approval mechanism は、read-first foundation の dogfooding により concrete requirement が確認された場合に、後続 requirement / ADR / spec / work item で再判断する。

### 4. 外部 tool の採用判断には LLM-mediated 利用コストを含める

OFT、Doorstop、または将来の外部 tool を build / buy / adapt 候補として評価する場合、機能の存在または成熟度だけで判断しない。

少なくとも以下の評価軸を含める。

| 評価軸 | 判断すること |
| --- | --- |
| functional fit | 必要な coverage / trace / lifecycle 能力を担えるか |
| ownership fit | Brewprint の requirements / spec / ADR / investigation / YAML の source-of-truth boundary を壊さないか |
| agent read I/O | LLM が必要な item、relation、diagnostic を semantic に取得・解釈できるか |
| diagnostic usability | validation failure を LLM が安全な修正候補へ変換しやすいか |
| identity mapping | `spec:` 等の canonical identity と外部 item identity が二重 owner にならないか |
| testability / maintenance cost | connector、projection、version 追従、failure handling の維持費が妥当か |

Write I/O、approval、追加 provenance の適合性は、read-first dogfooding によりそれらが必要と確認された場合の後段評価軸とする。

LLM-first は、外部 tool を排除してすべて自作する原則ではない。外部 backend を採用する場合でも、LLM が semantic に利用できる形へ接続する総コストと正確性を、native implementation と比較する原則である。

### 5. LLM-first は intent transfer problem を設計対象として扱う

brewprint は、ユーザーの曖昧な意図を直接 implementation へ飛ばすのではなく、artifact と tool boundary を通じて、レビュー可能かつ検証可能な設計状態へ変換する問題を扱う。

この intent transfer において、起票時点では既存 artifact model を基盤とする。新しい intent artifact layer は導入しない。Requirements / investigations / ADR / spec / YAML / MCP の運用を dogfooding し、意味の混同、欠落、誤読、修正不能な diagnostic などの実在する不足が確認された場合にのみ、metadata や tool contract の refinement を判断する。

## 理由

### なぜ LLM-first を明示するか

既存 policy と ADR は、spec-first、ownership、canonical reference、MCP といった個別の制約を既に持つ。しかし原則が明示されないままでは、新しい artifact や外部 tool を評価するとき、それらの制約が不要な手続きに見えたり、human-oriented な既存 workflow へ寄せることで LLM 利用時の再構成コストを後から負ったりする可能性がある。

LLM-first を明示することで、これらの制約が「LLM が曖昧な指示と session 非永続性の下でも、正しい artifact boundary を保って作業するための設計」であることを、後続判断の根拠として利用できる。

### なぜ read-first MVP を維持するか

LLM が主要な協働主体であることと、最初から write interface を設計することは同義ではない。

Brewprint は既に、front matter、semantic ref、canonical reference、artifact ownership、MCP query / validation という構造を持つ。まずこれらを最低限の実装と実際の運用で dogfooding し、どの場面で不足するかを観測する方が、未観測の writer workflow や追加 metadata を先取りするより合理的である。

したがって、M19 の read / resolve / validation foundation は LLM-first に反する縮小ではなく、LLM-first の必要十分性を実地に検証するための段階的実装である。

### なぜ外部 tool を機能だけで選ばないか

External tool が coverage や requirements management の機能を持っていても、LLM が利用するために identity bridge、diagnostic projection、artifact translation、connector maintenance を大きく再実装する必要があるなら、brewprint の目的に対する導入利益は小さくなり得る。

逆に、外部 tool が model-facing interface を native に持たなくても、安定した adapter を小さく構築でき、ownership boundary と canonical identity を維持できるなら、採用を排除する理由にはならない。

このため、build / buy / adapt の判断は、機能だけでなく LLM-mediated workflow の総コストと正確性を含めて行う。

## 却下した代替案

### 代替案A: doc-policy の運用ルールだけを維持し、LLM-first を判断として記録しない

既存ルールだけでも日々の作業は実行できる。しかし、新しい session、外部 tool 評価、artifact 拡張、MCP 拡張でトレードオフを判断する際に、なぜその制約を維持すべきかという根拠が残らない。

そのため採用しない。

### 代替案B: LLM を secondary actor とし、人間中心の artifact workflow を主設計とする

人間中心の tool や document workflow が成熟している場合、それらを採用できる利点はある。しかし brewprint が重視するのは、LLM が session をまたいで project state を再取得し、曖昧な intent を適切な artifact と structured model に落とし、検証可能に扱えることである。

外部 tool は引き続き候補になり得るが、LLM-mediated 利用コストを考慮せず、人間向け primary workflow を brewprint の中心に置く案は採用しない。

### 代替案C: LLM-first を理由に writer tool と追加 intent / provenance mechanism を直ちに MVP へ追加する

LLM-first が write capability の将来価値を示すことはあり得るが、現時点では read-first foundation と既存 structured artifact が実運用で不足することをまだ観測していない。

未確認の不足を前提に MVP scope を広げると、V01-ADR-088 の scope reduction と矛盾し、過剰設計になる。まず dogfooding で concrete requirement を確認するため採用しない。

### 代替案D: LLM-native interface を持たない外部 tool を一律に排除する

外部 tool が native MCP interface を持たなくても、adapter cost が小さく、ownership と identity の境界を保てる場合には有効な backend 候補となり得る。

LLM-first は native-only 原則ではなく、LLM 利用時の総コストと正確性を比較する原則であるため、一律排除は採用しない。

## 影響

### Documentation policy への影響

`docs/doc-policy.md` は、MCP 経由の操作、spec-first、artifact ownership、canonical reference といった既存ルールの判断根拠として、本 ADR を参照し得る。

本 ADR は起票時点で policy 文面を直接変更しない。参照追記や LLM-first の入口説明が必要かは、後続の docs 同期判断で扱う。

### Project artifact model / traceability spec への影響

`spec:project-artifact-model` は、human intent、LLM mediation、requirements / investigation / ADR / spec / YAML、MCP query / validation の関係を current ownership boundary と矛盾しない形で明示する refinement 候補となる。

`spec:trace.artifact-refs` と `spec:trace.coverage-mapping` が所有する canonical identity および deferred realization coverage の境界は維持する。External tool interop や LLM-mediated evaluation のために spec への追記が必要かは、V01-INV-DOCS-004 / V01-INV-DOCS-005 の後続評価を含めて判断する。

### M19 / dogfooding への影響

本 ADR は、M19 の read / resolve / validation scope を拡大しない。M19 により成立する canonical reference foundation は、LLM-first workflow の initial substrate として dogfooding 対象になる。

Dogfooding により、現行 front matter、semantic ref、canonical reference、query、diagnostic の不足が具体的に観測された場合、必要な requirement / work item / ADR / spec refinement を別途起票する。

### External tool evaluation への影響

V01-INV-DOCS-004 が扱う OFT 評価、および V01-INV-DOCS-005 が扱う Doorstop 評価では、機能適合性・ownership 適合性に加えて、agent read I/O、diagnostic usability、identity mapping、testability、connector maintenance cost を判断材料に含める候補が生じる。

Write I/O、approval、追加 provenance の評価は、read-first dogfooding により必要性が確認された場合に追加する後段論点とする。

### 将来の MCP writer tool への影響

本 ADR は MCP writer tool の導入を決定しない。

Read-first dogfooding により controlled write の concrete requirement が確認された場合、LLM-first は、安全な mutation、validation-before-commit、人間の承認境界、必要な provenance / conflict handling を判断する根拠となる。

## Evidence

- commit: tbd
- impl commit: 該当なし
- investigation: V01-INV-DOCS-006
- decision context: V01-ADR-081, V01-ADR-083, V01-ADR-087, V01-ADR-088
