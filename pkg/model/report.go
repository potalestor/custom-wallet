package model

import (
	"encoding/csv"
	"io"
	"time"
)

// Reports array of report.
type Reports []*Report

// CSV convert to CSV.
func (r Reports) CSV(w io.Writer) error {
	writer := csv.NewWriter(w)
	defer writer.Flush()

	for _, report := range r {
		if err := writer.Write(report.Strings()); err != nil {
			return err
		}
	}

	return nil
}

// Report from transaction.
type Report struct {
	Operation Operation
	Created   time.Time
	Amount    USD
}

// Strings converts to []string.
func (r *Report) Strings() []string {
	return []string{
		r.Operation.String(),
		r.Created.Format(time.RFC3339),
		r.Amount.String(),
	}
}
