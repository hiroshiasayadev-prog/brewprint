# DRMCP-INV-MCP-001: multi-root / multi-namespace 環境における MCP tool contract 調査

- **id**: DRMCP-INV-MCP-001
- **status**: concluded
- **date**: 2026-06-10
- **trigger**: DRMCP-REQ-MCP-001
- **scope**: multi-root / multi-namespace 環境において既存 MCP tool contract に生じた空白・不整合箇所の特定、および contract 定義案の提示。対象ツールは `list_records`、`get_records`、`validate_records`、`resolve_reference`
- **non_scope**: multi-root index 自体の実装変更、spec / implementation ファイルの実際の変更（本 INV は調査フェーズ）、suggest_next_record の廃止処理、authoring routing（semantic ref spec の namespace 定義で対応済み）
- **source_refs**:
  - DRMCP-REQ-MCP-001
  - V01-REQ-MCP-033
- **follow_up_candidates**:
  - DRMCP-WORK-MCP-new

## 背景

V01-WORK-MCP-030 で Design Records MCP に multi-root / multi-namespace index を実装した。`*/records/` を glob auto-detect し、namespace prefix は親ディレクトリ名から `strings.ToUpper(dir) + "-"` で導出する。bpdsl / drmcp / product / v01 の 4 namespace、計 405 records を統合 index で管理している。

既存 MCP tool contract（`SPEC-design-records-mcp-tools`）は single namespace 前提で書かれていたため、multi-root 環境での各ツールの動作 contract が未定義のままとなった。本 INV はその空白・不整合を 3 axis で調査した結果を記録する。

## Axis 1: namespace filter / query scope

### Spec 空白・不整合箇所

| ファイル | セクション | 問題の種類 | 内容 |
|---|---|---|---|
| tools.md | list_records Request (L148–205) | 空白 | Request schema に namespace を絞る field がない。kind / status / id / id_range / order_by / order / limit のみ。multi-root 環境でのスコープ規定なし |
| tools.md | list_records Response (L206–230) | 空白 | mixed namespace を返す場合の ordering 規則に言及なし。public ID に namespace prefix を含む場合の mixed namespace ordering 例が無い |
| tools.md | get_records Request (L317–342) | 前提ズレ | exact lookup key のみと明記。「全 namespace を一つの index と見なして lookup する」明文が無い |
| tools.md | validate_records Request (L654–675) | 空白 | request 空時に「MVP index 対象の全 record を検証する」とだけ。「全 record = 全 namespace 横断」かどうかが未定義 |
| overview.md | multi-root スキャン (L138–148) | 空白 | 「異なる app namespace の record を同一クエリで取得・参照解決できる」とは書かれているが、ツール側 default scope（横断 vs 単一）と filter API の有無が言及されていない |
| tools.md | id_range 規則 (L185–204) | 不整合 | Mixed family / mixed domain は禁止だが mixed namespace（from と to の namespace prefix が異なる場合）の扱いが未定義 |

### Contract 定義案

- `list_records` / `validate_records` の Request に optional `namespace` field を追加する
  - 型: `string` または `array of string`（複数 namespace 横断クエリを許容）
  - 値: 末尾 `-` を含む namespace prefix 形式（例: `V01-`、`DRMCP-`）を canonical とする
  - default: 未指定 = 全 namespace 横断
  - 未知 namespace は `invalid_request` エラー
- `get_records` には追加しない。`ids[]` は完全修飾 public ID で exact lookup する設計と整合するため、namespace filter は意味を持たない。spec 側に「ids[] は全 namespace 横断 index に対して exact lookup される」旨を明記する
- `id_range` の mixed namespace 規則を追加: from / to の namespace prefix が異なる場合は `invalid_request`
- `list_records` Response の mixed namespace ordering 規則を明記: `order_by: id` は public ID（namespace prefix 込み）の ASCII lexical order とする（`DRMCP-` < `PRODUCT-` < `V01-`）

## Axis 2: 完全修飾 ID-as-ref における namespace 判断 contract

### Spec 空白・不整合箇所

| ファイル | セクション | 問題の種類 | 内容 |
|---|---|---|---|
| tools.md | resolve_reference Supported input table (L572–585) | 前提ズレ | 「MVP では namespace_prefix = V01- となる」と明記されており、複数 namespace 前提を排除した記述。実装は multi-root に対応済み |
| tools.md | resolve_reference Purpose (L556–558) | 空白 | multi-root index に対する prefix matching algorithm が無い |
| tools.md | resolve_reference Response (L587–637) | 空白 | 同一 normalized ID が複数 namespace に存在する場合の挙動（ambiguous_reference か別カテゴリか）が未定義。実装は public ID 全体を key にするため異 namespace 間では衝突しないが、spec 側にこの不変条件の言及が無い |
| schema.md | Public ID (L440–459) | 空白 | 「namespace prefix が異なる record の public ID は識別子として常に区別される（衝突しない）」という invariant が明文化されていない |

### 実装（resolver.go）と spec の差分

- 実装は登録済み namespace prefix の集合（`idx.RecordsEntries` 由来）に対する prefix match を行う。spec の表は `<namespace_prefix>` のような placeholder 表記で、何が valid namespace prefix なのかを規定しない
- 実装は backward-compat path として `idx.NamespacePrefix`（single-root mode）を fallback する（L41–45）。spec はこの 2 段階 lookup を記述しない
- 実装は未知 prefix（例: `FOO-ADR-001`）を `refKindUnsupported` として `status: "unsupported"` を返す。spec の「grammar に合わない ID form は unsupported」と一致するが、「known namespace に含まれない prefix も unsupported に分類される」が明文化されていない
- ambiguous 判定は `normalizeRecordID` が public ID 全体を upper-case 化するため、異 namespace 間の同一 bare ID は別 key となり ambiguous にならない。spec 側で invariant として未記載

### Contract 定義案

- resolve_reference Purpose / Request セクションに以下を追記する
  - **Known namespace set の定義**: 「index に登録された全 namespace prefix の集合を known namespace set とする。resolve_reference の record ID-as-ref 入力は、known namespace set のいずれかの prefix で始まり、ストリップ後の文字列が record kind ごとの bare ID grammar に合致した場合のみ supported とする」
  - **Unsupported になる条件の明示**: 「prefix が known namespace set に含まれない場合、または prefix ストリップ後の文字列が bare ID grammar に合致しない場合、status: "unsupported" を返す」
  - **Response invariant**: 「public ID は namespace prefix 込みの完全形であり、異 namespace 間で衝突しない。同一 namespace 内で重複した場合のみ ambiguous_reference を返す」
- tools.md Supported input table の「MVP では namespace_prefix = V01- となる」を削除し、「namespace_prefix は known namespace set から選ばれる」とする
- schema.md Public ID セクションに non-collision invariant を追加: 「異なる namespace_prefix を持つ record の public ID は常に区別される。したがって list_records / get_records / resolve_reference / validate_records の lookup key としての public ID は multi-root 環境でも一意である」

## Axis 3: cross-namespace validation scope と record identity / relation の定義

### Spec 空白・不整合箇所

| ファイル | セクション | 問題の種類 | 内容 |
|---|---|---|---|
| tools.md | validate_records Purpose (L640–652) | 前提ズレ | 「namespace_prefix 付き完全形。MVP では V01-ADR-* / ...」と single-namespace 前提で記述。cross-namespace 参照の許容/禁止が未定義 |
| tools.md | Diagnostic categories unresolved_workflow_relation / unresolved_source_ref | 空白 | 解決規則として resolve_reference と同一とあるが、「異 namespace の record を参照した場合 valid か」が未定義 |
| tools.md | duplicate_id 関連 | 空白 | multi-root index で「同一 normalized ID が複数 namespace に存在」する場合の扱いが明示なし（実装は発生しないが spec 側未記載） |
| schema.md | ID normalization model (L436–476) | 空白 | multi-root で record identity が namespace prefix 込みで一意になるという invariant が未記載 |
| schema.md | Discovery and index inclusion (L414–435) | 空白 | 複数 records_root を統合した index の identity contract が無い |
| overview.md | multi-root スキャン L144 | 空白 | 「cross-namespace relation は index が保持し、resolve / validate が解決できる」のみ。validation の valid/error/warning 判定基準が無い |

### Contract 定義案

- **Record identity invariant**（schema.md ID normalization model に追記）
  - 「multi-root 環境において、record の identity は public ID（namespace prefix 込み）で一意に決まる。namespace prefix が異なる record は、bare ID が同一でも別の record として扱う」
  - 「異 namespace 間で同一 bare ID が存在することは正常状態であり、いかなる diagnostic も出さない。duplicate_id diagnostic は同一 namespace 内の同一 public ID 重複に対してのみ発火する」
- **Cross-namespace relation validation scope**（tools.md validate_records Purpose に追記）
  - 既定方針: cross-namespace 参照は valid として扱う。理由: V01-WORK-MCP-030 から DRMCP-REQ-MCP-001 を参照するケースが現実に発生する正常 use case
  - 解決規則: resolve_reference の known namespace set lookup に従う。target が解決可能であれば valid
  - error 判定: target が known namespace set 内に存在しない → 既存 unresolved_workflow_relation / unresolved_source_ref (error)
  - 新カテゴリは追加しない。既存の unresolved_* / invalid_workflow_relation_target で cross-namespace の異常ケースをカバーできる
  - ただし invalid_workflow_relation_target の説明に「target が他 namespace に存在することそれ自体は invalid 条件ではない」を明記する
- **validate_records request scope の multi-root 拡張**
  - request 空時の default scope を「全 namespace 横断の全 record」と明文化する
  - Axis 1 で提案した `namespace` filter を validate_records でも受ける。ただし relation 解決時の lookup index は全 namespace 横断のままにする（指定 namespace の record から外 namespace を参照する relation が valid 判定できるよう）

## 後続 artifact 候補

### DRMCP-WORK-MCP-new: multi-root / multi-namespace MCP tool contract 補完

推奨 TASK 構成（WORK/TASK 分割案）:

| TASK | 内容 | 概算 | 依存 |
|---|---|---|---|
| T01 | schema.md: Public ID non-collision invariant と ID normalization model の multi-root identity 規則を追記 | 0.5d | — |
| T02 | overview.md: multi-root スキャン節に default scope と cross-namespace relation validation 方針を追記 | 0.5d | T01 |
| T03 | tools.md: resolve_reference を multi-root 対応に書き換え（known namespace set 概念、Supported input table 修正、ambiguous 不発生 invariant） | 1d | T01 |
| T04 | tools.md: list_records / validate_records Request に namespace filter を導入。get_records に scope 注記追加。id_range mixed namespace 規則を追加 | 1d | T01, T03 |
| T05 | tools.md: validate_records Purpose / Diagnostic categories に cross-namespace 参照の valid 扱いと既存 diagnostic への影響を明記 | 1d | T01, T03, T04 |
| T06 | spec 更新を反映した実装差分 review（resolver.go 等の挙動が新 spec 文言と一致するか確認）。追加変更が必要な場合は別 WORK を起票 | 0.5d | T05 |

全 T01–T05 が spec update のためレビューゲートあり。概算合計 4–5d。

T03〜T05 を 1 タスクに統合してレビューを 1 回にまとめる選択肢もある。ただし resolve_reference と validate_records は性質が異なるため、axis 別に分けることを推奨。
