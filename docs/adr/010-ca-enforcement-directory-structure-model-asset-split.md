# 010: CA強制・ディレクトリ構造・model/asset分離

- **status**: accepted
- **date**: 2026-04-19
- **supersedes**: ADR-002（dag.yaml/er.yamlのビュー別ファイル分け）, ADR-007（asset = 型定義の定義）

## 背景

YAMLの階層構造とディレクトリ構造の1:1対応設計を検討する中で、以下の問題が明確になった。

1. ADR-002が示した `dag.yaml` / `er.yaml` というビュー別ファイル分けは、型別サブディレクトリ設計と軸が異なる
2. ADR-007で定義した「asset = 型定義」は、DagsterのSoftware-Defined Assets（データ成果物の実体）を参照根拠としていたが、両者は別概念だった
3. 結果として `asset` に「型の定義」と「フロー上の存在」という2つの責務が混在していた

## 決定

### 1. brewprintはClean Architectureを構造的に強制する

brewprintを使うほど大規模なアプリケーションにはCAが必要、という前提に立つ。
「brewprintを使う = CAを採用した」という宣言とし、YAMLスキーマがCAの層構造を強要する。

自由度を元の言語に委ねると、Claude Codeが実装配置で迷う。構造的に強制することで
「どこに何を置くか」の判断を不要にし、自動実装の精度を上げることが目的。

### 2. ディレクトリ構造：1ファイル=1メインノード、型サブディレクトリ方式

> **ADR-011により改訂**：「1ファイル=1ノード」は「1ファイル=1メインノード（+ 複数のサブノード可）」に上書きされた。

```
modules/
  auth/
    model/              ← 型定義
      user.yaml
      token.yaml
      credential.yaml
    task/               ← 処理（assetノードの暗黙的宣言を含む）
      login.yaml
      logout.yaml
    store/              ← 実行時インスタンス
      user_db.yaml
    state.yaml          ← state machine宣言（必要な場合のみ）
    sequence.yaml       ← sequenceシナリオ宣言（必要な場合のみ）
```

- 1ファイル = 1ノード
- YAMLの階層 = ディレクトリ構造 = 実装ファイル構造が1:1対応
- これが「設計と実装の乖離を構造的に防ぐ」仕組みとなる

### 3. dag.yaml / er.yaml を廃止：ビューはノード定義から自動導出

| ビュー | 導出方法 |
|--------|---------|
| DAG | `task` の `params`/`returns` 参照関係から自動導出 |
| ER | `store.of` / `model.fields` の参照関係から自動導出 |
| Class | `endpoint: true` なtaskを自動収集（ADR-005） |
| State | `state.yaml` に明示宣言（自動導出不可） |
| Sequence | `sequence.yaml` に明示宣言（自動導出不可） |

ADR-002が示したビュー別ファイル（`dag.yaml` / `er.yaml`）は廃止。
名前空間ルール（フォルダ = モジュール）はADR-002から継続して有効。

### 4. model / asset の責務分離

`asset` に混在していた2つの責務を分離する。

| 概念 | ファイル | DAGに登場 | 本質 |
|------|---------|----------|------|
| `model` | ✅ `model/*.yaml` | ❌ | 型定義（どんな形か） |
| `asset` | ❌ taskから導出 | ✅ | フロー上の存在（役割） |

**model**：型定義に徹する。DAGには登場しない。`model/` サブディレクトリに1ファイル=1定義で置く。

**asset**：`task` の `returns` 宣言から暗黙的に生まれる。独立ファイルは持たない。
assetの意味はtaskによって生産されて初めて生まれるため、task定義の外に独立する情報がない。

```yaml
# task/login.yaml
- id: login
  type: task
  params:
    - name: credentials
      model: credential      # model IDを参照
  returns:
    name: auth_token
    model: token             # ← このassetノードがDAG上に暗黙的に生える
```

### 5. state machineのstateはstoreと別概念

state machineの `state`（「システムが今いる場所」）は、
`store`（「データの実体」）とは本質的に別物。新規ノード種別として扱う。
詳細はADR-011に委譲。

## 理由

**CA強制**：brewprint使用対象は「CAが必要な規模のアプリ」に絞る。
自由度を残すとClaude Codeが実装配置で迷い、自動実装の品質が下がる。
構造的に強制することで「どこに何を置くか」の判断を不要にする。

**1ファイル=1ノード**：実装ファイルとの1:1対応を成立させるための最小単位。

**ビュー自動導出**：Dagsterの思想「依存関係をノードと一緒に追跡することでグラフを自動推論する」を参照。
ただしDagsterの `asset`（データ成果物の実体）とbrewprintの `model`（型定義）は別概念であり、
思想のみ借用する。

**model/asset分離**：
- `asset`（フロー上の存在）はtaskの `returns` から意味が生まれる。task定義があれば独立ファイルは冗長
- `model`（型定義）はDAGと独立して存在すべき情報。ER / class diagramにも影響する
- 分離することでMCPの問いが一意になる：「形は？→ model」「どこを流れる？→ task参照」

却下した代替案：
- `asset/` ディレクトリを独立で持つ → taskから導出できる情報の二重管理になる
- ビュー別ファイル（dag.yaml / er.yaml）継続 → 型別サブディレクトリと軸が異なり混乱する
- CA非強制（自由度を言語に委ねる） → Claude Codeが実装配置で迷う。brewprintの価値が薄れる

## 影響

- ADR-002の `dag.yaml` / `er.yaml` 例示は本ADRにより廃止。名前空間ルール（フォルダ=モジュール）は継続有効
- ADR-007の「asset = 型定義」は `model` に移管。`store` の定義（kind: db / session / collection / context）は継続有効
- ADR-009の `params` / `returns` は `model` ID を参照する形に更新（旧: asset ID）
- `spec/nodes.md`（未作成）にて `model` / `task` のフィールド定義を詳細化する
- `spec/overview.md` のノード種別テーブルを更新する必要がある ✅
- **ADR-011により「1ファイル=1ノード」は「1ファイル=1メインノード」に改訂。その他の決定は継続有効**

## Evidence
- commit: 09b7d25
- impl commit: tbd
- 参考: Clean Architecture層構造、Dagsterの依存関係自動推論思想参考
