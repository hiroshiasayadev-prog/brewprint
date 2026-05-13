package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/hiroshiasayadev-prog/brewprint/internal/designrecords"
	"github.com/hiroshiasayadev-prog/brewprint/internal/designrecordsmcp"
)

func main() {
	if err := run(os.Args[1:], os.Stdin, os.Stdout, os.Stderr); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(args []string, stdin io.Reader, stdout io.Writer, stderr io.Writer) error {
	flags := flag.NewFlagSet("design-records-mcp", flag.ContinueOnError)
	flags.SetOutput(stderr)
	root := flags.String("root", "", "repository root; defaults to current working directory")
	summary := flags.Bool("summary", false, "print a human-readable index summary and exit")
	if err := flags.Parse(args); err != nil {
		return err
	}

	cfg, err := designrecords.NewConfig(*root)
	if err != nil {
		return err
	}
	if !*summary {
		return designrecordsmcp.NewServer(cfg).ServeJSONRPCLines(stdin, stdout)
	}

	idx, err := designrecords.BuildIndex(context.Background(), cfg)
	if err != nil {
		return err
	}

	fmt.Fprintln(stdout, "design-records-mcp ready")
	fmt.Fprintf(stdout, "root: %s\n", idx.Root)
	fmt.Fprintf(stdout, "records: %d\n", len(idx.Records))
	return nil
}
