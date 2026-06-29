package designrecords

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// BuildIndex discovers design record and workflow artifact Markdown records
// from explicit current roots only. It discovers sequential records (ADR, investigations,
// requirements, work-items, tasks) and specs, using current-format parsers.
// Duplicate canonical IDs create conflict groups with no arbitrary winner.
// All views are deterministically sorted. No legacy or V01 sources are loaded.
func BuildIndex(ctx context.Context, cfg Config) (*Index, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	normalized, err := normalizeConfig(cfg)
	if err != nil {
		return nil, err
	}
	idx := &Index{
		Root:               normalized.Root,
		NamespacePrefix:    normalized.NamespacePrefix(),
		RecordsRoot:        normalized.primaryRecordsRoot(),
		RecordsEntries:     normalized.RecordsRoots,
		Records:            []Record{},
		Diagnostics:        []Diagnostic{},
		Candidates:         []RecordCandidate{},
		ParseIssues:        []ParseIssue{},
		PathIssues:         []PathIssue{},
		SemanticRefs:       []SemanticRefDecl{},
		SemanticRefSources: []SemanticRefSource{},
	}
	for _, entry := range normalized.RecordsRoots {
		ns := entry.NamespacePrefix
		recordsRootAbs := filepath.Join(normalized.Root, filepath.FromSlash(entry.RecordsRoot))
		if err := validateCurrentRecordsRoot(normalized.Root, recordsRootAbs); err != nil {
			return nil, fmt.Errorf("current records root %q: %w", entry.RecordsRoot, err)
		}
		if err := discoverCurrentADRRecords(ctx, normalized.Root, recordsRootAbs, ns, idx); err != nil {
			return nil, err
		}
		if err := discoverCurrentSpecRecords(ctx, normalized.Root, recordsRootAbs, entry.RecordsRoot, entry.AppNamespace, idx); err != nil {
			return nil, err
		}
		if err := discoverCurrentInvestigationRecords(ctx, normalized.Root, recordsRootAbs, ns, idx); err != nil {
			return nil, err
		}
		if err := discoverCurrentRequirementRecords(ctx, normalized.Root, recordsRootAbs, ns, idx); err != nil {
			return nil, err
		}
		if err := discoverCurrentWorkItemRecords(ctx, normalized.Root, recordsRootAbs, ns, idx); err != nil {
			return nil, err
		}
		if err := discoverCurrentTaskRecords(ctx, normalized.Root, recordsRootAbs, ns, idx); err != nil {
			return nil, err
		}
	}
	separateConflictGroups(idx)
	sortIndexDeterministically(idx)
	return idx, nil
}

func validateCurrentRecordsRoot(root, path string) error {
	rel, err := filepath.Rel(root, path)
	if err != nil {
		return fmt.Errorf("resolve relative path: %w", err)
	}
	current := root
	for _, component := range splitPathComponents(rel) {
		current = filepath.Join(current, component)
		info, err := os.Lstat(current)
		if err != nil {
			return err
		}
		if info.Mode()&os.ModeSymlink != 0 {
			return fmt.Errorf("path component %q must not be a symlink", relativePath(root, current))
		}
	}
	info, err := os.Lstat(path)
	if err != nil {
		return err
	}
	if !info.IsDir() {
		return fmt.Errorf("must be a directory")
	}
	if _, err := os.ReadDir(path); err != nil {
		return fmt.Errorf("read directory: %w", err)
	}
	return nil
}

func splitPathComponents(path string) []string {
	clean := filepath.Clean(path)
	if clean == "." {
		return nil
	}
	components := []string{}
	for clean != "." && clean != string(filepath.Separator) {
		dir, base := filepath.Split(clean)
		if base != "" {
			components = append(components, base)
		}
		clean = filepath.Clean(dir)
	}
	for i, j := 0, len(components)-1; i < j; i, j = i+1, j-1 {
		components[i], components[j] = components[j], components[i]
	}
	return components
}

func currentDirectoryAvailable(path string) (bool, error) {
	info, err := os.Lstat(path)
	if os.IsNotExist(err) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	if info.Mode()&os.ModeSymlink != 0 || !info.IsDir() {
		return false, nil
	}
	return true, nil
}

func separateConflictGroups(idx *Index) {
	byNorm := make(map[string][]Record, len(idx.Records))
	for _, r := range idx.Records {
		if r.NormalizedID != "" {
			byNorm[r.NormalizedID] = append(byNorm[r.NormalizedID], r)
		}
	}
	singles := idx.Records[:0]
	for _, records := range byNorm {
		if len(records) == 1 {
			singles = append(singles, records[0])
			continue
		}
		sort.Slice(records, func(i, j int) bool {
			if records[i].ID != records[j].ID {
				return records[i].ID < records[j].ID
			}
			return records[i].Path < records[j].Path
		})
		sources := make([]string, 0, len(records))
		for _, r := range records {
			sources = append(sources, r.Path)
		}
		sort.Strings(sources)
		idx.ConflictGroups = append(idx.ConflictGroups, CurrentConflict{Ref: records[0].ID, Sources: sources})
	}
	idx.Records = singles
}

func discoverCurrentADRRecords(ctx context.Context, root, recordsRootAbs, ns string, idx *Index) error {
	adrRoot := filepath.Join(recordsRootAbs, "adr")
	available, err := currentDirectoryAvailable(adrRoot)
	if err != nil {
		return fmt.Errorf("discover adr records: %w", err)
	}
	if !available {
		return nil
	}
	matches, err := filepath.Glob(filepath.Join(adrRoot, ns+"ADR-*.md"))
	if err != nil {
		return fmt.Errorf("discover flat adr records: %w", err)
	}
	domains, err := os.ReadDir(adrRoot)
	if err != nil {
		return fmt.Errorf("discover adr records: %w", err)
	}
	for _, domain := range domains {
		if err := ctx.Err(); err != nil {
			return err
		}
		if domain.Type()&os.ModeSymlink != 0 || !domain.IsDir() {
			continue
		}
		domainMatches, err := filepath.Glob(filepath.Join(adrRoot, domain.Name(), ns+"ADR-*-*.md"))
		if err != nil {
			return fmt.Errorf("discover adr records: %w", err)
		}
		matches = append(matches, domainMatches...)
	}
	sort.Strings(matches)
	for _, path := range matches {
		if err := ctx.Err(); err != nil {
			return err
		}
		rel := relativePath(root, path)
		info, err := os.Lstat(path)
		if err != nil {
			idx.PathIssues = append(idx.PathIssues, PathIssue{Path: rel, Operation: "stat", Err: err})
			continue
		}
		if info.Mode()&os.ModeSymlink != 0 || !info.Mode().IsRegular() {
			continue
		}
		content, err := os.ReadFile(path)
		if err != nil {
			idx.PathIssues = append(idx.PathIssues, PathIssue{Path: rel, Operation: "read", Err: err})
			continue
		}
		record, candidate, issues := parseCurrentADRRecord(rel, string(content), ns)
		idx.Candidates = append(idx.Candidates, candidate)
		idx.ParseIssues = append(idx.ParseIssues, issues...)
		if record != nil {
			idx.Records = append(idx.Records, *record)
		}
	}
	return nil
}

func discoverCurrentInvestigationRecords(ctx context.Context, root, recordsRootAbs, ns string, idx *Index) error {
	investigationRoot := filepath.Join(recordsRootAbs, "investigations")
	available, err := currentDirectoryAvailable(investigationRoot)
	if err != nil {
		return fmt.Errorf("discover investigation records: %w", err)
	}
	if !available {
		return nil
	}
	domains, err := os.ReadDir(investigationRoot)
	if err != nil {
		return fmt.Errorf("discover investigation records: %w", err)
	}
	for _, domain := range domains {
		if err := ctx.Err(); err != nil {
			return err
		}
		if domain.Type()&os.ModeSymlink != 0 || !domain.IsDir() {
			continue
		}
		matches, err := filepath.Glob(filepath.Join(investigationRoot, domain.Name(), ns+"INV-*-*.md"))
		if err != nil {
			return fmt.Errorf("discover investigation records: %w", err)
		}
		sort.Strings(matches)
		for _, path := range matches {
			if err := ctx.Err(); err != nil {
				return err
			}
			rel := relativePath(root, path)
			info, err := os.Lstat(path)
			if err != nil {
				idx.PathIssues = append(idx.PathIssues, PathIssue{Path: rel, Operation: "stat", Err: err})
				continue
			}
			if info.Mode()&os.ModeSymlink != 0 || !info.Mode().IsRegular() {
				continue
			}
			content, err := os.ReadFile(path)
			if err != nil {
				idx.PathIssues = append(idx.PathIssues, PathIssue{Path: rel, Operation: "read", Err: err})
				continue
			}
			record, candidate, issues := parseCurrentInvestigationRecord(rel, string(content), ns)
			idx.Candidates = append(idx.Candidates, candidate)
			idx.ParseIssues = append(idx.ParseIssues, issues...)
			if record != nil {
				idx.Records = append(idx.Records, *record)
			}
		}
	}
	return nil
}

func discoverCurrentSpecRecords(ctx context.Context, root, recordsRootAbs, recordsRoot, appNamespace string, idx *Index) error {
	specRoot := filepath.Join(recordsRootAbs, "spec")
	available, err := currentDirectoryAvailable(specRoot)
	if err != nil {
		return fmt.Errorf("discover spec records: %w", err)
	}
	if !available {
		return nil
	}
	err = filepath.WalkDir(specRoot, func(path string, entry os.DirEntry, walkErr error) error {
		if err := ctx.Err(); err != nil {
			return err
		}
		rel := relativePath(root, path)
		if walkErr != nil {
			idx.PathIssues = append(idx.PathIssues, PathIssue{Path: rel, Operation: "walk", Err: walkErr})
			return nil
		}
		if entry.Type()&os.ModeSymlink != 0 {
			return nil
		}
		if entry.IsDir() || strings.ToLower(filepath.Ext(path)) != ".md" {
			return nil
		}
		info, err := os.Stat(path)
		if err != nil {
			idx.PathIssues = append(idx.PathIssues, PathIssue{Path: rel, Operation: "stat", Err: err})
			return nil
		}
		if !info.Mode().IsRegular() {
			return nil
		}
		content, err := os.ReadFile(path)
		if err != nil {
			idx.PathIssues = append(idx.PathIssues, PathIssue{Path: rel, Operation: "read", Err: err})
			return nil
		}
		record, candidate, issues := parseCurrentSpecRecord(rel, string(content), recordsRoot, appNamespace)
		if candidate.Path != "" {
			idx.Candidates = append(idx.Candidates, candidate)
		}
		idx.ParseIssues = append(idx.ParseIssues, issues...)
		if record != nil {
			idx.Records = append(idx.Records, *record)
		}
		return nil
	})
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return fmt.Errorf("discover spec records: %w", err)
	}
	return nil
}

func discoverCurrentRequirementRecords(ctx context.Context, root, recordsRootAbs, ns string, idx *Index) error {
	return discoverCurrentWorkflowRecords(ctx, root, idx,
		filepath.Join(recordsRootAbs, "requirements"),
		ns+"REQ-*-*.md", ns, RecordKindRequirement,
		func(p, r, n string) (*Record, RecordCandidate, []ParseIssue) {
			return parseCurrentRequirementRecord(p, r, n)
		})
}

func discoverCurrentWorkItemRecords(ctx context.Context, root, recordsRootAbs, ns string, idx *Index) error {
	return discoverCurrentWorkflowRecords(ctx, root, idx,
		filepath.Join(recordsRootAbs, "work-items"),
		ns+"WORK-*-*.md", ns, RecordKindWorkItem,
		func(p, r, n string) (*Record, RecordCandidate, []ParseIssue) {
			return parseCurrentWorkItemRecord(p, r, n)
		})
}

func discoverCurrentTaskRecords(ctx context.Context, root, recordsRootAbs, ns string, idx *Index) error {
	return discoverCurrentWorkflowRecords(ctx, root, idx,
		filepath.Join(recordsRootAbs, "tasks"),
		ns+"TASK-*-*.md", ns, RecordKindTask,
		func(p, r, n string) (*Record, RecordCandidate, []ParseIssue) {
			return parseCurrentTaskRecord(p, r, n)
		})
}

func discoverCurrentWorkflowRecords(ctx context.Context, root string, idx *Index, baseRoot, pattern, ns string, kind RecordKind, parser func(string, string, string) (*Record, RecordCandidate, []ParseIssue)) error {
	available, err := currentDirectoryAvailable(baseRoot)
	if err != nil {
		return fmt.Errorf("discover workflow %s records: %w", kind, err)
	}
	if !available {
		return nil
	}
	domains, err := os.ReadDir(baseRoot)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return fmt.Errorf("discover workflow %s records: %w", kind, err)
	}
	for _, domain := range domains {
		if err := ctx.Err(); err != nil {
			return err
		}
		if domain.Type()&os.ModeSymlink != 0 || !domain.IsDir() {
			continue
		}
		matches, err := filepath.Glob(filepath.Join(baseRoot, domain.Name(), pattern))
		if err != nil {
			return fmt.Errorf("discover workflow %s records: %w", kind, err)
		}
		sort.Strings(matches)
		for _, path := range matches {
			if err := ctx.Err(); err != nil {
				return err
			}
			if strings.ToLower(filepath.Ext(path)) != ".md" {
				continue
			}
			rel := relativePath(root, path)
			info, err := os.Lstat(path)
			if err != nil {
				idx.PathIssues = append(idx.PathIssues, PathIssue{Path: rel, Operation: "stat", Err: err})
				continue
			}
			if info.Mode()&os.ModeSymlink != 0 || !info.Mode().IsRegular() {
				continue
			}
			content, err := os.ReadFile(path)
			if err != nil {
				idx.PathIssues = append(idx.PathIssues, PathIssue{Path: rel, Operation: "read", Err: err})
				continue
			}
			record, candidate, issues := parser(rel, string(content), ns)
			idx.Candidates = append(idx.Candidates, candidate)
			idx.ParseIssues = append(idx.ParseIssues, issues...)
			if record != nil {
				idx.Records = append(idx.Records, *record)
			}
		}
	}
	return nil
}

func relativePath(root, path string) string {
	rel, err := filepath.Rel(root, path)
	if err != nil {
		return filepath.ToSlash(filepath.Clean(path))
	}
	return filepath.ToSlash(rel)
}

func sortIndexDeterministically(idx *Index) {
	sort.Slice(idx.Records, func(i, j int) bool {
		if idx.Records[i].ID == idx.Records[j].ID {
			return idx.Records[i].Path < idx.Records[j].Path
		}
		return idx.Records[i].ID < idx.Records[j].ID
	})
	sort.Slice(idx.Candidates, func(i, j int) bool {
		if idx.Candidates[i].ID == idx.Candidates[j].ID {
			return idx.Candidates[i].Path < idx.Candidates[j].Path
		}
		return idx.Candidates[i].ID < idx.Candidates[j].ID
	})
	sort.Slice(idx.ParseIssues, func(i, j int) bool {
		left, right := idx.ParseIssues[i], idx.ParseIssues[j]
		if left.Path != right.Path {
			return left.Path < right.Path
		}
		if left.RecordID != right.RecordID {
			return left.RecordID < right.RecordID
		}
		if left.Category != right.Category {
			return left.Category < right.Category
		}
		if left.Message != right.Message {
			return left.Message < right.Message
		}
		return parseIssueDetailsKey(left.Details) < parseIssueDetailsKey(right.Details)
	})
	sort.Slice(idx.PathIssues, func(i, j int) bool {
		left, right := idx.PathIssues[i], idx.PathIssues[j]
		if left.Path != right.Path {
			return left.Path < right.Path
		}
		if left.Operation != right.Operation {
			return left.Operation < right.Operation
		}
		return errorText(left.Err) < errorText(right.Err)
	})
	for i := range idx.ConflictGroups {
		sort.Strings(idx.ConflictGroups[i].Sources)
	}
	sort.Slice(idx.ConflictGroups, func(i, j int) bool {
		return idx.ConflictGroups[i].Ref < idx.ConflictGroups[j].Ref
	})
	sort.Slice(idx.SemanticRefs, func(i, j int) bool {
		if idx.SemanticRefs[i].Ref == idx.SemanticRefs[j].Ref {
			return idx.SemanticRefs[i].Path < idx.SemanticRefs[j].Path
		}
		return idx.SemanticRefs[i].Ref < idx.SemanticRefs[j].Ref
	})
	sort.Slice(idx.SemanticRefSources, func(i, j int) bool {
		return idx.SemanticRefSources[i].Path < idx.SemanticRefSources[j].Path
	})
}

func parseIssueDetailsKey(details map[string]string) string {
	keys := make([]string, 0, len(details))
	for key := range details {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	var key strings.Builder
	for _, name := range keys {
		key.WriteString(name)
		key.WriteByte(0)
		key.WriteString(details[name])
		key.WriteByte(0)
	}
	return key.String()
}

func errorText(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}
