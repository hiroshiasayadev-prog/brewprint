# V01-ADR-031: actorのグローバル定義

- **status**: accepted
- **date**: 2026-04-22
- **supersedes**: V01-ADR-025

## 背景

V01-ADR-025では `actor` ノードは独立ファイルを持たず、参照元ファイル（`state.yaml` 等）のサブノードとして同居すると定めた。

しかし `stripe` や `scheduler` のような外部システム・外部actorは、複数のstate diagram / sequence diagramから参照されうる。V01-ADR-025の設計では同一actorが複数ファイルに重複定義される問題が生じる。

また、`source=external` のeventには発火元actorを明示する必要があり（V01-ADR-018改訂）、actorがモジュールスコープに縛られていると参照解決が複雑になる。

## 決定

### actorはプロジェクトglobalな存在として定義する

actorはモジュールに属さず、プロジェクト全体でユニークなIDを持つ。
任意のファイル名・任意の配置で `nodes:` リストに定義できる。

```yaml
# actors.yaml（ファイル名は任意）
nodes:
  - id: stripe
    type: actor
    note: "外部決済サービス"

  - id: scheduler
    type: actor
    note: "cronスケジューラー"

  - id: end_user
    type: actor
    note: "サービスを利用するエンドユーザー"
```

### モジュールパス不要

V01-ADR-027のsentinel方式による名前解決は適用しない。
actor IDはプロジェクト内でフラットにユニークであり、参照は常にIDのみで行う。

```yaml
# event定義での参照（V01-ADR-018）
- id: payment_webhook_received
  type: event
  source: external
  actor: stripe    # モジュールパスなし、ID直参照
```

### パーサーによる重複チェック

同一IDのactorが複数ファイルに定義されている場合、パーサーはエラーとする。

## 理由

### グローバル化の根拠

actorは「システム外部の存在」であり、特定モジュールの関心事ではない。
Stripeという外部サービスは認証モジュール・決済モジュール双方から参照されうるが、
それはStripeというエンティティが1つであることを意味する。モジュールスコープに縛ると
同一actorの重複定義を招き、整合性の保証が困難になる。

### 重複チェックの必要性

global IDである以上、同一IDの複数定義はパーサーレベルで検出・エラーにする必要がある。
ファイルを分けて定義したい場合（`common_actors.yaml` / `service_actors.yaml` 等）は
ID衝突がなければ問題ない。

却下した代替案：
- actorを参照元ファイルに同居（V01-ADR-025）→ 複数ファイルへの重複定義問題が生じる
- actorをモジュールスコープで管理 → 外部システムの性質と合わない

## 影響

- V01-ADR-025は本ADRにsupersededされる
- V01-ADR-027のsentinel予約語から `actor` を削除する（適用外のため）
- `spec/nodes.md` のactorセクションのファイル配置を更新する
- `spec/nodes.md` のノード種別一覧のactor行のファイル配置を更新する
- `source=external` のeventの `actor:` フィールドはglobal actor IDを直参照する（V01-ADR-018）

## Evidence
- commit: ae16bba
- impl commit: tbd
- 参考: 特になし
