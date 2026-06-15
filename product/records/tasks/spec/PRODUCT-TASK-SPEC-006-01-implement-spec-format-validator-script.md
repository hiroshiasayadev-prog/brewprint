# PRODUCT-TASK-SPEC-006-01: Implement spec format validator script

- **id**: PRODUCT-TASK-SPEC-006-01
- **status**: done
- **date**: 2026-06-16
- **work_item**: PRODUCT-WORK-SPEC-006
- **source_requirement**: PRODUCT-REQ-SPEC-001
- **estimate**: 0.5d
- **depends_on**:
- **outputs**:
  - `product/src/tools/validate_spec.py`

## Goal

Implement the temporary PRODUCT-level spec format validator as a standalone Python script. The script must cover all validation checks needed for PRODUCT-WORK-SPEC-005 migration review without patching the current DRMCP codebase.

## Work

| area | required work |
|---|---|
| path-derived ID resolver | Given a spec file path, compute the canonical `spec:` ID (index.md omission, hyphen-to-underscore, prefix from app namespace). |
| parser-aware H1 counter | Count real ATX H1s while ignoring YAML front matter and fenced code blocks. |
| H1 format check | Validate `# <SpecKind>: <Title>` against accepted kind set. |
| H1-adjacent metadata | Check presence of `id`, `status`, `date`, `parent`; `contract_class` required/prohibited per kind. |
| ID mismatch check | Compare declared `id` against path-derived canonical ref; report mismatch. |
| parent grammar check | Validate parent value against `root | - | spec:<...>` grammar. |
| section matrix checks | Check required/prohibited sections by kind and contract_class. |
| Topics table checks | Check required columns (`title`, `kind`, `ref`, `summary`); flag `file` used as canonical column; detect duplicate refs. |
| inventory vs strict mode | `--strict` escalates migration-phase warnings to errors; default is inventory (warning) mode. |
| output form | Standalone Python script; no DRMCP dependency; no MCP surface. |

## Done condition

| item | done when |
|---|---|
| resolver available | `derive_spec_id(path, repo_root)` returns correct canonical ref for all spec tree cases. |
| validator available | All checks above implemented and exercised. |
| diagnostics usable | Output is file-grouped with ERROR/WARN/INFO labels and diagnostic codes. |
| boundary clean | No DRMCP implementation code is changed. |

## Verification

- Inventory mode: existing pre-migration specs produce only WARNs, not ERRORs.
- New-format specs (`product/records/spec/concepts/spec-format/`, `authoring-standards/`) pass clean in both inventory and strict mode.
- Strict mode: pre-migration specs escalate to ERRORs.
- Exit code: 0 if no errors, 1 if any errors.
- ID derivation verified against spec examples from `spec:product.concepts.spec_format.spec_id_as_ref`.

## Evidence

- Created `product/tools/validate_spec.py`.
- Inventory run: 10 pre-migration spec files produce 49 warnings (YAML front matter, hidden front matter keys, H1 format); 0 errors.
- New-format specs (10 files): 0 issues in both inventory and strict mode.
- Strict run on all files: same 49 diagnostics escalate to errors, confirming strict escalation path works.
- ID derivation spot-check: `index.md` omission and hyphen-to-underscore normalization match spec examples.
- No DRMCP implementation files changed.
- No `v01/records/**` files changed.
