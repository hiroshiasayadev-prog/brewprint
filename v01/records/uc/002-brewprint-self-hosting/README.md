# UC-002: brewprint Self-hosting

## 概要

UC-002 は、brewprint 自身を brewprint YAML で表現する self-hosting ユースケースである。

このUCでは、v1.0.0-spec の表現力と MCP coverage が、brewprint自身という実用規模のソフトウェアを表現するに足るかを検証する。
また、self-hosting 作業中に見つかった editor / viewer 要件を蓄積し、M14以降の仕様化候補として整理する。

作業は M14 の方針に従い、まず MCP 公開 contract を blueprint 化し、その後に内部レイヤーへ広げる。

## ファイル構成

```text
docs/uc/002-brewprint-self-hosting/
  README.md              ← このファイル
  TASKS-UC-002.md        ← UC-002固有のspec gap / 作業追跡
  editor-viewer-notes.md ← self-hosting中に見つかったeditor / viewer要件メモ
  render_index.yaml      ← render出力のgroup定義
  yaml/                  ← brewprint YAML群（single source of truth）
  renders/               ← Go rendererによる生成物。人間が直接編集しない
```

Phase A では `yaml/mcp/` 配下に MCP 公開 contract の YAML を追加していく想定。
`renders/` は renderer 出力先として先に作成しているが、手書きのrender成果物は置かない。

## render fixture 再生成

`renders/` は Go renderer の canonical fixture として管理する予定。
手編集ではなく、原則として以下のコマンドで再生成する。

```powershell
brewprint render --yaml-root docs/uc/002-brewprint-self-hosting/yaml --out docs/uc/002-brewprint-self-hosting/renders --clean
```

再生成後は以下を確認する。

```powershell
go test ./...
```

現時点では Phase A の YAML が未作成のため、render fixture は未生成。

## ドキュメント

| ファイル | 内容 |
|---|---|
| README.md | UC-002の概要とファイル構成 |
| TASKS-UC-002.md | UC-002固有の作業・spec gap・MCP coverage不足ログの追跡 |
| editor-viewer-notes.md | self-hosting中に見つかったeditor / viewer要件の蓄積メモ |
| render_index.yaml | 初期groupとしてMCP公開contractを定義 |
| docs/phase-a-mcp-contract.md | Phase AでMCP公開contractをblueprint化する表現方針 |
| docs/phase-a-work-split.md | Phase Aを並列作業するためのレーン分割・命名規約・merge前チェックリスト |
| renders/index.md | 未生成。Phase A render生成後に追加される想定 |

`docs/coverage.md` 相当のカバレッジ表は、Phase A の YAML が入った後に起票する。
MCP tool / node-field / render coverage を、実例YAMLに対応づけて追跡する。

## TODO / spec gap

未解決の仕様差分・追跡タスクは `TASKS-UC-002.md` を参照。

初期TODO:

- [ ] Phase A: MCP公開contract の blueprint 化を開始する
- [ ] Phase A render を生成し、`renders/` 配下に canonical fixture を配置する
- [ ] Phase A の YAML が入った後に `docs/coverage.md` を起票する
