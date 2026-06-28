package designrecords

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

// TestNewConfigNormalizesRoot confirms that a non-clean root path is resolved to
// its absolute clean form.
func TestNewConfigNormalizesRoot(t *testing.T) {
	root := t.TempDir()
	if err := os.MkdirAll(filepath.Join(root, "product", "records"), 0755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	cfg, err := NewConfig(filepath.Join(root, "."), "product/records")
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

// TestNewConfigSingleRoot confirms that a single explicit records root is accepted
// and produces correct AppNamespace, RecordsRoot, and NamespacePrefix.
func TestNewConfigSingleRoot(t *testing.T) {
	root := t.TempDir()
	if err := os.MkdirAll(filepath.Join(root, "drmcp", "records"), 0755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	cfg, err := NewConfig(root, "drmcp/records")
	if err != nil {
		t.Fatalf("NewConfig: %v", err)
	}
	if len(cfg.RecordsRoots) != 1 {
		t.Fatalf("RecordsRoots len = %d, want 1", len(cfg.RecordsRoots))
	}
	e := cfg.RecordsRoots[0]
	if e.AppNamespace != "drmcp" {
		t.Errorf("AppNamespace = %q, want %q", e.AppNamespace, "drmcp")
	}
	if e.RecordsRoot != "drmcp/records" {
		t.Errorf("RecordsRoot = %q, want %q", e.RecordsRoot, "drmcp/records")
	}
	if e.NamespacePrefix != "DRMCP-" {
		t.Errorf("NamespacePrefix = %q, want %q", e.NamespacePrefix, "DRMCP-")
	}
	if cfg.NamespacePrefix() != "DRMCP-" {
		t.Errorf("cfg.NamespacePrefix() = %q, want %q", cfg.NamespacePrefix(), "DRMCP-")
	}
	if cfg.primaryRecordsRoot() != "drmcp/records" {
		t.Errorf("cfg.primaryRecordsRoot() = %q, want %q", cfg.primaryRecordsRoot(), "drmcp/records")
	}
}

// TestNewConfigEmptyRecordsRootRejected confirms that auto-discovery is removed:
// an empty records_root is rejected.
func TestNewConfigEmptyRecordsRootRejected(t *testing.T) {
	root := t.TempDir()
	_, err := NewConfig(root, "")
	if err == nil {
		t.Fatal("NewConfig with empty records_root: want error, got nil")
	}
}

// TestNewConfigShapeMismatchRejected confirms that a single-component records_root
// (no app_namespace parent) is rejected.
func TestNewConfigShapeMismatchRejected(t *testing.T) {
	root := t.TempDir()
	_, err := NewConfig(root, "records")
	if err == nil {
		t.Fatal("NewConfig with single-component records_root: want error, got nil")
	}
}

// TestCurrentRootTwoRootsAccepted covers fixture case C08:
// two unique current app roots with distinct app_namespace values are accepted.
func TestCurrentRootTwoRootsAccepted(t *testing.T) {
	root := t.TempDir()
	for _, p := range []string{"product/records", "drmcp/records"} {
		if err := os.MkdirAll(filepath.Join(root, filepath.FromSlash(p)), 0755); err != nil {
			t.Fatalf("mkdir %s: %v", p, err)
		}
	}
	cfg, err := NormalizeConfig(root, []CurrentRoot{
		{AppNamespace: "product", RecordsRoot: "product/records"},
		{AppNamespace: "drmcp", RecordsRoot: "drmcp/records"},
	})
	if err != nil {
		t.Fatalf("NormalizeConfig two roots: %v", err)
	}
	if len(cfg.RecordsRoots) != 2 {
		t.Fatalf("RecordsRoots len = %d, want 2", len(cfg.RecordsRoots))
	}
	if cfg.RecordsRoots[0].AppNamespace != "product" {
		t.Errorf("[0].AppNamespace = %q, want %q", cfg.RecordsRoots[0].AppNamespace, "product")
	}
	if cfg.RecordsRoots[0].NamespacePrefix != "PRODUCT-" {
		t.Errorf("[0].NamespacePrefix = %q, want %q", cfg.RecordsRoots[0].NamespacePrefix, "PRODUCT-")
	}
	if cfg.RecordsRoots[1].AppNamespace != "drmcp" {
		t.Errorf("[1].AppNamespace = %q, want %q", cfg.RecordsRoots[1].AppNamespace, "drmcp")
	}
	if cfg.RecordsRoots[1].NamespacePrefix != "DRMCP-" {
		t.Errorf("[1].NamespacePrefix = %q, want %q", cfg.RecordsRoots[1].NamespacePrefix, "DRMCP-")
	}
	if cfg.NamespacePrefix() != "PRODUCT-" {
		t.Errorf("cfg.NamespacePrefix() = %q, want %q", cfg.NamespacePrefix(), "PRODUCT-")
	}
}

// TestCurrentRootCurrentOnlyConfig covers fixture case C10:
// a configuration with only current roots and no legacy_roots is accepted.
func TestCurrentRootCurrentOnlyConfig(t *testing.T) {
	root := t.TempDir()
	if err := os.MkdirAll(filepath.Join(root, "product", "records"), 0755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	cfg, err := NormalizeConfig(root, []CurrentRoot{
		{AppNamespace: "product", RecordsRoot: "product/records"},
	})
	if err != nil {
		t.Fatalf("NormalizeConfig current-only: %v", err)
	}
	if len(cfg.RecordsRoots) != 1 {
		t.Fatalf("RecordsRoots len = %d, want 1", len(cfg.RecordsRoots))
	}
}

// TestCurrentRootMissingDirectory covers fixture case R08:
// a current root pointing to a missing directory fails the complete configuration.
func TestCurrentRootMissingDirectory(t *testing.T) {
	root := t.TempDir()
	// product/records is intentionally not created.
	_, err := NormalizeConfig(root, []CurrentRoot{
		{AppNamespace: "product", RecordsRoot: "product/records"},
	})
	if err == nil {
		t.Fatal("NormalizeConfig missing directory: want error, got nil")
	}
}

// TestCurrentRootDuplicateDeclaration covers fixture case R10:
// declaring the same current records_root twice fails the complete configuration.
func TestCurrentRootDuplicateDeclaration(t *testing.T) {
	root := t.TempDir()
	if err := os.MkdirAll(filepath.Join(root, "product", "records"), 0755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	_, err := NormalizeConfig(root, []CurrentRoot{
		{AppNamespace: "product", RecordsRoot: "product/records"},
		{AppNamespace: "product", RecordsRoot: "product/records"},
	})
	if err == nil {
		t.Fatal("NormalizeConfig duplicate declaration: want error, got nil")
	}
}

// TestCurrentRootContainmentViolation confirms that a records_root whose resolved
// absolute path escapes the repository root is rejected.
func TestCurrentRootContainmentViolation(t *testing.T) {
	root := t.TempDir()
	// app_namespace "../sibling" → records_root "../sibling/records"
	// shape check passes; containment check must reject.
	_, err := NormalizeConfig(root, []CurrentRoot{
		{AppNamespace: "../sibling", RecordsRoot: "../sibling/records"},
	})
	if err == nil {
		t.Fatal("NormalizeConfig containment violation: want error, got nil")
	}
}

// TestCurrentRootAppNamespaceMismatch confirms that a records_root whose path
// does not match <app_namespace>/records is rejected.
func TestCurrentRootAppNamespaceMismatch(t *testing.T) {
	root := t.TempDir()
	// app_namespace is "product" but records_root is "other/records".
	_, err := NormalizeConfig(root, []CurrentRoot{
		{AppNamespace: "product", RecordsRoot: "other/records"},
	})
	if err == nil {
		t.Fatal("NormalizeConfig app_namespace/records_root mismatch: want error, got nil")
	}
}

// TestNormalizeConfigEmptyRoot confirms that an empty repository root is rejected.
func TestNormalizeConfigEmptyRoot(t *testing.T) {
	_, err := NormalizeConfig("", []CurrentRoot{
		{AppNamespace: "product", RecordsRoot: "product/records"},
	})
	if err == nil {
		t.Fatal("NormalizeConfig empty root: want error, got nil")
	}
}

// TestNormalizeConfigEmptyRoots confirms that an empty roots list is rejected.
func TestNormalizeConfigEmptyRoots(t *testing.T) {
	root := t.TempDir()
	if _, err := NormalizeConfig(root, nil); err == nil {
		t.Fatal("NormalizeConfig nil roots: want error, got nil")
	}
	if _, err := NormalizeConfig(root, []CurrentRoot{}); err == nil {
		t.Fatal("NormalizeConfig empty roots slice: want error, got nil")
	}
}

// TestNormalizeConfigMissingAppNamespace confirms that an empty app_namespace is rejected.
func TestNormalizeConfigMissingAppNamespace(t *testing.T) {
	root := t.TempDir()
	_, err := NormalizeConfig(root, []CurrentRoot{
		{AppNamespace: "", RecordsRoot: "product/records"},
	})
	if err == nil {
		t.Fatal("NormalizeConfig missing app_namespace: want error, got nil")
	}
}

// TestNormalizeConfigMissingRecordsRoot confirms that an empty records_root is rejected.
func TestNormalizeConfigMissingRecordsRoot(t *testing.T) {
	root := t.TempDir()
	_, err := NormalizeConfig(root, []CurrentRoot{
		{AppNamespace: "product", RecordsRoot: ""},
	})
	if err == nil {
		t.Fatal("NormalizeConfig missing records_root: want error, got nil")
	}
}

// TestNormalizeConfigValidEmptyTree confirms that a current root that is an existing
// but empty directory is accepted.
func TestNormalizeConfigValidEmptyTree(t *testing.T) {
	root := t.TempDir()
	if err := os.MkdirAll(filepath.Join(root, "product", "records"), 0755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	cfg, err := NormalizeConfig(root, []CurrentRoot{
		{AppNamespace: "product", RecordsRoot: "product/records"},
	})
	if err != nil {
		t.Fatalf("NormalizeConfig valid empty tree: %v", err)
	}
	if len(cfg.RecordsRoots) != 1 {
		t.Fatalf("RecordsRoots len = %d, want 1", len(cfg.RecordsRoots))
	}
}

// TestNormalizeConfigRootNormalization confirms that a non-clean root path is
// resolved to its absolute clean form by NormalizeConfig.
func TestNormalizeConfigRootNormalization(t *testing.T) {
	root := t.TempDir()
	if err := os.MkdirAll(filepath.Join(root, "product", "records"), 0755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	cfg, err := NormalizeConfig(filepath.Join(root, "."), []CurrentRoot{
		{AppNamespace: "product", RecordsRoot: "product/records"},
	})
	if err != nil {
		t.Fatalf("NormalizeConfig: %v", err)
	}
	want, _ := filepath.Abs(root)
	if cfg.Root != filepath.Clean(want) {
		t.Fatalf("Root = %q, want %q", cfg.Root, filepath.Clean(want))
	}
}

// TestNormalizeConfigNotADirectory confirms that a records_root that resolves to
// a file rather than a directory is rejected.
func TestNormalizeConfigNotADirectory(t *testing.T) {
	root := t.TempDir()
	if err := os.MkdirAll(filepath.Join(root, "product"), 0755); err != nil {
		t.Fatalf("mkdir product: %v", err)
	}
	// Create a file named "records" where a directory is expected.
	f, err := os.Create(filepath.Join(root, "product", "records"))
	if err != nil {
		t.Fatalf("create file: %v", err)
	}
	f.Close()
	_, err = NormalizeConfig(root, []CurrentRoot{
		{AppNamespace: "product", RecordsRoot: "product/records"},
	})
	if err == nil {
		t.Fatal("NormalizeConfig records_root is file: want error, got nil")
	}
}

// TestNormalizeConfigRejectsEmptyRoots confirms that the private normalizeConfig
// rejects a Config with empty RecordsRoots (auto-discovery removed).
func TestNormalizeConfigRejectsEmptyRoots(t *testing.T) {
	cfg := Config{Root: t.TempDir()}
	_, err := normalizeConfig(cfg)
	if err == nil {
		t.Fatal("normalizeConfig empty RecordsRoots: want error, got nil")
	}
}

// TestCurrentRootUnreadableDirectory confirms that a current root whose directory
// cannot be read fails the complete configuration.
// This test uses the package-level readDir seam and must not run in parallel.
func TestCurrentRootUnreadableDirectory(t *testing.T) {
	root := t.TempDir()
	if err := os.MkdirAll(filepath.Join(root, "product", "records"), 0755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	absTarget := filepath.Clean(filepath.Join(root, "product", "records"))

	orig := readDir
	readDir = func(path string) ([]os.DirEntry, error) {
		if filepath.Clean(path) == absTarget {
			return nil, errors.New("simulated unreadable directory")
		}
		return orig(path)
	}
	defer func() { readDir = orig }()

	_, err := NormalizeConfig(root, []CurrentRoot{
		{AppNamespace: "product", RecordsRoot: "product/records"},
	})
	if err == nil {
		t.Fatal("NormalizeConfig unreadable directory: want error, got nil")
	}
}
