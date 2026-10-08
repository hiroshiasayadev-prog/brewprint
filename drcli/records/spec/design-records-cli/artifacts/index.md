# Overview: Design record artifacts

- **id**: `spec:drcli.design_records_cli.artifacts`
- **status**: draft
- **date**: 2026-10-07
- **parent**: `spec:drcli.design_records_cli`

## What this is

Defines the DRCLI Specification area for supported Design Record artifact kinds.

Each artifact-specific Specification set defines its artifact according to the shared definitions and templates under `base`.

## Current contract

DRCLI uses the shared base declarations and artifact-specific projections in this area to interpret supported Design Record sources.
Where a projection cites PRODUCT authority, the cited PRODUCT Specification remains the semantic source of truth.

## Topics

| title | kind | ref | summary |
|---|---|---|---|
| Artifact base | Index | `spec:drcli.design_records_cli.artifacts.base` | Shared artifact definitions and templates. |
| Specification artifact | Index | `spec:drcli.design_records_cli.artifacts.spec` | Specification artifact Specification set. |
| Requirement artifact | Index | `spec:drcli.design_records_cli.artifacts.requirement` | Requirement artifact Specification set. |
| Decision artifact | Index | `spec:drcli.design_records_cli.artifacts.decision` | Decision artifact Specification set. |
| Investigation artifact | Index | `spec:drcli.design_records_cli.artifacts.investigation` | Investigation artifact Specification set. |
| Work Item artifact | Index | `spec:drcli.design_records_cli.artifacts.work_item` | Work Item artifact Specification set. |
| Task artifact | Index | `spec:drcli.design_records_cli.artifacts.task` | Task artifact Specification set. |
