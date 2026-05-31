# credential

## Public model

| property | value |
|---|---|
| kind | struct |
| visibility | public |
| source | yaml/auth/model/credential.yaml |

### Fields

| field | type | note |
|---|---|---|
| username | str | ユーザー名（PK）。他モジュールからの user_id FK 参照先 |
| password | str | パスワード。実装ではハッシュ化して保存する（本UCでは型だけ示す） |
| created_at | datetime | アカウント作成日時 |

