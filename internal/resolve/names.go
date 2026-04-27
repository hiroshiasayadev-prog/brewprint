package resolve

import (
	"path"
	"strings"

	"github.com/hiroshiasayadev-prog/brewprint/internal/semantic"
)

var nodeKindSentinels = map[string]struct{}{
	"task":   {},
	"asset":  {},
	"store":  {},
	"event":  {},
	"state":  {},
	"branch": {},
	"fork":   {},
	"join":   {},
	"model":  {},
}

func moduleForFileID(fileID semantic.FileID) string {
	parts := strings.Split(fileID.String(), "/")
	if len(parts) == 0 {
		return ""
	}
	parts[len(parts)-1] = strings.TrimSuffix(parts[len(parts)-1], path.Ext(parts[len(parts)-1]))
	for i, part := range parts[:len(parts)-1] {
		if _, ok := nodeKindSentinels[part]; ok {
			return strings.Join(parts[:i], ".")
		}
	}
	if len(parts) == 1 {
		return ""
	}
	return strings.Join(parts[:len(parts)-1], ".")
}

func qidFor(module, kind, id string) semantic.QualifiedID {
	if module == "" {
		return semantic.QualifiedID(kind + "." + id)
	}
	return semantic.QualifiedID(module + "." + kind + "." + id)
}

func actorQID(id string) semantic.QualifiedID {
	return semantic.QualifiedID("actor." + id)
}

func localStoreQID(taskQID semantic.QualifiedID, name string) semantic.QualifiedID {
	return semantic.QualifiedID(taskQID.String() + ".store." + name)
}

func isFullQID(ref string, kind string) bool {
	parts := strings.Split(ref, ".")
	for _, part := range parts {
		if part == kind {
			return true
		}
	}
	return false
}

func resolveModelQID(module, ref string) semantic.QualifiedID {
	if ref == "" {
		return ""
	}
	if isFullQID(ref, "model") {
		return semantic.QualifiedID(ref)
	}
	return qidFor(module, "model", ref)
}

func resolveSameModuleStoreQID(module, ref string) semantic.QualifiedID {
	if ref == "" {
		return ""
	}
	if isFullQID(ref, "store") {
		return semantic.QualifiedID(ref)
	}
	return qidFor(module, "store", ref)
}

func shortName(qid semantic.QualifiedID) string {
	parts := strings.Split(qid.String(), ".")
	if len(parts) == 0 {
		return qid.String()
	}
	return parts[len(parts)-1]
}
