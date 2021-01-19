package service

import "github.com/potalestor/custom-wallet/pkg/repo"

type Report struct {
	repository repo.Repository
}

func NewReport(repository repo.Repository) *Report {
	return &Report{repository: repository}
}

// func (r *Report) Report(src, dst string, amount model.USD) error {

// }
