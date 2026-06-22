# Contract: API Table render rules

- **id**: `spec:bpdsl.views.api_table`
- **status**: draft
- **date**: 2026-06-17
- **parent**: `spec:bpdsl.views.overview`
- **contract_class**: `format`

## What this is

API Table is an Application-layer view that gives humans and LLMs an overview list of all tasks with `endpoint: true`. It has no Mermaid rendering — it renders as Markdown and is also exposed via the `list_endpoints` MCP tool.

API Table is not an independent node and not an implicit aggregation either. It has its own YAML, which explicitly states which modules are aggregated, how, and under which HTTP root they are shown.

> Source: V01-ADR-002, V01-ADR-005, V01-ADR-009, V01-ADR-017, V01-ADR-028, V01-ADR-030

## Current contract

### YAML schema

```yaml
as: api_table
id: auth_api
note: Authentication API listing
http_root_path: /api
modules:
  - module: app.auth
    include_submodules: true
  - module: app.admin.audit
    include_submodules: false
```

| field | required | type | content |
|---|---|---|---|
| `as` | ✓ | string | File-kind declaration, used to identify a view definition file. Fixed: `api_table` (V01-ADR-030). |
| `id` | ✓ | string | API Table identifier. Used as the Markdown H1. |
| `note` | optional | string | Description of the API Table as a whole. Output directly under the H1. |
| `http_root_path` | ✓ | string | HTTP root path this API Table is responsible for, e.g. `/api` (includes the leading `/`). |
| `modules` | ✓ | list\<module-entry\> | List of modules to aggregate. |

### module-entry object

| field | required | type | content |
|---|---|---|---|
| `module` | ✓ | string | Absolute module path. |
| `include_submodules` | optional | bool | When `true`, also collects endpoint tasks from submodules. Defaults to `false`. |

### Output format

```markdown
# {id}

{note}

## Routes

- [{module-1}](#{module-1-anchor})
- [{module-2}](#{module-2-anchor})

## {module-1}

| task id | method | path | params | returns |
|---|---|---|---|---|
| login | POST | /api/login | credential | token |
| oauth/start | GET | /api/oauth/start | - | oauth_redirect |

## {module-2}

| task id | method | path | params | returns |
|---|---|---|---|---|
| list_logs | GET | /api/admin/audit/list_logs | - | audit_log_list |
```

- H1 is `id`.
- If `note` is present, it is output directly under the H1; omitted otherwise.
- `## Routes` comes first, listing links to each section.
- The body lists `##` sections in `modules[]` order.
- Each section has the endpoint table for its corresponding module-entry.

### table column rendering

| column | content |
|---|---|
| `task id` | Path relative to the section's anchor module (absolute path if needed). |
| `method` | The endpoint task's `method`. |
| `path` | The full route composed by the API Table (see §Route composition rule). |
| `params` | `params[].model` listed in declaration order, joined by `<br/>`. |
| `returns` | `returns.model` (single value). |

Missing values render as `-` for both `params` (absent) and `returns` (absent).

`params` example: declaring

```yaml
params:
  - name: refresh_token
    model: refresh_token
  - name: client_info
    model: client_info
```

renders as `refresh_token<br/>client_info`.

`returns` is single-valued; only `returns.model` is shown. Example: `returns: { name: auth_token, model: token }` renders as `token`.

### Render example

YAML:

```yaml
as: api_table
id: auth_api
note: Authentication API listing
http_root_path: /api
modules:
  - module: app.auth
    include_submodules: true
```

Target endpoint tasks:

```yaml
# app.auth
- id: login
  type: task
  endpoint: true
  method: POST
  path: login
  params:
    - name: credentials
      model: credential
  returns:
    name: auth_token
    model: token

# app.auth.oauth
- id: start
  type: task
  endpoint: true
  method: GET
  path: start
  returns:
    name: oauth_redirect
    model: oauth_redirect

# app.auth.oauth
- id: callback
  type: task
  endpoint: true
  method: GET
  path: callback
  params:
    - name: code
      model: oauth_code
    - name: state
      model: oauth_state
  returns:
    name: auth_token
    model: token
```

Markdown output:

```markdown
# auth_api

Authentication API listing

## Routes

- [app.auth](#appauth)

## app.auth

| task id | method | path | params | returns |
|---|---|---|---|---|
| login | POST | /api/login | credential | token |
| oauth/callback | GET | /api/oauth/callback | oauth_code<br/>oauth_state | token |
| oauth/start | GET | /api/oauth/start | - | oauth_redirect |
```

## Rules

### Collection rule

Only tasks satisfying all of the following are included in the API Table:

- `type: task`
- `endpoint: true`
- Belongs to a module listed in `modules[]`.
- If `include_submodules: true`, submodules under that module are included too.

A task without `endpoint: true` is excluded entirely.

### Per-module sections

The API Table output has one section per `modules[]` entry.

- If `include_submodules: false`, only endpoint tasks directly under that module are targeted.
- If `include_submodules: true`, endpoint tasks of submodules are included in the same section — submodules never get their own separate section.
- A module-entry with zero collected endpoints does not output a section. This is not a parser error or warning.

Example: with `app.auth` as the section's anchor module and `app.auth.oauth` included, `login`, `oauth/start`, and `oauth/callback` all render into the same section.

### Route composition rule

An endpoint task's `path` is a leaf path, not a full path. `path` is optional; when omitted, `task.id` is used as the leaf name. `path` allows only a single segment with no `/` (e.g. `stripe`, `login`) — a multi-segment path containing `/` is invalid; URL hierarchy is expressed via the module directory structure instead.

The API Table composes each endpoint's final route from:

```text
{http_root_path}/{module path relative to the section's anchor module}/{task.path}
```

- `http_root_path` comes from the API Table YAML.
- The relative module path is computed from the section's anchor module.
- `task.path` is the endpoint task's leaf path.

Example: with `http_root_path: /api` and section anchor `app.auth` (`include_submodules: true`):

```yaml
# app.auth
- id: login
  type: task
  endpoint: true
  method: POST
  path: login

# app.auth.oauth
- id: start
  type: task
  endpoint: true
  method: GET
  path: start
```

renders routes `login` → `/api/login` and `oauth/start` → `/api/oauth/start`.

### Display name rule

A task's display name within a section uses the path relative to the section's anchor module:

- Directly under the same module: `login`.
- Under a submodule: `oauth/start`.
- If something in the same section can't be relativized from the anchor module, show its absolute path.

### Routes listing

`## Routes` enumerates heading links for each section, using the same display text as the section heading. The order matches the order sections are output, i.e. `modules[]` declaration order. A section omitted because it has zero collected endpoints is also omitted from the Routes listing. Anchor targets follow the Markdown renderer's heading-ID generation rule.

### Section heading

A section heading uses the absolute module path of `modules[].module`, e.g. `## app.auth`, `## app.admin.audit`.

### Sort order

Rows within each section are sorted by `task id` in ASCII ascending order, e.g. `login`, `oauth/callback`, `oauth/start`. The section order itself (`modules[]` order) follows YAML declaration order, not sorting.

### Exclusion rule

The following are never included in the API Table:

- A task without `endpoint: true`.
- `model`, `store`, `actor`, `event`, `state`, `branch`, `fork`, `join`.
- `asset` (derived from a task's `returns`, but not a table target).

## Validation rules

- `path` containing `/` is invalid — single-segment leaf paths only.
- A module-entry with zero collected endpoints is not an error or warning — it simply produces no section and no Routes row.
- `modules[]` entries are processed in declaration order; this order is never re-sorted by the renderer.

## Related specs

| ref | relation |
|---|---|
| `spec:bpdsl.views.overview` | Parent overview; view kind catalog. |
| `spec:bpdsl.dsl.nodes.processing` | `task` node definition including `endpoint` field. |
| `spec:bpdsl.mcp.tools.list_endpoints` | MCP tool exposing the same route-composition rule. |
