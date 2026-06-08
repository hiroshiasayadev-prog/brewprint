package designrecords

import (
	"fmt"
	"os"
	"path/filepath"
)

// Config contains repository-local Design Records MCP configuration.
type Config struct {
	// Root is the absolute repository root used for discovery. Response paths
	// are expected to be relative to this root.
	Root string
}

// NewConfig resolves root from an explicit path or the current working
// directory when root is empty.
func NewConfig(root string) (Config, error) {
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
	return Config{Root: filepath.Clean(abs)}, nil
}

func normalizeConfig(cfg Config) (Config, error) {
	return NewConfig(cfg.Root)
}
