package designrecords

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"gopkg.in/yaml.v3"
)

const authoringRetentionDays = 3

const (
	ProposalStateProposed    = "proposed"
	ProposalStateAccepted    = "accepted"
	ProposalStateDiscarded   = "discarded"
	ProposalStateFailedFinal = "failed_final"

	ProposalOperationCreate = "create"
	ProposalOperationUpdate = "update"

	UpdateTypeMetadataBlockReplace = "metadata_block_replace"
	UpdateTypeNamedSectionReplace  = "named_section_replace"
)

const noWriteProposalNote = "No repository files have been written. Call accept_proposed_write with this proposal_id to apply the diff."

var (
	createDecisionPlaceholderPattern    = regexp.MustCompile(`^ADR-new$`)
	createRequirementPlaceholderPattern = regexp.MustCompile(`^REQ-([A-Z0-9](?:[A-Z0-9-]*[A-Z0-9])?)-new$`)
	createWorkItemPlaceholderPattern    = regexp.MustCompile(`^WORK-([A-Z0-9](?:[A-Z0-9-]*[A-Z0-9])?)-new$`)
	createTaskPlaceholderPattern        = regexp.MustCompile(`^TASK-([A-Z0-9](?:[A-Z0-9-]*[A-Z0-9])?)-(\d{3})-new$`)
)

type AuthoringStore struct {
	mu        sync.Mutex
	now       func() time.Time
	nextID    int
	proposals map[string]*StoredProposal
	bodies    map[string]BodyCacheEntry
}

type StoredProposal struct {
	ProposalID              string
	State                   string
	Operation               string
	TargetKind              RecordKind
	Target                  AuthoringTarget `json:"target"`
	ExpiresAt               time.Time
	RetentionDays           int
	Diff                    Diff
	Validation              ValidateRecordsResponse
	Diagnostics             []Diagnostic
	Note                    string
	Files                   []ProposedFile
	RequiredFollowUpUpdates []RequiredFollowUpUpdate
	ProposalCreatedAt       time.Time
}

type ProposedFile struct {
	Path       string
	Change     string
	RecordID   string
	RecordKind RecordKind
	BaseHash   string
	BaseID     string
	BaseKind   RecordKind
	Content    string
}

type BodyCacheEntry struct {
	BodyCacheID   string    `json:"body_cache_id"`
	ExpiresAt     time.Time `json:"expires_at"`
	RetentionDays int       `json:"retention_days"`
	Body          string    `json:"-"`
}

type ProposeRecordCreateRequest struct {
	Kind                 RecordKind     `json:"kind"`
	ID                   string         `json:"id"`
	Domain               string         `json:"domain,omitempty"`
	ParentID             string         `json:"parent_id,omitempty"`
	Title                string         `json:"title"`
	Fields               map[string]any `json:"fields"`
	Body                 *string        `json:"body,omitempty"`
	BodyCacheID          string         `json:"body_cache_id,omitempty"`
	ReciprocalUpdateMode string         `json:"reciprocal_update_mode,omitempty"`
}

type ProposeRecordUpdateRequest struct {
	Kind        RecordKind    `json:"kind"`
	ID          string        `json:"id"`
	Update      UpdateRequest `json:"update"`
	Body        *string       `json:"body,omitempty"`
	BodyCacheID string        `json:"body_cache_id,omitempty"`
}

type UpdateRequest struct {
	Type            string           `json:"type"`
	Metadata        map[string]any   `json:"metadata,omitempty"`
	SectionSelector *SectionSelector `json:"section_selector,omitempty"`
}

type SectionSelector struct {
	Heading string `json:"heading"`
	Match   string `json:"match,omitempty"`
	Level   *int   `json:"level,omitempty"`
}

type AuthoringTarget struct {
	RequestedID string     `json:"requested_id"`
	ResolvedID  string     `json:"resolved_id"`
	Kind        RecordKind `json:"kind"`
	Domain      string     `json:"domain,omitempty"`
	ParentID    string     `json:"parent_id,omitempty"`
	Path        string     `json:"path"`
}

type Diff struct {
	Format string     `json:"format"`
	Files  []DiffFile `json:"files"`
	Text   string     `json:"text"`
}

type DiffFile struct {
	Path       string     `json:"path"`
	Change     string     `json:"change"`
	RecordID   string     `json:"record_id,omitempty"`
	RecordKind RecordKind `json:"record_kind,omitempty"`
}

type ProposeRecordResponse struct {
	ProposalCreated         bool                     `json:"proposal_created"`
	ProposalID              string                   `json:"proposal_id,omitempty"`
	State                   string                   `json:"state,omitempty"`
	Operation               string                   `json:"operation,omitempty"`
	TargetKind              RecordKind               `json:"target_kind,omitempty"`
	Target                  *AuthoringTarget         `json:"target,omitempty"`
	ExpiresAt               *time.Time               `json:"expires_at,omitempty"`
	RetentionDays           int                      `json:"retention_days,omitempty"`
	Diff                    *Diff                    `json:"diff,omitempty"`
	Validation              ValidateRecordsResponse  `json:"validation"`
	Diagnostics             []Diagnostic             `json:"diagnostics"`
	Note                    string                   `json:"note,omitempty"`
	BodyCache               *BodyCacheEntry          `json:"body_cache,omitempty"`
	RequiredFollowUpUpdates []RequiredFollowUpUpdate `json:"required_follow_up_updates,omitempty"`
}

type GetProposedWriteRequest struct {
	ProposalID string `json:"proposal_id"`
}

type GetProposedWriteResponse struct {
	ProposalID              string                   `json:"proposal_id,omitempty"`
	State                   string                   `json:"state,omitempty"`
	Operation               string                   `json:"operation,omitempty"`
	TargetKind              RecordKind               `json:"target_kind,omitempty"`
	Target                  *AuthoringTarget         `json:"target,omitempty"`
	ExpiresAt               *time.Time               `json:"expires_at,omitempty"`
	RetentionDays           int                      `json:"retention_days,omitempty"`
	Diff                    *Diff                    `json:"diff,omitempty"`
	Validation              ValidateRecordsResponse  `json:"validation"`
	Diagnostics             []Diagnostic             `json:"diagnostics"`
	Note                    string                   `json:"note,omitempty"`
	RequiredFollowUpUpdates []RequiredFollowUpUpdate `json:"required_follow_up_updates,omitempty"`
}

type AcceptProposedWriteRequest struct {
	ProposalID string `json:"proposal_id"`
}

type AcceptProposedWriteResponse struct {
	ProposalID     string                  `json:"proposal_id"`
	State          string                  `json:"state"`
	Written        bool                    `json:"written"`
	FilesWritten   []WrittenFile           `json:"files_written"`
	Validation     ValidateRecordsResponse `json:"validation"`
	RepairGuidance []string                `json:"repair_guidance"`
	Diagnostics    []Diagnostic            `json:"diagnostics"`
}

type WrittenFile struct {
	Path       string     `json:"path"`
	RecordID   string     `json:"record_id,omitempty"`
	RecordKind RecordKind `json:"record_kind,omitempty"`
}

type DiscardProposedWriteRequest struct {
	ProposalID string `json:"proposal_id"`
}

type DiscardProposedWriteResponse struct {
	ProposalID  string       `json:"proposal_id"`
	State       string       `json:"state"`
	Discarded   bool         `json:"discarded"`
	Written     bool         `json:"written"`
	Diagnostics []Diagnostic `json:"diagnostics"`
}

type RequiredFollowUpUpdate struct {
	RecordID string     `json:"record_id"`
	Kind     RecordKind `json:"kind"`
	Field    string     `json:"field"`
	Value    string     `json:"value"`
	Message  string     `json:"message"`
}

type authoringPreparation struct {
	target                  AuthoringTarget
	files                   []ProposedFile
	requiredFollowUpUpdates []RequiredFollowUpUpdate
	diagnostics             []Diagnostic
}

func NewAuthoringStore() *AuthoringStore {
	return &AuthoringStore{
		now:       time.Now,
		proposals: map[string]*StoredProposal{},
		bodies:    map[string]BodyCacheEntry{},
	}
}

func (s *AuthoringStore) SetClockForTest(now func() time.Time) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.now = now
}

func ProposeRecordCreate(ctx context.Context, cfg Config, idx *Index, store *AuthoringStore, req ProposeRecordCreateRequest) (ProposeRecordResponse, error) {
	if err := ctx.Err(); err != nil {
		return ProposeRecordResponse{}, err
	}
	if idx == nil {
		return ProposeRecordResponse{}, newToolError(ErrorCodeInvalidRequest, "index is nil")
	}
	if store == nil {
		return ProposeRecordResponse{}, newToolError(ErrorCodeInvalidRequest, "authoring store is nil")
	}
	if err := validateCreateKind(req.Kind); err != nil {
		return failedProposalResponse(nil, nil, authoringDiagnostic(ErrorCodeUnsupportedKind, err.Error())), nil
	}
	body, bodyCache, diagnostics, ok := resolveBodySource(store, req.Body, req.BodyCacheID, false)
	if !ok {
		return failedProposalResponse(nil, nil, diagnostics...), nil
	}
	prep, err := prepareCreate(ctx, cfg, idx, req, body)
	if err != nil {
		if body != nil {
			bodyCache = store.cacheBody(*body)
		}
		return failedProposalResponse(bodyCache, nil, authoringDiagnostic(ErrorCodeInvalidRequest, err.Error())), nil
	}
	if len(prep.diagnostics) > 0 {
		if body != nil {
			bodyCache = store.cacheBody(*body)
		}
		return failedProposalResponse(bodyCache, nil, prep.diagnostics...), nil
	}
	return persistProposal(ctx, cfg, idx, store, ProposalOperationCreate, prep, bodyCache)
}

func ProposeRecordUpdate(ctx context.Context, cfg Config, idx *Index, store *AuthoringStore, req ProposeRecordUpdateRequest) (ProposeRecordResponse, error) {
	if err := ctx.Err(); err != nil {
		return ProposeRecordResponse{}, err
	}
	if idx == nil {
		return ProposeRecordResponse{}, newToolError(ErrorCodeInvalidRequest, "index is nil")
	}
	if store == nil {
		return ProposeRecordResponse{}, newToolError(ErrorCodeInvalidRequest, "authoring store is nil")
	}
	if err := validateUpdateKind(req.Kind); err != nil {
		return failedProposalResponse(nil, nil, authoringDiagnostic(ErrorCodeUnsupportedKind, err.Error())), nil
	}
	if hasSequenceNewToken(req.ID, req.Kind) {
		return failedProposalResponse(nil, nil, authoringDiagnostic(ErrorCodeInvalidRequest, "new placeholder is invalid for update operations")), nil
	}
	bodyRequired := req.Update.Type == UpdateTypeNamedSectionReplace
	bodyForbidden := req.Update.Type == UpdateTypeMetadataBlockReplace
	if bodyForbidden && (req.Body != nil || req.BodyCacheID != "") {
		return failedProposalResponse(nil, nil, authoringDiagnostic(ErrorCodeInvalidBodySource, "metadata_block_replace must not include body or body_cache_id")), nil
	}
	body, bodyCache, diagnostics, ok := resolveBodySource(store, req.Body, req.BodyCacheID, bodyRequired)
	if !ok {
		return failedProposalResponse(nil, nil, diagnostics...), nil
	}
	prep, err := prepareUpdate(ctx, cfg, idx, req, body)
	if err != nil {
		if body != nil {
			bodyCache = store.cacheBody(*body)
		}
		return failedProposalResponse(bodyCache, nil, authoringDiagnostic(ErrorCodeInvalidRequest, err.Error())), nil
	}
	if len(prep.diagnostics) > 0 {
		if body != nil {
			bodyCache = store.cacheBody(*body)
		}
		return failedProposalResponse(bodyCache, nil, prep.diagnostics...), nil
	}
	return persistProposal(ctx, cfg, idx, store, ProposalOperationUpdate, prep, bodyCache)
}

func GetProposedWrite(ctx context.Context, store *AuthoringStore, req GetProposedWriteRequest) (GetProposedWriteResponse, error) {
	if err := ctx.Err(); err != nil {
		return GetProposedWriteResponse{}, err
	}
	proposal, diag := store.lookupProposal(req.ProposalID)
	if diag != nil {
		return GetProposedWriteResponse{Diagnostics: []Diagnostic{*diag}, Validation: ValidateRecordsResponse{OK: false, Diagnostics: []Diagnostic{*diag}}}, nil
	}
	target := proposal.Target
	expiresAt := proposal.ExpiresAt
	diff := proposal.Diff
	return GetProposedWriteResponse{
		ProposalID:              proposal.ProposalID,
		State:                   proposal.State,
		Operation:               proposal.Operation,
		TargetKind:              proposal.TargetKind,
		Target:                  &target,
		ExpiresAt:               &expiresAt,
		RetentionDays:           proposal.RetentionDays,
		Diff:                    &diff,
		Validation:              proposal.Validation,
		Diagnostics:             cloneDiagnostics(proposal.Diagnostics),
		Note:                    proposal.Note,
		RequiredFollowUpUpdates: append([]RequiredFollowUpUpdate{}, proposal.RequiredFollowUpUpdates...),
	}, nil
}

func AcceptProposedWrite(ctx context.Context, cfg Config, idx *Index, store *AuthoringStore, req AcceptProposedWriteRequest) (AcceptProposedWriteResponse, error) {
	if err := ctx.Err(); err != nil {
		return AcceptProposedWriteResponse{}, err
	}
	if idx == nil {
		return AcceptProposedWriteResponse{}, newToolError(ErrorCodeInvalidRequest, "index is nil")
	}
	proposal, diag := store.lookupProposal(req.ProposalID)
	if diag != nil {
		return acceptRejected(req.ProposalID, "", *diag), nil
	}
	if proposal.State == ProposalStateDiscarded {
		return acceptRejected(proposal.ProposalID, proposal.State, authoringDiagnostic(ErrorCodeProposalDiscarded, "proposal was discarded")), nil
	}
	if proposal.State == ProposalStateAccepted {
		return acceptRejected(proposal.ProposalID, proposal.State, authoringDiagnostic(ErrorCodeProposalAlreadyAccepted, "proposal was already accepted")), nil
	}
	if proposal.State == ProposalStateFailedFinal {
		return acceptRejected(proposal.ProposalID, proposal.State, authoringDiagnostic(ErrorCodeInvalidRequest, "proposal had a partial write failure and cannot be accepted again")), nil
	}
	preWriteValidation := proposal.Validation
	if len(proposal.RequiredFollowUpUpdates) > 0 {
		if !requiredFollowUpsSatisfied(idx, proposal.RequiredFollowUpUpdates) {
			return acceptRejected(proposal.ProposalID, proposal.State, authoringDiagnostic(ErrorCodeRequiredFollowUpNotSatisfied, "required follow-up updates are not satisfied")), nil
		}
		revalidated, err := validateProposedFiles(ctx, idx, proposal.Files)
		if err != nil {
			return AcceptProposedWriteResponse{}, err
		}
		preWriteValidation = revalidated
	}
	if hasErrorDiagnostics(preWriteValidation.Diagnostics) {
		return acceptRejected(proposal.ProposalID, proposal.State, authoringDiagnostic(ErrorCodeInvalidRequest, "proposal validation has error diagnostics")), nil
	}
	if diagnostics := acceptTimeDiagnostics(cfg, idx, proposal); len(diagnostics) > 0 {
		return AcceptProposedWriteResponse{
			ProposalID:     proposal.ProposalID,
			State:          proposal.State,
			Written:        false,
			FilesWritten:   []WrittenFile{},
			Validation:     preWriteValidation,
			RepairGuidance: []string{},
			Diagnostics:    diagnostics,
		}, nil
	}

	var written []WrittenFile
	for _, file := range proposal.Files {
		abs := filepath.Join(cfg.Root, filepath.FromSlash(file.Path))
		if err := os.MkdirAll(filepath.Dir(abs), 0o755); err != nil {
			diagnostic := authoringDiagnostic(ErrorCodeInvalidRequest, fmt.Sprintf("create parent directory for %s: %v", file.Path, err))
			if len(written) > 0 {
				return acceptPartialWriteFailure(store, proposal, written, diagnostic), nil
			}
			return acceptRejected(proposal.ProposalID, proposal.State, diagnostic), nil
		}
		if err := os.WriteFile(abs, []byte(file.Content), 0o644); err != nil {
			diagnostic := authoringDiagnostic(ErrorCodeInvalidRequest, fmt.Sprintf("partial write failure after earlier files were written; write %s: %v", file.Path, err))
			if len(written) > 0 {
				return acceptPartialWriteFailure(store, proposal, written, diagnostic), nil
			}
			return acceptRejected(proposal.ProposalID, proposal.State, diagnostic), nil
		}
		written = append(written, WrittenFile{Path: file.Path, RecordID: file.RecordID, RecordKind: file.RecordKind})
	}
	store.markAccepted(proposal.ProposalID)

	nextIdx, err := BuildIndex(ctx, cfg)
	if err != nil {
		return AcceptProposedWriteResponse{
			ProposalID:     proposal.ProposalID,
			State:          ProposalStateAccepted,
			Written:        true,
			FilesWritten:   written,
			Validation:     ValidateRecordsResponse{OK: false, Diagnostics: []Diagnostic{authoringDiagnostic(ErrorCodeInvalidRequest, fmt.Sprintf("post-write index rebuild failed: %v", err))}},
			RepairGuidance: []string{"Create a repair proposal or inspect the written files manually; the accepted write was not rolled back."},
			Diagnostics:    []Diagnostic{},
		}, nil
	}
	validation, err := ValidateRecords(ctx, nextIdx, ValidateRecordsRequest{})
	if err != nil {
		return AcceptProposedWriteResponse{}, err
	}
	return AcceptProposedWriteResponse{
		ProposalID:     proposal.ProposalID,
		State:          ProposalStateAccepted,
		Written:        true,
		FilesWritten:   written,
		Validation:     validation,
		RepairGuidance: repairGuidance(validation.Diagnostics),
		Diagnostics:    []Diagnostic{},
	}, nil
}

func DiscardProposedWrite(ctx context.Context, store *AuthoringStore, req DiscardProposedWriteRequest) (DiscardProposedWriteResponse, error) {
	if err := ctx.Err(); err != nil {
		return DiscardProposedWriteResponse{}, err
	}
	proposal, diag := store.lookupProposal(req.ProposalID)
	if diag != nil {
		return DiscardProposedWriteResponse{ProposalID: req.ProposalID, State: "", Discarded: false, Written: false, Diagnostics: []Diagnostic{*diag}}, nil
	}
	if proposal.State == ProposalStateAccepted {
		return DiscardProposedWriteResponse{ProposalID: proposal.ProposalID, State: proposal.State, Discarded: false, Written: false, Diagnostics: []Diagnostic{authoringDiagnostic(ErrorCodeProposalAlreadyAccepted, "accepted proposal cannot be discarded")}}, nil
	}
	store.markDiscarded(proposal.ProposalID)
	return DiscardProposedWriteResponse{ProposalID: proposal.ProposalID, State: ProposalStateDiscarded, Discarded: true, Written: false, Diagnostics: []Diagnostic{}}, nil
}

func persistProposal(ctx context.Context, cfg Config, idx *Index, store *AuthoringStore, operation string, prep authoringPreparation, bodyCache *BodyCacheEntry) (ProposeRecordResponse, error) {
	validation, err := validateProposedFiles(ctx, idx, prep.files)
	if err != nil {
		return ProposeRecordResponse{}, err
	}
	now := store.currentTime()
	expiresAt := now.Add(authoringRetentionDays * 24 * time.Hour)
	diff := buildDiff(prep.files)
	proposal := &StoredProposal{
		State:                   ProposalStateProposed,
		Operation:               operation,
		TargetKind:              prep.target.Kind,
		Target:                  prep.target,
		ExpiresAt:               expiresAt,
		RetentionDays:           authoringRetentionDays,
		Diff:                    diff,
		Validation:              validation,
		Diagnostics:             cloneDiagnostics(prep.diagnostics),
		Note:                    noWriteProposalNote,
		Files:                   append([]ProposedFile{}, prep.files...),
		RequiredFollowUpUpdates: append([]RequiredFollowUpUpdate{}, prep.requiredFollowUpUpdates...),
		ProposalCreatedAt:       now,
	}
	store.saveProposal(proposal)
	target := proposal.Target
	return ProposeRecordResponse{
		ProposalCreated:         true,
		ProposalID:              proposal.ProposalID,
		State:                   proposal.State,
		Operation:               proposal.Operation,
		TargetKind:              proposal.TargetKind,
		Target:                  &target,
		ExpiresAt:               &expiresAt,
		RetentionDays:           proposal.RetentionDays,
		Diff:                    &diff,
		Validation:              proposal.Validation,
		Diagnostics:             cloneDiagnostics(proposal.Diagnostics),
		Note:                    proposal.Note,
		BodyCache:               bodyCache,
		RequiredFollowUpUpdates: append([]RequiredFollowUpUpdate{}, proposal.RequiredFollowUpUpdates...),
	}, nil
}

func prepareCreate(ctx context.Context, cfg Config, idx *Index, req ProposeRecordCreateRequest, body *string) (authoringPreparation, error) {
	if strings.TrimSpace(req.ID) == "" {
		return authoringPreparation{}, fmt.Errorf("id is required")
	}
	if strings.TrimSpace(req.Title) == "" && body == nil {
		return authoringPreparation{}, fmt.Errorf("title is required for template-driven create")
	}
	mode := req.ReciprocalUpdateMode
	if mode == "" {
		mode = "include_required"
	}
	if mode != "include_required" && mode != "report_required_follow_up" {
		return authoringPreparation{}, fmt.Errorf("unsupported reciprocal_update_mode %q", req.ReciprocalUpdateMode)
	}
	resolved, domain, err := resolveCreateID(idx, req.Kind, req.ID, req.Domain, req.ParentID)
	if err != nil {
		return authoringPreparation{}, err
	}
	if findRecordByIDKind(idx, resolved, req.Kind) != nil {
		return authoringPreparation{}, fmt.Errorf("record %s already exists", resolved)
	}

	content := ""
	if body != nil {
		content = *body
	} else {
		var renderErr error
		content, renderErr = renderCreateBody(idx, req.Kind, resolved, req.Title, req.Fields, req.ParentID)
		if renderErr != nil {
			return authoringPreparation{}, renderErr
		}
	}
	path := createRecordPath(req.Kind, resolved, req.Title, domain)
	file := ProposedFile{Path: path, Change: "create", RecordID: resolved, RecordKind: req.Kind, BaseHash: "", Content: ensureTrailingNewline(content)}
	prep := authoringPreparation{
		target: AuthoringTarget{RequestedID: req.ID, ResolvedID: resolved, Kind: req.Kind, Domain: domain, ParentID: req.ParentID, Path: path},
		files:  []ProposedFile{file},
	}
	reciprocalFiles, followUps, err := requiredReciprocalUpdates(ctx, cfg, idx, req.Kind, resolved, req.Fields, req.ParentID, mode)
	if err != nil {
		return authoringPreparation{}, err
	}
	prep.files = append(prep.files, reciprocalFiles...)
	prep.requiredFollowUpUpdates = append(prep.requiredFollowUpUpdates, followUps...)
	return prep, nil
}

func prepareUpdate(ctx context.Context, cfg Config, idx *Index, req ProposeRecordUpdateRequest, body *string) (authoringPreparation, error) {
	if req.ID == "" {
		return authoringPreparation{}, fmt.Errorf("id is required")
	}
	record := findRecordByIDKind(idx, req.ID, req.Kind)
	if record == nil {
		return authoringPreparation{}, fmt.Errorf("record %s was not found", req.ID)
	}
	raw, err := readRepoFile(cfg, record.Path)
	if err != nil {
		return authoringPreparation{}, err
	}
	var updated string
	var diagnostics []Diagnostic
	switch req.Update.Type {
	case UpdateTypeMetadataBlockReplace:
		updated, diagnostics, err = replaceMetadataBlock(*record, raw, req.Update.Metadata)
	case UpdateTypeNamedSectionReplace:
		if req.Update.SectionSelector == nil {
			return authoringPreparation{}, fmt.Errorf("section_selector is required for named_section_replace")
		}
		if body == nil {
			return authoringPreparation{}, fmt.Errorf("body is required for named_section_replace")
		}
		updated, diagnostics, err = replaceNamedSection(raw, *req.Update.SectionSelector, *body)
	default:
		return authoringPreparation{}, fmt.Errorf("unsupported update.type %q", req.Update.Type)
	}
	if err != nil {
		return authoringPreparation{}, err
	}
	if len(diagnostics) > 0 {
		return authoringPreparation{diagnostics: diagnostics}, nil
	}
	file := ProposedFile{
		Path:       record.Path,
		Change:     "modify",
		RecordID:   record.ID,
		RecordKind: record.Kind,
		BaseHash:   contentHash(raw),
		BaseID:     record.ID,
		BaseKind:   record.Kind,
		Content:    ensureTrailingNewline(updated),
	}
	return authoringPreparation{
		target: AuthoringTarget{
			RequestedID: req.ID,
			ResolvedID:  record.ID,
			Kind:        record.Kind,
			Domain:      workflowDomain(record.ID),
			Path:        record.Path,
		},
		files: []ProposedFile{file},
	}, nil
}

func validateCreateKind(kind RecordKind) error {
	switch kind {
	case RecordKindDecision, RecordKindRequirement, RecordKindWorkItem, RecordKindTask:
		return nil
	case RecordKindSpec:
		return fmt.Errorf("SPEC-new and spec skeleton creation are outside the MVP")
	case RecordKindInvestigation:
		return fmt.Errorf("investigation creation is outside the MVP")
	default:
		return fmt.Errorf("unsupported create kind %q", kind)
	}
}

func validateUpdateKind(kind RecordKind) error {
	switch kind {
	case RecordKindDecision, RecordKindSpec, RecordKindRequirement, RecordKindWorkItem, RecordKindTask:
		return nil
	case RecordKindInvestigation:
		return fmt.Errorf("investigation update is outside the MVP")
	default:
		return fmt.Errorf("unsupported update kind %q", kind)
	}
}

func resolveBodySource(store *AuthoringStore, body *string, bodyCacheID string, required bool) (*string, *BodyCacheEntry, []Diagnostic, bool) {
	if body != nil && bodyCacheID != "" {
		return nil, nil, []Diagnostic{authoringDiagnostic(ErrorCodeInvalidBodySource, "body and body_cache_id are mutually exclusive")}, false
	}
	if body == nil && bodyCacheID == "" {
		if required {
			return nil, nil, []Diagnostic{authoringDiagnostic(ErrorCodeInvalidBodySource, "body or body_cache_id is required")}, false
		}
		return nil, nil, nil, true
	}
	if body != nil {
		entry := store.cacheBody(*body)
		resolved := *body
		return &resolved, entry, nil, true
	}
	entry, diag := store.lookupBody(bodyCacheID)
	if diag != nil {
		return nil, nil, []Diagnostic{*diag}, false
	}
	bodyValue := entry.Body
	publicEntry := entry
	publicEntry.Body = ""
	return &bodyValue, &publicEntry, nil, true
}

func resolveCreateID(idx *Index, kind RecordKind, requestedID, requestedDomain, parentID string) (string, string, error) {
	switch kind {
	case RecordKindDecision:
		if createDecisionPlaceholderPattern.MatchString(requestedID) {
			next, _ := nextDecisionID(idx)
			return next, "", nil
		}
		if _, ok := decisionRecordNumber(requestedID); ok {
			return requestedID, "", nil
		}
		return "", "", fmt.Errorf("invalid decision create ID %q", requestedID)
	case RecordKindRequirement:
		if match := createRequirementPlaceholderPattern.FindStringSubmatch(requestedID); match != nil {
			domain := match[1]
			if requestedDomain != "" && requestedDomain != domain {
				return "", "", fmt.Errorf("domain %q does not match ID domain %q", requestedDomain, domain)
			}
			return nextWorkflowID(idx, RecordKindRequirement, domain), domain, nil
		}
		if !validWorkflowIDForKind(requestedID, kind) {
			return "", "", fmt.Errorf("invalid requirement create ID %q", requestedID)
		}
		domain := workflowDomain(requestedID)
		if requestedDomain != "" && requestedDomain != domain {
			return "", "", fmt.Errorf("domain %q does not match ID domain %q", requestedDomain, domain)
		}
		return requestedID, domain, nil
	case RecordKindWorkItem:
		if match := createWorkItemPlaceholderPattern.FindStringSubmatch(requestedID); match != nil {
			domain := match[1]
			if requestedDomain != "" && requestedDomain != domain {
				return "", "", fmt.Errorf("domain %q does not match ID domain %q", requestedDomain, domain)
			}
			return nextWorkflowID(idx, RecordKindWorkItem, domain), domain, nil
		}
		if !validWorkflowIDForKind(requestedID, kind) {
			return "", "", fmt.Errorf("invalid work item create ID %q", requestedID)
		}
		domain := workflowDomain(requestedID)
		if requestedDomain != "" && requestedDomain != domain {
			return "", "", fmt.Errorf("domain %q does not match ID domain %q", requestedDomain, domain)
		}
		return requestedID, domain, nil
	case RecordKindTask:
		if strings.TrimSpace(parentID) == "" {
			return "", "", fmt.Errorf("parent_id is required for task create")
		}
		parent := findRecordByIDKind(idx, parentID, RecordKindWorkItem)
		if parent == nil {
			return "", "", fmt.Errorf("parent work item %s was not found", parentID)
		}
		parentDomain := workflowDomain(parentID)
		parentSeq := workflowSequence(parentID)
		if match := createTaskPlaceholderPattern.FindStringSubmatch(requestedID); match != nil {
			if match[1] != parentDomain || match[2] != parentSeq {
				return "", "", fmt.Errorf("task placeholder ID must match parent work item domain and sequence")
			}
			if requestedDomain != "" && requestedDomain != parentDomain {
				return "", "", fmt.Errorf("domain %q does not match parent domain %q", requestedDomain, parentDomain)
			}
			return nextTaskID(idx, parentDomain, parentSeq), parentDomain, nil
		}
		if !validWorkflowIDForKind(requestedID, kind) {
			return "", "", fmt.Errorf("invalid task create ID %q", requestedID)
		}
		if workflowDomain(requestedID) != parentDomain || workflowSequence(requestedID) != parentSeq {
			return "", "", fmt.Errorf("task ID must match parent work item domain and sequence")
		}
		return requestedID, parentDomain, nil
	default:
		return "", "", fmt.Errorf("unsupported kind %q", kind)
	}
}

func renderCreateBody(idx *Index, kind RecordKind, id, title string, fields map[string]any, parentID string) (string, error) {
	switch kind {
	case RecordKindDecision:
		num := strings.TrimPrefix(id, "ADR-")
		meta, err := renderADRMetadata(fields)
		if err != nil {
			return "", err
		}
		return "# " + num + ": " + title + "\n\n" + meta + "\n", nil
	case RecordKindRequirement:
		meta, err := renderRequirementMetadata(id, fields)
		if err != nil {
			return "", err
		}
		return "# " + id + ": " + title + "\n\n" + meta + "\n\n## Requirement\n\n## Evidence\n", nil
	case RecordKindWorkItem:
		meta, err := renderWorkItemMetadata(id, fields)
		if err != nil {
			return "", err
		}
		return "# " + id + ": " + title + "\n\n" + meta + "\n\n## Goal\n\n## Boundary\n\n## Evidence\n", nil
	case RecordKindTask:
		if scalarField(fields, "work_item") != parentID {
			return "", fmt.Errorf("task create requires explicit fields.work_item equal to parent_id")
		}
		meta, err := renderTaskMetadata(id, fields)
		if err != nil {
			return "", err
		}
		return "# " + id + ": " + title + "\n\n" + meta + "\n\n## Goal\n\n## Work\n\n## Done condition\n\n## Verification\n\n## Evidence\n", nil
	default:
		return "", fmt.Errorf("unsupported kind %q", kind)
	}
}

func requiredReciprocalUpdates(ctx context.Context, cfg Config, idx *Index, kind RecordKind, id string, fields map[string]any, parentID, mode string) ([]ProposedFile, []RequiredFollowUpUpdate, error) {
	switch kind {
	case RecordKindWorkItem:
		reqID := scalarField(fields, "source_requirement")
		if reqID == "" {
			return nil, nil, nil
		}
		parent := findRecordByIDKind(idx, reqID, RecordKindRequirement)
		if parent == nil || parent.Requirement == nil || containsString(parent.Requirement.WorkItems, id) {
			return nil, nil, nil
		}
		followUp := RequiredFollowUpUpdate{RecordID: parent.ID, Kind: RecordKindRequirement, Field: "work_items", Value: id, Message: "add new work item to source requirement work_items"}
		if mode == "report_required_follow_up" {
			return nil, []RequiredFollowUpUpdate{followUp}, nil
		}
		return []ProposedFile{reciprocalMetadataFile(cfg, *parent, append(parent.Requirement.WorkItems, id))}, nil, nil
	case RecordKindTask:
		parent := findRecordByIDKind(idx, parentID, RecordKindWorkItem)
		if parent == nil || parent.WorkItem == nil || containsString(parent.WorkItem.Tasks, id) {
			return nil, nil, nil
		}
		followUp := RequiredFollowUpUpdate{RecordID: parent.ID, Kind: RecordKindWorkItem, Field: "tasks", Value: id, Message: "add new task to parent work item tasks"}
		if mode == "report_required_follow_up" {
			return nil, []RequiredFollowUpUpdate{followUp}, nil
		}
		return []ProposedFile{reciprocalMetadataFile(cfg, *parent, append(parent.WorkItem.Tasks, id))}, nil, nil
	default:
		return nil, nil, nil
	}
}

func reciprocalMetadataFile(cfg Config, record Record, values []string) ProposedFile {
	raw, _ := readRepoFile(cfg, record.Path)
	fields := map[string]any{}
	switch record.Kind {
	case RecordKindRequirement:
		fields = map[string]any{
			"id":          record.ID,
			"status":      string(record.Status),
			"date":        workflowMetadataScalar(record, "date"),
			"source_refs": record.Requirement.SourceRefs,
			"work_items":  values,
		}
	case RecordKindWorkItem:
		fields = map[string]any{
			"id":                 record.ID,
			"status":             string(record.Status),
			"date":               workflowMetadataScalar(record, "date"),
			"source_requirement": record.WorkItem.SourceRequirement,
			"impact_refs":        record.WorkItem.ImpactRefs,
			"tasks":              values,
		}
	}
	updated, _, _ := replaceMetadataBlock(record, raw, fields)
	return ProposedFile{
		Path:       record.Path,
		Change:     "modify",
		RecordID:   record.ID,
		RecordKind: record.Kind,
		BaseHash:   contentHash(raw),
		BaseID:     record.ID,
		BaseKind:   record.Kind,
		Content:    ensureTrailingNewline(updated),
	}
}

func replaceMetadataBlock(record Record, raw string, metadata map[string]any) (string, []Diagnostic, error) {
	if metadata == nil {
		return "", nil, fmt.Errorf("metadata is required")
	}
	var rendered string
	var err error
	switch record.Kind {
	case RecordKindDecision:
		rendered, err = renderADRMetadata(metadata)
	case RecordKindRequirement:
		rendered, err = renderRequirementMetadata(record.ID, metadata)
	case RecordKindWorkItem:
		rendered, err = renderWorkItemMetadata(record.ID, metadata)
	case RecordKindTask:
		rendered, err = renderTaskMetadata(record.ID, metadata)
	case RecordKindSpec:
		return replaceSpecMetadata(record.ID, raw, metadata)
	default:
		return "", nil, fmt.Errorf("metadata replacement does not support kind %q", record.Kind)
	}
	if err != nil {
		return "", []Diagnostic{metadataRenderErrorDiagnostic(err)}, nil
	}
	lines := splitMarkdownLines(raw)
	if len(lines) == 0 {
		return "", nil, fmt.Errorf("record body is empty")
	}
	start := 1
	end := len(lines)
	for i := 1; i < len(lines); i++ {
		line := trimLineEnd(lines[i])
		if strings.HasPrefix(line, "##") || strings.HasPrefix(line, ">") {
			end = i
			break
		}
	}
	replacement := []string{"", rendered, ""}
	out := append([]string{}, lines[:start]...)
	out = append(out, replacement...)
	out = append(out, lines[end:]...)
	return strings.Join(out, "\n"), nil, nil
}

func replaceSpecMetadata(expectedID, raw string, metadata map[string]any) (string, []Diagnostic, error) {
	fm, rest, ok := extractFrontMatter(raw)
	if !ok {
		return "", []Diagnostic{authoringDiagnostic(ErrorCodeInvalidRequest, "spec front matter is required")}, nil
	}
	var node map[string]any
	if err := yaml.Unmarshal([]byte(fm), &node); err != nil {
		return "", nil, err
	}
	if node == nil {
		node = map[string]any{}
	}
	design, _ := node["design_record"].(map[string]any)
	if design == nil {
		design = map[string]any{}
	}
	required := []string{"scope", "status"}
	for _, key := range required {
		if strings.TrimSpace(scalarField(metadata, key)) == "" {
			return "", []Diagnostic{missingMetadataDiagnostic(key)}, nil
		}
		node[key] = scalarField(metadata, key)
	}
	inDesign, _ := metadata["design_record"].(map[string]any)
	for _, key := range []string{"id", "kind", "status"} {
		value := scalarField(inDesign, key)
		if strings.TrimSpace(value) == "" {
			return "", []Diagnostic{missingMetadataDiagnostic("design_record." + key)}, nil
		}
		if key == "id" && value != expectedID {
			return "", []Diagnostic{{
				Category: DiagnosticInvalidMetadataValue,
				Severity: DiagnosticSeverityError,
				Field:    "design_record.id",
				Value:    value,
				Message:  fmt.Sprintf("design_record.id must remain %s", expectedID),
			}}, nil
		}
		if key == "kind" && value != string(RecordKindSpec) {
			return "", []Diagnostic{{
				Category: DiagnosticInvalidMetadataValue,
				Severity: DiagnosticSeverityError,
				Field:    "design_record.kind",
				Value:    value,
				Message:  "design_record.kind must be spec",
			}}, nil
		}
		design[key] = value
	}
	deps, ok := listField(inDesign, "depends_on")
	if !ok {
		return "", []Diagnostic{missingMetadataDiagnostic("design_record.depends_on")}, nil
	}
	design["depends_on"] = deps
	node["design_record"] = design
	rendered, err := yaml.Marshal(node)
	if err != nil {
		return "", nil, err
	}
	return "---\n" + string(rendered) + "---\n" + strings.TrimPrefix(rest, "\n"), nil, nil
}

func replaceNamedSection(raw string, selector SectionSelector, body string) (string, []Diagnostic, error) {
	if strings.TrimSpace(selector.Heading) == "" {
		return "", nil, fmt.Errorf("section_selector.heading is required")
	}
	if selector.Match != "" && selector.Match != "exact" {
		return "", nil, fmt.Errorf("section_selector.match supports exact only")
	}
	if selector.Level != nil && (*selector.Level < 1 || *selector.Level > 6) {
		return "", nil, fmt.Errorf("section_selector.level must be 1 through 6")
	}
	sections := markdownSections(raw)
	var matches []markdownSection
	for _, section := range sections {
		if section.Heading.Text != selector.Heading {
			continue
		}
		if selector.Level != nil && section.Heading.Level != *selector.Level {
			continue
		}
		matches = append(matches, section)
	}
	if len(matches) == 0 {
		return "", []Diagnostic{sectionSelectorDiagnostic(ErrorCodeSectionSelectorNoMatch, "section selector matched no ATX section", sections)}, nil
	}
	if len(matches) > 1 {
		return "", []Diagnostic{sectionSelectorDiagnostic(ErrorCodeSectionSelectorAmbiguous, "section selector matched multiple ATX sections", matches)}, nil
	}
	match := matches[0]
	lines := splitMarkdownLines(raw)
	replacementBody := strings.TrimSuffix(body, "\n")
	replacement := []string{strings.Repeat("#", match.Heading.Level) + " " + match.Heading.Text}
	if replacementBody != "" {
		replacement = append(replacement, replacementBody)
	}
	out := append([]string{}, lines[:match.StartLine]...)
	out = append(out, replacement...)
	out = append(out, lines[match.EndLine:]...)
	return strings.Join(out, "\n"), nil, nil
}

type markdownSection struct {
	Heading   Heading
	StartLine int
	EndLine   int
	Ordinal   int
}

func markdownSections(raw string) []markdownSection {
	lines := splitMarkdownLines(raw)
	start := 0
	if hasOpeningFrontMatter(lines) {
		for i := 1; i < len(lines); i++ {
			if strings.TrimSpace(trimLineEnd(lines[i])) == "---" {
				start = i + 1
				break
			}
		}
	}
	type foundHeading struct {
		Heading
		Line int
	}
	var headings []foundHeading
	inFence := false
	fenceMarker := ""
	for i := start; i < len(lines); i++ {
		trimmed := strings.TrimSpace(trimLineEnd(lines[i]))
		if isFenceLine(trimmed) {
			marker := fencePrefix(trimmed)
			if !inFence {
				inFence = true
				fenceMarker = marker
			} else if marker == fenceMarker {
				inFence = false
				fenceMarker = ""
			}
			continue
		}
		if inFence {
			continue
		}
		match := atxHeadingPattern.FindStringSubmatch(trimLineEnd(lines[i]))
		if match == nil {
			continue
		}
		headings = append(headings, foundHeading{Heading: Heading{Level: len(match[1]), Text: strings.TrimSpace(match[2])}, Line: i})
	}
	sections := make([]markdownSection, 0, len(headings))
	for i, h := range headings {
		end := len(lines)
		for _, next := range headings[i+1:] {
			if next.Level <= h.Level {
				end = next.Line
				break
			}
		}
		sections = append(sections, markdownSection{Heading: h.Heading, StartLine: h.Line, EndLine: end, Ordinal: i + 1})
	}
	return sections
}

func validateProposedFiles(ctx context.Context, idx *Index, files []ProposedFile) (ValidateRecordsResponse, error) {
	hyp := &Index{
		Root:               idx.Root,
		Records:            []Record{},
		Candidates:         append([]RecordCandidate{}, idx.Candidates...),
		ParseIssues:        append([]ParseIssue{}, idx.ParseIssues...),
		PathIssues:         append([]PathIssue{}, idx.PathIssues...),
		SemanticRefs:       []SemanticRefDecl{},
		SemanticRefSources: []SemanticRefSource{},
	}
	replaced := map[string]bool{}
	for _, file := range files {
		replaced[file.Path] = true
	}
	for _, record := range idx.Records {
		if !replaced[record.Path] {
			hyp.Records = append(hyp.Records, record)
		}
	}
	for _, file := range files {
		record, candidate, issues := parseRecordByPath(file.Path, file.Content)
		if candidate.Path != "" {
			hyp.Candidates = append(hyp.Candidates, candidate)
		}
		hyp.ParseIssues = append(hyp.ParseIssues, issues...)
		if record != nil {
			hyp.Records = append(hyp.Records, *record)
			if record.Kind == RecordKindSpec {
				if source, ok := parseSpecSemanticRefSource(file.Path, file.Content); ok {
					hyp.SemanticRefSources = append(hyp.SemanticRefSources, source)
				}
			}
		}
	}
	for _, record := range hyp.Records {
		hyp.SemanticRefs = append(hyp.SemanticRefs, record.SemanticRefs...)
	}
	return ValidateRecords(ctx, hyp, ValidateRecordsRequest{})
}

func parseRecordByPath(path, content string) (*Record, RecordCandidate, []ParseIssue) {
	switch {
	case strings.HasPrefix(path, "docs/adr/"):
		return parseADRRecord(path, content)
	case strings.HasPrefix(path, "docs/spec/"):
		return parseSpecRecord(path, content)
	case strings.HasPrefix(path, "docs/requirements/"):
		return parseRequirementRecord(path, content)
	case strings.HasPrefix(path, "docs/work-items/"):
		return parseWorkItemRecord(path, content)
	case strings.HasPrefix(path, "docs/tasks/"):
		return parseTaskRecord(path, content)
	default:
		return nil, RecordCandidate{}, nil
	}
}

func acceptTimeDiagnostics(cfg Config, idx *Index, proposal *StoredProposal) []Diagnostic {
	var diagnostics []Diagnostic
	for _, file := range proposal.Files {
		switch file.Change {
		case "create":
			if findRecordByIDKind(idx, file.RecordID, file.RecordKind) != nil {
				diagnostics = append(diagnostics, authoringDiagnostic(ErrorCodeIDCollision, fmt.Sprintf("record %s already exists", file.RecordID)))
			}
			if _, err := os.Stat(filepath.Join(cfg.Root, filepath.FromSlash(file.Path))); err == nil {
				diagnostics = append(diagnostics, authoringDiagnostic(ErrorCodeIDCollision, fmt.Sprintf("target path %s already exists", file.Path)))
			}
		case "modify":
			current := findRecordByIDKind(idx, file.BaseID, file.BaseKind)
			if current == nil || current.Path != file.Path {
				diagnostics = append(diagnostics, authoringDiagnostic(ErrorCodeTargetChanged, fmt.Sprintf("target record %s no longer resolves to the same path", file.BaseID)))
				continue
			}
			raw, err := readRepoFile(cfg, file.Path)
			if err != nil {
				diagnostics = append(diagnostics, authoringDiagnostic(ErrorCodeTargetChanged, fmt.Sprintf("read current target %s: %v", file.Path, err)))
				continue
			}
			if contentHash(raw) != file.BaseHash {
				diagnostics = append(diagnostics, authoringDiagnostic(ErrorCodeProposalStale, fmt.Sprintf("target file %s changed after proposal creation", file.Path)))
			}
		}
	}
	return diagnostics
}

func requiredFollowUpsSatisfied(idx *Index, followUps []RequiredFollowUpUpdate) bool {
	for _, followUp := range followUps {
		record := findRecordByIDKind(idx, followUp.RecordID, followUp.Kind)
		if record == nil {
			return false
		}
		switch followUp.Field {
		case "work_items":
			if record.Requirement == nil || !containsString(record.Requirement.WorkItems, followUp.Value) {
				return false
			}
		case "tasks":
			if record.WorkItem == nil || !containsString(record.WorkItem.Tasks, followUp.Value) {
				return false
			}
		default:
			return false
		}
	}
	return true
}

func buildDiff(files []ProposedFile) Diff {
	diffFiles := make([]DiffFile, 0, len(files))
	var text strings.Builder
	for _, file := range files {
		diffFiles = append(diffFiles, DiffFile{Path: file.Path, Change: file.Change, RecordID: file.RecordID, RecordKind: file.RecordKind})
		oldPath := file.Path
		if file.Change == "create" {
			oldPath = "/dev/null"
		}
		text.WriteString("--- " + oldPath + "\n")
		text.WriteString("+++ " + file.Path + "\n")
		for _, line := range strings.Split(strings.TrimSuffix(file.Content, "\n"), "\n") {
			text.WriteString("+" + line + "\n")
		}
	}
	return Diff{Format: "unified", Files: diffFiles, Text: text.String()}
}

func failedProposalResponse(bodyCache *BodyCacheEntry, validation *ValidateRecordsResponse, diagnostics ...Diagnostic) ProposeRecordResponse {
	if validation == nil {
		v := ValidateRecordsResponse{OK: false, Diagnostics: diagnostics}
		validation = &v
	}
	return ProposeRecordResponse{
		ProposalCreated: false,
		Validation:      *validation,
		Diagnostics:     diagnostics,
		BodyCache:       bodyCache,
	}
}

func (s *AuthoringStore) currentTime() time.Time {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.now == nil {
		return time.Now()
	}
	return s.now()
}

func (s *AuthoringStore) saveProposal(proposal *StoredProposal) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.cleanupLocked()
	s.nextID++
	proposal.ProposalID = fmt.Sprintf("pw_%06d", s.nextID)
	s.proposals[proposal.ProposalID] = proposal
}

func (s *AuthoringStore) lookupProposal(id string) (*StoredProposal, *Diagnostic) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if strings.TrimSpace(id) == "" {
		diag := authoringDiagnostic(ErrorCodeProposalNotFound, "proposal_id is required")
		return nil, &diag
	}
	proposal, ok := s.proposals[id]
	if !ok {
		diag := authoringDiagnostic(ErrorCodeProposalNotFound, fmt.Sprintf("proposal %s was not found", id))
		return nil, &diag
	}
	if !s.nowLocked().Before(proposal.ExpiresAt) {
		delete(s.proposals, id)
		diag := authoringDiagnostic(ErrorCodeProposalExpired, fmt.Sprintf("proposal %s is expired", id))
		return nil, &diag
	}
	copy := *proposal
	copy.Diagnostics = cloneDiagnostics(proposal.Diagnostics)
	copy.Files = append([]ProposedFile{}, proposal.Files...)
	copy.RequiredFollowUpUpdates = append([]RequiredFollowUpUpdate{}, proposal.RequiredFollowUpUpdates...)
	return &copy, nil
}

func (s *AuthoringStore) markAccepted(id string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if proposal, ok := s.proposals[id]; ok {
		proposal.State = ProposalStateAccepted
	}
}

func (s *AuthoringStore) markDiscarded(id string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if proposal, ok := s.proposals[id]; ok {
		proposal.State = ProposalStateDiscarded
	}
}

func (s *AuthoringStore) markFailedFinal(id string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if proposal, ok := s.proposals[id]; ok {
		proposal.State = ProposalStateFailedFinal
	}
}

func (s *AuthoringStore) cacheBody(body string) *BodyCacheEntry {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.cleanupLocked()
	s.nextID++
	entry := BodyCacheEntry{
		BodyCacheID:   fmt.Sprintf("bc_%06d", s.nextID),
		ExpiresAt:     s.nowLocked().Add(authoringRetentionDays * 24 * time.Hour),
		RetentionDays: authoringRetentionDays,
		Body:          body,
	}
	s.bodies[entry.BodyCacheID] = entry
	public := entry
	public.Body = ""
	return &public
}

func (s *AuthoringStore) lookupBody(id string) (BodyCacheEntry, *Diagnostic) {
	s.mu.Lock()
	defer s.mu.Unlock()
	entry, ok := s.bodies[id]
	if !ok {
		diag := authoringDiagnostic(ErrorCodeBodyCacheNotFound, fmt.Sprintf("body cache %s was not found", id))
		return BodyCacheEntry{}, &diag
	}
	if !s.nowLocked().Before(entry.ExpiresAt) {
		delete(s.bodies, id)
		diag := authoringDiagnostic(ErrorCodeBodyCacheExpired, fmt.Sprintf("body cache %s is expired", id))
		return BodyCacheEntry{}, &diag
	}
	return entry, nil
}

func (s *AuthoringStore) cleanupLocked() {
	now := s.nowLocked()
	for id, proposal := range s.proposals {
		if !now.Before(proposal.ExpiresAt) {
			delete(s.proposals, id)
		}
	}
	for id, body := range s.bodies {
		if !now.Before(body.ExpiresAt) {
			delete(s.bodies, id)
		}
	}
}

func (s *AuthoringStore) nowLocked() time.Time {
	if s.now == nil {
		return time.Now()
	}
	return s.now()
}

func authoringDiagnostic(code ErrorCode, message string) Diagnostic {
	return Diagnostic{Category: DiagnosticCategory(code), Severity: DiagnosticSeverityError, Message: message}
}

func missingMetadataDiagnostic(field string) Diagnostic {
	return Diagnostic{Category: DiagnosticMissingRequiredMetadata, Severity: DiagnosticSeverityError, Field: field, Message: fmt.Sprintf("missing required metadata field %s", field)}
}

const missingRequiredMetadataMessagePrefix = "missing required metadata field "

func metadataRenderErrorDiagnostic(err error) Diagnostic {
	message := err.Error()
	if strings.HasPrefix(message, missingRequiredMetadataMessagePrefix) {
		field := strings.TrimSpace(strings.TrimPrefix(message, missingRequiredMetadataMessagePrefix))
		return missingMetadataDiagnostic(field)
	}
	return authoringDiagnostic(ErrorCodeInvalidRequest, message)
}

func sectionSelectorDiagnostic(code ErrorCode, message string, sections []markdownSection) Diagnostic {
	diagnostic := authoringDiagnostic(code, message)
	for _, section := range sections {
		diagnostic.CandidateHeadings = append(diagnostic.CandidateHeadings, CandidateHeading{
			Heading: section.Heading.Text,
			Level:   section.Heading.Level,
			Ordinal: section.Ordinal,
		})
	}
	return diagnostic
}

func acceptRejected(proposalID, state string, diagnostic Diagnostic) AcceptProposedWriteResponse {
	return AcceptProposedWriteResponse{
		ProposalID:     proposalID,
		State:          state,
		Written:        false,
		FilesWritten:   []WrittenFile{},
		Validation:     ValidateRecordsResponse{OK: false, Diagnostics: []Diagnostic{diagnostic}},
		RepairGuidance: []string{},
		Diagnostics:    []Diagnostic{diagnostic},
	}
}

func acceptPartialWriteFailure(store *AuthoringStore, proposal *StoredProposal, written []WrittenFile, diagnostic Diagnostic) AcceptProposedWriteResponse {
	store.markFailedFinal(proposal.ProposalID)
	return AcceptProposedWriteResponse{
		ProposalID:     proposal.ProposalID,
		State:          ProposalStateFailedFinal,
		Written:        true,
		FilesWritten:   written,
		Validation:     ValidateRecordsResponse{OK: false, Diagnostics: []Diagnostic{diagnostic}},
		RepairGuidance: []string{"Inspect and repair the partially written files manually or create a repair proposal; the accepted write was not rolled back and this proposal cannot be retried."},
		Diagnostics:    []Diagnostic{diagnostic},
	}
}

func hasErrorDiagnostics(diagnostics []Diagnostic) bool {
	for _, diagnostic := range diagnostics {
		if diagnostic.Severity == DiagnosticSeverityError {
			return true
		}
	}
	return false
}

func repairGuidance(diagnostics []Diagnostic) []string {
	if !hasErrorDiagnostics(diagnostics) {
		return []string{}
	}
	return []string{"Create a repair proposal for the reported diagnostics; the accepted write was not rolled back."}
}

func cloneDiagnostics(in []Diagnostic) []Diagnostic {
	out := make([]Diagnostic, len(in))
	copy(out, in)
	return out
}

func contentHash(raw string) string {
	sum := sha256.Sum256([]byte(raw))
	return hex.EncodeToString(sum[:])
}

func readRepoFile(cfg Config, rel string) (string, error) {
	data, err := os.ReadFile(filepath.Join(cfg.Root, filepath.FromSlash(rel)))
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func createRecordPath(kind RecordKind, id, title, domain string) string {
	slug := slugifyRecordTitle(title)
	suffix := ".md"
	if slug != "" {
		suffix = "-" + slug + ".md"
	}
	switch kind {
	case RecordKindDecision:
		return "docs/adr/" + strings.TrimPrefix(id, "ADR-") + suffix
	case RecordKindRequirement:
		return "docs/requirements/" + strings.ToLower(domain) + "/" + id + suffix
	case RecordKindWorkItem:
		return "docs/work-items/" + strings.ToLower(domain) + "/" + id + suffix
	case RecordKindTask:
		return "docs/tasks/" + strings.ToLower(domain) + "/" + id + suffix
	default:
		return id + suffix
	}
}

func findRecordByIDKind(idx *Index, id string, kind RecordKind) *Record {
	for i := range idx.Records {
		if idx.Records[i].ID == id && idx.Records[i].Kind == kind {
			return &idx.Records[i]
		}
	}
	return nil
}

func nextDecisionID(idx *Index) (string, int) {
	maxNum := 0
	for _, record := range idx.Records {
		if record.Kind != RecordKindDecision {
			continue
		}
		num, ok := decisionRecordNumber(record.ID)
		if ok && num > maxNum {
			maxNum = num
		}
	}
	return fmt.Sprintf("ADR-%03d", maxNum+1), maxNum + 1
}

func nextWorkflowID(idx *Index, kind RecordKind, domain string) string {
	maxNum := 0
	for _, record := range idx.Records {
		if record.Kind != kind || workflowDomain(record.ID) != domain {
			continue
		}
		seq, err := strconv.Atoi(workflowSequence(record.ID))
		if err == nil && seq > maxNum {
			maxNum = seq
		}
	}
	prefix := "REQ"
	if kind == RecordKindWorkItem {
		prefix = "WORK"
	}
	return fmt.Sprintf("%s-%s-%03d", prefix, domain, maxNum+1)
}

func nextTaskID(idx *Index, domain, workSeq string) string {
	maxNum := 0
	for _, record := range idx.Records {
		if record.Kind != RecordKindTask || workflowDomain(record.ID) != domain || workflowSequence(record.ID) != workSeq {
			continue
		}
		parts := strings.Split(record.ID, "-")
		if len(parts) == 4 {
			seq, err := strconv.Atoi(parts[3])
			if err == nil && seq > maxNum {
				maxNum = seq
			}
		}
	}
	return fmt.Sprintf("TASK-%s-%s-%02d", domain, workSeq, maxNum+1)
}

func workflowDomain(id string) string {
	parts := strings.Split(id, "-")
	if len(parts) < 3 {
		return ""
	}
	return parts[1]
}

func workflowSequence(id string) string {
	parts := strings.Split(id, "-")
	if len(parts) < 3 {
		return ""
	}
	return parts[2]
}

func hasSequenceNewToken(id string, kind RecordKind) bool {
	switch kind {
	case RecordKindDecision:
		return id == "ADR-new"
	case RecordKindRequirement, RecordKindWorkItem:
		parts := strings.Split(id, "-")
		return len(parts) == 3 && parts[2] == "new"
	case RecordKindTask:
		parts := strings.Split(id, "-")
		return len(parts) == 4 && parts[3] == "new"
	default:
		return strings.Contains(id, "-new")
	}
}

func renderADRMetadata(fields map[string]any) (string, error) {
	for _, key := range []string{"status", "date", "depends_on", "supersedes", "migrated_to_spec"} {
		if _, ok := fields[key]; !ok {
			return "", fmt.Errorf("missing required metadata field %s", key)
		}
	}
	lines := []string{
		"- **status**: " + scalarField(fields, "status"),
		"- **date**: " + scalarField(fields, "date"),
		"- **depends_on**: " + strings.Join(listFieldOrEmpty(fields, "depends_on"), ", "),
		"- **supersedes**: " + strings.Join(listFieldOrEmpty(fields, "supersedes"), ", "),
		"- **migrated_to_spec**: " + scalarField(fields, "migrated_to_spec"),
	}
	return strings.Join(lines, "\n"), nil
}

func renderRequirementMetadata(id string, fields map[string]any) (string, error) {
	return renderWorkflowMetadata(id, fields, []string{"id", "status", "date", "source_refs", "work_items"})
}

func renderWorkItemMetadata(id string, fields map[string]any) (string, error) {
	return renderWorkflowMetadata(id, fields, []string{"id", "status", "date", "source_requirement", "impact_refs", "tasks"})
}

func renderTaskMetadata(id string, fields map[string]any) (string, error) {
	return renderWorkflowMetadata(id, fields, []string{"id", "status", "date", "work_item", "source_requirement", "estimate", "depends_on", "outputs"})
}

func renderWorkflowMetadata(id string, fields map[string]any, keys []string) (string, error) {
	out := make([]string, 0, len(keys)*2)
	for _, key := range keys {
		if _, ok := fields[key]; !ok {
			return "", fmt.Errorf("missing required metadata field %s", key)
		}
		if key == "id" {
			out = append(out, "- **id**: "+id)
			continue
		}
		if isListMetadataKey(key) {
			out = append(out, "- **"+key+"**:")
			for _, value := range listFieldOrEmpty(fields, key) {
				out = append(out, "  - "+value)
			}
			continue
		}
		out = append(out, "- **"+key+"**: "+scalarField(fields, key))
	}
	return strings.Join(out, "\n"), nil
}

func isListMetadataKey(key string) bool {
	switch key {
	case "depends_on", "outputs", "source_refs", "work_items", "impact_refs", "tasks":
		return true
	default:
		return false
	}
}

func scalarField(fields map[string]any, key string) string {
	if fields == nil {
		return ""
	}
	value, ok := fields[key]
	if !ok || value == nil {
		return ""
	}
	switch typed := value.(type) {
	case string:
		return typed
	default:
		return fmt.Sprint(typed)
	}
}

func listFieldOrEmpty(fields map[string]any, key string) []string {
	values, ok := listField(fields, key)
	if !ok {
		return []string{}
	}
	return values
}

func listField(fields map[string]any, key string) ([]string, bool) {
	if fields == nil {
		return nil, false
	}
	value, ok := fields[key]
	if !ok || value == nil {
		return nil, false
	}
	switch typed := value.(type) {
	case []string:
		return append([]string{}, typed...), true
	case []any:
		out := make([]string, 0, len(typed))
		for _, item := range typed {
			out = append(out, fmt.Sprint(item))
		}
		return out, true
	default:
		text := strings.TrimSpace(fmt.Sprint(typed))
		if text == "" {
			return []string{}, true
		}
		return splitCommaList(text), true
	}
}

func workflowMetadataScalar(record Record, key string) string {
	if record.WorkflowMeta == nil || record.WorkflowMeta.Fields == nil {
		return ""
	}
	return record.WorkflowMeta.Fields[key].Value
}

func ensureTrailingNewline(s string) string {
	return strings.TrimRight(s, "\n") + "\n"
}
