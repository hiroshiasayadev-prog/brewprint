# V01-TASK-DATA-001-03: V01-ADR-069 minimum を spec / implementation / tests に反映する

- **id**: V01-TASK-DATA-001-03
- **status**: done
- **date**: 2026-05-29
- **work_item**: V01-WORK-DATA-001
- **source_requirement**: V01-REQ-DATA-001
- **estimate**: 1.5d
- **depends_on**:
  - V01-TASK-DATA-001-01
- **outputs**:
  - V01-ADR-069 §10 の M15 minimum scope に対応する spec 更新
  - parser safety limit / opaque_type_ref warning の implementation と tests
  - bare any + note debt を非対象とする境界 evidence

## Goal

V01-ADR-069 が M15 v1.1.0-spec に含めると定めた TypeRef safety / debt visibility の minimum baseline を反映し、V01-ADR-067 enum minimum 実装に先立って共通 spec surface を整列する。

## Work

- V01-ADR-069 §10 と現行 `docs/spec/type-ref.md` / `docs/spec/diagnostics.md` / 関連 implementation / tests の差分を確認する。
- 次の M15 minimum scope を spec に反映する。
  - nested `list<T>` / `dict<T>` の維持
  - parser safety limit
  - anonymous inline struct の非導入
  - `opaque_type_ref` warning
- parser safety limit と `opaque_type_ref` warning の implementation / tests を追加または更新する。
- `opaque_type_ref` は container TypeRef 内の `any` を対象とし、bare `any + note` の主要 response shape を一括で警告・解消するものではない境界を維持する。
- `unclear_dict_key` / `deep_type_ref` 等、V01-ADR-069 が将来候補として残すものを本 task の実装 scope に拡張しない。
- 後続 `V01-TASK-DATA-001-04` が共有 spec surface を更新する前提となる差分・commit boundary を evidence に記録する。

## Done condition

- V01-ADR-069 minimum scope が spec / implementation / tests に矛盾なく反映されている。
- `opaque_type_ref` warning と parser safety limit の expected behavior が local tests で確認されている。
- UC-002 の bare `any + note` debt 全体を本 task の救済範囲として誤記していない。
- V01-ADR-070 helper model または V01-ADR-073 tagged union を暗黙に引き込んでいない。

## Verification

- V01-ADR-069 §10 と spec 差分を照合する。
- TypeRef parser / diagnostic に関連する単体テストを local environment で実行する。
- 必要に応じて representative container type を用い、warning / non-warning boundary を確認する。
- 後続 enum spec update と衝突し得る更新 surface を evidence に明記する。

## Evidence

- Spec-only draft reflection performed for Codex review.
- Updated `docs/spec/type-ref.md` with V01-ADR-069 minimum boundary: nested `list<T>` / `dict<T>` remain valid, container nesting depth > 16 is `invalid_type_ref`, anonymous inline struct TypeRef is not introduced, `dict<T>` key semantics should be clarified by field name / model name / note, and `opaque_type_ref` is limited to container TypeRef containing `any`.
- Updated `docs/spec/diagnostics.md` with `opaque_type_ref` warning and clarified that bare `any + note` debt is outside this warning boundary.
- Codex review returned OK with minor fixes; no V01-ADR-069 blocking issue reported.
- Implemented parser safety limit as container nesting depth > 16 producing `invalid_type_ref`.
- Implemented `opaque_type_ref` warning for container TypeRef containing `any`; bare `any` remains outside this warning boundary.
- Added tests for parser safety limit, opaque container warning, and non-container `any` non-warning.
- Local verification found `TestForeachReturnsAnyCollectedSourceOK` still expected zero diagnostics while `list<any>` now intentionally emits `opaque_type_ref` warnings.
- Updated the foreach collected-source test to require no error diagnostics while accepting the expected `opaque_type_ref` warning.
- Re-run completed after `gofmt` on `internal/resolve/foreach_returns_test.go`.
- `go test ./internal/resolve`: passed.
- `go test ./...`: passed.
- Note: named list/dict models containing `any` also emit `opaque_type_ref` at their own TypeRef definition site; this matches the debt visibility baseline for container TypeRef containing `any`.
