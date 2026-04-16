# 001: ノード型分け

- **status**: accepted
- **date**: 2026-04-17

## 背景

ASTのノード表現として、union型 + kindフィールドによるswitch処理と、型ごとに別structに分ける方式の2択があった。

## 決定

ノードは型ごとに別structに分ける。

```go
type StructNode struct { ... }
type ScalarNode struct { ... }
type ProcedureInlineNode struct { ... }
type ProcedureRefNode struct { ... }
```

## 理由

MCPで`get_signature`を返すとき、union型でkindを見てswitchするより、型そのものがシグネチャになっている方がLLMが推論しやすい。`StructNode`が来たら`fields`がある、それだけで確定する。

validation実装コストは度外視し、**LLMが楽できる仕様**を優先する方針とした。

却下した代替案：
- union型 + kindフィールド → switch処理が必要になりLLMの推論コストが上がる

## 影響

- GoのAST定義はノード種別ごとにstructを用意する
- MCPの`get_signature`はstructをそのままJSONで返せばよい
