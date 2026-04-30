---
scope: docs/spec/mcp/tools/list-endpoints.md
status: draft
last_updated: 2026-04-30
summary: >
  list_endpoints toolの仕様を定義する。
  API Table view YAMLに基づくendpoint一覧を返す。
  ADR-028のroute合成規則に従うfull pathを規定する。
depends_on:
  - docs/adr/028-api-table-route-composition.md
---

# `list_endpoints`

## 1. Purpose

`list_endpoints` は、API Table view YAMLに基づいてendpoint一覧を返す。

`task(endpoint=true)` を単純列挙するだけではなく、ADR-028のroute合成規則に従いfull pathを返す。

## 2. Input

```json
{
  "api_table_id": "ec_api"
}
```

| フィールド | 必須 | 内容 |
|---|---:|---|
| `api_table_id` | 任意 | API Table view ID。省略時はproject内の全API Tableを返す |

API Tableが複数存在し、`api_table_id` が省略された場合は、全API Tableを `tables[]` に分けて返す。

## 3. Output

```json
{
  "tables": [
    {
      "id": "ec_api",
      "http_root_path": "/api",
      "sections": [
        {
          "module": "auth",
          "include_submodules": true,
          "endpoints": [
            {
              "method": "POST",
              "path": "/api/login",
              "leaf_path": "login",
              "task": "auth.task.login",
              "params": "auth.model.login_request",
              "returns": "auth.model.token",
              "source": {
                "file": "auth/task/login.yaml"
              }
            }
          ]
        }
      ]
    }
  ],
  "diagnostics": []
}
```

この例ではsection起点moduleが `auth` のため、ADR-028のroute合成規則により、section起点moduleからの相対module pathは空になる。
そのためfull pathは `/api/auth/login` ではなく `/api/login` になる。
`/api/auth/login` を返したい場合は、API Table view側で `http_root_path: /api/auth` とするか、section起点moduleを上位moduleにする。

## 4. endpoint object

| フィールド | 必須 | 内容 |
|---|---:|---|
| `method` | ✓ | HTTP method |
| `path` | ✓ | API Table viewにより合成されたfull path |
| `leaf_path` | ✓ | task側のleaf path。省略時はtask.id由来 |
| `task` | ✓ | endpoint task QualifiedID |
| `params` | 任意 | request model QualifiedID |
| `returns` | 任意 | response model QualifiedID |
| `source` | 任意 | endpoint taskのSourceLocation |

---
