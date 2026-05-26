package designrecords

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// BuildIndex discovers design record and workflow artifact Markdown records.
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
		Records:            []Record{},
		Diagnostics:        []Diagnostic{},
		Candidates:         []RecordCandidate{},
		ParseIssues:        []ParseIssue{},
		PathIssues:         []PathIssue{},
		SemanticRefs:       []SemanticRefDecl{},
		SemanticRefSources: []SemanticRefSource{},
	}
	if err := discoverADRRecords(ctx, normalized.Root, idx); err != nil {
		return nil, err
	}
	if err := discoverSpecRecords(ctx, normalized.Root, idx); err != nil {
		return nil, err
	}
	if err := discoverInvestigationRecords(ctx, normalized.Root, idx); err != nil {
		return nil, err
	}
	if err := discoverRequirementRecords(ctx, normalized.Root, idx); err != nil {
		return nil, err
	}
	if err := discoverWorkItemRecords(ctx, normalized.Root, idx); err != nil {
		return nil, err
	}
	if err := discoverTaskRecords(ctx, normalized.Root, idx); err != nil {
		return nil, err
	}
	return idx, nil
}

func discoverADRRecords(ctx context.Context, root string, idx *Index) error {
	adrRoot := filepath.Join(root, "docs", "adr")
	if _, err := os.Stat(adrRoot); os.IsNotExist(err) {
		return nil
	}
	pattern := filepath.Join(adrRoot, "*.md")
	matches, err := filepath.Glob(pattern)
	if err != nil {
		return fmt.Errorf("discover adr records: %w", err)
	}
	sort.Strings(matches)
	for _, path := range matches {
		if err := ctx.Err(); err != nil {
			return err
		}
		rel := relativePath(root, path)
		if _, err := os.Stat(path); err != nil {
			idx.PathIssues = append(idx.PathIssues, PathIssue{Path: rel, Operation: "stat", Err: err})
			continue
		}
		content, err := os.ReadFile(path)
		if err != nil {
			idx.PathIssues = append(idx.PathIssues, PathIssue{Path: rel, Operation: "read", Err: err})
			continue
		}
		record, candidate, issues := parseADRRecord(rel, string(content))
		idx.Candidates = append(idx.Candidates, candidate)
		idx.ParseIssues = append(idx.ParseIssues, issues...)
		if record != nil {
			idx.Records = append(idx.Records, *record)
		}
	}
	return nil
}

func discoverInvestigationRecords(ctx context.Context, root string, idx *Index) error {
	investigationRoot := filepath.Join(root, "docs", "investigations")
	if _, err := os.Stat(investigationRoot); os.IsNotExist(err) {
		return nil
	}
	err := filepath.WalkDir(investigationRoot, func(path string, entry os.DirEntry, walkErr error) error {
		if err := ctx.Err(); err != nil {
			return err
		}
		rel := relativePath(root, path)
		if walkErr != nil {
			idx.PathIssues = append(idx.PathIssues, PathIssue{Path: rel, Operation: "walk", Err: walkErr})
			return nil
		}
		if entry.IsDir() || strings.ToLower(filepath.Ext(path)) != ".md" {
			return nil
		}
		if investigationFilenameID(path) == "" {
			return nil
		}
		if _, err := os.Stat(path); err != nil {
			idx.PathIssues = append(idx.PathIssues, PathIssue{Path: rel, Operation: "stat", Err: err})
			return nil
		}
		content, err := os.ReadFile(path)
		if err != nil {
			idx.PathIssues = append(idx.PathIssues, PathIssue{Path: rel, Operation: "read", Err: err})
			return nil
		}
		record, candidate, issues := parseInvestigationRecord(rel, string(content))
		idx.Candidates = append(idx.Candidates, candidate)
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
		return fmt.Errorf("discover investigation records: %w", err)
	}
	return nil
}

func discoverSpecRecords(ctx context.Context, root string, idx *Index) error {
	specRoot := filepath.Join(root, "docs", "spec")
	if _, err := os.Stat(specRoot); os.IsNotExist(err) {
		return nil
	}
	err := filepath.WalkDir(specRoot, func(path string, entry os.DirEntry, walkErr error) error {
		if err := ctx.Err(); err != nil {
			return err
		}
		rel := relativePath(root, path)
		if walkErr != nil {
			idx.PathIssues = append(idx.PathIssues, PathIssue{Path: rel, Operation: "walk", Err: walkErr})
			return nil
		}
		if entry.IsDir() || strings.ToLower(filepath.Ext(path)) != ".md" {
			return nil
		}
		if _, err := os.Stat(path); err != nil {
			idx.PathIssues = append(idx.PathIssues, PathIssue{Path: rel, Operation: "stat", Err: err})
			return nil
		}
		content, err := os.ReadFile(path)
		if err != nil {
			idx.PathIssues = append(idx.PathIssues, PathIssue{Path: rel, Operation: "read", Err: err})
			return nil
		}
		record, candidate, issues := parseSpecRecord(rel, string(content))
		if source, ok := parseSpecSemanticRefSource(rel, string(content)); ok {
			idx.SemanticRefSources = append(idx.SemanticRefSources, source)
			idx.SemanticRefs = append(idx.SemanticRefs, source.Decls...)
		}
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

func discoverRequirementRecords(ctx context.Context, root string, idx *Index) error {
	return discoverWorkflowRecords(ctx, root, idx, filepath.Join("docs", "requirements"), "REQ-*.md", RecordKindRequirement, parseRequirementRecord)
}

func discoverWorkItemRecords(ctx context.Context, root string, idx *Index) error {
	return discoverWorkflowRecords(ctx, root, idx, filepath.Join("docs", "work-items"), "WORK-*.md", RecordKindWorkItem, parseWorkItemRecord)
}

func discoverTaskRecords(ctx context.Context, root string, idx *Index) error {
	return discoverWorkflowRecords(ctx, root, idx, filepath.Join("docs", "tasks"), "TASK-*.md", RecordKindTask, parseTaskRecord)
}

func discoverWorkflowRecords(ctx context.Context, root string, idx *Index, rootRel, pattern string, kind RecordKind, parser func(string, string) (*Record, RecordCandidate, []ParseIssue)) error {
	baseRoot := filepath.Join(root, rootRel)
	if _, err := os.Stat(baseRoot); os.IsNotExist(err) {
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
		if !domain.IsDir() {
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
			if workflowFilenameID(path, kind) == "" {
				continue
			}
			rel := relativePath(root, path)
			if _, err := os.Stat(path); err != nil {
				idx.PathIssues = append(idx.PathIssues, PathIssue{Path: rel, Operation: "stat", Err: err})
				continue
			}
			content, err := os.ReadFile(path)
			if err != nil {
				idx.PathIssues = append(idx.PathIssues, PathIssue{Path: rel, Operation: "read", Err: err})
				continue
			}
			record, candidate, issues := parser(rel, string(content))
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
