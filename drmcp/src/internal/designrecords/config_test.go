package designrecords

import (
	"os"
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

func TestNewConfigSingleRoot(t *testing.T) {
	root := t.TempDir()
	cfg, err := NewConfig(root, "v01/records")
	if err != nil {
		t.Fatalf("NewConfig: %v", err)
	}
	if len(cfg.RecordsRoots) != 1 {
		t.Fatalf("RecordsRoots len = %d, want 1", len(cfg.RecordsRoots))
	}
	if cfg.RecordsRoots[0].RecordsRoot != "v01/records" {
		t.Fatalf("RecordsRoot = %q, want v01/records", cfg.RecordsRoots[0].RecordsRoot)
	}
	if cfg.RecordsRoots[0].NamespacePrefix != "V01-" {
		t.Fatalf("NamespacePrefix = %q, want V01-", cfg.RecordsRoots[0].NamespacePrefix)
	}
	if cfg.NamespacePrefix() != "V01-" {
		t.Fatalf("NamespacePrefix() = %q, want V01-", cfg.NamespacePrefix())
	}
	if cfg.primaryRecordsRoot() != "v01/records" {
		t.Fatalf("primaryRecordsRoot() = %q, want v01/records", cfg.primaryRecordsRoot())
	}
}

func TestNewConfigAutoDetectMultiRoot(t *testing.T) {
	root := t.TempDir()
	// Create two */records/ directories.
	for _, p := range []string{"v01/records", "drmcp/records"} {
		if err := os.MkdirAll(filepath.Join(root, filepath.FromSlash(p)), 0755); err != nil {
			t.Fatalf("mkdir %s: %v", p, err)
		}
	}

	cfg, err := NewConfig(root, "")
	if err != nil {
		t.Fatalf("NewConfig: %v", err)
	}
	if len(cfg.RecordsRoots) != 2 {
		t.Fatalf("RecordsRoots len = %d, want 2: %+v", len(cfg.RecordsRoots), cfg.RecordsRoots)
	}
	// Both entries must have correct namespace prefixes.
	prefixes := map[string]bool{}
	for _, e := range cfg.RecordsRoots {
		prefixes[e.NamespacePrefix] = true
	}
	if !prefixes["V01-"] || !prefixes["DRMCP-"] {
		t.Fatalf("expected prefixes V01- and DRMCP-, got %+v", cfg.RecordsRoots)
	}
}

func TestNewConfigAutoDetectFallback(t *testing.T) {
	// No */records/ dirs → falls back to v01/records default.
	root := t.TempDir()
	cfg, err := NewConfig(root, "")
	if err != nil {
		t.Fatalf("NewConfig: %v", err)
	}
	if len(cfg.RecordsRoots) != 1 {
		t.Fatalf("RecordsRoots len = %d, want 1 (fallback)", len(cfg.RecordsRoots))
	}
	if cfg.RecordsRoots[0].RecordsRoot != "v01/records" {
		t.Fatalf("fallback RecordsRoot = %q, want v01/records", cfg.RecordsRoots[0].RecordsRoot)
	}
}

func TestMakeRecordsEntry(t *testing.T) {
	cases := []struct {
		input   string
		wantRoot string
		wantNS   string
	}{
		{"v01/records", "v01/records", "V01-"},
		{"drmcp/records", "drmcp/records", "DRMCP-"},
		{"docs", "docs", ""},
	}
	for _, c := range cases {
		e := makeRecordsEntry(c.input)
		if e.RecordsRoot != c.wantRoot {
			t.Errorf("makeRecordsEntry(%q).RecordsRoot = %q, want %q", c.input, e.RecordsRoot, c.wantRoot)
		}
		if e.NamespacePrefix != c.wantNS {
			t.Errorf("makeRecordsEntry(%q).NamespacePrefix = %q, want %q", c.input, e.NamespacePrefix, c.wantNS)
		}
	}
}
