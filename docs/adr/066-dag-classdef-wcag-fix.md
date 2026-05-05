# 066: DAG node classDef の WCAG コントラスト比修正

- **status**: accepted
- **date**: 2026-05-05

> このADRは起票時点での決定を記録したスナップショットである。
> 現在の仕様は spec を参照すること。

## 背景

ADR-022 で DAG render の classDef 配色が決まり、`docs/spec/views/dag.md` でも「WCAG 2.1 Level AA（コントラスト比 4.5:1 以上）に準拠」と明記されている。

ADR-064 起票検討の過程で initStoreNode 新設時のコントラスト比を試算した際、**既存 classDef の大半が AA 基準を満たしていない**ことが判明した。

WCAG 2.1 アルゴリズム（IEC 61966-2-1 sRGB Linearization → 相対輝度 → コントラスト比）で各 classDef の `fill` と `color` の比を計算すると以下のとおり：

| classDef | fill | color | ratio | AA (4.5)? |
|---|---|---|---|---|
| taskNode | `#4A90D9` | `#fff` | 3.34 | **NG** |
| assetNode | `#5BA55B` | `#fff` | 3.01 | **NG** |
| storeNode | `#E8A838` | `#fff` | 2.08 | **NG** |
| branchNode | `#9B6BBD` | `#fff` | 4.01 | **NG** |
| forkNode | `#8A8A8A` | `#fff` | 3.45 | **NG** |
| terminalNode | `#2C2C2C` | `#fff` | 13.97 | OK |
| boundaryNode | `#2D7D9A` | `#fff` | 4.65 | OK |
| external | `#E0E0E0` | `#555` | 5.65 | OK |

8 classDef のうち 5 つ（taskNode / assetNode / storeNode / branchNode / forkNode）が AA 違反。
ADR-022 §決定 §「公知技術を根拠とする」および `spec/views/dag.md` §ノードの色付け の WCAG 2.1 AA 準拠記述は、実態と乖離している。

本ADRはこの乖離を解消する。

## 決定

### 1. 違反 5 classDef の文字色を `#fff` → `#000` に変更する

`fill` 色は据え置き、`color` のみ `#fff` から `#000` に変更する。

| classDef | fill（変更なし） | color 変更 | 新 ratio |
|---|---|---|---|
| taskNode | `#4A90D9` | `#fff` → `#000` | 6.28 |
| assetNode | `#5BA55B` | `#fff` → `#000` | 6.98 |
| storeNode | `#E8A838` | `#fff` → `#000` | 10.09 |
| branchNode | `#9B6BBD` | `#fff` → `#000` | 5.23 |
| forkNode | `#8A8A8A` | `#fff` → `#000` | 6.08 |

5 classDef すべて AA 基準（4.5:1）を満たす。

### 2. AA 準拠の他 classDef は変更しない

terminalNode（13.97）/ boundaryNode（4.65）/ external（5.65）は AA 基準を既に満たしているため変更しない。
コントラスト変更は最小限とし、WCAG 違反 classDef のみを修正する。

### 3. 修正後の classDef 一覧

```
classDef taskNode     fill:#4A90D9,stroke:#2C5F8A,color:#000
classDef assetNode    fill:#5BA55B,stroke:#3A6B3A,color:#000
classDef storeNode    fill:#E8A838,stroke:#B07820,color:#000
classDef branchNode   fill:#9B6BBD,stroke:#6B3D8F,color:#000
classDef forkNode     fill:#8A8A8A,stroke:#5A5A5A,color:#000
classDef terminalNode fill:#2C2C2C,stroke:#000,color:#fff
classDef boundaryNode fill:#2D7D9A,stroke:#1A5068,color:#fff
classDef external     fill:#E0E0E0,stroke:#999,color:#555
```

`stroke` も変更しない。`fill` の縁取りとしての可読性は ADR-022 の選択をそのまま維持する。

### 4. AAA 基準（7.0:1）への引き上げは目指さない

ADR-022 / spec/views/dag.md は AA 基準を準拠目標として明記している。AAA 引き上げを行うと assetNode（6.98）/ taskNode（6.28）/ branchNode（5.23）/ forkNode（6.08）/ boundaryNode（4.65）が再修正対象になり、`fill` 色の変更が必要になる。

ADR-022 の配色は ISO 5807 / UML / BPMN といった公知技術根拠ではなく brewprint render 実装上の選択（ADR-022 §決定 §「ノードの色付け」）であり、`fill` を変更しても公知技術根拠は崩れない。しかし AA を満たした時点で WCAG 準拠記述との乖離は解消されるため、本ADRでは AAA 引き上げは行わない。

将来 AAA 引き上げが必要になった場合は別 ADR で扱う。

## 理由

### なぜ文字色変更で済ませるか

`fill` 色を変更すると ADR-022 の配色全体（taskNode は青系・assetNode は緑系等）の見た目が大きく変わり、既存 UC-001 の視覚的同一性が失われる。文字色を `#fff` → `#000` に切り替えるだけで AA 基準を満たせるため、変更範囲を最小化する。

`#000` 文字は ADR-022 §決定 §「ノードの色付け」で external classDef（薄灰背景）に既に採用されており、brewprint render の語彙として既知である。

### なぜ AA 違反だけ直して AAA は目指さないか

ADR-022 / spec/views/dag.md の準拠目標は AA であり、AA を満たせば spec 記述との乖離は解消される。AAA は WCAG 2.1 が「強化レベル」と位置付けるオプショナル基準であり、本ADRのスコープ（WCAG 違反の修正）を超える。

### なぜ ADR-022 を遡及修正しないか

ADR-050 §3 / doc-policy.md §3 「ADR は遡及修正しない」原則により、ADR-022 §決定 §「ノードの色付け」の表は起票時点のスナップショットとして残す。本ADR-066 が補追として現行配色を上書きする。

ADR-022 の冒頭には partial superseded annotation を追加し、§「ノードの色付け」のうち `color` 列が ADR-066 で更新されたことを明示する（§影響 参照）。

### 却下した代替案

#### 代替案A: `fill` 色も変更し AAA 基準まで引き上げる

- 利点: アクセシビリティが最大化される
- 欠点: ADR-022 の配色体系（青=task / 緑=asset / オレンジ=store / 紫=branch / 灰=fork）が崩れる。既存 UC-001 の視覚的同一性が失われる。本ADRのスコープ（WCAG 違反の修正）を超える

→ 却下。

#### 代替案B: 違反は spec 記述側を修正して「best-effort」と書き換える

- 利点: 配色変更ゼロ
- 欠点: アクセシビリティが実態として AA 違反のまま放置される。ADR-022 の「WCAG 2.1 Level AA」明記の意図を裏切る

→ 却下。

#### 代替案C: ADR-022 を遡及修正して色テーブルを直接書き換える

- 利点: ADR-022 を読めば最新の値が分かる
- 欠点: ADR-050 §3 / doc-policy.md §3 の「ADR は遡及修正しない」原則に反する

→ 却下。本ADR-066 を新規起票して補追する。

## 影響

### spec への影響

- `docs/spec/views/dag.md` §ノードの色付け の classDef 5 個（taskNode / assetNode / storeNode / branchNode / forkNode）の `color` を `#000` に書き換える
- §ノードの色付け に「WCAG 2.1 AA 準拠コントラスト比は ADR-066 で再検証済」旨の由来注記を追加する
- 全 render 例の classDef 行を新値に置換する

### 既存 ADR への影響

- **ADR-022**: 冒頭に partial superseded annotation を追加。「§決定 §ノードの色付け の `color` 列のみ ADR-066 で更新」旨を明示。本文は遡及修正しない
- ADR-024: `boundaryNode` は AA 準拠のため修正対象外。影響なし

### 既存実装への影響

- `internal/render/dag` の classDef 出力に新 `color` 値を反映する
- 出力の Mermaid 文字列が変わるため、文字列比較ベースの test が変わる

### 既存 UC への影響

- UC-001 の renders/ 配下 Mermaid 出力の classDef 行が変わる
- レンダリング結果は文字色のみ変わる（白文字 → 黒文字）。`fill` 色は同一のため視覚的な違和感は最小

### v1.1 への影響

本ADRは bug fix 相当（spec 記述と実態の乖離解消）であり、v1.1.0-spec の凍結対象に含める。

### 後続 ADR への影響

- ADR-064 で新設される `initStoreNode` classDef は本ADRの方針に従い `color: #000` を採用する

## Evidence

- commit: 96aa78c
- impl commit: tbd
- 参考: WCAG 2.1（W3C）, IEC 61966-2-1 sRGB Linearization, ADR-022 DAG node shape and edge types
