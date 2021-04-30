package accounts

import (
	"errors"
	"fmt"
)

// Account struct
type Account struct {
	owner 	string
	balance int
}

var errNoMoney = errors.New("Can't withdraw.")

// private struct를 활용하여 변수를 만들기 위해
//constructor를 대신해서 structf를 만드는 Public function을 생성

// NewAccount creates Account
func NewAccount(owner string) *Account {
	//* : pointer -> private struct를 가리키는 중
	account := Account{owner: owner, balance: 0}
	return &account // return the value of new obj with same memory addr w/ account struct
}

// Balance of your account balance
func (a Account) Balance() int {
	return a.balance
}

// Deposit x amount on your account
func (a *Account) Deposit(amount int) {
	// * : DON'T COPY THE RECEIVER
	a.balance += amount  // Use the account and add the amount on balance
}

// Withdraw from your account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil // error의 value : nil(에러가 없을 경우)
}

// ChangeOwner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner of the account
func (a Account) Owner() string {
	return a.owner
}


func (a Account) String() string {
	// GO's internal function -> GO calls the function before calling the object
	return fmt.Sprint(a.Owner(), "'s account.\nHas: ", a.Balance())
}