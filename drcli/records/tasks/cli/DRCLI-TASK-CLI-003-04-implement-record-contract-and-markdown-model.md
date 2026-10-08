# DRCLI-TASK-CLI-003-04: Implement artifact grammar and record model

- **status**: not_started
- **date**: 2026-10-08
- **work_item**: DRCLI-WORK-CLI-003
- **task_type**: implementation
- **estimate**: 3d
- **depends_on**:
  - DRCLI-TASK-CLI-003-03
- **outputs**:
  - drcli/go.mod
  - drcli/internal/model/types.go
  - drcli/internal/model/identity.go
  - drcli/internal/model/markdown.go
  - drcli/internal/model/identity_test.go
  - drcli/internal/model/markdown_test.go

## Goal

Implement artifact grammar and record model in a bounded, separately testable Go package.

## Work

### Read

- This Task: `DRCLI-TASK-CLI-003-04`.
  - `drcli/records/spec/design-records-cli/artifacts/base/definitions/h1-adjacent-metadata.md`
  - `drcli/records/spec/design-records-cli/artifacts/base/definitions/identity-declaration.md`
  - `drcli/records/spec/design-records-cli/artifacts/base/definitions/index.md`
  - `drcli/records/spec/design-records-cli/artifacts/base/definitions/record-structure.md`
  - `drcli/records/spec/design-records-cli/artifacts/base/definitions/source.md`
  - `drcli/records/spec/design-records-cli/artifacts/base/index.md`
  - `drcli/records/spec/design-records-cli/artifacts/base/templates/h1-adjacent-metadata.md`
  - `drcli/records/spec/design-records-cli/artifacts/base/templates/identity-and-structure/index.md`
  - `drcli/records/spec/design-records-cli/artifacts/base/templates/identity-and-structure/sequential.md`
  - `drcli/records/spec/design-records-cli/artifacts/base/templates/identity-and-structure/tree.md`
  - `drcli/records/spec/design-records-cli/artifacts/base/templates/index.md`
  - `drcli/records/spec/design-records-cli/artifacts/base/templates/source.md`
  - `drcli/records/spec/design-records-cli/artifacts/decision/h1-adjacent-metadata.md`
  - `drcli/records/spec/design-records-cli/artifacts/decision/identity-and-structure.md`
  - `drcli/records/spec/design-records-cli/artifacts/decision/index.md`
  - `drcli/records/spec/design-records-cli/artifacts/decision/source.md`
  - `drcli/records/spec/design-records-cli/artifacts/index.md`
  - `drcli/records/spec/design-records-cli/artifacts/investigation/h1-adjacent-metadata.md`
  - `drcli/records/spec/design-records-cli/artifacts/investigation/identity-and-structure.md`
  - `drcli/records/spec/design-records-cli/artifacts/investigation/index.md`
  - `drcli/records/spec/design-records-cli/artifacts/investigation/source.md`
  - `drcli/records/spec/design-records-cli/artifacts/requirement/h1-adjacent-metadata.md`
  - `drcli/records/spec/design-records-cli/artifacts/requirement/identity-and-structure.md`
  - `drcli/records/spec/design-records-cli/artifacts/requirement/index.md`
  - `drcli/records/spec/design-records-cli/artifacts/requirement/source.md`
  - `drcli/records/spec/design-records-cli/artifacts/spec/h1-adjacent-metadata.md`
  - `drcli/records/spec/design-records-cli/artifacts/spec/identity-and-structure.md`
  - `drcli/records/spec/design-records-cli/artifacts/spec/index.md`
  - `drcli/records/spec/design-records-cli/artifacts/spec/source.md`
  - `drcli/records/spec/design-records-cli/artifacts/task/h1-adjacent-metadata.md`
  - `drcli/records/spec/design-records-cli/artifacts/task/identity-and-structure.md`
  - `drcli/records/spec/design-records-cli/artifacts/task/index.md`
  - `drcli/records/spec/design-records-cli/artifacts/task/source.md`
  - `drcli/records/spec/design-records-cli/artifacts/work-item/h1-adjacent-metadata.md`
  - `drcli/records/spec/design-records-cli/artifacts/work-item/identity-and-structure.md`
  - `drcli/records/spec/design-records-cli/artifacts/work-item/index.md`
  - `drcli/records/spec/design-records-cli/artifacts/work-item/source.md`
- Consume only the completed predecessor exported APIs named below; do not reread siblings or change their code.

### Change

- `drcli/go.mod`
- `drcli/internal/model/types.go`
- `drcli/internal/model/identity.go`
- `drcli/internal/model/markdown.go`
- `drcli/internal/model/identity_test.go`
- `drcli/internal/model/markdown_test.go`

### Implement

- Fixed private Go API: `model.Kind, model.Record, model.MetadataField, model.Section, model.Source, model.TreeNode; ParseMarkdown([]byte) (Document,error); ParseSequentialID(string,Kind) (Identity,error); TreeRef(kind,app,relativePath string) (string,error); ArtifactRule(Kind) Rule; ValidateArtifactShape(Document,Rule) []Finding.`
- Fixed behavior: Use PRODUCT-anchored six artifact kinds: ADR, REQ, WORK, TASK, INV, and tree spec. H1 is the only sequential identity; tree path is the only tree identity. Preserve UTF-8 source bytes, H1 first-line title, exact H2 slices, metadata scalar/inline/indented forms, kind-specific field, heading, segment and filename-projection rules. A bad metadata field never repairs H1 identity; a mismatched tree metadata id is a finding, not an identity source. Parse without source mutation.
- Fix `drcli/go.mod` as `module github.com/hiroshiasayadev-prog/brewprint/drcli` with `go 1.22`, a standalone nested module and standard-library-only initial implementation. Do not edit repository root `go.mod` or `go.sum`.
- Model routing: Codex/Sonnet: format and identity semantics.
- Preserve exact canonical refs, error/warning order, null-versus-omitted fields, and accepted PRODUCT authority; no inferred aliases or behavior.

### Test

- Assertions: Test sequential identity work/task segment widths and order, invalid H1 and misplaced source, nonidentity filename/metadata faults, tree root/index/leaf identity, all six artifact field allow-lists and H2 inventories, duplicate headings, CRLF and UTF-8 preserved source bytes.
- Focused command: `go -C drcli test ./internal/model -count=1`; require exit 0.
- Aggregate owner: DRCLI-TASK-CLI-003-18. Independent code verdict: DRCLI-TASK-CLI-003-19.

### Stop

- Return BLOCKED with exact missing input, field, symbol, or owning Task for absent predecessor output, a needed out-of-Change edit, a contract conflict, or a failed environment command. Do not infer missing semantics.

### Prohibited operations

- No modifications to PRODUCT, DRMCP, accepted DRCLI Specifications or closed DRCLI-WORK-CLI-002; no other executor's files, no repository-wide traversal, no independent review, no lifecycle synchronization, and no stage/commit.

### Output

- Record changed files, focused command and exit code, blocking diagnostics if any, and handoff to T18.

## Implementation contract

| target | required change | acceptance criterion | verification |
|---|---|---|---|
| `drcli/go.mod` | Implement artifact grammar and record model API and its named ownership boundary. | Use PRODUCT-anchored six artifact kinds: ADR, REQ, WORK, TASK, INV, and tree spec. | `go -C drcli test ./internal/model -count=1` |
| `drcli/internal/model/types.go` | Implement artifact grammar and record model API and its named ownership boundary. | Use PRODUCT-anchored six artifact kinds: ADR, REQ, WORK, TASK, INV, and tree spec. | `go -C drcli test ./internal/model -count=1` |
| `drcli/internal/model/identity.go` | Implement artifact grammar and record model API and its named ownership boundary. | Use PRODUCT-anchored six artifact kinds: ADR, REQ, WORK, TASK, INV, and tree spec. | `go -C drcli test ./internal/model -count=1` |
| `drcli/internal/model/markdown.go` | Implement artifact grammar and record model API and its named ownership boundary. | Use PRODUCT-anchored six artifact kinds: ADR, REQ, WORK, TASK, INV, and tree spec. | `go -C drcli test ./internal/model -count=1` |
| `drcli/internal/model/identity_test.go` | Add focused table-driven regression cases. | Test sequential identity work/task segment widths and order, invalid H1 and misplaced source, nonidentity filename/metadata faults, tree root/index/leaf identity, all six artifact field allow-lists and H2 inventories, duplicate headings, CRLF and UTF-8 preserved source bytes.. | `go -C drcli test ./internal/model -count=1` |
| `drcli/internal/model/markdown_test.go` | Add focused table-driven regression cases. | Test sequential identity work/task segment widths and order, invalid H1 and misplaced source, nonidentity filename/metadata faults, tree root/index/leaf identity, all six artifact field allow-lists and H2 inventories, duplicate headings, CRLF and UTF-8 preserved source bytes.. | `go -C drcli test ./internal/model -count=1` |

## Done condition

- All declared API, behavior, and test-contract rows pass without modifying other owners. `go -C drcli test ./internal/model -count=1` exits 0.

## Verification

- Run `go -C drcli test ./internal/model -count=1` and inspect only declared changed files and test assertions. T18 owns full integration.

## Evidence

TBD
