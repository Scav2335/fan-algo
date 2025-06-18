package priority_queue

import "time"

type Transaction struct {
	Name   string
	Date   time.Time
	Amount float64
}

func (t Transaction) Get() Transaction {
	return t
}

func (t Transaction) LessThan(a Transaction) bool {
	return t.Amount < a.Amount
}

func (t Transaction) GreaterThan(a Transaction) bool {
	return t.Amount > a.Amount
}

func (t Transaction) CompareTo(a Transaction) int {
	if t.Amount < a.Amount {
		return -1
	} else if t.Amount > a.Amount {
		return 1
	} else {
		return 0
	}
}
