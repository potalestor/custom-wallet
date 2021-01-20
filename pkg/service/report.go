package service

import (
	"context"

	"github.com/potalestor/custom-wallet/pkg/model"
	"github.com/potalestor/custom-wallet/pkg/repo"
)

type Report struct {
	repository repo.Repository
}

func NewReport(repository repo.Repository) *Report {
	return &Report{repository: repository}
}

func (r *Report) Report(filter *model.Filter) (model.Reports, error) {
	return r.repository.Report(context.Background(), filter)
}
