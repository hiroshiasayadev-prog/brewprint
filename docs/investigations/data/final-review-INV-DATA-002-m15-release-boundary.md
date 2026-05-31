# Final review: M15 v1.1 release boundary after notes retreat inventory

- **status**: final review
- **date**: 2026-05-28
- **trigger**: 前回 review (`review-INV-DATA-002-m15-minimum-expressiveness-boundary.md`) の判断基準に対する重要疑義の検証と、M15 / v1.1.0-spec release boundary の最終推奨確定
- **scope**: 前回 review の判断軸の妥当性検証、M15 historical intent と UC-002 contract dogfooding 価値の評価、`F0` / `F1` / `F2` の焦点比較、ADR status の事実整合の再確認
- **non_scope**: 新規 REQ / WORK / TASK / INV / ADR の起票、spec / implementation / fixture / YAML の変更、`INV-DATA-001` / `INV-DATA-002` / 前回 review report / legacy M15 文書の編集、git commit、inventory 再実施
- **source_refs**:
  - INV-DATA-001
  - INV-DATA-002
  - docs/investigations/data/review-INV-DATA-002-m15-minimum-expressiveness-boundary.md
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

> このレビューは boundary 確定の候補を提示する文書であり、新規 ADR / spec / 実装変更を含まない。
> 採用の最終判断は後続 requirement / work item / ADR が所有する。

---

## 1. Final recommendation

| item | conclusion |
|---|---|
| Recommended release definition | **`minimum-expressiveness release`** |
| Recommended option | **`F1: B4 + enum minimum`**（`F2` は強い follow-up 候補、`F0` は flow-safety release として再定義する場合のみ valid） |
| Should previous O0 recommendation be retained? | **`partially`**（flow-safety release として割り切るなら valid。ただし M15 historical intent / Phase C 派生 ADR 起票者の前提との整合は不十分） |
| Must implement before M15 close | ADR-060〜064 完遂（DAG renderer 実装、UC-001 golden 更新、Phase A〜B3 単体テスト存在の独立確認を含む）+ ADR-069 §10 の M15 確定範囲（`opaque_type_ref` warning の spec / impl 反映）+ ADR-067 §7 の初期移行対象（`mcp_object_type` / `mcp_diagnostic_severity` / `reference_tree_direction` の 3 enum model）導入と UC-002 該当箇所への適用 |
| Explicitly deferred after M15 | ADR-067 §7 の `impact_entry` / `impact_severity` / `impact_fixability` 系を含む UC-002 enum migration の残部、ADR-070 / ADR-071 / ADR-072 / ADR-073 / ADR-074 / ADR-075 系（narrow helper-shape capability は強い follow-up 候補）、ADR-078〜080 系 |
| ADR-078〜080 treatment | **`separate domain`**（前回 review と同じ。MCP semantic identity / state machine identity の独立 requirement に集約） |
| Human decision needed next | **「v1.1 を `flow-safety release` として `F0` で切るか、`minimum-expressiveness release` として `F1` で切るか」** |

推奨の核となる三つの理由:

1. **M15 historical intent は `minimum-expressiveness release` 側にある**。M15 task file タイトルは "data layer expressiveness"、Phase C は当初から構想に含まれており、ADR-067 / ADR-070 は本文の §M15 への影響で「M15 Phase C の…を定める」と明示している。`flow-safety release` として M15 を閉じる場合、release 名称と historical intent との乖離を release notes 等で明示する必要がある。
2. **ADR-067 §7 は ADR 自身が "初期移行対象" として 3 enum model に narrow に scope している**。19 件 enum-like の中で central な direction / object / severity 系の retreat を解消でき、`impact_entry` model 化など重い migration は同 §7 が follow-up へ送ることを明文化している。Critical path の決定論性を著しく下げずに minimum-expressiveness 性質を獲得できる唯一の option である。
3. **ADR-069 §10 が「M15 v1.1.0-spec で確定する内容」を本文明示している**。`opaque_type_ref` warning と parser safety limit の spec 反映は、ADR-067 / ADR-070 の判断とは独立に M15 で扱うべきと ADR 起票者自身が記録している。前回 review はこの evidence を見落としていた。

---

## 2. Review of the previous review's reasoning

| question | assessment | rationale |
|---|---|---|
| Was defining blocker only as "breaks ADR-060〜064 capability" appropriate? | **狭すぎた** | M15 タイトルは "data layer expressiveness"。M15 が解決すべき問題範囲を ADR-060〜064 が導入した capability の有効性に限定したのは循環的な定義であり、M15 そのものの目的を矮小化している。 |
| Did the previous review adequately account for Phase C historical origin? | **不十分** | Phase C は「派生」ではなく M15 当初構想の一部。ADR-067 / ADR-070 本文は「本ADRは M15 Phase C の…を定める」と明示。前回 review は Phase C を pragmatic に follow-up 化したが、historical intent との整合性を独立に評価していない。 |
| Does "any is valid wildcard" suffice to justify O0? | **不十分** | `any` が ADR-060 上 valid wildcard であることは type system level の事実だが、UC-002 MCP contract の中央 model（`change` / `signature` / `impacts` / `summary` / `coverage`）の note には「v1 model では表せない」と書かれている。これは M15 / v1.1 が解消すべき問題そのものであり、`any` の type system 上の正当性を release 判断の根拠にするのは循環。 |
| Is O0 still supportable under a broader dogfooding criterion? | **条件付きで yes** | M15 を `flow-safety release` として明示的に再定義する場合に限り valid。release notes / spec / migration guide で「v1.1 は flow wiring の型安全性を確立する release であり、data-layer 表現力の改善は v1.2 へ送る」ことを明示する必要がある。再定義しない限り、M15 の release 名称が historical intent と乖離する。 |
| Was ADR-064 status represented consistently? | **不整合あり** | 前回 review §6.1 は「ADR-064: `returns.source` の DAG render ルール（proposed → accepted へ進める）」と記載したが、ADR-064 は既に **accepted**（`commit: eb891f2`、`impl commit: tbd`）。INV-DATA-001 Q2 / Q4 の記述とも矛盾していた。M15 close 前に必要なのは acceptance 判断ではなく implementation / verification の完遂である。 |

前回 review が妥当だった点:

- ADR-078〜080 を `separate domain` として分離する判断
- `notes retreat ゼロ化を goal にしない`、`件数だけで判断しない` の判断軸
- ADR-073（proposed、ADR 起票者自身が「M15 から外す judgment を許容」と本文記載）を follow-up に送る判断
- ADR-074 / ADR-075 / ADR-071 / ADR-072 を M15 critical path から外す判断（DAG hint / model file render / catalog view 等は v1.1 必須範囲外）
- ADR-070 narrower boundary の余地を §7 で観測した点（ただし「ADR-070 全体が巨大」と評価したのは過大）

前回 review が狭すぎた点:

- `release_blocker` を ADR-060〜064 capability の破壊に限定したことで、M15 の historical intent が release boundary 判断から事実上排除された
- ADR-064 を proposed と誤認したことで、M15 close path の見積もりが過大化した（accepted ADR の implementation 完遂 vs proposed ADR の acceptance + 実装 + 反映、では critical path の規模が異なる）
- ADR-069 §10 の「M15 v1.1.0-spec で確定する内容」明示を見落とした
- ADR-067 §7 の「初期移行対象」narrow scope を見落とした（ADR-067 全体を critical path に置く必要があると暗黙に仮定した）
- ADR-070 §9 / §10 が MCP exposure / render exposure を既存方針流用 / 後続 ADR 委譲としていることを見落とし、「ADR-070 全体が large blast radius」と評価した

---

## 3. Historical intent of M15

### 3.1 M15 Phase A〜B4 が解決しようとした問題

- ADR-060: TypeRef の導入と flow wiring の型互換性ルールの欠落（`any` の代入互換性が未定義、`list<T>` / `dict<T>` の inline 表現が不可、`foreach` の `$item` 型解決が未明文）
- ADR-061: foreach.returns collected asset を後続 flow から参照する場合の source id 解決ルールの欠落
- ADR-062: task return signature と内部 flow source の接続を明示する手段の欠落
- ADR-063: initialized source を `returns.source` および flow 内部 wiring の bare token source として参照する手段の欠落
- ADR-064: `returns.source` / initialized source 参照の DAG render ルールの欠落

これらはすべて「flow wiring の型安全性 + return source の明示性 + DAG render の整合性」を確立する課題群であり、`flow-safety` 性質に強く倒れている。

### 3.2 M15 Phase C が生えた理由

M15 task file（`docs/tasks/m15-data-layer-expressiveness.md`）は当初から Phase C を「enum / discriminated object / inline struct / 深さ制限」として明記している。Phase C は ADR-060〜064 が壊れて発生したのではなく、`docs/tasks/m15-data-layer-expressiveness.md` Context が次を明示している:

> enum / discriminated object / inline struct など、UC-002 self-hosting で必要になる data layer の表現力が不足

つまり Phase C は「ADR-060〜064 だけでは UC-002 contract を machine-readable に十分表現できない」ことが M15 起票時点で既に予見されていた帰結である。Phase C 派生 ADR の本文がこれを補強している:

- **ADR-067 §M15 への影響**: 「M15 Phase C では、本ADRを受けて enum model の spec / implementation / UC-002 migration を扱う」
- **ADR-070 §M15 への影響**: 「本ADRは M15 Phase C の model visibility / helper model 方針を定める」
- **ADR-069 §10 "M15 v1.1.0-spec への含め方"**: `opaque_type_ref` warning を含む 8 項目を「M15 v1.1.0-spec で確定する」と明示

これらは ADR 起票者自身が M15 Phase C を当然の構想範囲として記述している evidence である。

### 3.3 UC-002 notes retreat の性質

`INV-DATA-002` の inventory および本 review の YAML spot check（`analyze_impact_request.yaml` / `analyze_impact_response.yaml` / `get_signature_response.yaml` / `common.yaml`）の結果、UC-002 MCP contract の中央 model の note には次が共通して書かれている:

- 「v1 model では optional / enum / discriminated object / payload 相関を型として表せないため、change は any + note で保持する」
- 「brewprint v1 model では union / discriminated object を厳密表現できないため any + note で保持する」
- 「brewprint v1 model では enum list を厳密表現できないため any + note で保持する」

これらの note は単なる enhancement demand ではなく、**M15 / v1.1 が解消すべき "v1 model の表現力不足" を YAML 著者が自己診断して記録したもの**である。したがって notes retreat は M15 本来目的の未達を示す evidence であり、release 判断の中核に据えるべきデータである。

### 3.4 v1.1 を `flow-safety release` として切る場合に意図的に残す debt

- enum-like 19 件すべて（direction / object / kind / reference / severity / fallback / detail / error code / status / result）
- helper-shape 14 件すべて（impacts / nodes / edges / tables / objects / coverage / snippet 等）
- tagged-union 5 件すべて（change / signature / members）
- recursive ObjectRef、untagged union list、selector combination matrix、numeric range / default、cross-response behavior、usage-site-dependent vocabulary
- identity series 7 件（FileID / QualifiedID / synthetic ID / file-local ID / semantic object registry / transition / state machine identity）

これらは v1.1 ではすべて prose 依存のまま残る。「v1.1 は flow-safety release である」という再定義が release notes / spec / migration guide で明示されない限り、M15 タイトル "data layer expressiveness" と release 実体が乖離する。

### 3.5 v1.1 を `minimum-expressiveness release` とする場合に最低限解消すべき recurring debt

- 中央 contract の direction / object / severity 系 enum-like 退避（ADR-067 §7 初期移行対象が網羅）
- ADR-069 §10 が M15 確定範囲とした `opaque_type_ref` warning と parser safety limit
- できれば（F2 で扱うべきだが必須ではない）主要 response entry shape の named 化（`impact_entry` / `reference_tree_node_entry` / `endpoint_entry` 等）

これらが解消されると、UC-002 中央 MCP contract の prose 依存が「すべての中央 model が prose」から「主要 enum は machine-readable、helper shape は引き続き prose（follow-up 明示）」へ移行する。release 名称と実体の乖離は許容可能なレベルに収まる。

---

## 4. Critical inventory patterns

| pattern | representative items | count / scale | centrality to UC-002 contract | capability candidate | can defer from M15? | rationale |
|---|---|---:|---|---|---:|---|
| recurring enum-like closed vocabulary（direction / object / kind / severity / reference kind） | N-010 / N-013 / N-017 / N-019 / N-030 / N-031 / N-034 / N-036 / N-037 / N-041 / N-042 / N-045 / N-046 | central 8〜10 / total 19 | **high**（selector / ObjectRef / Reference / Diagnostic は全 MCP tool 横断の中央 model） | ADR-067 §7 初期移行対象（3 enum model） | **partially**（central 部分は M15 で解消、`impact_entry` 等は ADR-067 §7 自身が follow-up へ送る） | central enum を解消すれば UC-002 contract dogfooding の主要 contract が machine-readable になる。ADR-067 §7 が narrow scope を明示している。 |
| central response entry shapes hidden behind `any` | N-005 (`impacts`) / N-014 (`nodes`) / N-015 (`edges`) / N-029 (`tables`) / N-033 (`objects`) | central 5 / total 14 | **high**（MCP tool の主要 response の中身） | ADR-070 §1〜§8 core + 各 helper model の導入 | **conditionally**（ADR-070 既 accepted で spec / impl 反映は可能。ただし F2 を採用する場合のみ M15 critical path に含める。F1 では follow-up へ送る） | helper shape を named 化すると UC-002 contract が大幅に machine-readable になるが、UC-002 YAML 変更が widespread になる。M15 critical path 確定性とのトレードオフが大きい。 |
| discriminator / kind-specific payload hidden behind `any` | N-001 / N-003 (`change`) / N-021 / N-026 (`signature`) / N-027 (`members`) | 5 | **high**（analyze_impact / get_signature / inspect の中央 payload） | ADR-073 tagged union | **yes** | ADR-073 は proposed であり、ADR 起票者自身が「acceptance 前に実装コスト見合いを確認」「実装コストが大きすぎる場合、M15 では `any + note` 継続を選び、tagged union は後続 milestone に送ってよい」と本文明示している。F1 / F2 のいずれでも follow-up へ送ってよい。 |
| unsupported constraints / recursion / untagged union | N-009 (`diagnostic.related` 未タグ union list) / N-011 (`depth` numeric range) / N-020 / N-040 (selector combination matrix) / N-044 (`object_ref.parent` 再帰) / N-051 (`string_list` usage-site-dependent) | 7 | mixed | 既存 ADR では救えない | **yes** | 既存 ADR 候補単独では扱えない長期 follow-up。M15 release boundary に乗せる候補ではない。 |
| MCP / state identity semantics | N-032 (`FileID`) / N-035 (`object_selector.id`) / N-038 (`object_selector.file`) / N-039 (`object_selector.local_id`) / N-043 (`object_ref` identity variants) / N-047 (`resolved_project.semantic_objects`) / N-050 (`SourceLocation.file`) | 7 | high（identity domain） | ADR-078 / ADR-079 / ADR-080 | **yes**（separate domain） | data-layer expressiveness とは domain 距離が大きい。前回 review の `separate_domain` 判定を維持する。 |

---

## 5. Focused option comparison

| option | release definition satisfied | debt reduced | central contract improvement | additional decision needed | blast radius | recommendation | rationale |
|---|---|---|---|---|---|---|---|
| **`F0: B4 only`** | `flow-safety release`（再定義必要） | 0（M15 capability は notes retreat を直接解消しない） | なし（UC-002 中央 MCP contract は全面 prose 依存のまま） | release notes / spec / migration guide で `flow-safety release` への再定義を明示すること（ADR / 公式コミュニケーション） | small | **acceptable**（M15 を flow-safety release として再定義する場合のみ） | flow wiring 型安全性 + return source + DAG render の確立で v1.1 を切る。M15 タイトルと release 実体の乖離を意図的に受け入れる選択肢。 |
| **`F1: B4 + enum minimum`** | `minimum-expressiveness release` | central enum 8〜10 件 + ADR-069 §10 warning + ADR-064 implementation 反映 | 中央 contract の direction / object / severity / reference kind が machine-readable に。selector / ObjectRef / Reference / Diagnostic 横断モデルに効く | ADR-067 の acceptance 判断（proposed → accepted）。ただし ADR 本文が「M15 Phase C で扱う」と前提しており、acceptance を M15 で行うこと自体は ADR 起票者の前提と整合する | medium | **recommended** | minimum-expressiveness release を成立させる現実的最小境界。ADR-067 §7 が "初期移行対象" を 3 enum model に narrow scope しており、`impact_entry` 等は ADR §7 自身が follow-up に明示している。critical path 確定性と historical intent の両立点。 |
| **`F2: B4 + enum minimum + narrow helper-shape capability`** | `minimum-expressiveness release`（より完成形） | central enum 8〜10 件 + central helper shape 5 件 + ADR-069 §10 warning + ADR-064 implementation 反映 | F1 に加えて主要 response entry shape（`impact_entry` / `reference_tree_node_entry` / `endpoint_entry` 等）が named model 化 | ADR-067 acceptance + ADR-070 §1〜§8 core の spec / impl 反映。ADR-070 §9（MCP exposure 既存方針流用）/ §10（render exposure は後続）はそのまま | large | **acceptable**（M15 critical path の確定性を一段下げてでも minimum-expressiveness を完成させたい場合） | UC-002 中央 contract の machine-readability が大幅向上する。ただし helper model migration は UC-002 YAML 変更が widespread になり、Phase A〜B4 完遂 + ADR-067 acceptance との同時並行が必要。F1 完了後の早期 follow-up に送ることで同等効果を得られる。 |
| `F3: B4 + helper-shape only` | partial | central helper 5 件 + ADR-064 / ADR-069 §10 | enum 退避が残ることで selector / ObjectRef / Reference / Diagnostic 横断 model の prose 依存は維持 | ADR-070 §1〜§8 反映 | medium-large | not_recommended | enum は ADR-067 §7 narrow scope のおかげで F1 の adoption コストが小さく、helper だけ入れて enum を残す比較優位がない。 |
| `F4: B4 + tagged union minimum` | partial | 5 件のみ | `analyze_impact.change` / `get_signature_response.signature` / `inspect_response.members` の中央 payload が machine-readable に | ADR-073 acceptance（proposed → accepted） | very_large | not_recommended | ADR-073 起票者自身が「M15 から外す judgment を許容」と本文明示。F1 / F2 と同時採用する必要性も観測できない。 |
| `F5: all Phase C lineage` | maximalist | ほぼ全 debt | UC-002 中央 contract が広範に machine-readable に | ADR-067 / ADR-073 acceptance + ADR-070 / ADR-071 / ADR-072 / ADR-075 の spec / impl 反映 | very_large | reject | M15 を完了不能にするリスクが極大。`INV-DATA-001` の scope discipline 方針と矛盾。 |
| `F6: MCP/state identity included` | other | identity 7 件 | data-layer ではなく MCP / state identity domain | ADR-078 acceptance / ADR-079 / ADR-080 acceptance | very_large | reject | domain 距離が大きく、M15 を data-layer release として切る目的と矛盾。 |

### 5.1 重要観察

- `F1` の "ADR-067 acceptance を critical path に置くリスク" は、前回 review が想定したほど大きくない。ADR-067 本文 §M15 への影響と §7 初期移行対象は ADR 起票者が "M15 Phase C で扱う" 前提で書いている。acceptance 判断は M15 close path 上で行うのが ADR 起票者の前提に整合する。
- `F2` は理想だが、helper model migration は UC-002 YAML が widespread に変わる。M15 close 確定性とのトレードオフが大きいため、F1 を採用後の早期 follow-up に送る方が運用上安全である。
- `F0` は release 名称を `flow-safety release` に再定義する公式コミュニケーションが取れる場合に限り valid である。再定義しない場合、M15 historical intent との乖離が release notes / spec の信頼性を損なう。

---

## 6. Minimum capability boundary, if not F0

### 6.1 Capability included before M15 close（`F1` 採用時）

#### 対象となる機能

- ADR-060〜064 完遂（Phase A〜B4。ADR-064 は既に accepted、impl / DAG renderer / golden 更新が残課題）
- ADR-069 §10 で M15 確定範囲とされた内容:
  - `opaque_type_ref` warning diagnostic の spec / impl 反映
  - parser safety limit の値の spec 反映
  - `unclear_dict_key` / `deep_type_ref` は将来 lint 候補として spec に残置（実装は不要）
- ADR-067 §7 初期移行対象 3 enum model の導入:
  - `mcp_object_type`（values: `node` / `view` / `transition` / `asset` / `field` / `file` / `primitive`）
  - `mcp_diagnostic_severity`（values: `error` / `warning` / `info` / `hint`）
  - `reference_tree_direction`（values: `out` / `in` / `both`）
- 上記 3 enum を使用する UC-002 YAML 該当箇所の `str + note` から enum named model 参照への切替（`object_selector.object` / `object_ref.object` / `diagnostic.severity` / `get_reference_tree_request.direction` / `get_reference_tree_response.direction`）

#### 解消対象となる representative inventory items

- N-010 / N-013 / N-017 / N-019 / N-046（`reference_tree_direction`）
- N-030 / N-036 / N-041（`mcp_object_type`）
- `mcp_diagnostic_severity` 対応の severity 退避（INV-DATA-002 では `diagnostic.severity` 専用行を立てていないが、ADR-067 §7 で central とされている）

ADR-067 §7 自身が follow-up に送ると明示する対象（`impact_entry` model 化に依存する `impact_severity` / `impact_fixability`、object-dependent kind vocabulary、reference_kind の細目）は M15 close では扱わない。

#### 必要な spec surface

- `docs/spec/nodes.md` の model section に `kind: enum` の構文を追加
- `docs/spec/diagnostics.md` に `invalid_enum_model` / `duplicate_enum_value` を追加（ADR-067 §6）
- `docs/spec/diagnostics.md` に `opaque_type_ref` warning を追加（ADR-069 §10）
- `docs/spec/views/dag.md` に ADR-064 採用案の確定反映（既存反映の補完）
- 上記 3 enum model の YAML を UC-002 model 群に追加

#### 必要な implementation / fixture surface

- `internal/semantic` に enum model 表現を追加（既存 model 系列に variant を足す）
- enum 互換性検証ロジック（ADR-067 §4: nominal）
- `opaque_type_ref` warning diagnostic 出力ロジック
- parser safety limit の値の実装
- ADR-064 DAG renderer 採用案の実装と golden 更新
- UC-001 / UC-002 該当箇所の YAML 修正（enum 参照への切替）
- 単体テスト: enum 互換ケース / 非互換ケース / parser safety limit 超過 / `opaque_type_ref` warning ケース

#### 明示的に含めないもの

- `impact_entry` model 化と `impact_severity` / `impact_fixability` 系の enum 導入（ADR-067 §7 自身が follow-up へ送ると明示）
- file-private helper model（ADR-070 §1〜§8 の M15 critical path 投入）
- tagged union model（ADR-073 系）
- private model render（ADR-071）/ model catalog view（ADR-072）/ model file render（ADR-075）/ DAG asset TypeRef hint（ADR-074）
- MCP semantic identity / state machine identity 系（ADR-078〜080）
- selector combination matrix / numeric range / default behavior / recursive ObjectRef / untagged union list / usage-site-dependent vocabulary

### 6.2 Why this is still bounded

#### なぜ Phase C 全体を戻すことにならないか

- ADR-067 §7 が "初期移行対象" を 3 enum model に narrow scope している。Phase C 全体ではなく ADR-067 内の最小 subset のみが critical path に乗る。
- ADR-069 §10 は warning と spec 反映のみで、新規 capability 導入を伴わない。
- ADR-070 / ADR-073 / ADR-074 / ADR-075 系は M15 critical path に含めない。

#### なぜ tagged union / catalog / MCP identity / state machine identity を巻き込まないか

- ADR-073: 起票者自身が「M15 から外す judgment を許容」と本文明示。
- ADR-072 / ADR-075: catalog view と model file render は v1.1 必須範囲外の view / render artifact。
- ADR-074: DAG readability 改善であり data-layer expressiveness ではない。
- ADR-078〜080: MCP query layer / state machine semantic object domain。data-layer から domain 距離が大きい。

#### なぜ notes retreat ゼロ化を要求しないか

- helper shape 14 件、tagged union 5 件、constraint_not_covered 7 件、identity 7 件は M15 close 後も残置する。これらの note retreat 残置は ADR 起票者前提と整合する（ADR-067 §7、ADR-073 §M15 への影響、`INV-DATA-002` §6 capability coverage map "not covered by existing ADRs"）。
- F1 は "minimum-expressiveness release" の最低水準を満たすが、完成形ではない。完成形に近づけるには F2 を follow-up として早期に進める。

### 6.3 Need for a narrower clarification / ADR

- **F1 では narrower clarification ADR は不要**。ADR-067 §7 自身が "初期移行対象" を narrow scope しており、追加 ADR を起票せずに既存 ADR 本文の範囲で M15 close を成立させられる。
- ADR-067 acceptance 判断は M15 close path 上で行う。ADR 本文の §M15 への影響と §7 がこの順序を ADR 起票者前提として記録している。
- **F2 を採用する場合**: ADR-070 §1〜§8 core は既に accepted で本文上 narrow に scope されており、追加 ADR は不要。ただし `impact_entry` 等の具体 helper model 設計は task file または UC-002 側で扱う（ADR-070 §UC-002 への影響と整合）。

#### F0 を推奨する場合の根拠

F0 を採用するには次の二条件のいずれかを満たす必要がある:

1. **v1.1 を明示的に `flow-safety release` と再定義**: release notes / spec / migration guide で「v1.1 は flow wiring の型安全性確立 release。data-layer 表現力の改善は v1.2 へ送る」と明示する。M15 task file タイトル "data layer expressiveness" と release 名称の乖離を意図的に受け入れる。
2. **historical intent からの逸脱を承知の上で release 速度を優先**: ADR-067 / ADR-070 起票者前提を一旦保留し、UC-002 中央 contract が引き続き prose 依存である状態を v1.1 として release する。dogfooding 上の不完全さを follow-up で速やかに解消する確約を別途行う。

これらいずれの根拠も明示されない F0 採用は、historical intent との乖離を release コミュニケーションが吸収しきれない可能性がある。

---

## 7. ADR boundary judgment

| ADR | verified status | relation to inventory | before M15 close? | disposition | rationale |
|---|---|---|---:|---|---|
| ADR-064 | **accepted** | renderer 仕様で notes retreat とは直交 | yes | `required_for_m15_close` | Phase B4 主源。既に accepted。implementation 完遂と spec views/dag.md 反映、UC-001 golden 更新が残課題。 |
| ADR-067 | **proposed** | enum-like 19 件のうち central 8〜10 件 | yes (F1 採用時) | `required_only_if_F1`（または `required_only_if_F2`） | ADR §7 が 3 enum model に narrow scope。acceptance 判断は ADR 起票者前提どおり M15 close path 上で行う。`impact_entry` 系は §7 自身が follow-up に明示。 |
| ADR-069 | **accepted** | opaque container warning 対象 15 件 | yes（F0 / F1 / F2 すべて） | `required_for_m15_close` | ADR §10 が「M15 v1.1.0-spec で確定する内容」を明示しており、`opaque_type_ref` warning と parser safety limit の spec / impl 反映は release boundary 判定によらず M15 で扱う。前回 review はこの evidence を見落としていた。 |
| ADR-070 | **accepted** | helper shape 14 件のうち central 5 件 | partially (F2 採用時のみ) | `required_only_if_F2` | ADR §1〜§8 core は narrow scope。§9 MCP exposure は既存方針流用、§10 render exposure は後続 ADR 委譲。「ADR-070 全体が large」と評価したのは前回 review の過大評価。 |
| ADR-071 | **accepted** | DAG Markdown detail section（render exposure） | no | `follow_up_after_m15` | ADR-070 §10 が後続に委譲した render exposure の具体策。v1.1 必須範囲外。 |
| ADR-072 | **accepted** | catalog view（render artifact） | no | `follow_up_after_m15` | view 追加であり data-layer 表現力ではない。v1.1 必須範囲外。 |
| ADR-073 | **proposed** | tagged union 5 件（central payload） | no | `follow_up_after_m15` | ADR 本文が「M15 から外す judgment を許容」と明示。F1 / F2 採用時いずれでも follow-up。 |
| ADR-074 | **proposed** | DAG hint（readability） | no | `follow_up_after_m15` | DAG readability 改善であり data-layer expressiveness ではない。 |
| ADR-075 | **proposed** | model file render | no | `follow_up_after_m15` | `depends_on: ADR-070 / ADR-071 / ADR-072 / ADR-073` でいずれも follow-up または部分採用。M15 critical path に乗らない。 |
| ADR-078 | **accepted** | MCP semantic identity 7 件 | no | `separate_domain` | data-layer から domain 距離。MCP semantic identity domain として独立 requirement へ。 |
| ADR-079 | **proposed** | state transition identity | no | `separate_domain` | state machine identity domain。 |
| ADR-080 | **proposed** | state machine semantic object | no | `separate_domain` | state machine domain。設計規模が大きく M15 から完全分離。 |

---

## 8. Corrections required before artifact execution

| artifact | correction needed? | correction content | when to apply |
|---|---:|---|---|
| `INV-DATA-001` | partial | 結論「Phase A〜B4 close 推奨」は historical intent の評価と独立に成立する pragmatic 推奨である点を、後続 work item 起票時に補記して明示する（INV 本体の編集ではなく、後続 chain での参照時の注釈として扱う） | 後続 data-layer requirement / work item 起票時 |
| `INV-DATA-002` | yes | `follow_up_candidates: なし` を、本 final review の 3 follow-up domain（data-layer close work item / data-layer extension follow-up / MCP semantic identity follow-up）に修正する。本 review は本体編集を行わないため、修正は後続 work item 起票前の別アクションで行う | 後続 work item 起票前 |
| 前回 review report (`review-INV-DATA-002-m15-minimum-expressiveness-boundary.md`) | yes | (1) ADR-064 status を「accepted」に統一（前回 §6.1 "proposed → accepted へ進める" の誤記訂正）。(2) `release_blocker` 定義に「ADR-060〜064 capability の破壊」に加えて「M15 historical intent が要請する expressiveness 達成の阻害」を含める旨の補記。(3) ADR-069 §10 / ADR-067 §7 / ADR-070 §9 / §10 を見落としていた旨の補記。本 review は前回 review 本体の編集を行わない | 本 review が確定した後、最終 boundary 判断と合わせて前回 review に corrigendum として補記するか、本 review を superseding document として位置づけるかを判断する |
| legacy M15 task record (`docs/tasks/m15-data-layer-expressiveness.md`) | yes | 最終 boundary 決定後、M15 close 時に historical record として次を記録: (1) M15 の実 close boundary（F0 / F1 / F2 のいずれを採用したか）、(2) `flow-safety release` か `minimum-expressiveness release` のいずれとして release したか、(3) ADR-067 / ADR-070 系のうち M15 で扱った範囲と follow-up へ送った範囲。ADR-091 の legacy milestone-shaped work record 方針に整合する形で行う | M15 close work 完了時 |

### 8.1 `INV-DATA-002.follow_up_candidates: なし` の扱い

INV 本体は変更しない。後続 work item / requirement 起票前に、INV へ canonical reference 候補を補記する独立判断を行う。補記すべき 3 follow-up domain は本 review §1 / §6.3 で明示している。

### 8.2 前回 review の ADR-064 status 記述の扱い

前回 review §6.1 「ADR-064: `returns.source` の DAG render ルール（proposed → accepted へ進める）」は事実誤認。ADR-064 は既に accepted（`commit: eb891f2`、`impl commit: tbd`）。本 review が訂正版を提示する。前回 review 本体を編集するか、本 review を superseding document として扱うかは、最終 boundary 判断と同時に決める。

### 8.3 最終 boundary 決定後の M15 historical record

最終 boundary の選択結果（F0 / F1 / F2）と release 名称定義（`flow-safety release` / `minimum-expressiveness release` / その他）は、legacy `docs/tasks/m15-data-layer-expressiveness.md` の Evidence または historical note section に追記し、ADR-091 の legacy milestone-shaped work record 方針に整合する形で残す。F0 を採用する場合は「v1.1 を flow-safety release として明示的に再定義した」旨を release notes / spec / migration guide に追記する公式コミュニケーション計画を、後続 work item で扱う必要がある。

---

## 9. Single human decision

**「v1.1 / M15 を `flow-safety release` として `F0: B4 only` で切るか、`minimum-expressiveness release` として `F1: B4 + enum minimum` で切るか」を確定する**。

この一点が確定すれば、後続は次の通り自動的に従う:

- `F0` 採用の場合: `flow-safety release` への再定義を明示する公式コミュニケーション計画を含む M15 close work item を起票。data-layer extension follow-up は v1.2 以降に送る。
- `F1` 採用の場合: ADR-067 §7 初期移行対象を含む M15 close work item を起票。ADR-067 acceptance 判断は M15 close path 上で行う。`impact_entry` 系 enum 拡大、ADR-070 / ADR-073 系は data-layer extension follow-up に送る。

`F2` は「`F1` を選んだ後、helper-shape を critical path に追加するか」の独立判断として扱う。本 review では `F2` を M15 critical path の単独候補としては推奨せず、「`F1` 完了後の早期 follow-up」として位置づける。`F2` を critical path に直接置くかどうかは、`F1` 採用が確定した後の二次的判断である。

複数 work item の同時開始や、ADR / REQ / WORK / TASK の大量起票はこの decision の責務外である。
