package designrecordsmcp

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/hiroshiasayadev-prog/brewprint/internal/designrecords"
)

type ToolsCallParams struct {
	Name      string          `json:"name"`
	Arguments json.RawMessage `json:"arguments,omitempty"`
}

type ToolsCallResult struct {
	Content []ToolContent `json:"content"`
	IsError bool          `json:"isError"`
}

type ToolContent struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type toolErrorEnvelope struct {
	Error designrecords.ToolError `json:"error"`
}

func (s *Server) HandleToolsCall(params json.RawMessage) ToolsCallResult {
	var call ToolsCallParams
	if err := decodeJSONObject(params, &call); err != nil {
		return toolErrorResult(designrecords.ErrorCodeInvalidRequest, fmt.Sprintf("invalid tools/call params: %v", err))
	}

	ctx := context.Background()
	idx, buildErr := s.buildIndex(ctx, s.cfg)

	if call.Name == "" {
		return toolErrorResult(designrecords.ErrorCodeInvalidRequest, "tool name is required")
	}

	switch call.Name {
	case "list_records":
		var req designrecords.ListRecordsRequest
		if err := decodeToolArguments(call.Arguments, &req); err != nil {
			return toolErrorResult(designrecords.ErrorCodeInvalidRequest, fmt.Sprintf("invalid list_records arguments: %v", err))
		}
		if buildErr != nil {
			return toolBuildIndexErrorResult(buildErr)
		}
		return toolHandlerResult(designrecords.ListRecords(ctx, idx, req))
	case "validate_records":
		var req designrecords.ValidateRecordsRequest
		if err := decodeToolArguments(call.Arguments, &req); err != nil {
			return toolErrorResult(designrecords.ErrorCodeInvalidRequest, fmt.Sprintf("invalid validate_records arguments: %v", err))
		}
		if buildErr != nil {
			return toolBuildIndexErrorResult(buildErr)
		}
		return toolHandlerResult(designrecords.ValidateRecords(ctx, idx, req))
	case "get_record":
		var req designrecords.GetRecordRequest
		if err := decodeToolArguments(call.Arguments, &req); err != nil {
			return toolErrorResult(designrecords.ErrorCodeInvalidRequest, fmt.Sprintf("invalid get_record arguments: %v", err))
		}
		if buildErr != nil {
			return toolBuildIndexErrorResult(buildErr)
		}
		return toolHandlerResult(designrecords.GetRecord(ctx, idx, req))
	case "get_records":
		var req designrecords.GetRecordsRequest
		if err := decodeToolArguments(call.Arguments, &req); err != nil {
			return toolErrorResult(designrecords.ErrorCodeInvalidRequest, fmt.Sprintf("invalid get_records arguments: %v", err))
		}
		if buildErr != nil {
			return toolBuildIndexErrorResult(buildErr)
		}
		return toolHandlerResult(designrecords.GetRecords(ctx, idx, req))
	case "list_authoring_guides":
		var req designrecords.ListAuthoringGuidesRequest
		if err := decodeToolArguments(call.Arguments, &req); err != nil {
			return toolErrorResult(designrecords.ErrorCodeInvalidRequest, fmt.Sprintf("invalid list_authoring_guides arguments: %v", err))
		}
		return toolHandlerResult(designrecords.ListAuthoringGuides(ctx, s.cfg, req))
	case "get_authoring_guidance":
		var req designrecords.GetAuthoringGuidanceRequest
		if err := decodeToolArguments(call.Arguments, &req); err != nil {
			return toolErrorResult(designrecords.ErrorCodeInvalidRequest, fmt.Sprintf("invalid get_authoring_guidance arguments: %v", err))
		}
		return toolHandlerResult(designrecords.GetAuthoringGuidance(ctx, s.cfg, req))
	case "resolve_reference":
		var req designrecords.ResolveReferenceRequest
		if err := decodeToolArguments(call.Arguments, &req); err != nil {
			return toolErrorResult(designrecords.ErrorCodeInvalidRequest, fmt.Sprintf("invalid resolve_reference arguments: %v", err))
		}
		if buildErr != nil {
			return toolBuildIndexErrorResult(buildErr)
		}
		return toolHandlerResult(designrecords.ResolveReference(ctx, idx, req))
	case "suggest_next_record":
		var req designrecords.SuggestNextRecordRequest
		if err := decodeToolArguments(call.Arguments, &req); err != nil {
			return toolErrorResult(designrecords.ErrorCodeInvalidRequest, fmt.Sprintf("invalid suggest_next_record arguments: %v", err))
		}
		if buildErr != nil {
			return toolBuildIndexErrorResult(buildErr)
		}
		return toolHandlerResult(designrecords.SuggestNextRecord(ctx, idx, req))
	case "propose_record_create":
		var req designrecords.ProposeRecordCreateRequest
		if err := decodeToolArguments(call.Arguments, &req); err != nil {
			return toolErrorResult(designrecords.ErrorCodeInvalidRequest, fmt.Sprintf("invalid propose_record_create arguments: %v", err))
		}
		if buildErr != nil {
			return toolBuildIndexErrorResult(buildErr)
		}
		return toolHandlerResult(designrecords.ProposeRecordCreate(ctx, s.cfg, idx, s.authoringStore, req))
	case "propose_record_update":
		var req designrecords.ProposeRecordUpdateRequest
		if err := decodeToolArguments(call.Arguments, &req); err != nil {
			return toolErrorResult(designrecords.ErrorCodeInvalidRequest, fmt.Sprintf("invalid propose_record_update arguments: %v", err))
		}
		if buildErr != nil {
			return toolBuildIndexErrorResult(buildErr)
		}
		return toolHandlerResult(designrecords.ProposeRecordUpdate(ctx, s.cfg, idx, s.authoringStore, req))
	case "get_proposed_write":
		var req designrecords.GetProposedWriteRequest
		if err := decodeToolArguments(call.Arguments, &req); err != nil {
			return toolErrorResult(designrecords.ErrorCodeInvalidRequest, fmt.Sprintf("invalid get_proposed_write arguments: %v", err))
		}
		return toolHandlerResult(designrecords.GetProposedWrite(ctx, s.authoringStore, req))
	case "accept_proposed_write":
		var req designrecords.AcceptProposedWriteRequest
		if err := decodeToolArguments(call.Arguments, &req); err != nil {
			return toolErrorResult(designrecords.ErrorCodeInvalidRequest, fmt.Sprintf("invalid accept_proposed_write arguments: %v", err))
		}
		if buildErr != nil {
			return toolBuildIndexErrorResult(buildErr)
		}
		return toolHandlerResult(designrecords.AcceptProposedWrite(ctx, s.cfg, idx, s.authoringStore, req))
	case "discard_proposed_write":
		var req designrecords.DiscardProposedWriteRequest
		if err := decodeToolArguments(call.Arguments, &req); err != nil {
			return toolErrorResult(designrecords.ErrorCodeInvalidRequest, fmt.Sprintf("invalid discard_proposed_write arguments: %v", err))
		}
		return toolHandlerResult(designrecords.DiscardProposedWrite(ctx, s.authoringStore, req))
	default:
		return toolErrorResult(designrecords.ErrorCodeInvalidRequest, fmt.Sprintf("unknown tool %q", call.Name))
	}
}

func toolBuildIndexErrorResult(err error) ToolsCallResult {
	return toolErrorResult(designrecords.ErrorCodeInvalidRequest, fmt.Sprintf("build design records index: %v", err))
}

func decodeToolArguments(raw json.RawMessage, out any) error {
	if len(raw) == 0 || bytes.Equal(bytes.TrimSpace(raw), []byte("null")) {
		raw = []byte("{}")
	}
	return decodeJSONObject(raw, out)
}

func decodeJSONObject(raw json.RawMessage, out any) error {
	decoder := json.NewDecoder(bytes.NewReader(raw))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(out); err != nil {
		return err
	}
	var extra struct{}
	if err := decoder.Decode(&extra); !errors.Is(err, io.EOF) {
		return fmt.Errorf("multiple JSON values")
	}
	return nil
}

func toolHandlerResult(data any, err error) ToolsCallResult {
	if err != nil {
		var toolErr *designrecords.ToolError
		if errors.As(err, &toolErr) {
			return toolErrorResult(toolErr.Code, toolErr.Message)
		}
		return toolErrorResult(designrecords.ErrorCodeInvalidRequest, err.Error())
	}
	text, err := jsonText(data)
	if err != nil {
		return toolErrorResult(designrecords.ErrorCodeInvalidRequest, fmt.Sprintf("encode tool response: %v", err))
	}
	return ToolsCallResult{
		Content: []ToolContent{{Type: "text", Text: text}},
		IsError: false,
	}
}

func toolErrorResult(code designrecords.ErrorCode, message string) ToolsCallResult {
	text, err := jsonText(toolErrorEnvelope{Error: designrecords.ToolError{Code: code, Message: message}})
	if err != nil {
		text = `{"error":{"code":"invalid_request","message":"failed to encode tool error"}}`
	}
	return ToolsCallResult{
		Content: []ToolContent{{Type: "text", Text: text}},
		IsError: true,
	}
}

func jsonText(v any) (string, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
