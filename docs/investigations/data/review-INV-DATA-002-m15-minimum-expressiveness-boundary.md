# Review of INV-DATA-002: M15 minimum expressiveness boundary

- **status**: review
- **date**: 2026-05-28
- **trigger**: M15 / v1.1.0-spec release boundary を確定するため、`INV-DATA-002` の notes retreat inventory を boundary judgment へ変換するレビュー
- **scope**: `INV-DATA-001` 当初結論の維持可否、M15 close boundary の最終推奨、ADR-067 / ADR-069 / ADR-070 / ADR-071 / ADR-072 / ADR-073 / ADR-074 / ADR-075 / ADR-078 / ADR-079 / ADR-080 の M15 boundary 上の扱い分類、追加の限定 ADR の正当性判断、investigation chain への補記要否
- **non_scope**: 新規 REQ / WORK / TASK / INV の起票、ADR 本文作成 / 修正 / status 変更、spec / implementation / fixture / YAML の変更、`INV-DATA-001` / `INV-DATA-002` 本体の編集、git commit、追加 inventory の再実施
- **source_refs**:
  - INV-DATA-001
  - INV-DATA-002
  - docs/tasks/m15-data-layer-expressiveness.md
  - ADR-060
  - ADR-061
  - ADR-062
  - ADR-063
  - ADR-064
  - ADR-067
  - ADR-069
  - ADR-070
  - ADR-071
  - ADR-072
  - ADR-073
  - ADR-074
  - ADR-075
  - ADR-078
  - ADR-079
  - ADR-080

> このレビューは boundary judgment の候補を整理する文書であり、新規 ADR / spec / 実装変更を含まない。
> 採用の最終判断は後続 requirement / work item / ADR が所有する。

---

## 1. Executive recommendation

| 項目 | 推奨 |
|---|---|
| Recommended option | **O0: B4 only**（ADR-060〜064 完遂のみで M15 close） |
| M15 / v1.1.0-spec close boundary | Phase A〜B4 完遂、Phase C 派生はすべて follow-up |
| M15 close 前に実装すべき capability | TypeRef / flow wiring type validation、foreach.returns、task return source、initialized source wiring、`returns.source` の DAG render（ADR-064 採用案）、UC-001 golden 更新 |
| M15 close から明示的に外す capability | enum model（ADR-067）、opaque TypeRef warning（ADR-069）、file-private helper model（ADR-070）、private model render（ADR-071）、model catalog view（ADR-072）、tagged union model（ADR-073）、DAG asset TypeRef hint（ADR-074）、model file render（ADR-075） |
| ADR-078 / ADR-079 / ADR-080 の扱い | M15 から完全に分離。MCP semantic identity / state machine identity の独立 requirement へ送る |
| `INV-DATA-001` Phase A〜B4 close 結論 | **retain** |

推奨の核は次の三点である。

1. Inventory の 54 件 schema / contract debt の中に、M15 が導入する型安全性・flow 表現を実質的に無効化する `release_blocker` 級の退避は確認できない。大半は `acceptable_note` / `accepted_debt` の閾値内に収まる。
2. ADR-067 / ADR-073 は proposed であり、acceptance 自体が未確定である。これらを M15 close blocker にすることは、acceptance 判断を release boundary の critical path に置き換えることになり、`v1.1.0-spec` の発行可能性を低下させる。
3. ADR-070 は accepted だが、ADR 本文が扱う visibility / 名前解決 / MCP exposure / render exposure / `main: true` for model の範囲は inventory が直接要求する helper shape capability より広く、ADR 全体を M15 boundary に戻すと blast radius が large 化する。Inventory に効くのは helper shape 部分のみという `INV-DATA-002` §6 capability coverage map の観測とも整合する。

---

## 2. Review of inventory quality

| 観点 | 評価 |
|---|---|
| 重大な分類誤り | 観測されない。`primary category` と `primary candidate capability` の組合せは ADR-067〜075 / ADR-078+ の domain 境界と整合している。 |
| 件数集計の読み方 | **件数は重要度ではない**ことを明示する必要がある。`enum_like_closed_vocabulary: 19` は最大カテゴリだが、後述する severity 分類では大半が `acceptable_note` / `accepted_debt` 側に倒れる。「件数最多 = M15 close blocker」と読まれない warning を boundary 判断側で持つ必要がある。 |
| `follow_up_candidates: なし` metadata | 内容と**不整合**。本文 §6〜§7 では ADR-067 / ADR-069 / ADR-070 / ADR-073 / ADR-078+ など複数 follow-up domain が既に観測されている。後続 requirement / work item を起票する前に、INV 補記対象として扱うべき（§8 参照）。 |
| 追加調査なしで boundary 判断へ進めるか | 進める。`INV-DATA-002` §3〜§5 の分類粒度と §6 capability coverage map は boundary 判断に十分な分解能を持つ。 |
| 追加確認が必要な範囲 | 本レビューでは不要。後続 work item 起票時には ADR-064 §論点 1〜6 採用案が `docs/spec/views/dag.md` に反映済みかを別途確認する必要がある（M15 close acceptance のため、`INV-DATA-001` §未確定点でも observed）。これは boundary 確定の前提条件ではなく close work の対象である。 |

inventory 自体の boundary judgment 材料としての適合性は高い。唯一の懸念は `follow_up_candidates: なし` の metadata 不整合だが、これは inventory 結論ではなく後続 chain の起票前に補記すべき軽微な不整合である。

---

## 3. Notes retreat severity model

### 3.1 Severity 定義

| severity | definition | release handling |
|---|---|---|
| `acceptable_note` | 人間向け説明、key semantics の解説、価値の小さい補足。machine-readable contract に必須ではなく、note が残っていても M15 が導入した型安全性・flow 表現が実質的に成立する。 | release 可。follow-up 明示不要。 |
| `accepted_debt` | machine-readable ではないため LLM / tooling が直接 contract を読み取れないが、主要 flow / contract の理解や validate phase の検証を致命的には妨げない退避。後続 capability で解消予定であり、後続 requirement / work item の責務として記録する。 | follow-up を明示して release 可。 |
| `release_blocker` | 主要 contract、または M15 が導入した capability（TypeRef compatibility、flow wiring、foreach.returns、task return source、initialized source wiring、`returns.source` DAG render）を実質的に無効化する退避。`note` で代替できず、M15 内 capability の有効性を破壊するため close 前に解消必要。 | M15 close 前に解消必要。 |

### 3.2 重要な含意

- `release_blocker` は、ADR-060〜064 が確定した capability の有効性を**直接破壊する**退避にのみ適用する。M15 で導入されない data-layer 表現力（enum / tagged union / helper shape など）の不在は、定義上 `release_blocker` ではない。
- M15 が導入する型安全性は「TypeRef compatibility に基づく wiring type validation」であり、`any` は両方向 wildcard として正当に許容される（ADR-060 §3）。したがって `any`-family retreat 24 件は、ADR-060 ルール上は valid な型表現であり、`release_blocker` には該当しない。
- `accepted_debt` の release 許容は、後続 work item での解消を約束することと同義である。`v1.1.0-spec` を release する場合、§6 で示す data-layer extension follow-up と MCP semantic identity follow-up の存在が前提条件となる。

### 3.3 Category ごとの暫定扱い

| inventory category | likely severity | rationale | requires capability before M15 close? |
|---|---|---|---|
| `enum_like_closed_vocabulary` (19) | 大半 `acceptable_note`、`code` / `kind` 等の closed contract vocabulary は `accepted_debt` | 値集合が note に記載されていれば人間は contract を読める。ADR-060 上は `str` として valid。型安全性を破壊しない。 | No |
| `opaque_container_shape` (12) | `accepted_debt` | tool-specific response shape を `any` で保持する退避。`any` は ADR-060 上 wildcard として valid であり、wiring type validation を阻害しない。 | No |
| `tagged_union_candidate` (5) | `accepted_debt` | discriminator payload を `any` に退避するため、LLM 側からの contract 読み取りは困難。ただし M15 が導入する型表現力の外側であり、`any` として ADR-060 上は valid。 | No |
| `named_or_helper_shape_candidate` (2) | `accepted_debt` | local helper model 化で解消できるが、unnamed `any` のままでも flow / contract の主要動作は維持される。 | No |
| `dict_key_semantics` (2) | `accepted_debt` | key 意味論を note に退避。LLM tooling 側の読み取りには不利だが、M15 capability を破壊しない。 | No |
| `identity_or_reference_semantics` (7) | `accepted_debt`（domain 違い） | 別 domain（ADR-078+ MCP semantic identity）で解消すべき退避であり、M15 data-layer の責務外。 | No |
| `constraint_not_covered` (7) | `accepted_debt`（長期）| 既存 ADR 候補だけでは解消できない退避（recursive ObjectRef、untagged union list、selector combination matrix、numeric range / default / unknown-value behavior、cross-response behavior、usage-site-dependent vocabulary）。M15 で扱う対象ではない。 | No |

このレビューの観測範囲では、`release_blocker` に該当する inventory item は存在しない。

---

## 4. Capability option comparison

| option | capability content | notes debt potentially reduced | important debt left unresolved | spec / impl / migration blast radius | M15 close suitability | rationale |
|---|---|---:|---|---|---|---|
| `O0: B4 only` | ADR-060〜064 完遂のみ | 0（M15 capability は notes retreat を直接解消しない設計） | enum 19 / opaque 12 / tagged union 5 / helper 2 / dict 2 / identity 7 / not-covered 7 | small | **recommended** | M15 が導入する capability の本体（TypeRef + flow wiring + foreach.returns + task return source + initialized source + DAG render）は inventory が指す debt と直交する。release 可能性が最も高い。 |
| `O1: B4 + enum` | ADR-067 最小導入 | 最大 19（実際は acceptable_note 級が大半） | opaque 12 / tagged union 5 / helper 2 / dict 2 / identity 7 / not-covered 7 | medium | acceptable | recurring pattern を 1 capability で削減できるが、ADR-067 が proposed のままで acceptance 判断は未実施。boundary 確定の前提として acceptance を critical path に置く必要があり、release 不能リスクが立ち上がる。 |
| `O2: B4 + helper shape` | ADR-070 の helper shape 部分（visibility / render exposure / MCP exposure は narrower boundary） | 約 14（helper shape + nested entry shape） | enum 19 / tagged union 5 / dict 2 / identity 7 / not-covered 7 | large（ADR-070 を絞り込まないと very_large） | not_recommended | ADR-070 全体は visibility / render exposure / catalog 連携を含み blast radius が large。Helper shape 部分だけを切り出すには narrower boundary ADR が必要（§7 参照）。同時に M15 で acceptance 判断を行うのは過剰負担。 |
| `O3: B4 + enum + helper shape` | ADR-067 + ADR-070 narrower | 約 33 | tagged union 5 / dict 2 / identity 7 / not-covered 7 | large | not_recommended | ADR-067 の acceptance 判断と ADR-070 narrower boundary の判断が同時に M15 critical path に乗る。release boundary の決定論性が大きく低下する。 |
| `O4: B4 + tagged union` | ADR-073 追加 | 約 5 | enum 19 / opaque 12 / helper 2 / dict 2 / identity 7 / not-covered 7 | very_large | reject | proposed であり、ADR-073 自身が「accepted へ進める前に実装コスト見合いを確認」と明記している。M15 close boundary に乗せる候補ではない。 |
| `O5: B4 + enum + helper shape + tagged union` | ADR-067 + ADR-070 narrower + ADR-073 | 約 38 | dict 2 / identity 7 / not-covered 7 | very_large | reject | 3 件の Phase C 派生 ADR を同時に M15 critical path に置くと、M15 が事実上閉じない。`INV-DATA-001` の「M15 を完了不能にしないための scope discipline」とも矛盾する。 |
| `O6: limited new boundary` | 既存 ADR とは別に、最小 capability（例: enum minimum-subset） | 不確定 | 不確定 | medium〜large | not_recommended | 新規 ADR を M15 critical path に置くこと自体が release boundary の確定性を低下させる。`INV-DATA-001` の M15 scope 拡大警戒方針とも矛盾する。 |
| `O7: all Phase C lineage` | ADR-067 / 069 / 070 / 071 / 072 / 073 / 074 / 075 を全部 M15 へ戻す | 約 45（estimated） | identity 7 / not-covered 7 | very_large | reject | `INV-DATA-001` が明示的に却下した case。release 不能リスクが極大。 |

### 評価上の留意

- 「notes debt potentially reduced」は capability で対応可能な item 数の上限であり、severity が `acceptable_note` の item を含む。実際の release blocker 削減量はこれより少ない。
- ADR-067 / ADR-073 が proposed である事実は capability suitability に直接影響する。acceptance 判断を M15 close の前段に置くか、後続 follow-up に置くかは独立した判断であり、現状では「後続 follow-up に置く」方が M15 close を確定論的に成立させやすい。

---

## 5. ADR-by-ADR boundary judgment

| ADR | capability | inventory evidence related | include before M15 close? | handling | rationale |
|---|---|---:|---:|---|---|
| ADR-064 | `returns.source` の DAG render ルール | 0（renderer 仕様で notes retreat とは直交） | yes | `required_for_m15_close` | Phase B4 主源。`returns.source` / initialized source の DAG render 確定が v1.1 release の整合性に直結する。`INV-DATA-001` の Phase A〜B4 close 結論を維持する核。 |
| ADR-067 | enum model | 19 件（acceptable_note 多数を含む） | no | `follow_up_after_m15` | proposed のままで acceptance 自体が未確定。19 件のうち release_blocker 級は確認できず、M15 close 前段に置く根拠は弱い。data-layer extension follow-up で acceptance 判断と spec / impl 反映を扱う。 |
| ADR-069 | opaque TypeRef warning | 15 件（warning 対象） | no | `follow_up_after_m15` | warning を出す方針であり、退避を直接解消する capability ではない。M15 close blocker ではない。 |
| ADR-070 | file-private helper model | 14 件（helper shape + nested entry） | no | `needs_narrower_boundary` | accepted だが ADR 全体は visibility / render exposure / MCP exposure / `main: true` for model まで広い。Inventory が要求するのは helper shape capability の一部のみ（`INV-DATA-002` §6 capability coverage map と整合）。data-layer extension follow-up で narrower boundary 判断を行う候補（§7 参照）。 |
| ADR-071 | private model の render exposure | 0（render 範囲、notes retreat とは直交） | no | `follow_up_after_m15` | ADR-070 / ADR-072 後続の render exposure。M15 close blocker ではない。 |
| ADR-072 | model / schema catalog view | 0（view 範囲、notes retreat とは直交） | no | `follow_up_after_m15` | catalog view は v1.1 release の必須範囲外。`docs/spec/views/model-catalog.md` 不在のままでも M15 capability は成立する。 |
| ADR-073 | tagged union model | 5 件 | no | `follow_up_after_m15` | proposed。ADR 本文に「accepted に進める前に実装コスト見合いを確認」と明記。M15 close boundary に乗せる候補ではない。 |
| ADR-074 | DAG asset TypeRef hint | 0（DAG readability、notes retreat とは直交） | no | `follow_up_after_m15` | DAG readability 改善であり data-layer 表現力ではない。`INV-DATA-001` Q5 観測と整合。 |
| ADR-075 | model file render | 0（model file の render 範囲） | no | `follow_up_after_m15` | depends_on: ADR-070 / ADR-071 / ADR-072 / ADR-073。前提 ADR がすべて follow-up であるため M15 close blocker にできない。 |
| ADR-078 | MCP semantic anchor synthetic ID | 7 件（identity series） | no | `separate_domain` | MCP query layer / semantic identity policy。`INV-DATA-001` Q3 の「M15 起点で露出した可能性はあるが、別 domain として分離」結論を inventory も支持する。 |
| ADR-079 | transition ID policy | 0（state transition identity） | no | `separate_domain` | proposed。state machine identity 系の別 requirement に集約する。 |
| ADR-080 | state machine semantic object | 0（state machine semantic object 導入） | no | `separate_domain` | proposed。設計判断の規模が大きく、data-layer domain を越えて state machine 表現に踏み込む。M15 から完全に分離する。 |

---

## 6. Candidate minimum release boundary

### 6.1 Must complete before v1.1.0-spec

#### 必須 capability

- ADR-060: TypeRef + flow wiring type compatibility
- ADR-061: foreach.returns collected asset 参照ルール
- ADR-062: task return source（`returns.source`）の明示化
- ADR-063: initialized source の wiring source 化
- ADR-064: `returns.source` / initialized source 参照の DAG render ルール（proposed → accepted へ進める）

#### 必須 spec reflection

- `docs/spec/type-ref.md` の TypeRef 構文（反映済み）
- `docs/spec/edges.md` の §型互換性ルール、§foreach.over、§wiring source 型解決表、§task return wiring（反映済み）
- `docs/spec/diagnostics.md` の Phase A〜B3 diagnostic（`incompatible_wiring_type` / `invalid_wiring_source` / `invalid_foreach_over_type` / `invalid_type_ref` / `invalid_foreach_returns` / `duplicate_flow_source` / `unresolved_wiring_source` / `unresolved_return_source` / `invalid_return_source` / `incompatible_return_type`）（反映済み）
- `docs/spec/views/dag.md` への ADR-064 採用案反映（M15 close 内で完遂）

#### 必須 implementation / verification

- `internal/semantic` の TypeRef 表現と `internal/resolve` の `validateFlowWiringTypes` 完遂
- foreach.returns collected asset source の semantic 登録と diagnostic 完備
- task return source resolver と TypeRef compatibility 検証完遂
- initialized source wiring source 解決と TypeRef 抑制ルール完遂
- DAG renderer に ADR-064 採用案実装
- UC-001 golden 更新と回帰確認（特に `validate_cart.yaml` で `returns.source: validated_items` を含むケース）

#### 代表 fixture / contract で確認すべきこと

- UC-001 全体が新 wiring type validation で valid 判定される
- ADR-061 / ADR-062 / ADR-063 の `returns.source` / initialized source 参照を含む task が renderer / inspect / MCP `get_signature` で contract どおりに露出する
- Phase A〜B3 の単体テストが impl commit に実際に含まれているかを照合（`INV-DATA-001` 未確定点）

### 6.2 Accepted debt at v1.1.0-spec

#### release 時点で残してよい notes retreat category

| category | items | 残置の正当化 |
|---|---:|---|
| `enum_like_closed_vocabulary` | 19 | M15 capability の有効性を破壊しない。値集合は note に記載されており、人間 contract は維持される。 |
| `opaque_container_shape` | 12 | `any` は ADR-060 上 wildcard として valid。wiring type validation を阻害しない。 |
| `tagged_union_candidate` | 5 | M15 が導入する型表現力の外側。`any` として ADR-060 上 valid。 |
| `named_or_helper_shape_candidate` | 2 | unnamed `any` のままでも flow / contract の主要動作は維持される。 |
| `dict_key_semantics` | 2 | key 意味論は note に退避。M15 capability を破壊しない。 |
| `constraint_not_covered` | 7 | 既存 ADR 候補単独では扱えない（recursive ObjectRef、untagged union list、selector combination matrix、numeric range / default / unknown-value behavior、cross-response behavior、usage-site-dependent vocabulary）。長期 follow-up で扱う。 |

#### 明示的に残す representative debt

- N-001 / N-003 (`analyze_impact.change` の discriminator payload)
- N-005 / N-014 / N-029 / N-033 (`analyze_impact_response.impacts` / `get_reference_tree_response.nodes` / `list_endpoints_response.tables` / `list_objects_response.objects` の nested entry shape)
- N-009 (`diagnostic.related` の untagged union list)
- N-011 (`get_reference_tree_request.depth` の numeric range)
- N-021 / N-026 / N-027 (signature / inspect の kind-specific payload)
- N-044 (`object_ref.parent` の recursive ObjectRef)

#### なぜ blocker ではないか

これらの退避は、M15 が導入する capability（TypeRef + flow wiring + foreach.returns + task return source + initialized source + DAG render）の有効性を破壊しない。`any` は ADR-060 上 wildcard として valid であり、`str + note` の closed vocabulary は ADR-060 上の primitive `str` として valid である。LLM tooling 側からの contract 読み取りは劣化するが、これは M15 capability の本体ではなく、後続 capability で段階的に解消する debt である。

### 6.3 Explicitly separate from M15

#### data-layer follow-up に送るもの

- ADR-067（enum model、proposed）: acceptance 判断 + spec 反映 + implementation
- ADR-069（opaque TypeRef warning、accepted）: spec / impl 反映
- ADR-070（file-private helper model、accepted）: narrower boundary 判断（§7）+ spec 反映 + implementation
- ADR-071（private model render、accepted）: spec / impl 反映
- ADR-072（model catalog view、accepted）: `docs/spec/views/model-catalog.md` 起票 + implementation
- ADR-073（tagged union model、proposed）: acceptance 判断 + 採否確定後の spec / impl
- ADR-074（DAG asset TypeRef hint、proposed）: acceptance 判断 + spec / impl
- ADR-075（model file render、proposed）: depends_on 解消後の判断

#### MCP / state identity 系へ送るもの

- ADR-078（MCP semantic anchor synthetic ID、accepted）: `docs/spec/mcp/*` への反映
- ADR-079（transition ID policy、proposed）: acceptance 判断 + spec / impl
- ADR-080（state machine semantic object、proposed）: acceptance 判断 + 採否確定後の spec / impl / 設計規模に応じた sub-domain 分割

#### M15 と混ぜてはいけない理由

- ADR-078〜080 の主題（MCP query layer の semantic identity、state machine semantic object）は data-layer 表現力（TypeRef + flow wiring + return source + DAG render）と domain 距離が大きい。
- ADR-080 は state machine semantic object 導入により、UC-002 self-hosting / 既存 state DSL の境界に踏み込む。data-layer scope を越える。
- `INV-DATA-001` Q3 / `INV-DATA-002` §4.5 / §6 が「これらは MCP semantic identity / state-machine identity style work に近く、純粋な data-layer expressiveness ではない」と一致して観測している。

---

## 7. Whether a narrower or additional ADR is justified

### 7.1 ADR-070 narrower boundary

ADR-070 全体は visibility / 名前解決 / MCP exposure / render exposure / `main: true` for model を含み、blast radius が large である。一方、`INV-DATA-002` が要求する capability は helper shape（nested response entry array を local helper model として named 化する）の部分のみである。

このため、ADR-070 をそのまま M15 close に含めるのは過大である。Helper shape capability だけを切り出す narrower boundary ADR を追加判断する余地はある。ただし、その判断は data-layer extension follow-up の責務であり、本 review では起票しない。

追加 ADR が解決すべき問いの候補（本 review では本文を作成しない）:

- ADR-070 が扱う helper shape のうち、`local helper model 化された response entry / snippet / coverage / endpoint 形状` を M15 後の最小 capability として切り出すか。
- 切り出した場合、visibility / 名前解決 / MCP exposure / render exposure は別 ADR / 別 work item に分けるか。
- ADR-070 自体を supersede するか、`clarification` として narrower boundary の adoption rule を追記するか。

### 7.2 ADR-073 narrower boundary

ADR-073 は proposed のままで、本文に「accepted に進める前に実装コスト見合いを確認」と明記されている。Inventory に効くのは 5 件（discriminator payload + kind-specific signature / inspect payload）であり、件数は少ないが構造的に重要な退避である。

ただし、ADR-073 narrower boundary（例: `analyze_impact.change` の discriminator pattern だけを minimal tagged union として扱う）は、ADR 本文の規模が大きく、acceptance 判断を M15 close path に置くのは boundary 確定性を下げる。Narrower boundary の判断は data-layer extension follow-up の責務として残す。

### 7.3 既存 ADR を supersede / clarify する必要があるか

- ADR-070 は narrower boundary ADR で `clarify` する余地がある（§7.1）。Supersede までは不要。
- ADR-067 / ADR-073 は proposed のままであり、supersede 対象ではない。Acceptance 判断時に必要なら scope を narrower に調整する。
- ADR-078〜080 については、M15 から分離する判断記録は本 review が果たすため、追加 ADR は不要。

### 7.4 本 review が新 ADR を起票しない理由

- task 指示で ADR / spec / implementation / YAML / legacy M15 文書の変更が禁止されている。
- Boundary 判断の最小 next action（§9）は capability option の確定であり、新 ADR の起票はその後段である。

---

## 8. Required correction to existing investigation chain

### 8.1 `INV-DATA-001` 結論の修正要否

修正不要（**retain**）。

Inventory は当初結論「M15 を Phase A〜B4 / ADR-060〜064 で close し、Phase C 派生を後続へ送る」を**反証していない**。むしろ:

- 19 件の enum-like / 12 件の opaque container / 5 件の tagged union は、いずれも `release_blocker` 級ではなく `accepted_debt` の閾値内に収まる。
- ADR-070 全体が inventory に対して過剰範囲を持つことが §6 capability coverage map で観測されており、ADR-070 を M15 boundary に含めるのは過大という `INV-DATA-001` 推奨案と整合する。
- ADR-078〜080 は別 domain であるという `INV-DATA-001` Q3 結論を、`INV-DATA-002` §4.5 / §6 が同じ方向で確認している。

したがって、後続 work item / requirement 起票時に `INV-DATA-001` の Phase A〜B4 close 結論を直接の根拠として参照できる。

### 8.2 `INV-DATA-002.follow_up_candidates: なし` の修正対象判断

修正対象である（**revise**）。

`INV-DATA-002` 本文 §6〜§7 では複数 follow-up domain が既に観測されている。本 review が candidates を確定したわけではないが、最低限以下の domain は metadata に追記する余地がある:

- M15 Phase A〜B4 の close と `v1.1.0-spec` release を扱う data-layer requirement / work item
- ADR-067 / ADR-069 / ADR-070 / ADR-071 / ADR-072 / ADR-073 / ADR-074 / ADR-075 を扱う data-layer model representation extension requirement / work item
- ADR-078 / ADR-079 / ADR-080 を扱う MCP semantic identity / state machine identity requirement / work item

ただし本 review では `INV-DATA-002` 本体の編集を行わない。INV 補記は後続 work item 起票前の独立判断として扱う。

### 8.3 後続 requirement / work item 起票前に INV に補記すべき結論

本 review の §1 Executive recommendation と §6 minimum release boundary が、後続 chain の起点として参照される。INV 補記が必要な内容は次の通り:

- **`INV-DATA-002` への補記候補**: `follow_up_candidates` を「なし」から、§8.2 の 3 domain 名（または canonical reference 候補）に修正する。
- **`INV-DATA-001` への補記候補**: 本 review が `INV-DATA-001` Phase A〜B4 close 結論を `retain` 判定したことの evidence として、本 review path を「後続 boundary judgment レビューによって裏付けられた」旨で記録する。

これらの補記は、後続 requirement / work item を起票する work が所有する。本 review では補記内容の提示までを行い、実際の編集は行わない。

---

## 9. Smallest next action

**`O0: B4 only` で M15 / v1.1.0-spec release boundary を固定してよいか**、を確定する。

これが yes であれば、後続として:

- M15 Phase A〜B4 close と `v1.1.0-spec` release を扱う data-layer requirement / work item を起票し、ADR-064 acceptance / DAG renderer 実装 / UC-001 golden 更新 / Phase A〜B3 単体テスト存在の独立確認を完遂する。
- ADR-067 / ADR-069 / ADR-070 / ADR-071 / ADR-072 / ADR-073 / ADR-074 / ADR-075 を扱う data-layer model representation extension requirement と、ADR-078 / ADR-079 / ADR-080 を扱う MCP semantic identity / state machine identity requirement を別 work として起票する。

これが no であれば、`O1: B4 + enum` または ADR-070 narrower boundary（§7.1）に進むかの二択を独立判断する必要がある。それ以外の option（O4 / O5 / O6 / O7）は本 review で reject / not_recommended として除外している。

複数の work を同時に開始する判断は本 next action の責務外である。
