# V01-REQ-MCP-001: Design Records MCP に semantic trace support を実装する

- **id**: V01-REQ-MCP-001
- **status**: accepted
- **date**: 2026-05-23
- **source_refs**:
  - V01-ADR-087
  - V01-ADR-088
  - spec:trace.resolve-and-validation
- **work_items**:
  - V01-WORK-MCP-001

## 要求

Design Records MCP は、V01-ADR-087 / V01-ADR-088 と traceability spec に従い、canonical reference resolution foundation と investigation record integration を実装として提供する必要がある。

## 必要な結果

- `decision` / `spec` に加えて `investigation` record を index / query / validate 対象として扱える。
- record response は common fields と kind-specific detail object を分離する contract に追従する。
- active `spec:` semantic ref と record ID-as-ref (`ADR-*` / `SPEC-*` / `INV-*`) の resolve が、traceability spec の canonical reference rule に従って動作する。
- investigation の `source_refs` および記載済み `follow_up_results` の unresolved を validation error として検出できる。
- `follow_up_candidates` の未解決は存在しうる候補として扱い、それ自体を error にしない。

## 境界

- 本 requirement は実装追従を要求するものであり、現行仕様そのものを所有しない。
- resolve tool の具体 request / response schema と diagnostic category の確定が必要な場合は、関連 spec を更新してから実装する。
- `internal-design:` / `coverage:` / `COV-*` の resolve、semantic realization relation、coverage mapping query、および MCP writer tools はこの requirement の必須範囲には含めない。

## 根拠

- V01-ADR-087 は semantic/artifact ref resolve と investigation integration を Design Records MCP の責務として accepted にした。
- V01-ADR-088 は MVP resolve / validation を `spec:` semantic ref、record ID-as-ref、investigation canonical references に限定した。
- M18 は docs 上の最小運用基盤までを扱い、実装追従は M19 に切り出す。
