# V01-ADR-002: フォルダ階層=モジュール階層=名前空間

- **status**: accepted
- **date**: 2026-04-17
- **migrated_to_spec**: 2026-04-29
- **supersedes**:

> このADRの現行仕様詳細は [docs/spec/naming.md](../spec/naming.md) §1 を参照。

> 補足: V01-ADR-010により `dag.yaml` / `er.yaml` のビュー別ファイル分けは廃止。V01-ADR-043により `master.yaml` は廃止。名前空間ルール（フォルダ = モジュール）は継続有効。

## 背景

クロスエッジはノードIDの参照で成立するため、ID体系を先に確立する必要があった。
候補はフラット・種別prefix・モジュール階層の3択。

## 決定

フォルダ階層をそのまま名前空間とする（仕様詳細: [naming.md](../spec/naming.md) §1）。

加えて、リポジトリルートに `master.yaml` を置く方針を当初定めた（modules一覧・エントリーポイント・グローバルなクロスエッジ・履歴管理）。

> ※ `master.yaml` は V01-ADR-043 でプロジェクトルートレイアウトを再定義した際に廃止された。現行仕様では `master.yaml` は存在しない（[project-layout.md](../spec/project-layout.md) §1）。

## 理由

- フォルダ構造が名前空間になることで、MCPの`get_deps`時に「このIDがどのモジュールに属するか」がパスから自明になる
- YAMLのフォルダ階層・renderされたmirror mdのフォルダ階層・実装が一致するため、人間にとっても追いやすい

却下した代替案：
- フラット → ID衝突リスクがあり、大規模になると破綻する
- 種別prefix（`proc_xxx`など） → モジュール間の関係が見えにくい

## 影響

- クロスエッジのID参照は`モジュール名.ノード種別.ID`の3階層フルパス
- master.yamlはgo.mod的存在として別途スキーマを定義する
- ADR 003（名前解決ルール）に依存する

## Evidence
- commit: 152e83e
- impl commit: tbd
- 参考: Goのgo.mod台帳構造参考
