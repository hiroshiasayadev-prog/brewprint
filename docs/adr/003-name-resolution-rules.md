# 003: 名前解決ルール

- **status**: accepted
- **date**: 2026-04-17
- **depends on**: ADR 002

## 背景

同モジュール内のノード参照を毎回フルパスで書くと冗長になる。
相対参照（`./procedure.login`）を許すか、ID直書きで解決するかの2択があった。

## 決定

同モジュール内はID直書きで解決する。モジュールを跨ぐ場合はフルパスを要求する。

```yaml
# 同モジュール内（auth/dag.yaml内での参照）
edges:
  - from: login        # auth.procedure.loginに解決される
    to: token

# モジュール跨ぎ
edges:
  - from: auth.procedure.login
    to: analysis.state.session
    kind: trigger
```

## 理由

Rustのmod内・Pythonの同パッケージimportと同じ感覚で自然。
言語仕様として一般的であり、学習コストが低い。

却下した代替案：
- 相対参照（`./procedure.login`） → パス文字列の処理が複雑になる割にメリットが薄い
- 常にフルパス → 同モジュール内の記述が冗長になる

## 影響

- GoのAST構築時に、エッジのfromがフルパスかどうかで名前解決ルートを分岐する
- バリデーションはASTビルド時に解決失敗をエラーとして返す
