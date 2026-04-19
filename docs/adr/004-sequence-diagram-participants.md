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
| UI | イベント発火点 | 既存：`event`（source=ui） |
| API | バックエンドエンドポイント単位 | 既存：`task`（endpoint=true） |
| DB | 永続化層 | 既存：`store`（kind=db） |

### participantのリンク

Mermaidのsequence diagramでは矢印にリンクを貼ることができない。
リンクはparticipant単位で付与する。

- `API` participant → class diagram（endpoint viewのURL）
- `DB` participant → ER diagram（storeのURL）
- `Actor` / `UI` → リンク不要

### 矢印のラベル

矢印のラベルにはtask IDを記載する（`auth.task.login` など）。
リンクにはならないが、IDがあればMCP経由で詳細を参照できる。

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
- commit: db44639
- impl commit: tbd
- 参考: UML標準Actor定義参考
