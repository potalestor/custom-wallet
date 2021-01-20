package service

import (
	"context"

	"github.com/potalestor/custom-wallet/pkg/model"
	"github.com/potalestor/custom-wallet/pkg/repo"
)

// Report service.
type Report struct {
	repository repo.Repository
}

// NewReport returns new instance.
func NewReport(repository repo.Repository) *Report {
	return &Report{repository: repository}
}

// Report use case.
func (r *Report) Report(filter *model.Filter) (model.Reports, error) {
	return r.repository.Report(context.Background(), filter)
}
