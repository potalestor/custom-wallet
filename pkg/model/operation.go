package model

// Operation with the wallet account.
type Operation int

const (
	// Unsupported is undefined operation.
	Unsupported Operation = iota
	// Deposit money.
	Deposit
	// Withdraw money.
	Withdraw
	// Both.
	Both
)

func (o Operation) String() string {
	switch o {
	case Unsupported:
		return "Unsupported"
	case Deposit:
		return "Deposit"
	case Withdraw:
		return "Withdraw"
	case Both:
		return "Both"
	default:
		return "Unsupported"
	}
}
