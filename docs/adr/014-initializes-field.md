# 014: initializesフィールド設計

- **status**: accepted
- **date**: 2026-04-19

## 背景

ADR-013にて「foreachの結果をstoreの特定フィールドへ蓄積するパターン」をADR-014に委譲した。

議論の過程で、`initializes` はノード種別・foreachのフラグ・apply task内の責務のいずれでもなく、**main nodeのフィールド**として持つべきという結論に至った。

## 決定

### initializesはmain nodeのフィールド

```yaml
- id: process_report
  type: task
  main: true
  initializes:
    - name: cache
      model: result_cache
      note: "空のresult_cacheで初期化。params.configからdefault_ttlを設定する"
    - name: report
      model: report
      note: "report.itemsは空リスト、report.summaryは空文字で初期化"
  params:
    - name: config
      model: app_config
```

| フィールド | 必須 | 内容 |
|-----------|------|------|
| `name` | ✓ | このファイル内で参照するstore名 |
| `model` | ✓ | storeの型（modelのID参照） |
| `note` | 任意 | 初期値・初期化方法を自然言語で記述 |

### 初期値の記述方針

初期値の詳細（デフォルト値・paramsからの参照・dataclassのフィールド初期化）は `note` に自然言語で記述する。YAMLフィールドとして初期値をサポートしない。

### スコープ

`initializes` で宣言されたstoreは**ファイル内にprivate**。同ファイル内のsub taskおよびforeachから参照可能。外部からの参照不可（ADR-011のスコープ規則と同じ）。

## 理由

### main nodeのフィールドとして持つ理由

「関数内で使う変数は最初に宣言する」の精神をYAML構造として表現する。`initializes` をファイルの先頭（main node）に集約することで、「このファイルで使うstoreは全てここを見れば分かる」状態になる。

apply task内で副作用としてstoreを更新するパターン（return None + append）より、DAGの読み手にとって意図が明確になる。

### 初期値をnoteで書く理由

初期値をYAMLフィールドとして厳密にサポートすると、型システムの作り込みが必要になり仕様が際限なく複雑化する。brewprintの方針「図で表せないことはnoteで」（spec/overview.md）と一致するため、自然言語記述で割り切る。

却下した代替案：
- `initializes` をノード種別として独立させる → 「ファイル内で使うstoreの宣言」はノードではなくファイルの属性であるため却下
- `initializes` をforeachのフラグとする → foreach外でstoreを使い回すパターンが表現できないため却下
- apply task内でNone returnでstoreに蓄積する → DAGの読み手に意図が伝わらないため却下
- `value` フィールドで初期値を構造化する → 型システムの複雑化を招くため却下

## 影響

- `spec/overview.md` のノード種別表への追記は不要（`initializes` はノード種別ではなくフィールド）
- `spec/nodes.md`（未作成）にてmain nodeのフィールド定義を詳細化する際に `initializes` を含める
- storeへの参照・更新のedge表現はADR-015（ファイル内edge記述構造）に委譲

## Evidence
- commit: tbd
- impl commit: tbd
- 参考: 特になし
