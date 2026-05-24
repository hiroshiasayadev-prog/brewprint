# INV-DOCS-005: Doorstop artifact model overlap and adoption boundary

- **status**: investigating
- **date**: 2026-05-24
- **trigger**: INV-DOCS-004 による OpenFastTrace (OFT) 評価の開始後、Doorstop が requirements / design / test document tree と traceability validation を提供する既存 OSS 候補として判明し、Brewprint の requirements / spec / future trace artifact ownership と競合し得る範囲を別途確認する必要が生じた
- **scope**: Doorstop の document / item / link / review / reference model が Brewprint の requirements layer、spec-first ownership、semantic trace foundation、および deferred realization coverage に与える影響と代替可能性
- **non_scope**: Doorstop 採用決定、OFT との最終比較結論、既存 artifact layer の廃止、ADR-088 の撤回、M19 implementation の変更、Doorstop project tree の試験導入
- **source_refs**:
  - ADR-081
  - ADR-083
  - ADR-088
  - INV-DOCS-004
  - spec:project-artifact-model
  - spec:trace.artifact-refs
  - spec:trace.coverage-mapping
- **follow_up_candidates**:
  - ADR refinement for external requirements / traceability tool ownership boundary
  - spec:project-artifact-model
  - spec:trace.artifact-refs
  - spec:trace.coverage-mapping
  - requirements layer boundary clarification for external trace tools
  - comparative recommendation across OFT / Doorstop / Brewprint-native continuation
  - Doorstop compatibility spike only if artifact ownership fit remains plausible

> 本 investigation は Doorstop の採用を決定するものではない。Doorstop は OFT と同一の問いではなく、coverage validator 候補に留まらず、requirements / document tree / review lifecycle の owner になり得るため、別 investigation として扱う。
> ADR-088 により deferred とされた realization coverage を MVP に戻す判断、または ADR-081 が定める requirements / spec の責務境界を変更する判断は、本 artifact では行わない。

## 調査スコープ

- Doorstop の公式 documentation に基づき、document tree、item identity、parent link、review / suspect link、external reference、validation / publish のモデルを確認する。
- Doorstop が Brewprint の deferred realization coverage の候補に留まるのか、`docs/requirements/` / `docs/spec/` / future assurance artifact まで代替対象として侵入するのかを切り分ける。
- Brewprint の `REQ-*` capture layer と Doorstop の normative / traceable item model を同一視できるかを評価する。
- Brewprint の `spec:` canonical ref と Doorstop item UID / document hierarchy / referenced implementation path が共存可能かを確認する。
- OFT と Doorstop を同一採用候補として比較すべき範囲と、別々に評価すべき責務境界を整理する。

## 非スコープ

- Doorstop の document tree を Brewprint repository に作成しない。
- `docs/requirements/`、`docs/spec/`、`docs/internal-design/`、`docs/work-items/` の ownership を変更しない。
- `REQ-*` / `SPEC-*` / `INV-*` / `spec:` の現行 identity contract を変更しない。
- M19 Phase A / implementation / acceptance test を更新しない。
- OFT と Doorstop のどちらを採用するかを本 investigation の起票時点で決定しない。
- ライセンス、配布、組み込み方式は利用形態が候補として成立した場合のみ後続判断へ渡す。

## 背景

ADR-081 は、`docs/requirements/` を要求・不足・要望・spec gap 候補の捕捉 layer とし、現行仕様の source of truth ではないと定めた。Accepted requirement であっても、外部 contract とするには `docs/spec/` へ反映する必要がある。

ADR-083 と現行 `spec:project-artifact-model` は、requirements、spec、work-items、internal design、trace mechanism、brewprint DSL YAML の ownership を分けている。ADR-088 は、semantic trace MVP を `spec:` と record ID-as-ref および investigation canonical reference の解決・検証 foundation に縮小し、coverage / realization relation を future decision へ送った。

INV-DOCS-004 は、OFT を deferred realization coverage の backend 候補として評価し、`spec:` identity と外部 coverage item identity / revision が衝突しないかを調査対象とした。

Doorstop は、起票時点で確認した公式 documentation 上、単に implementation / test coverage を検証する engine ではなく、requirements 等の document と item を repository 内で管理し、document hierarchy と item links を検証・公開する requirements management tool として位置付けられている。このモデルは OFT よりも Brewprint の requirements / spec / trace artifact system の所有境界に広く重なる可能性があるため、OFT 調査へ混在させず独立に評価する。

## 起票時点で確認した external sources

以下は Doorstop の公式資料として起票時点で確認した対象である。これらは Brewprint の canonical `source_refs` metadata ではなく、本文調査の external source として扱う。

- Doorstop documentation: `https://doorstop.readthedocs.io/`
- Doorstop document reference: `https://doorstop.readthedocs.io/en/latest/reference/document.html`
- Doorstop item reference: `https://doorstop.readthedocs.io/en/latest/reference/item.html`
- Doorstop validation / CLI documentation: `https://doorstop.readthedocs.io/en/latest/cli/validation.html`
- Doorstop repository: `https://github.com/doorstop-dev/doorstop`

## 起票時点の確認結果

### Q1: Doorstop はどの種類の tool か

#### 観測事実

- Doorstop は version control とともに要求文書を管理する requirements management tool として公開されている。
- Doorstop の document は directory 単位で作成され、prefix を持つ。Document は親 document を指定して hierarchy を構成できる。
- Doorstop の item は document 内の UID を持つ個別管理単位であり、item 間の link によって上位 document の item を参照する。
- Doorstop は validation と publish を提供し、要求・設計・テスト等の traceable document tree を管理する用途に向く。

#### 評価

Doorstop は OFT と同じく traceability に関係するが、評価対象となる責務は異なる。OFT は将来の coverage relation validator / backend 候補として評価しやすい一方、Doorstop は document / item / review workflow の owner 候補であり、Brewprint の artifact ownership そのものとの重複を先に評価する必要がある。

### Q2: Doorstop の document / item identity は Brewprint の artifact model とどこで重なるか

#### 観測事実

- Doorstop は document prefix と item UID により、requirement / design / test 等の管理単位を識別する。
- Brewprint は `REQ-*` を要求・不足・要望の捕捉 identity、`SPEC-*` を spec record identity、`spec:` を現行 design spec の semantic identity として分離している。
- Brewprint の requirements は source of truth ではなく、normative な現行仕様は spec が所有する。

#### 重複候補

| Brewprint layer / concern | Doorstop overlap candidate | 注意点 |
| --- | --- | --- |
| `docs/requirements/` と `REQ-*` | Requirement document / item UID | Brewprint requirement は pre-spec capture を含み、Doorstop item の normative / trace workflow と同一視できない可能性がある |
| `docs/spec/` と `spec:` | Design / requirement document item | `spec:` canonical identity と Doorstop UID が二重化し得る |
| future coverage / assurance | Parent-child link validation / suspect link workflow | OFT より ownership 範囲が広い可能性がある |
| tests / external implementation evidence | references / test document links | physical path や external reference の扱いが Brewprint の canonical ref 原則と衝突し得る |

#### 評価

Doorstop は deferred coverage mechanism だけでなく、すでに Brewprint が明示的に ownership を定めた requirements / spec artifact model にも重なり得る。採用可能性の評価には、機能充足ではなく、どの layer の source of truth を Doorstop に渡すことになるかの確認が必要である。

### Q3: Doorstop は Brewprint の `docs/requirements/` を置き換え得るか

#### 観測事実

- ADR-081 の requirement は、追加要件、要望、不足、spec gap 候補、deferred / rejected の履歴を捕捉する layer であり、accepted でもそれ自体は現行仕様ではない。
- Doorstop は traceable item と document hierarchy を扱い、validation / review の対象として item を管理する。

#### 候補

- Candidate A: Brewprint の `docs/requirements/` 全体を Doorstop requirement document として置換する。
- Candidate B: Adopted / normative な spec item または verification 対象のみ Doorstop 管理対象とし、capture requirements は Brewprint に残す。
- Candidate C: Doorstop を requirement / spec owner としては採用せず、参考比較対象に留める。

#### 評価

Candidate A は、Brewprint requirement の capture / decision-needed / rejected 履歴と、traceable normative item の境界を曖昧にする危険が高い。Doorstop を評価する場合でも、少なくとも初期仮説としては `docs/requirements/` 全体の直接置換を前提にすべきではない。

### Q4: Doorstop と Brewprint の canonical reference 原則は整合するか

#### 観測事実

- Brewprint の semantic trace MVP は、physical path を canonical reference にしない境界を明示している。
- Doorstop は item identity と links を持つ一方、外部 artifact / source への reference を管理する機構も備える。
- Doorstop が implementation evidence との接続に用いる reference / review modelが、Brewprint の `spec:` canonical ref と将来の YAML / implementation identity をどのように扱えるかは追加確認が必要である。

#### 判断に必要な観点

- Doorstop item UID を採用する場合、`spec:` とどちらを canonical identity とするか。
- External file reference は evidence / locator に限定し、canonical identity と分離できるか。
- Doorstop の suspect link / review lifecycle が Brewprint の future relation lifecycle として十分か、または ownership が過剰か。

#### 後続判断先

- Doorstop compatibility が成立し得る場合のみ、minimal document / item / link spike の要否を判断する。

### Q5: Doorstop と OFT は同じ比較軸で評価すべきか

#### 観測事実

- INV-DOCS-004 の中心問いは、OFT が deferred realization coverage backend になり得るか、および external tool identity / revision と Brewprint `spec:` identity の接続境界である。
- Doorstop の中心問いは、requirements / design / test document tree と item lifecycle を外部 tool に所有させた場合に、Brewprint の既存 artifact model を保てるかである。

#### 評価

OFT と Doorstop は最終的に比較対象にはなり得るが、最初から同じ investigation に統合すると責務境界が混ざる。先に INV-DOCS-004 で OFT の coverage backend 適合性、INV-DOCS-005 で Doorstop の artifact ownership overlap を独立に評価し、その後に比較判断へ集約する方が妥当である。

## 横断的な観測事実

### 1. Doorstop の発見は ADR-088 を即時に覆さない

Doorstop が traceability / requirements management 機能を提供していても、semantic realization relation を Brewprint MVP へ再導入する理由には直結しない。ADR-088 の MVP 縮小方針は、実装を先取りしない判断として引き続き成立する。

一方で、Doorstop の artifact model が requirements / spec ownership の代替候補になり得る場合、将来 coverage を再導入するか以前に、既存の Brewprint artifact model を維持するのかという別の判断が必要になる可能性がある。

### 2. Doorstop は OFT より Brewprint への侵入範囲が広い可能性がある

OFT 評価では、Brewprint の canonical identity と docs ownership を保持しながら coverage validation を外部化できるかが主題になる。Doorstop 評価では、document / item identity、links、review lifecycle、publish workflow が Brewprint の requirements / spec / future assurance ownership と競合し得る。

このため、Doorstop の star 数や利用実績が高い場合でも、それだけを理由に採用候補として優先できない。Brewprint との境界で重要なのは、外部 tool の成熟度に加え、既存の source-of-truth separation を保ったまま利用できるかである。

### 3. M19 への直接影響は OFT より間接的だが、判断は無視できない

M19 の直接 scope は `spec:` / record ID-as-ref / investigation canonical reference の resolve / validation foundation であり、Doorstop document tree の導入ではない。

ただし Doorstop 採用が spec / requirement item identity の owner を変更する可能性を持つなら、M19 が `spec:` contract を確定する前後で、Doorstop を採用する場合にも `spec:` を Brewprint canonical identity として維持するのかを確認する必要がある。

## 候補比較

| candidate | meaning | benefit | risk / question |
| --- | --- | --- | --- |
| Candidate A: Doorstop を Brewprint artifact system の owner 候補として評価する | requirements / spec / test trace workflow を Doorstop へ大きく寄せる可能性を検証する | 既存 OSS の文書管理・検証 workflow を広く活用できる可能性 | Brewprint の spec-first / capture requirements / semantic ref boundary を大きく変更し得る |
| Candidate B: Doorstop を normative trace subset のみの候補として評価する | Capture requirements と canonical Brewprint identity は維持し、限定的な verification tree だけ適合性を見る | ownership 衝突を限定できる可能性 | Doorstop の強みを部分利用する価値が十分か確認が必要 |
| Candidate C: Doorstop は比較対象に留め、OFT または Brewprint-native boundary を優先する | Doorstop に artifact ownership を渡さない | 現行境界を維持しやすい | Doorstop が提供し得る review / publish / trace workflow を活用しない |

## 起票時点の推奨案

本 investigation の起票時点では、**Candidate B を成立可能性の確認対象としつつ、Candidate A は採用前提にしない**ことが妥当と考えられる。

理由は以下である。

1. Doorstop は Brewprint の deferred coverage より広い artifact ownership に関与し得るため、全面置換を初期仮説にすると ADR-081 / ADR-083 / ADR-088 の責務分離を先に崩してしまう。
2. Brewprint の `docs/requirements/` は normative requirement set ではなく capture layer を含むため、Doorstop requirement item と直結する前提は危険である。
3. Doorstop の文書 tree / review / validation workflow に価値がある場合でも、normative spec または限定的 verification subset で試せるかを先に確認すべきである。
4. OFT と Doorstop は役割が異なるため、各調査の結果を得た後に、coverage backend、artifact system owner、または不採用の候補として比較する方が判断を誤りにくい。

## 後続判断に渡す候補

- Doorstop の公式 item / document / reference / review model を追加確認し、Brewprint の `spec:` canonical identity を維持した限定利用が成立するか。
- Doorstop item と `docs/requirements/` の意味差を踏まえ、capture requirement を外部 tool 管理対象から外す boundary を維持すべきか。
- Doorstop と OFT の調査結果がまとまった段階で、external tool comparative recommendation を別 investigation または本件の follow-up として作成すべきか。
- Doorstop compatibility spike を行う場合、既存 docs を移行せず、isolated sample document tree に限定すべきか。
- Doorstop 採用可能性が低い場合でも、Doorstop が提供する suspect link / review / publish workflow から Brewprint future requirements へ抽出すべき要件があるか。

## 後続 artifact 候補

| category | candidate artifact / action | purpose |
| --- | --- | --- |
| Investigation update | 本 investigation の更新 | Doorstop official model と Brewprint overlap の詳細結論を記録 |
| Comparative investigation | OFT / Doorstop / Brewprint-native comparative recommendation | 両 OSS 調査の結論を採用判断へ集約する場合 |
| ADR | external requirements / traceability tool ownership boundary refinement | artifact ownership へ影響する判断が必要と判明した場合のみ起票 |
| Spec | `spec:project-artifact-model`, `spec:trace.artifact-refs`, `spec:trace.coverage-mapping` | external tool boundary または non-goal を明文化する必要がある場合 |
| Spike | isolated Doorstop document-tree experiment | 限定利用の成立性が一次情報だけで判断できない場合 |

## 未確定点

- Doorstop の current item / document formats と、既存 Brewprint Markdown spec を直接または生成経由で扱える範囲。
- Doorstop link / review / fingerprint mechanism が、Brewprint が将来必要とし得る semantic realization lifecycle と一致するか。
- Doorstop の external reference が implementation / test evidence に利用される場合、physical path 非 canonical 原則と分離できるか。
- `spec:` を canonical identity に残したまま Doorstop の限定 subset を利用する価値が、OFT coverage backend 候補より高いか。
- Doorstop の publish / review workflow が Brewprint の project artifact needs として実際に必要か、単に既存 tool の機能であるだけか。
- OFT / Doorstop の個別評価後に、採用判断用の比較 investigation を新たに起票すべきか。

## 起票時点の制約記録

- 本 artifact の `source_refs` は、MVP canonical reference rule に従い、Brewprint 内の record ID-as-ref または active `spec:` ref に限定した。Doorstop の外部文書は本文調査の external source として扱い、metadata の `source_refs` には含めない。
- 起票時点で確認できた外部事実は Doorstop 公式 documentation / repository に基づく初期観測であり、Doorstop compatibility spike や adoption decision を意味しない。
- 本起票では、新規 investigation artifact の作成と investigation 一覧の同期以外に、ADR、spec、task、requirement、work item、implementation の変更は行わない。
