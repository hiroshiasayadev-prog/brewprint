# V01-REQ-MCP-029: third-party dependency license notice management for authoring diff library

- **id**: V01-REQ-MCP-029
- **status**: accepted
- **date**: 2026-06-07
- **source_refs**:
  - V01-REQ-MCP-027
  - V01-TASK-MCP-023-03
- **work_items**:
  - V01-WORK-MCP-024

## Requirement

brewprint MUST track and preserve required third-party license notices when adding external runtime or build dependencies to implement MCP authoring behavior.

This requirement is captured because V01-TASK-MCP-023-03 adopted `github.com/pmezard/go-difflib/difflib` for git-style unified diff generation. The dependency uses the BSD 3-Clause license, which permits source and binary redistribution but requires retaining the copyright notice, license conditions, and disclaimer. Binary redistribution also requires reproducing those notices in documentation or other materials.

The project needs an explicit policy and repository location for third-party dependency notices so future dependency additions do not rely on ad-hoc memory.

## Evidence

During V01-REQ-MCP-027 implementation, the project chose to use `github.com/pmezard/go-difflib/difflib` rather than hand-rolling complex unified diff logic.

The pasted license text for the dependency is BSD 3-Clause style:

- Copyright notice: `Copyright (c) 2013, Patrick Mezard`.
- Redistribution in source and binary forms is permitted.
- Source redistribution must retain the copyright notice, conditions, and disclaimer.
- Binary redistribution must reproduce the copyright notice, conditions, and disclaimer in documentation or other materials.
- Contributor names must not be used for endorsement without prior written permission.

This is license-compatible for the intended use, but the repository currently needs an explicit notice-management path before distribution hygiene can be considered complete.

## Required Outcome

The project has a documented, repeatable way to manage third-party dependency license notices.

Acceptance criteria:

- A repository location or policy is defined for third-party license notices, for example `THIRD_PARTY_NOTICES.md`, `NOTICE`, `docs/licenses/`, or an explicitly documented equivalent.
- `github.com/pmezard/go-difflib/difflib` is listed with its module path, license family, notice requirement summary, and pointer to upstream or module license text.
- The policy distinguishes Go module dependency metadata (`go.mod` / `go.sum`) from redistribution notice obligations.
- The policy covers both source distribution and binary distribution expectations at a minimum level.
- The policy avoids inventing endorsement language using contributor names.
- Future dependency additions have a clear checklist for notice capture.

## Explicitly Excluded Scope

- Replacing `go-difflib` with a hand-written diff implementation.
- Performing a full legal review of all historical dependencies.
- Changing the project license.
- Automating SBOM generation.
- Vendorizing dependencies unless a later work item explicitly chooses that approach.

## Boundary

This requirement belongs to the MCP/project maintenance boundary because the immediate trigger is the MCP authoring diff implementation, but the resulting notice-management policy should be reusable for future brewprint dependencies.

## Close Evidence

- V01-WORK-MCP-024 added `docs/third-party-notices.md` as the repository location for third-party dependency license notices.
- `docs/third-party-notices.md` lists `github.com/pmezard/go-difflib v1.0.0` / `github.com/pmezard/go-difflib/difflib` with module path, purpose, BSD 3-Clause style license family, notice requirement summary, and upstream license pointer.
- `docs/doc-policy.md` links to `docs/third-party-notices.md` and distinguishes `go.mod` / `go.sum` module metadata from redistribution notice obligations.
- The work stayed within the explicitly excluded scope: no dependency replacement, full legal review, project license change, SBOM automation, or vendoring was introduced.
