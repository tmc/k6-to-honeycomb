package main

import (
	"fmt"

	libhoney "github.com/honeycombio/libhoney-go"
	k6tohoneycomb "github.com/tmc/k6-to-honeycomb"
)

func shipToHC(hcAPIKey string, hcDataset string, results []k6tohoneycomb.K6DataPoint) error {
	if hcAPIKey == "" {
		return fmt.Errorf("missing Honeycomb write API key")
	}
	libhoney.Init(libhoney.Config{
		WriteKey: hcAPIKey,
		Dataset:  hcDataset,
	})
	defer libhoney.Close()

	for _, dp := range results {
		e := libhoney.NewEvent()
		e.Add(dp)
		e.Timestamp = dp.Data.Time

		// TODO: wrap and contextualize errors
		if err := e.Send(); err != nil {
			return err
		}
	}
	return nil
}
