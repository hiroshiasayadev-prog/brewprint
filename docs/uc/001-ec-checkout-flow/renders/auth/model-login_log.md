# login_log

## Public model

| property | value |
|---|---|
| kind | struct |
| visibility | public |
| source | yaml/auth/model/login_log.yaml |

### Fields

| field | type | note |
|---|---|---|
| id | str | ログID（PK） |
| username | str | 認証試行ユーザー名 |
| success | bool | 認証成否フラグ（bool primitive の例） |
| attempted_at | datetime | 認証試行日時 |

