# 069: TypeRef container complexity と anonymous inline struct 不採用

- **status**: accepted
- **date**: 2026-05-10

> このADRは起票時点での決定を記録したスナップショットである。
> 現在の仕様は spec を参照すること。

## 背景

ADR-060 / `docs/spec/type-ref.md` では、`list<T>` / `dict<T>` を built-in container TypeRef として導入した。
TypeRef は再帰的に定義されるため、構文上は以下のような nested TypeRef が valid である。

```txt
list<diagnostic>
dict<list<diagnostic>>
list<dict<impact_entry>>
```

これは、UC-002 self-hosting で観測された MCP 公開 contract の nested schema 需要に対応するために必要である。
一方で、深い nested TypeRef や `any` を含む opaque な container TypeRef は、後から読んだときに shape の意味が追いにくい。

特に `dict<T>` は、value 型だけでは設計意図を十分に表現できない。
`dict<T>` の key は常に `str` だが、その key が FileID / object id / code / severity / module id のどれを表すかは TypeRef だけでは分からない。
そのため、`dict<T>` は `list<T>` よりも key semantics の明示が重要になる。

この問題への対処案として、TypeRef に anonymous inline struct を追加する案も考えられる。
しかし brewprint は実装言語の型システムではなく、人間と LLM が設計意図を共有するための設計層である。
設計層で匿名 shape を許容すると、数日後または別 session の LLM が読んだときに、その shape が何の概念だったのかを復元しにくくなる。

M15 Phase C では、`list<T>` / `dict<T>` の深さ制限または lint 方針を決めることになっている。
本ADRは、nested TypeRef を validation error にするのではなく、意味が失われやすい container TypeRef を warning diagnostic の候補として扱う方針を定める。
また、TypeRef variant として anonymous inline struct を v1.1 では導入しないことを定める。

## 決定

### 1. nested `list<T>` / `dict<T>` は構文上 valid のまま維持する

`list<T>` / `dict<T>` の recursive TypeRef 構文は維持する。

以下は構文上 valid とする。

```txt
list<diagnostic>
dict<diagnostic>
list<dict<diagnostic>>
dict<list<diagnostic>>
```

通常の nested TypeRef は validation error にしない。

### 2. error にするのは構文不正と parser safety limit のみ

TypeRef として解釈できない文字列は、既存の `invalid_type_ref` error とする。

例:

```txt
list<
dict<>
list<user
```

また、parser / implementation safety のために定める最大深さを超えた TypeRef は hard error とする。
これは設計上の可読性 lint ではなく、実装保護のための hard limit として扱う。

parser safety limit は、TypeRef container nesting depth が 16 を超えた場合に `invalid_type_ref` error とする。

ここでいう container nesting depth は、`list<T>` / `dict<T>` の入れ子数で数える。
primitive / named model は depth 0 とする。

| TypeRef | depth |
|---|---:|
| `diagnostic` | 0 |
| `list<diagnostic>` | 1 |
| `dict<list<diagnostic>>` | 2 |
| `list<dict<list<any>>>` | 3 |

この limit は可読性 lint ではなく、parser / implementation safety のための hard limit である。
通常の設計上の深さ警告は warning diagnostic 側で扱う。

### 3. anonymous inline struct TypeRef は v1.1 では導入しない

v1.1 の TypeRef variant は ADR-060 の最小構成を維持する。

- primitive
- named model
- inline `list<T>`
- inline `dict<T>`

以下のような anonymous inline struct TypeRef は導入しない。

```txt
list<{ id: str, severity: str }>
```

brewprint では、設計上意味を持つ shape には名前を付けることを優先する。
anonymous inline struct は実装言語では便利な場合があるが、brewprint の責務は実装そのものではなく、実装を明瞭にするための設計契約を残すことである。

したがって、M15 v1.1 では anonymous inline struct ではなく named model への切り出しを推奨する。

### 4. 読みにくい TypeRef は validation error ではなく warning diagnostic の候補とする

本ADRで「lint」と呼ぶものは、実装上は `severity: warning` の diagnostic として扱う。
validation error にはしない。

理由:

- 一時的な task 間受け渡しで inline container が有用な場面がある
- すべての shape に public model を強制すると model file が増えすぎる
- 一方で、無警告にすると `any` や opaque nested container が設計負債化する
- warning により、後から見直すべき箇所を明示できる

warning は validation 成功扱いとする。
`brewprint validate` は warning を出してよいが、error count には含めない。
MCP / render / inspect は、warning を含む project でも動作可能とする。

### 5. 深さそのものより「意味が named model に回収されているか」を重視する

単純な nesting depth だけで判断しない。

以下のように、内部に named model がある場合は比較的安全である。

```txt
list<diagnostic>
dict<list<diagnostic>>
list<dict<impact_entry>>
```

問題にするのは、以下のような opaque container である。

```txt
list<any>
dict<any>
list<dict<any>>
dict<list<any>>
list<dict<list<any>>>
```

特に `any` が nested container 内に入る場合は、shape の意味が失われやすい。
このような TypeRef は warning diagnostic の対象にできる。

### 6. `dict<T>` は key semantics の明示を推奨する

`list<T>` は「同じ shape が並ぶ」ことを表すため、`T` が named model であれば意味を追いやすい。

一方で `dict<T>` は、value 型だけでは不十分である。
key が何を表すかが TypeRef だけでは分からないため、`dict<T>` を使う場合は key semantics を明示することを推奨する。

明示方法は以下のいずれかとする。

- field 名で明らかにする
- model 名で明らかにする
- `note` で明記する

例:

```yaml
fields:
  - name: diagnostics_by_file
    type: dict<list<diagnostic>>
    note: "keyはFileID"
```

または named dict model として切る。

```yaml
nodes:
  - id: diagnostics_by_file
    type: model
    kind: dict
    value: list<diagnostic>
    note: "keyはFileID。valueはそのfileに紐づくdiagnostic一覧。"
```

key semantics が field 名 / model 名 / note から追いにくい `dict<T>` は warning diagnostic の対象にできる。
ただし、この判定は構造だけで決まらない heuristic であるため、初期実装対象に含めるかどうかは後続の lint 実装時に判断してよい。

### 7. helper shape は named model へ切り出すことを推奨する

`list<dict<any>>` のような opaque TypeRef は、できるだけ named model に切り出す。

例:

```txt
list<dict<any>>
```

ではなく、以下のように名前を持つ model を導入する。

```txt
list<impact_entry>
```

named model にすることで、shape に以下を与えられる。

- stable ID
- fields / element / value
- `note`
- MCP query / reference / future catalog render の対象

ただし、named model を public model として置くか、file-private helper model として置けるようにするかは、本ADRでは決めない。
file-private helper model の visibility / 名前解決 / MCP exposure / render exposure は、後続ADRで扱う。

### 8. M15 で追加する warning diagnostic

本ADRでは warning diagnostic の方向性を定める。
M15 v1.1.0-spec では、以下の warning diagnostic を追加する。

| code | severity | condition |
|---|---|---|
| `opaque_type_ref` | warning | container TypeRef の内部に `any` が含まれ、shape の意味が named model に回収されていない |

`opaque_type_ref` は validation 成功扱いとする。
message では、必要に応じて named model への切り出しを提案する。

例:

```txt
opaque container TypeRef `list<dict<any>>`; consider introducing a named model for the element shape.
```

### 9. 将来 lint 候補

以下は有用だが、初期実装では heuristic / 閾値調整が必要なため、M15 の必須実装には含めない。

| code | reason |
|---|---|
| `unclear_dict_key` | key semantics の判定が field 名 / model 名 / note に依存し、構造だけで確定しにくい |
| `deep_type_ref` | 単純な depth だけでは可読性問題を判断できず、named model によって意味が回収されている場合がある |

優先度は以下とする。

```txt
opaque_type_ref
unclear_dict_key
deep_type_ref
```

この優先順は、意味の消失が大きいものを先に扱うためである。
`any` を含む opaque container は shape そのものを失いやすい。
`dict<T>` の key semantics は重要だが、判定は field 名 / model 名 / note に依存するため heuristic になりやすい。
単純な nesting depth は表面的な指標であり、named model によって意味が回収されている場合は必ずしも問題ではない。

### 10. M15 v1.1.0-spec への含め方

M15 v1.1.0-spec では、以下を確定する。

- nested `list<T>` / `dict<T>` を error にしない
- malformed TypeRef と parser safety limit 超過のみ error とする
- parser safety limit の値を spec に反映する
- anonymous inline struct TypeRef は v1.1 では導入しない
- `opaque_type_ref` warning diagnostic を spec に追加する
- `unclear_dict_key` / `deep_type_ref` は将来 lint 候補として残す
- warning 時は named model への切り出しを促す
- file-private helper model と model catalog view は後続ADRで扱う

## 理由

### なぜ error にしないか

nested container TypeRef は、MCP / API / tool I/O のような JSON contract を扱う上で実用上必要になる。
これを validation error にすると、ユーザーは named model へ切る前に `any` へ逃げる可能性が高い。

brewprint の目的は、型システムを強制することではなく、人間と LLM が設計意図を共有しやすくすることである。
したがって、通常の nested TypeRef は許容しつつ、意味が失われやすい箇所を warning として表面化する方が適切である。

### なぜ anonymous inline struct を導入しないか

anonymous inline struct は、実装言語では便利な場合がある。
しかし brewprint は実装言語ではなく、設計層である。

設計層で匿名 shape を許容すると、後から読んだときにその shape が何の概念を表していたのか分からなくなりやすい。
特に LLM は session を跨ぐと過去の暗黙文脈を保持しないため、名前のない shape は解釈が不安定になる。

brewprint では、実装の自由度を実装言語に委ねる一方で、設計上の概念には名前を与える。
そのため、anonymous inline struct ではなく named model を使う方がプロダクトのミッションに合う。

### なぜ depth だけで判断しないか

`list<diagnostic>` や `dict<list<diagnostic>>` のように、内部に named model がある TypeRef は比較的読みやすい。
一方で、浅くても `dict<any>` のように shape や key semantics が不明な TypeRef は危険である。

問題の本質は nesting depth そのものではなく、意味が named model / field 名 / note に回収されているかである。

### なぜ dict を特別視するか

`dict<T>` は value 型だけでなく key semantics が重要である。
key が FileID なのか、object id なのか、diagnostic code なのかは TypeRef だけでは表現できない。

そのため、`dict<T>` は `list<T>` よりも note / model 名 / field 名による説明を強く推奨する。

### なぜ file-private helper model を本ADRで定義しないか

file-private helper model は、anonymous inline struct を避けるための有力な解決策である。
しかし、これを仕様化するには以下の論点を決める必要がある。

- public model と file-private model の visibility 境界
- model file 内の main model / helper model の扱い
- task file 内 helper model と model file 内 helper model の両方を認めるか
- file-private model の名前解決
- MCP query / inspect / reference tree での露出
- render / future model catalog view での露出
- public model への昇格ルール

これらは TypeRef container complexity より大きいスコープであり、本ADRに含めると論点が膨らみすぎる。
したがって、本ADRでは named model への切り出しを推奨するに留め、file-private helper model は後続ADRで扱う。

## 却下した代替案

### 代替案A: nested `list<T>` / `dict<T>` を validation error にする

- 利点: TypeRef が浅く保たれる
- 欠点: MCP / API / tool I/O の nested schema 需要に合わない。ユーザーが `any` に逃げる可能性が高い

→ 却下。通常の nested container は valid とし、可読性問題は warning diagnostic 領域で扱う。

### 代替案B: 深さ閾値を超えた TypeRef を hard error にする

- 利点: parser / implementation は単純になる
- 欠点: 深さだけでは意味の有無を判断できない。`list<dict<diagnostic>>` のように named model が意味を回収しているケースまで拒否しうる

→ 却下。parser safety limit は hard error とするが、設計上の深さは warning diagnostic 領域で扱う。

### 代替案C: anonymous inline struct TypeRef を導入する

```txt
list<{ id: str, severity: str }>
```

- 利点: 小さな one-off shape を手軽に書ける
- 欠点: TypeRef variant が増える。compatibility / nesting / lint の論点が増える。shape に名前がないため、設計意図が後から追いにくくなる

→ 却下。brewprint は実装言語ではなく設計層であり、設計上の shape には named model を使う。

### 代替案D: すべての helper shape を public model として強制する

- 利点: 参照・検索・MCP exposure が単純になる
- 欠点: 小さな one-off helper shape まで public model file として増え、model file が爆発する。公開 surface が広がりすぎる

→ 却下。helper shape の visibility は後続ADRで扱う。少なくとも本ADRでは、public model 強制を前提にしない。

## 影響

### spec への影響

本ADR受理後、以下を更新する。

- `docs/spec/type-ref.md`
  - nested `list<T>` / `dict<T>` は valid のまま維持すること
  - malformed TypeRef / parser safety limit 超過のみ `invalid_type_ref` error とすること
  - parser safety limit の container nesting depth を 16 とすること
  - anonymous inline struct TypeRef は v1.1 では導入しないこと
  - opaque container は warning diagnostic 対象であること
  - unclear dict key / deep TypeRef は将来 lint 候補であること
  - `dict<T>` は key semantics を field 名 / model 名 / note で明示することを推奨すること

- `docs/spec/diagnostics.md`
  - M15 で `opaque_type_ref` を warning diagnostic として追加する
  - `unclear_dict_key` / `deep_type_ref` は将来 lint 候補として残し、初期 diagnostics spec への必須追加対象にはしない

- `docs/tasks/m15-data-layer-expressiveness.md`
  - Phase C の inline struct 検討項目は、本ADRにより v1.1 では不採用と整理する
  - file-private helper model と model catalog view は後続ADR候補として追加する

### 実装への影響

- TypeRef parser は nested `list<T>` / `dict<T>` を valid として扱う
- malformed TypeRef と parser safety limit 超過は `invalid_type_ref` error とする
- TypeRef parser に anonymous inline struct variant は追加しない
- `opaque_type_ref` warning diagnostic は、TypeRef AST に対して `any` を含む opaque container を検査する
- `unclear_dict_key` / `deep_type_ref` は将来 lint 候補とし、M15 初期実装の必須対象にはしない
- warning は validation 成功扱いとする

### UC-002 への影響

UC-002 で現在 `any` に逃がしている nested schema は、M15 以降、以下のいずれかへ寄せられる。

- public named model
- 後続ADRで許可された場合の file-private helper model
- warning を受け入れた上での inline container / `any`

anonymous inline struct TypeRef へは移行しない。

### 後続ADR候補

本ADRから、少なくとも以下の後続ADR候補が発生する。

- public model と file-private helper model の visibility / 名前解決 / MCP exposure / render exposure
- model catalog / schema catalog view による model 群の俯瞰 render

これらは、anonymous inline struct を導入しない方針を実用上成立させるための補完策である。

## Evidence

- commit: tbd
- impl commit: tbd
- 参考: UC-002 MCP公開contract YAML における nested schema / `any` 暫定表現
