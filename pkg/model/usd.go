package model

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// USD is a currency US.
type USD float64

const sUSD = `^\$?([1-9]{1}[0-9]{0,2}(\,[0-9]{3})*` +
	`(\.[0-9]{0,2})?|[1-9]{1}[0-9]{0,}(\.[0-9]{0,2})?` +
	`|0(\.[0-9]{0,2})?|(\.[0-9]{1,2})?)$`

var (
	reUSD            = regexp.MustCompile(sUSD)
	ErrNegativeValue = errors.New(`the USD can't be less than zero`)
)

// ErrInvalidUSDFormat is invalid USD format.
type ErrInvalidUSDFormat struct {
	format string
}

// NewErrInvalidUSDFormat returns new ErrInvalidUSDFormat instance.
func NewErrInvalidUSDFormat(format string) *ErrInvalidUSDFormat {
	return &ErrInvalidUSDFormat{format: format}
}

func (e *ErrInvalidUSDFormat) Error() string {
	return fmt.Sprintf("invalid USD format: %s. Required: $12.34 or $1,234.56", e.format)
}

// parse a string that can be a USD.
func (u *USD) Parse(s string) error {
	rawUSD := reUSD.FindStringSubmatch(s)
	if len(rawUSD) < 1 {
		return NewErrInvalidUSDFormat(s)
	}

	floatUSD, err := strconv.ParseFloat(strings.ReplaceAll(rawUSD[1], ",", ""), 64)
	if err != nil {
		return errors.Wrap(NewErrInvalidUSDFormat(s), err.Error())
	}

	*u = USD(floatUSD)

	return u.Validate()
}

// Validate USD.
func (u USD) Validate() error {
	if u < 0 {
		return ErrNegativeValue
	}

	return nil
}

func (u USD) String() string {
	return fmt.Sprintf("%.2f", u)
}

// Float64 returns USD as float64.
func (u USD) Float64() float64 {
	return float64(u)
}
