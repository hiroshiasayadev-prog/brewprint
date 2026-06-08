package designrecords

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type authoringGuide struct {
	ID       string
	Title    string
	Abstract string
	Content  string
}

func ListAuthoringGuides(ctx context.Context, cfg Config, req ListAuthoringGuidesRequest) (ListAuthoringGuidesResponse, error) {
	if err := ctx.Err(); err != nil {
		return ListAuthoringGuidesResponse{}, err
	}
	guides, err := loadAuthoringGuides(ctx, cfg)
	if err != nil {
		return ListAuthoringGuidesResponse{}, err
	}
	out := make([]AuthoringGuideSummary, 0, len(guides))
	for _, guide := range guides {
		out = append(out, AuthoringGuideSummary{
			ID:       guide.ID,
			Title:    guide.Title,
			Abstract: guide.Abstract,
		})
	}
	return ListAuthoringGuidesResponse{Guides: out}, nil
}

func GetAuthoringGuidance(ctx context.Context, cfg Config, req GetAuthoringGuidanceRequest) (GetAuthoringGuidanceResponse, error) {
	if err := ctx.Err(); err != nil {
		return GetAuthoringGuidanceResponse{}, err
	}
	if req.ID == "" {
		return GetAuthoringGuidanceResponse{}, newToolError(ErrorCodeInvalidRequest, "id is required")
	}
	guides, err := loadAuthoringGuides(ctx, cfg)
	if err != nil {
		return GetAuthoringGuidanceResponse{}, err
	}
	for _, guide := range guides {
		if guide.ID == req.ID {
			return GetAuthoringGuidanceResponse{
				ID:      guide.ID,
				Title:   guide.Title,
				Content: guide.Content,
			}, nil
		}
	}
	return GetAuthoringGuidanceResponse{}, newToolError(ErrorCodeGuideNotFound, fmt.Sprintf("authoring guide %s was not found", req.ID))
}

func loadAuthoringGuides(ctx context.Context, cfg Config) ([]authoringGuide, error) {
	normalized, err := normalizeConfig(cfg)
	if err != nil {
		return nil, newToolError(ErrorCodeInvalidRequest, err.Error())
	}
	guideRoot := filepath.Join(normalized.Root, filepath.FromSlash(normalized.RecordsRoot), "guides")
	if _, err := os.Stat(guideRoot); os.IsNotExist(err) {
		return []authoringGuide{}, nil
	} else if err != nil {
		return nil, newToolError(ErrorCodeInvalidRequest, fmt.Sprintf("stat authoring guide root: %v", err))
	}

	matches, err := filepath.Glob(filepath.Join(guideRoot, "*.md"))
	if err != nil {
		return nil, newToolError(ErrorCodeInvalidRequest, fmt.Sprintf("discover authoring guides: %v", err))
	}
	guides := make([]authoringGuide, 0, len(matches))
	for _, path := range matches {
		if err := ctx.Err(); err != nil {
			return nil, err
		}
		content, err := os.ReadFile(path)
		if err != nil {
			return nil, newToolError(ErrorCodeInvalidRequest, fmt.Sprintf("read authoring guide %s: %v", filepath.Base(path), err))
		}
		id := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
		raw := string(content)
		title, _, _ := parseSpecH1(raw)
		guides = append(guides, authoringGuide{
			ID:       id,
			Title:    title,
			Abstract: extractAuthoringGuideAbstract(raw),
			Content:  raw,
		})
	}
	sort.SliceStable(guides, func(i, j int) bool {
		return guides[i].ID < guides[j].ID
	})
	return guides, nil
}

func extractAuthoringGuideAbstract(raw string) string {
	lines := splitMarkdownLines(raw)
	startLine := 0
	if hasOpeningFrontMatter(lines) {
		for i := 1; i < len(lines); i++ {
			if strings.TrimSpace(trimLineEnd(lines[i])) == "---" {
				startLine = i + 1
				break
			}
		}
	}

	start := -1
	end := len(lines)
	inFence := false
	fenceMarker := ""
	for i := startLine; i < len(lines); i++ {
		trimmed := strings.TrimSpace(trimLineEnd(lines[i]))
		if isFenceLine(trimmed) {
			marker := fencePrefix(trimmed)
			if !inFence {
				inFence = true
				fenceMarker = marker
			} else if marker == fenceMarker {
				inFence = false
				fenceMarker = ""
			}
			continue
		}
		if inFence {
			continue
		}

		match := atxHeadingPattern.FindStringSubmatch(trimLineEnd(lines[i]))
		if match == nil {
			continue
		}
		if start == -1 && len(match[1]) == 2 && strings.TrimSpace(match[2]) == "Abstract" {
			start = i + 1
			continue
		}
		if start != -1 && len(match[1]) <= 2 {
			end = i
			break
		}
	}
	if start == -1 {
		return ""
	}
	return strings.TrimSpace(strings.Join(lines[start:end], "\n"))
}
