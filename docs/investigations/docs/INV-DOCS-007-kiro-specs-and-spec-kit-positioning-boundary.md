# INV-DOCS-007: Kiro Specs / GitHub Spec Kit と LLM-first architecture artifact positioning の境界

- **status**: investigating
- **date**: 2026-05-24
- **trigger**: INV-DOCS-006 において Kiro Specs と GitHub Spec Kit が LLM-first intent-to-artifact workflow の近接事例として確認され、ADR-089 により LLM-first を設計原則として採用した後の brewprint の positioning と競合境界を、OFT / Doorstop とは別軸で評価する必要が生じた
- **scope**: Kiro Specs / GitHub Spec Kit の公式 workflow・artifact・model-facing interface・対象フェーズを確認し、brewprint の LLM-first architecture artifact system としての重複領域、差別化候補、将来競合リスク、および positioning への含意を評価する
- **non_scope**: Kiro または Spec Kit の採用・連携決定、brewprint の positioning の最終確定、既存 spec / ADR / roadmap の変更、M19 scope の変更、connector / integration の実装、第三者製品の網羅的市場調査
- **source_refs**:
  - INV-DOCS-006
  - ADR-089
  - ADR-088
- **follow_up_candidates**:
  - comparative investigation across Kiro Specs / GitHub Spec Kit / brewprint positioning
  - project overview or positioning refinement for LLM-mediated architecture artifact system
  - spec:project-artifact-model refinement candidate for LLM-mediated design flow
  - external adjacent-tool evaluation rubric for workflow overlap and domain overlap
  - dogfooding evaluation of brewprint architecture-query value beyond feature-spec workflows

> 本 investigation は、Kiro Specs / GitHub Spec Kit を競合と確定するものではない。また、brewprint がこれらより優位であることを前提にしない。公式情報から確認できる workflow と artifact の重複、および確認できない差分を分けて記録し、brewprint の positioning 判断材料を作る。
>
> ADR-089 に従い、LLM-first は「native でないものを排除する」原則ではない。本 investigation では、OFT / Doorstop のような backend / trace capability 候補ではなく、intent-to-artifact workflow 自体が近い adjacent tooling として Kiro Specs / GitHub Spec Kit を評価する。

## 調査スコープ

- Kiro Specs の公式 documentation / announcement / product surface から、自然言語の意図を requirements / design / tasks へ変換し agent execution へ接続する workflow、artifact の種類、steering / MCP 等の model-facing context surface、および対象フェーズを確認する。
- GitHub Spec Kit の公式 repository / documentation から、constitution / specify / clarify / plan / tasks / implement を中心とする spec-driven agent workflow、生成 artifact、agent との関係、および対象フェーズを確認する。
- brewprint の ADR-089 と現行 artifact model に照らし、LLM-first workflow としての重複、architecture semantic DSL / MCP query / canonical ownership boundary としての差分候補を整理する。
- 「思想レベルで近い」と「同じ product category / domain を代替する」を分離して評価する。
- 将来 Kiro Specs / Spec Kit が architecture design 領域へ拡張した場合にも維持できる brewprint の価値仮説を、断定ではなく後続判断候補として整理する。

## 非スコープ

- Kiro IDE、Kiro CLI、Spec Kit の実導入・試用・ベンチマークを本起票では実施しない。
- Kiro Specs / Spec Kit が brewprint を代替できる、または代替できないと最終判断しない。
- Brewprint の YAML DSL、MCP tools、doc-policy、spec、M19 task、ADR-089 を本 artifact により変更しない。
- Feature development tooling、AI IDE、spec-driven development tooling の全市場調査は行わない。
- 公開資料で明記されていない内部 architecture、保存形式、validation implementation、roadmap を推測しない。
- Writer tool や implementation automation の有無を、read-first dogfooding を優先する brewprint MVP の欠陥として扱わない。

## 背景

INV-DOCS-004 / INV-DOCS-005 は、OFT / Doorstop を traceability または requirements management の外部 tool 候補として評価した。INV-DOCS-006 は、その比較に LLM I/O suitability と connector cost が不足していることを整理し、Kiro Specs と GitHub Spec Kit を「自然言語の意図から structured artifact を形成し agent execution へ渡す」近接事例として挙げた。

その後、ADR-089 は LLM-first を brewprint の設計原則として採用した。ADR-089 における LLM-first は、人間が意図と最終判断を担い、LLM が structured artifact の起草・探索・検証・許可された操作を担う協働モデルである。同時に、M19 の read / resolve / validation foundation をまず dogfooding し、writer や追加 metadata は具体的な不足が観測された場合に再判断する境界を維持した。

Kiro Specs と GitHub Spec Kit は、OFT / Doorstop と異なり、LLM / agent を workflow の中心に据え、ユーザーの idea や feature description から persistent な requirement / design / plan / task artifact を生成することを公式 surface として示している。この点で brewprint と設計思想が近い可能性がある。

一方、brewprint は、対象 system / design model の意味構造を YAML DSL として保持し、Mermaid view と MCP semantic query interface を提供する architecture artifact layer を目指している。したがって、intent-to-artifact workflow が重なっても、対象 domain、artifact ownership、semantic query、DSL、validation の粒度が異なれば棲み分けが成立し得る。

本 investigation は、Kiro Specs / Spec Kit を「LLM-first だから即競合」または「feature tooling だから無関係」と扱わず、workflow overlap と domain / semantic-model overlap を分離して確認するために起票する。

## 起票時点で確認した external sources

以下は、本文の観測事実を整理するために参照した公開一次情報である。Brewprint 内 canonical `source_refs` metadata には含めない。

- Kiro official documentation: Specs overview / workflow documentation — spec が requirements、design、implementation tasks を通じて idea を実装へ導く workflow の確認対象
- Kiro official documentation: Requirements / Design / Tasks documentation — EARS-style requirements、design artifact、task execution surface の確認対象
- Kiro official documentation: Steering and MCP documentation — persistent project context と external tools / services への接続 surface の確認対象
- Kiro official documentation: CLI / supported IDE integrations / model and credits documentation — Kiro の利用 surface と provider / platform boundary の確認対象
- GitHub `github/spec-kit` official repository README and documentation — Spec-Driven Development toolkit、CLI / slash-command workflow、constitution / specification / plan / tasks / implementation artifacts、および supported agent integrations / generic agent option の確認対象
- Model Context Protocol official documentation — client / server protocol boundary と、server を対応 client から利用する model の確認対象

## 調査項目ごとの確認結果

### Q1: Kiro Specs は何をするプロダクトか

#### 確認対象

- Kiro 公式 documentation に記載される Specs workflow と artifact structure
- Natural-language idea から agent execution へ至る流れ
- Steering / MCP 等の persistent context または model-facing interface
- Architecture design と feature implementation のどちらを主な対象として提示しているか

#### 観測事実

- Kiro の公式 Specs documentation は、spec を「idea から implementation まで」を進める structured workflow として説明し、requirements、design、implementation tasks の三段階を中心に扱う。
- Kiro Specs は、ユーザーが述べた feature または task の高位な要望から `requirements.md`、`design.md`、`tasks.md` を生成・反復し、tasks を agent が実行可能な単位へ展開する workflow を公式に示している。
- Requirements surface は EARS 形式を用いた acceptance criteria の表現を公式に示している。これは曖昧な意図を testable な requirement へ近づける設計であり、INV-DOCS-006 で扱った intent transfer problem に近い。
- Design artifact は、requirements を満たすための technical design を記録する。公式 documentation 上、architecture、components、interfaces、data models 等を design の内容として扱い得るが、起票時点で確認した資料からは、独立した architecture description DSL や project-wide semantic model を正本として提供するとは確認できない。
- Tasks artifact は、承認された requirements / design に基づく implementation step を記録し、Kiro agent が個別 task または一連の tasks を実行する入口となる。
- Kiro は steering files による persistent project context を持ち、agent が project knowledge・standards・workflow guidance を継続的に参照する surface を公式に提供している。
- Kiro は MCP servers を接続可能な surface として公式に説明している。ただし、この事実は Kiro Specs artifact 自体が brewprint のような architecture semantic query model を提供することを意味しない。MCP は external tools / data source を agent に接続する mechanism として確認できる範囲である。
- Kiro は IDE に加えて CLI、Web、ACP-compatible IDE integrations、automations を含む利用 surface を公式に示している。ただし、これらは Kiro platform 上の操作 surface の展開であり、Kiro のアカウント、agent / model、credits / pricing boundary から独立するものとは確認できない。特に Kiro CLI は、任意の MCP 対応 chat client から Kiro Specs を利用可能にする interface ではなく、Kiro の terminal-oriented surface として扱うべきである。
- 公式説明の中心は software feature の specification から implementation execution までの workflow である。Design phase に architecture 内容を含め得るが、architecture model 自体を independent implementation source または semantic query layer として所有することは、起票時点の確認範囲では確認できない。

#### 候補

| candidate | interpretation | brewprint との関係 |
| --- | --- | --- |
| Candidate A: Kiro Specs は feature-level spec-to-implementation workflow | requirements / design / tasks を agent 実行へ接続する product surface | Intent-to-artifact flow は近いが、architecture semantic model は別領域となり得る |
| Candidate B: Kiro Specs は architecture artifact system まで実質的に担う | design artifact が architecture を含むため brewprint と広く重複する | DSL / semantic query / ownership の確認なしに断定するのは早い |
| Candidate C: Kiro Specs は brewprint の upstream / downstream workflow と連携し得る | Kiro が feature delivery、brewprint が architecture semantic context を担う | 実接続 surface と ownership boundary の検討が必要 |

#### 判断に必要な観点

- Kiro の design artifact が architecture を記述できることと、architecture meaning を query / validate 可能な canonical model として所有することを区別できるか。
- Kiro MCP surface から brewprint MCP を利用する構成が成立する場合、brewprint は競合ではなく architecture-context provider として位置付けられるか。
- Kiro が将来 architecture-specific structured model / validation / query を拡張した場合に、brewprint の固有価値がどこに残るか。

#### 後続判断先

- Kiro Specs を workflow competitor、integration host、または両方の可能性を持つ adjacent product として扱う比較判断。
- 必要な場合のみ、Kiro agent が brewprint MCP を利用する限定 spike の候補化。

### Q2: GitHub Spec Kit は何をする toolkit か

#### 確認対象

- GitHub 公式 repository / documentation に記載される Spec-Driven Development workflow
- 生成・更新される artifact と agent interaction
- Kiro Specs との設計差分
- Architecture DSL / semantic query / model-facing tool surface の有無

#### 観測事実

- GitHub Spec Kit は、公式 repository において Spec-Driven Development のための toolkit として公開され、AI coding agent とともに specification を中心に development を進める workflow を示している。
- 公式 workflow は、project principles を定める constitution、機能要件を記述する specification、technical plan、tasks、implementation へ進む段階を持つ。公式 command surface には `/speckit.constitution`、`/speckit.specify`、`/speckit.clarify`、`/speckit.plan`、`/speckit.tasks`、`/speckit.implement` 等が示されている。
- Spec Kit は、ユーザーの intent を specification としてまず明確化し、plan / tasks を介して agent に実装させることを主眼とする。Clarification surface が明示されている点は、曖昧な意図を artifact 化する前後で確認する workflow として Brewprint の intent transfer 問題に近い。
- Kiro Specs が IDE/product の integrated agent workflow として requirements / design / tasks を提示するのに対し、Spec Kit は repository に導入して複数の coding agent で利用可能な toolkit / templates / command workflow として提示されている。
- 起票時点で確認した公式資料では、Spec Kit が独自の architecture meaning DSL、architecture graph の semantic query API、canonical ref resolution、または Brewprint のような MCP server を提供することは確認できない。
- Spec Kit の plan / design-related artifacts が architecture decisions や data model を含み得るとしても、公式 positioning の中心は software feature development の specification-driven execution であり、architecture semantic model の長期的 source of truth を独立に所有することは確認できない。
- GitHub Spec Kit の公式 repository は、多数の supported agent integrations に加えて `generic` option を示している。これは GitHub Copilot のみに固定されず、repository 内の files / commands を扱える coding agent 間で workflow を移植できることを示す。一方、Claude Desktop のような chat client を含む任意の MCP 対応 client から Spec Kit の capability を service として利用する設計であるとは確認できない。したがって、Spec Kit の coding-agent portability と、brewprint が MCP server / semantic interface boundary を通じて意図する client independence は別の概念として評価する。

#### 候補

| candidate | interpretation | brewprint との関係 |
| --- | --- | --- |
| Candidate A: Spec Kit は agent向け feature delivery workflow toolkit | Constitution / spec / plan / tasks / implement の repository workflow | Intent-to-artifact pattern は近く、semantic architecture layer は未確認 |
| Candidate B: Spec Kit を brewprint artifact workflow の代替として扱う | 仕様・計画・task の作成に重複がある | Requirements / ADR / DSL / MCP query ownership の差異確認が必要 |
| Candidate C: Spec Kit の downstream agent workflow が brewprint architecture context を利用する | spec-driven execution と architecture model を分離 | Integration の価値と重複コストの検証が必要 |

#### 判断に必要な観点

- Spec Kit の constitution / specification が brewprint の doc-policy / requirements / spec / ADR と意味的にどこまで重なるか。
- Spec Kit の generated artifacts を brewprint 内へ取り込むと ownership が二重化するか、それとも外部 implementation workflow として分離できるか。
- Brewprint の MCP semantic query が coding agent へ architecture context を供給する場合、Spec Kit と補完関係を作れるか。

#### 後続判断先

- Spec Kit を competing artifact workflow としてのみ見るか、architecture context consumer 候補としても扱うかの比較判断。
- Integration を評価する場合も、既存 source-of-truth boundary を壊さない条件の整理。

### Q3: Brewprint との重複領域はどこか

#### 確認対象

- 主役、intent-to-artifact flow、artifact 種別、対象フェーズ、ownership 規律、MCP / model-facing interface、DSL、query / validation の比較
- ADR-089 により確定した LLM-first と、Kiro / Spec Kit の公式 workflow の接点

#### 観測事実

| 軸 | brewprint | Kiro Specs | GitHub Spec Kit | 起票時点の評価 |
| --- | --- | --- | --- | --- |
| 主役 | 人間が意図・最終判断、LLM が structured artifact の起草・探索・検証・許可された操作を担う（ADR-089） | Agent が spec artifact の生成・反復・task execution を支援する公式 workflow | AI coding agent が specification-driven workflow を進める前提 | 三者とも LLM / agent-mediated workflow で近接する |
| 意図→artifact変換 | requirements / investigations / ADR / spec / YAML の責務分離を通じて精密化する方針 | Idea / feature request → requirements → design → tasks → execution | Constitution → specify / clarify → plan → tasks → implement | Intent-to-artifact flow は本質的に重なる |
| artifact の種類 | Capture requirement、investigation、ADR、current spec、architecture YAML、trace metadata | `requirements.md`、`design.md`、`tasks.md`、steering context | Constitution、specification、plan、tasks、agent-oriented project files | Brewprint は判断・調査・architecture model をより分離している候補 |
| 対象フェーズ | Architecture / design model の記述・照会が核。実装への関係は段階的 | Feature specification から implementation execution が公式中心 | Feature specification から implementation workflow が公式中心 | 現時点では domain / lifecycle center が異なる可能性が高い |
| artifact ownership 規律 | spec-first、capture requirement 非正本、ADR decision history、YAML primary implementation source、canonical refs | Requirements / design / tasks の段階 artifact と steering を確認。Brewprint と同一の source-of-truth separation は未確認 | Constitution / spec / plan / tasks の workflow を確認。Brewprint と同一の artifact ownership rule は未確認 | Brewprint の ownership 分離は差分候補だが優位性は未判断 |
| MCP / model-facing interface | MCP semantic query、Design Records MCP resolve / validation を設計対象とする | MCP servers 接続 surface を公式に確認。Specs artifact 自体の architecture query API は未確認 | Agent command workflow を確認。Dedicated MCP architecture-query surface は未確認 | Brewprint の architecture-query surface は差分候補 |
| LLM クライアント独立性 | MCP server / semantic interface を境界とし、chat client を含む MCP 対応 client から利用されることを意図する。特定 model provider を正本にしない | CLI / Web / ACP-compatible IDE 等は Kiro platform 上の利用 surface であり、Kiro の提供モデル・認証・credits から独立した任意 MCP client / provider 差替えは未確認 | 複数の coding agent integrations と `generic` option を確認。Repository files / commands を扱う coding agent 間の移植性であり、chat client を含む任意 MCP client からの service 利用とは異なる | Brewprint の MCP client independence は Kiro と Spec Kit の双方に対する差分候補。ただし Spec Kit は coding-agent portability という別種の選択性を持つ |
| LLM アクセスモード依存 | 利用者が選択した MCP 対応 client / host の利用形態を前提に接続する設計候補。Brewprint 自身は個人向け chat、開発向け app、API、enterprise のいずれかを必須化しない | Specs 実行に利用できる surface は複数あるが、いずれも Kiro の agent / model / credits 境界を利用する | Coding agent を実行 surface とする workflow であり、chat client を含む MCP host へ capability を公開する方式は未確認 | Brewprint は access mode の選択範囲を MCP client へ広げる差分候補を持つ。Spec Kit は coding-agent 間移植性として分離評価する |
| 追加プラットフォームコスト | Brewprint 自身は特定 LLM platform の追加契約を要求しない設計候補。ただし MCP 対応 client の利用可否・プラン料金・model 利用料金は利用者の選択に依存する | Kiro 利用に伴う plan / credits / overage 等の cost surface がある | Spec Kit toolkit 自体を特定の Copilot 契約必須とは扱えない。選択する agent / client / model の費用は別途生じ得る | 「Brewprint は常に追加費用ゼロ」「Spec Kit は GitHub 追加費用必須」とは断定しない |
| プラットフォーム依存 | MCP protocol と Brewprint artifact / query contract を中心に据える。個別 LLM client の製品 lifecycle を primary ownership としない設計候補 | Kiro platform の Specs / agent / model-credit surface に依存する範囲がある | Spec Kit toolkit 自体は複数の coding agent に展開可能。ただし chat client を含む MCP service boundary ではなく、各 coding-agent integration の command / repository workflow に依存する | Kiro と Spec Kit の双方に対して Brewprint の MCP client boundary は差分候補。ただし Spec Kit は coding-agent portability を持つ |
| DSL | Architecture meaning を YAML DSL で記述 | Requirements / design Markdown workflow。Architecture DSL は未確認 | Specification / plan templates workflow。Architecture DSL は未確認 | Brewprint の dedicated architecture DSL は明確な差分候補 |
| query / validation | DSL / records を structured query / resolve / validation 可能にする方針。M19 は read-first foundation | Spec workflow / agent execution / steering / MCP は確認。Architecture semantic validation は未確認 | Workflow commands と analysis / clarification は確認。Architecture graph query / validation は未確認 | Semantic architecture query / validation が差分候補 |

- Brewprint と Kiro / Spec Kit は、「曖昧な意図を LLM が persistent artifact へ精密化し、後続作業へ接続する」という workflow レベルでは重なる。
- 一方、起票時点の公式情報で確認できる Kiro / Spec Kit の中心は feature specification と implementation workflow であり、brewprint が目指す architecture meaning の YAML DSL、semantic query、canonical ref / record validation と同一の surface は確認できない。
- Brewprint の `docs/spec/` と Kiro / Spec Kit の specification artifacts は「spec」という語が共通していても、ownership と scope が同一とは限らない。Brewprint では現行 architecture / design contract の唯一の正であり、capture requirement・investigation・ADR・YAML と明示的に分離される。

#### 候補

| candidate | overlap assessment | implication |
| --- | --- | --- |
| Candidate A: 同一カテゴリの直接競合 | LLM-first artifact workflow と design artifact が重なるため同一市場として扱う | Architecture semantic layer の差分を過小評価する可能性 |
| Candidate B: Workflow-level competitor / domain-level complement | Intent-to-artifact の思想は競合するが、feature delivery と architecture semantic model で当面分離する | 競合警戒と統合可能性の両方を追跡できる |
| Candidate C: 無関係な参照事例 | 対象 domain が違うため競争関係を持たない | 将来 architecture 領域へ拡張された場合のリスクを見落とす |

#### 判断に必要な観点

- Brewprint が primary value として提供すべきものが architecture model / semantic query なのか、汎用の intent-to-artifact workflow まで含むのか。
- Kiro / Spec Kit を利用した実装 workflow に brewprint MCP が architecture context を供給する補完構成が成立するか。
- Brewprint の own requirements / investigations / ADR / spec workflow が、feature workflow tools と競合する範囲をどこまで意図的に持つべきか。

#### 後続判断先

- Brewprint positioning の refinement と、workflow scope を広げるか architecture layer に集中するかの判断候補。
- Kiro / Spec Kit との integration spike を行う場合の ownership / I/O boundary の事前整理。

### Q4: 棲み分けできるか、それとも本質的に競合するか

#### 確認対象

- Flow overlap と domain overlap の分離
- Brewprint の architecture meaning DSL / MCP semantic query / spec-first ownership / canonical refs の差分候補
- Kiro / Spec Kit が architecture design scope を広げた場合の positioning risk

#### 観測事実

- Kiro Specs / Spec Kit と brewprint は、LLM が user intent を persistent structured artifacts へ変換するという flow で近い。この意味では、ADR-089 の LLM-first を Brewprint のみの差別化語として用いることはできない。
- Kiro Specs は design artifact を生成し、Kiro は MCP server 接続により外部 context / tools を利用できるため、単純に「feature tool なので architecture とは無関係」と切り離すのも不適切である。
- Spec Kit は toolkit として agent workflow に組み込まれ、constitution / specification / plan / tasks を提供するため、Brewprint の docs workflow の一部と概念的な重複がある。
- ただし、起票時点の公式情報では、Kiro Specs / Spec Kit に Brewprint の architecture meaning YAML DSL と同等の dedicated model、Mermaid 等への architecture rendering、architecture graph の MCP semantic query、canonical semantic reference resolution / validation、investigation / ADR / current spec / implementation-source の責務分離が確認できない。
- したがって現時点では、Brewprint は LLM-first そのものでは差別化できないが、**architecture semantics を machine-readable source と query layer として所有すること**に差別化候補がある。
- 追加の差分候補として、Brewprint は MCP を protocol boundary として architecture semantics を client / model から切り離して提供することを意図できる。Kiro は利用 surface が IDE に限られない一方、Kiro platform の agent / model / credits boundary との結合が確認できる。Spec Kit は多数の agent integration と generic option を持つため、client independence の観点で Brewprint の単純な優位とは扱えず、protocol boundary・semantic query・domain ownership の差として再評価が必要である。

#### 候補

| candidate | positioning implication | risk |
| --- | --- | --- |
| Candidate A: Brewprint を汎用 spec-driven development workflow として広げる | Kiro / Spec Kit と直接競合する領域へ進む | 成熟した product / toolkit との重複が大きくなる |
| Candidate B: Brewprint を LLM-mediated architecture semantic layer として明確化する | Kiro / Spec Kit の implementation workflow へ context を供給する補完可能性を残す | Feature-level workflow の一部を外部に依存する可能性 |
| Candidate C: Architecture design から implementation workflow まで一貫所有する | 独自の end-to-end surface を構成できる | Scope が拡大し、重複・保守・競合リスクが最大化する |

#### 判断に必要な観点

- Brewprint の YAML DSL / MCP query / validation が、Kiro / Spec Kit では代替しにくい実務価値を dogfooding で示せるか。
- Kiro / Spec Kit の artifact を brewprint source of truth に取り込まず、implementation-side consumer として接続できるか。
- 将来 Kiro / Spec Kit が architecture model を持つ場合に、Brewprint の open protocol / DSL / tool-agnostic semantic layer としての価値が残るか。

#### 後続判断先

- Project overview / positioning statement を「LLM-first」だけでなく architecture semantic layer に寄せて refine するか。
- Integration / competition の双方を検証する small spike の要否。

### Q5: Brewprint のポジショニングへの含意は何か

#### 確認対象

- INV-DOCS-006 の positioning 候補が Kiro / Spec Kit との比較後も有効か
- ADR-089 による LLM-first の固有の意味
- 競合警戒、参照事例、補完統合の扱い

#### 観測事実

- `LLM-mediated architecture artifact system` は、Kiro / Spec Kit と共通する `LLM-mediated` だけでなく、`architecture artifact` を含むため、起票時点の差分候補を表す positioning として依然有効である可能性がある。
- `LLM-native architecture specification and query layer` は、Brewprint の YAML DSL と MCP query / resolve / validation に焦点を当てる点で、Kiro / Spec Kit の feature workflow との重複を避けやすい候補である。
- ADR-089 の LLM-first は、brewprint が唯一 LLM を使うという主張ではない。LLM が session 間で再取得できる ownership / canonical identity / validation boundary を project-level constraint として置く判断であり、Kiro / Spec Kit が agent-first workflow を持つことと両立する。
- Brewprint が MCP server / semantic interface を boundary とすることは、architecture semantics の owner を特定の agent product や model provider に置かず、chat client を含む MCP 対応 client へ公開する positioning 候補を支える。これは Kiro の platform-integrated workflow との比較では差分候補となり、Spec Kit との比較でも、repository を操作する coding agent 間の portability とは異なる MCP client independence として差分候補になり得る。ただし対 Spec Kit では architecture domain と query / validation surface を伴って評価すべきである。
- Kiro Specs は、intent-to-artifact flow と MCP / persistent context surface の両面で、brewprint の将来 positioning に影響し得る最も近い比較対象の一つと考えられる。Spec Kit は product integration より toolkit / process pattern として、brewprint の workflow 境界・artifact discipline、および coding-agent portability と MCP client independence の違いを比較する参照対象として強い。
- 起票時点で、両者を無関係と扱う根拠も、Brewprint の直接代替と扱う根拠も不足している。Workflow-level competition と architecture-layer complement の双方を残す評価が妥当と見られる。

#### 候補

| candidate | positioning | evaluation |
| --- | --- | --- |
| Candidate A: LLM-first design tool | LLM-first 自体を主差別化にする | Kiro / Spec Kit も近い flow を持つため弱い |
| Candidate B: LLM-mediated architecture artifact system | Architecture artifacts と LLM mediation の双方を示す | 有効候補だが workflow scope が広く見える可能性 |
| Candidate C: LLM-native architecture specification and query layer | Architecture DSL / structured query / validation を中心に示す | 差分が明確になりやすいが、人間との intent transfer story を別途補う必要がある |
| Candidate D: Architecture context backend for agentic development workflows | Kiro / Spec Kit 等へ context を供給する補完的位置付け | Integration feasibility と product aspiration の判断が必要 |

#### 判断に必要な観点

- Brewprint の primary audience / usage は、独立した design authoring workflow を求める利用者か、Kiro / Spec Kit 等の agent workflow に architecture context を供給したい利用者か、両方か。
- Positioning は implementation workflow の代替を意図するのか、architecture semantics の source / query / validation を中心にするのか。
- Client / provider independence を外部 positioning へ含める場合、Kiro に対する差分と Spec Kit に対する非差分を誤解なく説明できるか。
- Dogfooding で、YAML DSL と MCP query が feature-level spec artifact だけでは得られない価値を実証できるか。

#### 後続判断先

- Project overview / README / spec overview の positioning refinement 候補。
- MCP boundary と client / provider independence を positioning に含める場合の表現精査候補。
- Kiro / Spec Kit を comparison set とした dogfooding evaluation plan の候補。

## 横断的な観測事実

### 1. Kiro Specs / Spec Kit は OFT / Doorstop と評価カテゴリが異なる

OFT / Doorstop は、主に既存の traceability / requirements management capability を Brewprint の backend または artifact ownership 候補として評価する対象であった。Kiro Specs / Spec Kit は、ユーザー意図から structured artifacts を生成し agent execution へ接続する workflow を公式 surface とする点で、ADR-089 の LLM-first workflow に近い adjacent tooling である。

このため、Kiro / Spec Kit は connector cost だけでなく、Brewprint の product positioning と workflow scope に直接影響する比較対象として扱う必要がある。

### 2. LLM-first 自体は Brewprint の固有差分にならない

Kiro Specs と Spec Kit は、公式資料上、agent が requirements / design / specification / plan / tasks 等の artifact を生成・利用して implementation を進める workflow を持つ。したがって、Brewprint が LLM-first を採用したことだけを market / product differentiation として用いるのは不十分である。

ADR-089 の価値は差別化 slogan ではなく、Brewprint 内の ownership、canonical identity、query / validation、build / buy / adapt 判断を統一する設計原則にある。

### 3. 現時点の差分候補は architecture semantic model と query / validation boundary にある

Brewprint は、architecture meaning を YAML DSL に保持し、人間向け render と LLM 向け MCP query / resolve / validation を提供することを project の核としている。Kiro Specs / Spec Kit は、起票時点で確認した公式資料上、feature-to-implementation の spec-driven workflow を中心に説明しており、同等の dedicated architecture meaning DSL / semantic query layer / canonical trace boundary は確認できない。

この差分は採用判断ではなく positioning 仮説であり、Brewprint の dogfooding と追加調査で実効価値を確認する必要がある。

### 4. 補完関係は成立し得るが、ownership を混ぜると危険である

Kiro が MCP server を利用可能であること、また Spec Kit が agent-oriented repository workflow であることから、将来的に Brewprint が architecture context provider として利用される構成は候補になり得る。

ただし、Kiro / Spec Kit の generated spec / design / task artifact と Brewprint の requirements / investigations / ADR / current spec / YAML の source-of-truth boundary を混在させると、同名 artifact の意味が二重化し得る。Integration を評価する場合は、Brewprint が architecture semantic source を所有し、外部 workflow は consumer または downstream execution layer となるのかを先に整理する必要がある。

### 5. Brewprint の MCP 境界選択は LLM client / provider から architecture semantics を分離する差分候補である

Brewprint が MCP を interface boundary として選ぶ意味は、単に「LLM から呼べる」ことではなく、architecture artifact と semantic query / validation contract を、特定の IDE、agent product、model provider、料金体系から切り離して提供できる設計にある。LLM provider ごとの system prompt、能力、料金、提供機能は用途や時期により変わり得るため、どの provider / client が最適かを Brewprint が固定しないという判断と整合する。

Kiro は、IDE 以外にも CLI、Web、ACP-compatible IDE integrations、automations を公式に持つが、これらは Kiro platform 上での操作 surface の展開である。Kiro CLI が存在しても、Kiro の agent / model / credits boundary を離れて Claude Desktop 等の任意 MCP 対応 client から Specs capability を利用できることを意味しない。したがって、Kiro との比較では、brewprint が architecture semantics を MCP server として client / provider から切り離して提供する意図は明確な差分候補となり得る。

GitHub Spec Kit は、Copilot 固定とは言えない。公式 repository は多数の supported agents と `generic` option を提供しており、coding agent 間で workflow を移植できる余地を持つ。しかし、この agent portability は、repository files / commands を扱う coding agent を前提とするものであり、chat client を含む任意の MCP 対応 client が semantic service を照会できるという Brewprint の client independence とは別の概念である。

したがって、Brewprint の MCP client independence は Kiro と Spec Kit の双方との比較で差分候補となり得る。ただし Spec Kit については、coding-agent portability という別の利点を持つため、単純に flexibility が低いとは扱わず、MCP protocol boundary、architecture semantic DSL、query / validation、canonical ownership と合わせて評価すべきである。

### 6. Brewprint の MCP 境界は LLM アクセスモードとコスト構造の選択を利用者側に残す設計候補である

LLM 利用の現実的な選択肢は、model provider の違いだけではなく、どの access mode を通して利用するかでも異なる。個人向け chat application、coding / agent application、API direct usage、enterprise-managed offering では、料金体系、利用可能な tool / MCP surface、system prompt や周辺 orchestration の性質、管理・compliance 条件が異なり得る。各製品・plan の具体条件は変化し得るため、以下は固定的な価格表ではなく、Brewprint が介入しない選択領域を整理する評価枠である。

| access mode | 典型的な利用形態 | Brewprint が固定しない論点 |
| --- | --- | --- |
| 個人向け chat application | Subscription-based chat client | MCP 対応有無、利用制限、会話向け orchestration、利用料金 |
| Coding / agent application | IDE / CLI / coding agent | Repository 操作能力、coding-oriented orchestration、model / credit / subscription 条件 |
| API direct usage | Programmatic model API | 従量課金、system prompt / tool orchestration の自前設計、運用負担 |
| Enterprise-managed offering | Organization-managed client / API | Data governance、compliance、管理 controls、契約条件 |

Brewprint が MCP を interface boundary として選ぶ場合の設計上の意味は、任意の利用者が既に利用可能な **MCP 対応 client / host** があるなら、その access mode 上で architecture semantics を照会できる余地を残すことである。Brewprint 自身が Kiro plan、特定 coding agent、特定 model API、または特定 enterprise contract を必須の primary interface として要求しないことが差分候補となる。

ただし、この方針は「すべての既存 chat subscription で追加費用なく直ちに動く」ことを意味しない。MCP server を接続可能か、どの tool surface が利用できるか、model 利用に別料金・制限があるかは、選択した client / provider / plan の条件に依存する。したがって、正確な表現は **Brewprint は追加の特定 LLM platform 契約を設計上必須化しない** であり、利用者の総コストが常にゼロになるという主張ではない。

Kiro は複数の利用 surface を提供するが、起票時点で確認した公式情報上、Specs / agent 実行は Kiro の model / credits / pricing boundary の内側にある。このため、Kiro との比較では、Brewprint が architecture semantics を特定 product の利用料金・model availability から切り離して提供し得ることが差分候補となる。

GitHub Spec Kit は、公式 repository 上で複数の supported coding agent integrations と `generic` option を示しているため、特定 GitHub Copilot 契約のみを必須と扱ってはならない。一方、その portability は repository files / commands を扱う coding agent 間のものであり、chat client を含む任意の MCP 対応 client から capability を service として利用できることを意味しない。Spec Kit との比較では、選択した coding agent ごとの cost boundary に加え、Brewprint の MCP-based semantic query / ownership boundary と access-mode independence の違いを評価する必要がある。

## 候補比較

### Positioning comparison candidate

| axis | brewprint | Kiro Specs | GitHub Spec Kit | implication |
| --- | --- | --- | --- | --- |
| LLM / agent mediated workflow | ADR-089 により採用 | 公式 Specs workflow の中心 | 公式 Spec-Driven Development workflow の中心 | LLM-first のみでは差別化不可 |
| Persistent intent artifacts | requirements / investigations / ADR / spec | requirements / design / tasks / steering | constitution / spec / plan / tasks | Artifact pipeline に概念重複あり |
| Decision history ownership | ADR を独立所有 | 同等の ADR ownership は未確認 | Constitution / planning はあるが同等の ADR ownership は未確認 | Brewprint 差分候補 |
| Current contract ownership | `docs/spec/` を唯一の正とする | Requirements / design / tasks workflow は確認、同一 boundary は未確認 | Specification / plan workflow は確認、同一 boundary は未確認 | Brewprint 差分候補 |
| Architecture meaning source | YAML DSL | Dedicated architecture DSL は未確認 | Dedicated architecture DSL は未確認 | Brewprint の中心差分候補 |
| Model-facing structured query | MCP semantic query / resolve / validation | MCP 接続 surface は確認、Specs semantic query は未確認 | Agent command workflow は確認、architecture query surface は未確認 | Brewprint の中心差分候補 |
| LLM client / provider independence | MCP protocol boundary を中心とし、chat client を含む MCP 対応 client へ architecture semantic service を公開する設計候補 | IDE / CLI / Web / ACP integration は Kiro platform 上の surface。Kiro agent / model / credits boundary から独立した MCP client 利用は未確認 | 複数の coding agent integrations と `generic` option を確認。Coding agent 間の移植性であり、任意 MCP client 向け semantic service ではない | Brewprint の MCP client independence は双方に対する差分候補。Spec Kit の coding-agent portability は別軸で評価 |
| LLM access-mode dependence | MCP 対応 client / host の範囲で、利用者が選ぶ chat / coding app / API / enterprise access mode に architecture service を接続する設計候補 | Kiro surface と提供される agent / model / credit regime に乗る必要がある範囲を確認 | Repository を操作する coding agent を前提とする workflow。Chat client 経由の service 利用は未確認 | Brewprint は特定 access mode を必須化しない差分候補。ただし実際の client 対応可否は検証対象 |
| Additional platform cost imposed by tool | Brewprint 自身は特定 LLM platform への追加契約を必須化しない。既存 client の対応状況・利用料は別問題 | Plan / credits / overage を含む Kiro cost surface が存在する | Toolkit が単一 paid agent subscription を必須化するとは確認できない。利用 agent 側の費用は残り得る | Cost 差分は Kiro では評価可能。Spec Kit では選択 agent 単位で比較が必要 |
| Platform dependency | Brewprint artifact / MCP contract を owner とし、chat client を含む MCP 対応 client からの利用を意図する方向 | Kiro product platform との結合が確認できる | Toolkit は複数の coding agent に展開可能だが、各 coding-agent integration の repository / command workflow に依存する | Brewprint の MCP client boundary は双方との差分候補。Spec Kit の coding-agent portability は別軸で評価 |
| Implementation execution | Brewprint MVP の中心ではない | Tasks から agent execution へ接続 | Implement workflow へ接続 | Kiro / Spec Kit の中心価値 |
| Integration potential | architecture context provider となり得る | MCP 経由 consumer の可能性 | Agent workflow の context consumer の可能性 | 要 spike / ownership整理 |

### Strategy candidates

| candidate | meaning | benefit | risk / question |
| --- | --- | --- | --- |
| Candidate A: Kiro / Spec Kit と同じ end-to-end spec workflow を目指す | Intent capture から implementation tasks まで広く所有する | 一貫体験を構築し得る | 直接競合と scope 拡大が大きい |
| Candidate B: Architecture semantic layer に集中し、agent workflows へ接続可能にする | DSL / query / validation を中心価値とする | 差分を保ちつつ補完関係を作り得る | Integration usefulness を dogfooding で検証する必要がある |
| Candidate C: Comparisons を参照に留め、positioning を変えない | 現行設計を進める | 即時 scope 変更を避ける | Product boundary の曖昧さを残す |

## 起票時点の推奨案

本 investigation の起票時点では、**Candidate B: Brewprint を architecture semantic layer として明確化しつつ、Kiro / Spec Kit のような agentic implementation workflow へ接続可能な補完位置を検証する**ことが妥当と考えられる。

理由は以下である。

1. Kiro Specs / Spec Kit は、human intent を structured artifacts へ変換し agent execution に接続する workflow を公式に示しており、LLM-first という flow 自体では brewprint の固有差分にならない。
2. 一方、起票時点で確認した公式資料上、両者には Brewprint の architecture meaning YAML DSL、MCP semantic query / resolve / validation、canonical ref と artifact ownership boundary に相当する一体の surface は確認できない。
3. Brewprint が汎用 feature-to-implementation workflow へ広く進むと、Kiro / Spec Kit との直接重複を増やし、現行 DSL / query layer の核がぼやける可能性がある。
4. Kiro が MCP 接続 surface を持つことは、Brewprint が architecture context provider として補完可能になる仮説を支持する。また、Kiro との比較では、特定 product / model-credit boundary に architecture semantics を閉じず MCP server として提供することが差分候補となる。ただし、実用性と ownership boundary は spike または dogfooding で確認すべきである。
5. Spec Kit は多数の coding agent integration と `generic` option を公式に持つが、これは coding agent 間の portability であり、chat client を含む任意 MCP client に semantic service を公開する Brewprint の意図とは別の拡張性である。したがって MCP client independence は対 Spec Kit でも差分候補となり得るが、Spec Kit の portability を過小評価せず、architecture semantic DSL / MCP query / canonical ownership と組み合わせて評価すべきである。
6. ADR-089 に従い、現時点で新たな integration / writer / workflow expansion を決定せず、まず Brewprint の read-first MCP と architecture artifact model の価値を dogfooding で観測する方が適切である。

## 後続判断に渡す候補

- Brewprint の positioning を `LLM-mediated architecture artifact system` または `LLM-native architecture specification and query layer` のどちらへ寄せるかを、Kiro / Spec Kit comparison と dogfooding 結果を基に判断するか。
- `LLM-first` は internal design principle として維持し、external positioning では architecture semantic DSL / query / validation を中心に表現するか。
- Brewprint の requirements / investigations / ADR / spec workflow を、汎用 feature implementation workflow として拡張するのではなく、architecture semantic source の ownership と refinement に限定するか。
- Kiro の MCP server 接続 surface または agent-oriented workflow から Brewprint MCP を参照する small integration spike を、ownership を混在させない条件付きで評価するか。
- Spec Kit の constitution / spec / plan / task flow から、Brewprint の dogfooding evaluation または authoring guidance に取り入れるべき workflow practice があるか。
- Project overview / spec overview / README に positioning を反映する必要が成立した場合、ADR または spec refinement を起票するか。

## 後続 artifact 候補

| category | candidate artifact / action | purpose |
| --- | --- | --- |
| Investigation update | 本 investigation の更新 | Kiro / Spec Kit official sources の追加確認、比較結論の精緻化 |
| Comparative investigation | Brewprint / Kiro Specs / GitHub Spec Kit positioning comparison | Positioning 判断を独立して深掘りする場合 |
| Project overview / spec refinement | LLM-mediated architecture semantic layer positioning | Current ownership と external positioning を整合させる場合 |
| Dogfooding work item | Architecture YAML / MCP query value evaluation against feature-spec workflow | Brewprint 固有価値の実運用観測を行う場合 |
| Integration spike | Kiro or agent workflow consumes Brewprint MCP without ownership transfer | 補完可能性を限定検証する場合のみ |
| Evaluation rubric update | Adjacent agentic workflow comparison axes | 外部 tool 評価を backend と workflow competitor に分類する場合 |

## 未確定点

- Kiro Specs の design artifacts が、大規模・長期的な architecture source of truth としてどの程度運用されることを公式に想定しているか。
- Kiro の MCP server 接続 surface が、Brewprint MCP のような architecture context provider を実用上取り込めるか、その場合の query / context / validation 負担がどの程度か。
- Kiro が将来 API-first、既存 client からの接続、または model / client-agnostic な interface を公式に提供した場合、MCP 境界および access-mode 非介入による Brewprint の差分候補がどの程度縮小するか。
- MCP 対応が各 chat / coding / enterprise client に広く標準提供される場合、Brewprint の client / cost non-intervention が独立した差分ではなく前提条件へ変化するか。
- GitHub Spec Kit の supported agent integrations / `generic` option が実際にどの LLM client / provider / access mode でどの程度同等に動作するか、および Brewprint の MCP protocol boundary と実務上どう異なるか。
- GitHub Spec Kit の constitution / specification / plan と、Brewprint の doc-policy / requirements / ADR / spec の ownership を混ぜずに併用できるか。
- Kiro / Spec Kit が将来 architecture-specific DSL、semantic model、query / validation surface を拡張する可能性と、その場合の Brewprint positioning への影響。
- Brewprint の architecture meaning YAML DSL と MCP query が、feature-oriented spec artifacts だけでは解けないどの実務課題に価値を持つか。
- `LLM-mediated architecture artifact system` と `LLM-native architecture specification and query layer` のどちらが、実装範囲と product aspiration を誤解なく示すか。

## 起票時点の制約記録

- 本 artifact の `source_refs` は、Brewprint 内の record ID-as-ref に限定した。Kiro / Spec Kit の外部資料は本文の調査対象としてのみ記載し、metadata には含めない。
- 本 investigation は、ADR-089 が採用した LLM-first 原則を前提として comparison scope を定めるが、Brewprint の外部 positioning や product strategy を決定しない。
- Kiro Specs / Spec Kit の観測事実は、起票時点で確認した公式 documentation / repository / announcement の記載範囲に限定する。確認できない architecture DSL、query API、validation semantics、roadmap は存在しないと断定しない。
- Kiro / Spec Kit が LLM-first workflow を持つことは Brewprint の価値を否定しないが、LLM-first のみを Brewprint 固有差分とする根拠もない。
- 本起票では、新規 investigation artifact の作成と investigation 一覧の同期以外に、ADR、spec、doc-policy、task、requirement、work item、implementation を変更しない。
