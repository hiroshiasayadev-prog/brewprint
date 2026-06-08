# V01-ADR-087: Design Records MCP investigation support and semantic ref resolve

- **status**: accepted
- **date**: 2026-05-23
- **depends_on**: V01-ADR-076, V01-ADR-077, V01-ADR-081, V01-ADR-084, V01-ADR-085, V01-ADR-086
- **supersedes**:
- **migrated_to_spec**:

> このADRは起票時点での決定を記録したスナップショットである。
> 現在の仕様は spec を参照すること。

## 背景

V01-ADR-076 / V01-ADR-077 は、Design Records MCP を ADR / spec の machine-readable metadata に対する read-only index / query / validation layer として導入し、`list_records` / `get_record` / `validate_records` を中心とする tool boundary を決定した。

その後、V01-ADR-085 / V01-ADR-086 により `docs/investigations/` が導入され、investigation artifact の責務、format、lifecycle が確定した。investigation は決定や現行仕様を所有しない一方、調査根拠、未確定点、選択肢、推奨案、後続 artifact 候補を保存する。V01-ADR-086 は、investigation を MCP index / query / validate 対象にする可能性を認めたが、Design Records MCP に統合するか、別 MCP interface とするかは後続判断とした。

また、M18 semantic traceability foundation では、Design Records MCP に semantic ref resolve を入れるか、investigation をどう扱うか、investigation metadata を MCP 側で補完・検証するかが後続判断として残されている。

設計対話では、ADR / spec と investigation を横断して調査根拠と後続判断を辿る必要がある。investigation だけを別 interface に分離すると、record 探索と参照解決の導線が分断される。一方、investigation は decision ではないため、lifecycle や kind 固有 metadata を decision/spec と平坦に混在させてもならない。

さらに、V01-ADR-086 は初期運用として investigation の `source_refs` / `follow_up_candidates` / `follow_up_results` に human-readable な artifact ID または path を許容した。しかし semantic traceability は physical layout 変更と trace の安定性を分離する方針であり、調査根拠となる `source_refs` を physical path に依存させ続ける根拠はない。

## 決定

### 1. Design Records MCP は investigation を record kind として追加する

Design Records MCP の既存対象である `decision` / `spec` に加えて、`investigation` を index / query / validation 対象として追加する。

investigation は Design Records MCP と別の MCP interface に分離しない。
ADR / spec / investigation は責務と lifecycle が異なるが、設計対話で参照・探索・検証される artifact record としては同一 index 上で扱う。

本ADRは Design Records MCP の record kind 全体を閉じた列挙として確定するものではない。V01-ADR-081 で将来拡張として予約された `kind: requirement` を含め、後続判断により他の artifact kind を追加する可能性を制限しない。

この決定は investigation を decision と同一視するものではない。V01-ADR-085 / V01-ADR-086 の責務境界は維持する。

- decision は判断履歴を所有する
- spec は現行仕様を所有する
- investigation は調査結果、根拠、選択肢、推奨案、後続 artifact 候補を所有するが、決定を所有しない

### 2. 既存の record-oriented tool 名を維持する

Design Records MCP の既存tool名は維持する。

- `list_records`
- `get_record`
- `validate_records`

これらの `record` は ADR/spec に限定された呼称ではなく、Design Records MCP が query / validate する artifact record を意味するものとして明確化する。

`list_adrs` のような ADR 固有名への rename は行わない。investigation と spec を含む index の名称として不適切になるためである。

### 3. record response は共通 field と kind 固有 detail を分離する

Design Records MCP の record response は、全kindに共通する field と、kind 固有の detail object を分離する構造を採用する。

共通 field は少なくとも以下を持つ。

- `id`
- `kind`
- `title`
- `status`
- `path`

本ADR時点で対象となる `decision` / `spec` / `investigation` の kind 固有 metadata は、それぞれの detail object に保持する。後続判断により追加される artifact kind も、同じ拡張方式で固有 detail object を持てる。

概念例:

```json
{
  "id": "V01-INV-MCP-001",
  "kind": "investigation",
  "title": "Design Records MCP investigation support",
  "status": "concluded",
  "path": "docs/investigations/mcp/INV-MCP-001-design-records-mcp-investigation-support.md",
  "investigation": {
    "trigger": "V01-ADR-086",
    "scope": "investigation MCP integration",
    "non_scope": "writer tools",
    "source_refs": ["V01-ADR-086"],
    "follow_up_candidates": ["V01-ADR-087"],
    "follow_up_results": ["V01-ADR-087"]
  }
}
```

Go implementation では、共通 record struct と kind 固有の optional detail struct を組み合わせる形を第一候補とする。具体的な Go type と JSON schema は spec / implementation task で定める。

この response 構造への移行は、既存 `decision` / `spec` response の flat field contract の変更を伴う。Design Records MCP は brewprint 内部向けの開発中 interface であるため、本ADRでは新しい response 構造への破壊的な切替を許容し、既存 flat field と detail object の併存期間または compatibility adapter を要求しない。

ただし、contract が不整合な中間状態を作らないため、Design Records MCP spec を新しい response 構造へ更新した後、実装および対応する tests は同一の切替単位で追従する。spec が旧 flat response のままである状態で、実装だけを先行して変更してはならない。

### 4. Design Records MCP は semantic ref resolve の責務を持つ

Design Records MCP は、record の list / get / validate に加えて、docs artifact 間の semantic/artifact ref を解決する責務を持つ。

semantic ref の解決を上位 agent や利用者の文字列照合に委ねない。resolver は、参照元 metadata から参照先 artifact または semantic anchor を機械的に解決し、validation がその結果を利用できる境界を提供する。

resolver が解決に利用する lookup source と、`list_records` / `get_record` が公開する record kind は同一集合である必要はない。resolver が requirement / work item / internal-design / coverage その他の artifact を解決対象に含めても、それだけで当該 artifact を Design Records MCP の record kind として list / get 対象に追加する決定にはならない。record kind の追加は、各 artifact の lifecycle と query 需要を踏まえて別途判断する。

本ADRでは resolve 用toolの名称、request / response schema、prefix grammar の完全形、diagnostic response shape は確定しない。これらは traceability spec および Design Records MCP spec で定義する。

### 5. investigation の `source_refs` は canonical reference として解決・検証する

investigation の `source_refs` は調査根拠を表す required metadata である。
そのため、`source_refs` に記載された参照は resolver が解決可能でなければならず、解決不能な参照は validation error とする。

`source_refs` の canonical reference は、resolver が解決可能な artifact ID または semantic ref とする。
physical path は canonical reference として扱わない。

V01-ADR-086 が初期運用の便宜として許容した path-based reference は、既存documentのcompatibility inputとして読み取る必要がある場合に限り許容してよい。ただし、canonical form ではなく、validator は noncanonical reference として診断できるものとする。

### 6. investigation の `follow_up_results` は存在するartifactへの参照として解決・検証する

`follow_up_results` は、この investigation を根拠に実際に作成・更新された artifact を記録する optional metadata である。

記載されている場合、参照先 artifact は存在する前提であるため、resolver が解決可能でなければならず、解決不能な参照は validation error とする。

`follow_up_results` についても canonical reference は resolver が解決可能な artifact ID または semantic ref とし、physical path は canonical reference として扱わない。

### 7. investigation の `follow_up_candidates` は存在検査を要求しない

`follow_up_candidates` は、調査結果から起票・更新されうる artifact を表す required metadata である。候補には、まだ存在しない後続 artifact が含まれうる。

そのため、`follow_up_candidates` は記法・参照形式の検査対象にはできるが、参照先 artifact の存在は required validation としない。

candidate が実際に作成・更新された場合は、`follow_up_results` に記録し、その時点から解決・存在検査の対象とする。

### 8. validation と具体contractの分担

本ADRで必須とする validation rule は以下である。

- investigation の `source_refs` の unresolved は error とする
- investigation の `follow_up_results` が記載されている場合、その unresolved は error とする
- `source_refs` / `follow_up_results` に physical path を canonical reference として使用しない
- investigation の `follow_up_candidates` は未作成 artifact を指しうるため、存在しないこと自体を error としない

以下は本ADRでは確定しない。

- semantic ref resolve tool の名称と入出力schema
- active prefix / artifact ID grammar の完全形
- noncanonical path reference diagnostic のcategory名とseverity
- `trigger` / `related_*` の resolve / validation rule
- coverage mapping query のDesign Records MCPへの追加可否
- MCP writer tools
- 新しい response 構造への切替に必要な具体的な実装手順および test 更新内容

既存 flat response との互換adapterまたは併存期間を設けるかどうかは未決事項ではない。本ADRは、それらを設けず新しい response 構造へ切り替える方針を決定する。

これらは spec または後続 ADR / task で扱う。

## 理由

### なぜ investigation を同一 MCP に統合するか

investigation は決定を所有しないが、ADR/specを根拠として読み、後続ADR/spec/task等へ判断材料を渡す artifact である。

設計対話において、調査結果だけを別interfaceから取得し、根拠となるADR/specを別indexから辿る構成は不要な分断になる。Design Records MCP は設計に関わる record の探索・取得・検証を担うため、investigation も同じrecord indexに載せる方が自然である。

### なぜ kind 固有 detail を分離するか

`decision` / `spec` / `investigation` は status vocabulary も metadata semantics も異なる。

investigation の `trigger` / `source_refs` / `follow_up_candidates` を、decision の `depends_on` / `migrated_to_spec` と同一の平坦field群として扱うと、どのfieldがどのkindで有効かが曖昧になる。

共通fieldとkind固有detailを分けることで、record横断queryとartifact種別固有の責務境界を両立できる。

### なぜ semantic ref resolve を MCP responsibility にするか

semantic/artifact ref は設計対話で機械的に辿るための参照であり、上位agentの自然言語推定や都度のpath探索に委ねるべきではない。

MCPが resolver と validation の境界を所有すれば、異なるagentやclientでも同一の解決規則と参照切れ診断を利用できる。

### なぜ `source_refs` の参照切れを error にするか

`source_refs` は調査根拠である。存在しない根拠を指す investigation は、調査内容を再検証できず、後続判断の信頼性を損なう。

参照切れを許容して後から追跡する運用は、investigation をmachine-queryableにする目的と矛盾する。そのため unresolved `source_refs` は必須検査で error とする。

### なぜ physical path を canonical reference にしないか

semantic traceability は、file rename や section move といった physical layout 変更から trace identity を分離する必要がある。

physical path を根拠参照のcanonical formにすると、artifactの意味が変わっていなくてもfile移動だけで参照が壊れる。pathはsource locationや移行補助には使えても、長期安定参照の正規形には適さない。

### なぜ `follow_up_candidates` の存在を要求しないか

`follow_up_candidates` は未来に起票されうるartifactの候補を記録するfieldであり、調査 concluded 時点でも未作成であることが正当である。

存在検査を必須にすると、候補と結果の区別が崩れ、V01-ADR-086が分離した `follow_up_candidates` / `follow_up_results` の意味を壊す。

## 却下した代替案

### 代替案A: investigation は別 MCP interface で扱う

- 利点: investigation固有のlifecycleを隔離しやすい
- 欠点: ADR/specとの根拠参照、後続判断、validation導線が分断される

→ 却下。責務はkindで分離し、index / query / validationはDesign Records MCPに統合する。

### 代替案B: MCP tool を `list_adrs` 等へrenameする

- 利点: 初期MVPの実態に近い名称になる
- 欠点: specとinvestigationを含む横断record queryの名称として成立しない

→ 却下。`list_records` / `get_record` / `validate_records` を維持し、recordの定義を明確化する。

### 代替案C: 全kindのmetadataを平坦なrecord responseに載せる

- 利点: JSON shapeが一見単純になる
- 欠点: kind別責務と有効fieldが曖昧になり、誤用しやすい

→ 却下。共通field + kind固有detailを採用する。

### 代替案D: `source_refs` の参照切れはwarningに留める

- 利点: 移行中の運用は軽い
- 欠点: 調査根拠が存在しない状態を正当化し、後続判断時の追跡負荷を残す

→ 却下。unresolved `source_refs` は error とする。

### 代替案E: `source_refs` にphysical pathをcanonical formとして残す

- 利点: semantic ref未整備artifactを即座に指しやすい
- 欠点: file layout変更で根拠参照が壊れ、semantic traceの安定性方針に反する

→ 却下。canonical referenceはresolverが解決可能なartifact IDまたはsemantic refとし、pathはcompatibility扱いに限定する。

## 影響

### Design Records MCP spec への影響

Design Records MCP spec では、少なくとも以下を反映する必要がある。

- `kind: investigation` の追加
- `list_records` / `get_record` / `validate_records` が investigation を扱うこと
- 共通field + kind固有detail object のresponse schema
- 既存 flat response との併存を行わず、新しい response schema に spec / implementation / tests を同一変更単位で切り替える方針
- investigation metadata のparse / validation rule
- semantic ref resolve のtool contractまたは既存toolとの組合せ
- unresolved `source_refs` / `follow_up_results` のdiagnostic
- noncanonical path reference のdiagnostic

### Traceability spec への影響

`docs/spec/concepts/traceability/resolve-and-validation.md` は、resolver lookup source と orphan detection 対象に investigation metadata を追加する必要がある。

少なくとも以下を整理する必要がある。

- investigation `source_refs`
- investigation `follow_up_results`
- investigation `follow_up_candidates` の存在検査除外
- artifact ID / semantic ref とphysical pathのcanonicality境界

### docs/investigations/README.md への影響

READMEは、V01-ADR-086の初期運用として記載した path許容を本ADRに合わせてrefineする必要がある。

- `source_refs` / `follow_up_results` は canonical reference としてartifact IDまたはsemantic refを使う
- physical pathはcanonical referenceとしない
- `follow_up_candidates` は未作成artifact候補を記述できる

### M18 task への影響

`docs/tasks/m18-semantic-traceability-foundation.md` の Phase F に残っていた以下の判断は、本ADRにより解決される。

- Design Records MCP に semantic ref resolve を入れるか
- Design Records MCP に `kind: investigation` を追加するか、別 MCP interface とするか
- investigation の `source_refs` / `follow_up_results` を MCP 側で検証するか

一方で、具体 request / response schema、coverage mapping query、writer tools、migration作業は引き続き後続作業として扱う。

### V01-ADR-086 への影響

V01-ADR-086 の investigation format / lifecycle と、investigation が決定を所有しないという境界は維持する。

本ADRは、V01-ADR-086が後続判断に送ったMCP interface方針と、初期運用で暫定許容したreference記法のcanonicalityをrefineする。
V01-ADR-086本文は起票時点のスナップショットとして遡及修正しない。

### 実装への影響

具体的な実装項目、移行順序、完了条件はDesign Records MCP implementation taskで追跡する。

実装上は、少なくとも以下の検討が必要になる。

- investigation document parser / index entry
- kind固有detail response serialization
- semantic/artifact ref resolver
- unresolved / noncanonical reference diagnostics
- existing decision/spec response compatibility

## Evidence

- commit: 81deb03
- impl commit: tbd
- 参考: V01-ADR-076 Design Records MCP、V01-ADR-077 MVP boundary and tool prioritization、V01-ADR-084 semantic trace MVP scope and artifact boundary、V01-ADR-085 investigation artifact boundary、V01-ADR-086 investigation artifact format and lifecycle、`docs/tasks/m18-semantic-traceability-foundation.md`、`docs/spec/concepts/traceability/resolve-and-validation.md`、`docs/spec/concepts/traceability/semantic-ref.md`
