# 026: FKカーディナリティとN:M表現方針

- **status**: accepted
- **date**: 2026-04-20

## 背景

ER図生成に際し、`fk:` フィールドだけではカーディナリティ（1:1 vs N:1）が判定できず、
またN:Mリレーションの表現方法が未定だった。
複合PKの扱いも論点として浮上した。

## 決定

### 1. `fk:` のデフォルトカーディナリティは many-to-one

`fk:` を持つフィールドは、明示がない限り many-to-one として扱う。

### 2. 1:1 は `unique: true` で明示

```yaml
- name: profile_id
  type: str
  fk: profile.id
  unique: true    # 1:1 を明示。FK側カラムにUNIQUE制約を付与する
```

`unique: true` は省略可能（optional）。付与しない場合は many-to-one とみなす。

### 3. N:M は中間 model + surrogate key で表現

N:M リレーションを直接表現する構文は持たない。
中間 model（FK を2本持つ struct）として明示的に定義する。
中間 model には surrogate key（`pk: true` の `id` フィールド）を持たせることを推奨する。

```yaml
# model/user_tag.yaml
- id: user_tag
  type: model
  kind: struct
  fields:
    - name: id
      type: str
      pk: true
    - name: user_id
      type: str
      fk: user.id
    - name: tag_id
      type: str
      fk: tag.id
```

中間 model もDBテーブルとして実在するため、対応する `store.kind: db` を定義する必要がある。
ER図の描画対象は `store.kind: db` から辿れる model に限定されるため（`spec/views/er.md`）、
`store.kind: db` を定義しない中間 model はER図に登場しない。

```yaml
# store/user_tag_db.yaml
- id: user_tag_db
  type: store
  kind: db
  of: user_tag
```

### 4. 複合 PK は現時点で非サポート

`pk: true` の「1 struct に1つ」制約（ADR-021）を維持する。
複合 PK が必要なケースは surrogate key で代替する。
将来的に需要があればissue経由で追加する。

## 理由

- surrogate key は Django・Rails・Prisma 等のデフォルト方針と一致しており、現実的なユースケースをカバーする
- 複合 PK のサポートは ADR-021 の改訂を伴う割に、初期ユースケースで必要性が低い
- brewprint は「まず動くものを出す」フェーズであり、実需がなければ仕様を増やさない

## 影響

- `spec/nodes.md` の field オブジェクト定義に `unique` フィールドを追加する
- `spec/views/er.md` でカーディナリティの導出ルールを定義する

## Evidence
- commit: tbd
- impl commit: tbd
- 参考: Django/Rails/PrismaのSurrogate key慣習参考
