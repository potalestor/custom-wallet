package model

import (
	"fmt"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/pkg/errors"
)

var (
	ErrInvalidDataRange = errors.New("invalid data range")

	tenYearsAgo = -24 * 365 * 10 * time.Hour
	tomorrow    = 24 * time.Hour
)

const formatFilter = `wallet=%s, operation=%v, range=%s-%s`

// Filter for report.
type Filter struct {
	WalletName string
	Operation  Operation
	DateRange  [2]time.Time
}

// NewFilter returns new instance.
func NewFilter() *Filter {
	f := &Filter{Operation: Both}
	f.DateRange[0] = time.Now().Add(tenYearsAgo)
	f.DateRange[1] = time.Now().Add(tomorrow)

	return f
}

// Validate returns error if filter incorrect.
func (f Filter) Validate() error {
	return validation.ValidateStruct(&f,
		validation.Field(&f.WalletName, validation.Required, is.Alphanumeric),
		validation.Field(&f.Operation, validation.Required),
		validation.Field(&f.DateRange,
			validation.By(
				func(value interface{}) error {
					r, _ := value.([2]time.Time)

					if r[0].After(r[1]) {
						return errors.Wrap(ErrInvalidDataRange, fmt.Sprintf("%v-%v", r[0], r[1]))
					}

					return nil
				})))
}

func (f *Filter) String() string {
	return fmt.Sprintf(formatFilter,
		f.WalletName,
		f.Operation.String(),
		f.DateRange[0].Format("02.01.2006 15:04:05"),
		f.DateRange[1].Format("02.01.2006 15:04:05"),
	)
}
