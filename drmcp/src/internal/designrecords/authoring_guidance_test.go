package designrecords

import (
	"context"
	"encoding/json"
	"testing"
)

func TestListAuthoringGuides(t *testing.T) {
	root := t.TempDir()
	writeTestFile(t, root, "docs/guides/zeta.md", "# Zeta Guide\n\n## Abstract\n\nZeta summary.\n\n## Body\n\nZeta body.\n")
	writeTestFile(t, root, "docs/guides/alpha.md", "# Alpha Guide\n\n## Abstract\n\nAlpha summary line 1.\n\nAlpha summary line 2.\n\n### Detail\n\nKept in abstract.\n\n## Body\n\nAlpha body.\n")
	writeTestFile(t, root, "docs/guides/not-markdown.txt", "# Ignored\n\n## Abstract\n\nIgnored.\n")

	cfg, err := NewConfig(root, "")
	if err != nil {
		t.Fatalf("NewConfig: %v", err)
	}
	resp, err := ListAuthoringGuides(context.Background(), cfg, ListAuthoringGuidesRequest{})
	if err != nil {
		t.Fatalf("ListAuthoringGuides: %v", err)
	}
	if len(resp.Guides) != 2 {
		t.Fatalf("guides len = %d, want 2: %#v", len(resp.Guides), resp.Guides)
	}
	if resp.Guides[0].ID != "alpha" || resp.Guides[1].ID != "zeta" {
		t.Fatalf("guide order = %#v, want ASCII lexical ID order", resp.Guides)
	}
	if resp.Guides[0].Title != "Alpha Guide" {
		t.Fatalf("title = %q, want Alpha Guide", resp.Guides[0].Title)
	}
	wantAbstract := "Alpha summary line 1.\n\nAlpha summary line 2.\n\n### Detail\n\nKept in abstract."
	if resp.Guides[0].Abstract != wantAbstract {
		t.Fatalf("abstract = %q, want %q", resp.Guides[0].Abstract, wantAbstract)
	}
	assertJSONKeys(t, resp.Guides[0], []string{"abstract", "id", "title"})
}

func TestGetAuthoringGuidance(t *testing.T) {
	root := t.TempDir()
	content := "# ADR Authoring Guide\n\n## Abstract\n\nADR summary.\n\n## Body\n\nKeep this Markdown exactly.\n"
	writeTestFile(t, root, "docs/guides/adr-authoring.md", content)

	cfg, err := NewConfig(root, "")
	if err != nil {
		t.Fatalf("NewConfig: %v", err)
	}
	resp, err := GetAuthoringGuidance(context.Background(), cfg, GetAuthoringGuidanceRequest{ID: "adr-authoring"})
	if err != nil {
		t.Fatalf("GetAuthoringGuidance: %v", err)
	}
	if resp.ID != "adr-authoring" || resp.Title != "ADR Authoring Guide" || resp.Content != content {
		t.Fatalf("response = %#v", resp)
	}
	assertJSONKeys(t, resp, []string{"content", "id", "title"})
}

func TestGetAuthoringGuidanceErrors(t *testing.T) {
	root := t.TempDir()
	writeTestFile(t, root, "docs/guides/adr-authoring.md", "# ADR Authoring Guide\n\n## Abstract\n\nADR summary.\n")
	cfg, err := NewConfig(root, "")
	if err != nil {
		t.Fatalf("NewConfig: %v", err)
	}

	_, err = GetAuthoringGuidance(context.Background(), cfg, GetAuthoringGuidanceRequest{})
	assertToolErrorCodeDirect(t, err, ErrorCodeInvalidRequest)

	_, err = GetAuthoringGuidance(context.Background(), cfg, GetAuthoringGuidanceRequest{ID: "unknown-guide"})
	assertToolErrorCodeDirect(t, err, ErrorCodeGuideNotFound)
}

func TestBuildIndexIgnoresAuthoringGuides(t *testing.T) {
	root := t.TempDir()
	writeTestFile(t, root, "docs/guides/adr-authoring.md", "# ADR Authoring Guide\n\n## Abstract\n\nADR summary.\n")

	cfg, err := NewConfig(root, "")
	if err != nil {
		t.Fatalf("NewConfig: %v", err)
	}
	idx, err := BuildIndex(context.Background(), cfg)
	if err != nil {
		t.Fatalf("BuildIndex: %v", err)
	}
	if len(idx.Records) != 0 {
		t.Fatalf("records = %#v, want no authoring guides in record index", idx.Records)
	}
}

func assertJSONKeys(t *testing.T, value any, want []string) {
	t.Helper()
	data, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	var got map[string]any
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if len(got) != len(want) {
		t.Fatalf("JSON keys = %#v, want %#v", got, want)
	}
	for _, key := range want {
		if _, ok := got[key]; !ok {
			t.Fatalf("JSON missing key %q in %#v", key, got)
		}
	}
}

func assertToolErrorCodeDirect(t *testing.T, err error, want ErrorCode) {
	t.Helper()
	toolErr, ok := err.(*ToolError)
	if !ok {
		t.Fatalf("error = %T %v, want *ToolError", err, err)
	}
	if toolErr.Code != want {
		t.Fatalf("tool error code = %q, want %q", toolErr.Code, want)
	}
}
