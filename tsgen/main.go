package main

import (
	"flag"
	"fmt"
	"os"

	dsl "school/project/dslgen"
)

func main() {
	dslFile := flag.String("dsl", "", "Path to DSL file (required)")
	outFile := flag.String("out", "", "Output .go file")
	flag.Parse()

	if *dslFile == "" {
		fmt.Fprintln(os.Stderr, "error: -dsl is required")
		flag.Usage()
		os.Exit(2)
	}

	src, err := os.ReadFile(*dslFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "read error %s: %v\n", *dslFile, err)
		os.Exit(1)
	}

	generated, err := dsl.GenerateFromDSL(string(src))
	if err != nil {
		fmt.Fprintf(os.Stderr, "generate error: %v\n", err)
		os.Exit(1)
	}

	if *outFile != "" && *outFile != generated {
		if err := os.Rename(generated, *outFile); err != nil {
			fmt.Fprintf(os.Stderr, "rename %s→%s: %v\n", generated, *outFile, err)
			os.Exit(1)
		}
		fmt.Println(*outFile)
		return
	}

	fmt.Println(generated)
}
