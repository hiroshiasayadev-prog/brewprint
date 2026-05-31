# get_items

**API**: [GET /api/catalog_items](../_cross/api.md)

販売中の商品一覧を返す。
item_collection の note にある `list_available` クエリ（is_available = true）を使用。
params なし（フィルタ無しで全件返す簡易仕様。検索/ページングは別endpointへ）。

```mermaid
flowchart TD
  _start([Start]) ==> get_items[get_items]

  item_collection[(item_collection)] -- "read" --> get_items

  get_items --> items([items])
  get_items ==> _end([End])

  classDef taskNode     fill:#4A90D9,stroke:#2C5F8A,color:#fff
  classDef storeNode    fill:#E8A838,stroke:#B07820,color:#fff
  classDef assetNode    fill:#5BA55B,stroke:#3A6B3A,color:#fff
  classDef terminalNode fill:#2C2C2C,stroke:#000,color:#fff
  classDef boundaryNode fill:#2D7D9A,stroke:#1A5068,color:#fff
  class get_items taskNode
  class item_collection storeNode
  class items assetNode
  class _start,_end terminalNode
```

## Tasks

### get_items

#### Returns

| name | model | source |
|---|---|---|
| items | item_list | — |

#### Store access

| access | store |
|---|---|
| read | item_collection |
