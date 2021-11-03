// Command k6-to-honeycomb streams k6 load test output into a Honeycomb dataset.
//
// This program is meant to be used to operate over the streamed output produced by the
// https://k6.io/ load testing tool.
//
// Configuration
//
// The Honeycomb write API key can be provided with the `-k` command line flag. If not provided it will attempt to be read from the `HC_API_KEY` environment variable.
//
// Example Usage
//
// Here is an exmaple of running k6 with streaming JSON output that this program can ingest and ship off to honeycomb.
//
// Example k6 invocation (note --out json= parameter):
//
// $ k6 run -u 1 -d 10s --out json=./output.json
//
// You can then ship these events to honeycomb like so:
//
// $ k6-to-honeycomb output.json
//
// By default this will send to a `k6-load-tests` dataset but this is configurable with the -d parameter like so:
//
// $ k6-to-honeycomb -d k6-load-test-results output.json
//
package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	flagDataset = flag.String("d", "k6-load-tests", "If provided, overrides the name of the honeycomb dataset.")
	flagAPIKey  = flag.String("k", "", "The Honeycomb write API key to use (will read from HC_API_KEY if not provided).")
)

func main() {
	flag.Parse()
	fmt.Println("k6-to-honeycomb")
	fmt.Println("dataset:", *flagDataset)
	fmt.Println("api key:", *flagAPIKey)

	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// run processes the provided input file(s) and ships them to the configured honeycomb target.
func run(inputs []string) error {
	parsed, err := parseK6Results(inputs)
	if err != nil {
		return fmt.Errorf("issue parsing k6 results: %w", err)
	}
	fmt.Println(len(parsed), "results")

	if *flagAPIKey == "" {
		*flagAPIKey = os.Getenv("HC_API_KEY")
	}

	if err := shipToHC(*flagAPIKey, *flagDataset, parsed); err != nil {
		return fmt.Errorf("issue shipping events to Honeycomb: %w", err)
	}
	return nil
}
