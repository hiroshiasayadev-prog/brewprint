# Reference: Discovery

- **id**: `spec:drmcp.design_records_mcp.schema.discovery`
- **status**: draft
- **date**: 2026-06-27
- **parent**: `spec:drmcp.design_records_mcp.schema.overview`

## What this is

Defines DRMCP-specific index inclusion conditions and source-format boundary rules per record kind.

Kind-level path-pattern conventions are defined by `spec:product.design_records.repository_layout.record_discovery_paths`. This specification defines how DRMCP applies those patterns and what happens when a candidate is invalid.

## Current spec inclusion

### Candidate paths

Every regular Markdown file under `<records_root>/spec/**/*.md` is a current spec candidate.

- Discovery is recursive to arbitrary depth below `spec/`.
- Both `index.md` files and non-index leaf files are candidates.
- Non-Markdown files are outside current spec discovery.
- Auxiliary or non-spec Markdown files are not allowed inside the current spec tree.
- Symlinked files are not current spec candidates.
- Symlinked directories and filesystem aliases are not traversed for current spec discovery.
- A canonicalized source that resolves outside the configured spec tree is not adopted through the visible alias path.

The path pattern is consumed from `spec:product.design_records.repository_layout.record_discovery_paths`. This specification defines DRMCP-specific inclusion and boundary rules only.

### Candidate versus valid source

A current spec candidate is not automatically a valid current source.

A candidate becomes a valid current source when:

- it is a regular file (not a symlink);
- it contains exactly one real ATX H1 outside fenced code blocks;
- the H1 uses the format `# <SpecKind>: <Title>`;
- the extracted spec kind is in the PRODUCT accepted spec kind set defined by `spec:product.design_records.spec_format.document_shape`;
- it carries a valid H1-adjacent metadata block;
- the path-derived canonical ref is unique across all configured current roots;
- it does not contain YAML front matter.

A candidate that fails any of these conditions is an invalid current source. It is never silently excluded from the discovery scope. Missing or malformed metadata does not reclassify a candidate as a non-spec file. Unknown and deferred spec kinds fail the valid-source condition; DRMCP does not replicate the accepted kind list and does not validate against deferred-kind names independently — it consumes the PRODUCT accepted-kind set as the sole authority. A candidate with an unknown or deferred kind but a unique path-derived canonical ref remains addressable through that ref and remains a validation input.

### YAML front matter

A current spec candidate that contains YAML front matter is a current source-format violation.

- DRMCP does not read YAML front matter as a metadata source.
- DRMCP does not ignore YAML front matter and continue with H1-adjacent metadata.
- Matching YAML and H1-adjacent values do not make the source valid.
- Empty YAML front matter is still invalid.
- YAML values do not supplement missing H1-adjacent metadata.

This rule applies to leading YAML front matter (delimited by `---` at the start of the file). Later thematic breaks in Markdown body content are not subject to this rule.

### Invalid source behavior

Individual current spec source invalidity does not fail server startup or the complete active read surface.

A candidate with a unique path-derived canonical ref remains addressable through that ref even when its metadata, H1, or document shape is invalid. Invalid metadata, invalid H1, and invalid document shape do not remove read availability. This behavior supports inspection and repair without hiding the source.

A candidate whose canonical identity cannot be determined uniquely is not addressable through normal ID-based retrieval. This includes:

- candidates with invalid path segments (segments that do not match `[a-z0-9][a-z0-9_]*`);
- candidates whose canonical identity is conflicted by a duplicate.

Such unaddressable candidates remain validation inputs and may be reported with source provenance.

Repository validation reports all invalid sources. Exact diagnostic identifiers and severity levels are owned by the validation and diagnostics contracts (`DRMCP-WORK-MCP-006`).

### Duplicate canonical identity

When two or more current spec candidates produce the same path-derived canonical ref, that canonical identity becomes conflicted.

- A conflicted identity has no single addressable active-index entry.
- Both conflicting sources remain validation inputs with source provenance.
- Unrelated uniquely identified records remain available to normal read operations.
- Filesystem traversal order never selects a duplicate winner.
- A topic index may coexist with child specs below it when each child derives a distinct canonical ref.

A leaf spec and a topic `index.md` may produce the same canonical ref when the leaf file stem matches the directory name. This is also a conflicted identity subject to the rules above.

Exact diagnostic identifiers and read response representation for conflicted identities are owned by `DRMCP-WORK-MCP-006` and `DRMCP-WORK-MCP-004`.

### No legacy alias fallback

Current spec identity is path-derived. The following do not provide fallback identity or alternate lookup keys:

- `SPEC-*` bare IDs;
- `V01-SPEC-*` public IDs;
- metadata `id` values that do not match the path-derived canonical ref;
- any alias or redirect not defined by an explicit compatibility design.

## Other record kind inclusion

Every regular Markdown file that matches the PRODUCT discovery path for a sequential record kind is a candidate for that kind.

| kind | candidate condition |
|---|---|
| `decision` | Matches the PRODUCT ADR discovery path, including the PRODUCT-defined flat ADR compatibility pattern. |
| `investigation` | Matches the current investigation discovery path. |
| `requirement` | Matches the current requirement discovery path. |
| `work_item` | Matches the current work-item discovery path. |
| `task` | Matches the current task discovery path. |

Path patterns are consumed from `spec:product.design_records.repository_layout.record_discovery_paths`.
H1 or metadata validity is not a candidate-inclusion gate.

For these sequential candidates:

- a valid complete canonical ID in H1 establishes canonical identity;
- invalid or missing metadata makes the candidate invalid but does not remove addressability while H1 identity remains determinable and unique;
- metadata `id` on requirement, work-item, and task records is a required consistency value, not a fallback identity source;
- filename ID prefixes are consistency values, not fallback identity sources;
- a candidate without a valid complete canonical ID in H1 remains validation-only;
- duplicate canonical identity creates no winner;
- every invalid or conflicting candidate retains source provenance for validation and repair diagnostics.

Shared addressability behavior is defined by `spec:drmcp.design_records_mcp.schema.record_model` and `spec:drmcp.design_records_mcp.schema.id_normalization`.

## Explicit exclusions

This specification does not define:

- H1-adjacent metadata block grammar or field normalization (see `spec:drmcp.design_records_mcp.schema.metadata_grammar`);
- path-derived canonical ref derivation steps (see `spec:drmcp.design_records_mcp.schema.id_normalization`);
- exact diagnostic identifiers or validation response shapes (see `DRMCP-WORK-MCP-006`);
- read response representation for invalid or conflicted sources (see `DRMCP-WORK-MCP-004`);
- query filters, ordering, ranges, or pagination;
- resolver fallback order;
- fixture design or runtime implementation.

## Related specs

| ref | relation |
|---|---|
| `spec:product.design_records.repository_layout.record_discovery_paths` | Authority for record path patterns by kind. |
| `spec:product.design_records.spec_format.document_shape` | Authority for current spec H1, metadata block, and document shape requirements. |
| `spec:product.design_records.spec_format.spec_id_as_ref` | Authority for path-derived canonical spec ref derivation. |
| `spec:drmcp.design_records_mcp.namespace_scanning` | Configured current roots and active-index construction contract. |
| `spec:drmcp.design_records_mcp.schema.record_source` | DRMCP source representation for discovered records. |
| `spec:drmcp.design_records_mcp.schema.metadata_grammar` | H1-adjacent metadata block grammar. |
| `spec:drmcp.design_records_mcp.schema.id_normalization` | Path-derived canonical ref derivation steps. |

## Sources

- `DRMCP-TASK-MCP-003-03`: Accepted decisions 1, 9–12, 16–17, 20.
- `DRMCP-TASK-MCP-003-02`: Accepted duplicate canonical identity handling.
- `spec:product.design_records.repository_layout.record_discovery_paths`: Path pattern authority.
- `spec:product.design_records.spec_format.document_shape`: Document shape and H1 requirements.
