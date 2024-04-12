package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	// Verify and parse arguments
	op := flag.String("op", "", "Operation to be executed.")
	column := flag.Int("col", 1, "CSV column on which to execute operation.")
	flag.Parse()

	if err := run(flag.Args(), *op, *column, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(filenames []string, op string, column int, out io.Writer) error {
	var opFunc statsFunc

	// Validations
	if len(filenames) == 0 {
		return ErrNoFiles
	}

	if column < 1 {
		return fmt.Errorf("%w: %d", ErrInvalidColumn, column)
	}

	switch op {
	case "sum":
		opFunc = sum
	case "avg":
		opFunc = avg
	default:
		return fmt.Errorf("%w: %s", ErrInvalidOperation, op)
	}

	consolidate := make([]float64, 0)
	for _, fname := range filenames {
		// Open the file for reading
		f, err := os.Open(fname)
		if err != nil {
			return fmt.Errorf("cannot open file: %w", err)
		}

		// Parse the CSV into a slice of float64 numbers
		data, err := csv2float(f, column)
		if err != nil {
			return err
		}

		if err = f.Close(); err != nil {
			return err
		}

		// Append the data to consolidate
		consolidate = append(consolidate, data...)
	}

	_, err := fmt.Fprintln(out, opFunc(consolidate))
	if err != nil {
		return err
	}

	return nil
}
