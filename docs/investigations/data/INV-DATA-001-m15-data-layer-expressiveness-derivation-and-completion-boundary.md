# INV-DATA-001: M15 data layer expressiveness の派生 ADR と completion boundary の再整理

- **status**: concluded
- **date**: 2026-05-27
- **trigger**: legacy open record `docs/tasks/m15-data-layer-expressiveness.md` を、ADR-091 / ADR-092 が確定した `requirement -> work item -> task` workflow artifact フローへ移行するための前提整理
- **scope**: M15 が抱える派生 ADR 群の系列特定、M15 close boundary 候補の再定義、後続 requirement / work item / task 起票に渡す境界の整理
- **non_scope**: 新規 `REQ-*` / `WORK-*` / `TASK-*` の起票、ADR の新規起票・修正・status 変更、spec / implementation / fixture の変更、legacy M15 文書の編集、git commit、未反映 spec / implementation 工程の網羅完了確認
- **source_refs**:
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
  - ADR-091
  - ADR-092
- **follow_up_candidates**:
  - M15 Phase A〜B4 の close と `v1.1.0-spec` release を扱う data-layer requirement / work item (`REQ-DATA-*` / `WORK-DATA-*` 候補)
  - ADR-067 / ADR-069 / ADR-070 / ADR-071 / ADR-072 / ADR-073 / ADR-074 / ADR-075 を扱う data-layer model representation extension requirement / work item (`REQ-DATA-*` / `WORK-DATA-*` 候補)
  - ADR-078 / ADR-079 / ADR-080 を扱う MCP semantic identity / state machine identity requirement / work item (`REQ-MCP-*` / `WORK-MCP-*` 候補)

> この investigation は、legacy open record である M15 を新 workflow artifact フローへ移行する前提として、M15 の close boundary と派生 ADR 系列を canonical evidence として記録するために起票された。
> 本文中の boundary は調査結果に基づく候補であり、採用判断は後続 requirement / work item / ADR が所有する。

## 調査スコープ

この investigation は、以下を判断材料として整理する。

- legacy M15 が当初保持していた completion scope の再構成
- M15 作業中に accepted / proposed された ADR 群の系列分類
- M15 minimum close boundary に必須な ADR と、後続 follow-up に分離すべき ADR の境界
- M15 から domain として分離すべき ADR (MCP / state semantic identity) の境界
- M15 を新 workflow artifact (`REQ-*` / `WORK-*` / `TASK-*`) へ移行する際の分割境界候補
- legacy `docs/tasks/m15-data-layer-expressiveness.md` の historical record としての扱い

## 非スコープ

この investigation では以下を行わない / 確定しない。

- 新規 `REQ-*` / `WORK-*` / `TASK-*` artifact の起票
- ADR の新規起票、status 変更、本文修正、Evidence 更新
- spec / implementation / fixture / golden の変更
- legacy `docs/tasks/m15-data-layer-expressiveness.md` の編集
- 個別 ADR の spec 反映状況 / git commit 状態の網羅的確認
- ADR-064 採用案の DAG renderer 実装範囲の確定
- ADR-067 / ADR-073 (proposed) の acceptance 判断
- ADR-078 の `spec/mcp/*` への反映完了確認
- v1.1.0-spec tag 発行条件の最終確定
- git commit

## 背景

ADR-091 は、`requirement -> work item -> task` の workflow artifact 三層を active layer として確定し、従来 `milestone` と呼んでいた実行計画の役割を work item に統合した。
ADR-091 は同時に、既存 `docs/tasks/m*.md` を legacy milestone-shaped work record として残存させ、移行対象として明示的に計画してから扱う方針を採った。
特に M15 は大きな open legacy record であり、機械的に分割しないと記録された。

ADR-092 は、ADR-091 が確定した workflow artifact を Design Records MCP の record / resolver / validation 対象として扱う最小 public contract を確定した。
これにより `REQ-*` / `WORK-*` / `TASK-*` が canonical reference として運用可能になった。

M15 は data layer expressiveness を主題とする legacy milestone-shaped work record である。
M15 起票時点では Phase A (ADR-060 TypeRef) を起点に、Phase B〜B4 を ADR-061〜064 へ分解し、Phase C を「enum / discriminated object / inline struct / 深さ制限」の検討範囲として未分解のまま保持していた。
その後、実作業を通じて Phase C は ADR-067 / ADR-069 / ADR-070 / ADR-071 / ADR-072 / ADR-073 / ADR-074 / ADR-075 の 8 件へ派生した。
さらに ADR-070 (file-private helper model) の MCP exposure 議論を契機として、ADR-078 / ADR-079 / ADR-080 が MCP query layer / state machine semantic identity 領域に派生した。

legacy `docs/tasks/m15-data-layer-expressiveness.md` は `last_updated: 2026-05-09` 以降更新されておらず、派生 ADR 群の status と Phase C / 後続 domain の進捗を反映していない。
本 investigation は、M15 を workflow artifact フローへ移行するために必要な境界判断を、本日時点の証拠から整理する。

## 調査したもの

調査根拠として参照した artifact は以下である。

- legacy `docs/tasks/m15-data-layer-expressiveness.md` 本文 (Phase 構成 / Tag 方針 / Non-goals / Phase 配下 task / Evidence)
- ADR-060 / ADR-061 / ADR-062 / ADR-063 / ADR-064 (M15 Phase A〜B4 主源)
- ADR-067 / ADR-069 / ADR-070 / ADR-071 / ADR-072 / ADR-073 / ADR-074 / ADR-075 (M15 Phase C 派生)
- ADR-078 / ADR-079 / ADR-080 (M15 起点で露出した MCP / state identity domain)
- ADR-091 / ADR-092 (workflow artifact / MCP record boundary)
- `docs/spec/type-ref.md` 存在確認
- `docs/spec/diagnostics.md` Flow validation 節 (Phase A〜B3 由来 diagnostic 反映状況)
- `docs/spec/nodes.md` model section (ADR-067 / ADR-070 / ADR-073 spec 反映の不在確認)
- `docs/spec/views/dag.md` の見出し一覧 (ADR-064 returns.source data line 存在、ADR-071 `## Private models` 不在の確認)
- `docs/spec/views/` directory listing (ADR-072 `model-catalog.md` / ADR-075 `model-file.md` 不在の確認)
- Design Records MCP `list_records` による ADR-060〜080 の status / depends_on 取得

実装の単体テスト存在、ADR 本文 `commit: tbd` と実際の git state の乖離、`spec/mcp/*` への ADR-078 反映状況の網羅確認は、本 investigation の証拠範囲外として扱った。

## 調査項目ごとの確認結果

### Q1: legacy M15 が元々保持していた completion scope は何か

#### 確認対象

- legacy `docs/tasks/m15-data-layer-expressiveness.md`

#### 観測事実

- Phase 構成は Phase A (ADR-060) / Phase B (ADR-061) / Phase B2 (ADR-062) / Phase B3 (ADR-063) / Phase B4 (ADR-064) / Phase C の 6 phase
- Phase A〜B3 は ADR ごとに spec 反映 / 実装 / テストの個別 task が列挙されている
- Phase B4 は ADR-064 proposed 前提で「ADR を accepted へ進める」「UC-001 サンプル整備」「DAG renderer 実装」「golden 更新」が task として残る
- Phase C は「enum / discriminated object / inline struct / 深さ制限」の 4 検討項目しか持たず、ADR 起票 → 議論 → accepted → spec / 実装反映の方針記述のみ
- Tag 方針として「M15 完了時点で `v1.1.0-spec` タグを発行する」と明記
- Non-goals に「UC-002 self-hosting の完了」「subtyping / generics の導入」を明示

#### 候補

- Phase A〜B4 を minimum close、Phase C を後続 requirement へ分離する候補
- Phase A〜C 全完了まで M15 を抱える候補
- Phase A〜B3 のみ minimum close、Phase B4 (DAG render) も後続へ送る候補

#### 判断に必要な観点

- v1.1.0-spec として release するために必要な最小実用範囲
- DAG render が brewprint v1 系の必須 view であることの重み
- Phase C 派生 ADR の規模 (8 件) と M15 完了不能リスク

#### 後続判断先

- M15 close boundary を確定する data-layer requirement / work item

### Q2: ADR-060〜075 のうち、M15 close に必須なものはどれか

#### 確認対象

- ADR-060 / ADR-061 / ADR-062 / ADR-063 / ADR-064 / ADR-067 / ADR-069 / ADR-070 / ADR-071 / ADR-072 / ADR-073 / ADR-074 / ADR-075

#### 観測事実

- ADR-060: accepted、`commit: f507485`、`impl commit: 01e7127`、`spec/type-ref.md` 存在、`spec/diagnostics.md` Flow validation 節に `incompatible_wiring_type` / `invalid_wiring_source` / `invalid_foreach_over_type` / `invalid_type_ref` 反映済み
- ADR-061: accepted、`commit: a5032d1`、`impl commit: 01e7127`、`spec/diagnostics.md` に `invalid_foreach_returns` / `duplicate_flow_source` / `unresolved_wiring_source` 反映済み
- ADR-062: accepted、`commit: a5032d1`、`impl commit: e7b8292`、`spec/diagnostics.md` に `unresolved_return_source` / `invalid_return_source` / `incompatible_return_type` 反映済み、M15 Phase B2 内 spec / 実装 task は `[x]`
- ADR-063: accepted、`commit: ee0a48c`、`impl commit: e7b8292`、`spec/diagnostics.md` で `unresolved_wiring_source` / `duplicate_flow_source` が initialized source 対応に拡張、M15 Phase B3 内 spec / 実装 task は `[x]`
- ADR-064: accepted、`commit: eb891f2`、`impl commit: tbd`、`spec/views/dag.md` に `### returns.source の data line` (L306) / `### initialized source を含むDAG` (L750) が存在
- ADR-067: proposed、`commit: 693e3c0`、`impl commit: tbd`、`spec/nodes.md` model section に kind: enum 記述なし、`spec/diagnostics.md` に `invalid_enum_model` / `duplicate_enum_value` なし
- ADR-069: accepted、`commit: tbd`、`impl commit: tbd`、`spec/diagnostics.md` に `opaque_type_ref` warning なし
- ADR-070: accepted、`commit: 49391ff`、`impl commit: tbd`、`spec/nodes.md` model section は「1ファイル=1定義」のまま、`main: true` for model 記述なし
- ADR-071: accepted、`commit: 476a4f4`、`impl commit: tbd`、`spec/views/dag.md` headings に `## Private models` セクションなし
- ADR-072: accepted、`commit: tbd`、`impl commit: tbd`、`spec/views/model-catalog.md` 不在
- ADR-073: proposed、`commit: tbd`、`impl commit: tbd`、`spec/nodes.md` に kind: tagged_union 記述なし
- ADR-074: proposed、`commit: tbd`、`impl commit: tbd`、ADR 本体は DAG asset node label の TypeRef hint
- ADR-075: proposed、depends_on: ADR-070 / ADR-071 / ADR-072 / ADR-073、`spec/views/model-file.md` 不在
- 各 ADR 本文の「M15 Phase C」「M15 v1.1.0-spec」言及により Phase C lineage が明示

#### 候補

- M15 close 必須: ADR-060 / ADR-061 / ADR-062 / ADR-063 / ADR-064 (Phase A〜B4 主源)
- Phase C lineage だが M15 close blocker から外す: ADR-067 / ADR-069 / ADR-070 / ADR-071 / ADR-072 / ADR-073 / ADR-074 / ADR-075

#### 判断に必要な観点

- v1.1.0-spec release に必須な型表現力の最小実用範囲
- model visibility / helper model render / catalog view を v1.1 release に含める価値と、含めた場合の completion 不能リスク
- proposed ADR を M15 close 条件に含めることの妥当性

#### 後続判断先

- M15 close work item と、data-layer extension follow-up work item

### Q3: ADR-078〜080 は M15 blocker か、別 domain か

#### 確認対象

- ADR-078 / ADR-079 / ADR-080

#### 観測事実

- ADR-078: accepted、`depends_on: ADR-047 / ADR-048 / ADR-049 / ADR-054 / ADR-058 / ADR-070`、主題は MCP synthetic ID を file path から semantic anchor へ寄せる方針
- ADR-078 本文「MCP は Raw YAML AST ではなく ResolvedProject 上の semantic object を LLM へ公開する query layer である」と明示し、主題が MCP query layer の semantic identity policy
- ADR-079: proposed、`depends_on: ADR-035 / ADR-048 / ADR-078`、主題は transition ID の非 file-path 制約
- ADR-080: proposed、`depends_on: ADR-030 / ADR-032 / ADR-035 / ADR-078 / ADR-079`、主題は state machine semantic object 導入と file-path-free scenario reference
- legacy M15 task file に ADR-078 / ADR-079 / ADR-080 への言及は存在しない (M15 起票者本人が M15 直系として扱っていない証拠)
- ADR-078 のトリガーは ADR-070 (Phase C 派生) だが、主軸 dependency は MCP query layer 系 (ADR-047 / ADR-048 / ADR-049 / ADR-054)

#### 候補

- M15 close から完全に分離し、MCP semantic identity / state machine identity の別 requirement として扱う
- M15 close に含める (M15 Phase C 副作用として抱える)

#### 判断に必要な観点

- M15 の主題 (data layer expressiveness for v1.1 TypeRef) と ADR-078〜080 の主題 (MCP query layer / state model semantic identity) の domain 距離
- state machine semantic object 導入 (ADR-080) が data layer expressiveness の boundary を超える事実
- M15 を完了不能にしないための scope discipline

#### 後続判断先

- MCP semantic identity / state machine identity の独立 requirement / work item

### Q4: accepted だが spec / implementation 反映が未完の判断

#### 確認対象

- ADR-064 / ADR-069 / ADR-070 / ADR-071 / ADR-072 / ADR-078

#### 観測事実

- ADR-064 (accepted, `impl commit: tbd`): `spec/views/dag.md` に部分反映あり、DAG renderer 実装と golden 更新が未完
- ADR-069 (accepted, `commit: tbd`): `spec/diagnostics.md` への `opaque_type_ref` warning 反映なし
- ADR-070 (accepted, `commit: 49391ff`, `impl commit: tbd`): `spec/nodes.md` model section への visibility / `main: true` for model 反映なし
- ADR-071 (accepted, `commit: 476a4f4`, `impl commit: tbd`): `spec/views/dag.md` への `## Private models` section 反映なし
- ADR-072 (accepted, `commit: tbd`): `spec/views/model-catalog.md` ファイル不在
- ADR-078 (accepted, `commit: tbd`): `spec/mcp/*` への反映状況は本 investigation の証拠範囲外

#### 候補

- 上記すべてを、それぞれが属する後続 work item の責務として扱う
- ADR-070 / ADR-071 を adopted decision evidence として M15 由来に残す一方、spec / implementation の未反映作業は後続 work が所有

#### 判断に必要な観点

- ADR-091 が定めた「work item は requirement 解消フロー全体を所有」する責務に従い、未反映 spec / implementation を後続 work item へ集約する妥当性
- legacy M15 が adopted decision を historical record として保持する意義

#### 後続判断先

- 各 ADR を所有する後続 work item

### Q5: proposed のまま再判断が必要な ADR

#### 確認対象

- ADR-067 / ADR-073 / ADR-074 / ADR-075 / ADR-079 / ADR-080

#### 観測事実

- ADR-067 (enum model, proposed): UC-002 self-hosting で必要だが acceptance 判断は未実施
- ADR-073 (tagged union model, proposed): 本文に「accepted に進める前に実装コスト見合いを確認」と明記
- ADR-074 (DAG asset type hint, proposed): DAG readability 改善であり data layer 表現力ではない
- ADR-075 (model file render, proposed): `depends_on: ADR-070 / ADR-071 / ADR-072 / ADR-073`、ADR-070 後続 render exposure
- ADR-079 (transition ID 非 file-path, proposed): ADR-078 後続
- ADR-080 (state machine semantic object, proposed): ADR-078 / ADR-079 後続、設計判断の規模が大きい

#### 候補

- すべて後続 work item の acceptance 判断対象とし、M15 close boundary に含めない

#### 判断に必要な観点

- proposed ADR を milestone close 条件に含めることの妥当性
- 各 ADR の影響域と必要性

#### 後続判断先

- data-layer extension work item / MCP semantic identity work item

### Q6: M15 の `v1.1.0-spec` close に必要な最小 completion boundary

#### 確認対象

- legacy M15 Tag 節、ADR-060〜064 の Evidence

#### 観測事実

- M15 Tag 方針: 「M15 完了時点で `v1.1.0-spec` タグを発行する」
- Phase A〜B3 は spec / 実装ともに反映済み (`impl commit: 01e7127` / `e7b8292`)
- Phase B4 は ADR-064 accepted だが renderer 実装 / golden が未完
- Phase C 派生 ADR は accepted 4 件 / proposed 4 件、spec 反映ほぼゼロ、implementation ゼロ

#### 候補

- minimum close boundary を Phase A〜B4 (ADR-060〜064 完遂) に縮小
- Phase A〜B3 のみで close し、Phase B4 (DAG render) も後続へ送る
- 当初計画通り Phase A〜C 全完了で close

#### 判断に必要な観点

- DAG render が brewprint v1 系の必須 view であり、`returns.source` を renderer 側で表現できないと v1.1 release として整合性が崩れる
- Phase C 派生は派生 ADR の規模が大きく、抱え続けると M15 が事実上閉じない

#### 後続判断先

- M15 close work item の completion 条件設計

### Q7: M15 を workflow artifact へ載せる際の分割境界候補

#### 確認対象

- ADR-091 / ADR-092 / 本 investigation の Q1〜Q6 結果

#### 観測事実

- ADR-091 は requirement / work item / task の三層と task 短期粒度方針を確定
- ADR-091 は「M15 は大きな open legacy record であり、機械的に変更・分割しない」と明示
- 派生 ADR は data-layer domain 系列と MCP semantic identity domain 系列に分かれる
- ADR-070 / ADR-071 は accepted かつ adopted decision として記録されているが、spec / implementation 反映は未完

#### 候補

- 3 requirement に分割:
  - M15 Phase A〜B4 close と `v1.1.0-spec` release を扱う data-layer requirement
  - ADR-067 / ADR-069 / ADR-070 / ADR-071 / ADR-072 / ADR-073 / ADR-074 / ADR-075 を扱う data-layer model representation extension requirement
  - ADR-078 / ADR-079 / ADR-080 を扱う MCP / state semantic identity requirement
- data-layer extension requirement をさらに sub-domain (visibility / helper model / catalog / enum / tagged union / lint / DAG type hint) に分割

#### 判断に必要な観点

- 1 requirement あたりの scope と task 粒度
- ADR-070 / ADR-071 を分離すると ADR-072 / ADR-075 との依存関係をどう扱うか

#### 後続判断先

- 各 requirement / work item の起票判断

## 横断的な観測事実

- legacy M15 task file の `last_updated: 2026-05-09` は本日時点 (2026-05-27) から 18 日経過。派生 ADR の accepted 化 / 追加発生 / status 変更は task file に反映されていない。
- Phase A〜B3 の単体テストが impl commit に実際に含まれているかは task file 内 checkbox が `[ ]` のままであり、checkbox の stale 状態と実装証跡の乖離が観測される。test 存在の独立確認は本 investigation の証拠範囲外。
- ADR 本文の `commit: tbd` 表記は、実際の git state と必ずしも一致しない (ADR-070 / ADR-071 は本文上 commit hash 入りだが `impl commit: tbd`)。git status / git log を直接確認できないため、各 ADR の actual git state は本 investigation の証拠範囲外。
- ADR-091 が定めた「task status の正本は task artifact」「work item に checkbox を複製しない」方針に従えば、legacy M15 の checkbox は今後手動更新の対象にしない。

## 後続判断に渡す候補

### M15 close boundary 候補

- minimum close boundary を Phase A〜B4 (ADR-060〜064 完遂) に縮小し、`v1.1.0-spec` tag をその時点で発行する
- ADR-060〜063 は完了 evidence の照合対象として data-layer close work item に含める
- ADR-064 は M15 close に残る implementation / verification 対象として data-layer close work item に含める

### Phase C 派生 ADR の分離候補

- ADR-067 / ADR-069 / ADR-070 / ADR-071 / ADR-072 / ADR-073 / ADR-074 / ADR-075 は M15 由来の data-layer follow-up candidate として記録するが、M15 close blocker には含めない
- ADR-070 / ADR-071 は accepted かつ commit 済みの判断として M15 由来 adopted decision evidence に残す。未反映の spec / implementation 作業は後続 work が所有

### MCP / state identity domain の分離候補

- ADR-078 / ADR-079 / ADR-080 は M15 起点で露出した可能性はあるが、MCP semantic identity / state machine identity 系の別 requirement candidate として分離し M15 close blocker には含めない

### legacy M15 文書の扱い候補

- legacy `docs/tasks/m15-data-layer-expressiveness.md` は historical record として維持
- M15 close work 完了時に migrated / superseded 境界を記録する候補とする
- ADR-091 の方針に従い、新規 task は work item 配下の短期 task として起票し、legacy M15 の checkbox を正本として手動更新しない

## 推奨案

本 investigation の証拠範囲では、以下の workflow artifact 構成が妥当と見られる。

1. **M15 minimum close boundary を Phase A〜B4 へ縮小する** ことが、`v1.1.0-spec` release を現実的に成立させる最小実用範囲として妥当と見られる。
2. **Phase C 派生 ADR 群 (ADR-067 / ADR-069 / ADR-070 / ADR-071 / ADR-072 / ADR-073 / ADR-074 / ADR-075) を data-layer model representation extension の別 requirement へ集約する** ことが、M15 を完了可能な単位に保ちつつ、accepted ADR の未反映作業を canonical な responsible work item に渡す方法として妥当と見られる。
3. **ADR-078 / ADR-079 / ADR-080 を MCP semantic identity / state machine identity の独立 requirement として切り出す** ことが、domain discipline と M15 完了可能性の両方を満たす方法として妥当と見られる。
4. **legacy `docs/tasks/m15-data-layer-expressiveness.md` を historical record として維持し、close 完了時に migrated / superseded 境界を記録する** ことが、ADR-091 の legacy milestone-shaped work record 方針に整合する扱いとして妥当と見られる。

これらは本 investigation の証拠に基づく候補であり、採用判断は後続 requirement / work item / ADR が所有する。

## 後続 artifact 候補

- M15 Phase A〜B4 の close と `v1.1.0-spec` release を扱う data-layer requirement / work item (`REQ-DATA-*` / `WORK-DATA-*` 候補)
- ADR-067 / ADR-069 / ADR-070 / ADR-071 / ADR-072 / ADR-073 / ADR-074 / ADR-075 を扱う data-layer model representation extension requirement / work item (`REQ-DATA-*` / `WORK-DATA-*` 候補)
- ADR-078 / ADR-079 / ADR-080 を扱う MCP semantic identity / state machine identity requirement / work item (`REQ-MCP-*` / `WORK-MCP-*` 候補)

## 未確定点

- Phase A〜B3 の単体テストが impl commit に実際に含まれているかどうか。M15 task file 上 checkbox は `[ ]`、impl commit は hash 入りであり、checkbox の stale 状態が観測される。実装テスト存在の独立確認は data-layer close work item の責務として残る。
- ADR-064 採用案が `spec/views/dag.md` に十分反映されているかどうか。`returns.source の data line` セクションは存在するが、ADR-064 §論点 1〜6 採用案がすべて spec に反映済みかは内容突合が必要。
- ADR-078 の `spec/mcp/*` への反映状況。本 investigation の証拠範囲外。M15 から分離する境界判断には影響しないが、ADR-078 close 判定には別途確認が必要。
- 各 ADR 本文の `commit: tbd` 表記と actual git state の乖離範囲。本 investigation の証拠範囲では確認できない。
- data-layer model representation extension requirement を 1 本にするか、sub-domain (visibility / helper model render / catalog / enum / tagged union / lint / DAG type hint) に分割するかの粒度判断。
- ADR-080 の state machine semantic object 導入が、data-layer domain にも波及するかの確認。
