# Milestone 9: QueryService state/event/scenario expansion

- **status**: closed
- **scope**: QueryService / MCP
- **source**: explicit milestone marker for `docs/impl/go-m9-summary.md`
- **last_updated**: 2026-04-30

---

## Summary

M9 は、M3 QueryService の後続拡張として実施された state / event / scenario / transition / field / file selector 周辺の QueryService / MCP 拡張を指す。

TASKS上で M9 が欠番に見えると実装再開時に迷うため、このファイルは closed milestone marker として置く。

実装詳細は以下を参照する。

- `docs/impl/go-m9-summary.md`

---

## Completed scope

- state / event / scenario 系 reverse lookup index
- transition / event direct references
- inspect(state)
- inspect(event)
- inspect(sequence_diagram scenario)
- GetReferences(transition)
- GetSignature(transition)
- Inspect(transition)
- GetReferences(field)
- GetReferences(file: state_file)
- same-module bare FK normalization
- GetSignature(field)
- Inspect(field)
- `docs/spec/mcp/tools/get-signature.md` / `docs/spec/mcp/tools/inspect.md` の transition / field spec追随

---

## Notes

- M3 は QueryService の初期vertical slice。
- M9 はその後続拡張として扱う。
- 今後のMCP設計対話coverage拡張は M10〜M12 に分けて管理する。
