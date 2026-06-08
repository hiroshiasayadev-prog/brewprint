# V01-ADR-092: Design Records MCP workflow artifact record and relation boundary

- **status**: accepted
- **date**: 2026-05-27
- **depends_on**: V01-ADR-087, V01-ADR-088, V01-ADR-090, V01-ADR-091
- **supersedes**:
- **migrated_to_spec**:

> このADRは起票時点での決定を記録したスナップショットである。
> 現在の仕様は spec を参照すること。

## 背景

V01-ADR-087 は、Design Records MCP に `investigation` を既存の record-oriented surface の record kind として追加し、`decision` / `spec` / `investigation` を同一の index / query / validation boundary で扱う判断を行った。また、investigation の `source_refs` / 記載済み `follow_up_results` を canonical reference として resolve / validate し、`follow_up_candidates` は canonical form を検査する一方で未作成候補の存在は要求しない境界を確定した。

V01-ADR-088 は semantic trace MVP を canonical reference resolution foundation に限定し、active semantic ref を `spec:` に絞りつつ、Design Records MCP が扱う record ID-as-ref と investigation canonical references の resolve / validation を operational scope とした。

V01-ADR-090 は、選択済みの複数 record を同一 representation で取得する `get_records` を追加したが、その時点の成功対象は `decision` / `spec` / `investigation` に限定し、requirement / work item / task support は後続判断へ送った。

V01-ADR-091 は、active workflow artifact を `requirement -> work item -> task` の三層とし、workflow artifact 間の canonical relation は `REQ-*` / `WORK-*` / `TASK-*` ID-as-ref のみを用いることを確定した。Physical path、`req:` / `work:` / `task:` semantic prefix、および新しい milestone relation は導入しない。また、この workflow artifact の MCP support 自体は `V01-REQ-MCP-003` の対象として後続判断へ送った。

`V01-REQ-MCP-003` / `V01-WORK-MCP-003` の dogfooding では、LLM が requirement から work item、work item から task と dependency、さらに判断・検証 evidence まで辿るには、workflow artifact を MCP 経由で取得・解決・検証する経路が必要であることを整理した。現行 MCP では `REQ-*` / `WORK-*` / `TASK-*` は `get_record(s)` の成功対象ではなく、`resolve_reference` の supported record ID-as-ref にも含まれず、workflow relation validation も存在しない。

加えて、`spec:project-artifact-model` は investigation から requirement への要求候補および work item への後続候補という relation を既に定義している。したがって workflow artifact support は workflow chain 内部だけの問題ではなく、既存 investigation relation を canonical reference / validation contract に接続する問題でもある。

本判断では、workflow artifact を既存 record-oriented surface へ追加する最小 public contract と、その範囲で必要な relation validation boundary を確定する。

## 決定

### 1. `requirement` / `work_item` / `task` を既存 record-oriented surface に追加する

Design Records MCP は、既存の `decision` / `spec` / `investigation` に加えて、以下を public record kind として追加する。

- `requirement`
- `work_item`
- `task`

これらは別の workflow 専用 MCP interface に分離しない。V01-ADR-087 が investigation integration で採用した既存 record-oriented extension と同様に、artifact の責務・lifecycle・metadata semantics の差分は kind 固有 detail object で表現し、探索・取得・解決・検証の tool surface は共有する。

Record response は既存の共通 field と kind 固有 detail object の構造を維持する。Workflow artifact が保持する具体 metadata field と status vocabulary の public schema は spec で定義する。

### 2. `list_records` / `get_record` / `get_records` は workflow artifact を対象に含める

`list_records`、`get_record`、`get_records` は、`requirement` / `work_item` / `task` record を取得対象に含める。

`get_record` / `get_records` は既存責務を維持し、指定された exact record ID の detail retrieval のみを担う。Workflow relation の traversal、progress 集約、graph projection、validation は兼務しない。

`get_records` の partial result、duplicate requested ID、raw body、response ordering の contract は V01-ADR-090 の既存判断を維持し、workflow artifact にも同じ retrieval behavior を適用する。

`list_records` の workflow kind 追加に伴う filter / ordering の具体 contract は spec に定める。一方、ADR 用の `id_range` を domain-scoped workflow ID へ一般化することは本判断の必須範囲に含めない。

### 3. `resolve_reference` は workflow ID-as-ref を supported record ID input として扱う

`resolve_reference` の supported record ID-as-ref に、以下の workflow artifact ID を追加する。

- `REQ-<DOMAIN>-NNN`
- `WORK-<DOMAIN>-NNN`
- `TASK-<DOMAIN>-<WORK-SEQUENCE>-<TASK-SEQUENCE>`

本ADRは、workflow artifact を MCP public record / resolver target として追加するにあたり、現行 authoring convention の上記 ID form を MCP-supported workflow record ID grammar として採用する。詳細な parser rule、filename / ID consistency rule、invalid input diagnostic は spec で定義する。

Resolved workflow artifact ID-as-ref は、既存 record ID resolution と同様に record target へ解決される。

Resolver は task ID の文字列構造または physical path から parent work item / source requirement / dependency relation を推測しない。Workflow relation は artifact metadata に明示された ID-as-ref を読み、validation がその整合性を確認する。

Workflow artifact の section-level addressing は本判断の対象外であり、`req:` / `work:` / `task:` semantic prefix は導入しない。

### 4. Workflow relation の宣言済み integrity validation を MVP に含める

`validate_records` は、workflow artifact の canonical relation metadata について、少なくとも参照先の存在確認を行う。

| source kind | field | required target form |
|---|---|---|
| `requirement` | `work_items` | `WORK-*` |
| `work_item` | `source_requirement` | `REQ-*` |
| `work_item` | `tasks` | `TASK-*` |
| `task` | `work_item` | `WORK-*` |
| `task` | `source_requirement` | `REQ-*` |
| `task` | `depends_on` | `TASK-*` |

また、metadata が宣言する workflow structural relation について、以下の相互整合性確認を MVP に含める。

- `requirement.work_items` が指す work item の `source_requirement` は、その requirement と一致しなければならない。
- `work_item.source_requirement` が指す requirement の `work_items` は、その work item を含まなければならない。
- `work_item.tasks` が指す task の `work_item` は、その work item と一致しなければならない。
- `task.work_item` が指す work item の `tasks` は、その task を含まなければならない。
- task の `source_requirement` は、その task の parent work item が持つ `source_requirement` と一致しなければならない。

`task.depends_on` の参照先存在確認は MVP に含める。一方、dependency が同一 work item 内に閉じるべきか、cycle detection、dependency からの実行順 projection などの追加 graph rule は、具体的な必要性と spec refinement に委ね、本判断では必須化しない。

これらは既に宣言された canonical relation の integrity validation であり、未接続 artifact を探索する orphan diagnostics とは区別する。

### 5. Investigation canonical reference validation は requirement / work item ID-as-ref を扱えるようにする

V01-ADR-087 が validation 対象とした investigation metadata field は、`spec:project-artifact-model` が既に定義する investigation から requirement / work item への relation を operationalize するため、`REQ-*` / `WORK-*` を canonical record ID-as-ref として扱えるようにする。

対象 field と既存 validation semantics は以下のとおり維持する。

| investigation field | `REQ-*` / `WORK-*` treatment |
|---|---|
| `source_refs` | 記載された `REQ-*` / `WORK-*` は resolve 可能でなければならず、unresolved は error とする |
| `follow_up_results` | 記載された `REQ-*` / `WORK-*` は resolve 可能でなければならず、unresolved は error とする |
| `follow_up_candidates` | `REQ-*` / `WORK-*` の canonical form を許容し、未作成 target の unresolved 自体は error とせず info とする |

`spec:project-artifact-model` が明示する investigation から requirement / work item への relation は、これにより MCP 上で canonical reference として表現・検証できる。

`TASK-*` は workflow artifact 間の relation と direct resolver input としては support するが、investigation metadata field の canonical reference 対象には追加しない。Investigation から task への relation は `spec:project-artifact-model` に定義されておらず、task は work item 配下の短期 concrete work として扱う V01-ADR-091 の責務境界を維持する。Investigation metadata に `TASK-*` が現れた場合の concrete diagnostic behavior は、この unsupported boundary に従って spec で定義する。

Investigation の `trigger` および optional `related_*` field の resolve / validation rule は現行 contract と同じく本判断の対象外とし、今回同時に operationalize しない。

### 6. Physical path と workflow semantic prefix は support しない

Workflow artifact 自身の identity と workflow artifact 間 relation は、`REQ-*` / `WORK-*` / `TASK-*` ID-as-ref を用いる。Investigation metadata から workflow artifact への canonical reference は、`spec:project-artifact-model` が定義する relation boundary に従い、`REQ-*` / `WORK-*` に限定する。

Physical path は canonical relation / canonical reference として support しない。`req:` / `work:` / `task:` semantic prefix も導入しない。

この境界は V01-ADR-091 の決定および既存 canonical reference foundation と整合させるためである。

### 7. Orphan diagnostics と progress projection は MVP に含めない

以下は本判断の MVP public contract に含めず、dogfooding で concrete requirement が確認された場合の後続判断へ送る。

- orphan requirement / orphan work item / orphan task diagnostics
- task status から work item progress を導出する projection
- workflow 専用 traversal / tree / graph query tool
- task dependency cycle detection や execution order projection
- investigation の `trigger` / optional `related_*` の resolve / validation 拡張
- section-level workflow reference または workflow semantic prefix

Orphan diagnostics は宣言済み relation の integrity 検査ではなく、未接続 artifact を検出する運用診断である。Progress projection は task status を正本とした derived view であり、有用性はあるが record retrieval / canonical resolution / declared relation integrity の成立条件ではない。

## 理由

### なぜ既存 record-oriented surface に統合するか

V01-ADR-087 は、investigation が decision / spec と異なる責務と relation semantics を持つにもかかわらず、設計対話で横断的に探索・取得・検証する必要があるため、既存 record-oriented surface への統合を採用した。

Workflow artifact も、requirements、進行中の work、短期 task と evidence を設計対話で辿る対象である。別 interface に分離すると、investigation から workflow artifact、requirement から work item、work item から task を確認する導線が不必要に分断される。

Kind 固有 detail object を用いれば、各 artifact の責務と lifecycle を混同せずに同一 surface 上で扱える。

### なぜ retrieval と resolution の両方を追加するか

`get_record(s)` のみを追加しても、workflow metadata に書かれた ID-as-ref や investigation から workflow artifact への canonical reference は機械的に解決・検証できない。

逆に resolver だけを追加して record retrieval を公開しないと、LLM が relation target の metadata、task dependency、evidence、status を既存 query surface 上で確認できない。

Workflow artifact chain を読み、根拠を辿り、relation を検証するためには、record retrieval と ID-as-ref resolution を同一の最小 extension として導入する必要がある。

### なぜ宣言済み relation の相互整合性を MVP に含めるか

V01-ADR-091 により、workflow relation は authoring convention ではなく canonical ID-as-ref relation として定義された。参照先の存在だけを検査し、`REQ.work_items` と `WORK.source_requirement`、`WORK.tasks` と `TASK.work_item` の矛盾を許容すると、MCP が取得できる workflow chain の意味が不定になる。

相互整合性確認は、新しい運用診断や集約 view ではなく、宣言済み relation が同一 chain を表していることを保証する最小 integrity validation である。

### なぜ investigation metadata に requirement / work item ID-as-ref を追加するか

`spec:project-artifact-model` は investigation が requirement の要求候補や work item の後続候補を示しうる関係を既に定義している。Workflow artifact の record / resolver support を追加しても、investigation metadata が `REQ-*` / `WORK-*` を canonical reference として扱えなければ、既存 artifact model の relation が MCP 上で分断されたまま残る。

V01-ADR-087 が定めた field ごとの validation semantics を変えず、解決可能な canonical record ID-as-ref の対象を requirement / work item へ拡張することで、既存 boundary と整合した最小 extension になる。`TASK-*` は investigation relation に追加せず、V01-ADR-091 が定める work item 配下の concrete work という境界を維持する。

### なぜ orphan diagnostics と progress projection を含めないか

Orphan diagnostics は、既に宣言された relation の correctness ではなく、本来 relation を持つべき artifact が接続されていないかを診断する追加の運用 policy を必要とする。

Progress projection は、task status を集約して work item 上の派生表示を提供する capability であり、V01-ADR-091 が将来候補として認める一方、aggregation semantics と response surface の追加判断を要する。

今回の目的は workflow artifact を record / reference / integrity validation として利用可能にすることであり、運用診断と derived view を同時に取り込まない方が MVP boundary が明確である。

## 却下した代替案

### 代替案A: Workflow artifact は別 MCP interface で扱う

却下する。Investigation から workflow artifact、requirement から work item、work item から task を辿る設計対話の query / validation 導線を分断する。Kind 固有 detail object によって責務差分は既存 surface 内で表現できる。

### 代替案B: `get_record(s)` だけを追加し、resolver / validation は後続へ送る

却下する。Metadata に記載された canonical ID-as-ref を取得後に上位 agent が手作業で解釈することになり、V01-ADR-087 / V01-ADR-091 の canonical relation 方針を operationalize できない。

### 代替案C: Resolver のみ追加し、workflow artifact を public record kind にしない

却下する。Relation target の metadata、status、task evidence を同一 record query surface で取得できず、V01-REQ-MCP-003 が求める workflow chain の確認経路として不十分である。

### 代替案D: Existence validation のみを行い、相互整合性は後続へ送る

却下する。参照先が存在しても双方向 metadata が異なる requirement / work item / task chain を許容し、canonical workflow relation の意味が壊れる。

### 代替案E: Orphan diagnostics と progress projection も同時に追加する

却下する。Declared relation integrity と、未接続 artifact の運用診断または status-derived view は責務が異なる。Concrete requirement の確認前に同時採用すると public contract を過度に広げる。

### 代替案F: Physical path または workflow semantic prefix を relation identity として追加する

却下する。V01-ADR-091 が定めた ID-as-ref 方針と矛盾し、identity を二重化するか file relocation で relation を stale にする。

## 影響

### Traceability / project artifact model spec への影響

`docs/spec/concepts/project-artifact-model/index.md` および必要な traceability leaf spec は、workflow artifact の public record / ID-as-ref resolve / relation validation が後続判断ではなく採用済み boundary となる場合に追従する必要がある。

Investigation から requirement / work item への concept relation と、MCP が扱う canonical reference 対象の整合も明文化する必要がある。

### Design Records MCP spec への影響

`docs/spec/design-records-mcp/overview.md` / `schema.md` / `tools.md` は、少なくとも以下を定義する必要がある。

- `requirement` / `work_item` / `task` record kind と kind 固有 detail representation
- discovery / metadata parse / status vocabulary の contract
- `list_records` / `get_record` / `get_records` の workflow artifact 対応
- `resolve_reference` の `REQ-*` / `WORK-*` / `TASK-*` input behavior
- workflow structural relation validation と concrete diagnostic category
- investigation validated field における `REQ-*` / `WORK-*` の扱いと、`TASK-*` を対象外とする境界
- MVP 外とした orphan diagnostics / progress projection / workflow traversal boundary

### Implementation / tests / verification への影響

Implementation は、workflow document discovery、metadata parsing、record serialization、resolver、relation validation、tests、runtime verification の更新を必要とする。

具体的な実装順序、test case、runtime evidence、完了条件は `V01-WORK-MCP-003` 配下の短期 task で追跡する。

### Requirement / work item への影響

本判断が accepted となった場合、`V01-REQ-MCP-003` は workflow artifact MCP support の採用結果を反映し、`V01-WORK-MCP-003` は public contract の spec 反映、implementation、tests、runtime verification、close evidence へ進む。

## Evidence

- commit: tbd
- impl commit: tbd
- review: Codex review で major 二件・minor 一件を反映後、再 review で `OK to proceed to acceptance` を確認（2026-05-27）
- 参考: V01-ADR-087, V01-ADR-088, V01-ADR-090, V01-ADR-091, V01-REQ-MCP-003, V01-WORK-MCP-003, V01-TASK-MCP-003-01, V01-TASK-MCP-003-02, `spec:project-artifact-model`, `SPEC-design-records-mcp-tools`
