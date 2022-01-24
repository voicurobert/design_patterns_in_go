package memento

import "fmt"

type BankAccount2 struct {
	balance int
	changes []*Memento
	current int
}

func NewBankAccount2(balance int) *BankAccount2 {
	b := &BankAccount2{balance: balance}
	b.changes = append(b.changes, &Memento{balance})
	return b
}

func (b *BankAccount2) Deposit2(amount int) *Memento {
	b.balance += amount
	m := Memento{b.balance}
	b.changes = append(b.changes, &m)
	b.current++
	fmt.Println("Deposited", amount, ", balance is now", b.balance)
	return &m
}

func (b *BankAccount2) Restore(m *Memento) {
	if m != nil {
		b.balance = m.Balance
		b.changes = append(b.changes, m)
		b.current = len(b.changes) - 1
	}
}

func (b *BankAccount2) Undo() *Memento {
	if b.current > 0 {
		b.current--
		m := b.changes[b.current]
		b.balance = m.Balance
		return m
	}
	return nil
}

func (b *BankAccount2) Redo() *Memento {
	if b.current+1 < len(b.changes) {
		b.current++
		m := b.changes[b.current]
		b.balance = m.Balance
		return m
	}
	return nil
}

func (b *BankAccount2) String() string {
	return fmt.Sprintf("Balance = %d, current = %d", b.balance, b.current)
}

func UndeRedoExample() {
	ba := NewBankAccount2(100)
	ba.Deposit2(50)
	ba.Deposit2(25)
	fmt.Println(ba)

	ba.Undo()
	fmt.Println("Undo 1", ba)
	ba.Undo()
	fmt.Println("Undo 1", ba)
	ba.Redo()
	fmt.Println("Redo 1", ba)
}
