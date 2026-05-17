# 084: semantic trace MVP scope と artifact boundary

- **status**: accepted
- **date**: 2026-05-17
- **depends_on**: ADR-050, ADR-068, ADR-081, ADR-082, ADR-083
- **supersedes**:
- **migrated_to_spec**:

> このADRは起票時点での決定を記録したスナップショットである。
> 現在の仕様は spec を参照すること。

## 背景

ADR-081 により `docs/requirements/` layer と semantic traceability の方針が導入された。
ADR-083 により、brewprint project の artifact boundary と YAML as primary implementation source が整理され、trace layer は physical path ではなく semantic ref を primary key とする方針が定まった。

一方で、semantic ref / trace schema の具体仕様を一度に広く定義すると、requirements、work-items、internal design、coverage、YAML、fixture、golden、test harness、MCP query interface まで巻き込み、MVP として重くなりすぎる。

特に golden fixture / render expected output は、brewprint processor / renderer / validator の debug / regression test asset であり、一般 brewprint project の semantic trace graph と同列に扱うと責務境界が曖昧になる。

また YAML は brewprint project における primary DSL source だが、semantic ref の粒度、YAML entity-level anchor、resolver rule はまだ未定義である。
仕様が固まっていない YAML entity を trace schema の一級対象にすると、後続の YAML schema / naming / resolver 仕様を先取りしてしまう。

そのため本ADRでは、semantic trace MVP で一級対象にする artifact、reserve only とする対象、明示的に scope 外とする対象を定める。
具体的な schema、front matter field、coverage mapping format、resolver validation rule は後続 spec に移す。

## 決定

### 1. semantic trace MVP は spec / internal-design / coverage を一級対象にする

semantic trace MVP で active semantic ref prefix として扱う対象は以下に限定する。

```yaml
active_prefixes:
  - spec
  - internal-design
  - coverage
```

`spec:` は design spec の document / section level semantic ref を表す。
`internal-design:` は internal design topic / wiring story の semantic ref を表す。
`coverage:` は coverage mapping set / mapping group の semantic ref を表す。

MVP は、まず spec semantics、internal design wiring route、coverage mapping の対応関係を安定して扱えることを目的とする。
ADR-083 §7 が coverage に課した問いのうち、本ADRの MVP coverage が扱うのは YAML を含まない問い、すなわち spec と internal design の対応関係に限る。
`spec ↔ YAML`、`internal-design ↔ YAML`、および internal design 変更時にどの YAML を compile / validate / render すべきかという問いは、`yaml:` の active 化までは MVP coverage では答えない。

本ADR MVP は、active prefix を持つ semantic ref が実在 artifact / section / mapping に resolve 可能であることを前提とする。
resolve の具体的挙動、すなわち ref grammar、section anchor lookup、duplicate / orphan detection、解決失敗時の error / warning 方針は後続 spec が定義する。
`spec:` ref の解決は、ADR-081 §7〜§8 の spec front matter `sections` mapping と anchor append-only 規則を前提とする。

MVP coverage edge の source / target に置けるのは、active prefix を持つ semantic ref、すなわち `spec:` / `internal-design:` / `coverage:` のみとする。
requirement / work item の ID-as-ref である `REQ-*` / `WORK-*` は、requirement metadata や work item metadata 内で関連 artifact を参照する用途に使えるが、本MVPの coverage edge の source / target には置かない。

### 2. YAML semantic ref は reserve only とする

`yaml:` prefix は将来拡張のために予約するが、MVP の active prefix には含めない。

```yaml
reserved_prefixes:
  - yaml
```

ここでいう `yaml:` は、対象 system / design model の brewprint DSL YAML、すなわち DAG / UML / state machine / sequence などの design model 表現 YAML を指す semantic ref として予約する。
ADR-083 が primary DSL source と呼ぶ YAML はこの brewprint DSL YAML である。

一方で、spec front matter、coverage mapping YAML、requirement metadata、work item metadata、internal-design metadata などの trace metadata YAML は、`yaml:` prefix の対象ではない。
trace metadata YAML は、個別 artifact の semantic ref または ID-as-ref によって参照され、その schema validation や authoring contract は relation vocabulary ではなく後続 traceability spec / MCP tool contract の責務として扱う。

MVP 時点では、brewprint DSL YAML の entity-level semantic ref、logical unit の粒度、YAML 内 anchor、resolver rule を定義しない。
そのため `yaml:` reserve only は、prefix 名を将来用に確保することを意味する。
MVP schema が `yaml:` ref の記述を許容するか、許容する場合に resolver 未解決を warning / error / ignored のどれにするかは後続 spec で定義する。
YAML と spec / internal design の詳細な traceability は、後続 requirement と work item によって導入する。

### 3. fixture / golden は semantic trace MVP の一級対象にしない

`fixture:` prefix は MVP に含めない。
`fixture:` は reserve only でもなく、project-level semantic trace graph に将来 active 化する予定の prefix としては扱わない。

fixture / golden は brewprint processor / renderer / validator の検証資産であり、一般 brewprint project の source-of-truth chain には含めない。
fixture は固定入力 YAML 群であり、golden / render expected output はその入力から期待される出力である。

新規 fixture では ADR-082 に従い、期待 render 出力を `render_expected/` に置ける。
ただし、render expected output と actual output の比較方式、test harness schema、fixture-local coverage schema は本ADRでは定義しない。

fixture / golden output と docs / spec semantic ref の対応が必要になった場合は、後続 requirement と work item により、project-level trace graph とは別の layer として fixture-level traceability を導入する。
ADR-082 §8 の fixture-local coverage は本ADR MVP の project-level trace graph には含めない。

### 4. relation vocabulary の MVP は `maps_to` / `covers` のみにする

semantic trace MVP の relation vocabulary は以下の 2 つに限定する。

```yaml
relations:
  - maps_to
  - covers
```

`maps_to` は、異なる artifact layer 間で、同じ概念・責務・実装写像に対応することを示す。

`covers` は、ある artifact が対象 semantic ref の意味範囲を表現・包含していることを示す。

`validates` relation は MVP には含めない。
`validates` は fixture、test harness、diagnostic verification、renderer regression test と結びつきやすく、今回の project-level semantic trace schema の範囲を超えるためである。

本ADRで scope 外とする `validates` は、coverage relation vocabulary としての validation relation である。
以下は relation vocabulary の問題ではなく、後続 spec / MCP tool contract の責務として扱う。

- semantic ref の resolve 可否
- duplicate semantic ref の検出
- orphan ref の検出
- coverage mapping の source / target の解決可否
- trace metadata YAML の schema validation
- MCP writer による artifact 生成 contract

これらは validation 軸が異なるため、`validates` relation の不在によって運用不能にはならない。
検証関係を trace graph に含める必要が出た場合は、後続 requirement と work item により relation vocabulary を拡張する。

### 5. physical path は semantic trace の primary key にしない

ADR-083 の trace layer common principle に従い、semantic trace MVP でも physical path を primary key にしない。

physical path、Markdown heading、directory layout、file split / merge は、semantic ref を解決するための implementation detail である。

file rename、document split、document merge、section move が発生しても、semantic ref が同一概念を指す限り trace は維持されるべきである。

### 6. requirement / work item / coverage mapping ID の形式は MVP で予約する

requirement ID は ADR-081 §5 で示された domain-scoped sequence に従い、ADR 番号と結合しない。
work item は ADR-083 に従い、source requirement を必ず持つ。
本ADR MVP では、requirement / work item は `requirement:` / `work-item:` prefix を持つ semantic ref ではなく、`REQ-*` / `WORK-*` の ID-as-ref として扱う。

MVP では以下の ID 形式を採用候補として予約する。

```yaml
id_forms:
  requirement: "REQ-<DOMAIN>-NNN"
  work_item: "WORK-<DOMAIN>-NNN"
  coverage_mapping: "COV-<DOMAIN>-NNN"
```

`requirement:` / `work-item:` prefix は MVP では採用しない。
将来 trace graph 上で requirement / work item を coverage edge の source / target として参照する必要が出た場合、後続 ADR により `requirement:` / `work-item:` prefix の reserve または active 化を判断する。

`coverage:` semantic ref は coverage mapping set / mapping group を指す semantic key である。
`COV-<DOMAIN>-NNN` は個別 coverage mapping の ID 予約である。
coverage mapping ID は coverage relation を個別に参照する必要がある場合のために予約する。
coverage mapping set / group / 個別 mapping の階層関係は後続 spec で確定する。

具体的な required field、status field、validation rule は後続 spec で定義する。

### 7. schema 詳細は後続 spec に移す

本ADRは semantic trace MVP の scope と artifact boundary を決める。
以下の詳細仕様は後続 spec が所有する。

- semantic ref grammar
- allowed character set
- active / reserved prefix vocabulary
- spec front matter の `semantic_refs` / `sections` field
- coverage mapping schema
- resolver rule
- duplicate ref detection
- orphan ref detection
- redirect / superseded mapping
- prefix-ref と ID-as-ref の混在パターン
- field ごとに許可する ref 種別
- `internal-design:` ref の解決単位
- `coverage:` mapping set / group / individual mapping の階層関係
- coverage edge の source / target に置ける semantic ref の組み合わせ
- relation vocabulary の最終命名
- trace metadata YAML schema validation
- MCP writer による artifact 生成 contract
- MCP query interface への露出

後続 spec は `docs/spec/traceability.md` として新設することを第一候補とする。

## 理由

### なぜ MVP scope を絞るか

semantic trace は requirements、work-items、internal design、coverage、YAML、fixture、MCP query interface にまたがる。
最初から全 artifact を一級対象にすると、schema が重くなり、責務境界が曖昧になる。

まず spec / internal-design / coverage に限定することで、project-level の design trace を小さく始められる。
不足が見つかった場合は ADR-081 の requirements layer に捕捉し、work item によって拡張できる。
特に YAML coverage が必要になる self-hosting / UC-002 再構築のタイミング、または brewprint DSL YAML entity ref / resolver rule の spec 化が進んだ時点で、`yaml:` の active 化を後続 ADR / requirement で判断する。

### なぜ YAML を reserve only にするか

YAML は brewprint project における primary DSL source であり、将来的には semantic trace の重要な対象になる。

しかし、YAML entity-level ref、YAML logical unit、YAML 内 anchor、resolver rule はまだ定義されていない。
ここを先に決めると、YAML schema / naming / resolver 仕様を semantic trace 側が先取りしてしまう。

そのため MVP では `yaml:` を予約し、active trace 対象にはしない。
この判断は brewprint DSL YAML の semantic ref を対象とするものであり、front matter や coverage mapping YAML などの trace metadata YAML schema を scope 外にするものではない。

### なぜ fixture / golden を外すか

fixture / golden は brewprint 本体の debug / regression test asset である。
これは「brewprint で記述される project の semantic trace」とは別の関心である。

fixture input と render expected output を比較するだけであれば、簡易 script や test harness で足りる。
それを project-level semantic trace graph に含めると、fixture coverage、validation relation、golden update workflow、render diff semantics まで MVP に混ざってしまう。

そのため fixture / golden は MVP scope 外とする。

### なぜ `validates` を外すか

`validates` は一見有用だが、実際には fixture、test harness、diagnostic verification、renderer regression test など検証資産側の語彙に寄りやすい。

今回の MVP は project-level artifact 間の semantic mapping を扱うものであり、brewprint 本体の検証資産を扱うものではない。

検証関係を trace graph に含める必要が出たら、fixture-level traceability または validation relation を後続 requirement と work item で導入する。
一方で、semantic ref resolver、duplicate / orphan detection、trace metadata YAML schema validation、MCP writer contract は relation vocabulary とは別の後続 spec / tool contract として扱う。

### なぜ `maps_to` / `covers` だけにするか

relation vocabulary は増やしやすいが、増やしすぎると coverage が ontology 設計になり、運用が重くなる。

MVP では、異なる artifact layer 間の対応を表す `maps_to` と、意味範囲の包含を表す `covers` のみに絞る。
実例から不足が出た場合に relation を追加する方が安全である。

## 却下した代替案

### 代替案A: fixture / golden を `fixture:` prefix として MVP に含める

- 利点: fixture と spec / docs の対応を早期に trace できる
- 欠点: brewprint 本体の debug / regression test asset が project-level semantic trace graph に混ざる。test harness、render expected comparison、fixture-local coverage まで MVP に入ってしまう

→ 却下。fixture-level traceability は必要になったら後続 requirement と work item で導入する。

### 代替案B: YAML semantic ref を MVP の active prefix にする

- 利点: spec / internal-design / YAML の対応をすぐに追える
- 欠点: YAML entity-level ref、logical unit、resolver rule が未定義のまま固定される。YAML schema / naming / resolver 仕様を先取りしてしまう

→ 却下。`yaml:` は reserve only とする。

### 代替案C: `validates` relation を MVP に含める

- 利点: fixture / test / diagnostic による検証関係を表現できる
- 欠点: validation relation は検証資産側の責務を MVP に引き込みやすい。今回の project-level semantic mapping の範囲を超える

→ 却下。MVP relation は `maps_to` / `covers` に限定する。

### 代替案D: relation vocabulary を広く定義する

- 利点: 将来の表現力を先に確保できる
- 欠点: 実例に基づかない relation が増え、coverage schema が過剰になる

→ 却下。relation は実例から不足が出たときに追加する。
既存 2 語が複数の意味で運用されていることが coverage review で確認された場合、relation vocabulary の追加または rename を検討する。

### 代替案E: physical path を trace key として使う

- 利点: 実装が簡単
- 欠点: file rename、split、merge、section move で trace が壊れる。ADR-083 の trace layer common principle と衝突する

→ 却下。semantic ref を primary key とする。

## 影響

### docs/spec への影響

`docs/spec/traceability.md` を新設し、semantic ref / trace schema の最小仕様を定義する必要がある。

この spec では、active / reserved prefix、relation vocabulary、front matter schema、coverage mapping schema、resolver rule、validation rule を扱う。
また、trace metadata YAML schema、prefix-ref と ID-as-ref の field-level 分担、coverage edge の source / target 制約、`internal-design:` ref の解決単位、`coverage:` の mapping hierarchy、MCP writer artifact generation contract も後続 spec / tool contract の対象とする。

### docs/doc-policy.md への影響

`docs/doc-policy.md` は ADR-081〜083 と本ADRを踏まえて更新する必要がある。

特に以下を反映する。

- `docs/requirements/`
- `docs/work-items/`
- `docs/internal-design/`
- `docs/coverage/`
- YAML は対象 system / design model の primary DSL source であること
- semantic trace MVP は physical path ではなく semantic ref を primary key とすること
- fixture / golden は MVP の project-level semantic trace schema に含めないこと

### docs/adr-authoring-guide.md への影響

責務表に requirements、work-items、internal-design、coverage を追加する必要がある。

また、UC docs / fixture が gap discovery log や migration state を所有するように読める既存記述は、ADR-082 / ADR-083 / 本ADRと整合するように更新する必要がある。

### Design Records MCP への影響

MVP 直後の必須変更ではない。

ただし将来的に semantic ref を query interface として扱う場合、MCP の外部 contract は physical file schema ではなく semantic query interface とする。
Design Records MCP の next record suggestion / index が stale な場合は、filesystem の実体確認または index rebuild が必要になる。

### fixture / golden への影響

本ADRは fixture / golden の既存運用を変更しない。

ADR-082 に従い、新規 fixture では `render_expected/` を使える。
ただし fixture-level traceability、validation relation、render expected comparison semantics は本ADRの scope 外である。
fixture-local coverage は、必要に応じて fixture 側の補足として扱えるが、本ADR MVP の project-level trace graph には含めない。

## Evidence

- commit: 17d6910
- impl commit: tbd
- 参考: ADR-050 spec-first documentation policy, ADR-068 ADR authoring guide, ADR-081 project requirements layer と semantic traceability, ADR-082 golden fixture と self-hosting requirement の責務境界, ADR-083 project artifact boundary と YAML as primary implementation source
