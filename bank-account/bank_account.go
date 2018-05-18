package account

import (
	"sync"
)

// Account 1
type Account struct {
	balance int64
	open    bool
	mux     sync.Mutex
}

// Open 1
func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	newAccount := Account{balance: initialDeposit, open: true}
	return &newAccount
}

// Close 1
func (a *Account) Close() (int64, bool) {
	a.mux.Lock()
	defer a.mux.Unlock()
	if a.open {
		a.open = false
		return a.balance, true
	}
	return 0, false
}

// Balance 1
func (a *Account) Balance() (int64, bool) {
	if a.open {
		return a.balance, true
	}
	return 0, false
}

// Deposit 1
func (a *Account) Deposit(amount int64) (int64, bool) {
	a.mux.Lock()
	defer a.mux.Unlock()
	if !a.open {
		return 0, false
	}
	newBalance := a.balance + amount
	if amount < 0 {
		if newBalance < 0 {
			return 0, false
		}
		a.balance = newBalance
		return a.balance, true
	}
	a.balance = newBalance
	return a.balance, true
}
