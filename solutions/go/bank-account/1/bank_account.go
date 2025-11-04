package account

import "sync"

type Account struct {
    mu			sync.Mutex
    Bal 		int64
    IsClosed	bool
}

func Open(amount int64) *Account {
	if amount < 0 { return nil }
    a := &Account{
        Bal: amount,
        IsClosed: false,
    }
    return a
}

func (a *Account) Balance() (int64, bool) {
    a.mu.Lock()
    defer a.mu.Unlock()
    
	if a.IsClosed { return 0, false }
    return a.Bal, true
}

func (a *Account) Deposit(amount int64) (int64, bool) {
    a.mu.Lock()
    defer a.mu.Unlock()
    
    if a.IsClosed {
        return 0, false
    }
    if a.Bal+amount < 0 {
        return 0, false
    }
    a.Bal += amount
    return a.Bal, true
}

func (a *Account) Close() (int64, bool) {
    a.mu.Lock()
    defer a.mu.Unlock()

    if a.IsClosed {
        return 0, false
    }
    closingBalance := a.Bal
    a.IsClosed = true
    a.Bal = 0
    return closingBalance, true
}
