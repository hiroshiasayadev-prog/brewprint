## brewprint プロジェクト設定

- ローカルリポジトリ: `C:\Users\imved\projects\brewprint`（filesystem / md-section MCPで参照可能）

### 会話開始時に必ず行うこと

1. `docs/doc-policy.md` を読む
2. `docs/adr/` の一覧を取得し、acceptedなADRのタイトルを把握する
3. 作業に関連するspec / ucを必要に応じて読む

**全docを最初から読まなくていい。** ADRタイトルで文脈を把握し、必要なものだけ読む。

### 会話のスコープを定義すること

contextは有限。この会話でどこまでの範囲を進めるかを最初に決める。
スコープの作業完了後はcommitを勧めること。

### ユーザーを完全に信頼しない

人間は不完全。docsをすべて正確に把握しているわけではないし、整合が取れているとも限らない。
docsや技術的な観点での不整合・腑に落ちない部分があれば列挙してユーザーの判断を仰ぎ、必ずdocsに明記すること。

### docsを完全に信頼しない

書いたのはAI。必ずしもユーザーの意図を全て正確に反映しているとは限らない。
ユーザーの発言・意図と不整合があるように感じられたら列挙してユーザーの判断を仰ぎ、必ずdocsに明記すること。

### 会話中に適宜行うこと

- **commitの提案** — 1つのトピックの議論が完了したタイミングでcommitを勧め、同意が得られたら設計判断をdocsに明記し、git commitコマンドを出すこと
- **ADRの起票** — 設計決定が確定したら自動的にADRを書く（proposed → accepted）
- **既存ADRを覆す決定をした場合** — 旧ADRをsuperseededに更新し新ADRを起票する

### ファイル操作

- 新規ファイル → `filesystem:write_file`
- 既存ファイルの部分更新 → `str-replace:str_replace_in_file` 優先
- 複数箇所同時更新 → `filesystem:edit_file`

### 回答は日本語
