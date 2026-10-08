package mcp

import (
	puppy_mcp "github.com/hiroshiasayadev-prog/brewprint/drmcp/puppydsl/go_native/puppygen/drmcp_runtime/mcp"
	"github.com/hiroshiasayadev-prog/brewprint/puppydsl/src/native"
)

func Serve(call native.Call) error {
	return nil
}

var serveBinding = puppy_mcp.BindServe(Serve)

// Bindings returns the native bindings owned by the DRMCP MCP package.
func Bindings() []native.Binding {
	return []native.Binding{serveBinding}
}
