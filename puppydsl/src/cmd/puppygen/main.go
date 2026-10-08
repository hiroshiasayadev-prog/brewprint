package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hiroshiasayadev-prog/brewprint/puppydsl/src/generator"
)

func main() {
	root := flag.String("root", "", "PuppyDSL application root containing .puppy.yaml and .manifest.yaml sources")
	flag.Parse()

	if *root == "" {
		fmt.Fprintln(os.Stderr, "error: -root is required")
		os.Exit(2)
	}
	if err := generator.Generate(*root); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("generated %s/go_native/puppygen\n", *root)
}
