# 024: DAGの境界ノード（params / returns subgraph）

- **status**: accepted
- **date**: 2026-04-20

## 背景

DAGのMermaid renderにおいて、メインノードの外部入力（`$params`）と最終出力（`returns`）をDAG内でどう表現するかが未定義だった。

これらを通常のassetと同じ扱いにすると、「DAGの中で生成されるデータ」と「DAGの境界を越えて入出力されるデータ」が視覚的に区別できない。

## 決定

### 1. params / returns を subgraph で囲む

メインノードの `$params`（外部入力）と `returns`（最終出力）のassetを、それぞれ `subgraph params` / `subgraph returns` で囲んでrenderする。

```
subgraph params
  config([config])
end

subgraph returns
  result([result])
end
```

これによりDAGの境界（どこから来てどこへ出るか）が一目で分かる。

### 2. 境界ノードの形状

通常のassetと同じスタジアム形状（`([label])`）を使う。形状の区別ではなく、subgraphのグループ化で境界を表現する。

### 3. 境界ノードの色付け（classDef）

`boundaryNode` クラスを追加する。通常のassetNode（緑）と区別できるティール系。

```
classDef boundaryNode fill:#2D7D9A,stroke:#1A5068,color:#fff
```

WCAG 2.1 Level AA（コントラスト比4.8:1）準拠（ADR-022と同基準）。

params / returns 両方の境界ノードに同じクラスを適用する。subgraphラベルで入力・出力の区別が既に明示されているため、クラスは1種類で十分。

### 4. Start / End ノードのID

Start / End ノードのMermaid IDは `_start` / `_end` とする（Mermaid組み込みキーワードとの衝突を避けるためアンダースコアプレフィックスを付与）。

```
_start([Start])
_end([End])
```

## 理由

### subgraphによる境界表現

UML 2.x Activity Diagramの **Activity Parameter Node**（OMG UML 2.x）に対応する概念。ActivityのI/Oを明示する記法であり、DAGの境界入出力をsubgraphで視覚化することはこの意味に沿う。

subgraphを使わない代替案（境界ノードに専用形状を使う等）も考えられるが、Mermaidのsubgraphはグループの「名前」を付けられるため、`params` / `returns` という意味ラベルを自然に表現できる。専用形状を導入するとノード形状の種類が増え、ADR-022で確立した形状体系が複雑になる。

### boundaryNodeクラスの単一化

params（入力）とreturns（出力）で別クラスにする案もあったが、方向はsubgraphのラベルで既に明示されており、色による二重表現は冗長。1クラスに統一することでclassDef管理が簡素になる。

### `_start` / `_end` のID

MermaidではいくつかのキーワードがID名として使えない場合がある。アンダースコアプレフィックスにより衝突リスクを排除する。ラベル（`Start` / `End`）は人間向け表示のため通常の英字を使う。

## 影響

- `spec/views/dag.md` のノードrender（start/end）・色付け・render例を本ADRに基づき更新する
- ADR-022の色付けテーブルに `boundaryNode` を追記する

## Evidence
- commit: ae532b0
- impl commit: tbd
- 参考: OMG UML 2.x Activity Diagram（Activity Parameter Node）、WCAG 2.1 Level AA
