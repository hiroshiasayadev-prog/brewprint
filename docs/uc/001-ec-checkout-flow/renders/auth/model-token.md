# token

## Public model

| property | value |
|---|---|
| kind | struct |
| visibility | public |
| source | yaml/auth/model/token.yaml |

### Fields

| field | type | note |
|---|---|---|
| access_token | str | アクセストークン文字列（PK） |
| user_id | str | credentialへの1:1 FK。1ユーザーにつき有効トークンは1つ（ADR-026） |
| expires_at | datetime | トークン有効期限 |

