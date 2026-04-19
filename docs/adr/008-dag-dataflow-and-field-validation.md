# 008: DAGデータフロー構造とfieldのvalidation層構造

- **status**: accepted
- **date**: 2026-04-17

## 背景

taskのI/Oとして何を参照するか、およびassetのフィールド定義においてどのレベルのvalidationを行うかを決定する必要があった。

## 決定

### 1. assetはDAG上の共有ノード

assetはDAGにおいて独立したノードとして定義され、複数のtaskから参照される。
taskはassetノードのIDを `input` / `outputs` に指定することで接続を宣言する。

```
[task: login] --produces--> (auth_token) --consumed_by--> [task: dashboard]
```

- 同一のassetノードを複数のtaskが参照できる
- taskがasset IDを参照した時点で接続の整合性が保証される（同一IDであることが型一致の根拠）
- structural check（フィールド構造の機械的照合）は不要

### 2. fieldのvalidation層構造

assetのフィールドは以下の3要素で構成する：

```yaml
fields:
  - type: str        # 型（primitive or asset ID）
    name: token      # フィールド名
    comment: ログイントークン  # セマンティクスの記述
```

表示上のイメージ：
```
[str] token     ログイントークン
[User] user     ログインユーザー
```

| 要素 | validation対象 | 役割 |
|------|--------------|------|
| `type` | ✓ 機械的validation | 存在チェック（primitive or 定義済みasset ID） |
| `name` | ✓ 機械的validation | struct内でのユニーク性 |
| `comment` | LLM semantic contract | 意味的整合性の根拠。人間にとってはdocstring、LLMにとってはsemantic validationの根拠 |

`comment` は機械的validationの対象外だが、MCPの `inspect` ツールがLLMによる意味的整合性チェックの根拠として使用する。

## 理由

**shared asset nodeパターン**：LuigiのTargetコンセプト（タスクの出力を独立したノードとして定義）およびDagsterのAsset Graphに倣った。taskが入出力を独立に宣言してエッジを導出する方式より、assetを共有ノードとして扱う方式の方がDAGの接続整合性をシンプルに保証できる。

**field構造**：
- `type + name` による機械的validationはPython PEP 484・TypeScript・Goの型アノテーション慣習に倣う
- `comment` をLLMのsemantic contractとして位置づけるのはbrewprint固有の設計。「型・名前では表現できないセマンティクス」を自然言語で記述し、LLMが意味的整合性を判断する根拠とする。人間向けdocstringとLLM向けcontractを兼ねる。
- フィールドの型まで機械的に厳密一致させるvalidationは、GoやOpenAPI等の実装ツールの責務であり、brewprintのスコープ外とする

参照した公知概念：
- Luigi Target / Dagster Asset Graph（shared asset nodeパターンの根拠）
- Python PEP 484・TypeScript type annotation（field型宣言の慣習）
- Python docstring・JSDoc（commentのdocstring的役割の根拠）

## 影響

- `spec/nodes.md` のfield定義は `type / name / comment` の3要素とする
- MCPの `inspect` ツールはcommentをLLMに渡してsemantic validationを行う
- assetのIDは参照の一致をもって型チェックの根拠とするため、alias・rename系の機能を導入する場合は別途ADRで決定する

## Evidence
- commit: f911107
- impl commit: tbd
- 参考: Luigi Target概念・Dagster Asset Graph、Python PEP 484・TypeScript型アノテーション、Python docstring慣習参考
