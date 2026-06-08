# V01-ADR-020: クロスエッジの管理方式

- **status**: accepted
- **date**: 2026-04-19
- **supersedes**: V01-ADR-015

## 背景

`spec/overview.md` にてクロスエッジ（レイヤー間エッジ）として5種が定義されていた：
`write` / `read` / `trigger` / `reflect` / `hydrate`

これらの管理方式（YAML上の記述場所・構造）が未確定だった。
また、V01-ADR-015にてflow:ステップの`reads`/`writes`フィールドとして暫定的に置いていたが、
「taskの性質であってflowが持つべき情報ではない」という観点から見直しが必要だった。

## 決定

### 1. クロスエッジ種別を `write` / `read` の2種に絞る

以下の理由で残り3種を廃止する：

| kind | 廃止理由 |
|------|---------|
| `trigger` | `transition.action`（V01-ADR-019）で既に表現済み |
| `reflect` | event + transition の連鎖で既に表現済み |
| `hydrate` | event + transition の連鎖で既に表現済み |

### 2. `reads` / `writes` はtask nodeのフィールドとして書く

flow:ステップではなく、`nodes:`内のtask定義に直接持たせる。

```yaml
nodes:
  - id: process_payment
    type: task
    reads: [balance_store]
    writes: [transaction_store, balance_store]
    note: "transactionとbalanceは同一トランザクションで更新"
```

| フィールド | 対象 | 内容 |
|-----------|------|------|
| `reads` | store ID のリスト | このtaskが参照するstore |
| `writes` | store ID のリスト | このtaskが更新するstore |

### 3. 複数storeへの reads / writes を許容する

1task = 1store制約は設けない。複数storeにまたがる場合はtruncation境界を `note` に記述する。

```yaml
- id: process_payment
  type: task
  writes: [transaction_store, balance_store]
  note: "transactionとbalanceは同一トランザクションで更新"
```

機械的なtransaction境界の表現はbrewprintのスコープ外とし、noteによるLLMへのセマンティックcontractで担保する（V01-ADR-008と同じ思想）。

### 4. クロスエッジの追跡はGoのMCPツールが担う

「このstoreはどのtaskからwriteされているか」「このtaskはどのtransitionからtriggerされるか」等の逆引き追跡は、Go実装のMCPツールがASTを解析して提供する。YAML上に逆参照を手書きする構造は設けない。

## 理由

### `reads`/`writes` をtask nodeに置く

flow:ステップに書くと「このtaskはこのflow文脈ではこのstoreを読む」という意味になりかねない。しかしstoreの参照・更新はtaskの実装に紐づく性質であり、どのflowから呼ばれようとも変わらない。task nodeに置くことで責務が明確になる。

同じtaskを複数flowで再利用するケースが発生した場合も、task nodeに書いておけば一箇所で把握できる。

### trigger / reflect / hydrate の廃止

これらをクロスエッジとして別管理すると、既にV01-ADR-018（event）・V01-ADR-019（state/transition）で表現されている情報が二重管理になる。YAMLを読む際の追跡経路も増える。既存の構造で表現できるものは新たな概念を導入しない。

### 複数store許容

ログイン処理（session + last_login更新）・決済処理（transaction + balance更新）のように、1つのユースケースが複数storeをアトミックに更新するケースは実務上避けられない。1task = 1store制約を設けるとtask数が爆発し、トランザクション境界がYAMLから消える。

## 影響

- V01-ADR-015の「5. storeのreads/writes」（flow:ステップへの記述）は本ADRによりsupersede。task nodeに書く
- `spec/overview.md` のクロスエッジ種別テーブルを `write` / `read` の2種に更新する
- `spec/nodes.md`（未作成）にてtask nodeの `reads` / `writes` フィールドを定義する
- `spec/edges.md`（未作成）に本ADRの内容を反映する

## Evidence
- commit: 15e69b4
- impl commit: tbd
- 参考: 特になし
