package designrecords

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	adrIDPattern              = regexp.MustCompile(`^ADR-(\d{3})$`)
	requirementRangeIDPattern = regexp.MustCompile(`^REQ-([A-Z0-9](?:[A-Z0-9-]*[A-Z0-9])?)-(\d{3})$`)
	workItemRangeIDPattern   = regexp.MustCompile(`^WORK-([A-Z0-9](?:[A-Z0-9-]*[A-Z0-9])?)-(\d{3})$`)
	taskRangeIDPattern       = regexp.MustCompile(`^TASK-([A-Z0-9](?:[A-Z0-9-]*[A-Z0-9])?)-(\d{3})-(\d{2})$`)
)

type recordIDRange struct {
	kind     RecordKind
	domain   string
	hasScope bool
	workSeq  int
	hasWork  bool
	from     int
	hasFrom  bool
	to       int
	hasTo    bool
}

type recordIDEndpoint struct {
	kind    RecordKind
	domain  string
	workSeq int
	hasWork bool
	seq     int
}

func parseRecordIDRange(idRange IDRange, explicitKind RecordKind) (*recordIDRange, error) {
	parsed := &recordIDRange{}
	var fromEndpoint, toEndpoint *recordIDEndpoint

	if strings.TrimSpace(idRange.From) != "" {
		endpoint, err := parseRecordIDRangeEndpoint(idRange.From)
		if err != nil {
			return nil, err
		}
		fromEndpoint = endpoint
	}
	if strings.TrimSpace(idRange.To) != "" {
		endpoint, err := parseRecordIDRangeEndpoint(idRange.To)
		if err != nil {
			return nil, err
		}
		toEndpoint = endpoint
	}

	if fromEndpoint == nil && toEndpoint == nil {
		kind := explicitKind
		if kind == "" {
			kind = RecordKindDecision
		}
		if !idRangeKindSupportsRange(kind) {
			return nil, newToolError(ErrorCodeInvalidIDRange, fmt.Sprintf("id_range does not support kind %q", kind))
		}
		parsed.kind = kind
		return parsed, nil
	}

	reference := fromEndpoint
	if reference == nil {
		reference = toEndpoint
	}
	if explicitKind != "" && explicitKind != reference.kind {
		return nil, newToolError(ErrorCodeInvalidIDRange, fmt.Sprintf("id_range endpoint kind %s does not match kind %s", reference.kind, explicitKind))
	}
	parsed.kind = reference.kind
	parsed.domain = reference.domain
	parsed.hasScope = reference.domain != "" || reference.hasWork
	parsed.workSeq = reference.workSeq
	parsed.hasWork = reference.hasWork

	for _, endpoint := range []*recordIDEndpoint{fromEndpoint, toEndpoint} {
		if endpoint == nil {
			continue
		}
		if endpoint.kind != parsed.kind {
			return nil, newToolError(ErrorCodeInvalidIDRange, "id_range endpoints must use the same ID family")
		}
		if endpoint.domain != parsed.domain {
			return nil, newToolError(ErrorCodeInvalidIDRange, "workflow id_range endpoints must use the same domain")
		}
		if endpoint.hasWork != parsed.hasWork || endpoint.workSeq != parsed.workSeq {
			return nil, newToolError(ErrorCodeInvalidIDRange, "task id_range endpoints must use the same work sequence")
		}
	}
	if fromEndpoint != nil {
		parsed.from = fromEndpoint.seq
		parsed.hasFrom = true
	}
	if toEndpoint != nil {
		parsed.to = toEndpoint.seq
		parsed.hasTo = true
	}
	return parsed, nil
}

func parseRecordIDRangeEndpoint(value string) (*recordIDEndpoint, error) {
	normalized := normalizeRecordID(value)
	if match := adrIDPattern.FindStringSubmatch(normalized); match != nil {
		seq, err := atoiIDPart(value, match[1])
		if err != nil {
			return nil, err
		}
		return &recordIDEndpoint{kind: RecordKindDecision, seq: seq}, nil
	}
	if match := requirementRangeIDPattern.FindStringSubmatch(normalized); match != nil {
		seq, err := atoiIDPart(value, match[2])
		if err != nil {
			return nil, err
		}
		return &recordIDEndpoint{kind: RecordKindRequirement, domain: match[1], seq: seq}, nil
	}
	if match := workItemRangeIDPattern.FindStringSubmatch(normalized); match != nil {
		seq, err := atoiIDPart(value, match[2])
		if err != nil {
			return nil, err
		}
		return &recordIDEndpoint{kind: RecordKindWorkItem, domain: match[1], seq: seq}, nil
	}
	if match := taskRangeIDPattern.FindStringSubmatch(normalized); match != nil {
		workSeq, err := atoiIDPart(value, match[2])
		if err != nil {
			return nil, err
		}
		seq, err := atoiIDPart(value, match[3])
		if err != nil {
			return nil, err
		}
		return &recordIDEndpoint{kind: RecordKindTask, domain: match[1], workSeq: workSeq, hasWork: true, seq: seq}, nil
	}
	return nil, newToolError(ErrorCodeInvalidIDRange, fmt.Sprintf("invalid id_range endpoint %q", value))
}

func atoiIDPart(original string, value string) (int, error) {
	num, err := strconv.Atoi(value)
	if err != nil {
		return 0, newToolError(ErrorCodeInvalidIDRange, fmt.Sprintf("invalid id_range endpoint %q", original))
	}
	return num, nil
}

func idRangeKindSupportsRange(kind RecordKind) bool {
	switch kind {
	case RecordKindDecision, RecordKindRequirement, RecordKindWorkItem, RecordKindTask:
		return true
	default:
		return false
	}
}

func (r recordIDRange) containsRecord(record Record) bool {
	return r.containsID(record.Kind, record.ID)
}

func (r recordIDRange) containsID(kind RecordKind, id string) bool {
	if kind != r.kind {
		return false
	}
	endpoint, err := parseRecordIDRangeEndpoint(id)
	if err != nil {
		return false
	}
	if endpoint.kind != r.kind {
		return false
	}
	if r.hasScope && endpoint.domain != r.domain {
		return false
	}
	if r.hasScope && (endpoint.hasWork != r.hasWork || endpoint.workSeq != r.workSeq) {
		return false
	}
	return r.contains(endpoint.seq)
}

func (r recordIDRange) contains(num int) bool {
	if r.hasFrom && num < r.from {
		return false
	}
	if r.hasTo && num > r.to {
		return false
	}
	return true
}

func decisionRecordNumber(id string) (int, bool) {
	match := adrIDPattern.FindStringSubmatch(normalizeRecordID(id))
	if match == nil {
		return 0, false
	}
	num, err := strconv.Atoi(match[1])
	if err != nil {
		return 0, false
	}
	return num, true
}
