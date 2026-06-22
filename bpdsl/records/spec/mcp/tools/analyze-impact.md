# Contract: `analyze_impact`

- **id**: `spec:bpdsl.mcp.tools.analyze_impact`
- **status**: draft
- **date**: 2026-06-17
- **parent**: `spec:bpdsl.mcp.overview`
- **contract_class**: `interface`

## What this is

`analyze_impact` takes a target object and a change kind (`change`), and returns the impact range as an **interpreted impact list**.

On top of the raw reference information `get_references` / `get_reference_tree` return, it layers per-change-kind interpretation, severity judgment, fixability, and recommended action.

Returns: `summary` (counts by severity/kind/fixability), `impacts[]` (interpreted entries per affected location), `coverage` (what was analyzed / not analyzed), `assumptions` (the tool's premises/limits), `truncated` (truncation info from the return cap), `diagnostics`.

Does not return: the full YAML snippet (that's `get_source`'s responsibility), the raw reference graph (that's `get_reference_tree`'s responsibility), presentation detail inside rendered output markdown, semantic-contract compatibility judgment.

> Source: V01-ADR-056 §Decision §1

## Request

```json
{
  "selector": {
    "object": "field",
    "id": "auth.model.user.email"
  },
  "change": {
    "kind": "rename",
    "new_id": "auth.model.user.email_address"
  },
  "scope_modules": ["auth"],
  "max_impacts": 200
}
```

| field | required | content |
|---|---:|---|
| `selector` | ✓ | Target object selector. |
| `change` | ✓ | Change kind — a discriminated object (see §`change` discriminated object). |
| `scope_modules` | optional | Module list narrowing the analysis scope. |
| `max_impacts` | optional | Impact return cap. Defaults to `200`. |

There is no `depth` field in the request. The traversal depth of impact exploration is decided per change kind by the tool itself.

The selector's object/kind support range is governed by the selector support matrix in [`spec:bpdsl.mcp.schema`](../schema.md). If `analyze_impact` receives a selector marked `no` in the matrix, it does **not** return a tool error — it returns a normal response with empty `impacts`, a `coverage`, and an `unsupported_selector` diagnostic.

> Source: V01-ADR-056 §Decision §2

## `change` discriminated object

`change` is an object discriminated by `kind`. v1 handles the following kinds.

```jsonc
// rename
{ "kind": "rename", "new_id": "..." }

// remove
{ "kind": "remove" }

// change_type (a field's type change)
{ "kind": "change_type", "new_type": "auth.model.account" }

// change_contract (structural change to a task's params / returns etc.)
{ "kind": "change_contract", "note": "add session_id to params" }

// change_transition_target (change to a transition's to / action)
{
  "kind": "change_transition_target",
  "new_to": "order.state.shipped",
  "new_action": "order.task.notify_user"
}

// add (new addition — more a consistency analysis than existing-impact)
{ "kind": "add", "added_id": "auth.model.user.locale" }
```

When `add` is given, `coverage.analyzed` differs from the other kinds (see §coverage).

### Validation

Required payload per `change.kind`:

| kind | required payload | optional payload | validation |
|---|---|---|---|
| `rename` | `new_id` | - | `invalid_change_payload` if `new_id` is empty. |
| `remove` | - | - | Extra payload is not ignored — it's `invalid_change_payload`. |
| `change_type` | `new_type` | - | `invalid_change_payload` if `new_type` is empty. |
| `change_contract` | - | `note` | Valid even with no payload. |
| `change_transition_target` | - | `new_to`, `new_action` | At least one of `new_to` / `new_action` is required. |
| `add` | `added_id` | - | `invalid_change_payload` if `added_id` is empty. |

An invalid kind/payload combination is a `invalid_change_payload` tool error. An unsupported selector is not a tool error — it returns a normal response with empty `impacts`, an `unsupported_selector` diagnostic, and `coverage`.

> Source: V01-ADR-056 §Decision §2

## Response

```json
{
  "target": {
    "object": "field",
    "id": "auth.model.user.email"
  },
  "change": {
    "kind": "rename",
    "new_id": "auth.model.user.email_address"
  },
  "summary": {
    "by_severity": {
      "breaking": 3,
      "warning": 5,
      "info": 2
    },
    "by_fixability": {
      "mechanical": 3,
      "suggested": 2,
      "manual_review": 3,
      "unknown": 2
    },
    "by_kind": {
      "field_consumer": 4,
      "transition_action": 2,
      "render_output": 1
    }
  },
  "impacts": [
    {
      "id": "impact-001",
      "kind": "field_consumer",
      "severity": "breaking",
      "fixability": "mechanical",
      "object": {
        "object": "node",
        "kind": "task",
        "id": "auth.task.login"
      },
      "reason": "auth.task.login reads field 'email', so reference resolution will break after the rename",
      "via": ["reads", "model_field"],
      "source": {
        "file": "auth/task/login.yaml",
        "line": 42,
        "column": 7,
        "end_line": 42,
        "end_column": 18
      },
      "recommended_action": "Update the reads field reference to email_address",
      "suggested_fixes": [
        {
          "kind": "replace_reference",
          "confidence": "high",
          "from": "email",
          "to": "email_address",
          "source": {
            "file": "auth/task/login.yaml",
            "line": 42,
            "column": 7
          }
        }
      ]
    }
  ],
  "coverage": {
    "analyzed": [
      "direct_references",
      "reference_tree",
      "model_field_resolution",
      "transition_action_resolution",
      "type_signature_identity",
      "render_output_files"
    ],
    "not_analyzed": [
      "type_structural_compatibility",
      "semantic_contract_compatibility",
      "render_presentation_details",
      "wireframe_element_binding"
    ],
    "note": "v1 judges only reference-path reachability and type-signature identity."
  },
  "assumptions": [
    "ID collisions after rename are not validated",
    "Natural-language references inside notes are not analyzed"
  ],
  "truncated": false,
  "truncated_reasons": [],
  "diagnostics": []
}
```

| field | required | content |
|---|---:|---|
| `target` | ✓ | Target ObjectRef. |
| `change` | ✓ | The `change` given in the request, returned as-is. |
| `summary` | ✓ | Counts by severity / fixability / kind. |
| `impacts` | ✓ | List of affected locations. |
| `coverage` | ✓ | Explicit statement of analysis scope. |
| `assumptions` | ✓ | The tool's premises/limits. |
| `truncated` | ✓ | Whether truncated by `max_impacts`. |
| `truncated_reasons` | ✓ | Truncation reason. |
| `diagnostics` | ✓ | Diagnostic list. |

## Impact entry

Each `impacts[]` entry has the shape:

```jsonc
{
  "id": "impact-NNN",
  "kind": "field_consumer",
  "severity": "breaking",
  "fixability": "mechanical",
  "object": { /* ObjectRef */ },
  "reason": "...",
  "via": ["reads", "model_field"],
  "source": { /* SourceLocation */ },
  "recommended_action": "...",
  "suggested_fixes": [ /* SuggestedFix[] */ ]
}
```

| field | required | content |
|---|---:|---|
| `id` | ✓ | Unique impact ID within the response. |
| `kind` | ✓ | Semantic kind of the impact. |
| `severity` | ✓ | See §severity / fixability. |
| `fixability` | ✓ | See §severity / fixability. |
| `object` | ✓ | Affected ObjectRef, or a render output. |
| `reason` | ✓ | Natural-language explanation of why it's affected. |
| `via` | ✓ | Reference-kind path reached from root. |
| `source` | ✓ | SourceLocation of the affected spot. |
| `recommended_action` | ✓ | Human-facing recommended action. |
| `suggested_fixes` | optional | Mechanical fix candidates (see §recommended_action / suggested_fixes). |

`via` is a lightweight representation showing only the shortest reachable path. Call [`get_reference_tree`](get-reference-tree.md) separately if full path reconstruction is needed.

> Source: V01-ADR-056 §Decision §5, §9

## severity / fixability

`severity` and `fixability` are independent axes, judged separately.

### severity

| value | meaning |
|---|---|
| `breaking` | Making the change as-is is likely to cause semantic build / validation / render / query to fail. |
| `warning` | Won't necessarily break, but meaning, reachability, display, or design intent may change. |
| `info` | Presented as related information, but no action is needed or it's low priority. |

### fixability

| value | meaning |
|---|---|
| `mechanical` | Source location and replacement content are uniquely determined; can be fixed mechanically. |
| `suggested` | A fix direction can be suggested, but assumes human review. |
| `manual_review` | Requires a design decision; the tool should not decide how to fix it. |
| `unknown` | The tool cannot judge (insufficient info, out of coverage, missing source range, etc.). |

### Typical combinations by change kind (reference)

| change kind | typical severity | typical fixability |
|---|---|---|
| `rename` | breaking | mechanical |
| `remove` | breaking | manual_review |
| `change_type` (primitive) | breaking or warning | suggested |
| `change_contract` | breaking | suggested or manual_review |
| `change_transition_target` | warning | manual_review |
| `add` | info or warning | manual_review |

This is a guideline — the implementation judges per individual situation.

> Source: V01-ADR-056 §Decision §3

## coverage

`coverage` is a required output field.

```jsonc
{
  "coverage": {
    "analyzed": [...],
    "not_analyzed": [...],
    "note": "..."
  }
}
```

`analyzed` / `not_analyzed` are string enumerations. They let an LLM distinguish "analyzed but zero results" from "not analyzed" when presenting to a human.

### v1 standard set

`coverage.analyzed` v1 standard vocabulary:

- `direct_references`
- `reference_tree`
- `model_field_resolution`
- `transition_action_resolution`
- `flow_step_task_resolution`
- `flow_param_field_resolution`
- `sequence_step_task_resolution`
- `type_signature_identity`
- `render_output_files`

`coverage.not_analyzed` v1 standard vocabulary:

- `type_structural_compatibility`
- `semantic_contract_compatibility`
- `render_presentation_details`
- `wireframe_element_binding`

### coverage subsets per change.kind

`coverage.analyzed` may be a subset depending on `change.kind` and the target object kind.

Examples:

- `change.kind = "rename"` on a field target → `direct_references`, `reference_tree`, `model_field_resolution`, `flow_param_field_resolution`, `render_output_files`.
- `change.kind = "remove"` on a task target → `direct_references`, `reference_tree`, `transition_action_resolution`, `flow_step_task_resolution`, `sequence_step_task_resolution`, `render_output_files`.
- `change.kind = "add"` on a field target → name-collision check / type resolution / writer coverage etc. — not `direct_references`.

`coverage.analyzed` minimum for v1 when `add` is given:

- `name_collision`

`type_resolution` / `writer_coverage` are future `coverage` vocabulary specific to `add`, but since M13 v1 has no collector implementation for them, they go into `coverage.not_analyzed`.

### coverage mandatory rules

- `coverage` is required even when zero analysis targets are returned.
- `not_analyzed` must always be returned as an array, even if empty (cannot be omitted).
- `note` is optional, for supplementary explanation aimed at the LLM.

> Source: V01-ADR-056 §Decision §6, §7

## recommended_action / suggested_fixes

Each impact returns a two-tier recommendation: `recommended_action` (human-facing) and `suggested_fixes[]` (mechanical candidates).

### recommended_action

A human-facing natural-language recommended action. May be somewhat abstract — it's material the LLM uses to phrase what it tells the human.

### suggested_fixes

Mechanically-fixable candidate fixes. Whether they're emitted depends on `fixability`.

| fixability | suggested_fixes |
|---|---|
| `mechanical` | May emit a concrete fix with source location. |
| `suggested` | May emit a conceptual fix with confidence; source is not required. |
| `manual_review` | Empty, or non-destructive advisory only. |
| `unknown` | Empty. |

### SuggestedFix shape

```jsonc
{
  "kind": "replace_reference",
  "confidence": "high",
  "from": "email",
  "to": "email_address",
  "source": {
    "file": "auth/task/login.yaml",
    "line": 42,
    "column": 7
  },
  "note": "..."
}
```

| field | required | content |
|---|---:|---|
| `kind` | ✓ | Fix kind (implementation's discretion, e.g. `replace_reference`, `update_param_model`). |
| `confidence` | ✓ | `high` / `medium` / `low`. |
| `source` | conditional | Required when `mechanical`; optional otherwise. |
| `note` | optional | Supplementary explanation. |

Payload like `from` / `to` is `kind`-dependent.

### Requirements for `fixability=mechanical`

`fixability=mechanical` may be returned only if **all** of the following hold:

1. The replacement-target source location is uniquely identifiable (file / line / column range).
2. The pre-replacement token is unique at the source location (no mismatched match).
3. The post-replacement token resolves unambiguously to exactly one value (no collision).
4. The post-replacement reference resolves to the same target.
5. It is a simple token replacement that doesn't change YAML structure.

If even one is unmet, downgrade to at least `suggested`. Use `manual_review` or `unknown` if uncertainty is high.

Implementations must treat these five requirements as a judgment gate. Additional heuristics per individual change kind are implementation discretion.

> Source: V01-ADR-056 §Decision §4, §5

## SourceLocation

`source` carries file / line / column inline. It does not include the full YAML snippet.

```jsonc
{
  "file": "auth/task/login.yaml",
  "line": 42,
  "column": 7,
  "end_line": 42,
  "end_column": 18
}
```

| field | required | content |
|---|---:|---|
| `file` | ✓ | Project-root-relative path. |
| `line` | conditional | 1-based line number; always returned when obtainable. |
| `column` | conditional | 1-based column number; always returned when obtainable. |
| `end_line` | optional | Range end. |
| `end_column` | optional | Range end. |

When line/column cannot be obtained, the impact itself is not dropped — `source.file` alone may be returned, in which case that impact's `fixability` should be downgraded to `unknown` or `manual_review`, and `diagnostics[]` should include `source_location_unavailable`. `fixability=mechanical` requires the file/line/column range to be uniquely identifiable.

When the full YAML snippet is needed, the LLM calls `get_source` separately. A `source_preview` short line-range field may optionally be added, but is not required by `analyze_impact` v1.

> Source: V01-ADR-056 §Decision §9

## coverage scope details

### Flow wiring handling

`analyze_impact` v1 includes flow wiring in coverage. Specifically, it analyzes:

- A flow step's task reference (`flow_step_task_resolution`).
- A flow param's field reference (`flow_param_field_resolution`).

M13 v1's `flow_param_field_resolution` is minimum scope: limited to cases where the flow param wiring's target/source/return-asset name, or the source task/join's return-model identity, can be determined to relate to the target field / field model. Structural compatibility between models, or arbitrary expression analysis, is not performed.

The implementation internally reads a path equivalent to `inspect(task).members.flow.entries`. `get_reference_tree`'s reference kinds are not extended (V01-ADR-055 is maintained).

### Sequence step handling

`analyze_impact` v1 includes sequence steps in coverage. Specifically, it analyzes:

- A sequence step's task reference (`sequence_step_task_resolution`).

### Render output handling

`analyze_impact` v1 handles render output at **file granularity**:

- Identifying the render group containing the changed target.
- Returning the corresponding render output file path.

Presentation detail inside the markdown (DAG node shape changes, ER line changes, etc.) is excluded from v1, explicit in `coverage.not_analyzed` as `render_presentation_details`.

### Type-compatibility handling

`analyze_impact` v1 covers up to **type-signature identity comparison**.

Included (`type_signature_identity`):

- Primitive type match (comparing `str` / `int` / `bool` etc.).
- Model type match (comparing model-id identity).

Not included (explicit in `not_analyzed` as `type_structural_compatibility`):

- Subtyping judgment between models.
- Structural compatibility of model fields.
- Nullable / optional / required compatibility judgment.
- Semantic-contract compatibility.

> Source: V01-ADR-056 §Decision §7, §8

## M13 v1 implementation constraints

M13 closes not by satisfying the full spec at once, but as a v1 minimum implementation that upholds a strong public contract.

Implemented in M13 v1:

- `change` discriminated-object validation.
- Normal response + `unsupported_selector` diagnostic for unsupported selectors.
- Transition-action / flow-step / sequence-step impact for task rename/remove/change_contract.
- Direct/reference-tree-based impact for field rename/remove/change_type.
- Minimum `flow_param_field_resolution` for field changes.
- Resolution check of `new_to` / `new_action` for transition `change_transition_target`.
- File-granularity render-output impact.
- `name_collision` for `add`.
- The common judgment gate for `fixability=mechanical`.

M13 v1 known limitations:

- Field rename falls to `unknown` / `manual_review` rather than `mechanical` when source line/column is insufficient.
- `fixability=mechanical` is returned only when the gate is satisfied; a rename that doesn't satisfy it is `suggested` or `unknown`.
- Flow param analysis is limited to a minimum wiring-identity judgment.
- `add`'s `type_resolution` / `writer_coverage` remain in `coverage.not_analyzed`.
- Dedicated analyze collectors for state / event / actor / store are limited in M13 v1, handled mainly via the reference/render path.

## assumptions

`assumptions` enumerates the tool's own premises/limits as strings.

Examples:

- `"ID collisions after rename are not validated"`
- `"Natural-language references inside notes are not analyzed"`
- `"Transitive flow-wiring impact is limited to depth 1"`

These are information the LLM uses to tell a human "this analysis assumes X." Implementation-driven constraints are also made explicit, so the LLM doesn't give a human false confidence.

## Selector support

`analyze_impact`'s target selectors start from the set [`get_references`](get-references.md) supports.

Supported in v1:

- `node: task`
- `node: model`
- `node: store`
- `node: state`
- `node: event`
- `node: actor`
- `transition`
- `field` / `model_field`

Unsupported in v1 (returns empty impacts with the relevant kind included in `coverage.not_analyzed`):

- `view: api_table`
- `view: er_diagram`
- `view: sequence_diagram`
- `file: *`
- `asset`
- `primitive`
- `render_index`

This list matches the `analyze_impact` column of the shared selector support matrix. When an unsupported selector arrives, `diagnostics[]` includes `unsupported_selector`.

## Errors

| code | condition |
|---|---|
| `invalid_change_payload` | `change.kind`/payload combination is invalid, or a required payload field is missing. |
| `invalid_args` | Other malformed input. |
| `not_found` | Selector does not resolve to any object. |
| `ambiguous` | Selector resolves to multiple candidates. |
| `kind_mismatch` | Resolved kind does not match `selector.kind`. |

Unsupported selectors are **not** an error for this tool — see §Request and §Selector support.

## Related specs

| ref | relation |
|---|---|
| `spec:bpdsl.mcp.overview` | Parent overview; tool catalog and selection guidance. |
| `spec:bpdsl.mcp.schema` | Selector support matrix, ObjectRef, Reference shapes. |
| `spec:bpdsl.mcp.errors` | Error code catalog. |
| `spec:bpdsl.mcp.tools.get_reference_tree` | Raw reference graph this tool interprets on top of. |
| `spec:bpdsl.mcp.tools.get_source` | Full YAML snippet retrieval, not duplicated by this tool. |
| `spec:bpdsl.mcp.versioning` | Records `analyze_impact`'s promotion to v1. |
