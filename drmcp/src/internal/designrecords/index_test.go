package designrecords

import (
	"context"
	"testing"
)

func TestBuildIndexEmptyState(t *testing.T) {
	cfg, err := NewConfig(t.TempDir(), "")
	if err != nil {
		t.Fatalf("NewConfig: %v", err)
	}
	idx, err := BuildIndex(context.Background(), cfg)
	if err != nil {
		t.Fatalf("BuildIndex: %v", err)
	}
	if idx.Root != cfg.Root {
		t.Fatalf("Root = %q, want %q", idx.Root, cfg.Root)
	}
	if len(idx.Records) != 0 {
		t.Fatalf("Records len = %d, want 0", len(idx.Records))
	}
	if len(idx.Diagnostics) != 0 {
		t.Fatalf("Diagnostics len = %d, want 0", len(idx.Diagnostics))
	}
}
