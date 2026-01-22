// Package model
package model

import "sync"

type BankModel struct {
	Mu      sync.Mutex
	balance int
}

func (b *BankModel) GetBalance() int {

	b.Mu.Lock()
	defer b.Mu.Unlock()

	return b.balance
}
