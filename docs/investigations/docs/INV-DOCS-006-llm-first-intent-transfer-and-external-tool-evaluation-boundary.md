# INV-DOCS-006: LLM-first intent transfer and external tool evaluation boundary

- **status**: investigating
- **date**: 2026-05-24
- **trigger**: OFT / Doorstop の機能・artifact ownership 適合性を評価する過程で、brewprint が「人間が直接 artifact を編集・閲覧する」ことよりも「人間が意図を伝え、LLM が structured artifact を作成・探索・検証する」ことを主用途とするという設計意図が明示され、外部 tool 評価に LLM I/O suitability と connector cost の軸が不足していることが判明した
- **scope**: LLM-first 設計原則、intent transfer problem、brewprint artifact model / MCP boundary の役割、および OFT / Doorstop 等の外部 tool を LLM 主体の運用へ接続する際の評価軸と後続設計含意
- **non_scope**: LLM-first 原則の正式採用決定、OFT / Doorstop の採否決定、MCP writer tool の実装、M19 implementation の変更、外部 tool connector の試作、brewprint のカテゴリやロードマップの最終確定
- **source_refs**:
  - ADR-076
  - ADR-081
  - ADR-083
  - ADR-087
  - ADR-088
  - INV-DOCS-004
  - INV-DOCS-005
  - spec:project-artifact-model
  - spec:trace.artifact-refs
  - spec:trace.coverage-mapping
- **follow_up_candidates**:
  - ADR for LLM-first design principle and human-intent / agent-artifact boundary
  - spec:project-artifact-model
  - spec:trace
  - Design Records MCP read / write / validation capability roadmap
  - external tool evaluation rubric including agent I/O and connector cost
  - comparative investigation across Brewprint-native / OFT adapter / Doorstop adapter approaches
  - isolated agent-facing connector spike only after rubric-based screening
- **follow_up_results**:
  - ADR-089
  - INV-DOCS-007

> 本 investigation は、ユーザーが本会話で明示した「LLM が artifact の主たる読み書き主体であり、人間は意図と最終判断を担う」という設計意図を、新たに確認すべき設計仮説として記録する。これは現時点で accepted ADR または confirmed spec に昇格した原則ではない。
> 本 investigation は、OFT / Doorstop が LLM-first ではないと断定するためのものではない。公式に確認できる interface と workflow を基に、LLM が利用する場合の追加 adapter / orchestration / validation 負担を比較可能にするための調査である。

## 調査スコープ

- 現行 Brewprint docs に既に存在する LLM / MCP を中心とした責務境界と、今回明示された LLM-first 設計意図の一致点・不足点を確認する。
- OFT / Doorstop の公開 interface と workflow が、LLM による read / query / write / validate / explain loop にどの程度直接適合するかを確認する。
- 外部 tool が coverage / requirements management 機能を持つことと、LLM-first workflow へ組み込む総コストが低いことを区別する評価軸を定義する。
- 「人間の曖昧な意図から、LLM がレビュー可能で機械検証可能な artifact を形成する」問題を intent transfer problem として整理し、近接する先行アプローチを調べる。
- Brewprint の requirements / investigations / ADR / spec / YAML / MCP のうち、intent capture、ambiguity reduction、decision ownership、canonical contract、machine query / validation をどこが担うべきかを評価する。
- LLM-first architecture tooling としての Brewprint の位置付け候補、および外部 tool interop の採用基準を整理する。

## 非スコープ

- 本 investigation だけで LLM-first を正式な project principle として確定しない。
- OFT / Doorstop を不採用または採用と決定しない。
- OFT / Doorstop の MCP connector、CLI wrapper、import / export adapter を実装しない。
- M19 の resolver / validation contract、ADR-088 の MVP scope、`yaml:` / `internal-design:` / `coverage:` endpoint の active scope を変更しない。
- Prompt engineering 一般、LLM model quality 一般、agent framework 全般を網羅的に評価しない。
- 「brewprint が既存研究・製品に対して新規性を持つ」とする法的・学術的主張を確定しない。

## 背景

`docs/doc-policy.md` は、brewprint を「人間と LLM の共通設計言語」と定義し、人間向けには Mermaid render、LLM 向けには MCP 経由の signature / dep tree / inspect を提供すると記述している。`spec:project-artifact-model` も、Design Records MCP を design record / investigation の探索と canonical ref の resolve / validation を担う tool boundary と位置付ける。

一方、現行 docs において LLM / MCP は主に query interface または tool boundary として記述されている。人間が意図を対話で渡し、LLM が requirement / investigation / ADR / spec / YAML を正しい ownership boundary に従って起草・更新・検証することを project-wide の一次設計原則として定める記述、または外部 tool 選定で agent I/O suitability を必須評価軸とする記述は確認できない。

INV-DOCS-004 は OFT を deferred realization coverage backend 候補として、identity / revision / M19 contract との境界を調査対象とした。INV-DOCS-005 は Doorstop を requirements / spec / trace artifact ownership に重なり得る候補として分離評価した。これらはいずれも機能と ownership を主軸にしており、LLM が primary actor である運用において、外部 tool を読み書き・探索・検証可能にする connector / projection / error recovery のコストを主評価軸として明示していない。

本会話でユーザーは、brewprint の意図として、人間が各 artifact を直接扱うよりも、LLM と対話して意図を伝え、LLM が MCP 経由で構造化 artifact を作成・操作する運用を中心に考えている旨を明示した。この追加情報が accepted decision として成立するなら、外部 tool の評価は「同じ機能が存在するか」だけでなく、「LLM に対して semantic query / constrained write / actionable diagnostic / provenance を低摩擦で公開できるか」を含めなければ不十分となる。

## 起票時点で確認した external sources

以下は、本文の観測事実と後続調査候補を整理するために参照した公開情報である。Brewprint 内 canonical `source_refs` metadata には含めない。

- OpenFastTrace official repository and User Guide — specification items、`Needs` / `Covers`、revision、Markdown / source-code item format、CLI / report workflow の確認対象
- Doorstop official documentation and repository — documents、items、links、review fingerprints、validation / publication、command line と Python scripting interface の確認対象
- Model Context Protocol official documentation — server が tools / resources / prompts を model-facing primitive として公開する protocol model の確認対象
- GitHub Spec Kit official repository and documentation — AI coding agent と spec / plan / tasks artifact を用いる spec-driven development workflow の近接事例
- Kiro official documentation on Specs — natural-language idea から requirements / design / tasks を生成し agent execution に接続する product workflow の近接事例
- Mavin et al., *Easy Approach to Requirements Syntax (EARS)* — 曖昧な自然言語要求を定型構文へ寄せる requirements engineering の先行アプローチ

## 調査項目ごとの確認結果

### Q1: OFT / Doorstop との根本的なターゲット差異は何か

#### 確認対象

- OFT / Doorstop の公式 interface、primary workflow、artifact format、validation / reporting surface
- Brewprint の MCP / semantic query / artifact ownership の現行記述
- 「human-first」「LLM-first」を、推測ではなく観測可能な interface 特性へ分解できるか

#### 観測事実

- 現行 Brewprint docs は、LLM 向け interface として MCP query surface を明示している。MCP は単なる export report ではなく、LLM が design structure を照会するための tool boundary と位置付けられている。
- OFT の公式 User Guide は、requirements engineer、technical writer、developer、tester 等が specification items を記述し、OpenFastTrace が links / coverage を検査して report を生成する workflow を示す。起票時点で確認した公式資料では、MCP または model-facing structured query / write interface は確認できなかった。
- Doorstop の公式 documentation は、version-controlled requirements の document / item tree、link validation、review、publish、および CLI / Python scripting を示す。起票時点で確認した公式資料では、MCP interface は確認できなかったが、Python scripting interface があるため adapter surface は OFT と同一ではない。
- 「pre-LLM 時代の tool」「人間が主役の tool」という表現は、成立史や公式 design intention を別途立証しない限り断定に向かない。現時点で比較可能なのは、公式公開されている primary interface が human-authored document + CLI/report か、agent が直接利用できる structured query / constrained write / diagnostic API を備えるかである。

#### 候補

| candidate | position | implication |
| --- | --- | --- |
| Candidate A: 既存 tool と Brewprint は機能が重なるため同一カテゴリとして比較する | traceability / requirements 機能中心の比較 | LLM connector cost を過小評価する危険がある |
| Candidate B: Brewprint を LLM-mediated design artifact system として分け、外部 tool は backend / interoperability candidate として評価する | actor / interface / ownership を含む比較 | 機能重複と product category を分離できる |
| Candidate C: LLM interface は後付け adapter で十分として、既存 tool を primary model にする | external tool ownership を優先 | adapter が artifact semantics / error recovery / write safety を再実装する可能性がある |

#### 判断に必要な観点

- LLM が item を列挙・絞込・解決・関連探索・安全に変更・validation failure を修正できるために、何が native surface として必要か。
- CLI 出力や document file を MCP wrapper で露出すれば十分なのか、semantic write contract / provenance / lifecycle rule まで Brewprint が所有する必要があるのか。
- 外部 tool の機能成熟度と、LLM 主体の end-to-end workflow の総実装コストを別軸で比較できるか。

#### 後続判断先

- LLM-first を project design principle として ADR 化するか。
- External tool comparative investigation に agent I/O / connector cost rubric を必須項目として追加するか。

### Q2: 「意図を LLM に正確に伝える」問題は既存のアプローチでどう扱われているか

#### 確認対象

- Requirements engineering における曖昧さ抑制と structured requirement authoring
- AI agent / spec-driven development における idea-to-spec-to-task workflow
- Prompt だけではなく、persistent artifact、constraint、validation、trace を使うアプローチ

#### 観測事実

- Requirements engineering には、曖昧な自然言語要求を定型化してレビューしやすくする既存アプローチがある。EARS は要求文をイベント・状態・応答等の構文パターンへ寄せることで、自由文より精度の高い requirement 記述を目指す。ただし、これは LLM を primary actor とする仕組みではなく、人間の requirement authoring を改善する枠組みである。
- 近年の agent-oriented development tooling では、自然言語の feature idea を persistent な specification / plan / task artifact に分解し、agent implementation と接続する workflow が現れている。GitHub Spec Kit は spec-driven development を掲げ、AI coding agent が利用する specification、plan、tasks の流れを公開している。Kiro Specs も requirements、design、tasks を生成し agent execution に接続する workflow を公開している。
- 上記の近接事例は、「vague intent をそのまま一度の prompt で code にする」のではなく、中間 artifact を保持し、レビュー・再利用・実行の境界にする点で、brewprint の問題意識に近い。
- ただし、起票時点で確認した公開情報だけでは、architecture description DSL、canonical semantic ref、investigation / ADR / spec の ownership、MCP query interface を一体として扱う既存 OSS が Brewprint と同じ問題設定を完全に満たすとは確認できない。

#### 候補

| approach | intent transfer で担うこと | Brewprint への示唆 |
| --- | --- | --- |
| Structured requirement syntax (例: EARS) | 曖昧な要望をレビュー可能な要求文へ制約する | requirement / spec authoring の精密化 rule として参照可能 |
| Spec-driven agent workflow (例: Spec Kit / Kiro Specs) | idea を spec / plan / tasks へ展開し agent implementation に渡す | LLM-first 原則と artifact pipeline を明示する近接事例 |
| Tool protocol / MCP | model が external context / actions を構造化 interface で利用する | read / write / validate surface を protocol boundary として設計可能 |
| Traceability tool adapter (OFT / Doorstop wrapper) | 既存 validation / lifecycle capability を agent から利用可能にする | connector cost と semantic mismatch を測る必要がある |

#### 判断に必要な観点

- Brewprint が解決したい intent transfer は、単なる requirements wording の標準化か、architecture model の生成・照会・修正・検証まで含む closed loop か。
- 既存 spec-driven agent workflow を参照しつつ、Brewprint 固有の YAML DSL / MCP semantic query / canonical ref の価値をどこに置くか。
- LLM が起草した artifact の正確性を、人間承認だけで担保するのか、schema / validation / trace / review gate で担保するのか。

#### 後続判断先

- LLM-first 原則 ADR の中で intent transfer problem を定義するか。
- Requirement / investigation / spec authoring guidance に、曖昧さの明示・未確定点・validation-ready contract を追加するか。

### Q3: Brewprint の現行 artifact model は intent transfer interface として機能しているか

#### 確認対象

- `docs/requirements/`、`docs/investigations/`、`docs/adr/`、`docs/spec/`、YAML、MCP の現行 ownership
- 人間の意図が、LLM の構造化作業を経て現行仕様・implementation source へ至る変換過程で失われるものがないか

#### 観測事実

- `docs/requirements/` は要求・不足・要望・spec gap 候補を stable identity で捕捉する。これは、人間の意図を未確定のまま失わず保持する layer と解釈できる。
- `docs/investigations/` は根拠、不確実性、選択肢、後続候補を保存する。これは、LLM が曖昧な意図を即座に仕様化せず、検証可能な問いへ分解する layer と解釈できる。
- `docs/adr/` は判断理由と却下案を保持し、`docs/spec/` は現行 contract の唯一の正を保持する。したがって intent から contract への変換における decision ownership と canonical output は既に分離されている。
- brewprint DSL YAML は対象 design model の primary implementation source とされ、MCP は LLM に structured query interface を提供する。これは precise artifact を agent が読む側の設計として既に明確である。
- 現行 Design Records MCP は read / resolve / validation を中心としており、writer tools は future extension とされている。この read-only 境界は、LLM-first の不足を直ちに示すものではなく、まず最低限の artifact 運用と実装を成立させ、実際に dogfooding してから write contract の必要性と形を判断するための意図的な MVP narrowing と解釈すべきである。
- 現行 docs は既に front matter、semantic ref、canonical reference rule、artifact ownership を構造化している。したがって起票時点で問うべき不足候補は「writer tool がまだないこと」そのものではなく、現行の read / resolve / validation と構造化 metadata が、LLM と協働した実運用で必要十分かをまだ dogfooding により確認していないことである。
- 将来 write を導入する場合に dry-run / approval / validation / provenance が論点になる可能性はあるが、その具体 contract は、read-first foundation の実運用で観測された摩擦・誤読・不足項目を根拠に判断すべきであり、本 investigation の起票時点で必須機能または欠陥として扱わない。

#### 候補

| candidate | interpretation | gap |
| --- | --- | --- |
| Candidate A: 現行 artifact model と read-first MCP をまず dogfooding する | 既存 boundary と MVP narrowing を維持し、必要十分性を実運用で観測する | 評価観点や観測記録が曖昧だと不足を取り逃す |
| Candidate B: Dogfooding の観測に基づき既存 artifact / MCP contract を refine する | 不足が確認された項目だけを追加・変更する | 観測前に具体変更を前提化しない運用が必要 |
| Candidate C: 起票時点で writer contract または新 intent layer を追加前提にする | 将来像を先に整えられる | requirements / investigation との重複や過剰設計を招く危険が高い |

#### 判断に必要な観点

- 現行 front matter / semantic ref / investigation metadata / read-only MCP が、LLM と協働した実際の設計作業でどの問いに答えられ、どの誤読や手戻りを防げないか。
- Intent の捕捉は `docs/requirements/` と `docs/investigations/` の現行運用で十分か、それとも dogfooding で確認された不足に限って provenance / ambiguity / approval metadata を追加すべきか。
- Writer tools を再判断する trigger を、read-first 運用で観測されるどの摩擦・反復作業・安全上の制約として定義すべきか。

#### 後続判断先

- Project artifact model に LLM-mediated intent-to-artifact flow を追加する必要が dogfooding で確認されるか。
- M19 の read / resolve / validation foundation を用いた dogfooding の観測方法・評価項目を requirement / work item / investigation のどこで捕捉するか。
- Writer tool または metadata refinement は、観測された不足が成立した場合にのみ後続 requirement として捕捉するか。

### Q4: LLM-first architecture tooling として参照すべき先行事例は何か

#### 確認対象

- MCP を利用した design / architecture artifact tool の公開事例
- LLM / agent が primary consumer または author として spec / tasks を扱う公開 workflow
- Brewprint と同一ではないが、設計判断へ転用できる近接例

#### 観測事実

- MCP は、server が resources、tools、prompts を client / model に公開するための protocol であり、LLM に外部構造・操作を公開する設計基盤として Brewprint の MCP 方針と整合する。MCP 自体は architecture documentation の domain model を提供しないため、brewprint が何を query / validate / write 可能にするかは引き続き固有設計となる。
- GitHub Spec Kit は、AI coding agent が specification、planning、task breakdown を利用する spec-driven development workflow を公式に公開している。Architecture DSL / semantic trace tool の代替とは確認できないが、人間の要望から persistent artifacts を介して agent execution に渡す先行例として参照価値がある。
- Kiro Specs は、自然言語の feature description から requirements、design、tasks を生成し、agentic implementation へ接続する workflow を公式に説明している。これも Brewprint の DSL / MCP / trace model と同一ではないが、intent-to-artifact pipeline を product surface として扱う近接例である。
- 起票時点の公開情報調査では、Brewprint と同様に architecture description DSL、Mermaid human rendering、MCP semantic query、ADR / investigation / spec ownership、semantic trace foundation を一体で提供する OSS は確認できていない。これは「存在しない」という断定ではなく、本調査時点での検索結果の範囲で未確認という意味である。

#### 候補

| reference class | example | what to learn | what it does not establish |
| --- | --- | --- | --- |
| Protocol for model-facing tools | MCP | discoverable tools / resources / capability boundary | Architecture semantics や artifact ownership |
| Spec-driven AI development workflow | GitHub Spec Kit | spec / plan / tasks を agent workflow に置く方法 | Brewprint DSL や traceability backend の代替 |
| Productized agent specs workflow | Kiro Specs | requirements / design / tasks の生成・反復 UI | OSS backend compatibility または canonical ref model |
| Human-oriented trace tools to adapt | OFT / Doorstop | coverage / lifecycle / publication capability | LLM I/O の native suitability |

#### 判断に必要な観点

- Brewprint は architecture description language を核にした LLM-native tooling なのか、spec-driven development workflow の architecture layer なのか、または双方を接続する semantic artifact system なのか。
- 先行製品から workflow UX を学ぶ一方、Brewprint 固有の semantic model / canonical ref / MCP tool surface を失わない境界を置けるか。

#### 後続判断先

- Positioning / project overview の refinement 候補。
- External tool / adjacent product comparison を継続する場合の評価カテゴリ定義。

### Q5: LLM-first 原則は Brewprint の今後の設計へ何を要求するか

#### 確認対象

- 現行 spec / ADR / doc-policy と新しい設計意図の整合性
- OFT / Doorstop interop 評価への追加軸
- Roadmap / tool contract / artifact guidance への影響候補

#### 観測事実

- `docs/doc-policy.md` と `spec:project-artifact-model` は、MCP が LLM 向け structured interface であるという方向を既に含むため、LLM-first の思想とは矛盾しない。
- 現在の design records / semantic trace MVP は canonical read / resolve / validation foundation に意図的に限定されており、まず最低限の運用と実装を成立させて dogfooding する段階にある。LLM による artifact creation / update、human approval gate、追加の intent provenance / ambiguity metadata は、実運用で必要性が確認された場合の future decision として扱うべきである。
- INV-DOCS-004 / 005 は、OFT / Doorstop の identity / ownership 境界を評価しているが、外部 backend を LLM が使用するための tool exposure、structured diagnostics、incremental query、write safety、connector maintenance cost を独立評価項目としてまだ定義していない。

#### 候補

| candidate | meaning | benefit | risk / question |
| --- | --- | --- | --- |
| Candidate A: LLM-first を説明文に留める | 現行 scope を変えない | 追加設計を抑えられる | 外部 tool 判断と writer roadmap が従来の human-first 前提へ流れ得る |
| Candidate B: LLM-first を project principle として ADR で明文化し、まず read-first dogfooding と外部 tool 評価 rubric に反映する | 現行 MVP narrowing を維持しつつ意思決定軸を明確化 | connector cost と intent transfer の評価を first-class にできる | dogfooding 前に writer / metadata 拡張まで先取りしない境界が必要 |
| Candidate C: LLM-first を理由に外部 tool を原則排除する | native design を優先 | boundary は単純になる | 外部 validator / lifecycle engine の再利用価値を不当に捨てる |

#### 判断に必要な観点

- LLM-first は「すべて自作する」原則ではなく、「LLM が semantic に利用できる総コストと正確性を基準に backend / adapter / native implementation を選ぶ」原則として表現すべきか。
- External tool を採用する場合の評価条件を、read/query、diagnostics、identity mapping、testability、connector maintenance cost にまず分解し、write/update や追加 provenance は必要性が観測された場合の次段評価に分けるべきか。
- M19 の現行 read / resolve / validation foundation を進めて dogfooding し、その観測結果から writer / connector / intent-transfer refinement の判断を分離して捕捉できるか。

#### 後続判断先

- LLM-first principle ADR の起票候補。
- OFT / Doorstop 評価の再採点、および connector spike の gate 設定。
- Read-first dogfooding で不足が確認された場合の MCP writer / controlled mutation / review flow の requirement 化候補。

## 横断的な観測事実

### 1. LLM-first は競合排除の文言ではなく、build / buy / adapt 判断の評価軸である

OFT / Doorstop が model-facing native interface を公開していないとしても、そのことだけで不採用とは言えない。CLI や Python API を安定した MCP connector へ包み、Brewprint の canonical identity と validation boundary を侵害せず、保守コストが native implementation より低いなら採用候補になり得る。

逆に、外部 tool の coverage / document management 機能が充実していても、LLM が使うために identity mapping、write orchestration、diagnostic normalization、artifact projection、approval lifecycle を Brewprint 側で再実装する必要があるなら、表面的な機能重複だけで buy 判断をするのは不適切である。

### 2. OFT と Doorstop は LLM adapter cost でも同一評価にならない

OFT は Markdown / source-code markers と CLI / report workflow を中心に公式説明しているため、LLM に対する fine-grained query / controlled update を提供するには、parser / runner / diagnostic projection / identity bridge をどこまで作る必要があるかを調べる必要がある。

Doorstop は CLI だけでなく Python scripting interface と repository-resident item model を持つため、structured access の adapter を作る候補 surface は OFT より広い可能性がある。一方で、Doorstop は Brewprint の requirements / spec ownership により深く重なるため、connector が作りやすくても semantic ownership conflict が大きければ採用しにくい。

したがって、機能適合性、ownership 適合性、agent I/O 適合性、adapter cost は分離して評価すべきである。

### 3. Brewprint の現行 artifact model は intent transfer の基礎を既に持ち、まず dogfooding で十分性を測るべきである

Requirements は未確定の必要性を保持し、investigation は不確実性と根拠を構造化し、ADR は判断を、spec は有効な contract を、YAML は design model implementation source を所有する。Front matter、semantic ref、canonical reference rule、MCP read / resolve / validation も含め、この分離は曖昧な会話から即時に実装へ飛ばず、LLM が intermediate artifacts を通じて精密化するための substrate として既に相応に設計されている。

この段階で不足と断定できるのは writer tool の不在ではない。確認すべきなのは、意図的に read-first とした MVP を実際に LLM 協働で運用したとき、現行 metadata と query / validation contract が必要十分か、どこで誤読・反復・手作業・承認判断の曖昧さが発生するかである。LLM-first の明文化や将来の controlled write / provenance / ambiguity metadata は、その dogfooding 観測を受けて refinement 候補とするのが妥当と考えられる。

### 4. Brewprint の位置付けは既存カテゴリの単純な一つでは未確定である

Brewprint は architecture description language の性質を持つ一方、LLM 向け MCP query surface、artifact ownership、将来の controlled write / trace validation を重視している。起票時点では、単なる ADL、単なる requirements management tool、単なる traceability tool のいずれかに還元するより、**LLM-mediated architecture artifact system** または **LLM-native architecture specification and query layer** としての位置付け候補を検証する方が実態に近いと考えられる。

ただし、この名称やカテゴリは現段階の推奨候補に留まり、project positioning の確定は後続判断で行う。

## 候補比較

### External tool evaluation rubric candidate

| axis | question | Brewprint-native | OFT adapter candidate | Doorstop adapter candidate |
| --- | --- | --- | --- | --- |
| Functional fit | 必要な coverage / trace / lifecycle 能力を持つか | future scope は自作が必要 | deferred coverage に近い可能性 | trace / document workflow が広い可能性 |
| Ownership fit | `requirements` / `spec:` / ADR / investigation の owner を壊さないか | 現行 boundary を維持 | backend 限定なら維持しやすい可能性 | requirements / spec ownership と衝突しやすい可能性 |
| Read I/O for LLM | semantic item / relation / diagnostic を構造化 query できるか | MCP として設計可能 | connector / parser が必要か要確認 | Python API adapter が使えるか要確認 |
| Write I/O for LLM（後段評価） | read-first dogfooding 後に constrained create / update / review が必要と判明した場合、どう実現するか | 意図的に MVP 外。必要性を観測後に判断 | source marker 更新 orchestration が必要か要確認 | item CRUD / review flow の adapter が必要か要確認 |
| Diagnostic usability | LLM が修正行動へ変換しやすい error を返せるか | contract として設計可能 | report projection が必要か要確認 | validation result projection が必要か要確認 |
| Identity mapping | `spec:` と外部 item ID が二重 owner にならないか | 単一 owner を維持 | revision / item ID bridge が主要課題 | UID / document tree bridge が主要課題 |
| Provenance / approval（後段評価） | dogfooding で追加追跡が必要と判明した場合、人間の意図・LLM変更・承認をどう記録するか | 現行 artifact で不足が出るか先に観測 | external lifecycle との境界要確認 | review fingerprint 利用可能性はあるが ownership 要確認 |
| Connector cost | MCP 化・保守・version追従の費用は妥当か | engine 自作コスト | CLI / format adapter コスト | Python API adapter + semantic migration コスト |

### Strategy candidates

| candidate | description | benefit | risk |
| --- | --- | --- | --- |
| Candidate A: Functional-fit first | 既存 tool 機能が近ければ adapter を前提に採用評価する | 自作を避けやすい | LLM I/O / semantic bridge 費用を後から負う |
| Candidate B: LLM-first rubric first | 機能・ownership・agent I/O・connector cost を採用前に同じ重みで評価する | Brewprint の目的と整合し、比較が公平になる | 評価調査と小規模 spike が増える |
| Candidate C: Native-only | LLM-native でない外部 tool は採用対象外とする | interface 一貫性を保ちやすい | 再利用可能な validator / workflow を不当に排除する |

## 起票時点の推奨案

本 investigation の起票時点では、**Candidate B: LLM-first rubric first** を後続判断の前提候補とすることが妥当と考えられる。

理由は以下である。

1. Brewprint の現行 policy / artifact model は既に MCP と structured artifact ownership を持ち、今回明示された LLM-first 設計意図と方向性が整合する。
2. INV-DOCS-004 / 005 は外部 tool の機能と ownership の重要な論点を捕捉しているが、LLM が利用するための read / write / validate interface と connector maintenance cost を独立の採用軸として十分に扱っていない。
3. OFT / Doorstop は、公式に確認できた interface が異なるため、MCP native でないという一点だけで等しく不利と判定できない。Tool ごとに adapter surface、diagnostic projection、identity bridge、controlled write の費用を評価すべきである。
4. LLM-first は外部 tool 排除の原則ではなく、external backend を採用する場合にも LLM-mediated workflow 全体の正確性・摩擦・保守費用を評価する原則として扱う方が合理的である。
5. Brewprint artifact model と read-first MCP foundation は intent transfer の substrate として有望であり、まず dogfooding によって現行項目・参照・diagnostic の必要十分性を測るべきである。LLM による controlled write、追加の ambiguity / approval / provenance contract は、不足が観測された場合の後続判断として明示する価値がある。

## 後続判断に渡す候補

- LLM-first を project-wide design principle として ADR に記録し、「人間が意図と承認を担い、LLM が structured artifact の起草・探索・検証・許可された更新を担う」境界を正式化するか。
- `spec:project-artifact-model` に、intent capture → investigation / decision → current spec → YAML / implementation → validation / query という LLM-mediated flow を追加するか。
- INV-DOCS-004 / 005 の後続評価で、functional fit / ownership fit に加え、agent read I/O、agent write I/O、diagnostic usability、identity bridge、approval / provenance、connector cost を必須 scoring axis にするか。
- M19 の read / resolve / validation foundation を使った dogfooding を先に行い、現行 front matter / metadata / query surface / diagnostics の必要十分性と、LLM が詰まる箇所を観測するか。
- Dogfooding により具体的な不足が確認された場合に限り、Design Records MCP の future writer tools を dry-run diff、validation-before-commit、human approval、source attribution、rollback / conflict detection を含む controlled mutation boundary として requirement 化するか。
- `docs/requirements/` または `docs/investigations/` に、会話由来の intent、未確定前提、解消すべき ambiguity、ユーザー承認点を記録する追加 metadata / authoring guidance が必要かを、dogfooding の観測結果に基づいて判断するか。
- External backend connector spike を行う場合、OFT / Doorstop の機能だけでなく「LLM が一回の tool session で問いを解き、診断を理解し、安全な修正候補を出せるか」を acceptance criterion にするか。

## 後続 artifact 候補

| category | candidate artifact / action | purpose |
| --- | --- | --- |
| ADR | LLM-first design principle and intent transfer boundary | 今回明示された設計意図を project decision として正式化する場合 |
| Project artifact model spec | `spec:project-artifact-model` refinement | human intent / LLM mediation / artifact lifecycle の関係を ownership map に追加する場合 |
| Traceability / MCP spec | `spec:trace` / Design Records MCP future contract | agent query / controlled write / validation / external backend adapter boundary を扱う場合 |
| Requirement / work item | read-first LLM dogfooding evaluation | 現行 artifact / MCP contract の必要十分性と不足を観測する場合 |
| Requirement | MCP writer / controlled mutation / provenance / ambiguity handling requirement | dogfooding により追加能力の必要性が確認された場合 |
| Investigation update | INV-DOCS-004 / INV-DOCS-005 refinement | OFT / Doorstop を LLM-first rubric で再評価する場合 |
| Comparative investigation | external backend comparison with agent I/O scoring | OFT / Doorstop / Brewprint-native の採否判断へ進む場合 |
| Spike | minimal connector / agent task evaluation | adapter cost と実際の LLM usability を測定する場合のみ |

## 未確定点

- ユーザーが明示した LLM-first 設計意図を、どの粒度で accepted ADR / confirmed spec に反映すべきか。
- Read-first dogfooding をどの実タスク・評価項目・記録 artifact で行い、現行 front matter / canonical refs / validation の必要十分性を判断するか。
- Intent transfer の primary owner は既存 requirements / investigation の運用で足りるか、dogfooding 後に refinement または別 mechanism が必要と判明するか。
- Human approval と LLM-controlled write の境界は、write 必要性が確認された場合に MCP writer tool、git workflow、artifact lifecycle のどこで所有すべきか。
- OFT の CLI / report interface を MCP connector 化した場合の実装・保守コストと、OFT の coverage engine を再利用する価値の比較。
- Doorstop の Python scripting surface を MCP 化した場合の connector cost と、Doorstop document / item ownership conflict の比較。
- GitHub Spec Kit / Kiro Specs 等の workflow から学ぶべき要素と、Brewprint 固有の architecture DSL / semantic ref / MCP query model を維持すべき境界。
- 起票時点で未確認の、MCP を前提に architecture / design documentation を扱う既存 OSS / research が存在するか。

## 起票時点の制約記録

- 本 artifact の `source_refs` は、MVP canonical reference rule に従い、Brewprint 内の record ID-as-ref または active `spec:` semantic ref に限定した。外部 source は本文の調査対象としてのみ記載した。
- ユーザーが本会話で明示した LLM-first 設計意図は、本 investigation の trigger および設計仮説として扱い、accepted ADR / confirmed spec と同じ authority としては扱っていない。
- OFT / Doorstop については、公式資料で確認できる interface / workflow を観測事実とし、「human-first」「pre-LLM」等の設計思想は、公式に裏付けられない限り断定しない。
- 起票時点の external review では、OFT と Doorstop に公式 MCP surface があることは確認できなかった。一方、Doorstop は Python scripting surface を持つため、OFT と同一の adapter cost であるとは扱わない。
- 本起票では、新規 investigation artifact の作成と investigation 一覧の同期以外に、ADR、spec、task、requirement、work item、implementation を変更しない。
