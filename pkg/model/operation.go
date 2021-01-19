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
)
