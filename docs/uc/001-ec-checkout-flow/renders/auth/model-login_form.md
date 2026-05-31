# login_form

## Public model

| property | value |
|---|---|
| kind | struct |
| visibility | public |
| source | yaml/auth/model/login_form.yaml |

### Fields

| field | type | note |
|---|---|---|
| username | str | フォーム入力されたユーザー名 |
| password | str | フォーム入力されたパスワード |
| factors | login_factor_list | ログイン時に提示された追加認証要素 |
| status | login_form_status | フォーム入力の検証状態 |
| metadata | login_metadata | フォーム送信時の補助メタデータ |

## Private models

File-private helper schemas defined in this model YAML file.
Promote a helper model to a public model file when it needs to be reused from other YAML files.

| model | kind | shape | note |
|---|---|---|---|
| login_factor | struct | kind: str<br/>value: str | login_form 内だけで使う追加認証要素。 |
| login_factor_list | list | element: login_factor | login_form 内だけで使う追加認証要素リスト。 |
| login_form_status | enum | draft<br/>submitted<br/>rejected | login_form 内だけで使う入力状態。 |
| login_metadata | dict | value: str | login_form 内だけで使う送信メタデータ。 |

