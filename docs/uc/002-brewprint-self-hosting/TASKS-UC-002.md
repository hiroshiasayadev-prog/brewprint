# TASKS-UC-002: brewprint Self-hosting

UC-002 固有の作業・spec gap・editor / viewer 要件の発見ログを追跡するための作業台。
プロジェクト全体の入口は `docs/TASKS.md`、milestone単位の方針は `docs/tasks/m14-self-hosting.md` に置く。

---

## ステータス

- [x] UC-002 ディレクトリ骨格を作成
  - [x] `yaml/` を作成
  - [x] `renders/` を作成
- [x] `README.md` を起票
- [x] `render_index.yaml` を起票
- [x] `TASKS-UC-002.md` を起票
- [x] `editor-viewer-notes.md` を起票
- [ ] Phase A: MCP公開contract の blueprint 化を開始
- [ ] Phase A の YAML が入った後に `docs/coverage.md` を起票

---

## 作業方針

M14の方針に従い、UC-002は v1.0.0-spec を基準に self-hosting を進める。

作業順序:

1. MCP公開contract を blueprint 化する
2. Phase A render を生成・確認する
3. 内部レイヤーを layer 単位で blueprint 化する
4. UC-002全体ERとMCP coverage実用検証レポートをまとめる
5. self-hosting中に発見したspec gapを本ファイルへ記録する
6. editor / viewer 要件は `editor-viewer-notes.md` に蓄積する

---

## Phase A: MCP公開contract

- [ ] MCP tools の blueprint 表現方針を決める
- [ ] MCP server / client actor を定義する
- [ ] MCP tool request / response model を定義する
- [ ] MCP tool 呼び出しを task / flow として表現する
- [ ] Phase A範囲のrenderを生成・確認する

---

## Phase B: 内部レイヤー

- [ ] source layer を blueprint 化する
- [ ] rawyaml layer を blueprint 化する
- [ ] semantic layer を blueprint 化する
- [ ] resolve layer を blueprint 化する
- [ ] query layer を blueprint 化する
- [ ] render layer を blueprint 化する
- [ ] CLI（`brewprint render` / `brewprint validate`）を blueprint 化する

---

## spec gap 発見ログ

現時点では未記録。

発見時は以下の形式で追記する。

```markdown
### N. タイトル

- 対象: `docs/spec/...`
- 発見元: `docs/uc/002-brewprint-self-hosting/yaml/...`
- 状況:
- 論点:
- 暫定対応:
- 分類: v1範囲内修正 / v2向け構造変更 / 棄却 / 判断待ち
```

---

## coverage.md 起票方針

`docs/coverage.md` は、Phase A の YAML が入った後に起票する。
空の骨格段階ではなく、MCP公開contractの実例YAMLに対して、以下の観点を対応づけて記録する。

- MCP tool coverage
- node / field coverage
- render coverage
- spec gap の発見元

---

## MCP coverage / tool不足ログ

現時点では未記録。

MCP toolsだけでは探索・確認が難しかった点があれば、ここに記録する。

---

## 完了条件

- [ ] MCP公開contract が blueprint 化されている
- [ ] 内部レイヤーが layer 単位で blueprint 化されている
- [ ] Phase A / Phase B のrender結果を確認している
- [ ] `docs/coverage.md` が実例YAMLに追随している
- [ ] spec gap 発見ログをレビューし、v1範囲内修正 / v2向け構造変更 / 棄却に分類している
- [ ] MCP coverage 実用検証レポートをまとめている
- [ ] editor / viewer notes のspec昇格要否を判断している
