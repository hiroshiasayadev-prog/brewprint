# V01-ADR-059: task / join return が primitive を許容するよう validation を整える

- **status**: accepted
- **date**: 2026-05-02

> このADRは起票時点での決定を記録したスナップショットである。
> 現在の仕様は spec を参照すること。

## 背景

UC-002（brewprint self-hosting / M14）Phase Aで MCP tool task群をblueprint化したところ、
`returns.model: any` を持つサブtaskで `unresolved_model: any` エラーが発生することが判明した。

具体例: `mcp/task/get_references.yaml` のサブtask `query_service` が
`returns.model: any` を宣言した時点で error。

これは V01-ADR-021 §3 / spec/nodes.md「primitive予約語」が `any` を含む7語の primitive 予約語を定義していることと矛盾する。

### 仕様側の規定（再掲）

V01-ADR-021 §3 / spec/nodes.md primitive予約語節:

```
| primitive | 意味             |
| --------- | ---------------- |
| str       | 文字列           |
| int       | 整数             |
| float     | 浮動小数点数     |
| bool      | 真偽値           |
| bytes     | バイト列         |
| datetime  | 日時             |
| any       | 型不定（最小限） |
```

primitive は model field の `type:` でも、task の `params[].model` でも使用可能であるべき。

### 実装側の現状

`internal/resolve/validation.go` の以下2系統が**非対称**:

- `validateParams` は `modelOrPrimitiveExists(project, module, param.ModelName)` を呼んでおり、primitive を許容する
- `validateReturn` は `modelExists(project, ret.Model)` だけを呼んでおり、primitive を許容しない

つまり**task / join の returns で primitive (`str` / `int` / `bool` / `any` 等) を返す宣言**が
すべて `unresolved_model` エラーになる。

これはreturn shape表現として明らかな機能不足。
UC-001では task return は常に project内 model を返していたため顕在化しなかったが、
`returns.model: any` は spec gap #1, #2 の暫定対応でも「v1範囲では primitive any で逃がす」と決まっており、
実装がこれをサポートしていないこと自体が独立したバグである。

## 決定

### 1. task / join の returns は primitive を許容する

`task.returns.model` および `join.returns.model` は、project内 model のみならず primitive 予約語も指定可能とする。

具体的に以下を許容する:

- `str` / `int` / `float` / `bool` / `bytes` / `datetime` / `any` (V01-ADR-021 §3 のprimitive予約語7語)
- 同一moduleの bare model ID
- 別moduleの QualifiedID (full path)

これは `task.params[].model` の解決ロジックと対称となる。

### 2. validation の対称化

`internal/resolve/validation.go` の `validateReturn` を `validateParams` と同じく
`modelOrPrimitiveExists(project, module, raw)` ベースの判定に揃える。

### 3. 仕様側更新

本ADR受理後、以下のspecを更新する。

- **`docs/spec/nodes.md` task節 / `returns` オブジェクト節**: `model` に primitive を指定可能であることを明記
- **`docs/spec/diagnostics.md` の `unresolved_model` 説明**: primitive は対象外である旨を明記

## 理由

### 対称化の根拠

`params` と `returns` はどちらも task の I/O contract。
入力で primitive を許容し、出力で許容しないのは設計として非対称であり、その非対称を正当化する仕様根拠は存在しない。

### `any` の扱い

`any` の使用は最小限であるべき（V01-ADR-021 §3 の but）だが、現状 v1 model の表現力には限界があり、
spec gap #1 / #2（discriminated object / enum / nested list element model 不足）の暫定対応として
`any` を使うケースは正規の用法である。

v1.1（M15）で enum / discriminated object / inline struct を導入することで `any` の使用は減るが、
それでも「ResolvedProject 内部shape を MCP公開contractに出さない」等、
意図的に `any` を使う場面は v1.1 以降も残る。
よって `any` を含む primitive を return で許容することは中長期的にも必要。

### specを更新する根拠

V01-ADR-021 §3 は primitive の存在を定義しているが、
「primitive を return.model に書けるか」までは明示していない。
spec-first運用（V01-ADR-050）に従い、specで明示する。

### 却下した代替案

#### 代替案A: returnに primitive を許容しない方針を維持し、wrapper struct を強制する

- 例: `any` を返したい場合は `{value: any}` のstructを定義し、それをreturnする
- ボイラープレート爆発。MCP toolごとに wrapper struct を量産することになる
- params との非対称な設計を維持する積極的理由がない

#### 代替案B: `any` だけ特別扱いして許容、他のprimitiveは却下

- V01-ADR-021が定めたprimitive予約語の中で一部だけ扱いを変える根拠がない
- params / returns の対称化という単純で一貫した方針のほうが整合する

## 影響

### 既存実装への影響

- `internal/resolve/validation.go` の `validateReturn` 修正
- 関連テスト（`internal/resolve/validation_test.go` 等）に return primitive 許容ケースを追加

### 既存UCへの影響

- UC-001: 影響なし（task return がすべて project内 model のため）
- UC-002: 本修正は UC-002 Phase A render の前提条件を満たすが、UC-002 Phase A YAML は
  M15（data layer expressiveness v1.1）完了後に enum 等を使った形で再構築する方針のため、
  本ADR完了直後のUC-002 Phase A 直接render検証は M14a のスコープに含めない

### 型互換性ルールへの影響

本 ADR で task / join return が primitive を許容するようになった結果、
`any` を含む primitive type が wiring source として現れるケースが発生する。
ただし、wiring 時の型互換性ルール（`any` の代入互換挙動を含む）は本 ADR では規定しない。
**型互換性ルールは V01-ADR-060 で別途規定する。**

### v1.0.0-spec タグへの影響

V01-ADR-058 と同様、`v1.0.0-spec` 凍結条件のうち「`go test ./...` パス」は
M14a で本修正を入れる際に return primitive 許容テストを追加した状態で再度満たす必要がある。

V01-ADR-050 §7 / V01-ADR-057 Non-goals が禁じる「v1範囲のspec / ADRの遡及修正」には該当しない。
本修正は spec が認める表現を実装が拒否していたバグの修正である。

`v1.0.1-spec` タグは V01-ADR-058 / 本ADR の修正を含むpatch releaseとして発行する。

### M14への影響

- M14a で V01-ADR-058 (B1: subnode scope) と本ADR (B2: return primitive) を一括対応する
- M14（self-hosting）Phase A は M14a + M15 完了後に再構築する

## Evidence
- commit: tbd
- impl commit: tbd
- 参考: V01-ADR-021 §3 primitive予約語、`internal/resolve/validation.go` `validateParams` の実装
