# 086: investigation artifact format and lifecycle

- **status**: accepted
- **date**: 2026-05-19
- **depends_on**: ADR-050, ADR-068, ADR-081, ADR-083, ADR-084, ADR-085, INV-DOCS-001
- **supersedes**:
- **migrated_to_spec**:

> このADRは起票時点での決定を記録したスナップショットである。
> 現在の仕様は spec を参照すること。

## 背景

ADR-085 により、複雑な変更における調査結果、根拠、影響範囲、未確定点、選択肢、後続 artifact 候補を保存する artifact layer として `docs/investigations/` が導入された。

ADR-085 は、investigation が決定、現行仕様、要求そのもの、横断進捗、完了状態、具体的な作業手順を所有しないことを定めた。
また、investigation は requirement / work item / task / ADR / spec / internal design / coverage / 別 investigation の起票・更新前に必ず必要な gate ではなく、複雑な変更で調査結果を後続 artifact へ保存する必要がある場合にのみ使うことを定めた。

一方で、ADR-085 は `docs/investigations/` の詳細 format、status 語彙、ID 規則、MCP index 方針、lifecycle を確定しなかった。
これらは初回 investigation `INV-DOCS-001` で調査することになった。

`INV-DOCS-001` では、directory 運用、ID 規則、file name、metadata、status、scope / non-scope、起点 artifact と後続 artifact の記録方法、別 investigation の扱い、README / doc-policy / authoring guide の責務分担、MCP index 対象化について調査した。

本ADRは、その調査結果を踏まえ、investigation artifact の初期 format と lifecycle を決定する。

## 決定

### 1. directory は domain subdir で運用する

investigation は `docs/investigations/<domain>/` 配下に置く。

例:

```text
docs/investigations/docs/
docs/investigations/adr/
docs/investigations/spec/
```

`docs/investigations/README.md` は root に置き、directory の入口説明とする。

初回 investigation である `INV-DOCS-001` は `docs/investigations/docs/` に置く。

### 2. ID は `INV-<DOMAIN>-NNN` とする

investigation ID は domain ごとの連番で採番する。

```text
INV-<DOMAIN>-NNN
```

例:

```text
INV-DOCS-001
INV-TRACE-001
INV-MCP-001
INV-YAML-001
```

`DOMAIN` は uppercase の短い domain label とする。
investigation ID は、他 artifact の ID / 番号体系から独立して採番する。
`NNN` は investigation の `DOMAIN` ごとに 001 から始まる3桁ゼロ埋め連番とする。
ADR number、requirement ID、work item ID、task milestone、coverage mapping ID などとは結合しない。

### 3. file name は `INV-<DOMAIN>-NNN-<slug>.md` とする

investigation file name は ID を大文字で残す。

```text
INV-<DOMAIN>-NNN-<slug>.md
```

例:

```text
INV-DOCS-001-investigation-artifact-format-and-lifecycle.md
```

ID を file name に残すことで、人間が directory listing から ID を確認しやすくする。

### 4. metadata は required / optional を分ける

investigation は Markdown 冒頭に bullet metadata を置く。
初期運用では YAML front matter を必須化しない。

required metadata は以下とする。

- `status`
- `date`
- `trigger`
- `scope`
- `non_scope`
- `source_refs`
- `follow_up_candidates`

optional metadata は以下とする。

- `supersedes`
- `related_requirements`
- `related_work_items`
- `related_adrs`
- `related_specs`
- `related_internal_design`
- `related_coverage`
- `follow_up_results`

optional metadata は、該当する情報がある場合にのみ書く。
すべての investigation に空 field を義務付けない。

### 5. status は `investigating` / `concluded` / `superseded` とする

investigation の初期 status 語彙は以下とする。

| status | 意味 |
|---|---|
| `investigating` | 調査中 |
| `concluded` | 調査結果がまとまり、後続判断に渡せる状態 |
| `superseded` | 後続 investigation または別 artifact により置き換えられた状態 |

`proposed` は ADR status と混同しやすいため、investigation status には採用しない。
`archived` は `superseded` と異なる意味が必要になった場合に後続判断で追加する。

`concluded` は、後続 artifact の採用判断や実装完了を意味しない。
調査 artifact として判断材料がまとまったことだけを表す。

### 6. scope / non-scope は metadata と本文の両方で扱う

metadata には短い `scope` / `non_scope` を置く。
本文には `## 調査スコープ` / `## 非スコープ` を置き、詳細を書く。

これにより、directory listing や冒頭 metadata で調査対象を素早く把握しつつ、本文で調査範囲と非対象を明確にする。

investigation の scope が広がりすぎる場合は、元 investigation を無制限に拡張せず、別 investigation として切り出してよい。

### 7. 起点 artifact と後続 artifact は `trigger` / `source_refs` / `follow_up_candidates` / `follow_up_results` で分ける

investigation は、起点、調査根拠、後続候補、実際に生まれた artifact を区別する。

- `trigger`: この investigation が起票された理由または起点 artifact
- `source_refs`: 調査根拠として参照する artifact
- `follow_up_candidates`: 調査結果から起票・更新されうる artifact
- `follow_up_results`: 実際に作成・更新された artifact

`trigger` / `source_refs` / `follow_up_candidates` は required metadata とする。
`follow_up_results` は optional metadata とする。

`follow_up_results` は進捗管理 field ではない。
この investigation を根拠に実際に作成・更新された artifact の記録に限る。
作業状態や完了状態の管理は work item / task が所有する。

`follow_up_candidates` が結果として空になることは許容する。
後続 artifact を生まない結論も、investigation の正当な帰結である。

`source_refs` / `follow_up_candidates` / `follow_up_results` の記法は、初期運用では human-readable な artifact ID または path を許容する。
semantic ref 化や MCP による補完は後続判断とする。

optional の `related_*` も同様に、初期運用では investigation document 内の補助参照として human-readable な artifact ID または path を許容する。
これらは関連 artifact 側の primary trace edge を所有するものではなく、関連 artifact の trace は requirement / work item / coverage 等の primary holder が引き続き所有する。

### 8. investigation から別 investigation を起票してよい

investigation の調査中に別領域の調査が必要になった場合、別 investigation を起票してよい。

ただし、別 investigation の起票は必須ではない。
軽微な追加確認は元 investigation 内に留めてよい。

別 investigation を起票する場合、元 investigation の `follow_up_candidates` または `follow_up_results` に記録する。
これにより、調査の分岐を追跡できる。

### 9. README / doc-policy / ADR authoring guide の責務を分ける

investigation に関する文書の責務は以下とする。

| artifact | 所有するもの |
|---|---|
| `docs/investigations/README.md` | investigation directory の入口説明、初期 format / lifecycle の実務ガイド |
| `docs/doc-policy.md` | docs layer としての `docs/investigations/` の存在と最小責務説明 |
| `docs/adr-authoring-guide.md` | ADR が判断前の探索ログや選択肢比較を抱え込まないための境界説明 |

`docs/doc-policy.md` と `docs/adr-authoring-guide.md` に investigation の完全 format を重複して書かない。
format / lifecycle の実務ガイドは `docs/investigations/README.md` が所有する。

### 10. MCP index 対象化は将来方針として認めるが、interface は後続判断とする

investigation は、将来的な MCP index / query / validate 対象になりうる。

ただし、本ADRでは Design Records MCP に `kind: investigation` を追加することは決定しない。
また、別 MCP interface として扱うかも決定しない。

Design Records MCP に統合するか、別 MCP interface とするか、また `source_refs` / `follow_up_candidates` / `follow_up_results` を MCP 側でどう補完・検証するかは後続 ADR / spec / implementation task で扱う。

## 理由

### なぜ domain subdir にするか

investigation は、docs、traceability、MCP、YAML、implementation など複数 domain に広がりうる。
root 直下にすべて置くと、数が増えたときに調査対象の把握が難しくなる。

一方で、時系列 directory は semantic ID と関係が薄く、調査対象の探索性を高めにくい。
そのため、初期から domain subdir を採用する。

### なぜ `INV-<DOMAIN>-NNN` にするか

ADR-081 / ADR-084 では、requirement / work item / coverage mapping の ID 候補として domain-scoped sequence が示されている。

investigation も同様に domain-scoped ID とすることで、artifact 種別と domain を ID から読み取れる。
ADR 番号と結合しないことで、ADR とは異なる lifecycle を保つ。

### なぜ required / optional metadata を分けるか

すべての field を required にすると、軽量な investigation でも空 field が増える。
一方で、最低限の trigger、scope、source_refs、follow_up_candidates がないと、なぜ起票されたか、何を調べたか、何に波及しうるかが追えない。

そのため、調査の追跡に必要な field を required とし、関連 artifact や follow_up_results は optional とする。

### なぜ `follow_up_results` を optional にするか

`follow_up_results` は、この investigation を根拠に実際に作成・更新された artifact を記録するには有用である。

しかし、これを required にすると、investigation が進捗管理や完了状態を所有しているように見える危険がある。
ADR-085 の責務境界に従い、作業状態や完了状態は work item / task が所有する。
そのため、`follow_up_results` は optional とし、作成・更新された artifact の記録に用途を限定する。

### なぜ `concluded` を採用するか

investigation は採用判断を所有しないため、ADR の `accepted` に相当する status は不要である。

一方で、調査が後続判断に渡せる状態になったことは表現したい。
そのため、調査結果がまとまった状態として `concluded` を採用する。

### なぜ README / doc-policy / ADR authoring guide の責務を分けるか

`docs/investigations/README.md` は investigation を書く人の入口であり、format / lifecycle の詳細を書くのに適している。

`docs/doc-policy.md` は docs 全体の入口であり、詳細 format を持つと肥大化する。
`docs/adr-authoring-guide.md` は ADR の書き方を所有する文書であり、investigation format 全体を持つと責務が広がりすぎる。

そのため、詳細 format は README に置き、doc-policy と ADR authoring guide には最小責務境界だけを置く。

## 却下した代替案

### 代替案A: `docs/investigations/` 直下に全 investigation を置く

- 利点: 初期実装が単純
- 欠点: investigation が増えたときに domain ごとの探索性が落ちる

→ 却下。初期から domain subdir を採用する。

### 代替案B: `INV-NNN` の単純連番にする

- 利点: ID が短い
- 欠点: domain が ID から分からず、requirements / work-items / coverage mapping の ID 方針とも揃わない

→ 却下。`INV-<DOMAIN>-NNN` を採用する。

### 代替案C: `follow_up_results` を required にする

- 利点: 調査結果から実際に生まれた artifact を常に追える
- 欠点: investigation が進捗管理や完了状態を所有しているように見える

→ 却下。`follow_up_results` は optional とし、作成・更新 artifact の記録に限定する。

### 代替案D: investigation format を doc-policy / ADR authoring guide に直接書く

- 利点: セッション開始時や ADR authoring 時に参照されやすい
- 欠点: doc-policy / ADR authoring guide の責務が膨らみ、同じ内容の重複管理になる

→ 却下。format / lifecycle は `docs/investigations/README.md` が所有する。

### 代替案E: Design Records MCP に直ちに `kind: investigation` を追加する

- 利点: investigation を list / get / validate できる
- 欠点: Design Records MCP の責務拡張と contract 変更を伴う。ADR-084 の semantic trace MVP scope も不用意に広がる

→ 却下。本ADRでは将来方針に留め、interface は後続判断とする。

## 影響

### docs/investigations への影響

`docs/investigations/README.md` は、本ADRの決定を反映して更新する必要がある。

`INV-DOCS-001` は、以下の path に置く。
既に移動済みの場合は、README と参照先がこの path と一致していることを確認する。

```text
docs/investigations/docs/INV-DOCS-001-investigation-artifact-format-and-lifecycle.md
```

`INV-DOCS-001` の status は `concluded` とする。
ただし、`follow_up_results` の記載は optional とする。

### docs/doc-policy.md への影響

`docs/doc-policy.md` には、`docs/investigations/` の存在と最小責務を反映する必要がある。
詳細 format / lifecycle は doc-policy に書かない。

### docs/adr-authoring-guide.md への影響

`docs/adr-authoring-guide.md` には、ADR が判断前の探索ログ、影響範囲調査、選択肢比較、未確定論点の蓄積を抱え込まず、必要に応じて investigation を参照するという境界を反映する必要がある。

ADR authoring guide に investigation format 全体は書かない。

### ADR-083 への影響

ADR-083 の artifact placement decision rule は、investigation layer を含む形で後続更新または後続ADRによる refinement が必要になる。

### M18 task への影響

`docs/tasks/m18-semantic-traceability-foundation.md` は、ADR-085 / ADR-086 / `docs/investigations/` の導入を踏まえて、references / scope / done criteria を更新する必要がある。

### MCP への影響

本ADRは MCP implementation の即時変更を要求しない。

将来的に investigation を MCP index / query / validate 対象にする場合は、Design Records MCP に統合するか、別 MCP interface とするかを後続 ADR / spec / task で決める。

## Evidence

- commit: tbd
- impl commit: tbd
- 参考: ADR-085 investigation artifact boundary, INV-DOCS-001 investigation artifact format and lifecycle
