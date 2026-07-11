# USDM requirement: Lifecycle and update rules

- **id**: `usdm:product.design_records.authoring_semantics.lifecycle_and_update_rules`
- **status**: draft
- **date**: 2026-07-10
- **kind**: requirement
- **parent**: `usdm:product.design_records.authoring_semantics`

## What this is

This record defines lifecycle requirements from Product Design Records authoring rules that DRMCP can check with static cross-record validation.

## Requirements: Task authoring
> source: spec:product.design_records.authoring_standards.task_authoring

| id | requirement | notes |
|---|---|---|
| R001 | DRMCP static validation must report an active Task whose `depends_on` references a cancelled Task unless the dependent Task status is `blocked`. | Active means the dependent Task is neither `done` nor `cancelled`. |
