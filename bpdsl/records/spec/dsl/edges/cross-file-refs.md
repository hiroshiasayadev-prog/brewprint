# Reference: Cross-file references

- **id**: `spec:bpdsl.dsl.edges.cross_file_refs`
- **status**: draft
- **date**: 2026-06-16
- **parent**: `spec:bpdsl.dsl.edges.overview`

## What this is

Rules for referencing nodes that live in other modules. Within-module main nodes are resolved by bare ID; cross-module references require a QualifiedID (V01-ADR-003).

## Reference syntax

| scope | syntax | example |
|---|---|---|
| Same-file sub-node or file-private source | bare ID (file-private resolution first) | `validate` |
| Same-module main node | bare ID (module fallback) | `validate` |
| Cross-module node | QualifiedID (full path) | `auth.task.login` |

```yaml
# Same-module reference (inside auth/task/login.yaml)
flow:
  - step: validate        # resolves to auth.task.validate

# Cross-module reference
transitions:
  - from: idle
    on: login_submitted
    to: loading
    action: auth.task.login   # full QualifiedID
```

Bare ID resolution order: same-file file-private sub-node / source first, then same-module main node fallback. See [`spec:bpdsl.dsl.naming`](../naming.md) §Bare ID resolution for full lookup rules.

## Related specs

| ref | relation |
|---|---|
| `spec:bpdsl.dsl.edges.overview` | Parent overview; edge kind summary. |
| `spec:bpdsl.dsl.naming` | QualifiedID format and bare ID resolution rules. |
