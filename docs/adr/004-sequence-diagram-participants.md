# 004: sequence diagramのparticipant設計

- **status**: accepted
- **date**: 2026-04-17

## 背景

sequence diagramを書くには「登場人物（participant）」が必要。
しかしbrewprintの既存ノード種別（task / asset / store）にはparticipantの概念がなく、
sequence diagramをどう表現するかが未定だった。

## 決定

sequence diagramのparticipantは以下の4種とする。

| participant | 粒度 | brewprintの実体 |
|---|---|---|
| Actor | 人間・外部システム | 新規：`actor` ノード |
| UI | イベント発火点 | `source=ui` なeventが存在する場合に暗黙生成 |
| API | バックエンドエンドポイント単位 | 既存：`task`（endpoint=true） |
| DB | 永続化層 | 既存：`store`（kind=db） |

### UI participantの暗黙生成

UI列はYAMLで明示的に宣言するノードを持たない。
`source=ui` なeventがsequence diagramのシナリオ内に存在する場合に、renderが自動でUI participantを生成する。

UIの種別（ブラウザ / モバイルアプリ等）の差異はstate diagram / wireframe側で表現する。
sequence diagramではUIを単一の参加者として扱い、種別を区別しない。

### participantのリンク

Mermaidのsequence diagramでは矢印にリンクを貼ることができない。
リンクはparticipant単位で付与する。

- `API` participant → class diagram（endpoint viewのURL）
- `DB` participant → ER diagram（storeのURL）
- `Actor` / `UI` → リンク不要

### 矢印のラベル

矢印の種別ごとにラベルの形式を定める。

| 矢印 | ラベル |
|------|--------|
| Actor → UI | event ID |
| UI → API | `METHOD path`（例: `POST /login`） |
| API → DB | `reads` または `writes`（方向で読み書きを表現） |
| API → UI | `returns.name`（returnsがない場合は `200 OK`） |

DB操作の詳細（どのtaskがどのstoreをいつ操作するか）はMermaid図の下にtableとして付記する。tableの `step` 列はシナリオの `steps:` インデックスと対応する。

### DB participantの粒度

`store.kind=db` のstoreは全て「DB」1列にまとめる。`kind=session` / `kind=collection` / `kind=context` はparticipant列に出ない（task内部で吸収）。

`store` はテーブル粒度の定義であり、スキーマ・DB単位の概念を持たないため、store IDごとに列を分けることはできない。複数DBを区別したい場合は `db_id` のような上位概念の追加が必要になるが、現時点ではスコープ外とする。

### happy pathのみ

sequence diagramはhappy pathのみを描画する。例外・エラーフローはnoteまたは別シナリオで表現する。taskの `returns` がhappy pathのレスポンスに対応する。

### 粒度の方針

sequence diagramは**レイヤー間の粗い粒度の流れ**を表現する。
関数呼び出しレベルの細かい粒度はDAGで表現する。

```
粗い粒度 → sequence diagram（Actor/UI/API/DB間のやりとり）
細かい粒度 → DAG（task内部の処理フロー）
```

## 理由

- sequence diagramに新規ノード種別を増やさず、既存要素のフラグ・属性追加で対応できる
- participantの種別が4つに限定されることで、「どの粒度で何を書くか」が明確になる
- APIをエンドポイント単位に絞ることで、Handler/Service/Repositoryの内部はDAG/ERに委譲できる

却下した代替案：
- participantを独立したノード種別として定義する → 既存要素との重複が生まれる
- sequence diagramをスコープ外にする → usecase的な「誰が何をするか」の表現手段がなくなる

## 影響

- `task` に `endpoint` フラグが追加される（ADR 005で詳細定義）
- `actor` ノードが新規追加される（軽量な定義で良い）
- sequence diagram viewのrenderロジックは、4種のparticipantをレイヤー順に並べる
- spec/overview.md の「書ける図の一覧」を更新する必要がある

## Evidence
- commit: db44639（初版）, ae16bba（DB participant・矢印ラベル・happy path追記）
- impl commit: tbd
- 参考: UML標準Actor定義参考
