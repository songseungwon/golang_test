package accounts

import "errors"

//Account struct
type Account struct {
	owner   string //예금주
	balance int    //잔고
}

//NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

//Deposit + amount on your account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

//Balance of your account
func (a *Account) Balance() int {
	return a.balance
}

var errNoMoney = errors.New("Can't withdraw. You are poor")

//Withdraw on your account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}
