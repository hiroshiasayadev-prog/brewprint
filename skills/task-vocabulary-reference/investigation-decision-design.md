# Task vocabulary reference: investigation / decision / design

General reference of real phrasing used by `investigation` Tasks (and closely related decision/design framing) in this repository's corpus. Not a boundary-violation list — see `skills/task-boundary-vocabulary/` for that.

Source: `PRODUCT-TASK-SPEC-025-06` through `-10` Finding logs, corpus-range extraction, 2026-07-03.

- `Distinguish [X] from [Y] findings` — separate responsibility domains so audit findings are not generation instructions (`PRODUCT-TASK-SPEC-013-01`)
- `Define the minimum [T06] verification cases` — specify bounded acceptance cases for a downstream implementation Task (`PRODUCT-TASK-SPEC-013-05`)
- `Classify each pointer as [six named outcomes]` — categorize each pointer without changing its owner (`PRODUCT-TASK-SPEC-015-01`)
- `Resolve facts already fixed by accepted authority without asking the user` — derive accepted facts rather than adopt a new decision (`PRODUCT-TASK-SPEC-016-01`, `-017-01`)
- `one explicit disposition` — route each source file or mixed section to one action, owner, or decision gate (`PRODUCT-TASK-SPEC-012-01`)
- `follow-up destination or deletion rationale` — assign every unadopted mechanism a terminal disposition (`PRODUCT-TASK-SPEC-012-05`)
- `four-way exit classification` — assign every staged item to one accepted exit route (`PRODUCT-TASK-SPEC-012-09`)
- `Produce one bounded implementation impact Investigation` — investigate repository seams, constraints, writers, and candidates without adopting choices (`PRODUCT-TASK-SPEC-021-03`)
- `Create one bounded Investigation of repository impact and conflicts` — investigate direct consumers, mismatches, owners, and writer conflicts (`PRODUCT-TASK-SPEC-023-02`)
- `Produce one Investigation that verifies [checklist coverage, authority alignment, ...]` — investigate one bounded question and record the result in an Investigation artifact (`PRODUCT-TASK-SPEC-020-03`)

Note: the corpus mostly expresses `investigation` responsibility with an explicit verb (`Investigate`, `Produce one ... Investigation`); disguised investigation phrasing is rarer than for `decision` or `correction` in this scan.
