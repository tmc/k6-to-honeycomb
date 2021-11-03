package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	k6tohoneycomb "github.com/tmc/k6-to-honeycomb"
)

func parseK6Results(inputs []string) ([]k6tohoneycomb.K6DataPoint, error) {
	var results []k6tohoneycomb.K6DataPoint

	for _, i := range inputs {
		r, err := parseK6Result(i)
		if err != nil {
			return results, fmt.Errorf("issue parsing k6 results in '%v': %w", i, err)
		}
		results = append(results, r...)
	}
	return results, nil
}

func parseK6Result(input string) ([]k6tohoneycomb.K6DataPoint, error) {
	var results []k6tohoneycomb.K6DataPoint
	// TODO: marshal each line of the file into a K6DataPoint.
	// TODO: consider allowing partial success.
	// TODO: wrap errors.
	f, err := os.Open(input)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		result, err := unmarshalK6Result(scanner.Bytes())
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	return results, nil
}

func unmarshalK6Result(in []byte) (k6tohoneycomb.K6DataPoint, error) {
	var result k6tohoneycomb.K6DataPoint
	return result, json.Unmarshal(in, &result)
}
