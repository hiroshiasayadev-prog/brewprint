package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/hiroshiasayadev-prog/brewprint/internal/mcp"
	"github.com/hiroshiasayadev-prog/brewprint/internal/query"
	"github.com/hiroshiasayadev-prog/brewprint/internal/resolve"
	"github.com/hiroshiasayadev-prog/brewprint/internal/semantic"
	"github.com/hiroshiasayadev-prog/brewprint/internal/source"
)

func main() {
	if err := run(os.Args[1:], os.Stdin, os.Stdout, os.Stderr); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(args []string, stdin io.Reader, stdout io.Writer, stderr io.Writer) error {
	if len(args) == 0 {
		return fmt.Errorf("%s", rootUsage())
	}
	switch args[0] {
	case "mcp":
		return runMCP(args[1:], stdin, stdout)
	case "validate":
		return runValidate(args[1:], stdout)
	default:
		return fmt.Errorf("unknown command: %s\n\n%s", args[0], rootUsage())
	}
}

func rootUsage() string {
	return "usage:\n  brewprint mcp --yaml-root <path>\n  brewprint validate --yaml-root <path> [--format text|json]"
}

func runMCP(args []string, stdin io.Reader, stdout io.Writer) error {
	flags := flag.NewFlagSet("mcp", flag.ContinueOnError)
	yamlRoot := flags.String("yaml-root", "", "path to brewprint yaml root")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if *yamlRoot == "" {
		return fmt.Errorf("--yaml-root is required\n\nusage: brewprint mcp --yaml-root <path>")
	}

	server, err := newMCPServer(*yamlRoot)
	if err != nil {
		return err
	}
	return server.ServeJSONRPCLines(stdin, stdout)
}

func runValidate(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("validate", flag.ContinueOnError)
	yamlRoot := flags.String("yaml-root", "", "path to brewprint yaml root")
	format := flags.String("format", "text", "output format: text or json")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if *yamlRoot == "" {
		return fmt.Errorf("--yaml-root is required\n\nusage: brewprint validate --yaml-root <path> [--format text|json]")
	}
	if *format != "text" && *format != "json" {
		return fmt.Errorf("unsupported validate format: %s", *format)
	}

	_, diagnostics, err := loadProject(*yamlRoot)
	if err != nil {
		return err
	}
	errors, warnings := countDiagnostics(diagnostics)
	if err := writeDiagnostics(stdout, diagnostics, errors, warnings, *format); err != nil {
		return err
	}
	if errors > 0 {
		return fmt.Errorf("validation failed: %d error(s), %d warning(s)", errors, warnings)
	}
	return nil
}

func newMCPServer(yamlRoot string) (*mcp.Server, error) {
	project, diagnostics, err := loadProject(yamlRoot)
	if err != nil {
		return nil, err
	}
	if diagnostic, ok := firstErrorDiagnostic(diagnostics); ok {
		return nil, fmt.Errorf("semantic diagnostic: %s", formatDiagnostic(diagnostic))
	}
	return mcp.NewServer(query.NewService(project)), nil
}

func loadProject(yamlRoot string) (*semantic.Project, []semantic.Diagnostic, error) {
	loader := source.Loader{}
	raw, err := loader.Load(yamlRoot)
	if err != nil {
		return nil, nil, fmt.Errorf("load yaml root: %w", err)
	}
	project, diagnostics := resolve.Build(raw)
	return project, diagnostics, nil
}

type validateOutput struct {
	Diagnostics  []semantic.Diagnostic `json:"diagnostics"`
	ErrorCount   int                   `json:"error_count"`
	WarningCount int                   `json:"warning_count"`
}

func writeDiagnostics(out io.Writer, diagnostics []semantic.Diagnostic, errors int, warnings int, format string) error {
	switch format {
	case "text":
		writeTextDiagnostics(out, diagnostics)
		return nil
	case "json":
		return json.NewEncoder(out).Encode(validateOutput{
			Diagnostics:  diagnostics,
			ErrorCount:   errors,
			WarningCount: warnings,
		})
	default:
		return fmt.Errorf("unsupported validate format: %s", format)
	}
}

func writeTextDiagnostics(out io.Writer, diagnostics []semantic.Diagnostic) {
	if len(diagnostics) == 0 {
		fmt.Fprintln(out, "ok")
		return
	}
	for _, diagnostic := range diagnostics {
		fmt.Fprintln(out, formatDiagnostic(diagnostic))
	}
}

func countDiagnostics(diagnostics []semantic.Diagnostic) (errors int, warnings int) {
	for _, diagnostic := range diagnostics {
		switch diagnostic.Severity {
		case semantic.SeverityError:
			errors++
		case semantic.SeverityWarning:
			warnings++
		}
	}
	return errors, warnings
}

func firstErrorDiagnostic(diagnostics []semantic.Diagnostic) (semantic.Diagnostic, bool) {
	for _, diagnostic := range diagnostics {
		if diagnostic.Severity == semantic.SeverityError {
			return diagnostic, true
		}
	}
	return semantic.Diagnostic{}, false
}

func formatDiagnostic(diagnostic semantic.Diagnostic) string {
	prefix := string(diagnostic.Severity)
	if diagnostic.Code != "" {
		prefix += " " + diagnostic.Code
	}
	if diagnostic.FileID != "" {
		return fmt.Sprintf("%s %s: %s", prefix, diagnostic.FileID, diagnostic.Message)
	}
	return fmt.Sprintf("%s: %s", prefix, diagnostic.Message)
}
