package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
)

// statsFunc defines a generic statistical function
type statsFunc func(data []float64) float64

func sum(data []float64) float64 {
	s := 0.0

	for _, v := range data {
		s += v
	}

	return s
}

func avg(data []float64) float64 {
	return sum(data) / float64(len(data))
}

func csv2float(r io.Reader, column int) ([]float64, error) {
	// Create the CSV Reader used to read in data from CSV files
	cr := csv.NewReader(r)

	// Adjusting for 0 based index
	column--

	// Read all CSV data
	allData, err := cr.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("cannot read data from file: %w", err)
	}

	var data []float64
	for i, row := range allData {
		if i == 0 {
			continue
		}

		// Checking number of columns in CSV file
		if len(row) <= column {
			return nil, fmt.Errorf("%w: file has only %d columns", ErrInvalidColumn, len(row))
		}

		// Try to convert data into a float number
		v, err := strconv.ParseFloat(row[column], 64)
		if err != nil {
			return nil, fmt.Errorf("%w: %s", ErrNotNumber, err)
		}

		data = append(data, v)
	}

	return data, nil
}
