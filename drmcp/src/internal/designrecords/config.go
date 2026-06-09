package designrecords

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// RecordsEntry describes a single app namespace records tree within a repository.
type RecordsEntry struct {
	// RecordsRoot is the records directory relative to Config.Root (e.g. "v01/records").
	RecordsRoot string
	// NamespacePrefix is derived from the app namespace directory name (e.g. "V01-").
	NamespacePrefix string
}

// Config contains repository-local Design Records MCP configuration.
type Config struct {
	// Root is the absolute repository root used for discovery. Response paths
	// are expected to be relative to this root.
	Root string
	// RecordsRoots is the ordered list of records trees to scan.
	// Built by NewConfig: single entry from --records-root, or auto-detected from */records/.
	RecordsRoots []RecordsEntry
}

// NamespacePrefix returns the namespace prefix of the primary (first) records tree,
// or empty string when RecordsRoots is empty.
func (c Config) NamespacePrefix() string {
	if len(c.RecordsRoots) == 0 {
		return ""
	}
	return c.RecordsRoots[0].NamespacePrefix
}

// primaryRecordsRoot returns the RecordsRoot of the first entry.
// Falls back to "v01/records" when RecordsRoots is empty (should not occur after normalization).
func (c Config) primaryRecordsRoot() string {
	if len(c.RecordsRoots) == 0 {
		return "v01/records"
	}
	return c.RecordsRoots[0].RecordsRoot
}

// NewConfig resolves root from an explicit path or the current working directory when
// root is empty. If recordsRoot is non-empty, a single-entry RecordsRoots is built for
// that path (single-root / backward-compat mode). When recordsRoot is empty, all
// */records/ directories under root are auto-detected (multi-root mode).
func NewConfig(root, recordsRoot string) (Config, error) {
	if root == "" {
		cwd, err := os.Getwd()
		if err != nil {
			return Config{}, fmt.Errorf("resolve cwd: %w", err)
		}
		root = cwd
	}
	abs, err := filepath.Abs(root)
	if err != nil {
		return Config{}, fmt.Errorf("resolve root: %w", err)
	}
	cleanRoot := filepath.Clean(abs)

	var entries []RecordsEntry
	if recordsRoot != "" {
		entries = []RecordsEntry{makeRecordsEntry(recordsRoot)}
	} else {
		entries = discoverRecordsEntries(cleanRoot)
		if len(entries) == 0 {
			// No */records/ dirs found; fall back to default.
			entries = []RecordsEntry{makeRecordsEntry("v01/records")}
		}
	}
	return Config{Root: cleanRoot, RecordsRoots: entries}, nil
}

// makeRecordsEntry builds a RecordsEntry from a records root path.
// NamespacePrefix is derived from the app namespace directory (parent of "records/"):
//
//	"v01/records"  → appNS "v01"  → "V01-"
//	"drmcp/records" → appNS "drmcp" → "DRMCP-"
//	"docs"          → appNS "."    → ""  (single-component path)
func makeRecordsEntry(recordsRoot string) RecordsEntry {
	clean := filepath.Clean(filepath.FromSlash(recordsRoot))
	appNS := filepath.Base(filepath.Dir(clean))
	ns := ""
	if appNS != "." && appNS != "" {
		ns = strings.ToUpper(appNS) + "-"
	}
	return RecordsEntry{
		RecordsRoot:     filepath.ToSlash(clean),
		NamespacePrefix: ns,
	}
}

// discoverRecordsEntries finds all */records/ directories directly under root.
func discoverRecordsEntries(root string) []RecordsEntry {
	matches, err := filepath.Glob(filepath.Join(root, "*", "records"))
	if err != nil {
		return nil
	}
	var entries []RecordsEntry
	for _, abs := range matches {
		info, err := os.Stat(abs)
		if err != nil || !info.IsDir() {
			continue
		}
		rel, err := filepath.Rel(root, abs)
		if err != nil {
			continue
		}
		entries = append(entries, makeRecordsEntry(filepath.ToSlash(rel)))
	}
	return entries
}

func normalizeConfig(cfg Config) (Config, error) {
	if cfg.Root == "" {
		cwd, err := os.Getwd()
		if err != nil {
			return Config{}, fmt.Errorf("resolve cwd: %w", err)
		}
		cfg.Root = cwd
	}
	abs, err := filepath.Abs(cfg.Root)
	if err != nil {
		return Config{}, fmt.Errorf("resolve root: %w", err)
	}
	result := cfg
	result.Root = filepath.Clean(abs)
	if len(result.RecordsRoots) == 0 {
		result.RecordsRoots = discoverRecordsEntries(result.Root)
		if len(result.RecordsRoots) == 0 {
			result.RecordsRoots = []RecordsEntry{makeRecordsEntry("v01/records")}
		}
	}
	return result, nil
}
