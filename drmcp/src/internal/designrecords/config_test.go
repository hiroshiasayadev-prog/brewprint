package designrecords

import (
	"path/filepath"
	"testing"
)

func TestNewConfigDefaultsToCWD(t *testing.T) {
	root := t.TempDir()
	t.Chdir(root)

	cfg, err := NewConfig("", "")
	if err != nil {
		t.Fatalf("NewConfig: %v", err)
	}
	want, err := filepath.Abs(root)
	if err != nil {
		t.Fatalf("Abs: %v", err)
	}
	if cfg.Root != filepath.Clean(want) {
		t.Fatalf("Root = %q, want %q", cfg.Root, filepath.Clean(want))
	}
}

func TestNewConfigNormalizesRoot(t *testing.T) {
	root := t.TempDir()
	cfg, err := NewConfig(filepath.Join(root, "."), "")
	if err != nil {
		t.Fatalf("NewConfig: %v", err)
	}
	want, err := filepath.Abs(root)
	if err != nil {
		t.Fatalf("Abs: %v", err)
	}
	if cfg.Root != filepath.Clean(want) {
		t.Fatalf("Root = %q, want %q", cfg.Root, filepath.Clean(want))
	}
}
