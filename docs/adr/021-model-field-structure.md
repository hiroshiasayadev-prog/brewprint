# 021: modelのフィールド構造確定（scalar廃止・primitive予約語・pk/fk・list/dict）

- **status**: accepted
- **date**: 2026-04-20
- **supersedes**: ADR-008（`fields[].comment` → `fields[].note` への改名のみ）

## 背景

`spec/nodes.md` に `⚠️ 未定` として残されていた以下の設計が未確定だった。

1. `model.kind: scalar` のYAML構造（フィールド名・必要性）
2. `model.kind: list` / `dict` のフィールド構造
3. `store.kind: db` が独自のフィールド定義を持つかどうか
4. `actor` ノードの固有フィールドの要否
5. `fields[].comment` という語がbrewprint内の `note` 慣習と不一致であること

## 決定

### 1. `fields[].comment` → `fields[].note` に改名

ADR-008では `fields[].comment` を採用していたが、brewprintの他フィールド（`task.note` / `store.note` 等）はすべて `note` で統一されている。一貫性を優先し `note` に改名する。

意味・役割は変わらない：人間向けdocstring兼LLMへのsemantic contract（ADR-008の定義を継承）。

### 2. `model.kind: scalar` を廃止

scalar modelは定義しない。primitive literalを `fields[].type`・`param.model`・`element`・`value` に直接書く。

```yaml
# NG: scalar modelを定義する
- id: user_id
  type: model
  kind: scalar
  base: str

# OK: struct fieldでprimitive literalを直接使う
fields:
  - name: user_id
    type: str
    note: "ユーザーID"
```

型エイリアス（`user_id = str`）が必要な意味付けは `note` で担う。

### 3. primitive予約語セット

以下の7語をprimitive予約語とする。model IDとして定義不可。

| primitive | 意味 |
|-----------|------|
| `str` | 文字列 |
| `int` | 整数 |
| `float` | 浮動小数点数 |
| `bool` | 真偽値 |
| `bytes` | バイト列 |
| `datetime` | 日時 |
| `any` | 型不定（使用は最小限に） |

`datetime` はDB固有フィールド（`created_at` / `updated_at`）を model に書く場合の可読性のために追加。

### 4. `model.kind: list` のフィールド構造

```yaml
- id: item_list
  type: model
  kind: list
  element: item      # model ID または primitive literal
```

| フィールド | 必須 | 内容 |
|-----------|------|------|
| `element` | ✓ | 要素の型。model ID または primitive literal |

`element: str`（文字列リスト）・`element: item`（model参照）どちらも有効。

### 5. `model.kind: dict` のフィールド構造

```yaml
- id: config_map
  type: model
  kind: dict
  value: config      # model ID または primitive literal
  # key: 定義しない → 常にstr
```

| フィールド | 必須 | 内容 |
|-----------|------|------|
| `value` | ✓ | 値の型。model ID または primitive literal |

`key` フィールドは定義しない。dictのkeyは常に `str` とする。

### 6. `model.kind: struct` のfield: `pk` / `fk` フラグ

```yaml
- id: user
  type: model
  kind: struct
  fields:
    - name: id
      type: str
      pk: true
      note: "ユーザーID（PK）"
    - name: role_id
      type: str
      fk: role.id
      note: "ロールID（FK → role.id）"
    - name: address
      type: address          # model ID参照・fkなし → JSON埋め込み扱い
      note: "配送先住所（DBではJSONカラム）"
    - name: tags
      type: tag_list         # list kindのmodel → DBではvariant/JSONカラム
      note: "タグリスト（DBではJSONカラム）"
    - name: created_at
      type: datetime
      note: "レコード作成日時（DB管理）"
```

#### `pk` フラグ

| フィールド | 必須 | 型 | 内容 |
|-----------|------|-----|------|
| `pk` | 任意 | bool | `true` でPKカラム。1 struct内に1つ |

ER図生成・DB schema生成の根拠としてClaude Codeが使用する。

#### `fk` フィールド

| フィールド | 必須 | 型 | 内容 |
|-----------|------|-----|------|
| `fk` | 任意 | `<model-id>.<field-name>` | 参照先のmodel IDとフィールド名 |

`fk: role.id` は「role modelのidフィールドを参照するFK」を意味する。

#### `fk` なしのmodel型参照 → JSON埋め込み

`type: address`（`fk:` なし）の場合、DBではJSONカラムとして埋め込む想定。brewprintはこの区別をYAML上で表現するが、実際のカラム型（`JSONB` / `TEXT` 等）の選択は実装依存でbrewprintのスコープ外。

#### list / dict kindのmodel型 → variantカラム

`type: tag_list`（`kind: list` のmodel）のようにlist/dict kindのmodelをfieldの型に使う場合、DBスキーマではvariant/JSONカラムになる。これは実装上の帰結であり、YAML上に追加フラグは不要。

### 7. `store.kind: db` は `of:` 参照のみ

storeは独自のフィールド定義を持たない。`of: <model-id>` でmodelを参照するのみ。

ER図は `store.of` → model → fields の参照を辿って列定義を描画する。DB固有フィールド（`created_at` 等）はmodel側に書く。

```yaml
- id: user_db
  type: store
  kind: db
  of: user
  note: "ユーザーテーブル"
```

### 8. `actor` の固有フィールドなし確定

`actor` は `id` / `type` / `note` のみ。`external:` / `role:` 等の固有フィールドは追加しない。Mermaid renderで人間/システムを描き分けたい場合は別途ADRで追加する。

## 理由

### scalar廃止

型エイリアスの意味付けは `note` が担える。brewprint全体で「機械検証できないセマンティクスはnoteで」という原則（ADR-008）を一貫させる。scalar kindが残ると「いつ定義するか」の判断コストが生じる。

### datetime追加

`str` で代替できるが、DBスキーマ生成でClaude Codeが判断に迷うケースが生じる（TEXT vs TIMESTAMP）。`datetime` を明示することで実装精度が上がる。

### `pk` をmodelに書く

PKはドメインエンティティの識別子であり、modelの関心範囲。store側に書くと `store → model → fields` の2ホップが必要になり、Claude Codeのスキーマ生成追跡コストが増す。

### `fk` の明示記法

`type: role`（model参照）だけではFK参照かJSON埋め込みかが判定できない。brewprintは「設計意図を明示する」方針（ADR-009のインライン定義禁止と同じ思想）のため、FK意図を `fk:` で明示する。

### dictのkeyをstrに固定

実用上dictのkeyはほぼ常に文字列。`int`キーが必要になった場合はその時点でADRを追加する。

### actorの軽量設計

ADR-004の「軽量な定義でよい」方針を踏襲。noteで意味付けできる情報に固有フィールドを追加しない。

## 影響

- `spec/nodes.md` の `fields[].comment` をすべて `fields[].note` に更新する
- `spec/nodes.md` の `model` セクションの `kindごとの追加フィールド`（⚠️未定）を本ADRの内容で確定する
- `spec/nodes.md` の `actor` セクションの ⚠️未定 を確定済みに更新する
- ADR-008の `fields[].comment` 言及は本ADRにより `fields[].note` に改名済みと注記する

## Evidence
- commit: 66e3c1a
- impl commit: tbd
- 参考: 特になし
