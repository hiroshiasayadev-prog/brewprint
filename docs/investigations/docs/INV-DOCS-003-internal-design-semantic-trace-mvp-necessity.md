# INV-DOCS-003: Internal-design endpoint necessity for semantic trace MVP

- **status**: concluded
- **date**: 2026-05-24
- **trigger**: INV-DOCS-002 Candidate B が external coverage artifact を外した後も `internal-design:` endpoint と `spec` から internal design への semantic realization relation を MVP 前提として残しているため、その前提自体を再検討する必要が生じた
- **scope**: M18 semantic traceability foundation における `internal-design:` active endpoint および `spec` と internal design 間の semantic realization relation の MVP 必要性
- **non_scope**: internal design artifact layer 自体の廃止、ADR / spec / README / task / implementation の変更、`yaml:` active 化または external coverage artifact の最終判断
- **source_refs**:
  - ADR-083
  - ADR-084
  - ADR-087
  - INV-DOCS-002
- **follow_up_candidates**:
  - ADR refinement for semantic trace MVP active endpoints and relation ownership
  - spec:project-artifact-model
  - spec:trace.artifact-refs
  - spec:trace.metadata-schema
  - spec:trace.coverage-mapping
  - spec:trace.resolve-and-validation
  - spec:trace.out-of-scope
  - M18 / M19 milestone alignment
  - Design Records MCP resolver acceptance target refinement
- **follow_up_results**:
  - ADR-088

> 本 investigation は、`docs/internal-design/` を不要と結論付けるものではない。対象は semantic trace MVP の active endpoint / operational relation として internal design を今すぐ要求する必要があるかである。
> `source_refs` は調査時点で authoritative だった accepted ADR と prior investigation に限定する。本文で観測した draft spec / task / requirement は、ADR-088 反映後の current artifact を provenance source として遡及参照しない。
> 以下の推奨は後続 ADR / spec / task 判断へ渡す候補であり、本 artifact 自体は決定または現行仕様を所有しない。

## 調査スコープ

- ADR-083 が internal design layer を導入した理由と、semantic trace relation が必要とされた理由を分離して確認する。
- ADR-084、M18 の未 commit relation narrowing draft、および INV-DOCS-002 Candidate B が MVP に残す `internal-design:` endpoint / relation の根拠を確認する。
- 現存する internal design example、M18 / M19 task、MCP requirement が、`spec` と internal design 間の operational semantic relation を MVP 必須とする実需を示しているかを確認する。
- internal design artifact layer を維持しつつ、その semantic trace active 化を deferred にできるかを評価する。

## 非スコープ

- `docs/internal-design/` layer の撤廃または責務再定義を確定しない。
- 後続の corrected ADR decision を起票・accept / reject しない。
- external coverage artifact の採否判断を置き換えない。
- relation metadata field、MCP request / response、diagnostic category、Go implementation を設計・変更しない。
- `yaml:` endpoint、fixture / golden traceability、`covers` relation を active 化しない。

## 背景

ADR-083 は、`docs/spec/` を現行 design spec の正本、brewprint DSL YAML を対象 design model の primary implementation source、`docs/internal-design/` を spec semantics と YAML model を target implementation へ落とす long-lived wiring route と位置付けた。この判断は、公開 contract を持つ spec に implementation component boundary、resolver order、index ownership、phase 分担を混在させないための artifact boundary として成立している。

一方、ADR-083 が external coverage artifact を導入した中心根拠は、spec / internal design / YAML の三層対応と、internal design と YAML の間の変更影響を追えることであった。ADR-084 は `yaml:` を reserve only としながら、semantic trace MVP の active endpoint に `spec:` / `internal-design:` / `coverage:` を残した。その後、M18 の未 commit relation narrowing draft は relation vocabulary を `maps_to` のみに縮小し、operational mapping を `spec:` → `internal-design:` の単方向に限定する案を記述した。

INV-DOCS-002 Candidate B は、YAML のない単方向 relation のためだけに external coverage artifact を維持する必然性は弱いとして、relation declaration を internal-design metadata 側に移す候補を示した。しかし、この候補は「external artifact が不要でも、MVP に `internal-design:` endpoint と semantic realization relation は必要である」という前提を置いている。今回の調査は、その残った前提を一段戻して評価する。

## 調査したもの

- Repository operation / authoring policy: `AGENTS.md`, `docs/prompt_chappy.md`, `docs/doc-policy.md`, `docs/spec-authoring-guide.md`, `docs/investigations/README.md`
- Accepted decisions and uncommitted M18 decision draft: ADR-083, ADR-084, ADR-087, M18 relation narrowing draft
- Prior investigation: INV-DOCS-002
- Concept / traceability drafts: `docs/spec/concepts/project-artifact-model/index.md`, `docs/spec/concepts/traceability/{index,artifact-refs,metadata-schema,coverage-mapping,resolve-and-validation,out-of-scope}.md`
- Internal design artifacts: `docs/internal-design/README.md`, `docs/internal-design/resolver/semantic-ref-index.md`
- Execution / implementation handoff: `docs/tasks/m18-semantic-traceability-foundation.md`, `docs/tasks/m19-design-records-semantic-trace-support.md`, `docs/TASKS.md`, `docs/requirements/mcp/REQ-MCP-001-design-records-semantic-trace-support.md`

## 調査項目ごとの確認結果

### Q1: internal design artifact layer と `internal-design:` active trace endpoint は同じ必要性で正当化されるか

#### 観測事実

- ADR-083 が internal design layer を導入した理由は、公開 contract を持つ spec とは別に、複数 component にまたがる implementation wiring / boundary / resolver order / index ownership を長期的に記録する必要があるためである。
- `docs/internal-design/README.md` も、internal design を「現行 spec semantics と brewprint DSL YAML model を target implementation へ写像する internal wiring / route」の置き場所として定義している。
- これらの理由は、semantic trace graph が MVP で active であるかどうかに依存しない。internal design 文書は、semantic relation を operationalize しなくても authoring / review 用の implementation-facing artifact として存続できる。
- 逆に、`internal-design:` prefix、relation declaration、resolver による逆引き、relation validation は、artifact が存在することに加えて machine-readable に追跡すべき問いと利用者がある場合に初めて必要になる。

#### 評価

Internal design artifact layer の必要性と、semantic trace MVP における `internal-design:` active endpoint の必要性は分離すべきである。ADR-083 は前者を強く根拠化しているが、YAML を scope 外にした MVP で後者を必須にする独立根拠は弱まっている。

### Q2: 現行 MVP に `spec` と internal design の relation を解決・検証しなければ答えられない問いがあるか

#### 観測事実

- ADR-084 は、YAML endpoint がない MVP では coverage が答える問いを「spec と internal design の対応関係」に限定した。
- M18 の未 commit relation narrowing draft は、その関係を `maps_to` による `spec:` → `internal-design:` のみへさらに縮小する案を記述した。
- INV-DOCS-002 Candidate B は、その単方向 relation を internal-design metadata の source `spec:` declaration で保持する案を示している。
- しかし、調査した一次 artifact には、MVP 利用者が「ある spec の internal design 実現先を機械的に列挙できなければならない」「internal design が source spec 宣言を持たなければ M18 が成立しない」という requirement は確認できなかった。
- `REQ-MCP-001` が要求するのは、docs artifact 間の semantic/artifact ref resolve と investigation record integration、特に investigation `source_refs` / `follow_up_results` の validation である。coverage mapping query は明示的に必須範囲外であり、`spec` → internal design relation の query / validation も要求結果として独立に列挙されていない。
- ADR-087 は、resolver が読む lookup source と Design Records MCP が公開する record kind は同一集合である必要がないとした。これは resolver 基盤を作ることが、すべての artifact relation を MVP active にすることを意味しないことを示す。

#### 評価

現在確認できる MVP の hard requirement は、canonical ref resolve と investigation validation であり、`spec` と internal design の operational semantic relation ではない。relation は将来価値を持ちうるが、現時点の M18 / M19 必須成立条件としては根拠が不足している。

### Q3: 現在の internal design example は active relation の必要性を立証しているか

#### 観測事実

- `docs/internal-design/` 配下で確認できた Markdown は README と `resolver/semantic-ref-index.md` のみであり、semantic ref を持つ本文 artifact は resolver example 一件である。
- `internal-design:resolver.semantic-ref-index` は、Design Records MCP の semantic/artifact ref resolver 自身の internal design を説明する文書である。
- 同文書は `spec:trace.semantic-ref` からの最小 mapping target とされ、現行では `COV-TRACE-001` が relation を所有すると記述している。
- M18 Phase E は、traceability spec 自身を最小 example として internal-design ref および `COV-TRACE-001` を追加した。M19 acceptance もこの example の resolve test を予定している。

#### 評価

現行 example は、semantic trace mechanism が自分自身の resolver design を trace する bootstrap / self-example である。これは schema と resolver の実装試験には使えるが、実プロジェクト上の spec-to-implementation navigation requirement を独立に示すものではない。

すなわち、現在は「relation が必要だから example がある」というより、「relation を MVP に置いたため、その relation を検証する example が作られた」という循環を含む。bootstrap example だけを根拠に active endpoint と relation を MVP 必須とするのは弱い。

### Q4: INV-DOCS-002 Candidate B を判断化する場合でも、internal-design declaration を残す必要があるか

#### 観測事実

- INV-DOCS-002 Candidate B は external coverage artifact を MVP active mechanism から外す案であり、その代替として internal-design metadata の source spec declaration と derived reverse graph を提示する。
- Candidate B の主たる理由は、external artifact、mapping set identity、mapping ID、専用 validator を一方向 relation のためだけに導入する負担を外すことにある。
- 同じ縮小原理を一段前提に戻して適用すると、関係の consumer requirement がまだ確認できない場合、relation metadata field、`internal-design:` resolver lookup、reverse graph 導出、relation validation も MVP の必須 contract にしない選択肢が成立する。
- external coverage artifact を外すことと、internal design artifact layer を消すことは別である。同様に、relation metadata を MVP から defer することと、internal design authoring を否定することも別である。

#### 評価

INV-DOCS-002 Candidate B は external artifact を外す点では整合的だが、internal-design relation declaration を残す点については追加の正当化が必要である。MVP を canonical ref / investigation validation foundation にさらに縮小するなら、後続 ADR は relation owner の移設ではなく、semantic realization relation 自体を後続 scope へ送る方向で判断すべきである。

### Q5: defer した場合に失われるものと、再導入 trigger は何か

#### 観測事実

- `internal-design:` を MVP active endpoint から外しても、internal design document は通常の docs artifact として作成・レビューできる。
- ADR-087 に基づく investigation record integration と canonical artifact/semantic ref resolution は、少なくとも ADR / spec / investigation refs を対象として進められる。
- ただし internal design を source ref として canonical に参照したい investigation が MVP 中に必要になる場合、`internal-design:` resolve 対応は再度必要になる。
- YAML trace を active 化し、spec semantics、internal wiring、DSL implementation source の対応を機械的に辿る場合、internal design endpoint の価値は再び強くなる。

#### 再導入 trigger 候補

- 実際の implementation-facing internal design document が複数生まれ、spec からそれらを機械的に辿る navigation / impact analysis requirement が確認された場合。
- Investigation / work item / MCP query が、internal design artifact を canonical ref として参照・検証する必要を持った場合。
- `yaml:` active 化により、spec / internal design / YAML の realization chain または cross-layer validation が operational requirement になった場合。
- Relation entry または reverse graph が、M19 の resolver 実装にとって単なる example ではなく acceptance requirement として別 requirement に捕捉された場合。

## 横断的な観測事実

### 1. 調査時点の authority と draft のずれ

- 調査時点で、accepted ADR-084 は `internal-design:` / `coverage:` を active prefix に含めていた。
- M18 の未 commit relation narrowing draft は、`spec:` → `internal-design:` mapping を MVP operational relation とする案を記述していた。
- 調査時点の `docs/spec/concepts/project-artifact-model/index.md` は draft でありながら INV-DOCS-002 Candidate B の方向を先取りして external coverage を MVP 外として記述していた。
- 一方、調査時点の traceability leaf draft specs と `docs/internal-design/README.md`、resolver example、M18 / M19 tasks は external coverage / `COV-TRACE-001` 前提を残していた。

したがって、調査時点の worktree は external coverage の扱いに関して既に draft 間不整合を含み、さらに internal-design relation の MVP 必須性は未判断であった。後続判断前に spec / task を部分同期すると、二段階の前提見直しが混ざる危険があると評価した。

### 2. MVP の実質的な価値軸

- 調査時点で ADR-087 / REQ-MCP-001 から確認できた実装価値は、design record と investigation の探索、canonical ref resolution、参照切れ validation であった。
- `spec` → internal design の semantic realization relation は、調査時点の一次情報ではその value axis を成立させる必須条件ではなく、追加された trace example / mapping validator の対象であった。
- MVP の目的を「将来 relation graph の完全な最小断面」ではなく「現在必要な canonical reference / validation foundation」と捉えるなら、internal-design relation の deferred は scope reduction として自然である。

## 候補比較

| criterion | Candidate A: INV-DOCS-002 Candidate B どおり internal-design metadata relation を active にする | Candidate B: `internal-design:` ref resolve のみ active にし、`spec` relation は defer | Candidate C: semantic trace MVP から `internal-design:` endpoint と relation を defer |
| --- | --- | --- | --- |
| internal design artifact layer の維持 | 維持 | 維持 | 維持 |
| MVP relation metadata / reverse graph | 必要 | 不要 | 不要 |
| internal-design canonical ref resolution | 必要 | 必要 | MVP 外 |
| 現行の concrete requirement による裏付け | bootstrap example 中心 | internal-design を参照する実需が確認できれば成立 | canonical spec / investigation validation foundation に集中できる |
| M19 実装負荷 | relation parse / validation / reverse query が必要 | ref index 追加だけが必要 | internal-design 対応を後続へ送れる |
| future YAML / impact trace 拡張 | 近い | 再度 relation 導入が必要 | endpoint と relation を同時に再判断 |
| 過剰前提の削減 | external coverage のみ削減 | relation を削減 | endpoint と relation を削減 |

### Candidate A: internal-design metadata relation を MVP active にする

INV-DOCS-002 Candidate B の候補である。External coverage artifact を外しつつ、internal-design artifact が source spec ref を宣言し、resolver / MCP が逆引き graph を導出する。

成立条件は、MVP において spec から internal design を辿る machine-readable navigation、または internal design の source-spec validation が実務上必要であることが確認されることである。現時点で確認できる直接証拠は bootstrap example とそれを前提にした task 記述に限られる。

### Candidate B: `internal-design:` identity resolve のみ active にする

Relation declaration は defer するが、internal-design document を investigation 等から canonical に参照する可能性に備え、document-level `internal-design:` ref の登録と resolve のみを MVP に残す案である。

これは、M18 中に internal design を参照する source/result metadata が実在し、その解決を acceptance に含める必要がある場合に成立する。しかし現存する investigation の internal-design reference は INV-DOCS-002 と本 investigation の調査根拠としての参照であり、semantic trace MVP が自ら作った artifact を根拠に active endpoint を必須化する bootstrap 性は残る。

### Candidate C: semantic trace MVP から `internal-design:` endpoint と relation を defer する

`docs/internal-design/` layer は維持し、文書作成も禁止しない。一方、MVP の active semantic trace contract / M19 required resolver input / relation validation から `internal-design:` と `spec` との semantic realization relation を外す。MVP は、確定済み value である spec semantic refs、artifact ID-as-ref、investigation canonical reference resolution / validation に集中する。

この案は、internal design を恒久的に trace しないという判断ではない。上記 trigger のいずれかが確認された時点で、endpoint、relation direction、metadata owner、MCP query / validation を実在 requirement とともに再判断する。

## 推奨案

本 investigation の結論時点では **Candidate C: semantic trace MVP から `internal-design:` endpoint と `spec` ↔ internal design semantic realization relation を defer する** ことが最も小さく、確認できた実需に一致する候補と考えられた。

理由は以下である。

1. Internal design artifact layer の存在理由は ADR-083 によって独立に成立しており、active trace endpoint 化を延期しても失われない。
2. YAML と external coverage を外した後に残る `spec` → internal design relation は、調査時点で確認できた一件の resolver self-example によって主に支えられており、MVP 利用者の独立 requirement は確認できなかった。
3. ADR-087 / REQ-MCP-001 の accepted value は investigation integration と canonical ref resolution / validation であり、`spec` → internal design relation を必須にしなくても維持できる。
4. Relation の active 化を後続へ送れば、実際の internal design navigation、impact analysis、YAML chain、または canonical internal-design ref 消費者が現れた時点で、endpoint と relation ownership を実需に基づいて設計できる。

調査時点では accepted ADR-084 が `internal-design:` / `coverage:` を MVP active endpoint に含めていた。したがって Candidate C を採用するには、spec や task の単独修正ではなく、後続 ADR による限定 refinement が必要であると結論付けた。後続の corrected ADR-088 がこの refinement を担った。

## 後続判断に渡す候補

- Corrected ADR-088 により、external coverage artifact の除外だけでなく、MVP の `internal-design:` endpoint と semantic realization relation 自体を defer するかを判断する。
- Candidate C を採用する場合、MVP active semantic ref / resolver input / validation acceptance から `internal-design:`、`coverage:`、`COV-*`、`spec` → internal design relation を外し、M19 を investigation canonical reference foundation に再同期する。
- Candidate B を採用する場合、relation declaration は defer し、なぜ internal-design document-level canonical resolve が現在必要かを concrete consumer artifact と acceptance test で明示する。
- Candidate A を採用する場合、bootstrap example 以外の実需、または bootstrap trace を MVP requirement とする明示理由を ADR / requirement に記録する。
- Internal design layer 自体は、いずれの候補でも ADR-083 の implementation-facing documentation boundary として維持する。

## 後続 artifact 候補

Candidate C を採用する場合の影響候補:

| category | affected artifact / handoff | refinement candidate |
| --- | --- | --- |
| ADR | corrected ADR-088 | external coverage だけでなく `internal-design:` endpoint / relation を MVP operational scope から外す判断を記録する |
| project artifact model | `spec:project-artifact-model` | internal design layer の存在と semantic trace active scope を分離して記述する |
| traceability spec | `spec:trace.artifact-refs`, `spec:trace.metadata-schema`, `spec:trace.coverage-mapping`, `spec:trace.resolve-and-validation`, `spec:trace.out-of-scope` | active endpoint / mapping / validator / future trigger を再構成する |
| authoring guidance | `docs/internal-design/README.md`, 必要なら policy docs | internal design は作成可能だが MVP semantic trace metadata を要求しない境界を明記する |
| example | `internal-design:resolver.semantic-ref-index`, `COV-TRACE-001` 相当 | M18/M19 required trace example から外すか、future/example-only として整理する |
| task / requirement | M18, M19, REQ-MCP-001 / WORK-MCP-001 | resolver / validation acceptance を relation graph ではなく confirmed canonical reference use case に合わせて更新する |

## 未確定点

- MVP で internal-design document を canonical `source_refs` / `follow_up_results` として解決すべき、bootstrap ではない concrete consumer が既に別 artifact に存在するか。
- `spec:` semantic ref 自体を investigation source ref resolution のための MVP active target として残す範囲、および spec section-level resolve の M19 acceptance 範囲。
- Candidate C 採用時に `docs/internal-design/resolver/semantic-ref-index.md` を単なる implementation design document として残すか、M19 contract 確定後に更新・置換するか。
- Corrected ADR-088 に external coverage と internal-design endpoint の両判断を統合する場合の spec / task 同期範囲。
- M18 の final independent review を、これらの scope 判断が確定する前に進めるべきか。

## 制約記録

今回利用可能な MCP tool 一覧には filesystem read/write と Design Records MCP の record tool は存在したが、`git status` / `git diff` を実行する git 操作 tool は確認できなかった。Design Records MCP の accepted record 一覧 query は安全性チェックにより実行できなかったため、関連 ADR の本文 status を filesystem MCP で直接確認した。

この制約の下で、調査実施時に新規に書き込んだ対象は本 investigation artifact のみであり、ADR、spec、README、task、example、implementation は調査の一環として変更していない。
