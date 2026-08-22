package main

import "fmt"

type BankAccount struct {
	name        string
	acc_no      int
	cur_balance float64
}

func (b *BankAccount) Deposit(add float64) {
	b.cur_balance += add
}

func (b *BankAccount) Withdraw(minus float64) {
	if b.cur_balance < minus {
		return
	}
	b.cur_balance -= minus
}

func (b *BankAccount) Balance() {

	fmt.Println("current balance is", b.cur_balance)
}

func (b *BankAccount) Info() {
	fmt.Println("name is :", b.name, "\n acc no is: ", b.acc_no, "\n balance is :", b.cur_balance)
}

func main() {
	b := BankAccount{name: "swastik", acc_no: 123, cur_balance: 1234.99}
	b.Deposit(99.01)
	b.Withdraw(88991.2)
	b.Balance()
	b.Info()
}
