# Third-party notices

This document tracks third-party dependency license notices that must be preserved for source or binary redistribution hygiene.

`go.mod` / `go.sum` record Go module dependency metadata, but they do not replace license notice obligations. When adding an external runtime or build dependency, add or update the relevant notice entry here before treating distribution hygiene as complete.

## Notice capture checklist

For each external runtime or build dependency, capture at minimum:

- dependency name and module path
- usage / purpose in brewprint
- license family
- copyright notice
- source redistribution notice expectation
- binary redistribution notice expectation
- endorsement restriction, when present
- pointer to the upstream license text or vendored/module license text

Do not invent endorsement wording using contributor or project names.

## Go dependencies

### github.com/pmezard/go-difflib v1.0.0

- Module path: `github.com/pmezard/go-difflib`
- Imported package: `github.com/pmezard/go-difflib/difflib`
- Used by: Design Records MCP authoring update diff generation for git-style unified diffs.
- Triggering records: REQ-MCP-027, TASK-MCP-023-03, REQ-MCP-029.
- License family: BSD 3-Clause style.
- Copyright notice: `Copyright (c) 2013, Patrick Mezard`.
- Source redistribution: retain the copyright notice, license conditions, and disclaimer.
- Binary redistribution: reproduce the copyright notice, license conditions, and disclaimer in documentation or other materials provided with the distribution.
- Endorsement restriction: contributor names must not be used to endorse or promote derived products without prior written permission.
- License text pointer: upstream module license text for `github.com/pmezard/go-difflib` v1.0.0.

