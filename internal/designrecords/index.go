package designrecords

import (
	"context"
)

// BuildIndex returns an empty index in Phase 0.
//
// TODO(Phase 1): discover docs/adr/*.md and docs/spec/**/*.md under Config.Root,
// parse Markdown metadata, and populate Records and Diagnostics.
func BuildIndex(ctx context.Context, cfg Config) (*Index, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	normalized, err := normalizeConfig(cfg)
	if err != nil {
		return nil, err
	}
	return &Index{
		Root:        normalized.Root,
		Records:     []Record{},
		Diagnostics: []Diagnostic{},
	}, nil
}
