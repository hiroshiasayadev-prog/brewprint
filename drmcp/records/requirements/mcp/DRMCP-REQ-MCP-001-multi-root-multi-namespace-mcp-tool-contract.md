# DRMCP-REQ-MCP-001: multi-root / multi-namespace 環境における MCP tool contract の再定義

- **id**: DRMCP-REQ-MCP-001
- **status**: captured
- **date**: 2026-06-10
- **source_refs**:
  - V01-REQ-MCP-033
- **work_items**:

## Requirement

`design-records-mcp` に multi-root / multi-namespace index を実装した（V01-WORK-MCP-030）後、既存 MCP tool contract の各 operation が複数 app namespace 環境でどう振る舞うべきかが未定義のままである。本 REQ は、影響範囲を axis ごとに調査し、contract の再定義または spec 補完の提案を求める。

## Investigation Axes

### 1. namespace filter / query scope

`list_records`、`get_records`、`validate_records` 等のクエリ系ツールは現状、全 namespace を混在して返す。multi-root 環境でどの単位でスコープを区切るか（全 namespace 横断 vs. caller 指定 namespace filter）の contract が未定義。

調査範囲:

- 既存 tool spec（`SPEC-design-records-mcp-tools`）が single namespace 前提で書かれている箇所の特定
- namespace filter / query scope パラメータの要否・形式・デフォルト動作の定義提案

### 2. 完全修飾 ID-as-ref における namespace 判断 contract

`resolve_reference` は現状、`idx.RecordsEntries` を iterate して known prefix を照合する実装を持つ。しかしこの動作は spec に明記されておらず未定義状態である。`DRMCP-WORK-MCP-001` や `PRODUCT-REQ-SPEC-001` のような cross-namespace ID-as-ref が `resolve_reference` でどう扱われるべきか、spec レベルで定義が必要。

調査範囲:

- 現行 `SPEC-design-records-mcp-tools` の `resolve_reference` 仕様と、multi-namespace 前提での不整合箇所の特定
- 完全修飾 ID-as-ref の acceptance criteria と resolver contract の補完提案

### 3. cross-namespace validation scope と record identity / relation の定義

`validate_records` が cross-namespace の参照（例: `V01-WORK-MCP-030` が `DRMCP-REQ-MCP-001` を source_requirement として参照する）を valid として扱うかどうかの定義がない。また multi-root index における record identity（同一 normalized ID が複数 namespace に存在する場合の扱い）も未定義。

調査範囲:

- 既存 spec（overview / tools / schema）で record identity・relation validation に関わる記述の特定
- cross-namespace relation の validation scope 定義提案（何を許容し、何を error / warning とするか）

## Required Outcome

- 上記 3 axis それぞれについて、既存 spec の不整合・空白箇所を列挙した調査結果を作成する
- 各 axis に対する contract 定義案または spec 補完案を提示する
- 調査結果および提案を WORK / TASK として追跡可能にする

## Explicitly Excluded Scope

- multi-root index 自体の実装変更（本 REQ は調査・定義フェーズ）
- suggest_next_record の廃止処理（`*-new` placeholder で機能的に代替済み。別途対応）
- authoring routing: semantic ref spec の namespace 定義で対応済みのため本 REQ の対象外
