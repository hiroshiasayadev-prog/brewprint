---
scope: docs/spec/mcp/overview.md
status: draft
last_updated: 2026-04-30
summary: >
  MCP spec全体の概要と設計原則、およびtoolの使い分けを定義する。
  各toolへの導線とLLM向けの選択指針を提供する。
depends_on:
  - docs/adr/047-go-semantic-model-query-layer-boundary.md
  - docs/adr/048-resolved-project-index-strategy.md
  - docs/adr/049-mcp-query-reference-vocabulary.md
  - docs/adr/054-mcp-query-coverage-for-design-conversation.md
  - docs/adr/055-mcp-reference-tree-traversal.md
  - docs/adr/056-mcp-analyze-impact-tool-design.md
---

# MCP仕様 Overview

## 1. Scope

このspecは、brewprintがMCP経由でLLMへ提供するquery toolの**外部I/O契約**を定義する。

対象:

- MCP tool名
- tool input schema
- tool output schema
- 共通ID表現
- 共通レスポンス語彙
- reference表現
- diagnostic / error表現
- LLMが各toolをどう使うべきかの意図

対象外:

- Go package名 / struct名 / interface名
- Raw YAML decode用struct
- `ResolvedProject` 内部のmap/index具体実装
- RendererのMermaid / HTML / Markdown出力仕様
- MCP transport実装の詳細
- transitive dependency graphの事前構築方式

本specは、ADR-047 / ADR-048で定義されたGo実装境界、およびADR-049で定義されたreference語彙統一を前提に、`QueryService` が外部へ返す情報の形を固定する。

---

## 2. Design principles

### 2.1 MCPはRaw YAML ASTを公開しない

brewprint MCPは、Raw YAML ASTをLLMへ公開するためのAPIではない。

MCP toolは常に、semantic buildを通過した `ResolvedProject` 上の情報を返す。

```text
YAML files
  ↓ load / classify / decode
Raw YAML structs
  ↓ validate / name resolution / derived model build / index build
ResolvedProject
  ↓
QueryService
  ↓
MCP response
```

MCP response内の参照は、原則として名前解決済みのIDを使う。
Raw YAMLに書かれた未解決文字列は、diagnosticやsource表示のために必要な場合のみ補助情報として返す。

### 2.2 Python inspect風の語彙を採用する

LLMが既に学習している一般的な introspection 語彙に寄せるため、MCP responseはPythonの `inspect` に近い操作感を持つ。

| 語彙 | 意味 |
|---|---|
| `signature` | 対象objectの外形。params / returns / fields / endpoint等 |
| `doc` | YAMLの `note` に由来する自然言語説明。semantic contractだが機械検証済み構造ではない |
| `source` | 定義元file / line / column等 |
| `members` | 対象objectが内包する要素。sub task / fields / transitions等 |
| `references` | 対象objectが参照する、または対象objectを参照する関係 |
| `diagnostics` | warning / error / hint 等の診断情報 |

ただし、brewprint MCPはPython AST互換APIではない。
公開対象はsyntax treeではなく、`ResolvedProject` 上のsemantic objectである。

### 2.3 dependenciesではなくreferencesを中心語彙にする

MCP responseでは、依存・参照・逆参照をまとめて `references` と呼ぶ。

理由:

- `dependency` は「ビルド依存」「実行依存」「型依存」などに意味が寄りやすい
- brewprintでは `reads` / `writes` / `transition.action` / `model field type` / `scenario step` など、依存というより参照として読むべき関係が多い
- Python / IDE文脈の `references` に近い語彙の方がLLMが解釈しやすい

ADR-049により、外部MCP tool名は `get_references`、内部QueryService method名は `GetReferences` に統一する。
`get_deps` / `GetDeps` は採用しない。

### 2.4 構造情報とdocを分離する

MCP responseは、機械的に確定した構造情報と、`note` 由来の自然言語説明を分けて返す。

```json
{
  "signature": {
    "params": [
      { "name": "credentials", "model": "auth.model.credential" }
    ]
  },
  "doc": "認証情報を検証しトークンを発行する"
}
```

`doc` はLLMへのsemantic contractとして重要だが、機械検証済みの事実として扱ってはならない。

### 2.5 v1のreferencesはdirectのみ

ADR-048 / ADR-049に従い、MCP v1では完全なtransitive dependency graphを事前構築しない。

そのため `get_references` は、初期仕様では**直接referenceのみ**を返す。

transitive closure / depth指定 / dependency graph cacheは、QueryService vertical sliceで実需が出た時点で別途拡張する。

### 2.6 設計対話 coverage を拡張判断基準にする

MCP / QueryService は、単なる実装補助APIではなく、DAG / State Diagram / Sequence Diagram / ER / API Table / Wireframe などの図・viewを見ながらLLMと設計対話するためのquery layerである。

そのため、renderされた図やviewに現れる主要semantic objectは、原則としてMCPからquery可能にする。

対象例:

- task / model / store / state / event / actor
- model field
- transition
- sequence scenario view
- API Table view
- ER Diagram view
- implicit asset
- file-local sub task / branch / fork / join
- flow entry / flow wiring
- source file

すべてをMCP v1で即時実装する必要はない。ただし、今後のMCP拡張では「そのobjectが図やview上で利用者に見えており、会話対象になりうるか」を優先判断基準とする。

MCP responseは引き続きRaw YAML ASTを公開せず、ResolvedProject上のsemantic object queryとして返す。source snippet取得が必要な場合も、semantic objectに対応するsource補助情報として扱う。

> 由来: ADR-054 §決定

---

## 3. Tool overview

MCP v1のquery toolは以下の8つとする。

| tool | 目的 | 主な利用場面 |
|---|---|---|
| [`list_objects`](tools/list-objects.md) | project内のsemantic object一覧を取得する | 実装・設計対話の起点として対象objectを探す |
| [`get_signature`](tools/get-signature.md) | object単体の外形を取得する | 実装前にtask/model/store等の型・I/Oを確認する |
| [`get_source`](tools/get-source.md) | semantic objectに対応するYAML snippetを取得する | 設計対話中に定義元YAMLを確認する |
| [`get_references`](tools/get-references.md) | objectの直接referenceを取得する | 影響範囲・依存・逆参照を確認する |
| [`get_reference_tree`](tools/get-reference-tree.md) | objectからdepth制限つきでreference graphを辿る | 変更影響範囲や周辺objectをN hopで確認する |
| [`analyze_impact`](tools/analyze-impact.md) | change kindを踏まえた意味づけ済み影響分析を取得する | 設計変更相談で「何が壊れるか / どう直すか」を判断する |
| [`inspect`](tools/inspect.md) | object kind別に実装判断用の文脈を取得する | Claude Code等が実装・修正時に読む |
| [`list_endpoints`](tools/list-endpoints.md) | API Table viewに基づくendpoint一覧を取得する | API実装・ルーティング確認 |

---

## 4. Tool selection guidance for LLM

LLMは以下の使い分けを基本とする。

| 状況 | 使うtool |
|---|---|
| 対象nodeのI/Oだけ確認したい | [`get_signature`](tools/get-signature.md) |
| 対象objectの定義元YAML snippetを確認したい | [`get_source`](tools/get-source.md) |
| 何に依存しているか / 何から参照されているか確認したい | [`get_references`](tools/get-references.md) |
| 変更影響範囲や周辺objectをN hopで確認したい | [`get_reference_tree`](tools/get-reference-tree.md) |
| 設計変更（rename / remove / 型変更等）の影響と直し方を判断したい | [`analyze_impact`](tools/analyze-impact.md) |
| 実装・修正・レビューのために周辺文脈が必要 | [`inspect`](tools/inspect.md) |
| API route一覧が必要 | [`list_endpoints`](tools/list-endpoints.md) |

原則:

- 実装前にはまず `inspect` を使う
- 小さな型確認だけなら `get_signature` を使う
- 直接参照確認では `get_references(direction="in")` または `both` を使う
- N hopの影響範囲確認では `get_reference_tree` を使い、`direction` と `depth` を明示する
- 設計変更相談では `analyze_impact` を使い、`change.kind` を明示する。 raw な reference 探索が必要な場合のみ `get_reference_tree` に降りる
- Raw YAMLを直接読む前に、まず `get_source` でsemantic objectに対応するsnippetを確認する

---
