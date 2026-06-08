# V01-ADR-088: Reduce semantic trace MVP to a canonical reference resolution foundation

- **status**: accepted
- **date**: 2026-05-24
- **depends_on**: V01-ADR-081, V01-ADR-083, V01-ADR-084, V01-ADR-087
- **supersedes**:
- **migrated_to_spec**:

> このADRは起票時点での決定を記録したスナップショットである。
> 現在の仕様は spec を参照すること。
> 本 ADR に基づく spec / README / example / task の同期は M18 で追跡する。

## 背景

V01-ADR-083 は、`docs/spec/` を現行 design spec の正本、brewprint DSL YAML を対象 design model の primary implementation source、`docs/internal-design/` を implementation-facing wiring / route の記録 layer と位置付けた。また、spec / internal design / YAML の対応を管理する external artifact として `docs/coverage/` を導入した。

V01-ADR-084 は semantic trace MVP の active semantic ref prefix に `spec:` / `internal-design:` / `coverage:` を含め、`yaml:` は reserve only とした。その後、M18 では relation vocabulary を `maps_to` に限定し、`spec:` → `internal-design:` の mapping を最小 example とする案が未 commit draft として検討された。

`V01-INV-DOCS-002: External coverage artifact necessity for semantic trace MVP` は、`yaml:` と `covers` を扱わない MVP で external coverage artifact を維持する必然性が弱いことを示し、MVP から `docs/coverage/` / `coverage:` / `COV-*` を外す候補を推奨した。

`V01-INV-DOCS-003: Internal-design endpoint necessity for semantic trace MVP` は、external coverage artifact を外した後も残る `internal-design:` endpoint と `spec` から internal design への semantic realization relation 自体を再評価した。調査の結果、`docs/internal-design/` layer の存在理由は V01-ADR-083 により独立に成立する一方、`internal-design:` を semantic trace MVP の active endpoint とし、`spec:` → `internal-design:` relation を operational に解決・検証する concrete requirement は確認できなかった。

現行の最小 mapping example は、semantic trace resolver の internal design を semantic trace 自身の endpoint として用いる bootstrap example である。これは実装試験には利用できるが、relation を MVP の必須 contract とする独立した実需の根拠にはならない。

本 ADR は、MVP を将来の realization graph の先取りではなく、現時点で必要性が確認された canonical reference resolution / validation foundation に縮小する判断を記録する。

## 決定

### 1. Semantic trace MVP は canonical reference resolution foundation に限定する

Semantic trace MVP の operational scope は、docs artifact を安定して参照し、investigation の根拠参照を解決・検証できる foundation に限定する。

MVP が扱うものは以下とする。

- `spec:` semantic ref の宣言、安定参照、および解決
- `ADR-*` / `SPEC-*` / `INV-*` など、Design Records MCP が扱う record artifact ID-as-ref の解決
- investigation の `source_refs` の canonical reference resolution / validation
- investigation に記載された `follow_up_results` の canonical reference resolution / validation
- investigation の `follow_up_candidates` に artifact reference を記載する場合の canonical form 検査。ただし参照先の存在は要求しない
- physical path を canonical reference としない boundary

MVP は、artifact 間の semantic realization graph、implementation impact graph、coverage matrix、evidence matrix を構築するものとはしない。

### 2. `docs/internal-design/` layer は維持するが、`internal-design:` は MVP active endpoint に含めない

`docs/internal-design/` は、spec に implementation internal を混在させないための long-lived implementation-facing design artifact layer として維持する。

Internal design は引き続き、resolver / index / cache の責務、component boundary、internal phase ordering、lookup route などを記録できる。

ただし semantic trace MVP では、以下を operational scope に含めない。

- `internal-design:` semantic ref prefix の active 化
- internal design document の semantic ref resolve / validation contract
- internal design metadata による source `spec:` relation declaration
- `spec:` から internal design への逆引き graph 導出
- `spec:` → `internal-design:` semantic realization relation の authoring / resolve / validation

Internal design layer の存在と、semantic trace graph の active endpoint としての採用は別判断である。

### 3. External coverage artifact と realization relation を MVP に含めない

Semantic trace MVP の active mechanism として、以下を要求しない。

- `docs/coverage/` に置く semantic realization mapping artifact
- `coverage:` semantic ref prefix
- `COV-*` individual mapping ID
- coverage mapping YAML
- coverage mapping validator / query contract
- `maps_to` / `covers` を用いた semantic realization relation の operational authoring / resolve / validation

`maps_to` は、MVP で external artifact から別 metadata owner へ移設する relation ではなく、semantic realization relation 自体が必要になった時点で、`covers` とあわせて後続判断する対象とする。

### 4. `yaml:` は MVP で active 化しないという判断を維持する

V01-ADR-084 が定めた以下の boundary は維持する。

- brewprint DSL YAML は対象 design model の primary implementation source である
- `yaml:` は MVP の active trace endpoint には含めない
- fixture / golden は project-level semantic trace MVP の一級対象にしない
- `validates` relation は MVP の relation vocabulary としない
- physical path は canonical semantic reference としない

`yaml:` active 化の要否と、その場合に `internal-design:` / realization relation / external artifact をどう導入するかは、具体的な requirement とともに後続判断する。

### 5. Deferred trace mechanisms は concrete requirement が成立した時点で再判断する

以下のいずれかが実務要件として捕捉された場合、後続 ADR / requirement / work item により、`internal-design:` endpoint、semantic realization relation、または external artifact の導入を再判断する。

- 複数の implementation-facing internal design document が存在し、spec からそれらを機械的に辿る navigation / impact analysis が必要になった場合
- investigation / work item / MCP query が internal design artifact を canonical reference として解決・検証する必要を持った場合
- `yaml:` active 化により、spec / internal design / YAML 間の realization chain または cross-layer validation が必要になった場合
- 未実装 spec、意図的非対応、gap、completeness を中央管理する必要が生じた場合
- fixture / golden / test evidence、reviewer sign-off、release baseline、audit snapshot を relation と結び付ける必要が生じた場合
- relation entry 自体に stable identity、approval、lifecycle、履歴管理が必要になった場合

再導入時には、endpoint identity、relation vocabulary、metadata owner、external artifact の有無、resolver / validation / MCP public contract を実在する利用要求に基づいて判断する。

### 6. V01-ADR-083 / V01-ADR-084 を限定的に refine する

本 ADR は以下を限定的に refine する。

- V01-ADR-083 が `docs/coverage/` に与えた MVP semantic realization mapping の役割を外す
- V01-ADR-084 が MVP active prefix に含めた `internal-design:` / `coverage:` を外す
- V01-ADR-084 が MVP に予約した `COV-*` mapping ID および coverage mapping schema の導入前提を外す
- V01-ADR-084 が定めた MVP relation vocabulary の operational 採用を後続判断へ送る

本 ADR は、V01-ADR-083 が定めた `docs/internal-design/` artifact layer 自体、YAML as primary implementation source、requirements / work-items / tasks の責務境界を supersede しない。

## 理由

V01-ADR-083 における external coverage の中心的な導入根拠は、spec / internal design / YAML の三層対応と YAML 変更影響の追跡であった。しかし、V01-ADR-084 後の MVP は `yaml:` endpoint を operational scope に含めず、M18 の relation narrowing draft でも `covers`、fixture / golden、evidence、completeness tracking は MVP 対象にならなかった。

V01-INV-DOCS-002 は、この縮小後の scope では external coverage artifact を維持する実需が確認できないことを示した。V01-INV-DOCS-003 は、さらに残存する `spec:` → `internal-design:` relation についても、現在の一次情報で確認できる根拠は resolver の bootstrap example と、それを前提に書かれた task に留まることを確認した。

一方、V01-ADR-087 によって確定している価値は、investigation record integration と canonical reference resolution / validation である。これは `internal-design:` endpoint や semantic realization relation を MVP に含めなくても成立する。

したがって、MVP は concrete requirement のない realization graph を先取りせず、現在必要な canonical reference resolution foundation に集中する方が、責務境界と実装コストの双方に対して妥当である。

## 却下した代替案

### 代替案A: External coverage artifact のみを外し、internal-design metadata relation を MVP に残す

V01-INV-DOCS-002 Candidate B に相当し、internal design artifact が source `spec:` ref を宣言し、resolver / MCP が逆引き graph を導出する案である。

External artifact の過剰さは解消できるが、relation 自体を MVP 必須とする concrete requirement が確認できない。Bootstrap example を理由に relation metadata / reverse graph / validation を MVP contract とすることになるため採用しない。

### 代替案B: `internal-design:` identity resolve のみを MVP に残し、relation は defer する

Internal design を canonical ref として参照する concrete consumer が既に必要であれば成立しうる。

しかし現時点で確認できる参照は、MVP 自身が作った example / investigation 文脈に依存しており、MVP 必須化の独立根拠としては弱い。Concrete consumer requirement が生じた時点で再判断する。

### 代替案C: `docs/internal-design/` layer 自体を導入しない

Internal design layer 自体は、公開 spec と implementation internal を分離するための V01-ADR-083 の判断として独立に成立している。Semantic trace endpoint の延期を理由に artifact layer まで外す必要はないため採用しない。

### 代替案D: External coverage mapping と `maps_to` relation を MVP に維持する

`yaml:` と `covers` を扱わない状態では、単方向 relation と bootstrap example のために独立 artifact / mapping ID / schema / validator を維持する根拠が不足するため採用しない。

## 影響

### Traceability spec への影響

本 ADR に基づき、`docs/spec/concepts/traceability/` は以下の方向で同期する。

- MVP active semantic prefix は `spec:` を中心とする canonical reference foundation に縮小する
- `internal-design:` / `coverage:` / `COV-*` を MVP active scope から外す
- coverage mapping schema / individual mapping / relation validator の MVP 記述を外す
- `maps_to` / `covers` を semantic realization scope の future decision として整理する
- investigation canonical reference resolution / validation と physical path boundary を MVP の中心に据える
- future trigger と再導入判断対象を out-of-scope / future extension に記録する

### Project artifact model / policy / README への影響

`docs/spec/concepts/project-artifact-model/index.md` は、`docs/internal-design/` layer を維持しつつ、MVP semantic trace endpoint とはしない境界を本 ADR に基づいて反映する。

`docs/doc-policy.md`、`docs/internal-design/README.md`、および必要な navigation / authoring guide は、internal design artifact layer の存在と semantic trace active scope を混同しないよう同期する。External coverage artifact の authoring entrance は MVP では設けず、将来導入を判断した時点で必要な directory / README を新設する。

### Example artifact への影響

`docs/internal-design/resolver/semantic-ref-index.md` は internal design document として残しうるが、`internal-design:` semantic trace endpoint または M18/M19 required relation example としては扱わない。

`docs/coverage/traceability/semantic-ref.yaml` と `COV-TRACE-001` は、MVP example / acceptance target から外す。

### M18 / M19 への影響

M18 は、本 ADR の spec / README / example / task への追従同期後に final independent review を行い、close 可否を判断する。

M19 は、以下を必須範囲から外す方向で再同期する。

- `internal-design:` / `coverage:` / `COV-*` lookup
- individual coverage mapping validation
- `spec:` → `internal-design:` relation resolve test
- `COV-TRACE-001` acceptance test

M19 の必須範囲は、V01-ADR-087 に基づく investigation record integration、record ID-as-ref resolution、`spec:` semantic ref resolution、investigation canonical reference validation、noncanonical physical path diagnostic の concrete contract / implementation / test とする。

## Evidence

- commit: 43bdfdb
- impl commit: 該当なし
- investigation: V01-INV-DOCS-002
- investigation: V01-INV-DOCS-003
- decision context: V01-ADR-083, V01-ADR-084, V01-ADR-087
