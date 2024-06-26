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

func unique(data []float64) float64 {
	uniq := make(map[float64]bool)
	for _, d := range data {
		uniq[d] = true
	}

	return float64(len(uniq))
}

func length(data []float64) float64 {
	return float64(len(data))
}

func avg(data []float64) float64 {
	return sum(data) / float64(len(data))
}

func min(data []float64) float64 {
	m := data[0]

	for i := 1; i < len(data); i++ {
		if data[i] < m {
			m = data[i]
		}
	}

	return m
}

func max(data []float64) float64 {
	m := data[0]

	for i := 1; i < len(data); i++ {
		if data[i] > m {
			m = data[i]
		}
	}

	return m
}

func csv2float(r io.Reader, column int) ([]float64, error) {
	// Create the CSV Reader used to read in data from CSV files
	cr := csv.NewReader(r)
	cr.ReuseRecord = true

	// Adjusting for 0 based index
	column--

	var data []float64
	for i := 0; ; i++ {
		row, err := cr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("cannot read data from file: %w", err)
		}

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
