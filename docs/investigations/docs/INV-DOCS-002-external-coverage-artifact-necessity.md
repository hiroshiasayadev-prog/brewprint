# INV-DOCS-002: External coverage artifact necessity for semantic trace MVP

- **status**: concluded
- **date**: 2026-05-24
- **trigger**: M18 で `yaml:` が active endpoint でない MVP に external coverage artifact が必要か再検討する必要が生じた
- **scope**: semantic trace MVP と将来 YAML trace における external coverage artifact の必要性、relation ownership、再導入 trigger の調査
- **non_scope**: ADR / spec / README / example / task / MCP implementation の変更、`yaml:` active 化仕様の確定
- **source_refs**:
  - ADR-081
  - ADR-083
  - ADR-084
  - ADR-085
  - ADR-086
  - ADR-087
- **follow_up_candidates**:
- **follow_up_results**:
  - INV-DOCS-003
  - ADR-088
  - spec:project-artifact-model
  - spec:trace.artifact-refs
  - spec:trace.coverage-mapping
  - spec:trace.resolve-and-validation
  - spec:trace.out-of-scope

> 本 investigation は `docs/coverage/` を残す前提で正当化するものではない。`source_refs` は調査時点で authoritative だった accepted ADR のみに限定し、本文で観測した M18 draft / example は後続同期後の current spec を provenance source として遡及参照しない。
> 以下の推奨は後続判断の候補であり、本 artifact 自体は決定を所有しない。

## 調査スコープ

- ADR-083 で `docs/coverage/` が導入された理由と、ADR-084 および M18 の relation narrowing draft 後の MVP での有効性を確認する。
- MVP の `spec:... --maps_to--> internal-design:...` の owner を external mapping と endpoint metadata の双方から評価する。
- 将来 `yaml:` が active 化された場合にも endpoint metadata で足りる条件と、external artifact が必要になる条件を整理する。
- Candidate A / B / C の影響範囲を ADR / spec / artifact / task / MCP handoff 単位で整理する。

## 非スコープ

- 今回は新規 investigation 以外のファイルを変更しない。
- external coverage artifact の採否を ADR として確定しない。
- YAML ref 粒度、fixture / golden traceability、evidence schema、MCP request / response の具体設計は確定しない。

## 背景

ADR-083 は、spec を正本、brewprint DSL YAML を対象 design model の primary implementation source、internal design を implementation route と位置付け、三者の対応を external coverage artifact が管理するとした。導入時の中心的な問いは「どの design / internal design がどの YAML で cover されるか」「internal design 変更時にどの YAML を処理すべきか」であり、YAML を含む三層対応であった。

ADR-084 は `yaml:` を reserve only とした。その後、M18 の未 commit relation narrowing draft は MVP relation を `maps_to` のみに縮小し、有効な individual mapping を `spec:` → `internal-design:` に限定する案を記述した。調査時点の draft spec、README、example はその案を反映して `docs/coverage/`、`coverage:`、`COV-*` を MVP mechanism として残していた。しかし YAML のない一方向 link に独立台帳が必要かは比較検討されていなかった。

## 調査時点の coverage 導入根拠

ADR-083 の根拠は、spec と二つの realization artifact（internal design と YAML）が独立して存在し、その相互対応と YAML 変更影響を外部から追跡する必要があることだった。

調査時点の MVP 案では YAML endpoint、`covers`、fixture / golden、test evidence、mapping query は operational scope 外であった。残る relation は、internal design がどの spec semantics を具体化するかという単方向対応であった。この対応は internal-design metadata に source spec ref を宣言する形でも保持できるため、ADR-083 の三層根拠は当該 MVP 案にそのまま継承できないと考えられた。

## Relation ownership の現状整理

| relation category | example | meaning | current owner | conflicts with coverage? |
| --- | --- | --- | --- | --- |
| decision provenance / dependency | `depends_on: ADR-084` | 仕様・判断の成立根拠 | ADR / spec / artifact metadata | coverage が由来や判断根拠を表そうとすると衝突する |
| investigation lineage | `source_refs`, `follow_up_results` | 調査の根拠・派生成果物 | investigation metadata | coverage が調査経路や結果記録を持つと衝突する |
| requirement / execution tracking | requirement → work-item → task | 必要性と進捗 | requirements / work-items / tasks | coverage が pending / done / impact を持つと衝突する |
| semantic realization mapping | spec → internal-design → YAML | 同じ design meaning の具体化対応 | 調査時点の draft では coverage | MVP 案の `spec → internal-design` は endpoint metadata でも所有可能であり再検討対象 |
| completeness / evidence coverage | 未実装、test evidence、coverage matrix 等 | 網羅性・証拠・gap 管理 | MVP では owner 未確定 | semantic mapping と混ぜると責務拡張になるが、external artifact の導入根拠にはなりうる |

semantic realization mapping は、存在する endpoint が自己申告すれば resolver で graph を合成できる。一方、未実装対象、意図的非対応、evidence、review sign-off、baseline は endpoint が存在しない状態や横断判定を含むため、endpoint metadata のみでは扱いにくい。

したがって、`docs/coverage/` の必要性を評価する際は、既存 endpoint 間の realization mapping と、completeness / evidence control matrix を分離して考えるべきである。後者が必要になった場合でも、semantic mapping と同じ artifact に収容すべきかは別判断である。

## Candidate A: external coverage artifact を維持する

`docs/coverage/` を semantic realization mapping 専用 artifact とし、MVP でも `COV-*` による `spec:` → `internal-design:` mapping と `coverage:` mapping set identity を保持する。

**得られるもの**:

- relation entry に独立 ID を付けられ、mapping set を一単位としてレビューできる。
- endpoint file を編集せず relation の追加・削除・注記変更を行える。
- 将来の YAML endpoint や many-to-many mapping を既存台帳へ拡張しやすい。

**負担**:

- MVP の情報量に対して directory、schema、identity、validation、MCP acceptance target が増える。
- internal design 本文と external mapping の二箇所を保守する必要がある。
- `coverage` が completeness / evidence / progress まで所有するとの誤読を招きやすい。

**成立条件**:

MVP から mapping entry 自体の承認・監査・個別参照、または横断 matrix / evidence の近接導入が実務要件として確認されること。読んだ一次情報からは、その必須性は確認できなかった。

## Candidate B: MVP では external coverage artifact を外す

MVP の relation は internal-design artifact が metadata で source spec semantic ref を自己申告し、resolver / MCP が逆引き graph を導出する。`docs/coverage/` / `coverage:` / `COV-*` は MVP active trace mechanism から外し、再導入条件が満たされたときに判断し直す。

**得られるもの**:

- relation owner が spec を具体化する endpoint と一致する。
- 一方向 relation のためだけの mapping set、individual ID、coverage YAML、mapping validator を除ける。
- ADR-084 と M18 の relation narrowing draft が YAML と `covers` を外して MVP scope を縮小する方向と整合する。

**失われるもの / 注意点**:

- 独立 mapping ID と mapping set identity は MVP から失われる。
- spec からの逆引きには derived query が必要になる。
- internal design がまだ存在しない spec の gap は endpoint metadata では表せない。

**成立条件**:

MVP を存在する internal design の source spec link 解決に限定し、gap / evidence / sign-off / YAML coverage を operational requirement にしないこと。調査時点の MVP 案はこの条件に一致した。

## Candidate C: endpoint metadata のみで運用する

将来 YAML trace が active 化された後も、internal design と YAML が自ら relation を宣言し、resolver が graph を構築する。存在する endpoint 間の realization relation だけであれば成立しうる。

ただし、未実装 gap、意図的非対応、many-to-many の承認集合、test evidence、reviewer sign-off、release baseline、監査可能な cross-layer matrix が必要になる場合、endpoint metadata のみでは弱い。将来要求が未確認の現時点で、external artifact を恒久的に否定する証拠は不足している。

## Candidate comparison

| criterion | Candidate A: external artifact 維持 | Candidate B: MVP では外す | Candidate C: 将来も導入しない |
| --- | --- | --- | --- |
| MVP `spec → internal-design` の表現力 | 十分 | 十分 | 十分 |
| MVP の artifact / schema / resolver コスト | 高い | 小さい | 小さい |
| relation owner と endpoint の近さ | 外部台帳に分離 | internal design に近接 | 各 endpoint に近接 |
| mapping ID / set identity | 保持 | MVP では持たない | 持たない |
| YAML active 化への拡張 | 台帳を拡張 | 実需時に再判断 | endpoint metadata を拡張 |
| many-to-many relation | 容易 | query で構築可能 | query で構築可能 |
| gap / evidence / sign-off | 拡張すれば保持可能 | 再導入または別 artifact が必要 | endpoint metadata のみでは不向き |
| 調査時点の一次情報による根拠 | 将来想定中心 | MVP scope と整合 | 恒久判断には不足 |

本 investigation が扱った external coverage の問いに限れば Candidate B が最小で責務に合う。Candidate A を維持するなら、mapping の独立 identity や central assurance を MVP で必要とする具体要求が別途必要である。Candidate C は、将来要求が分からない段階では結論を急げない。後続の INV-DOCS-003 と ADR-088 は、Candidate B に残る internal-design endpoint / relation 前提も再評価し、MVP をさらに縮小した。

## ADR / spec / artifact / task への影響

### Candidate B を採用する場合

| category | affected files / handoff | change candidate |
| --- | --- | --- |
| ADR | ADR-083 / ADR-084 | MVP relation owner を refine する後続 ADR が必要。accepted boundary の実質変更であり spec-only fix は不十分 |
| concept spec | `docs/spec/concepts/project-artifact-model/index.md` | MVP の relation owner を internal-design metadata / derived graph に変更する候補 |
| traceability spec | `index.md`, `artifact-refs.md`, `metadata-schema.md`, `coverage-mapping.md`, `resolve-and-validation.md`, `out-of-scope.md` | `coverage:` / `COV-*` / mapping YAML を MVP scope から外し、endpoint-declared relation と再導入 trigger を定義する候補 |
| README | `docs/coverage/README.md`, `docs/internal-design/README.md` | coverage を MVP authoring entrance としない整理と internal-design relation declaration guidance の追加候補 |
| example artifact | `docs/coverage/traceability/semantic-ref.yaml`, `docs/internal-design/resolver/semantic-ref-index.md` | `COV-TRACE-001` を MVP example / acceptance target から外し、internal-design metadata 宣言に置換する候補 |
| task / milestone | `docs/tasks/m18-semantic-traceability-foundation.md`, `docs/tasks/m19-design-records-semantic-trace-support.md`, `docs/TASKS.md` | coverage example と resolver acceptance criteria の差替え候補 |
| MCP implementation handoff | Design Records MCP follow-up | coverage mapping index / validation の前提を外し、endpoint metadata 由来 relation graph の resolver contract を追跡する候補 |

### Candidate A を採用する場合

| category | affected files / handoff | required narrowing candidate |
| --- | --- | --- |
| concept spec | `docs/spec/concepts/project-artifact-model/index.md` | coverage は semantic realization mapping のみを所有し、gap / evidence / progress / provenance を所有しないと明記する |
| traceability spec | `coverage-mapping.md`, `metadata-schema.md`, `resolve-and-validation.md` | mapping が completeness や evidence を示さず、既存 endpoint 間の relation entry に限ると明記する |
| README | `docs/coverage/README.md`, `docs/internal-design/README.md` | coverage と implementation route、work item impact、investigation lineage の非重複を明記する |
| task / MCP handoff | M18 / M19 | mapping query を必須にしない一方で mapping validation を実装する価値と acceptance target を明文化する |

## 推奨案

本 investigation の結論時点では **Candidate B: MVP では external coverage artifact を operational scope から外す** を推奨候補とした。

理由は、ADR-083 の external artifact 導入根拠が YAML を含む三層対応だった一方、ADR-084 と M18 の relation narrowing draft による MVP 案が YAML のない一方向 relation に縮小されていたためである。この relation は internal-design metadata の自己申告で表現でき、独立 mapping set / ID / schema / validator を MVP の必須機構とする実需は一次情報から確認できない。

この推奨自体は Candidate C の採用ではなかった。後続の INV-DOCS-003 / ADR-088 により、MVP の internal-design endpoint / relation も defer された。将来、endpoint self-declaration では扱えない completeness / evidence / central review requirement が生じた場合、external artifact を再判断すべきである。

1. **Verdict**: `remove external coverage artifact from MVP; reconsider later`
2. **Recommended MVP relation owner**: `docs/internal-design/` metadata
3. **External coverage artifact introduction trigger**:
   - 未実装 spec、意図的非対応、gap、approved many-to-many set、cross-layer matrix を中央管理する必要が生じた場合。
   - fixture / golden / test evidence、reviewer sign-off、release baseline、audit snapshot を relation と結び付ける必要が生じた場合。
   - relation entry 自体の独立 identity / lifecycle / approval が実務要件になった場合。
4. **ADR necessity**: 既存 ADR refinement が必要。accepted ADR-083 / ADR-084 が coverage を artifact boundary / MVP active scope として導入しているため、spec-only fix では足りず、後続 ADR で refinement を明示するのが妥当と見られる。
5. **Affected files if adopted**:
   - ADR: ADR-083 / ADR-084 に対する後続 refinement
   - spec: `docs/spec/concepts/project-artifact-model/index.md`, `docs/spec/concepts/traceability/{index,artifact-refs,metadata-schema,coverage-mapping,resolve-and-validation,out-of-scope}.md`
   - README: `docs/coverage/README.md`, `docs/internal-design/README.md`
   - example artifact: `docs/coverage/traceability/semantic-ref.yaml`, `docs/internal-design/resolver/semantic-ref-index.md`
   - task / milestone: `docs/tasks/m18-semantic-traceability-foundation.md`, `docs/tasks/m19-design-records-semantic-trace-support.md`, `docs/TASKS.md`
   - MCP implementation handoff: mapping index / validation から endpoint-declared relation resolution への切替候補

## 後続判断に渡す候補

- MVP relation declaration の metadata field 名、方向、cardinality、validation rule を後続 ADR / spec で判断する。
- `coverage:` / `COV-*` を削除するか、future-reserved concept として残すかを判断する。
- external artifact の再導入条件を YAML active 化そのものではなく、completeness / evidence / sign-off / matrix requirement の発生に置くか判断する。
- external artifact が必要になった場合、semantic mapping と assurance matrix を一つの `coverage` schema に入れるか分離するかを判断する。
- M19 実装着手前に resolver / validation acceptance target の前提を再確定する。

## 未確定点

- internal-design metadata 上の relation field の正式名称と schema。
- spec から internal design への逆引きを MCP public contract に含めるか。
- `coverage:` / `COV-*` を future reservation として保持する migration 価値。
- YAML active 化後に YAML が spec と internal-design の両方を参照するか、transitive relation を許すか。
- completeness / evidence 用の external artifact が必要になった場合、その名称と責務境界。
- 本 investigation の判断完了前に M18 final independent review を継続できる範囲。
