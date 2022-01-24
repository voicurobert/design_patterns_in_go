package memento

import "fmt"

type BankAccount struct {
	balance int
}

type Memento struct {
	Balance int
}

func NewBankAccount(balance int) (*BankAccount, *Memento) {
	return &BankAccount{balance: balance}, &Memento{Balance: balance}
}

func (b *BankAccount) Deposit(amount int) *Memento {
	b.balance += amount
	return &Memento{b.balance}
}

func (b *BankAccount) Restore(m *Memento) {
	b.balance = m.Balance
}

func Example() {
	ba, m0 := NewBankAccount(100)
	m1 := ba.Deposit(50)
	m2 := ba.Deposit(25)

	fmt.Println(ba)

	ba.Restore(m1)
	fmt.Println(ba)
	ba.Restore(m2)
	fmt.Println(ba)

	ba.Restore(m0)
	fmt.Println(ba)
}
