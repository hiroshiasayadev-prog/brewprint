package designrecords

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Config contains repository-local Design Records MCP configuration.
type Config struct {
	// Root is the absolute repository root used for discovery. Response paths
	// are expected to be relative to this root.
	Root string
	// RecordsRoot is the path to the records directory relative to Root
	// (e.g. "v01/records"). Defaults to "v01/records" when empty.
	RecordsRoot string
}

// NamespacePrefix derives the namespace prefix from the app namespace directory
// name — the parent of the final path component in RecordsRoot.
// "v01/records" → appNamespaceDir "v01" → "V01-"
func (c Config) NamespacePrefix() string {
	if c.RecordsRoot == "" {
		return ""
	}
	clean := filepath.Clean(filepath.FromSlash(c.RecordsRoot))
	appNS := filepath.Base(filepath.Dir(clean))
	if appNS == "." || appNS == "" {
		return ""
	}
	return strings.ToUpper(appNS) + "-"
}

// NewConfig resolves root from an explicit path or the current working
// directory when root is empty. recordsRoot is the path to the records
// directory relative to root (default: "v01/records").
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
	if recordsRoot == "" {
		recordsRoot = "v01/records"
	}
	return Config{Root: filepath.Clean(abs), RecordsRoot: recordsRoot}, nil
}

func normalizeConfig(cfg Config) (Config, error) {
	return NewConfig(cfg.Root, cfg.RecordsRoot)
}
