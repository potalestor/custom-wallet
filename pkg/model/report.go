package model

import (
	"encoding/csv"
	"io"
	"time"
)

type Reports []*Report

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

type Report struct {
	Operation Operation
	Created   time.Time
	Amount    USD
}

func (r *Report) Strings() []string {
	return []string{
		r.Operation.String(),
		r.Created.Format(time.RFC3339),
		r.Amount.String(),
	}
}
