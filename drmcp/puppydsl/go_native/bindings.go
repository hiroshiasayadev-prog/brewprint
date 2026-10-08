package gonative

import (
	drmcpmcp "github.com/hiroshiasayadev-prog/brewprint/drmcp/puppydsl/go_native/drmcp_runtime/mcp"
	"github.com/hiroshiasayadev-prog/brewprint/puppydsl/src/native"
)

// Bindings returns every native binding compiled into the DRMCP PuppyDSL runtime.
func Bindings() []native.Binding {
	return drmcpmcp.Bindings()
}

// NewRegistry validates and indexes every native binding compiled into DRMCP.
func NewRegistry() (*native.Registry, error) {
	return native.NewRegistry(Bindings())
}
