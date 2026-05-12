package designrecords

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// BuildIndex discovers ADR/spec Markdown records and builds the Phase 1 index.
func BuildIndex(ctx context.Context, cfg Config) (*Index, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	normalized, err := normalizeConfig(cfg)
	if err != nil {
		return nil, err
	}
	idx := &Index{
		Root:        normalized.Root,
		Records:     []Record{},
		Diagnostics: []Diagnostic{},
		Candidates:  []RecordCandidate{},
		ParseIssues: []ParseIssue{},
		PathIssues:  []PathIssue{},
	}
	if err := discoverADRRecords(ctx, normalized.Root, idx); err != nil {
		return nil, err
	}
	if err := discoverSpecRecords(ctx, normalized.Root, idx); err != nil {
		return nil, err
	}
	return idx, nil
}

func discoverADRRecords(ctx context.Context, root string, idx *Index) error {
	pattern := filepath.Join(root, "docs", "adr", "*.md")
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

func discoverSpecRecords(ctx context.Context, root string, idx *Index) error {
	specRoot := filepath.Join(root, "docs", "spec")
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

func relativePath(root, path string) string {
	rel, err := filepath.Rel(root, path)
	if err != nil {
		return filepath.ToSlash(filepath.Clean(path))
	}
	return filepath.ToSlash(rel)
}
