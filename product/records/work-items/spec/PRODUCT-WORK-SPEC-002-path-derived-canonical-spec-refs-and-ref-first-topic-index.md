# PRODUCT-WORK-SPEC-002: Path-derived canonical spec refs and ref-first topic index

- **id**: PRODUCT-WORK-SPEC-002
- **status**: done
- **date**: 2026-06-10
- **requirement_refs**:
  - PRODUCT-REQ-SPEC-001
- **investigation_refs**:
  - PRODUCT-INV-SPEC-001
- **source_work_items**:
  - PRODUCT-WORK-SPEC-001
- **task_refs**:

## Summary

Define the path-derived canonical `spec:` ref model, ref-first `## Topics` behavior, and any exceptional compatibility behavior still needed for aliases, redirects, stale refs, and legacy semantic reference metadata.

This work exists because `spec-format` now treats location-derived refs as canonical for new and migrated specs while existing records may still contain older refs such as `spec:trace.semantic-ref`. Compatibility must be explicit follow-up behavior, not the default identity model.

## Scope

| area | in scope |
|---|---|
| path-derived canonical refs | Define path-to-ref and ref-to-path derivation as the canonical model for new and migrated specs. |
| ref-first `## Topics` | Define how topic rows use `ref` as child identity and how tooling resolves `ref` to a path. |
| ID changes on move / rename | Define that moving or renaming a spec changes the canonical spec ID. |
| mismatch severity | Define warning-only behavior for inventory, migration, or transient working states and error behavior for new/migrated specs. |
| alias / redirect exceptions | Define alias, redirect, superseded, and compatibility mapping behavior only where the canonical model is insufficient. |
| split / merge stale-ref behavior | Define exceptional compatibility behavior for document restructuring when old refs must remain resolvable. |
| hyphen / underscore compatibility | Define coexistence for older hyphen-form IDs and new underscore default IDs. |
| legacy `semantic_refs` / `sections` | Define compatibility and migration policy for legacy front matter metadata. |

## Non-scope

| area | owner |
|---|---|
| DRMCP validation implementation | DRMCP-WORK-SPEC-001 / DRMCP-WORK-SPEC-002 |
| existing spec migration | PRODUCT-WORK-SPEC-005 |
| authoring guide update | PRODUCT-WORK-SPEC-003 |
| ownership relocation | PRODUCT-INV-SPEC-002 / PRODUCT-WORK-SPEC-004 |

## Dependencies

| dependency | reason |
|---|---|
| PRODUCT-WORK-SPEC-007 | Applies the corrected split spec-format contract that reframes this work item. |
| `spec:product.concepts.spec_format.spec_id_as_ref` | Defines path-derived canonical refs and defers exceptional compatibility here. |
| `spec:product.concepts.spec_format.topics_table` | Defines ref-first `## Topics` rows and ref-to-path resolution expectations. |

## Done condition

| item | done when |
|---|---|
| canonical ref contract | Path-derived canonical spec refs and one-to-one path/ref mapping are defined. |
| ref-first topic index | `## Topics` rows use `ref` as canonical child identity and tooling resolution rules are defined. |
| rename/move behavior | Move and rename behavior is explicit: canonical IDs change with location unless an accepted compatibility exception applies. |
| compatibility contract | Alias, redirect, stale-ref, split, merge, and legacy metadata behavior are defined as exceptions if still needed. |
| hyphen/underscore handling | Older IDs and new path-derived canonical refs have explicit coexistence or migration rules. |
| downstream gates | DRMCP validation and migration work can use the result without guessing. |
| task split | If the scope grows, alias/redirect, stale-ref, derived-topic-ref, and legacy compatibility are split into tasks. |

## Source records

| ref | role |
|---|---|
| PRODUCT-REQ-SPEC-001 | Requirement for MCP-readable spec format and topic tree support. |
| PRODUCT-INV-SPEC-001 | Investigation that identified stable ref and derived topic compatibility as a key issue. |
| PRODUCT-WORK-SPEC-001 | Work item that produced the initial format contract and follow-up split. |
| PRODUCT-WORK-SPEC-007 | Work item that corrected the split contract toward path-derived canonical refs and ref-first Topics. |

## Evidence

Investigation (2026-06-13): grep confirmed hyphen-form `spec:trace.*` refs and legacy YAML front matter exist in `product/records/spec/concepts/traceability/` and in all drmcp/bpdsl spec files. No pending splits, merges, or cross-references requiring alias/redirect were found. All compatibility classes resolved as "not needed now"; decisions recorded inline in `spec-id-as-ref.md` `## Boundary`.

Files written or updated:

| file | change |
|---|---|
| `product/records/spec/concepts/spec-format/spec-id-as-ref.md` | Removed "Until PRODUCT-WORK-SPEC-002..." interim qualifier from parent grammar. Replaced `## Boundary` deferred rows with inline "decided: not needed now" decision table. Updated `## Related specs`. |
| `product/records/spec/concepts/spec-format/follow-up-boundary.md` | Updated PRODUCT-WORK-SPEC-002 row to reflect done state. |
