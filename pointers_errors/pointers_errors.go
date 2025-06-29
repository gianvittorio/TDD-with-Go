package pointerserrors

import (
	"errors"
	"fmt"
)

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet interface {
	Deposit(amount Bitcoin)
	Withdraw(amount Bitcoin) error
	Balance() Bitcoin
}

type WalletImpl struct {
	balance Bitcoin
}

func (w *WalletImpl) Deposit(amount Bitcoin) {
	w.balance += amount
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *WalletImpl) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	
	w.balance -= amount
	
	return nil
}

func (w *WalletImpl) Balance() Bitcoin {
	return w.balance
}
