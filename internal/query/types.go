package query

import "github.com/hiroshiasayadev-prog/brewprint/internal/semantic"

type Selector struct {
	ID      string `json:"id"`
	Object  string `json:"object,omitempty"`
	Kind    string `json:"kind,omitempty"`
	File    string `json:"file,omitempty"`
	LocalID string `json:"local_id,omitempty"`
}

type ObjectRef struct {
	Object      string            `json:"object"`
	Kind        string            `json:"kind"`
	ID          string            `json:"id"`
	QualifiedID string            `json:"qualified_id,omitempty"`
	Label       string            `json:"label,omitempty"`
	Module      string            `json:"module,omitempty"`
	File        string            `json:"file,omitempty"`
	LocalID     string            `json:"local_id,omitempty"`
	Source      map[string]string `json:"source,omitempty"`
}

type AssetRef struct {
	Object    string `json:"object"`
	ID        string `json:"id,omitempty"`
	Name      string `json:"name"`
	Producer  string `json:"producer"`
	Model     string `json:"model"`
	ScopeFile string `json:"scope_file,omitempty"`
}

type FieldRef struct {
	Object string `json:"object"`
	ID     string `json:"id"`
	Model  string `json:"model"`
	Name   string `json:"name"`
}

type TransitionRef struct {
	Object    string `json:"object"`
	Kind      string `json:"kind"`
	ID        string `json:"id"`
	StateFile string `json:"state_file"`
	From      string `json:"from"`
	On        string `json:"on"`
	To        string `json:"to"`
	Guard     string `json:"guard,omitempty"`
	Action    string `json:"action,omitempty"`
}

type ScenarioStepRef struct {
	Index           int           `json:"index"`
	FromState       string        `json:"from_state"`
	Via             string        `json:"via"`
	Guard           string        `json:"guard,omitempty"`
	GuardExactMatch bool          `json:"guard_exact_match"`
	Transition      TransitionRef `json:"transition"`
	Action          *string       `json:"action"`
}

type ReferenceEndpoint struct {
	Object      string `json:"object,omitempty"`
	Kind        string `json:"kind,omitempty"`
	ID          string `json:"id,omitempty"`
	QualifiedID string `json:"qualified_id,omitempty"`
	Name        string `json:"name,omitempty"`
	Producer    string `json:"producer,omitempty"`
	Model       string `json:"model,omitempty"`
	ScopeFile   string `json:"scope_file,omitempty"`
	File        string `json:"file,omitempty"`
	LocalID     string `json:"local_id,omitempty"`
	StateFile   string `json:"state_file,omitempty"`
	FromState   string `json:"from,omitempty"`
	On          string `json:"on,omitempty"`
	ToState     string `json:"to,omitempty"`
	Guard       string `json:"guard,omitempty"`
	Action      string `json:"action,omitempty"`
}

type Reference struct {
	Kind      string            `json:"kind"`
	Direction string            `json:"direction"`
	From      ReferenceEndpoint `json:"from"`
	To        ReferenceEndpoint `json:"to"`
}

type ParamSignature struct {
	Name  string `json:"name"`
	Model string `json:"model"`
	Doc   string `json:"doc,omitempty"`
}

type ReturnSignature struct {
	Name  string    `json:"name"`
	Model string    `json:"model"`
	Asset *AssetRef `json:"asset,omitempty"`
}

type EndpointSignature struct {
	Method   string `json:"method"`
	LeafPath string `json:"leaf_path"`
}

type FieldSignature struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	PK     bool   `json:"pk,omitempty"`
	FK     string `json:"fk,omitempty"`
	Unique bool   `json:"unique,omitempty"`
	Doc    string `json:"doc,omitempty"`
}

type Signature map[string]any

type SourceLocation struct {
	File      string `json:"file"`
	Line      int    `json:"line,omitempty"`
	Column    int    `json:"column,omitempty"`
	EndLine   int    `json:"end_line,omitempty"`
	EndColumn int    `json:"end_column,omitempty"`
}

type SourceSnippet struct {
	Language string `json:"language"`
	Text     string `json:"text"`
}

type GetSourceRequest struct {
	Selector Selector `json:"selector"`
	Fallback string   `json:"fallback,omitempty"`
}

type GetSourceResponse struct {
	Object      ObjectRef             `json:"object"`
	Source      SourceLocation        `json:"source"`
	Snippet     SourceSnippet         `json:"snippet"`
	Fallback    string                `json:"fallback,omitempty"`
	Diagnostics []semantic.Diagnostic `json:"diagnostics"`
}

type GetSignatureRequest struct {
	Selector Selector `json:"selector"`
}

type GetSignatureResponse struct {
	Object      ObjectRef             `json:"object"`
	Signature   Signature             `json:"signature"`
	Doc         string                `json:"doc,omitempty"`
	Diagnostics []semantic.Diagnostic `json:"diagnostics"`
}

type GetReferencesRequest struct {
	Selector  Selector `json:"selector"`
	Direction string   `json:"direction,omitempty"`
	Kinds     []string `json:"kinds,omitempty"`
}

type GetReferencesResponse struct {
	Object      ObjectRef             `json:"object"`
	Direction   string                `json:"direction"`
	Depth       int                   `json:"depth"`
	References  []Reference           `json:"references"`
	Diagnostics []semantic.Diagnostic `json:"diagnostics"`
}

type GetReferenceTreeRequest struct {
	Selector Selector `json:"selector"`
	Direction string   `json:"direction"`
	Depth     int      `json:"depth"`
	Kinds     []string `json:"kinds,omitempty"`
	MaxNodes  int      `json:"max_nodes,omitempty"`
	MaxEdges  int      `json:"max_edges,omitempty"`
}

type ReferenceTreeNode struct {
	Object ObjectRef `json:"object"`
	Depth  int       `json:"depth"`
	Via    []string  `json:"via"`
}

type ReferenceTreeEdge struct {
	Kind      string            `json:"kind"`
	Direction string            `json:"direction"`
	From      ReferenceEndpoint `json:"from"`
	To        ReferenceEndpoint `json:"to"`
	Depth     int               `json:"depth"`
}

type GetReferenceTreeResponse struct {
	Root             ObjectRef             `json:"root"`
	Direction        string                `json:"direction"`
	Depth            int                   `json:"depth"`
	Nodes            []ReferenceTreeNode   `json:"nodes"`
	Edges            []ReferenceTreeEdge   `json:"edges"`
	Truncated        bool                  `json:"truncated"`
	TruncatedReasons []string              `json:"truncated_reasons"`
	Diagnostics      []semantic.Diagnostic `json:"diagnostics"`
}

type ListObjectsRequest struct {
	Object string `json:"object,omitempty"`
	Kind   string `json:"kind,omitempty"`
	Module string `json:"module,omitempty"`
	File   string `json:"file,omitempty"`
}

type ListObjectsResponse struct {
	Objects     []ObjectRef           `json:"objects"`
	Diagnostics []semantic.Diagnostic `json:"diagnostics"`
}

type InspectRequest struct {
	Selector Selector `json:"selector"`
	Detail   string   `json:"detail,omitempty"`
}

type InspectResponse struct {
	Object      ObjectRef             `json:"object"`
	Signature   Signature             `json:"signature"`
	Doc         string                `json:"doc,omitempty"`
	Source      map[string]string     `json:"source,omitempty"`
	Members     map[string]any        `json:"members,omitempty"`
	References  []Reference           `json:"references,omitempty"`
	Diagnostics []semantic.Diagnostic `json:"diagnostics"`
}
