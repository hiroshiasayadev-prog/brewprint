package query

import "github.com/hiroshiasayadev-prog/brewprint/internal/semantic"

type Selector struct {
	ID string `json:"id"`
}

type ObjectRef struct {
	Object      string `json:"object"`
	Kind        string `json:"kind"`
	ID          string `json:"id"`
	QualifiedID string `json:"qualified_id,omitempty"`
	Label       string `json:"label,omitempty"`
	File        string `json:"file,omitempty"`
	LocalID     string `json:"local_id,omitempty"`
}

type AssetRef struct {
	Object    string `json:"object"`
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
