# 022: DAGのノード形状とエッジ種別の根拠

- **status**: accepted
- **date**: 2026-04-20

## 背景

`spec/views/dag.md` にてDAGのMermaid renderルールを定義するにあたり、
ノード形状・エッジ種別の選択根拠を公知技術に基づいて明示する必要があった。

brewprintの設計が既存の標準・公知技術の組み合わせであることを記録するために本ADRを起票する。

## 決定

### ノード形状

| ノード種別 | 形状 | Mermaid記法 | 根拠 |
|-----------|------|------------|------|
| `task` | 矩形 | `[label]` | ISO 5807:1985 フローチャート Process記号 / UML Activity Diagram Action Node（OMG UML 2.x） |
| `asset` | スタジアム | `([label])` | UML Activity Diagram Object Node（OMG UML 2.x） |
| `store` | シリンダー | `[(label)]` | ISO 5807:1985 Stored Data記号 / BPMN 2.0 Data Store（OMG） |
| `branch` | ひし形 | `{label}` | ISO 5807:1985 Decision記号 / UML Activity Diagram Decision Node（OMG UML 2.x） / BPMN 2.0 Exclusive Gateway（OMG） |
| `fork` / `join` | 六角形 | `{{label}}` | **Mermaid制約による代替形状**（後述） |
| `start` / `end` | スタジアム | `([Start])` / `([End])` | ISO 5807:1985 Terminal記号 |

### fork / join の形状について

UML 2.x標準ではfork / joinを太いバー記号（━━）で表現する。
Mermaid flowchartはこのバー記号を再現できないため、視覚的に区別しやすい六角形（`{{label}}`）を代替として採用する。

六角形はUML標準形状ではなく、**Mermaidのレンダリングツール制約による選択**である。

### エッジ種別

| 種別 | Mermaid記法 | 根拠 |
|------|------------|------|
| データ線 | `-->` | UML Activity Diagram ObjectFlow（OMG UML 2.x） |
| 制御線 | `==>` | UML Activity Diagram ControlFlow（OMG UML 2.x） |

### 制御線ラベル

| ラベル | 用途 | 根拠 |
|--------|------|------|
| `"<label>"` | branchのcases[].label | ADR-023 |
| `"parallel"` | forkの各ブランチへのエッジ | BPMN 2.0 Parallel Gateway（OMG） |
| `"foreach"` | foreachのapply先taskへのエッジ | BPMN 2.0 Multi-Instance Activity（OMG） |

### ノードの色付け

種別ごとにclassDefで色分けする。配色はWCAG 2.1（W3C）Level AAのコントラスト比基準（4.5:1以上）に準拠する。
色の意味づけはbrewprintのrender実装上の選択であり、国際標準の規定外。

| ノード種別 | classDef名 | fill | stroke | color |
|-----------|-----------|------|--------|-------|
| `task` | `taskNode` | `#4A90D9` | `#2C5F8A` | `#fff` |
| `asset` | `assetNode` | `#5BA55B` | `#3A6B3A` | `#fff` |
| `store` | `storeNode` | `#E8A838` | `#B07820` | `#fff` |
| `branch` | `branchNode` | `#9B6BBD` | `#6B3D8F` | `#fff` |
| `fork` / `join` | `forkNode` | `#8A8A8A` | `#5A5A5A` | `#fff` |
| `start` / `end` | `terminalNode` | `#2C2C2C` | `#000` | `#fff` |
| 外部参照 | `external` | `#E0E0E0` | `#999` | `#555` |

### 公知技術を根拠とする

各形状・エッジ種別はISO 5807・OMG UML 2.x・OMG BPMN 2.0という国際標準・公開仕様に基づいている。
これらはいずれも本ADR起票以前から広く公知の技術であり、brewprint固有の創作ではない。

### fork/joinの代替形状

レンダリングツール（Mermaid）の制約上、標準形状を再現できない。
代替形状の採用はツール制約への対応であり、概念そのものはUML標準に従う。

## 影響

- `spec/views/dag.md` のノード形状・エッジ種別・色付けの根拠として本ADRを参照する
- fork/joinの六角形がUML標準と異なることを `spec/views/dag.md` に注記する
- 制御線ラベルの根拠として本ADRを `spec/views/dag.md` から参照する

## Evidence
- commit: ce50bcf
- impl commit: tbd
- 参考: ISO 5807:1985（フローチャート記号）、OMG UML 2.x Activity Diagram（ControlFlow / ObjectFlow / Action Node / Object Node / Decision Node / Fork / Join）、OMG BPMN 2.0（Data Store / Exclusive Gateway / Parallel Gateway / Multi-Instance Activity）
