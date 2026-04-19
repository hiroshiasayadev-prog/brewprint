# 006: ノード種別の命名変更

- **status**: accepted
- **date**: 2026-04-17

## 背景

初期仕様では `procedure / artifact / state / actor` という種別名を使っていた。
命名の一般性・他ツールとの整合性・LLMの推論しやすさを検討した結果、より広く使われている用語に変更する。

## 決定

| 旧名 | 新名 | 参照した公知概念 |
|------|------|----------------|
| `procedure` | `task` | Apache Airflow・Prefect・Luigi・Celeryにおける `task`（処理の単位の標準語彙） |
| `artifact` | `asset` | Dagster の Software-Defined Assets における `asset`（DAGを流れるデータ資産） |
| `state` | `store` | Redux・Vuex・Pinia・Zustand における `store`（実行時にデータを保持する実体） |
| `actor` | `actor` | UML標準の Actor（変更なし） |

## 理由

**`task`への変更**：`procedure` はストアドプロシージャ・手続き型プログラミングの文脈に寄りすぎていた。DAG/ワークフローツールで最も広く使われる `task` の方がLLM・人間ともに認識しやすい。

**`asset`への変更**：`artifact` はCI/CDの「ビルド成果物」の意味で定着しており、「DAGを流れるデータ」という意図と微妙にずれていた。DagsterのSoftware-Defined Assetsが同概念に `asset` を採用しており、これに倣う。

**`store`への変更**：`state` はstate machineの「状態ノード」と名前衝突する。brewprintはstate diagramも扱うため特に問題。「実行時にデータを保持する実体」という意味では、フロントエンド状態管理（Redux等）で広く定着している `store` の方が正確かつ一般的。

## 影響

- `spec/overview.md` のノード種別テーブルを更新する
- 以降のすべてのspec・uc・ADRは新名を使う
- GoのAST struct定義は `TaskNode / AssetNode / StoreNode / ActorNode` とする

## Evidence
- commit: f911107
- impl commit: tbd
- 参考: Apache Airflow/Prefect/LuigiのTask慣習、DagsterのSoftware-Defined Assets、ReduxのStore概念参考
