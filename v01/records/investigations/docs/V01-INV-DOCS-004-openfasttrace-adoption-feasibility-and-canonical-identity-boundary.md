# V01-INV-DOCS-004: OpenFastTrace adoption feasibility and canonical identity boundary

- **status**: concluded
- **date**: 2026-05-24
- **trigger**: V01-ADR-088 により MVP 外へ送った semantic realization coverage 領域について、OpenFastTrace (OFT) が外部実装候補になり得る可能性が判明し、M19 contract 確定前に canonical identity と将来接続境界への影響を確認する必要が生じた
- **scope**: OFT が deferred realization coverage の実装候補となり得るか、およびその可能性が semantic trace MVP foundation / M19 contract に与える設計上の影響
- **non_scope**: OFT 採用決定、V01-ADR-088 の撤回、coverage relation の MVP active 化、`yaml:` / `internal-design:` endpoint の active 化、M19 implementation の変更
- **source_refs**:
  - V01-ADR-081
  - V01-ADR-088
  - V01-INV-DOCS-002
  - V01-INV-DOCS-003
  - spec:trace.artifact-refs
  - spec:trace.coverage-mapping
- **follow_up_candidates**:

> 本 investigation は OFT の採用を決定するものではない。また、V01-ADR-088 が定めた「semantic realization relation を MVP operational scope に含めない」という現行方針を、この時点で変更するものではない。
> 対象は、将来の external coverage tool 候補が現実的に存在する場合に、現在確定しようとしている canonical reference foundation が不要な二重 identity や後戻りを生まないかである。

## 調査スコープ

- V01-ADR-088 により deferred とされた realization coverage 領域のうち、OFT が代替候補となり得る機能範囲を公式一次情報に基づいて確認する。
- Brewprint の `spec:` semantic ref / record ID-as-ref と、OFT が要求する specification item identity / revision model が共存可能かを確認する。
- OFT 利用を将来選択肢として残す場合に、M19 Phase A で確定する resolver / validation contract へ事前に課すべき制約があるかを確認する。
- `docs/requirements/`、`docs/spec/`、implementation / test、将来の YAML / internal-design endpoint のうち、OFT coverage chain に含める候補と含めるべきでない artifact layer を整理する。
- 外部 CLI / CI 利用による小さな compatibility spike で判断材料を得られるかを評価する。

## 非スコープ

- OFT の採用・不採用を本 investigation だけで決定しない。
- `maps_to` / `covers` / `validates` relation を Brewprint MVP に導入しない。
- `coverage:` / `COV-*` / external relation artifact の placement または schema を確定しない。
- `yaml:` / `internal-design:` を active semantic endpoint として再導入しない。
- M19 の implementation、task status、requirement / work item status を変更しない。
- OFT の fork、ライブラリ埋め込み、配布形態、ライセンス判断を確定しない。利用形態が判断対象となる場合のみ後続判断へ渡す。

## 背景

V01-ADR-088 は、semantic trace MVP を canonical reference resolution foundation に限定し、semantic realization relation、external coverage artifact、coverage matrix、`yaml:` active endpoint、および `internal-design:` endpoint を concrete requirement が成立するまで後続判断へ送った。

現行の `spec:trace.coverage-mapping` は、relation を operational に再導入する場合、identity、direction、owner、schema、validation を一緒に判断すると定めている。この境界は、自作 coverage mechanism を先取りしないためには整合的である。

一方で、deferred にした領域に対して外部 tool が有力な候補となり得る場合、将来 relation を導入する時点だけでなく、その前段の canonical reference foundation が tool interoperability を妨げないことも確認対象になる。特に、Brewprint が所有する `spec:` semantic ref と、外部 tool が所有または要求する item identity / revision / relation metadata が別々に存在する場合、将来の adapter または identity synchronization が過剰になる懸念がある。

M19 は coverage relation の導入を non-goal としているが、Phase A で active `spec:` ref と record ID-as-ref の resolver input / output contract、validation diagnostic、duplicate detection、および MCP response contract を確定する。したがって、OFT との接続可能性がこれらの contract に制約を与えるなら、implementation 着手後ではなく Phase A 前または Phase A の判断材料として調査すべきである。

## 調査したもの / 調査対象

### 起票時点で確認済みの Brewprint artifacts

- `AGENTS.md`
- `docs/prompt_chappy.md`
- `docs/doc-policy.md`
- `docs/TASKS.md`
- V01-ADR-081
- V01-ADR-088
- V01-INV-DOCS-002
- V01-INV-DOCS-003
- `docs/spec/concepts/project-artifact-model/index.md`
- `docs/spec/concepts/traceability/index.md`
- `docs/spec/concepts/traceability/artifact-refs.md`
- `docs/spec/concepts/traceability/coverage-mapping.md`
- `docs/tasks/m19-design-records-semantic-trace-support.md`
- `docs/investigations/README.md`

### これから確認する external / experimental sources

- OpenFastTrace official documentation and repository: item identity、artifact type、revision、coverage relation、validation / reporting、CLI / CI 利用境界
- Brewprint の current spec Markdown を入力にした最小の OFT compatibility spike の可否
- 必要な場合のみ、OFT のライセンスおよび利用形態に関する一次情報

## 調査項目ごとの確認計画

### Q1: OFT は Brewprint が deferred にした realization coverage 領域をどこまで代替し得るか

#### 確認対象

- OFT が扱う artifact / item model
- relation vocabulary と coverage validation
- revision または stale relation の扱い
- custom artifact type / format / input source の拡張性

#### 観測済みの Brewprint 側事実

- `spec:trace.coverage-mapping` は、`maps_to` / `covers`、external relation artifact、coverage matrix、relation validation を MVP 外へ送っている。
- V01-ADR-088 は、これらを必要性の否定ではなく concrete requirement 成立時の再判断対象としている。

#### 判断に必要な観点

- OFT が将来候補になり得る範囲を、自作しない方が合理的な engine 機能と、Brewprint が保持すべき artifact ownership / semantic model に分離できるか。
- OFT が coverage validator として適合しても、Brewprint 固有の canonical reference / investigation / requirement boundary を置き換えるものではないことを明確にできるか。

#### 後続判断先

- OFT が coverage engine 候補となる場合: external coverage tool interoperability に関する ADR refinement 候補
- 適合しない場合: V01-ADR-088 の deferred 判断を維持し、build / buy を後続 requirement 成立時に再判断

### Q2: `spec:` semantic ref と OFT item identity / revision は競合せず共存できるか

#### 確認対象

- OFT の item identifier と revision の意味・lifecycle
- Brewprint `spec:` ref の安定性 rule と front matter ownership
- 同一 normative concept に二つの identity が存在する場合の mapping / source of truth

#### 観測済みの Brewprint 側事実

- MVP では `spec:` が唯一の active semantic ref prefix であり、spec document / section の canonical identity を表す。
- `spec:trace.artifact-refs` は `yaml:` の resolve behavior を固定せず、`internal-design:` / `coverage:` の operational contract を deferred にしている。
- M19 は active `spec:` ref の resolve と duplicate detection を実装対象に含める。

#### 候補

- Candidate A: `spec:` を canonical identity とし、OFT ID は外部 coverage 実行用の派生 mapping とする。
- Candidate B: normative spec item に限って OFT identity を正本とし、Brewprint `spec:` は docs navigation / record reference 用の別 identity として明確に分離する。
- Candidate C: 同一 item への二重 identity が過剰となる場合、OFT を Brewprint の normative spec coverage backend としては採用しない。

#### 判断に必要な観点

- identity owner が曖昧にならないか。
- revision 変更と semantic ref 安定性が衝突しないか。
- M19 resolver が外部 ID を直接知る必要があるか、あるいは将来 adapter layer に閉じ込められるか。

#### 後続判断先

- 必要な場合のみ、traceability spec / Design Records MCP contract の refinement ADR または spec update 候補へ渡す。

### Q3: OFT 適合性を確認するために、どの artifact chain を最小 spike とすべきか

#### 確認対象

- `docs/spec/**` 内の normative section に外部 tool marker を置く場合の可読性と ownership
- implementation / unit test 側の marker placement と reviewability
- YAML / internal-design endpoint を active 化せずに coverage feasibility を確認できるか

#### 観測済みの Brewprint 側事実

- V01-ADR-088 は `yaml:` / `internal-design:` / external coverage artifact を MVP active scope に含めない。
- `docs/requirements/` は要求・不足・要望の捕捉 layer であり、現行仕様の正本ではない。

#### 最小 spike 候補

```text
normative spec section
  -> implementation marker
  -> unit test marker
```

この候補では、`docs/requirements/` を coverage endpoint にせず、`yaml:` / `internal-design:` を active endpoint として再導入しない。OFT が Brewprint の spec-first boundary と共存できるかを最小範囲で確認する。

#### 後続判断先

- spike を実施する場合は、M19 implementation と混ぜず、調査用の限定入力・出力・観測事項を本 investigation または派生 investigation に記録する。

### Q4: M19 contract は OFT 調査完了前にどこまで確定してよいか

#### 確認対象

- M19 Phase A の resolver / validation contract のうち、外部 coverage tool と無関係に確定できるもの
- OFT compatibility によって将来変更され得る identity / extension point / diagnostic boundary

#### 観測済みの Brewprint 側事実

- M19 の non-goal は coverage relation の実装であり、V01-ADR-088 の MVP 縮小方針と整合する。
- 一方、M19 Phase A は `spec:` resolver input / output contract と validation diagnostic を確定するため、将来 identity integration を閉じる設計でないかは確認が必要である。

#### 暫定的な切り分け候補

| M19 item | OFT 調査との関係 |
| --- | --- |
| `ADR-*` / `SPEC-*` / `INV-*` record ID-as-ref resolve | 原則として独立に進められる候補 |
| investigation `source_refs` / `follow_up_results` validation | 原則として独立に進められる候補 |
| physical path を canonical ref にしない boundary | 維持すべき候補 |
| `spec:` normative section identity の外部 tool 連携前提 | 調査結果を見て固定度を判断すべき候補 |
| 将来 external relation / external tool ID の extension point | 調査結果を M19 contract に反映する要否を判断すべき候補 |

#### 後続判断先

- M19 task / Design Records MCP spec の refinement が必要かを、調査結果に基づいて判断する。

### Q5: Requirements layer を OFT の coverage chain に含めるべきか

#### 確認対象

- OFT における requirement item の意味
- Brewprint `docs/requirements/` の捕捉 layer と `docs/spec/` の source-of-truth boundary

#### 観測済みの Brewprint 側事実

- V01-ADR-081 は、requirements を「要求・不足・要望・spec gap 候補」の所有 layer とし、現行仕様の source of truth ではないと定める。
- Accepted requirement であっても、それだけでは現行仕様にならず、外部 contract とするには spec への反映が必要である。

#### 判断に必要な観点

- OFT が normative requirement coverage を前提とする場合、Brewprint の pre-spec capture requirement と同一視してよいか。
- OFT chain の起点は `docs/spec/` 内の normative item とすべきか。

#### 後続判断先

- Requirements layer の境界を維持するか、OFT compatibility のために補足説明だけを追加すべきかを判断する。

## 横断的な観測事実

### 1. 今回の trigger は V01-ADR-088 と衝突するものではない

V01-ADR-088 は、concrete requirement のない realization graph を MVP で実装しない判断である。OFT の存在が確認されたとしても、ただちに coverage relation を MVP に戻す理由にはならない。

一方で、将来 realization coverage を導入する場合の実装候補が外部 tool になり得るなら、現在の canonical identity foundation がその選択肢を不必要に妨げないかは別途確認すべき問いになる。この問いは V01-INV-DOCS-002 / V01-INV-DOCS-003 では扱っていない。

### 2. 調査の適切な挿入点は M19 Phase A 前後である

M18 は V01-ADR-088 に基づく方針同期を完了して closed である。M19 は open であり、coverage relation 自体は non-goal としつつ、resolver / validation contract の確定を Phase A に含む。

したがって、本調査は M18 を再開するためではなく、M19 contract refinement が将来の OFT compatibility を不必要に閉じないかを確認する入力として位置付けるのが妥当と考えられる。

### 3. OFT を評価しても Brewprint の所有境界は消えない可能性が高い

少なくとも起票時点で確認できる Brewprint 側の責務は、requirements capture、spec-first ownership、ADR / investigation record、canonical ref resolve / validation である。OFT が realization coverage の一部または大部分を担える場合でも、これらの artifact ownership をそのまま外部 tool に委譲できるかは別途検証を要する。

## 候補比較

| candidate | meaning | benefit | risk / question |
| --- | --- | --- | --- |
| Candidate A: OFT を future coverage backend 候補として前提化する | Brewprint identity / docs ownership を保持し、relation validation engine は OFT 適合性を軸に後続設計する | 将来の自作重複を避けやすい | ID / revision / marker ownership の境界を今確認する必要がある |
| Candidate B: OFT は参考情報に留め、M19 / future coverage を Brewprint 独自前提で進める | 現行設計を最短で進める | 外部依存を増やさない | 後から OFT 相当へ寄せる場合に二重 identity / migration cost が生じ得る |
| Candidate C: OFT を canonical trace model の owner として再設計する | coverage tool との完全整合を優先できる可能性 | 自作 semantic coverage model を縮小できる可能性 | Brewprint の spec-first / investigation / requirement boundary を壊す恐れがあり、現時点で前提化する根拠不足 |

## 起票時点の推奨案

本 investigation の起票時点では、**Candidate A: OFT を future coverage backend 候補として前提化し得るかを検証する**ことが妥当と考えられる。

ただし、これは OFT 採用や OFT 準拠設計を決定する推奨ではない。次の順序で判断材料を作ることを意味する。

1. OFT の公式一次情報から、deferred realization coverage に適合し得る機能と identity / revision model を確認する。
2. Brewprint の `spec:` canonical identity と OFT identity を二重管理せずに扱える候補を整理する。
3. `docs/requirements/` を起点にせず、normative spec から implementation / test へ至る最小 chain で compatibility spike の必要性と実施方法を判断する。
4. M19 Phase A の contract 確定に先立ち、将来 external tool integration を閉じないための制約または extension point が必要かを判断する。
5. 具体的な変更が必要な場合のみ、ADR / spec / M19 task の refinement 候補へ渡す。

## 終結時の結論

本 investigation は、OFT を Brewprint の external coverage backend または canonical identity bridge の採用候補として追跡しない結論で終結する。

V01-ADR-088 により semantic realization coverage は MVP scope 外であり、現時点では evidence matrix、audit、sign-off、cross-layer coverage のような external backend を必要とする concrete requirement は存在しない。Brewprint の中心責務は、外部 trace tool との互換性ではなく、制約された architecture semantics と canonical query / validation boundary を所有することにある。

したがって OFT compatibility spike、M19 contract refinement、OFT interoperability 用の ADR / spec 更新は起票しない。将来、外部 assurance または coverage 管理の具体要求が成立した場合は、本 investigation の candidate を再利用せず、その要求を起点に改めて調査する。

## 起票時点の後続判断候補（終結により採用しない）

- OFT が coverage backend 候補になり得る場合、Brewprint が所有する canonical identity と外部 tool item identity の関係を ADR または traceability spec refinement で明文化するか。
- M19 Phase A を進める前に、`spec:` resolver contract が external item identity との mapping を妨げないことを acceptance / non-goal として追記するか。
- `docs/spec/**` 内の normative section を最小 OFT spike の起点とし、requirements / YAML / internal-design endpoint を早期に active 化しない境界を維持できるか。
- OFT を採用しない場合でも、外部 coverage tool 候補が再度現れたときに評価可能な extension boundary を残すべきか。

## 起票時点の後続 artifact 候補（終結により採用しない）

| category | candidate artifact / action | purpose |
| --- | --- | --- |
| Investigation result | 本 investigation の更新または派生 investigation | OFT official capability review / spike 結果の記録 |
| ADR | external coverage tool interoperability / identity ownership refinement | OFT 等を前提にした boundary が必要と判断された場合のみ起票 |
| Traceability spec | `spec:trace.artifact-refs`, `spec:trace.coverage-mapping`, `spec:trace.resolve-and-validation` | external tool compatibility に関する boundary / trigger / non-goal の反映候補 |
| M19 task / MCP spec | M19 Phase A contract refinement gate | resolver contract が future integration を閉じないことの確認候補 |
| Compatibility spike | minimal spec -> implementation -> unit test coverage experiment | OFT 適合性の限定検証候補 |

## 未確定点

- OFT の公式 item identity / revision / coverage relation model が、Brewprint の deferred mechanisms のどこまでに実際に適合するか。
- `spec:` を canonical identity としたまま、OFT item identity を派生または adapter-owned identity として扱えるか。
- OFT marker を source Markdown に直接埋め込む必要がある場合、`docs/spec/**` の可読性と source-of-truth boundary に許容可能な影響か。
- Minimal spike が YAML / internal-design endpoint の active 化なしで成立するか。
- OFT capability / lifecycle が M19 Phase A の resolver / diagnostic contract に事前変更を要求するか、それとも後続 adapter で閉じ込められるか。
- OFT 利用形態に応じて、CLI / CI 利用境界またはライセンス評価を別判断として扱う必要があるか。

## 起票時点の制約記録

- 本 artifact の `source_refs` は、MVP canonical reference rule に従い、Brewprint 内の record ID-as-ref または active `spec:` ref に限定した。OFT の外部文書は本文の調査対象として扱い、`source_refs` metadata には含めない。
- 本起票では、新規 investigation artifact の作成と investigation 一覧の同期以外に、ADR、spec、task、requirement、work item、implementation の変更は行わない。
- OFT の能力に関する確定的な観測結果は、公式一次情報および必要な compatibility spike を確認した後に本 artifact へ追記する。
