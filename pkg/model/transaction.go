package model

import "time"

// Transaction is a wallet operation.
type Transaction struct {
	ID        uint64
	Created   time.Time
	Operation Operation
	Wallet    uint64
	Amount    USD
}
