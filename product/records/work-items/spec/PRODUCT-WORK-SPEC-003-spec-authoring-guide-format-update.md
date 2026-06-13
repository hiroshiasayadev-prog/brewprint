# PRODUCT-WORK-SPEC-003: Spec authoring guide format update

- **id**: PRODUCT-WORK-SPEC-003
- **status**: not_started
- **date**: 2026-06-10
- **requirement_refs**:
  - PRODUCT-REQ-SPEC-001
- **investigation_refs**:
  - PRODUCT-INV-SPEC-001
- **source_work_items**:
  - PRODUCT-WORK-SPEC-001
- **task_refs**:

## Summary

Update spec authoring guidance so humans and agents can write new or migrated spec files using the accepted spec format contract.

## Scope

| area | in scope |
|---|---|
| H1 format | Explain `# <SpecKind>: <Title>` and parser-aware one-H1 expectations. |
| H1-adjacent metadata | Explain required `id/status/date/parent` markers. |
| spec ID-as-ref | Explain `spec:` ID-as-ref, underscore segment policy, `index.md` omission, and no separate `ref`. |
| spec kind selection | Provide guidance for `Overview`, `Index`, `Concept`, `Reference`, and `Contract`. |
| required sections | Provide authoring examples for required / recommended / prohibited sections. |
| `## What this is` | Provide examples and anti-patterns. |
| `## Topics` | Provide authoring examples for `title/kind/parent/file/summary`. |
| front matter policy | Explain that front matter is not the source of truth for new/migrated specs. |

## Non-scope

| area | owner |
|---|---|
| format contract changes | PRODUCT-WORK-SPEC-001 or successor work |
| ID/ref compatibility decisions | PRODUCT-WORK-SPEC-002 |
| DRMCP validation implementation | DRMCP-WORK-SPEC-001 / DRMCP-WORK-SPEC-002 |
| existing spec migration | PRODUCT-WORK-SPEC-005 |

## Dependencies

| dependency | reason |
|---|---|
| PRODUCT-WORK-SPEC-001 | Defines the accepted format. |
| PRODUCT-WORK-SPEC-002 | Recommended before final examples if derived topic / alias compatibility affects authoring. |

## Done condition

| item | done when |
|---|---|
| guide updated | Spec authoring guidance reflects the accepted format contract. |
| examples aligned | Examples match `spec-format/index.md`. |
| stale patterns avoided | Guide does not instruct authors to use front matter as topic/ref/source truth. |
| review completed | Guide update is reviewed against PRODUCT-WORK-SPEC-001 and PRODUCT-WORK-SPEC-002 decisions. |

## Source records

| ref | role |
|---|---|
| PRODUCT-REQ-SPEC-001 | Requirement for MCP-readable spec format and topic tree support. |
| PRODUCT-INV-SPEC-001 | Investigation evidence for the format. |
| PRODUCT-WORK-SPEC-001 | Defines the format contract and follow-up split. |

## Evidence

