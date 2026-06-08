package designrecordsmcp

import (
	"context"

	"github.com/hiroshiasayadev-prog/brewprint/drmcp/src/internal/designrecords"
)

type IndexBuilder func(context.Context, designrecords.Config) (*designrecords.Index, error)

type Server struct {
	cfg            designrecords.Config
	buildIndex     IndexBuilder
	authoringStore *designrecords.AuthoringStore
}

func NewServer(cfg designrecords.Config) *Server {
	return NewServerWithIndexBuilder(cfg, designrecords.BuildIndex)
}

func NewServerWithIndexBuilder(cfg designrecords.Config, buildIndex IndexBuilder) *Server {
	if buildIndex == nil {
		buildIndex = designrecords.BuildIndex
	}
	return &Server{
		cfg:            cfg,
		buildIndex:     buildIndex,
		authoringStore: designrecords.NewAuthoringStore(),
	}
}

func (s *Server) Config() designrecords.Config {
	if s == nil {
		return designrecords.Config{}
	}
	return s.cfg
}

func (s *Server) AuthoringStoreForTest() *designrecords.AuthoringStore {
	if s == nil {
		return nil
	}
	return s.authoringStore
}
