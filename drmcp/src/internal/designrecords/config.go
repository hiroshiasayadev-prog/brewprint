package designrecords

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// CurrentRoot is one explicitly declared current records root for NormalizeConfig.
type CurrentRoot struct {
	AppNamespace string // e.g. "product", "drmcp"
	RecordsRoot  string // repository-relative, e.g. "product/records"
}

// RecordsEntry describes one validated current records root within a repository.
type RecordsEntry struct {
	AppNamespace    string // explicit app namespace, e.g. "product"
	RecordsRoot     string // repository-relative records root, e.g. "product/records"
	NamespacePrefix string // derived: strings.ToUpper(AppNamespace) + "-", e.g. "PRODUCT-"
}

// Config contains repository-local Design Records MCP configuration.
type Config struct {
	Root         string         // absolute repository root
	RecordsRoots []RecordsEntry // validated current records roots
}

// NamespacePrefix returns the namespace prefix of the first records root, or "".
func (c Config) NamespacePrefix() string {
	if len(c.RecordsRoots) == 0 {
		return ""
	}
	return c.RecordsRoots[0].NamespacePrefix
}

// primaryRecordsRoot returns the RecordsRoot of the first entry, or "".
func (c Config) primaryRecordsRoot() string {
	if len(c.RecordsRoots) == 0 {
		return ""
	}
	return c.RecordsRoots[0].RecordsRoot
}

// readDir is the package-level seam for directory readability checks.
// Production default is os.ReadDir. Tests may replace it for a single test
// and must restore the original via defer before the test returns.
var readDir func(string) ([]os.DirEntry, error) = os.ReadDir

// NormalizeConfig validates an explicitly configured set of current roots.
//
// root must be non-empty and resolvable to an absolute path.
// roots must be non-empty.
// Each root must carry a non-empty app_namespace and a records_root that exactly
// matches <app_namespace>/records.
// Duplicate records_root declarations and duplicate app_namespace values fail the
// complete configuration.
// Every records_root is verified to exist as a readable directory within root.
// A valid but empty records directory is accepted.
// Any invalid root fails the complete configuration.
func NormalizeConfig(root string, roots []CurrentRoot) (Config, error) {
	if root == "" {
		return Config{}, errors.New("repository root must not be empty")
	}
	abs, err := filepath.Abs(root)
	if err != nil {
		return Config{}, fmt.Errorf("resolve root: %w", err)
	}
	cleanRoot := filepath.Clean(abs)

	if len(roots) == 0 {
		return Config{}, errors.New("at least one current root must be configured")
	}

	seenRoot := map[string]bool{}
	seenNS := map[string]bool{}
	entries := make([]RecordsEntry, 0, len(roots))

	for i, cr := range roots {
		if cr.AppNamespace == "" {
			return Config{}, fmt.Errorf("current_roots[%d]: app_namespace must not be empty", i)
		}
		if cr.RecordsRoot == "" {
			return Config{}, fmt.Errorf("current_roots[%d] (%q): records_root must not be empty", i, cr.AppNamespace)
		}

		// Exact app-root shape: records_root must equal <app_namespace>/records.
		want := cr.AppNamespace + "/records"
		cleanRel := filepath.ToSlash(filepath.Clean(cr.RecordsRoot))
		if cleanRel != want {
			return Config{}, fmt.Errorf("current_roots[%d]: records_root %q does not match required shape %q",
				i, cr.RecordsRoot, want)
		}

		// Root containment: resolve absolute path and confirm no parent escape.
		absDir := filepath.Join(cleanRoot, filepath.FromSlash(cleanRel))
		rel, relErr := filepath.Rel(cleanRoot, absDir)
		if relErr != nil || strings.HasPrefix(filepath.ToSlash(rel), "..") {
			return Config{}, fmt.Errorf("current_roots[%d]: records_root %q escapes repository root", i, cr.RecordsRoot)
		}
		normalizedRel := filepath.ToSlash(rel)

		// Duplicate records_root declaration.
		if seenRoot[normalizedRel] {
			return Config{}, fmt.Errorf("current_roots[%d]: duplicate records_root declaration %q", i, cr.RecordsRoot)
		}
		seenRoot[normalizedRel] = true

		// Duplicate app_namespace.
		if seenNS[cr.AppNamespace] {
			return Config{}, fmt.Errorf("current_roots[%d]: duplicate app_namespace %q", i, cr.AppNamespace)
		}
		seenNS[cr.AppNamespace] = true

		// Directory existence and readability.
		info, statErr := os.Stat(absDir)
		if statErr != nil {
			return Config{}, fmt.Errorf("current_roots[%d] (%q): records_root %q: %w",
				i, cr.AppNamespace, cr.RecordsRoot, statErr)
		}
		if !info.IsDir() {
			return Config{}, fmt.Errorf("current_roots[%d] (%q): records_root %q is not a directory",
				i, cr.AppNamespace, cr.RecordsRoot)
		}
		if _, rdErr := readDir(absDir); rdErr != nil {
			return Config{}, fmt.Errorf("current_roots[%d] (%q): records_root %q is not readable: %w",
				i, cr.AppNamespace, cr.RecordsRoot, rdErr)
		}

		entries = append(entries, RecordsEntry{
			AppNamespace:    cr.AppNamespace,
			RecordsRoot:     normalizedRel,
			NamespacePrefix: strings.ToUpper(cr.AppNamespace) + "-",
		})
	}

	return Config{Root: cleanRoot, RecordsRoots: entries}, nil
}

// NewConfig builds a Config from a repository root and a single explicit records root path.
// root may be empty (resolved to the current working directory).
// recordsRoot must be non-empty and match the <app_namespace>/records shape.
// Auto-discovery of */records directories is removed.
// The v01/records default fallback is removed.
func NewConfig(root, recordsRoot string) (Config, error) {
	if root == "" {
		cwd, err := os.Getwd()
		if err != nil {
			return Config{}, fmt.Errorf("resolve cwd: %w", err)
		}
		root = cwd
	}
	if recordsRoot == "" {
		return Config{}, errors.New("records_root must not be empty: auto-discovery has been removed")
	}

	// Derive app_namespace from the records_root shape: exactly <app_namespace>/records.
	cleanRel := filepath.ToSlash(filepath.Clean(recordsRoot))
	slashIdx := strings.Index(cleanRel, "/")
	if slashIdx <= 0 || cleanRel[slashIdx:] != "/records" {
		return Config{}, fmt.Errorf("records_root %q does not match required shape <app_namespace>/records", recordsRoot)
	}
	appNS := cleanRel[:slashIdx]

	return NormalizeConfig(root, []CurrentRoot{{AppNamespace: appNS, RecordsRoot: recordsRoot}})
}

// normalizeConfig resolves Root to an absolute path and verifies RecordsRoots is non-empty.
// Auto-discovery of */records directories and the v01/records fallback are removed.
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
		return Config{}, errors.New("current roots must not be empty: auto-discovery has been removed")
	}
	return result, nil
}
