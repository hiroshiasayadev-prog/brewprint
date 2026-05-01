# 056: MCP analyze_impact tool design

- **status**: accepted
- **date**: 2026-05-01
- **depends on**: ADR-047, ADR-048, ADR-049, ADR-054, ADR-055

> このADRは起票時点での決定を記録したスナップショットである。
> 現在の仕様は spec を参照すること。

## 背景

ADR-054 で MCP / QueryService の coverage 方針を「設計対話 coverage」基準に拡張すると決定した。
ADR-055 で `get_reference_tree` を別toolとして定義し、bounded reference graph traversal までを担うことが固定された。
変更種別ごとの意味づけ・severity 判定・recommended action は ADR-055 の対象外とし、`analyze_impact` 側で扱うと整理されている。

M12 task file `docs/tasks/m12-mcp-impact-traversal.md` には `analyze_impact` の input / output 候補が並んでいたが、以下が未確定だった。

- `change` 表現を enum とするか discriminated object とするか
- severity と「機械的に直せるかどうか」を同じ軸で扱うか分けるか
- `coverage` をどう扱うか
- `recommended_action` を一段で返すか、二段で返すか
- v1 で flow wiring / sequence step を含めるか
- render output mapping をどこまで扱うか
- `mechanical` 判定の基準を spec で固定するか

このADRはこれらを固定する。

`analyze_impact` は、人間と LLM が brewprint 上で設計変更相談を行うとき、LLM が「人間に提示できる回答」を作るための判断補助 tool として位置づける。
`get_references` / `get_reference_tree` が返す raw な reference 情報の上に、change kind ごとの解釈を載せて返すことで、LLM が次のような応答を作れるようにする。

- 確実に壊れる箇所と、要確認の箇所と、参考情報の箇所を区別して提示する
- 機械的に直せる箇所と、設計判断が要る箇所を区別して提示する
- どこまで分析しているか、何を分析していないかを明示する

## 決定

### 1. `analyze_impact` を MCP v1 tool として正式採用する

`docs/spec/mcp/versioning.md` の future 候補から `analyze_impact` を昇格させ、MCP v1 tool として正式採用する。
spec 詳細は `docs/spec/mcp/tools/analyze-impact.md` に定義する。

ADR-055 が決めた以下の責務分離は維持する。

```text
get_references       direct references API
get_reference_tree   bounded reference graph traversal API
analyze_impact       change kindを踏まえた意味づけ済み影響分析API
```

`analyze_impact` は ADR-055 を覆さず、補完する。

### 2. `change` は discriminated object 形式とする

`change` は単純 enum ではなく、`kind` discriminator を持つ object とする。

```jsonc
{
  "change": {
    "kind": "rename",
    "new_id": "auth.model.user.email_address"
  }
}
```

理由は、`rename` で `new_type` を渡すような不正組み合わせを schema レベルで弾けることと、kind ごとの payload を schema validation できること。

v1 で扱う `change.kind` は以下とする。

- `rename`
- `remove`
- `change_type`
- `change_contract`
- `change_transition_target`
- `add`

`add` は破壊的影響分析というより、name collision / type resolution / writer coverage 等の整合性分析として扱う。
よって `add` を投げたときの `coverage.analyzed` は他 kind と異なってよい。

### 3. severity と fixability を別軸として分ける

`severity` は「変更がどれくらい壊すか」を表す。
`fixability` は「直し方がどれくらい機械的に決まるか」を表す。
この2つを混ぜない。

`severity` enum:

- `breaking` — そのまま変更すると semantic build / validation / render / query のいずれかが失敗する可能性が高い
- `warning` — 壊れるとは限らないが、意味・到達経路・表示・設計意図が変わる可能性がある
- `info` — 関連情報として提示するが、変更対応は不要または低優先

`fixability` enum:

- `mechanical` — source location と置換内容が一意に決まり、機械的に直せる
- `suggested` — 修正方針は提案できるが、人間レビュー前提
- `manual_review` — 設計判断が必要で、tool では直し方を決めない方がよい
- `unknown` — tool が判断できない（情報不足、coverage 外、source range 欠落 等）

field rename は典型的に `severity=breaking, fixability=mechanical` であり、 「rename だから warning」 のように severity を緩めない。

### 4. `fixability=mechanical` の必要条件を spec で固定する

`fixability=mechanical` を返してよいのは、以下を**すべて**満たすときのみとする。

1. 置換対象 source location が一意に特定できる（file / line / column range）
2. 置換前 token が source 上で一意（誤一致しない）
3. 置換後 token が明確に1つに定まる（衝突なし）
4. 置換後の reference 解決先が変わらない
5. YAML 構造を変えない単純 token 置換である

ひとつでも欠ければ、最低でも `suggested` に下げる。
不確実性が高い場合は `manual_review` または `unknown` を返す。

これら5要件は spec で明示し、実装は judgement gate として持つ。
個別 change kind ごとの細かい heuristic は実装裁量に残す。

### 5. `recommended_action` と `suggested_fixes[]` を二段で返す

`recommended_action` は人間向け説明文。多少抽象でもよい。
`suggested_fixes[]` は機械的に直せそうな候補。`confidence` 必須。

```jsonc
{
  "recommended_action": "reads field reference を email_address へ更新する",
  "suggested_fixes": [
    {
      "kind": "replace_reference",
      "confidence": "high",
      "from": "email",
      "to": "email_address",
      "source": { "file": "...", "line": 42, "column": 7 }
    }
  ]
}
```

`fixability` が `manual_review` / `unknown` のときは `suggested_fixes[]` を空とするか、非破壊的 advisory のみに限定する。

### 6. `coverage` を必須出力とする

`analyze_impact` の output には `coverage` を必須で含める。

```jsonc
{
  "coverage": {
    "analyzed": [...],
    "not_analyzed": [...],
    "note": "..."
  }
}
```

これがない場合、LLM は「0件＝安全」と「未分析＝0件」を区別できず、人間に誤った安心感を与える。
`analyze_impact` は安心感を売る tool であり、どこまで見たか / 見ていないかを返す責務がある。

v1 の `coverage.analyzed` 標準セット:

- `direct_references`
- `reference_tree`
- `model_field_resolution`
- `transition_action_resolution`
- `flow_step_task_resolution`
- `flow_param_field_resolution`
- `sequence_step_task_resolution`
- `type_signature_identity`
- `render_output_files`

v1 の `coverage.not_analyzed` 標準セット:

- `type_structural_compatibility`
- `semantic_contract_compatibility`
- `render_presentation_details`
- `wireframe_element_binding`

実装は `change.kind` や対象 object kind に応じて、これらをサブセット化してよい。
ただし「分析対象だが結果0件」と「分析していない」を区別できる形で返すこと。

### 7. flow wiring / sequence step を v1 coverage に含める

flow wiring（`flow_step` / `flow_param`）と sequence step は、reference 経路の到達可能性レベルでは v1 で分析対象に含める。

ADR-055 §8 では「flow wiring を影響分析に含める場合は `analyze_impact` 側の補完材料として扱う」と保留されていたが、本ADRでこれを v1 で採用する。
ただし `get_reference_tree` の reference kind 拡張は行わない。
`analyze_impact` は内部で `inspect(task).members.flow.entries` 等を読んで、 flow wiring の参照先 task / field を抽出する。

sequence step は `inspect(view: sequence_diagram)` 系の経路で task 参照を抽出する。

これらは ADR-054 の設計対話 coverage 方針と整合する。
sequence diagram に render される object である sequence step が `analyze_impact` の死角になっていると、 設計対話 coverage 方針に反する。

flow param の type 整合性については、 v1 では「型 signature の identity 比較」（primitive 型一致 + model id一致）までを対象とし、 model 間の structural compatibility（subtyping 判定）は v1 除外とする。
これは「型コンパイルが通るか」と「要件通りに動くか」を別レイヤとして扱うため。

### 8. render output は file 粒度のみ対応する

`analyze_impact` は、変更対象を含む render group / render output file path までは返す。
md 内のどこがどう変わるかという presentation 詳細は v1 除外とする。

`render_output_files` impact の典型例:

```jsonc
{
  "kind": "render_output",
  "severity": "info",
  "object": {
    "object": "render_output",
    "path": "renders/auth/dag-login.md"
  },
  "reason": "この task を含む DAG render が再生成される",
  "via": ["render_output_of"],
  "fixability": "mechanical",
  "recommended_action": "brewprint render を再実行して renders/ を更新する"
}
```

これにより、LLM は「この変更で再生成される md は X / Y / Z」と人間に提示できる。
md 内の表現変化（DAG node shape の変化、ER の線の変化等）は renderer 内部詳細であり、 `analyze_impact` の責務外とする。

### 9. `source` は inline 必須、YAML snippet は optional

各 impact entry の `source` には、 file / line / column / end_line / end_column を inline で必須とする。
これにより、 LLM が `get_source` を別途呼ばなくても、 「この箇所を直す」と人間に提示できる。

YAML snippet 全文は `source` には含めない。
必要なら `source_preview` として短い行範囲のみを optional に持つ。
完全な snippet 取得は `get_source` の責務とする。

## 理由

### `analyze_impact` を実装補助 API ではなく判断補助 tool として定義する理由

ADR-054 で MCP / QueryService の coverage 拡張基準は「設計対話 coverage」と決まった。
`analyze_impact` はその設計対話 coverage を最も直接的に支える tool である。
人間が「この field を rename したい、何が壊れる？」と聞いたとき、 LLM が `get_references` の生 list を眺めて自力で意味づけするのではなく、 tool 側が change kind ごとの解釈を返すことで、 LLM が人間に直接提示できる形にする。

これは ADR-054 §決定 §1 の「MCPは設計対話用 query layer である」と整合する。

### severity と fixability を分ける理由

field rename は、参照側を機械的に直せるが、 旧 ID を残したままなら確実に壊れる。
逆に、 task 削除は「機械的に直せない」が「壊れることは確定」している。

severity と fixability を1軸にすると、 「rename は warning（直しやすいから）」 や 「remove は breaking（直しにくいから）」 のように両者が混ざり、 LLM が人間に正しく提示できない。

severity は「変更がどれくらい壊すか」、 fixability は「直し方がどれくらい確定するか」と別軸にすることで、 LLM は

> 確実に壊れる N 件、要確認 M 件、参考 K 件があります。
> このうち N 件は機械的に直せます。

という二軸サマリを作れる。

### `fixability=mechanical` の必要条件を spec で固定する理由

LLM はこの判定結果を人間に説明する。
実装裁量に任せすぎると、 同じ変更がある実装では mechanical、別実装では suggested になり、 LLM の発言が信頼できなくなる。

ただし、個別 change kind ごとの heuristic を全部固定すると、 v1 で実装が硬直化する。
そのため、 「judgement gate（5要件）」 だけを spec で固定し、 個別 rule は実装裁量とする。
これは brewprint MCP の他 spec でも採用しているレベル感（責務と最低条件は spec、 個別実装は裁量）と整合する。

### `coverage` を必須にする理由

`analyze_impact` が「0件＝安全」と「未分析＝0件」を区別できない返却をすると、 LLM は「安全です」と人間に伝えてしまい、 隠れた breaking change を見逃す。

`coverage.not_analyzed` を必須にすることで、 LLM は「flow wiring は v1 では未対応のため、 flow 内部の影響は別途確認が必要です」と人間に伝えられる。
これは `analyze_impact` の信頼性の根幹を成す。

### flow wiring / sequence step を v1 で含める理由

brewprint は近い将来、 brewprint 自身を brewprint で設計する self-hosting に進む。
editor / viewer は brewprint で描かれ、 MCP で操作される。
そのとき最頻出の設計変更は task の signature 変更や transition の rewiring であり、 これらは flow wiring / sequence step に高頻度で影響する。

flow wiring / sequence step を v1 coverage から外すと、 editor / viewer が `analyze_impact` を呼んだときに「ほぼ何も分析できない」状態になる。
それは設計対話 coverage 方針（ADR-054）と矛盾する。

flow wiring を扱う際、 `get_reference_tree` 側の reference kind 拡張は ADR-055 の方針上避け、 `analyze_impact` 内部で `inspect` 系を使って抽出する。
これにより ADR-055 を覆さず、補完する形で v1 を成立させる。

### render output を file 粒度に絞る理由

「どの md が再生成されるか」 は render index から逆引き可能であり、 LLM が人間に提示する上で必須情報。
一方、 「md 内のどこがどう変わるか」 は renderer 内部詳細であり、 これを `analyze_impact` で返すと renderer / QueryService の境界（ADR-047）を侵食する。

file 粒度に絞ることで、 ADR-047 の境界を維持しながら、 LLM が「再生成対象 md」を提示できるようにする。

### `source` を inline 必須にする理由

LLM 視点では、 各 impact について「ここを直せばいい」と人間に提示できることが価値の中心。
inline でなく `get_source` 呼び直しにすると、 N 件 impact があれば N 回の round trip が発生し、 設計対話の応答性が悪化する。

ただし YAML snippet 全文を inline で持つと output が膨れる。
file / line / column の identity 情報のみを inline 必須とし、 全文取得は `get_source` の責務に分離する。

## 影響

### MCP / QueryService への影響

`analyze_impact` は QueryService の新 method として実装される。
内部実装は `get_references` / `get_reference_tree` / `inspect` の結果を組み合わせて構築する。
ADR-047 の境界は維持し、 Raw YAML AST を直接公開しない。

flow wiring / sequence step impact 抽出のため、 QueryService は `inspect(task).members.flow.entries` 相当の経路に内部依存を持つ。
これは既存 inspect 系 spec の範囲内で済むはずであり、 新規 reference kind は導入しない（ADR-055 維持）。

### spec への影響

新規作成:

- `docs/spec/mcp/tools/analyze-impact.md`

更新:

- `docs/spec/mcp/overview.md` — tool overview / selection guidance に `analyze_impact` を追加
- `docs/spec/mcp/versioning.md` — future 候補から `analyze_impact` を外す

### TASKS への影響

`docs/tasks/m12-mcp-impact-traversal.md` の `analyze_impact` task は、 設計完了とともに spec 反映済みに mark する。
実装 task は別 milestone で管理する。

### ADR-055 との関係

ADR-055 は `get_reference_tree` の責務範囲を確定したものであり、 本ADRはこれを覆さない。
ADR-055 §8 の「flow wiring を `analyze_impact` 側の補完材料として扱う」という保留を、 本ADRが「v1 で coverage に含める」 と確定する。

### ADR-054 との関係

ADR-054 §影響 §MCP の future 候補リストに含まれていた `analyze_impact` を v1 採用に格上げする。
ADR-054 の方針（設計対話 coverage 基準）はそのまま継承する。

## Evidence
- commit: tbd
- impl commit: tbd
- 参考: 特になし
