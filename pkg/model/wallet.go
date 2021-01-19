package model

import "fmt"

// Wallet is used to store money.
type Wallet struct {
	ID      uint64
	Name    string
	Account USD
}

func (w *Wallet) String() string {
	return fmt.Sprintf(
		"id=%v name=%v account=%v",
		w.ID,
		w.Name,
		w.Account,
	)
}
