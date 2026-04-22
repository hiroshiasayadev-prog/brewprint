# 032: sequence diagramシナリオファイルのYAML schema

- **status**: accepted
- **date**: 2026-04-22

## 背景

ADR-004でsequence diagramのparticipant種別・矢印ラベル・DB participantの粒度を確定した。
次の課題として「どのシナリオを描くか」を表現するYAML schemaが未定義だった。

sequence diagramはノード定義ファイル（`nodes:` + `transitions:`）から自動導出できない。
「どのeventを起点に、どのstateからのtransitionを追うか」はシナリオ固有の意図であり、
YAMLから機械的に推論できない。したがって専用のシナリオ定義ファイルが必要と判断した。

## 決定

### シナリオファイルはview定義ファイルとして `as: sequence_diagram` で自己宣言する

ADR-030の `as:` フィールド方式を踏襲する。

```yaml
as: sequence_diagram
id: login_flow
title: "ログインフロー"
state_file: auth/state.yaml    # 参照するstate定義ファイル
steps:
  - from_state: idle
    via: login_submitted

  - from_state: session_expired
    via: login_submitted
```

### フィールド定義

| フィールド | 必須 | 内容 |
|-----------|------|------|
| `as` | ✓ | `sequence_diagram` 固定 |
| `id` | ✓ | シナリオID。プロジェクト内でユニーク |
| `title` | 任意 | 人間向けタイトル |
| `state_file` | ✓ | 参照するstate定義ファイルのパス |
| `steps` | ✓ | シナリオのステップリスト |

### step オブジェクト

| フィールド | 必須 | 内容 |
|-----------|------|------|
| `from_state` | ✓ | 遷移前のstate ID。省略不可 |
| `via` | ✓ | 発火するevent ID |

`(from_state, via)` のペアで `state_file` 内の `transitions:` を一意に特定する。
対応するtransitionが存在しない場合はパーサーエラー。

### バックエンドによる自動解決

シナリオYAMLに明示するのは `(from_state, via)` のみ。以下はバックエンドが自動解決する。

| 情報 | 解決元 |
|------|--------|
| 矢印の送信元participant | event の `source` / `actor` フィールド（ADR-018） |
| 呼び出されるtask | transition の `action` フィールド |
| UI → API の矢印ラベル | task の `method` / `path` |
| API → DB の矢印・方向 | task の `reads` / `writes`（kind=dbのstoreのみ） |
| API → UI の矢印ラベル | task の `returns.name`（なければ `200 OK`） |
| UI participantの生成 | `source=ui` のeventが存在する場合に暗黙生成（ADR-004） |

### DB操作tableの付記

Mermaid図の下に、DB操作の詳細をtableとして付記する。
`step` 列はシナリオの `steps:` の1-originインデックスと対応する。

```markdown
| step | task | store | 操作 |
|------|------|-------|------|
| 1 | auth.task.login | user_db | reads |
| 1 | auth.task.login | session_store | writes |
```

`kind=session` / `kind=collection` / `kind=context` のstoreはtableにも出力しない。

## 理由

### 専用シナリオファイル（案A）を選択

- **案B（state diagramから自動導出）**: 「どこで切るか」の粒度がYAMLから自動判定できない。却下。
- **案C（state.yamlに `scenarios:` セクション追加）**: 1ファイルがFSM定義とview定義の二役を担うことになり、ADR-030の原則と相性が悪い。また1シナリオ1ファイルの方がdiff・レビュー・参照が明確。却下。

### `from_state` 必須・省略なし

同一eventに対して `from_state` が異なる複数のtransitionが存在しうる（Mealy machine）。
`from_state` を省略すると `(from_state, via)` によるtransition特定ができなくなる。
省略時の曖昧性を許容するメリットがないため必須とする。

### `to:` フィールド不要

矢印の送信先はeventの `source` / `actor` からバックエンドが解決できる。
手書きは情報の重複であり不要。

## 影響

- ADR-030の `as:` 値一覧に `sequence_diagram` を追加する
- `spec/views/sequence-diagram.md` を新規作成する
- Goパーサーは `as: sequence_diagram` のファイルをsequence diagramシナリオとして処理する

## Evidence
- commit: e5d4c5d
- impl commit: tbd
- 参考: 特になし
