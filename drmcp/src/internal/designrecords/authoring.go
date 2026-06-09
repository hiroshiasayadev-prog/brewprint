package designrecords

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pmezard/go-difflib/difflib"
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

	UpdateTypeMetadataBlockReplace  = "metadata_block_replace"
	UpdateTypeMetadataFieldsReplace = "metadata_fields_replace"
	UpdateTypeNamedSectionReplace   = "named_section_replace"
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
	Path        string
	Change      string
	RecordID    string
	RecordKind  RecordKind
	BaseContent string
	BaseHash    string
	BaseID      string
	BaseKind    RecordKind
	Content     string
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
	DiffMode             string         `json:"diff_mode,omitempty"`
}

type ProposeRecordUpdateRequest struct {
	Kind        RecordKind        `json:"kind"`
	ID          string            `json:"id"`
	Update      UpdateRequest     `json:"update"`
	Operations  []UpdateOperation `json:"operations,omitempty"`
	Body        *string           `json:"body,omitempty"`
	BodyCacheID string            `json:"body_cache_id,omitempty"`
	DiffMode    string            `json:"diff_mode,omitempty"`
}

type UpdateRequest struct {
	Type            string           `json:"type"`
	Metadata        map[string]any   `json:"metadata,omitempty"`
	SectionSelector *SectionSelector `json:"section_selector,omitempty"`
}

type UpdateOperation struct {
	Type            string           `json:"type"`
	Metadata        map[string]any   `json:"metadata,omitempty"`
	SectionSelector *SectionSelector `json:"section_selector,omitempty"`
	Body            *string          `json:"body,omitempty"`
	BodyCacheID     string           `json:"body_cache_id,omitempty"`
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

type DiffMode string

const (
	DiffModeSummary DiffMode = "summary"
	DiffModePatch   DiffMode = "patch"
	DiffModeNone    DiffMode = "none"
)

type Diff struct {
	Format  string     `json:"format,omitempty"`
	Files   []DiffFile `json:"files,omitempty"`
	Text    string     `json:"text,omitempty"`
	Omitted bool       `json:"omitted,omitempty"`
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
	diffMode, err := validateAndResolveDiffMode(req.DiffMode)
	if err != nil {
		return failedProposalResponse(nil, nil, authoringDiagnostic(ErrorCodeInvalidRequest, err.Error())), nil
	}
	if req.Body != nil && req.BodyCacheID != "" {
		return failedProposalResponse(nil, nil, authoringDiagnostic(ErrorCodeInvalidBodySource, "body and body_cache_id are mutually exclusive")), nil
	}
	if req.Fields == nil {
		var bodyCache *BodyCacheEntry
		if req.Body != nil {
			bodyCache = store.cacheBody(*req.Body)
		}
		return failedProposalResponse(bodyCache, nil, authoringDiagnostic(ErrorCodeInvalidRequest, "fields is required for create")), nil
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
	if hasErrorDiagnostics(prep.diagnostics) {
		if body != nil {
			bodyCache = store.cacheBody(*body)
		}
		return failedProposalResponse(bodyCache, nil, prep.diagnostics...), nil
	}
	if subdomain, ok := req.Fields["subdomain"].(string); ok && subdomain != "" {
		prep.diagnostics = append(prep.diagnostics, subdomainAdvisoryDiagnostics(idx, prep.target.ResolvedID, subdomain)...)
	}
	return persistProposal(ctx, cfg, idx, store, ProposalOperationCreate, prep, bodyCache, diffMode)
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
	diffMode, err := validateAndResolveDiffMode(req.DiffMode)
	if err != nil {
		return failedProposalResponse(nil, nil, authoringDiagnostic(ErrorCodeInvalidRequest, err.Error())), nil
	}

	hasUpdate := req.Update.Type != ""
	hasOperations := req.Operations != nil

	if hasUpdate && hasOperations {
		return failedProposalResponse(nil, nil, authoringDiagnostic(ErrorCodeInvalidRequest, "update and operations are mutually exclusive")), nil
	}
	if !hasUpdate && !hasOperations {
		return failedProposalResponse(nil, nil, authoringDiagnostic(ErrorCodeInvalidRequest, "exactly one of update or operations must be present")), nil
	}

	if hasOperations {
		if len(req.Operations) == 0 {
			return failedProposalResponse(nil, nil, authoringDiagnostic(ErrorCodeInvalidRequest, "operations must not be empty")), nil
		}
		if req.Body != nil || req.BodyCacheID != "" {
			return failedProposalResponse(nil, nil, authoringDiagnostic(ErrorCodeInvalidBodySource, "body and body_cache_id must not be combined with operations")), nil
		}
		return proposeMultiOpUpdate(ctx, cfg, idx, store, req, diffMode)
	}

	// Single-op path (existing behaviour)
	bodyRequired := req.Update.Type == UpdateTypeNamedSectionReplace
	bodyForbidden := req.Update.Type == UpdateTypeMetadataBlockReplace || req.Update.Type == UpdateTypeMetadataFieldsReplace
	if bodyForbidden && (req.Body != nil || req.BodyCacheID != "") {
		return failedProposalResponse(nil, nil, authoringDiagnostic(ErrorCodeInvalidBodySource, req.Update.Type+" must not include body or body_cache_id")), nil
	}
	body, bodyCache, diagnostics, ok := resolveBodySource(store, req.Body, req.BodyCacheID, bodyRequired)
	if !ok {
		return failedProposalResponse(nil, nil, diagnostics...), nil
	}
	prep, err2 := prepareUpdate(ctx, cfg, idx, req, body)
	if err2 != nil {
		if body != nil {
			bodyCache = store.cacheBody(*body)
		}
		return failedProposalResponse(bodyCache, nil, authoringDiagnostic(ErrorCodeInvalidRequest, err2.Error())), nil
	}
	if hasErrorDiagnostics(prep.diagnostics) {
		if body != nil {
			bodyCache = store.cacheBody(*body)
		}
		return failedProposalResponse(bodyCache, nil, prep.diagnostics...), nil
	}
	if isNoOpUpdate(prep.files) {
		return noOpUpdateResponse(ctx, idx, prep)
	}
	if req.Update.Type == UpdateTypeMetadataFieldsReplace || req.Update.Type == UpdateTypeMetadataBlockReplace {
		if subdomain, ok := req.Update.Metadata["subdomain"].(string); ok && subdomain != "" {
			prep.diagnostics = append(prep.diagnostics, subdomainAdvisoryDiagnostics(idx, prep.target.ResolvedID, subdomain)...)
		}
	}
	return persistProposal(ctx, cfg, idx, store, ProposalOperationUpdate, prep, bodyCache, diffMode)
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
	if len(proposal.RequiredFollowUpUpdates) > 0 {
		if !requiredFollowUpsSatisfied(idx, proposal.RequiredFollowUpUpdates) {
			return acceptRejected(proposal.ProposalID, proposal.State, authoringDiagnostic(ErrorCodeRequiredFollowUpNotSatisfied, "required follow-up updates are not satisfied")), nil
		}
	}
	if diagnostics := acceptTimeDiagnostics(cfg, idx, proposal); len(diagnostics) > 0 {
		return AcceptProposedWriteResponse{
			ProposalID:     proposal.ProposalID,
			State:          proposal.State,
			Written:        false,
			FilesWritten:   []WrittenFile{},
			Validation:     proposal.Validation,
			RepairGuidance: []string{},
			Diagnostics:    diagnostics,
		}, nil
	}
	preWriteValidation, err := validateProposedFiles(ctx, idx, proposal.Files)
	if err != nil {
		return AcceptProposedWriteResponse{}, err
	}
	if hasErrorDiagnostics(preWriteValidation.Diagnostics) {
		return acceptRejected(proposal.ProposalID, proposal.State, authoringDiagnostic(ErrorCodeInvalidRequest, "proposal validation has error diagnostics")), nil
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
	validation, err := validateExistingAffectedFiles(ctx, nextIdx, proposal.Files)
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

func persistProposal(ctx context.Context, cfg Config, idx *Index, store *AuthoringStore, operation string, prep authoringPreparation, bodyCache *BodyCacheEntry, diffMode DiffMode) (ProposeRecordResponse, error) {
	validation, err := validateProposedFiles(ctx, idx, prep.files)
	if err != nil {
		return ProposeRecordResponse{}, err
	}
	now := store.currentTime()
	expiresAt := now.Add(authoringRetentionDays * 24 * time.Hour)
	fullDiff := buildDiff(prep.files)
	proposal := &StoredProposal{
		State:                   ProposalStateProposed,
		Operation:               operation,
		TargetKind:              prep.target.Kind,
		Target:                  prep.target,
		ExpiresAt:               expiresAt,
		RetentionDays:           authoringRetentionDays,
		Diff:                    fullDiff,
		Validation:              validation,
		Diagnostics:             cloneDiagnostics(prep.diagnostics),
		Note:                    noWriteProposalNote,
		Files:                   append([]ProposedFile{}, prep.files...),
		RequiredFollowUpUpdates: append([]RequiredFollowUpUpdate{}, prep.requiredFollowUpUpdates...),
		ProposalCreatedAt:       now,
	}
	store.saveProposal(proposal)
	target := proposal.Target
	responseDiff := shapeDiff(fullDiff, diffMode)
	return ProposeRecordResponse{
		ProposalCreated:         true,
		ProposalID:              proposal.ProposalID,
		State:                   proposal.State,
		Operation:               proposal.Operation,
		TargetKind:              proposal.TargetKind,
		Target:                  &target,
		ExpiresAt:               &expiresAt,
		RetentionDays:           proposal.RetentionDays,
		Diff:                    &responseDiff,
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
	// Detect target namespace from requested ID (supports namespace-prefixed placeholders, e.g. DRMCP-REQ-MCP-new).
	createNS, bareReqID := detectCreateNamespace(idx, req.ID)
	resolvedBare, domain, err := resolveCreateID(idx, req.Kind, bareReqID, req.Domain, req.ParentID, createNS)
	if err != nil {
		return authoringPreparation{}, err
	}
	// Apply namespace prefix so file content and paths use the public ID (e.g. V01-TASK-X-Y-01).
	// resolvedBare is kept without prefix for internal sequence comparisons (exactIDGapWarning).
	resolved := createNS + resolvedBare
	if err := validateCreateFieldsID(req.Kind, req.ID, req.Fields); err != nil {
		return authoringPreparation{}, err
	}
	if findRecordByIDKind(idx, resolved, req.Kind) != nil {
		return authoringPreparation{}, fmt.Errorf("record %s already exists", resolved)
	}

	// Batch required-field validation (REQ-MCP-028): report all missing fields at once.
	if batchDiag := validateCreateFieldsBatch(req.Kind, req.Fields); batchDiag != nil {
		return authoringPreparation{diagnostics: []Diagnostic{*batchDiag}}, nil
	}

	// Status value validation (REQ-MCP-024): emit allowed_values when status is invalid.
	if statusDiag := validateCreateStatusForCreate(req.Kind, req.Fields); statusDiag != nil {
		return authoringPreparation{diagnostics: []Diagnostic{*statusDiag}}, nil
	}

	content := ""
	if body != nil {
		var renderErr error
		content, renderErr = renderCreateBodyWithContent(idx, req.Kind, resolved, req.Title, req.Fields, req.ParentID, req.ID, *body)
		if renderErr != nil {
			return authoringPreparation{}, renderErr
		}
	} else {
		var renderErr error
		content, renderErr = renderCreateBody(idx, req.Kind, resolved, req.Title, req.Fields, req.ParentID)
		if renderErr != nil {
			return authoringPreparation{}, renderErr
		}
	}
	path := createRecordPath(recordsRootForNamespace(cfg, createNS), req.Kind, resolved, req.Title, domain)
	file := ProposedFile{Path: path, Change: "create", RecordID: resolved, RecordKind: req.Kind, BaseHash: "", Content: ensureTrailingNewline(content)}
	prep := authoringPreparation{
		target: AuthoringTarget{RequestedID: req.ID, ResolvedID: resolved, Kind: req.Kind, Domain: domain, ParentID: req.ParentID, Path: path},
		files:  []ProposedFile{file},
	}
	reciprocalFiles, followUps, reciprocalDiags, err := requiredReciprocalUpdates(ctx, cfg, idx, req.Kind, resolved, req.Fields, req.ParentID, mode)
	if err != nil {
		return authoringPreparation{}, err
	}
	prep.files = append(prep.files, reciprocalFiles...)
	prep.requiredFollowUpUpdates = append(prep.requiredFollowUpUpdates, followUps...)
	prep.diagnostics = append(prep.diagnostics, reciprocalDiags...)
	if d := exactIDGapWarning(idx, req.Kind, req.ID, resolvedBare, domain, req.ParentID); d != nil {
		prep.diagnostics = append(prep.diagnostics, *d)
	}
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
	case UpdateTypeMetadataFieldsReplace:
		if req.Update.Metadata == nil {
			return authoringPreparation{}, fmt.Errorf("metadata is required for metadata_fields_replace")
		}
		var base map[string]any
		base, err = currentMetadataAsMap(*record, raw)
		if err == nil {
			updated, diagnostics, err = replaceMetadataBlock(*record, raw, patchMetadataFields(base, req.Update.Metadata))
		}
	case UpdateTypeNamedSectionReplace:
		if req.Update.SectionSelector == nil {
			return authoringPreparation{}, fmt.Errorf("section_selector is required for named_section_replace")
		}
		if body == nil {
			return authoringPreparation{}, fmt.Errorf("body is required for named_section_replace")
		}
		updated, diagnostics, err = replaceNamedSection(raw, *req.Update.SectionSelector, *body, req.Kind)
	default:
		return authoringPreparation{}, fmt.Errorf("unsupported update.type %q", req.Update.Type)
	}
	if err != nil {
		return authoringPreparation{}, err
	}
	if hasErrorDiagnostics(diagnostics) {
		return authoringPreparation{diagnostics: diagnostics}, nil
	}
	file := ProposedFile{
		Path:        record.Path,
		Change:      "modify",
		RecordID:    record.ID,
		RecordKind:  record.Kind,
		BaseContent: raw,
		BaseHash:    contentHash(raw),
		BaseID:      record.ID,
		BaseKind:    record.Kind,
		Content:     ensureTrailingNewline(updated),
	}
	return authoringPreparation{
		target: AuthoringTarget{
			RequestedID: req.ID,
			ResolvedID:  record.ID,
			Kind:        record.Kind,
			Domain:      workflowDomain(bareRecordID(record.ID, namespacePrefixForID(idx, record.ID))),
			Path:        record.Path,
		},
		files:       []ProposedFile{file},
		diagnostics: diagnostics,
	}, nil
}

func proposeMultiOpUpdate(ctx context.Context, cfg Config, idx *Index, store *AuthoringStore, req ProposeRecordUpdateRequest, diffMode DiffMode) (ProposeRecordResponse, error) {
	ops := req.Operations
	bodies := make([]*string, len(ops))
	var firstBodyCache *BodyCacheEntry

	for i, op := range ops {
		bodyRequired := op.Type == UpdateTypeNamedSectionReplace
		bodyForbidden := op.Type == UpdateTypeMetadataBlockReplace || op.Type == UpdateTypeMetadataFieldsReplace
		if bodyForbidden && (op.Body != nil || op.BodyCacheID != "") {
			return failedProposalResponse(nil, nil, authoringDiagnostic(ErrorCodeInvalidBodySource,
				fmt.Sprintf("operations[%d] type %s must not include body or body_cache_id", i, op.Type))), nil
		}
		body, bc, diags, ok := resolveBodySource(store, op.Body, op.BodyCacheID, bodyRequired)
		if !ok {
			return failedProposalResponse(nil, nil, diags...), nil
		}
		bodies[i] = body
		if bc != nil && firstBodyCache == nil {
			firstBodyCache = bc
		}
	}

	prep, err := prepareMultiOpUpdate(ctx, cfg, idx, req, bodies)
	if err != nil {
		return failedProposalResponse(firstBodyCache, nil, authoringDiagnostic(ErrorCodeInvalidRequest, err.Error())), nil
	}
	if hasErrorDiagnostics(prep.diagnostics) {
		return failedProposalResponse(firstBodyCache, nil, prep.diagnostics...), nil
	}
	if isNoOpUpdate(prep.files) {
		return noOpUpdateResponse(ctx, idx, prep)
	}
	return persistProposal(ctx, cfg, idx, store, ProposalOperationUpdate, prep, firstBodyCache, diffMode)
}

func prepareMultiOpUpdate(ctx context.Context, cfg Config, idx *Index, req ProposeRecordUpdateRequest, bodies []*string) (authoringPreparation, error) {
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

	// Conflict detection before any application
	conflictDiags := detectOperationConflicts(req.Operations)
	if hasErrorDiagnostics(conflictDiags) {
		return authoringPreparation{diagnostics: conflictDiags}, nil
	}

	current := raw
	var allDiagnostics []Diagnostic

	// Pass 1: metadata operations (in array order)
	for i, op := range req.Operations {
		_ = i
		switch op.Type {
		case UpdateTypeMetadataBlockReplace:
			var updated string
			var diags []Diagnostic
			updated, diags, err = replaceMetadataBlock(*record, current, op.Metadata)
			if err != nil {
				return authoringPreparation{}, err
			}
			current = updated
			allDiagnostics = append(allDiagnostics, diags...)
		case UpdateTypeMetadataFieldsReplace:
			if op.Metadata == nil {
				return authoringPreparation{}, fmt.Errorf("metadata is required for metadata_fields_replace")
			}
			var base map[string]any
			base, err = currentMetadataAsMap(*record, current)
			if err != nil {
				return authoringPreparation{}, err
			}
			var updated string
			var diags []Diagnostic
			updated, diags, err = replaceMetadataBlock(*record, current, patchMetadataFields(base, op.Metadata))
			if err != nil {
				return authoringPreparation{}, err
			}
			current = updated
			allDiagnostics = append(allDiagnostics, diags...)
		}
	}

	// Pass 2: section operations (in array order)
	for i, op := range req.Operations {
		if op.Type != UpdateTypeNamedSectionReplace {
			continue
		}
		if op.SectionSelector == nil {
			return authoringPreparation{}, fmt.Errorf("section_selector is required for named_section_replace")
		}
		if bodies[i] == nil {
			return authoringPreparation{}, fmt.Errorf("body is required for named_section_replace")
		}
		var updated string
		var diags []Diagnostic
		updated, diags, err = replaceNamedSection(current, *op.SectionSelector, *bodies[i], req.Kind)
		if err != nil {
			return authoringPreparation{}, err
		}
		current = updated
		allDiagnostics = append(allDiagnostics, diags...)
	}

	if hasErrorDiagnostics(allDiagnostics) {
		return authoringPreparation{diagnostics: allDiagnostics}, nil
	}

	file := ProposedFile{
		Path:        record.Path,
		Change:      "modify",
		RecordID:    record.ID,
		RecordKind:  record.Kind,
		BaseContent: raw,
		BaseHash:    contentHash(raw),
		BaseID:      record.ID,
		BaseKind:    record.Kind,
		Content:     ensureTrailingNewline(current),
	}
	return authoringPreparation{
		target: AuthoringTarget{
			RequestedID: req.ID,
			ResolvedID:  record.ID,
			Kind:        record.Kind,
			Domain:      workflowDomain(bareRecordID(record.ID, namespacePrefixForID(idx, record.ID))),
			Path:        record.Path,
		},
		files:       []ProposedFile{file},
		diagnostics: allDiagnostics,
	}, nil
}

func detectOperationConflicts(ops []UpdateOperation) []Diagnostic {
	var diagnostics []Diagnostic

	// MVP constraint: at most one named_section_replace
	sectionCount := 0
	for _, op := range ops {
		if op.Type == UpdateTypeNamedSectionReplace {
			sectionCount++
		}
	}
	if sectionCount > 1 {
		diagnostics = append(diagnostics, authoringDiagnostic(ErrorCodeMultipleSectionReplaceNotSupported,
			"operations contains more than one named_section_replace; at most one is supported per operations array"))
		return diagnostics
	}

	// metadata_block_replace conflicts with any other metadata operation
	metadataBlockCount := 0
	metadataFieldsCount := 0
	for _, op := range ops {
		switch op.Type {
		case UpdateTypeMetadataBlockReplace:
			metadataBlockCount++
		case UpdateTypeMetadataFieldsReplace:
			metadataFieldsCount++
		}
	}
	if metadataBlockCount > 1 {
		diagnostics = append(diagnostics, authoringDiagnostic(ErrorCodeConflictingOperations,
			"operations contains more than one metadata_block_replace"))
	}
	if metadataBlockCount > 0 && metadataFieldsCount > 0 {
		diagnostics = append(diagnostics, authoringDiagnostic(ErrorCodeConflictingOperations,
			"operations combines metadata_block_replace with metadata_fields_replace"))
	}
	if hasErrorDiagnostics(diagnostics) {
		return diagnostics
	}

	// metadata_fields_replace: duplicate field keys across ops
	fieldCount := make(map[string]int)
	for _, op := range ops {
		if op.Type == UpdateTypeMetadataFieldsReplace && op.Metadata != nil {
			for k := range op.Metadata {
				fieldCount[k]++
			}
		}
	}
	for k, count := range fieldCount {
		if count > 1 {
			diagnostics = append(diagnostics, authoringDiagnostic(ErrorCodeConflictingOperations,
				fmt.Sprintf("operations contains conflicting metadata_fields_replace operations targeting field %q", k)))
		}
	}

	return diagnostics
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

func resolveCreateID(idx *Index, kind RecordKind, requestedID, requestedDomain, parentID, ns string) (string, string, error) {
	switch kind {
	case RecordKindDecision:
		if createDecisionPlaceholderPattern.MatchString(requestedID) {
			next, _ := nextDecisionID(idx, ns)
			return next, "", nil
		}
		if _, ok := decisionRecordNumber(requestedID); ok {
			return requestedID, "", nil
		}
		return "", "", fmt.Errorf("invalid decision create ID %q", requestedID)
	case RecordKindRequirement:
		if match := createRequirementPlaceholderPattern.FindStringSubmatch(requestedID); match != nil {
			domain := match[1]
			if !domainMatches(requestedDomain, domain) {
				return "", "", fmt.Errorf("domain %q does not match ID domain %q", requestedDomain, domain)
			}
			return nextWorkflowID(idx, RecordKindRequirement, domain, ns), domain, nil
		}
		if !validWorkflowIDForKind(requestedID, kind) {
			return "", "", fmt.Errorf("invalid requirement create ID %q", requestedID)
		}
		domain := workflowDomain(requestedID)
		if !domainMatches(requestedDomain, domain) {
			return "", "", fmt.Errorf("domain %q does not match ID domain %q", requestedDomain, domain)
		}
		return requestedID, domain, nil
	case RecordKindWorkItem:
		if match := createWorkItemPlaceholderPattern.FindStringSubmatch(requestedID); match != nil {
			domain := match[1]
			if !domainMatches(requestedDomain, domain) {
				return "", "", fmt.Errorf("domain %q does not match ID domain %q", requestedDomain, domain)
			}
			return nextWorkflowID(idx, RecordKindWorkItem, domain, ns), domain, nil
		}
		if !validWorkflowIDForKind(requestedID, kind) {
			return "", "", fmt.Errorf("invalid work item create ID %q", requestedID)
		}
		domain := workflowDomain(requestedID)
		if !domainMatches(requestedDomain, domain) {
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
		parentBare := strings.TrimPrefix(parentID, namespacePrefixForID(idx, parentID))
		parentDomain := workflowDomain(parentBare)
		parentSeq := workflowSequence(parentBare)
		if match := createTaskPlaceholderPattern.FindStringSubmatch(requestedID); match != nil {
			if match[1] != parentDomain || match[2] != parentSeq {
				return "", "", fmt.Errorf("task placeholder ID must match parent work item domain and sequence")
			}
			if !domainMatches(requestedDomain, parentDomain) {
				return "", "", fmt.Errorf("domain %q does not match parent domain %q", requestedDomain, parentDomain)
			}
			return nextTaskID(idx, parentDomain, parentSeq, ns), parentDomain, nil
		}
		if !validWorkflowIDForKind(requestedID, kind) {
			return "", "", fmt.Errorf("invalid task create ID %q", requestedID)
		}
		if workflowDomain(requestedID) != parentDomain || workflowSequence(requestedID) != parentSeq {
			return "", "", fmt.Errorf("task ID must match parent work item domain and sequence")
		}
		if !domainMatches(requestedDomain, parentDomain) {
			return "", "", fmt.Errorf("domain %q does not match parent domain %q", requestedDomain, parentDomain)
		}
		return requestedID, parentDomain, nil
	default:
		return "", "", fmt.Errorf("unsupported kind %q", kind)
	}
}

func domainMatches(requestedDomain, canonicalDomain string) bool {
	return strings.TrimSpace(requestedDomain) == "" || strings.EqualFold(strings.TrimSpace(requestedDomain), canonicalDomain)
}

func validateCreateFieldsID(kind RecordKind, requestedID string, fields map[string]any) error {
	fieldID := scalarField(fields, "id")
	if strings.TrimSpace(fieldID) == "" {
		return nil
	}
	if hasSequenceNewToken(requestedID, kind) {
		return fmt.Errorf("fields.id must be omitted when top-level id uses a new placeholder")
	}
	if normalizeRecordID(fieldID) != normalizeRecordID(requestedID) {
		return fmt.Errorf("fields.id %q does not match top-level id %q", fieldID, requestedID)
	}
	return nil
}

func renderCreateBody(idx *Index, kind RecordKind, id, title string, fields map[string]any, parentID string) (string, error) {
	header, err := renderCreateHeader(idx, kind, id, title, fields, parentID)
	if err != nil {
		return "", err
	}
	sections := defaultCreateSections(kind)
	if sections == "" {
		return header + "\n", nil
	}
	return header + "\n\n" + sections, nil
}

func renderCreateBodyWithContent(idx *Index, kind RecordKind, id, title string, fields map[string]any, parentID, requestedID, body string) (string, error) {
	if err := validateStructuredCreateContentBody(kind, requestedID, id, body); err != nil {
		return "", err
	}
	header, err := renderCreateHeader(idx, kind, id, title, fields, parentID)
	if err != nil {
		return "", err
	}
	return header + "\n\n" + normalizeStructuredCreateContentBody(body), nil
}

func renderCreateHeader(idx *Index, kind RecordKind, id, title string, fields map[string]any, parentID string) (string, error) {
	switch kind {
	case RecordKindDecision:
		meta, err := renderADRMetadata(fields)
		if err != nil {
			return "", err
		}
		return "# " + id + ": " + title + "\n\n" + meta, nil
	case RecordKindRequirement:
		meta, err := renderRequirementCreateMetadata(id, fields)
		if err != nil {
			return "", err
		}
		return "# " + id + ": " + title + "\n\n" + meta, nil
	case RecordKindWorkItem:
		meta, err := renderWorkItemCreateMetadata(id, fields)
		if err != nil {
			return "", err
		}
		return "# " + id + ": " + title + "\n\n" + meta, nil
	case RecordKindTask:
		if scalarField(fields, "work_item") != parentID {
			return "", fmt.Errorf("task create requires explicit fields.work_item equal to parent_id")
		}
		meta, err := renderTaskCreateMetadata(id, fields)
		if err != nil {
			return "", err
		}
		return "# " + id + ": " + title + "\n\n" + meta, nil
	default:
		return "", fmt.Errorf("unsupported kind %q", kind)
	}
}

func defaultCreateSections(kind RecordKind) string {
	switch kind {
	case RecordKindRequirement:
		return "## Requirement\n\n## Evidence\n"
	case RecordKindWorkItem:
		return "## Goal\n\n## Boundary\n\n## Evidence\n"
	case RecordKindTask:
		return "## Goal\n\n## Work\n\n## Done condition\n\n## Verification\n\n## Evidence\n"
	default:
		return ""
	}
}

func validateStructuredCreateContentBody(kind RecordKind, requestedID, resolvedID, body string) error {
	trimmed := strings.TrimSpace(body)
	if trimmed == "" {
		return fmt.Errorf("fields plus body create requires section content in body")
	}
	firstLine := firstNonEmptyLine(body)
	if strings.HasPrefix(firstLine, "# ") {
		return fmt.Errorf("fields plus body create body must omit H1; MCP generates the record heading")
	}
	if firstLine == "---" {
		return fmt.Errorf("fields plus body create body must omit YAML metadata; MCP generates the metadata block")
	}
	if strings.HasPrefix(firstLine, "- **") {
		return fmt.Errorf("fields plus body create body must omit the metadata block; MCP generates metadata from fields")
	}
	if containsMetadataIDLine(body) {
		return fmt.Errorf("fields plus body create body must omit metadata id; MCP uses the resolved target id")
	}
	if hasSequenceNewToken(requestedID, kind) && strings.Contains(body, resolvedID) {
		return fmt.Errorf("fields plus body create body must not include guessed resolved id %s", resolvedID)
	}
	return nil
}

func firstNonEmptyLine(body string) string {
	for _, line := range strings.Split(body, "\n") {
		trimmed := strings.TrimSpace(strings.TrimSuffix(line, "\r"))
		if trimmed != "" {
			return trimmed
		}
	}
	return ""
}

func containsMetadataIDLine(body string) bool {
	for _, line := range strings.Split(body, "\n") {
		if strings.HasPrefix(strings.TrimSpace(strings.TrimSuffix(line, "\r")), "- **id**:") {
			return true
		}
	}
	return false
}

func normalizeStructuredCreateContentBody(body string) string {
	return strings.TrimSpace(strings.ReplaceAll(body, "\r\n", "\n")) + "\n"
}

// requiredCreateFieldNames returns the field names that must be present in
// fields for a create request of the given kind. "id" is excluded because it
// is auto-managed for placeholder IDs.
func requiredCreateFieldNames(kind RecordKind) []string {
	switch kind {
	case RecordKindDecision:
		return []string{"status", "date", "depends_on", "supersedes", "migrated_to_spec"}
	case RecordKindRequirement:
		return []string{"status", "date", "source_refs", "work_items"}
	case RecordKindWorkItem:
		return []string{"status", "date", "source_requirement", "impact_refs", "tasks"}
	case RecordKindTask:
		return []string{"status", "date", "work_item", "source_requirement", "estimate", "depends_on", "outputs"}
	default:
		return nil
	}
}

// allowedStatusValuesForKind returns the status strings allowed for the given
// kind at create time. Returns nil for kinds with no restriction enforced here.
func allowedStatusValuesForKind(kind RecordKind) []string {
	switch kind {
	case RecordKindDecision:
		return []string{
			string(RecordStatusProposed),
			string(RecordStatusAccepted),
			string(RecordStatusSuperseded),
		}
	case RecordKindRequirement:
		return []string{
			string(RecordStatusCaptured),
			string(RecordStatusDecisionNeeded),
			string(RecordStatusAccepted),
			string(RecordStatusDeferred),
			string(RecordStatusRejected),
		}
	case RecordKindWorkItem:
		return []string{
			string(RecordStatusNotStarted),
			string(RecordStatusInProgress),
			string(RecordStatusBlocked),
			string(RecordStatusDone),
		}
	case RecordKindTask:
		return []string{
			string(RecordStatusNotStarted),
			string(RecordStatusInProgress),
			string(RecordStatusBlocked),
			string(RecordStatusDone),
		}
	default:
		return nil
	}
}

// validateCreateFieldsBatch checks that all required fields for the given kind
// are present in fields. Returns a single batch diagnostic listing all missing
// fields, or nil when all required fields are present.
func validateCreateFieldsBatch(kind RecordKind, fields map[string]any) *Diagnostic {
	required := requiredCreateFieldNames(kind)
	var missing []string
	for _, name := range required {
		if _, ok := fields[name]; !ok {
			missing = append(missing, name)
		}
	}
	if len(missing) == 0 {
		return nil
	}
	msg := fmt.Sprintf("fields is missing required metadata fields for kind %s: %s", kind, strings.Join(missing, ", "))
	d := Diagnostic{
		Category:       DiagnosticMissingRequiredMetadataBatch,
		Severity:       DiagnosticSeverityError,
		Message:        msg,
		RequiredFields: missing,
		TargetKind:     string(kind),
	}
	return &d
}

// validateCreateStatusForCreate checks the status field value against the
// allowed set for the kind at create time. Returns an invalid_metadata_value
// diagnostic with allowed_values when the value is invalid, or nil when valid
// or when the status key is absent (absence is caught by batch validation).
func validateCreateStatusForCreate(kind RecordKind, fields map[string]any) *Diagnostic {
	statusVal, ok := fields["status"]
	if !ok {
		return nil // absent; caught by batch validation
	}
	status := RecordStatus(scalarField(fields, "status"))
	if statusAllowedForKind(kind, status) {
		return nil
	}
	allowed := allowedStatusValuesForKind(kind)
	value := string(statusVal.(string))
	d := Diagnostic{
		Category:     DiagnosticInvalidMetadataValue,
		Severity:     DiagnosticSeverityError,
		Field:        "status",
		Value:        value,
		ValuePresent: true,
		Message:      fmt.Sprintf("status %q is not valid for kind %s", value, kind),
		AllowedValues: allowed,
	}
	if len(allowed) > 0 {
		d.RepairSuggestion = map[string]any{"status": allowed[0]}
	}
	return &d
}

func requiredReciprocalUpdates(ctx context.Context, cfg Config, idx *Index, kind RecordKind, id string, fields map[string]any, parentID, mode string) ([]ProposedFile, []RequiredFollowUpUpdate, []Diagnostic, error) {
	switch kind {
	case RecordKindWorkItem:
		reqID := scalarField(fields, "source_requirement")
		if reqID == "" {
			return nil, nil, nil, nil
		}
		parent := findRecordByIDKind(idx, reqID, RecordKindRequirement)
		if parent == nil || parent.Requirement == nil || containsString(parent.Requirement.WorkItems, id) {
			return nil, nil, nil, nil
		}
		followUp := RequiredFollowUpUpdate{RecordID: parent.ID, Kind: RecordKindRequirement, Field: "work_items", Value: id, Message: "add new work item to source requirement work_items"}
		if mode == "report_required_follow_up" {
			diag := reciprocalFollowUpModeRequiredDiagnostic()
			return nil, []RequiredFollowUpUpdate{followUp}, []Diagnostic{diag}, nil
		}
		inclDiag := reciprocalUpdateIncludedDiagnostic(parent.ID, RecordKindRequirement, "work_items", id)
		return []ProposedFile{reciprocalMetadataFile(cfg, *parent, append(parent.Requirement.WorkItems, id))}, nil, []Diagnostic{inclDiag}, nil
	case RecordKindTask:
		parent := findRecordByIDKind(idx, parentID, RecordKindWorkItem)
		if parent == nil || parent.WorkItem == nil || containsString(parent.WorkItem.Tasks, id) {
			return nil, nil, nil, nil
		}
		followUp := RequiredFollowUpUpdate{RecordID: parent.ID, Kind: RecordKindWorkItem, Field: "tasks", Value: id, Message: "add new task to parent work item tasks"}
		if mode == "report_required_follow_up" {
			diag := reciprocalFollowUpModeRequiredDiagnostic()
			return nil, []RequiredFollowUpUpdate{followUp}, []Diagnostic{diag}, nil
		}
		inclDiag := reciprocalUpdateIncludedDiagnostic(parent.ID, RecordKindWorkItem, "tasks", id)
		return []ProposedFile{reciprocalMetadataFile(cfg, *parent, append(parent.WorkItem.Tasks, id))}, nil, []Diagnostic{inclDiag}, nil
	default:
		return nil, nil, nil, nil
	}
}

func reciprocalFollowUpModeRequiredDiagnostic() Diagnostic {
	return Diagnostic{
		Category: DiagnosticReciprocalFollowUpModeRequired,
		Severity: DiagnosticSeverityWarning,
		Message:  `required reciprocal follow-up updates are present; use reciprocal_update_mode: "include_required" for a safe accept`,
		RepairSuggestion: map[string]any{
			"reciprocal_update_mode": "include_required",
		},
	}
}

func reciprocalUpdateIncludedDiagnostic(recordID string, kind RecordKind, field, value string) Diagnostic {
	return Diagnostic{
		Category: DiagnosticReciprocalUpdateIncluded,
		Severity: DiagnosticSeverityInfo,
		RecordID: recordID,
		Field:    field,
		Value:    value,
		Message:  fmt.Sprintf("reciprocal update included: %s.%s will receive %s", recordID, field, value),
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
		Path:        record.Path,
		Change:      "modify",
		RecordID:    record.ID,
		RecordKind:  record.Kind,
		BaseContent: raw,
		BaseHash:    contentHash(raw),
		BaseID:      record.ID,
		BaseKind:    record.Kind,
		Content:     ensureTrailingNewline(updated),
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

func requiredSectionsForKind(kind RecordKind) []string {
	switch kind {
	case RecordKindTask:
		return []string{"Goal", "Work", "Done condition", "Verification", "Evidence"}
	case RecordKindWorkItem:
		return []string{"Goal", "Boundary", "Evidence"}
	case RecordKindRequirement:
		return []string{"Requirement", "Required Outcome"}
	default:
		return nil
	}
}

func replaceNamedSection(raw string, selector SectionSelector, body string, kind ...RecordKind) (string, []Diagnostic, error) {
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
		// Case-only fallback: only for required headings of workflow artifact kinds.
		var recordKind RecordKind
		if len(kind) > 0 {
			recordKind = kind[0]
		}
		requiredSections := requiredSectionsForKind(recordKind)
		isRequired := false
		for _, s := range requiredSections {
			if s == selector.Heading {
				isRequired = true
				break
			}
		}
		if isRequired {
			var caseFallbacks []markdownSection
			for _, section := range sections {
				if section.Heading.Text == selector.Heading {
					continue
				}
				if !strings.EqualFold(section.Heading.Text, selector.Heading) {
					continue
				}
				if selector.Level != nil && section.Heading.Level != *selector.Level {
					continue
				}
				caseFallbacks = append(caseFallbacks, section)
			}
			if len(caseFallbacks) == 1 {
				match := caseFallbacks[0]
				lines := splitMarkdownLines(raw)
				strippedBody, sh, sl, stripped := stripBodyLeadingHeading(body, match.Heading)
				replacementBody := strings.TrimRight(strippedBody, "\n")
				var warnDiags []Diagnostic
				if stripped {
					warnDiags = append(warnDiags, sectionBodyHeadingStrippedDiagnostic(sh, sl))
				}
				replacement := []string{strings.Repeat("#", match.Heading.Level) + " " + selector.Heading}
				if replacementBody != "" {
					replacement = append(replacement, replacementBody)
				}
				out := append([]string{}, lines[:match.StartLine]...)
				out = append(out, replacement...)
				if match.EndLine < len(lines) {
					out = append(out, "")
				}
				out = append(out, lines[match.EndLine:]...)
				return strings.Join(out, "\n"), warnDiags, nil
			}
			if len(caseFallbacks) > 1 {
				return "", []Diagnostic{sectionSelectorDiagnostic(ErrorCodeSectionSelectorAmbiguous, "section selector matched multiple ATX sections by case-insensitive comparison", caseFallbacks)}, nil
			}
		}
		return "", []Diagnostic{sectionSelectorDiagnostic(ErrorCodeSectionSelectorNoMatch, "section selector matched no ATX section", sections)}, nil
	}
	if len(matches) > 1 {
		return "", []Diagnostic{sectionSelectorDiagnostic(ErrorCodeSectionSelectorAmbiguous, "section selector matched multiple ATX sections", matches)}, nil
	}
	match := matches[0]
	lines := splitMarkdownLines(raw)
	strippedBody, sh, sl, stripped := stripBodyLeadingHeading(body, match.Heading)
	replacementBody := strings.TrimRight(strippedBody, "\n")
	var warnDiags []Diagnostic
	if stripped {
		warnDiags = append(warnDiags, sectionBodyHeadingStrippedDiagnostic(sh, sl))
	}
	replacement := []string{strings.Repeat("#", match.Heading.Level) + " " + match.Heading.Text}
	if replacementBody != "" {
		replacement = append(replacement, replacementBody)
	}
	out := append([]string{}, lines[:match.StartLine]...)
	out = append(out, replacement...)
	if match.EndLine < len(lines) {
		out = append(out, "")
	}
	out = append(out, lines[match.EndLine:]...)
	return strings.Join(out, "\n"), warnDiags, nil
}

func stripBodyLeadingHeading(body string, heading Heading) (stripped string, headingText string, headingLevel int, wasStripped bool) {
	splitLines := strings.Split(strings.ReplaceAll(body, "\r\n", "\n"), "\n")
	firstNonEmptyIdx := -1
	for i, line := range splitLines {
		if strings.TrimSpace(line) != "" {
			firstNonEmptyIdx = i
			break
		}
	}
	if firstNonEmptyIdx < 0 {
		return body, "", 0, false
	}
	headingLine := strings.TrimRight(splitLines[firstNonEmptyIdx], "\r")
	match := atxHeadingPattern.FindStringSubmatch(headingLine)
	if match == nil {
		return body, "", 0, false
	}
	level := len(match[1])
	text := strings.TrimSpace(match[2])
	if level != heading.Level || text != heading.Text {
		return body, "", 0, false
	}
	newLines := make([]string, 0, len(splitLines)-1)
	newLines = append(newLines, splitLines[:firstNonEmptyIdx]...)
	newLines = append(newLines, splitLines[firstNonEmptyIdx+1:]...)
	return strings.Join(newLines, "\n"), text, level, true
}

func sectionBodyHeadingStrippedDiagnostic(headingText string, headingLevel int) Diagnostic {
	return Diagnostic{
		Category:        DiagnosticSectionReplacementBodyHeadingStripped,
		Severity:        DiagnosticSeverityWarning,
		Message:         fmt.Sprintf("replacement body leading heading %q (level %d) was stripped: it duplicates the resolved section heading", headingText, headingLevel),
		StrippedHeading: headingText,
		StrippedLevel:   headingLevel,
	}
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
	return validateAffectedRecordSet(ctx, buildHypotheticalIndex(idx, files), files)
}

func validateExistingAffectedFiles(ctx context.Context, idx *Index, files []ProposedFile) (ValidateRecordsResponse, error) {
	return validateAffectedRecordSet(ctx, idx, files)
}

func validateAffectedRecordSet(ctx context.Context, idx *Index, files []ProposedFile) (ValidateRecordsResponse, error) {
	validation, err := ValidateRecords(ctx, idx, ValidateRecordsRequest{})
	if err != nil {
		return ValidateRecordsResponse{}, err
	}
	affectedIDs, affectedPaths := affectedRecordSet(files)
	diagnostics := make([]Diagnostic, 0, len(validation.Diagnostics))
	for _, diagnostic := range validation.Diagnostics {
		if diagnosticInAffectedSet(diagnostic, affectedIDs, affectedPaths) {
			diagnostics = append(diagnostics, diagnostic)
		}
	}
	return ValidateRecordsResponse{OK: !hasErrorDiagnostics(diagnostics), Diagnostics: diagnostics}, nil
}

func buildHypotheticalIndex(idx *Index, files []ProposedFile) *Index {
	hyp := &Index{
		Root:            idx.Root,
		NamespacePrefix: idx.NamespacePrefix,
		RecordsRoot:     idx.RecordsRoot,
		RecordsEntries:  idx.RecordsEntries,
		Records:            []Record{},
		Candidates:         []RecordCandidate{},
		ParseIssues:        []ParseIssue{},
		PathIssues:         []PathIssue{},
		SemanticRefs:       []SemanticRefDecl{},
		SemanticRefSources: []SemanticRefSource{},
	}
	replaced := map[string]bool{}
	for _, file := range files {
		replaced[file.Path] = true
	}
	for _, candidate := range idx.Candidates {
		if !replaced[candidate.Path] {
			hyp.Candidates = append(hyp.Candidates, candidate)
		}
	}
	for _, issue := range idx.ParseIssues {
		if !replaced[issue.Path] {
			hyp.ParseIssues = append(hyp.ParseIssues, issue)
		}
	}
	for _, issue := range idx.PathIssues {
		if !replaced[issue.Path] {
			hyp.PathIssues = append(hyp.PathIssues, issue)
		}
	}
	for _, source := range idx.SemanticRefSources {
		if !replaced[source.Path] {
			hyp.SemanticRefSources = append(hyp.SemanticRefSources, source)
		}
	}
	for _, record := range idx.Records {
		if !replaced[record.Path] {
			hyp.Records = append(hyp.Records, record)
		}
	}
	for _, file := range files {
		fileNS, fileRecordsRoot := entryForPath(hyp, file.Path)
		record, candidate, issues := parseRecordByPath(file.Path, file.Content, fileNS, fileRecordsRoot)
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
	return hyp
}

func affectedRecordSet(files []ProposedFile) (map[string]bool, map[string]bool) {
	ids := map[string]bool{}
	paths := map[string]bool{}
	for _, file := range files {
		if file.RecordID != "" {
			ids[normalizeRecordID(file.RecordID)] = true
		}
		if file.BaseID != "" {
			ids[normalizeRecordID(file.BaseID)] = true
		}
		if file.Path != "" {
			paths[file.Path] = true
		}
	}
	return ids, paths
}

func diagnosticInAffectedSet(diagnostic Diagnostic, affectedIDs, affectedPaths map[string]bool) bool {
	if diagnostic.Path != "" && affectedPaths[diagnostic.Path] {
		return true
	}
	if diagnostic.RecordID != "" && affectedIDs[normalizeRecordID(diagnostic.RecordID)] {
		return true
	}
	return false
}

func parseRecordByPath(path, content, ns, recordsRoot string) (*Record, RecordCandidate, []ParseIssue) {
	root := filepath.ToSlash(recordsRoot)
	switch {
	case strings.HasPrefix(path, root+"/adr/"):
		return parseADRRecord(path, content, ns)
	case strings.HasPrefix(path, root+"/spec/"):
		return parseSpecRecord(path, content)
	case strings.HasPrefix(path, root+"/requirements/"):
		return parseRequirementRecord(path, content, ns)
	case strings.HasPrefix(path, root+"/work-items/"):
		return parseWorkItemRecord(path, content, ns)
	case strings.HasPrefix(path, root+"/tasks/"):
		return parseTaskRecord(path, content, ns)
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

func validateAndResolveDiffMode(raw string) (DiffMode, error) {
	switch DiffMode(raw) {
	case "", DiffModeSummary:
		return DiffModeSummary, nil
	case DiffModePatch:
		return DiffModePatch, nil
	case DiffModeNone:
		return DiffModeNone, nil
	default:
		return "", fmt.Errorf("invalid diff_mode %q: must be %q, %q, or %q", raw, DiffModeSummary, DiffModePatch, DiffModeNone)
	}
}

func shapeDiff(full Diff, mode DiffMode) Diff {
	switch mode {
	case DiffModePatch:
		return full
	case DiffModeNone:
		return Diff{Omitted: true}
	default: // DiffModeSummary
		return Diff{Format: full.Format, Files: full.Files}
	}
}

func buildDiff(files []ProposedFile) Diff {
	diffFiles := make([]DiffFile, 0, len(files))
	var text strings.Builder
	for _, file := range files {
		diffFiles = append(diffFiles, DiffFile{Path: file.Path, Change: file.Change, RecordID: file.RecordID, RecordKind: file.RecordKind})
		text.WriteString(buildFileDiff(file))
	}
	return Diff{Format: "unified", Files: diffFiles, Text: text.String()}
}

func buildFileDiff(file ProposedFile) string {
	oldContent := file.BaseContent
	oldName := "a/" + file.Path
	newName := "b/" + file.Path
	oldHeader := oldName
	if file.Change == "create" {
		oldContent = ""
		oldHeader = "/dev/null"
	}
	var text strings.Builder
	text.WriteString("diff --git a/" + file.Path + " b/" + file.Path + "\n")
	text.WriteString("index " + shortContentHash(oldContent) + ".." + shortContentHash(file.Content) + " 100644\n")
	text.WriteString("--- " + oldHeader + "\n")
	text.WriteString("+++ " + newName + "\n")
	hunks, err := difflib.GetUnifiedDiffString(difflib.UnifiedDiff{
		A:        difflib.SplitLines(oldContent),
		B:        difflib.SplitLines(file.Content),
		FromFile: oldHeader,
		ToFile:   newName,
		Context:  3,
	})
	if err != nil {
		return text.String()
	}
	text.WriteString(stripUnifiedDiffFileHeaders(hunks))
	return text.String()
}

func stripUnifiedDiffFileHeaders(diff string) string {
	lines := strings.Split(diff, "\n")
	if len(lines) >= 2 && strings.HasPrefix(lines[0], "--- ") && strings.HasPrefix(lines[1], "+++ ") {
		lines = lines[2:]
	}
	return strings.Join(lines, "\n")
}

func noOpUpdateResponse(ctx context.Context, idx *Index, prep authoringPreparation) (ProposeRecordResponse, error) {
	validation, err := validateProposedFiles(ctx, idx, prep.files)
	if err != nil {
		return ProposeRecordResponse{}, err
	}
	target := prep.target
	diagnostics := cloneDiagnostics(prep.diagnostics)
	diagnostics = append(diagnostics, Diagnostic{
		Category: DiagnosticNoOpUpdate,
		Severity: DiagnosticSeverityInfo,
		RecordID: target.ResolvedID,
		Path:     target.Path,
		Message:  "update produced no persisted content changes",
	})
	return ProposeRecordResponse{
		ProposalCreated: false,
		Operation:       ProposalOperationUpdate,
		TargetKind:      target.Kind,
		Target:          &target,
		Validation:      validation,
		Diagnostics:     diagnostics,
	}, nil
}

func isNoOpUpdate(files []ProposedFile) bool {
	if len(files) == 0 {
		return false
	}
	for _, file := range files {
		if file.Change != "modify" || file.BaseContent != file.Content {
			return false
		}
	}
	return true
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

func exactIDGapWarning(idx *Index, kind RecordKind, requestedID, resolvedID, domain, parentID string) *Diagnostic {
	if hasSequenceNewToken(requestedID, kind) {
		return nil
	}
	switch kind {
	case RecordKindRequirement, RecordKindWorkItem:
		requestedSeq, err := strconv.Atoi(workflowSequence(resolvedID))
		if err != nil {
			return nil
		}
		maxSeq := maxWorkflowSeq(idx, kind, domain)
		if requestedSeq <= maxSeq+1 {
			return nil
		}
		prefix := workflowKindPrefix(kind)
		nextAvailable := fmt.Sprintf("%s-%s-%03d", prefix, domain, maxSeq+1)
		d := Diagnostic{
			Category: DiagnosticExactIDSequenceGap,
			Severity: DiagnosticSeverityInfo,
			Message:  fmt.Sprintf("%s skips the next available sequence %s; prefer %s-%s-new unless this ID is intentional", resolvedID, nextAvailable, prefix, domain),
		}
		return &d
	case RecordKindTask:
		parts := strings.Split(resolvedID, "-")
		if len(parts) != 4 {
			return nil
		}
		requestedTaskSeq, err := strconv.Atoi(parts[3])
		if err != nil {
			return nil
		}
		parentBareParts := bareRecordID(parentID, namespacePrefixForID(idx, parentID))
		parentDomain := workflowDomain(parentBareParts)
		parentWorkSeq := workflowSequence(parentBareParts)
		maxSeq := maxTaskSeqForParent(idx, parentDomain, parentWorkSeq)
		if requestedTaskSeq <= maxSeq+1 {
			return nil
		}
		nextAvailable := fmt.Sprintf("TASK-%s-%s-%02d", parentDomain, parentWorkSeq, maxSeq+1)
		d := Diagnostic{
			Category: DiagnosticExactIDSequenceGap,
			Severity: DiagnosticSeverityInfo,
			Message:  fmt.Sprintf("%s skips the next available task sequence %s; prefer TASK-%s-%s-new unless this ID is intentional", resolvedID, nextAvailable, parentDomain, parentWorkSeq),
		}
		return &d
	default:
		return nil
	}
}

func maxWorkflowSeq(idx *Index, kind RecordKind, domain string) int {
	maxNum := 0
	for _, record := range idx.Records {
		bare := bareRecordID(record.ID, namespacePrefixForID(idx, record.ID))
		if record.Kind != kind || workflowDomain(bare) != domain {
			continue
		}
		seq, err := strconv.Atoi(workflowSequence(bare))
		if err == nil && seq > maxNum {
			maxNum = seq
		}
	}
	return maxNum
}

func maxTaskSeqForParent(idx *Index, domain, workSeq string) int {
	maxNum := 0
	for _, record := range idx.Records {
		bare := bareRecordID(record.ID, namespacePrefixForID(idx, record.ID))
		if record.Kind != RecordKindTask || workflowDomain(bare) != domain || workflowSequence(bare) != workSeq {
			continue
		}
		parts := strings.Split(bare, "-")
		if len(parts) == 4 {
			seq, err := strconv.Atoi(parts[3])
			if err == nil && seq > maxNum {
				maxNum = seq
			}
		}
	}
	return maxNum
}

func workflowKindPrefix(kind RecordKind) string {
	if kind == RecordKindWorkItem {
		return "WORK"
	}
	return "REQ"
}

// bareRecordID strips the namespace prefix from an ID before passing it to
// positional parsers (workflowDomain, workflowSequence, decisionRecordNumber)
// that expect bare IDs without namespace prefix. Safe to call when nsPrefix is
// empty or when id does not start with nsPrefix.
func bareRecordID(id, nsPrefix string) string {
	return strings.TrimPrefix(id, nsPrefix)
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

func shortContentHash(raw string) string {
	hash := contentHash(raw)
	if len(hash) <= 12 {
		return hash
	}
	return hash[:12]
}

func readRepoFile(cfg Config, rel string) (string, error) {
	data, err := os.ReadFile(filepath.Join(cfg.Root, filepath.FromSlash(rel)))
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func createRecordPath(root string, kind RecordKind, id, title, domain string) string {
	slug := slugifyRecordTitle(title)
	suffix := ".md"
	if slug != "" {
		suffix = "-" + slug + ".md"
	}
	switch kind {
	case RecordKindDecision:
		return root + "/adr/" + id + suffix
	case RecordKindRequirement:
		return root + "/requirements/" + strings.ToLower(domain) + "/" + id + suffix
	case RecordKindWorkItem:
		return root + "/work-items/" + strings.ToLower(domain) + "/" + id + suffix
	case RecordKindTask:
		return root + "/tasks/" + strings.ToLower(domain) + "/" + id + suffix
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

func nextDecisionID(idx *Index, ns string) (string, int) {
	maxNum := 0
	for _, record := range idx.Records {
		if record.Kind != RecordKindDecision {
			continue
		}
		if namespacePrefixForID(idx, record.ID) != ns {
			continue
		}
		num, ok := decisionRecordNumber(bareRecordID(record.ID, ns))
		if ok && num > maxNum {
			maxNum = num
		}
	}
	return fmt.Sprintf("ADR-%03d", maxNum+1), maxNum + 1
}

func nextWorkflowID(idx *Index, kind RecordKind, domain, ns string) string {
	maxNum := 0
	for _, record := range idx.Records {
		if namespacePrefixForID(idx, record.ID) != ns {
			continue
		}
		bare := bareRecordID(record.ID, ns)
		if record.Kind != kind || workflowDomain(bare) != domain {
			continue
		}
		seq, err := strconv.Atoi(workflowSequence(bare))
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

func nextTaskID(idx *Index, domain, workSeq, ns string) string {
	maxNum := 0
	for _, record := range idx.Records {
		if namespacePrefixForID(idx, record.ID) != ns {
			continue
		}
		bare := bareRecordID(record.ID, ns)
		if record.Kind != RecordKindTask || workflowDomain(bare) != domain || workflowSequence(bare) != workSeq {
			continue
		}
		parts := strings.Split(bare, "-")
		if len(parts) == 4 {
			seq, err := strconv.Atoi(parts[3])
			if err == nil && seq > maxNum {
				maxNum = seq
			}
		}
	}
	return fmt.Sprintf("TASK-%s-%s-%02d", domain, workSeq, maxNum+1)
}

func collectSubdomainValues(idx *Index, domain string) []string {
	seen := map[string]bool{}
	var values []string
	for _, r := range idx.Records {
		if workflowDomain(bareRecordID(r.ID, namespacePrefixForID(idx, r.ID))) != domain {
			continue
		}
		var sub *string
		switch {
		case r.Requirement != nil:
			sub = r.Requirement.Subdomain
		case r.WorkItem != nil:
			sub = r.WorkItem.Subdomain
		case r.Task != nil:
			sub = r.Task.Subdomain
		}
		if sub != nil && *sub != "" && !seen[*sub] {
			seen[*sub] = true
			values = append(values, *sub)
		}
	}
	sort.Strings(values)
	return values
}

func subdomainAdvisoryDiagnostics(idx *Index, id, subdomain string) []Diagnostic {
	domain := workflowDomain(bareRecordID(id, namespacePrefixForID(idx, id)))
	if domain == "" || subdomain == "" {
		return nil
	}
	existing := collectSubdomainValues(idx, domain)
	for _, v := range existing {
		if v == subdomain {
			return nil
		}
	}
	msg := fmt.Sprintf("subdomain value %q is new in domain %s", subdomain, domain)
	if len(existing) > 0 {
		quoted := make([]string, len(existing))
		for i, v := range existing {
			quoted[i] = fmt.Sprintf("%q", v)
		}
		msg += fmt.Sprintf(". Existing values: %s", strings.Join(quoted, ", "))
	} else {
		msg += ". No existing subdomain values in this domain."
	}
	return []Diagnostic{{
		Category:      DiagnosticNewSubdomainValue,
		Severity:      DiagnosticSeverityInfo,
		Field:         "subdomain",
		Value:         subdomain,
		ValuePresent:  true,
		Message:       msg,
		AllowedValues: existing,
	}}
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
	return renderWorkflowMetadata(id, fields, []string{"id", "status", "date", "source_refs", "work_items"}, true)
}

func renderWorkItemMetadata(id string, fields map[string]any) (string, error) {
	return renderWorkflowMetadata(id, fields, []string{"id", "status", "date", "source_requirement", "impact_refs", "tasks"}, true)
}

func renderTaskMetadata(id string, fields map[string]any) (string, error) {
	return renderWorkflowMetadata(id, fields, []string{"id", "status", "date", "work_item", "source_requirement", "estimate", "depends_on", "outputs"}, true)
}

func renderRequirementCreateMetadata(id string, fields map[string]any) (string, error) {
	return renderWorkflowMetadata(id, fields, []string{"id", "status", "date", "source_refs", "work_items"}, false)
}

func renderWorkItemCreateMetadata(id string, fields map[string]any) (string, error) {
	return renderWorkflowMetadata(id, fields, []string{"id", "status", "date", "source_requirement", "impact_refs", "tasks"}, false)
}

func renderTaskCreateMetadata(id string, fields map[string]any) (string, error) {
	return renderWorkflowMetadata(id, fields, []string{"id", "status", "date", "work_item", "source_requirement", "estimate", "depends_on", "outputs"}, false)
}

func renderWorkflowMetadata(id string, fields map[string]any, keys []string, requireFieldsID bool) (string, error) {
	out := make([]string, 0, len(keys)*2)
	for _, key := range keys {
		if key == "id" {
			if requireFieldsID {
				if _, ok := fields[key]; !ok {
					return "", fmt.Errorf("missing required metadata field %s", key)
				}
			}
			out = append(out, "- **id**: "+id)
			continue
		}
		if _, ok := fields[key]; !ok {
			return "", fmt.Errorf("missing required metadata field %s", key)
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

// patchMetadataFields merges patch fields on top of base, returning a new map.
// Fields present in patch overwrite or extend fields from base.
// Fields absent from patch are kept from base unchanged.
func patchMetadataFields(base, patch map[string]any) map[string]any {
	merged := make(map[string]any, len(base))
	for k, v := range base {
		merged[k] = v
	}
	for k, v := range patch {
		merged[k] = v
	}
	return merged
}

// currentMetadataAsMap builds a metadata map from an existing record's parsed data.
// The returned map is suitable for passing to replaceMetadataBlock after field patching.
func currentMetadataAsMap(record Record, raw string) (map[string]any, error) {
	switch record.Kind {
	case RecordKindSpec:
		return currentSpecMetadataAsMap(raw)
	case RecordKindDecision:
		m := map[string]any{
			"status": string(record.Status),
			"date":   rawMetadataScalarValue(raw, "date"),
		}
		if record.Decision != nil {
			m["depends_on"] = record.Decision.DependsOn
			m["supersedes"] = record.Decision.Supersedes
			migratedToSpec := ""
			if record.Decision.MigratedToSpec != nil {
				migratedToSpec = *record.Decision.MigratedToSpec
			}
			m["migrated_to_spec"] = migratedToSpec
		} else {
			m["depends_on"] = []string{}
			m["supersedes"] = []string{}
			m["migrated_to_spec"] = ""
		}
		return m, nil
	case RecordKindTask:
		m := map[string]any{
			"id":     record.ID,
			"status": string(record.Status),
			"date":   workflowMetadataScalar(record, "date"),
		}
		if record.Task != nil {
			m["work_item"] = record.Task.WorkItem
			m["source_requirement"] = record.Task.SourceRequirement
			m["estimate"] = record.Task.Estimate
			m["depends_on"] = record.Task.DependsOn
			m["outputs"] = record.Task.Outputs
		}
		return m, nil
	case RecordKindWorkItem:
		m := map[string]any{
			"id":     record.ID,
			"status": string(record.Status),
			"date":   workflowMetadataScalar(record, "date"),
		}
		if record.WorkItem != nil {
			m["source_requirement"] = record.WorkItem.SourceRequirement
			m["impact_refs"] = record.WorkItem.ImpactRefs
			m["tasks"] = record.WorkItem.Tasks
		}
		return m, nil
	case RecordKindRequirement:
		m := map[string]any{
			"id":     record.ID,
			"status": string(record.Status),
			"date":   workflowMetadataScalar(record, "date"),
		}
		if record.Requirement != nil {
			m["source_refs"] = record.Requirement.SourceRefs
			m["work_items"] = record.Requirement.WorkItems
		}
		return m, nil
	default:
		return nil, fmt.Errorf("metadata_fields_replace does not support kind %q", record.Kind)
	}
}

// currentSpecMetadataAsMap extracts recognized spec metadata fields from YAML front matter.
// The returned map is suitable for passing to replaceMetadataBlock after field patching.
func currentSpecMetadataAsMap(raw string) (map[string]any, error) {
	fm, _, ok := extractFrontMatter(raw)
	if !ok {
		return nil, fmt.Errorf("spec front matter is required")
	}
	var node map[string]any
	if err := yaml.Unmarshal([]byte(fm), &node); err != nil {
		return nil, err
	}
	if node == nil {
		node = map[string]any{}
	}
	result := map[string]any{}
	for _, key := range []string{"scope", "status"} {
		if v, ok2 := node[key]; ok2 {
			result[key] = v
		}
	}
	if v, ok2 := node["design_record"]; ok2 {
		result["design_record"] = v
	}
	return result, nil
}

// rawMetadataScalarValue reads a single scalar field from the metadata block in raw Markdown.
// Used for fields not stored in the parsed Record struct (e.g., ADR date).
func rawMetadataScalarValue(raw, key string) string {
	lines := splitMarkdownLines(raw)
	for _, line := range metadataBlock(lines) {
		k, v, ok := parseMetadataLine(line)
		if ok && k == key {
			return v
		}
	}
	return ""
}

// namespacePrefixForID returns the namespace prefix for the given record ID
// by finding the first entry in idx.RecordsEntries whose prefix is a prefix of the ID.
// Falls back to idx.NamespacePrefix when no entry matches.
func namespacePrefixForID(idx *Index, id string) string {
	for _, e := range idx.RecordsEntries {
		if e.NamespacePrefix != "" && strings.HasPrefix(id, e.NamespacePrefix) {
			return e.NamespacePrefix
		}
	}
	return idx.NamespacePrefix
}

// detectCreateNamespace determines the target namespace for a create request by
// stripping any matching namespace prefix from id. Returns the namespace prefix
// and the bare ID (without namespace). Falls back to the primary namespace.
func detectCreateNamespace(idx *Index, id string) (ns, bareID string) {
	for _, e := range idx.RecordsEntries {
		if e.NamespacePrefix != "" && strings.HasPrefix(id, e.NamespacePrefix) {
			return e.NamespacePrefix, strings.TrimPrefix(id, e.NamespacePrefix)
		}
	}
	return idx.NamespacePrefix, id
}

// recordsRootForNamespace returns the records root path (slash-separated) for the
// given namespace prefix. Falls back to the primary records root.
func recordsRootForNamespace(cfg Config, ns string) string {
	for _, e := range cfg.RecordsRoots {
		if e.NamespacePrefix == ns {
			return filepath.ToSlash(e.RecordsRoot)
		}
	}
	return filepath.ToSlash(cfg.primaryRecordsRoot())
}

// entryForPath returns the namespace prefix and records root for a proposed file
// by matching the file path against each entry's records root.
// Falls back to the primary namespace and records root.
func entryForPath(idx *Index, path string) (ns, recordsRoot string) {
	slashPath := filepath.ToSlash(path)
	for _, e := range idx.RecordsEntries {
		prefix := filepath.ToSlash(e.RecordsRoot) + "/"
		if strings.HasPrefix(slashPath, prefix) {
			return e.NamespacePrefix, e.RecordsRoot
		}
	}
	return idx.NamespacePrefix, idx.RecordsRoot
}
