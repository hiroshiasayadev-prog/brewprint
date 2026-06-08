# V01-ADR-030: YAMLファイル種別の自己宣言（`as:` フィールド）

- **status**: accepted
- **date**: 2026-04-22
- **migrated_to_spec**: 2026-04-29

> このADRの現行仕様詳細は [docs/spec/file-types.md](../spec/file-types.md) §1〜§3 を参照。

## 背景

brewprintのYAMLファイルには2種類が存在する。

- **ノード定義ファイル** — `task/*.yaml`, `model/*.yaml`, `state.yaml` 等。トップレベルに `nodes:` リストを持ち、各ノードが `type:` フィールドで種別を宣言する
- **View定義ファイル** — API Table YAML（V01-ADR-028）等。ノードリストを持たず、view固有のフィールド（`http_root_path:`, `modules:` 等）がトップレベルに並ぶ

ノード定義ファイルは各ノードの `type:` フィールドから種別が判定できる。
しかしview定義ファイルには `type:` も `nodes:` も存在しないため、パーサーがノード定義ファイルと区別できない。

ファイル名・ディレクトリ構造に依存した判定はV01-ADR-002の方針とも相性が悪く、
`common_actors.yaml` / `auth_api.yaml` のような任意のファイル名を許容するためにも、
ファイル自身が種別を宣言する仕組みが必要と判断した。

## 決定

**view定義ファイルのみ**、トップレベルに `as:` フィールドを必須とする（仕様詳細: [file-types.md](../spec/file-types.md) §1〜§2）。

```yaml
as: api_table
id: auth_api
http_root_path: /api
modules:
  - module: app.auth
    include_submodules: true
```

ノード定義ファイルには `as:` を置かない。パーサーはトップレベルに `as:` が存在するかどうかでノード定義 vs view定義を判別する。

### `as:` の値（起票時点スナップショット）

| 値 | 対応するview |
|----|-------------|
| `api_table` | API Table view（V01-ADR-028） |
| `sequence_diagram` | Sequence Diagramシナリオ（V01-ADR-032） |

> 現行の `as:` 値一覧は [file-types.md](../spec/file-types.md) §3 を参照。view 種別が追加された際は spec 側を更新する。

## 理由

- ノード定義ファイルは各ノードの `type:` で種別が判定できるため、ファイルレベルの宣言は不要
- view定義ファイルはノードリストを持たないため、`type:` による判定ができない。`as:` による明示宣言が唯一の判定手段
- ファイル名・ディレクトリ名への依存をなくすことで、`common_actors.yaml` のような任意の命名を許容できる

却下した代替案：
- すべてのYAMLファイルに `as:` を必須にする → ノード定義ファイルでは冗長。`type:` で十分判定できる
- ディレクトリ名で種別を判定する → 命名の自由度が下がり、V01-ADR-002の方針とも相性が悪い

## 影響

- view定義ファイルのYAML schemaに `as:` フィールドを必須として追加する
- `spec/views/api-table.md` のYAML schemaに `as: api_table` を追加する
- Goパーサーはトップレベルの `as:` の有無でノード定義 / view定義を振り分ける
- 新たなview種別を追加する際は本ADRの `as:` 値一覧を更新する

## Evidence
- commit: ec5e496
- impl commit: tbd
- 参考: 特になし
