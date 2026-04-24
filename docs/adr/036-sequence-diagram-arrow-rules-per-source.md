# 036: sequence diagram矢印ルールのevent source別拡張

- **status**: accepted
- **date**: 2026-04-24

## 背景

ADR-004でsequence diagramのparticipant種別・矢印ラベルを定義した時点では、`source: ui` を起点とするフロー（UI → API → DB → UI）が主な想定だった。

その後の拡張でeventの `source` は4種に増えた:

- `source: ui`（ADR-018）
- `source: external`（ADR-018）
- `source: er`（ADR-018）
- `source: internal`（ADR-034）

しかし `spec/views/sequence-diagram.md` の矢印ラベル表・participants生成条件表は主に `ui` を想定した状態のままで、以下のケースが未定義:

1. **`source: external` でUIを介さないフロー**（例: Stripe webhook、スケジューラー起動）
   - Actor が UI を経由せず直接 API を叩くが、`Actor → API` の矢印ラベルが未定義
   - API応答側の `API → Actor` も未定義
2. **`source: internal` のFSM内部駆動event**（例: task完了による `login_succeeded`）
   - 矢印の送信元participantが原理的に存在しない
3. **`source: er` のstore値変化駆動event**
   - 矢印の起点が未定義（監視対象storeから発火？APIの内部処理？）
   - `store.kind` による差も未整理

UC-001の `payment_webhook_flow` がまさに (1) のケースに該当し、既存ルールでは描画方針が決まらない。`sequence-diagram.md` 内にシナリオYAMLの例として `scenarios/payment_webhook_success.yaml` は存在するが、対応するMermaid出力イメージが欠けていた。

## 決定

sequence diagramの矢印ルール・participant生成条件をeventの `source` 4種に応じて以下の通り拡張する。

### sourceごとの矢印ルール

| `event.source` | participant追加 | 起点→API矢印 | ラベル | API応答矢印 |
|---|---|---|---|---|
| `ui` | UI暗黙生成 | `UI→API` | `METHOD path` | `API→UI`: `returns.name` / `200 OK` |
| `external` | Actor生成 | `Actor→API` | `METHOD path` | `API→Actor`: `returns.name` / `200 OK` |
| `internal` | なし | `API→API`（自己ループ） | event ID | なし |
| `er`（`watches`先が `store.kind=db`） | DB既存 | `DB→API` | event ID | なし |
| `er`（`watches`先が `store.kind≠db`） | なし | `API→API`（自己ループ） | event ID | なし |

### participant生成条件の整理

シナリオ内の全stepで参照されるevent・taskから以下のルールで生成:

| participant | 生成条件 |
|---|---|
| Actor | `source=external` のeventを参照するstepが1つでもある場合。対象actorは `event.actor` で指定されたglobal actor（ADR-031） |
| UI | `source=ui` のeventを参照するstepが1つでもある場合に暗黙生成 |
| API | シナリオのstepが参照するtaskの `endpoint=true` から生成 |
| DB | 以下いずれかの条件でstepが `store.kind=db` を参照する場合に生成<br>・taskの `reads` / `writes`<br>・`source=er` eventの `watches` |

表示順（左→右）: `Actor → UI → API → DB`（存在するもののみ）

### ラベル選択の原則

- **`Actor→API` / `UI→API` が `METHOD path`**: いずれも物理的にHTTPリクエストの到達を表す。`UI→API` との対称性を保つ。event IDはHTTPリクエストのbody/headerを解釈して初めて同定される情報であり、起点ラベルには適さない。
- **自己ループ / `DB→API` が event ID**: 内部駆動・データ変化駆動いずれもHTTPの物理表現が存在しない。「どのeventが駆動したか」を示すことが本質情報であり、event IDが唯一妥当なラベル。

### API応答矢印の有無

`ui` / `external` は呼び出し元（UI / Actor）が存在するためAPI応答矢印を描画する。
`internal` / `er` はFSM runtime駆動であり、呼び出し元に「返す」概念が存在しないため応答矢印を描画しない。

### `source=er` の矢印分岐根拠

`store.kind=db` のstoreは既にDB participantとして列に存在する（ADR-004）ため、`DB→API` として自然に描画できる。
`kind=session` / `kind=collection` / `kind=context` のstoreはparticipant列に出ない（ADR-004）ため、矢印の起点として表現できない。したがって `source: internal` と同様に自己ループで描画する。

## 理由

### 4 sourcesの網羅的定義

ADR-018で3種、ADR-034で4種目のsourceが追加されたが、sequence diagramの描画ルールは ADR-004 当時の想定（主にui起点）のまま取り残されていた。本ADRでこの不整合を解消する。

### `Actor→API` のラベルを `METHOD path` にした理由

候補:

1. `event ID`（`Actor→UI` との対称性）
2. `METHOD path`（`UI→API` との対称性）
3. 両方併記

webhook/スケジューラーにおける `Actor→API` は物理的にHTTPリクエストであり、`UI→API` と同じ「HTTPエンドポイントへの到達」を表す。event IDはHTTPリクエストを受信・解析して初めて得られる情報で、起点ラベルにはそぐわない。2を採用。併記案は情報過多かつMermaid上の可読性を損なうため却下。

### 自己ループの採用（`source: internal` / `er` で kind≠db）

`source: internal` はアプリ内部（FSM runtime）で発火するため、原理的に「外部起点」が存在しない。Mermaidのself-message記法（`API->>API`）で「API側で何かが起きた → 次のtaskが実行された」ことを表現する。event IDをラベルにすることで、何が駆動したかは図から読み取れる。

`er` で `kind≠db` のケースは「観測対象が図に出ていない」状況であり、矢印の起点を描けない点で `internal` と本質的に同じ。ルールを統一する。

### 却下した代替案

- **`source: internal` をnoteで表現**: Mermaidのnoteは時系列の一部として扱いにくく、stepの順序表現が崩れる。自動解決ルールがnote生成に特殊分岐を持つことになり実装も複雑化する。
- **`source: er` を一律 `DB→API` にする**: `kind=session` 等はDB participantに出ない設計（ADR-004）と矛盾する。
- **`source: er` を一律自己ループ**: `kind=db` のケースで「DB観測駆動」という本質情報を矢印で表現できなくなる。
- **ADR-004をsupersedeして書き直す案**: ADR-004のparticipant種別・DB粒度・happy path方針は変更不要。矢印ルールのみの拡張であり、独立したADRとして履歴を残すほうが辿りやすい（ADR-018 → ADR-034 と同じパターン）。

## 影響

- ADR-004 はparticipant種別・DB粒度・happy path方針は変更なし。矢印ラベル表のみ本ADRで拡張（supersedeはしない）
- `spec/views/sequence-diagram.md` の以下を更新:
  - `participants` セクションの生成条件表（DB生成条件にer経由の経路を追加）
  - `矢印ラベル` セクションの表を5パターンに拡張
  - `バックエンドによる自動解決` 表の拡張
  - `Mermaid出力イメージ` に `payment_webhook_success` のwebhookケース例を追加
- ADR-034 で追加された `source: internal` の可視化方針が本ADRで補完される

## Evidence
- commit: fdcd7a8
- impl commit: tbd
- 参考: UML 2.x Sequence Diagramのself-message記法、HTTP webhookにおけるreceiver-side modeling慣習
