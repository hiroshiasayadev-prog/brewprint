package designrecords

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestAuthoringProposalAcceptAndLifecycle(t *testing.T) {
	fx := newAuthoringFixture(t)
	targetPath := filepath.Join(fx.root, "docs", "tasks", "mcp", "TASK-MCP-001-01-first-task.md")
	before := readTestFile(t, targetPath)

	resp, err := ProposeRecordUpdate(context.Background(), fx.cfg, fx.idx, fx.store, ProposeRecordUpdateRequest{
		Kind: RecordKindTask,
		ID:   "TASK-MCP-001-01",
		Update: UpdateRequest{
			Type:            UpdateTypeNamedSectionReplace,
			SectionSelector: &SectionSelector{Heading: "Evidence"},
		},
		Body: stringPtr("accepted evidence\n"),
	})
	if err != nil {
		t.Fatalf("ProposeRecordUpdate: %v", err)
	}
	if !resp.ProposalCreated || resp.ProposalID == "" || resp.State != ProposalStateProposed {
		t.Fatalf("proposal response = %#v", resp)
	}
	if after := readTestFile(t, targetPath); after != before {
		t.Fatal("proposal creation wrote repository file")
	}

	accepted, err := AcceptProposedWrite(context.Background(), fx.cfg, mustBuildIndex(t, fx.cfg), fx.store, AcceptProposedWriteRequest{ProposalID: resp.ProposalID})
	if err != nil {
		t.Fatalf("AcceptProposedWrite: %v", err)
	}
	if !accepted.Written || accepted.State != ProposalStateAccepted || len(accepted.FilesWritten) != 1 {
		t.Fatalf("accept response = %#v", accepted)
	}
	if after := readTestFile(t, targetPath); !strings.Contains(after, "accepted evidence") {
		t.Fatalf("accepted write did not update file:\n%s", after)
	}

	second, err := AcceptProposedWrite(context.Background(), fx.cfg, mustBuildIndex(t, fx.cfg), fx.store, AcceptProposedWriteRequest{ProposalID: resp.ProposalID})
	if err != nil {
		t.Fatalf("second AcceptProposedWrite: %v", err)
	}
	if second.Written || !hasDiagnosticCategory(second.Diagnostics, DiagnosticCategory(ErrorCodeProposalAlreadyAccepted)) {
		t.Fatalf("second accept = %#v", second)
	}
}

func TestAuthoringDiscardExpiredAndUnknownProposal(t *testing.T) {
	fx := newAuthoringFixture(t)
	resp := proposeTaskSection(t, fx, "discarded\n")
	discarded, err := DiscardProposedWrite(context.Background(), fx.store, DiscardProposedWriteRequest{ProposalID: resp.ProposalID})
	if err != nil {
		t.Fatalf("DiscardProposedWrite: %v", err)
	}
	if !discarded.Discarded || discarded.Written {
		t.Fatalf("discard response = %#v", discarded)
	}
	accepted, err := AcceptProposedWrite(context.Background(), fx.cfg, fx.idx, fx.store, AcceptProposedWriteRequest{ProposalID: resp.ProposalID})
	if err != nil {
		t.Fatalf("AcceptProposedWrite discarded: %v", err)
	}
	if accepted.Written || !hasDiagnosticCategory(accepted.Diagnostics, DiagnosticCategory(ErrorCodeProposalDiscarded)) {
		t.Fatalf("discarded accept = %#v", accepted)
	}

	unknown, err := GetProposedWrite(context.Background(), fx.store, GetProposedWriteRequest{ProposalID: "pw_missing"})
	if err != nil {
		t.Fatalf("GetProposedWrite unknown: %v", err)
	}
	if !hasDiagnosticCategory(unknown.Diagnostics, DiagnosticCategory(ErrorCodeProposalNotFound)) {
		t.Fatalf("unknown proposal response = %#v", unknown)
	}

	now := time.Date(2026, 6, 2, 0, 0, 0, 0, time.UTC)
	fx.store.SetClockForTest(func() time.Time { return now })
	expiring := proposeTaskSection(t, fx, "expired\n")
	fx.store.SetClockForTest(func() time.Time { return now.Add(4 * 24 * time.Hour) })
	expired, err := GetProposedWrite(context.Background(), fx.store, GetProposedWriteRequest{ProposalID: expiring.ProposalID})
	if err != nil {
		t.Fatalf("GetProposedWrite expired: %v", err)
	}
	if !hasDiagnosticCategory(expired.Diagnostics, DiagnosticCategory(ErrorCodeProposalExpired)) {
		t.Fatalf("expired proposal response = %#v", expired)
	}
}

func TestAuthoringAcceptGuardsStaleTargetAndCreateCollision(t *testing.T) {
	fx := newAuthoringFixture(t)
	update := proposeTaskSection(t, fx, "stale\n")
	taskPath := filepath.Join(fx.root, "docs", "tasks", "mcp", "TASK-MCP-001-01-first-task.md")
	if err := os.WriteFile(taskPath, []byte(readTestFile(t, taskPath)+"\nexternal change\n"), 0o644); err != nil {
		t.Fatalf("external write: %v", err)
	}
	stale, err := AcceptProposedWrite(context.Background(), fx.cfg, mustBuildIndex(t, fx.cfg), fx.store, AcceptProposedWriteRequest{ProposalID: update.ProposalID})
	if err != nil {
		t.Fatalf("Accept stale: %v", err)
	}
	if stale.Written || !hasDiagnosticCategory(stale.Diagnostics, DiagnosticCategory(ErrorCodeProposalStale)) {
		t.Fatalf("stale accept = %#v", stale)
	}

	fx = newAuthoringFixture(t)
	create, err := ProposeRecordCreate(context.Background(), fx.cfg, fx.idx, fx.store, ProposeRecordCreateRequest{
		Kind:   RecordKindRequirement,
		ID:     "REQ-MCP-new",
		Domain: "MCP",
		Title:  "Collision",
		Fields: map[string]any{"status": "captured", "date": "2026-06-02", "source_refs": []any{}, "work_items": []any{}},
	})
	if err != nil {
		t.Fatalf("Propose create: %v", err)
	}
	writeTestFile(t, fx.root, "docs/requirements/mcp/"+create.Target.ResolvedID+"-other.md", "# "+create.Target.ResolvedID+": Other\n\n- **id**: "+create.Target.ResolvedID+"\n- **status**: captured\n- **date**: 2026-06-02\n- **source_refs**:\n- **work_items**:\n")
	collision, err := AcceptProposedWrite(context.Background(), fx.cfg, mustBuildIndex(t, fx.cfg), fx.store, AcceptProposedWriteRequest{ProposalID: create.ProposalID})
	if err != nil {
		t.Fatalf("Accept collision: %v", err)
	}
	if collision.Written || !hasDiagnosticCategory(collision.Diagnostics, DiagnosticCategory(ErrorCodeIDCollision)) {
		t.Fatalf("collision accept = %#v", collision)
	}
}

func TestAuthoringAcceptPartialMultiFileWriteReportsWrittenFiles(t *testing.T) {
	fx := newAuthoringFixture(t)
	resp, err := ProposeRecordCreate(context.Background(), fx.cfg, fx.idx, fx.store, ProposeRecordCreateRequest{
		Kind:     RecordKindTask,
		ID:       "TASK-MCP-001-new",
		Domain:   "MCP",
		ParentID: "WORK-MCP-001",
		Title:    "Partial write task",
		Fields:   map[string]any{"status": "todo", "date": "2026-06-02", "work_item": "WORK-MCP-001", "source_requirement": "REQ-MCP-001", "estimate": "0.5d", "depends_on": []any{}, "outputs": []any{"task"}},
	})
	if err != nil {
		t.Fatalf("Propose create: %v", err)
	}
	if !resp.ProposalCreated || resp.Target == nil {
		t.Fatalf("proposal response = %#v", resp)
	}

	parentPath := filepath.Join(fx.root, "docs", "work-items", "mcp", "WORK-MCP-001-first-work.md")
	if err := os.Chmod(parentPath, 0o444); err != nil {
		t.Fatalf("make parent read-only: %v", err)
	}
	defer func() {
		_ = os.Chmod(parentPath, 0o666)
	}()

	accepted, err := AcceptProposedWrite(context.Background(), fx.cfg, mustBuildIndex(t, fx.cfg), fx.store, AcceptProposedWriteRequest{ProposalID: resp.ProposalID})
	if err != nil {
		t.Fatalf("Accept partial write: %v", err)
	}
	if !accepted.Written || accepted.State != ProposalStateFailedFinal || len(accepted.FilesWritten) != 1 {
		t.Fatalf("partial write accept response = %#v", accepted)
	}
	if accepted.FilesWritten[0].Path != resp.Target.Path || accepted.FilesWritten[0].RecordID != resp.Target.ResolvedID {
		t.Fatalf("files_written = %#v, want first task file %s", accepted.FilesWritten, resp.Target.Path)
	}
	if len(accepted.Diagnostics) == 0 || len(accepted.RepairGuidance) == 0 {
		t.Fatalf("partial write response missing diagnostics/repair guidance: %#v", accepted)
	}
	if _, err := os.Stat(filepath.Join(fx.root, filepath.FromSlash(resp.Target.Path))); err != nil {
		t.Fatalf("first file was not written: %v", err)
	}

	retry, err := AcceptProposedWrite(context.Background(), fx.cfg, mustBuildIndex(t, fx.cfg), fx.store, AcceptProposedWriteRequest{ProposalID: resp.ProposalID})
	if err != nil {
		t.Fatalf("retry partial write proposal: %v", err)
	}
	if retry.Written || retry.State != ProposalStateFailedFinal {
		t.Fatalf("partial write proposal remained retryable: %#v", retry)
	}
}

func TestAuthoringCreateIDResolutionAndRejectedNewForms(t *testing.T) {
	fx := newAuthoringFixture(t)
	tests := []struct {
		name string
		req  ProposeRecordCreateRequest
		want string
	}{
		{
			name: "ADR-new",
			req:  ProposeRecordCreateRequest{Kind: RecordKindDecision, ID: "ADR-new", Title: "Next decision", Fields: map[string]any{"status": "proposed", "date": "2026-06-02", "depends_on": []any{}, "supersedes": []any{}, "migrated_to_spec": ""}},
			want: "ADR-002",
		},
		{
			name: "REQ-new",
			req:  ProposeRecordCreateRequest{Kind: RecordKindRequirement, ID: "REQ-MCP-new", Domain: "MCP", Title: "Next req", Fields: map[string]any{"status": "captured", "date": "2026-06-02", "source_refs": []any{}, "work_items": []any{}}},
			want: "REQ-MCP-002",
		},
		{
			name: "WORK-new",
			req:  ProposeRecordCreateRequest{Kind: RecordKindWorkItem, ID: "WORK-MCP-new", Domain: "MCP", Title: "Next work", Fields: map[string]any{"status": "implementation_pending", "date": "2026-06-02", "source_requirement": "REQ-MCP-001", "impact_refs": []any{}, "tasks": []any{}}},
			want: "WORK-MCP-002",
		},
		{
			name: "TASK-new",
			req:  ProposeRecordCreateRequest{Kind: RecordKindTask, ID: "TASK-MCP-001-new", Domain: "MCP", ParentID: "WORK-MCP-001", Title: "Next task", Fields: map[string]any{"status": "todo", "date": "2026-06-02", "work_item": "WORK-MCP-001", "source_requirement": "REQ-MCP-001", "estimate": "0.5d", "depends_on": []any{}, "outputs": []any{"task"}}},
			want: "TASK-MCP-001-02",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := ProposeRecordCreate(context.Background(), fx.cfg, fx.idx, fx.store, tt.req)
			if err != nil {
				t.Fatalf("ProposeRecordCreate: %v", err)
			}
			if !resp.ProposalCreated || resp.Target.ResolvedID != tt.want {
				t.Fatalf("resolved ID = %#v, want %s", resp, tt.want)
			}
		})
	}

	rejects := []ProposeRecordCreateRequest{
		{Kind: RecordKindSpec, ID: "SPEC-new", Title: "Spec", Fields: map[string]any{}},
		{Kind: RecordKindDecision, ID: "ADR-newish", Title: "Bad", Fields: map[string]any{}},
		{Kind: RecordKindRequirement, ID: "REQ-MCP-newer", Title: "Bad", Fields: map[string]any{}},
	}
	for _, req := range rejects {
		resp, err := ProposeRecordCreate(context.Background(), fx.cfg, fx.idx, fx.store, req)
		if err != nil {
			t.Fatalf("reject ProposeRecordCreate: %v", err)
		}
		if resp.ProposalCreated {
			t.Fatalf("malformed create produced proposal: %#v", resp)
		}
	}

	update, err := ProposeRecordUpdate(context.Background(), fx.cfg, fx.idx, fx.store, ProposeRecordUpdateRequest{Kind: RecordKindTask, ID: "TASK-MCP-001-new", Update: UpdateRequest{Type: UpdateTypeMetadataBlockReplace, Metadata: map[string]any{}}})
	if err != nil {
		t.Fatalf("ProposeRecordUpdate new: %v", err)
	}
	if update.ProposalCreated || !hasDiagnosticCategory(update.Diagnostics, DiagnosticCategory(ErrorCodeInvalidRequest)) {
		t.Fatalf("update new response = %#v", update)
	}
}

func TestAuthoringCreateInputContractNormalization(t *testing.T) {
	fx := newAuthoringFixture(t)
	body := "# REQ-MCP-050: Body\n\n- **id**: REQ-MCP-050\n- **status**: captured\n- **date**: 2026-06-02\n- **source_refs**:\n- **work_items**:\n"
	fields := map[string]any{"status": "captured", "date": "2026-06-02", "source_refs": []any{}, "work_items": []any{}}

	fieldsAndBody, err := ProposeRecordCreate(context.Background(), fx.cfg, fx.idx, fx.store, ProposeRecordCreateRequest{
		Kind:   RecordKindRequirement,
		ID:     "REQ-MCP-050",
		Domain: "mcp",
		Title:  "Fields and body",
		Fields: fields,
		Body:   &body,
	})
	if err != nil {
		t.Fatalf("fields plus body: %v", err)
	}
	if fieldsAndBody.ProposalCreated || fieldsAndBody.BodyCache != nil || !hasDiagnosticCategory(fieldsAndBody.Diagnostics, DiagnosticCategory(ErrorCodeInvalidRequest)) {
		t.Fatalf("fields plus body response = %#v", fieldsAndBody)
	}

	bodyAndCache, err := ProposeRecordCreate(context.Background(), fx.cfg, fx.idx, fx.store, ProposeRecordCreateRequest{
		Kind:        RecordKindRequirement,
		ID:          "REQ-MCP-051",
		Domain:      "mcp",
		Title:       "Body conflict",
		Body:        &body,
		BodyCacheID: "bc_other",
	})
	if err != nil {
		t.Fatalf("body plus cache: %v", err)
	}
	if bodyAndCache.ProposalCreated || bodyAndCache.BodyCache != nil || !hasDiagnosticCategory(bodyAndCache.Diagnostics, DiagnosticCategory(ErrorCodeInvalidBodySource)) {
		t.Fatalf("body plus cache response = %#v", bodyAndCache)
	}

	withoutFieldsID, err := ProposeRecordCreate(context.Background(), fx.cfg, fx.idx, fx.store, ProposeRecordCreateRequest{
		Kind:   RecordKindRequirement,
		ID:     "REQ-MCP-052",
		Domain: "mcp",
		Title:  "No fields id",
		Fields: fields,
	})
	if err != nil {
		t.Fatalf("without fields.id: %v", err)
	}
	if !withoutFieldsID.ProposalCreated || withoutFieldsID.Target.Domain != "MCP" || !strings.Contains(withoutFieldsID.Target.Path, "docs/requirements/mcp/REQ-MCP-052") || !strings.Contains(withoutFieldsID.Diff.Text, "- **id**: REQ-MCP-052") {
		t.Fatalf("without fields.id response = %#v\n%s", withoutFieldsID, withoutFieldsID.Diff.Text)
	}

	matchingFieldsID := map[string]any{"id": "req-mcp-053", "status": "captured", "date": "2026-06-02", "source_refs": []any{}, "work_items": []any{}}
	matching, err := ProposeRecordCreate(context.Background(), fx.cfg, fx.idx, fx.store, ProposeRecordCreateRequest{
		Kind:   RecordKindRequirement,
		ID:     "REQ-MCP-053",
		Domain: "MCP",
		Title:  "Matching fields id",
		Fields: matchingFieldsID,
	})
	if err != nil {
		t.Fatalf("matching fields.id: %v", err)
	}
	if !matching.ProposalCreated {
		t.Fatalf("matching fields.id response = %#v", matching)
	}

	mismatchedFieldsID := map[string]any{"id": "REQ-MCP-999", "status": "captured", "date": "2026-06-02", "source_refs": []any{}, "work_items": []any{}}
	mismatch, err := ProposeRecordCreate(context.Background(), fx.cfg, fx.idx, fx.store, ProposeRecordCreateRequest{
		Kind:   RecordKindRequirement,
		ID:     "REQ-MCP-054",
		Domain: "MCP",
		Title:  "Mismatched fields id",
		Fields: mismatchedFieldsID,
	})
	if err != nil {
		t.Fatalf("mismatched fields.id: %v", err)
	}
	if mismatch.ProposalCreated || !hasDiagnosticCategory(mismatch.Diagnostics, DiagnosticCategory(ErrorCodeInvalidRequest)) {
		t.Fatalf("mismatched fields.id response = %#v", mismatch)
	}

	placeholderFieldsID := map[string]any{"id": "REQ-MCP-new", "status": "captured", "date": "2026-06-02", "source_refs": []any{}, "work_items": []any{}}
	placeholder, err := ProposeRecordCreate(context.Background(), fx.cfg, fx.idx, fx.store, ProposeRecordCreateRequest{
		Kind:   RecordKindRequirement,
		ID:     "REQ-MCP-new",
		Domain: "MCP",
		Title:  "Placeholder fields id",
		Fields: placeholderFieldsID,
	})
	if err != nil {
		t.Fatalf("placeholder fields.id: %v", err)
	}
	if placeholder.ProposalCreated || !hasDiagnosticCategory(placeholder.Diagnostics, DiagnosticCategory(ErrorCodeInvalidRequest)) {
		t.Fatalf("placeholder fields.id response = %#v", placeholder)
	}

	domainMismatch, err := ProposeRecordCreate(context.Background(), fx.cfg, fx.idx, fx.store, ProposeRecordCreateRequest{
		Kind:   RecordKindRequirement,
		ID:     "REQ-MCP-055",
		Domain: "data",
		Title:  "Domain mismatch",
		Fields: fields,
	})
	if err != nil {
		t.Fatalf("domain mismatch: %v", err)
	}
	if domainMismatch.ProposalCreated || !hasDiagnosticCategory(domainMismatch.Diagnostics, DiagnosticCategory(ErrorCodeInvalidRequest)) {
		t.Fatalf("domain mismatch response = %#v", domainMismatch)
	}
}

func TestAuthoringTaskCreateRequiresExplicitParentMetadataAndReciprocalUpdate(t *testing.T) {
	fx := newAuthoringFixture(t)
	withoutParent, err := ProposeRecordCreate(context.Background(), fx.cfg, fx.idx, fx.store, ProposeRecordCreateRequest{
		Kind:   RecordKindTask,
		ID:     "TASK-MCP-001-new",
		Domain: "MCP",
		Title:  "No parent",
		Fields: map[string]any{"status": "todo", "date": "2026-06-02", "source_requirement": "REQ-MCP-001", "estimate": "0.5d", "depends_on": []any{}, "outputs": []any{}},
	})
	if err != nil {
		t.Fatalf("task without parent: %v", err)
	}
	if withoutParent.ProposalCreated {
		t.Fatalf("task without parent produced proposal: %#v", withoutParent)
	}

	withoutExplicitWorkItem, err := ProposeRecordCreate(context.Background(), fx.cfg, fx.idx, fx.store, ProposeRecordCreateRequest{
		Kind:     RecordKindTask,
		ID:       "TASK-MCP-001-new",
		Domain:   "MCP",
		ParentID: "WORK-MCP-001",
		Title:    "No explicit relation",
		Fields:   map[string]any{"status": "todo", "date": "2026-06-02", "source_requirement": "REQ-MCP-001", "estimate": "0.5d", "depends_on": []any{}, "outputs": []any{}},
	})
	if err != nil {
		t.Fatalf("task without explicit work_item: %v", err)
	}
	if withoutExplicitWorkItem.ProposalCreated {
		t.Fatalf("task without explicit work_item produced proposal: %#v", withoutExplicitWorkItem)
	}

	resp, err := ProposeRecordCreate(context.Background(), fx.cfg, fx.idx, fx.store, ProposeRecordCreateRequest{
		Kind:     RecordKindTask,
		ID:       "TASK-MCP-001-new",
		Domain:   "MCP",
		ParentID: "WORK-MCP-001",
		Title:    "Second task",
		Fields:   map[string]any{"status": "todo", "date": "2026-06-02", "work_item": "WORK-MCP-001", "source_requirement": "REQ-MCP-001", "estimate": "0.5d", "depends_on": []any{}, "outputs": []any{"task"}},
	})
	if err != nil {
		t.Fatalf("task reciprocal: %v", err)
	}
	if !resp.ProposalCreated || len(resp.Diff.Files) != 2 {
		t.Fatalf("reciprocal task proposal = %#v", resp)
	}
	if !strings.Contains(resp.Diff.Text, "TASK-MCP-001-02") || !strings.Contains(resp.Diff.Text, "WORK-MCP-001") {
		t.Fatalf("reciprocal diff missing expected IDs:\n%s", resp.Diff.Text)
	}

	reportOnly, err := ProposeRecordCreate(context.Background(), fx.cfg, fx.idx, fx.store, ProposeRecordCreateRequest{
		Kind:                 RecordKindTask,
		ID:                   "TASK-MCP-001-new",
		Domain:               "MCP",
		ParentID:             "WORK-MCP-001",
		Title:                "Follow-up task",
		Fields:               map[string]any{"status": "todo", "date": "2026-06-02", "work_item": "WORK-MCP-001", "source_requirement": "REQ-MCP-001", "estimate": "0.5d", "depends_on": []any{}, "outputs": []any{"task"}},
		ReciprocalUpdateMode: "report_required_follow_up",
	})
	if err != nil {
		t.Fatalf("task report follow-up: %v", err)
	}
	if !reportOnly.ProposalCreated || len(reportOnly.RequiredFollowUpUpdates) != 1 {
		t.Fatalf("report follow-up proposal = %#v", reportOnly)
	}
	rejected, err := AcceptProposedWrite(context.Background(), fx.cfg, fx.idx, fx.store, AcceptProposedWriteRequest{ProposalID: reportOnly.ProposalID})
	if err != nil {
		t.Fatalf("accept report follow-up: %v", err)
	}
	if rejected.Written || !hasDiagnosticCategory(rejected.Diagnostics, DiagnosticCategory(ErrorCodeRequiredFollowUpNotSatisfied)) {
		t.Fatalf("required follow-up accept = %#v", rejected)
	}
}

func TestAuthoringBodySourceAndCache(t *testing.T) {
	fx := newAuthoringFixture(t)
	body := "cached body\n"
	both, err := ProposeRecordUpdate(context.Background(), fx.cfg, fx.idx, fx.store, ProposeRecordUpdateRequest{
		Kind:        RecordKindTask,
		ID:          "TASK-MCP-001-01",
		Update:      UpdateRequest{Type: UpdateTypeNamedSectionReplace, SectionSelector: &SectionSelector{Heading: "Evidence"}},
		Body:        &body,
		BodyCacheID: "bc_other",
	})
	if err != nil {
		t.Fatalf("both body sources: %v", err)
	}
	if both.ProposalCreated || both.BodyCache != nil || !hasDiagnosticCategory(both.Diagnostics, DiagnosticCategory(ErrorCodeInvalidBodySource)) {
		t.Fatalf("both body sources response = %#v", both)
	}

	missing, err := ProposeRecordUpdate(context.Background(), fx.cfg, fx.idx, fx.store, ProposeRecordUpdateRequest{
		Kind:   RecordKindTask,
		ID:     "TASK-MCP-001-01",
		Update: UpdateRequest{Type: UpdateTypeNamedSectionReplace, SectionSelector: &SectionSelector{Heading: "Evidence"}},
	})
	if err != nil {
		t.Fatalf("missing body: %v", err)
	}
	if missing.ProposalCreated || !hasDiagnosticCategory(missing.Diagnostics, DiagnosticCategory(ErrorCodeInvalidBodySource)) {
		t.Fatalf("missing body response = %#v", missing)
	}

	proposed := proposeTaskSection(t, fx, body)
	if proposed.BodyCache == nil {
		t.Fatalf("successful body proposal did not return reusable body cache: %#v", proposed)
	}
	reuse, err := ProposeRecordUpdate(context.Background(), fx.cfg, fx.idx, fx.store, ProposeRecordUpdateRequest{
		Kind:        RecordKindTask,
		ID:          "TASK-MCP-001-01",
		Update:      UpdateRequest{Type: UpdateTypeNamedSectionReplace, SectionSelector: &SectionSelector{Heading: "Verification"}},
		BodyCacheID: proposed.BodyCache.BodyCacheID,
	})
	if err != nil {
		t.Fatalf("reuse body cache: %v", err)
	}
	if !reuse.ProposalCreated {
		t.Fatalf("body cache was not reusable: %#v", reuse)
	}

	unknown, err := ProposeRecordUpdate(context.Background(), fx.cfg, fx.idx, fx.store, ProposeRecordUpdateRequest{
		Kind:        RecordKindTask,
		ID:          "TASK-MCP-001-01",
		Update:      UpdateRequest{Type: UpdateTypeNamedSectionReplace, SectionSelector: &SectionSelector{Heading: "Evidence"}},
		BodyCacheID: "bc_unknown",
	})
	if err != nil {
		t.Fatalf("unknown body cache: %v", err)
	}
	if unknown.ProposalCreated || !hasDiagnosticCategory(unknown.Diagnostics, DiagnosticCategory(ErrorCodeBodyCacheNotFound)) {
		t.Fatalf("unknown body cache response = %#v", unknown)
	}

	now := time.Date(2026, 6, 2, 0, 0, 0, 0, time.UTC)
	fx.store.SetClockForTest(func() time.Time { return now })
	expiring := proposeTaskSection(t, fx, "expires\n")
	fx.store.SetClockForTest(func() time.Time { return now.Add(4 * 24 * time.Hour) })
	expired, err := ProposeRecordUpdate(context.Background(), fx.cfg, fx.idx, fx.store, ProposeRecordUpdateRequest{
		Kind:        RecordKindTask,
		ID:          "TASK-MCP-001-01",
		Update:      UpdateRequest{Type: UpdateTypeNamedSectionReplace, SectionSelector: &SectionSelector{Heading: "Verification"}},
		BodyCacheID: expiring.BodyCache.BodyCacheID,
	})
	if err != nil {
		t.Fatalf("expired body cache: %v", err)
	}
	if expired.ProposalCreated || !hasDiagnosticCategory(expired.Diagnostics, DiagnosticCategory(ErrorCodeBodyCacheExpired)) {
		t.Fatalf("expired body cache response = %#v", expired)
	}
}

func TestAuthoringMetadataReplacement(t *testing.T) {
	fx := newAuthoringFixture(t)
	task, err := ProposeRecordUpdate(context.Background(), fx.cfg, fx.idx, fx.store, ProposeRecordUpdateRequest{
		Kind: RecordKindTask,
		ID:   "TASK-MCP-001-01",
		Update: UpdateRequest{Type: UpdateTypeMetadataBlockReplace, Metadata: map[string]any{
			"id": "TASK-MCP-001-01", "status": "doing", "date": "2026-06-02", "work_item": "WORK-MCP-001", "source_requirement": "REQ-MCP-001", "estimate": "1d", "depends_on": []any{}, "outputs": []any{"updated"},
		}},
	})
	if err != nil {
		t.Fatalf("task metadata replace: %v", err)
	}
	if !task.ProposalCreated || !strings.Contains(task.Diff.Text, "- **status**: doing") {
		t.Fatalf("task metadata response = %#v", task)
	}

	spec, err := ProposeRecordUpdate(context.Background(), fx.cfg, fx.idx, fx.store, ProposeRecordUpdateRequest{
		Kind: RecordKindSpec,
		ID:   "SPEC-test",
		Update: UpdateRequest{Type: UpdateTypeMetadataBlockReplace, Metadata: map[string]any{
			"scope":  "docs/spec/test.md",
			"status": "draft",
			"design_record": map[string]any{
				"id": "SPEC-test", "kind": "spec", "status": "draft", "depends_on": []any{"ADR-001"},
			},
		}},
	})
	if err != nil {
		t.Fatalf("spec metadata replace: %v", err)
	}
	if !spec.ProposalCreated || !strings.Contains(spec.Diff.Text, "auxiliary: keep-me") {
		t.Fatalf("spec metadata did not preserve unknown field: %#v\n%s", spec, spec.Diff.Text)
	}
}

func TestAuthoringMetadataReplacementMissingRequiredFields(t *testing.T) {
	fx := newAuthoringFixture(t)
	adr, err := ProposeRecordUpdate(context.Background(), fx.cfg, fx.idx, fx.store, ProposeRecordUpdateRequest{
		Kind: RecordKindDecision,
		ID:   "ADR-001",
		Update: UpdateRequest{Type: UpdateTypeMetadataBlockReplace, Metadata: map[string]any{
			"status": "accepted", "depends_on": []any{}, "supersedes": []any{}, "migrated_to_spec": "",
		}},
	})
	if err != nil {
		t.Fatalf("adr missing metadata: %v", err)
	}
	if adr.ProposalCreated || !hasDiagnosticCategory(adr.Diagnostics, DiagnosticMissingRequiredMetadata) {
		t.Fatalf("adr missing metadata response = %#v", adr)
	}

	task, err := ProposeRecordUpdate(context.Background(), fx.cfg, fx.idx, fx.store, ProposeRecordUpdateRequest{
		Kind: RecordKindTask,
		ID:   "TASK-MCP-001-01",
		Update: UpdateRequest{Type: UpdateTypeMetadataBlockReplace, Metadata: map[string]any{
			"id": "TASK-MCP-001-01", "status": "doing", "date": "2026-06-02", "work_item": "WORK-MCP-001", "estimate": "1d", "depends_on": []any{}, "outputs": []any{"updated"},
		}},
	})
	if err != nil {
		t.Fatalf("task missing metadata: %v", err)
	}
	if task.ProposalCreated || !hasDiagnosticCategory(task.Diagnostics, DiagnosticMissingRequiredMetadata) {
		t.Fatalf("task missing metadata response = %#v", task)
	}
}

func TestAuthoringNamedSectionSelectors(t *testing.T) {
	fx := newAuthoringFixture(t)
	ok := proposeTaskSection(t, fx, "single match\n")
	if !ok.ProposalCreated || !strings.Contains(ok.Diff.Text, "single match") {
		t.Fatalf("section proposal = %#v", ok)
	}

	noMatch, err := ProposeRecordUpdate(context.Background(), fx.cfg, fx.idx, fx.store, ProposeRecordUpdateRequest{
		Kind:   RecordKindTask,
		ID:     "TASK-MCP-001-01",
		Update: UpdateRequest{Type: UpdateTypeNamedSectionReplace, SectionSelector: &SectionSelector{Heading: "Missing"}},
		Body:   stringPtr("body\n"),
	})
	if err != nil {
		t.Fatalf("no match: %v", err)
	}
	if noMatch.ProposalCreated || !hasDiagnosticCategory(noMatch.Diagnostics, DiagnosticCategory(ErrorCodeSectionSelectorNoMatch)) || noMatch.BodyCache == nil {
		t.Fatalf("no match response = %#v", noMatch)
	}
	noMatchDiagnostic := diagnosticByCategory(noMatch.Diagnostics, DiagnosticCategory(ErrorCodeSectionSelectorNoMatch))
	if noMatchDiagnostic == nil || len(noMatchDiagnostic.CandidateHeadings) == 0 {
		t.Fatalf("no match diagnostic missing candidate headings: %#v", noMatch.Diagnostics)
	}
	if noMatchDiagnostic.CandidateHeadings[0].Heading != "TASK-MCP-001-01: First task" || noMatchDiagnostic.CandidateHeadings[0].Level != 1 || noMatchDiagnostic.CandidateHeadings[0].Ordinal != 1 {
		t.Fatalf("no match candidate headings = %#v", noMatchDiagnostic.CandidateHeadings)
	}

	multiPath := filepath.Join(fx.root, "docs", "tasks", "mcp", "TASK-MCP-001-01-first-task.md")
	if err := os.WriteFile(multiPath, []byte(strings.Replace(readTestFile(t, multiPath), "## Evidence\n", "## Evidence\n\n## Evidence\n", 1)), 0o644); err != nil {
		t.Fatalf("make multi section: %v", err)
	}
	fx.idx = mustBuildIndex(t, fx.cfg)
	multi, err := ProposeRecordUpdate(context.Background(), fx.cfg, fx.idx, fx.store, ProposeRecordUpdateRequest{
		Kind:   RecordKindTask,
		ID:     "TASK-MCP-001-01",
		Update: UpdateRequest{Type: UpdateTypeNamedSectionReplace, SectionSelector: &SectionSelector{Heading: "Evidence"}},
		Body:   stringPtr("body\n"),
	})
	if err != nil {
		t.Fatalf("multi match: %v", err)
	}
	if multi.ProposalCreated || !hasDiagnosticCategory(multi.Diagnostics, DiagnosticCategory(ErrorCodeSectionSelectorAmbiguous)) {
		t.Fatalf("multi match response = %#v", multi)
	}
	multiDiagnostic := diagnosticByCategory(multi.Diagnostics, DiagnosticCategory(ErrorCodeSectionSelectorAmbiguous))
	if multiDiagnostic == nil || len(multiDiagnostic.CandidateHeadings) != 2 {
		t.Fatalf("multi match diagnostic missing candidate headings: %#v", multi.Diagnostics)
	}
	for _, candidate := range multiDiagnostic.CandidateHeadings {
		if candidate.Heading != "Evidence" || candidate.Level != 2 || candidate.Ordinal == 0 {
			t.Fatalf("ambiguous candidate heading = %#v", candidate)
		}
	}
}

func TestAuthoringSectionSelectorIgnoresFrontMatterAndFenceHeadings(t *testing.T) {
	fx := newAuthoringFixture(t)
	spec, err := ProposeRecordUpdate(context.Background(), fx.cfg, fx.idx, fx.store, ProposeRecordUpdateRequest{
		Kind:   RecordKindSpec,
		ID:     "SPEC-test",
		Update: UpdateRequest{Type: UpdateTypeNamedSectionReplace, SectionSelector: &SectionSelector{Heading: "Ignored"}},
		Body:   stringPtr("body\n"),
	})
	if err != nil {
		t.Fatalf("front matter/fence ignored: %v", err)
	}
	if spec.ProposalCreated || !hasDiagnosticCategory(spec.Diagnostics, DiagnosticCategory(ErrorCodeSectionSelectorNoMatch)) {
		t.Fatalf("ignored heading response = %#v", spec)
	}
}

func TestAuthoringProposalValidationIsAffectedRecordSetOnly(t *testing.T) {
	fx := newAuthoringFixture(t)
	writeTestFile(t, fx.root, "docs/requirements/mcp/REQ-MCP-999-bad.md", "# REQ-MCP-999: Bad\n\n- **id**: REQ-MCP-999\n- **status**: captured\n- **date**:\n- **source_refs**:\n- **work_items**:\n")
	fx.idx = mustBuildIndex(t, fx.cfg)
	resp := proposeTaskSection(t, fx, "proposal-local validation\n")
	if !resp.Validation.OK || len(resp.Validation.Diagnostics) != 0 {
		t.Fatalf("proposal validation included unrelated diagnostics: %#v", resp.Validation)
	}
	accepted, err := AcceptProposedWrite(context.Background(), fx.cfg, mustBuildIndex(t, fx.cfg), fx.store, AcceptProposedWriteRequest{ProposalID: resp.ProposalID})
	if err != nil {
		t.Fatalf("Accept with unrelated validation failure: %v", err)
	}
	if !accepted.Written || !accepted.Validation.OK || len(accepted.Validation.Diagnostics) != 0 || len(accepted.RepairGuidance) != 0 {
		t.Fatalf("accept validation included unrelated diagnostics: %#v", accepted)
	}
	taskPath := filepath.Join(fx.root, "docs", "tasks", "mcp", "TASK-MCP-001-01-first-task.md")
	if !strings.Contains(readTestFile(t, taskPath), "proposal-local validation") {
		t.Fatal("accepted write was not applied")
	}
}

func TestAuthoringHypotheticalIndexPreservesUnchangedSemanticRefSources(t *testing.T) {
	fx := newAuthoringFixture(t)
	fx.idx.SemanticRefSources = []SemanticRefSource{{
		Path:     "docs/spec/test.md",
		RecordID: "SPEC-test",
		Decls: []SemanticRefDecl{{
			Ref:        "spec:test",
			Path:       "docs/spec/test.md",
			TargetType: SemanticTargetDocument,
		}},
		Headings: []Heading{{Level: 1, Text: "Test spec"}},
	}}

	hyp := buildHypotheticalIndex(fx.idx, []ProposedFile{{
		Path:       "docs/tasks/mcp/TASK-MCP-001-01-first-task.md",
		Change:     "modify",
		RecordID:   "TASK-MCP-001-01",
		RecordKind: RecordKindTask,
		Content:    readTestFile(t, filepath.Join(fx.root, "docs", "tasks", "mcp", "TASK-MCP-001-01-first-task.md")),
	}})
	if len(hyp.SemanticRefSources) != 1 || hyp.SemanticRefSources[0].RecordID != "SPEC-test" {
		t.Fatalf("semantic ref sources were not preserved: %#v", hyp.SemanticRefSources)
	}
}

type authoringFixture struct {
	root  string
	cfg   Config
	idx   *Index
	store *AuthoringStore
}

func newAuthoringFixture(t *testing.T) authoringFixture {
	t.Helper()
	root := t.TempDir()
	writeTestFile(t, root, "docs/adr/001-one.md", "# 001: One\n\n- **status**: accepted\n- **date**: 2026-06-01\n- **depends_on**: \n- **supersedes**: \n- **migrated_to_spec**: \n")
	writeTestFile(t, root, "docs/requirements/mcp/REQ-MCP-001-first-req.md", "# REQ-MCP-001: First req\n\n- **id**: REQ-MCP-001\n- **status**: captured\n- **date**: 2026-06-01\n- **source_refs**:\n- **work_items**:\n  - WORK-MCP-001\n\n## Requirement\n")
	writeTestFile(t, root, "docs/work-items/mcp/WORK-MCP-001-first-work.md", "# WORK-MCP-001: First work\n\n- **id**: WORK-MCP-001\n- **status**: implementation_pending\n- **date**: 2026-06-01\n- **source_requirement**: REQ-MCP-001\n- **impact_refs**:\n- **tasks**:\n  - TASK-MCP-001-01\n\n## Goal\n")
	writeTestFile(t, root, "docs/tasks/mcp/TASK-MCP-001-01-first-task.md", "# TASK-MCP-001-01: First task\n\n- **id**: TASK-MCP-001-01\n- **status**: todo\n- **date**: 2026-06-01\n- **work_item**: WORK-MCP-001\n- **source_requirement**: REQ-MCP-001\n- **estimate**: 0.5d\n- **depends_on**:\n- **outputs**:\n  - initial\n\n## Goal\n\n## Verification\n\n## Evidence\nold evidence\n")
	writeTestFile(t, root, "docs/spec/test.md", "---\nscope: docs/spec/test.md\nstatus: draft\nauxiliary: keep-me\nsummary: '# Ignored'\ndesign_record:\n  id: SPEC-test\n  kind: spec\n  status: draft\n  depends_on:\n    - ADR-001\n---\n\n# Test spec\n\n```md\n## Ignored\n```\n\n## Real\n")
	cfg, err := NewConfig(root)
	if err != nil {
		t.Fatalf("NewConfig: %v", err)
	}
	return authoringFixture{root: root, cfg: cfg, idx: mustBuildIndex(t, cfg), store: NewAuthoringStore()}
}

func proposeTaskSection(t *testing.T, fx authoringFixture, body string) ProposeRecordResponse {
	t.Helper()
	resp, err := ProposeRecordUpdate(context.Background(), fx.cfg, fx.idx, fx.store, ProposeRecordUpdateRequest{
		Kind: RecordKindTask,
		ID:   "TASK-MCP-001-01",
		Update: UpdateRequest{
			Type:            UpdateTypeNamedSectionReplace,
			SectionSelector: &SectionSelector{Heading: "Evidence"},
		},
		Body: &body,
	})
	if err != nil {
		t.Fatalf("ProposeRecordUpdate: %v", err)
	}
	if !resp.ProposalCreated {
		t.Fatalf("proposal was not created: %#v", resp)
	}
	return resp
}

func mustBuildIndex(t *testing.T, cfg Config) *Index {
	t.Helper()
	idx, err := BuildIndex(context.Background(), cfg)
	if err != nil {
		t.Fatalf("BuildIndex: %v", err)
	}
	return idx
}

func readTestFile(t *testing.T, path string) string {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile %s: %v", path, err)
	}
	return string(data)
}

func hasDiagnosticCategory(diagnostics []Diagnostic, category DiagnosticCategory) bool {
	return diagnosticByCategory(diagnostics, category) != nil
}

func diagnosticByCategory(diagnostics []Diagnostic, category DiagnosticCategory) *Diagnostic {
	for _, diagnostic := range diagnostics {
		if diagnostic.Category == category {
			return &diagnostic
		}
	}
	return nil
}
