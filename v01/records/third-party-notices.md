# Third-party notices

This document tracks third-party dependency license notices that must be preserved for source or binary redistribution hygiene.

`go.mod` / `go.sum` record Go module dependency metadata, but they do not replace license notice obligations.

## Notice capture checklist

When adding an external runtime or build dependency, record:

- dependency name / module path
- purpose in brewprint
- license family
- notice requirement summary
- upstream license pointer

Do not copy full license text into this document unless a later requirement explicitly needs it.

## Go dependencies

### github.com/pmezard/go-difflib v1.0.0

- Module: `github.com/pmezard/go-difflib`
- Package: `github.com/pmezard/go-difflib/difflib`
- Used for: Design Records MCP git-style unified diff generation.
- License: BSD 3-Clause style.
- Notice requirement: retain/reproduce the upstream license notice and disclaimer for source or binary redistribution.
- Upstream license: module license text for `github.com/pmezard/go-difflib` v1.0.0.
