package accounts

import (
	"errors"
	"fmt"
)

// BankAccount struct
type BankAccount struct {
	owner   string
	name    string
	copName string
	balance int
}

var errNoMoney = errors.New("Can't Withdraw. Your Broke")

// In the below part, we are making a function that returns address(&)
// Because Go doesn't have a contructor

// NewAccount creates Account
func NewAccount(owner string, name string, copName string) *BankAccount {
	// any requirements will go into the parentheses
	account := BankAccount{owner: owner, balance: 0, name: name, copName: copName}
	return &account // returning an address
}

// Deposit (PointerReceiver) x amount on your account
func (b *BankAccount) Deposit(amount int) {
	b.balance += amount
}

// Withdraw (PointerReceiver)
func (b *BankAccount) Withdraw(amount int) error {
	if b.balance < amount {
		return errNoMoney
	}
	b.balance -= amount
	return nil
}

// SeeBalance (Receiver) shows the balance of the account
func (b BankAccount) SeeBalance() int {
	return b.balance
}

// Owner shows the owner
func (b *BankAccount) Owner() string {
	return b.owner
}

func (b BankAccount) String() string {
	return fmt.Sprintf("Current Account Owner : %s", b.Owner())
}

// ChangeOwner of the account
func (b *BankAccount) ChangeOwner(newOwner string) string {
	b.owner = newOwner
	return "The New Owner is : " + newOwner
}
