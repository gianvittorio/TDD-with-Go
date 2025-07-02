package arraysandsliceswithgenerics

func Reduce[T, U any](collection []T, aggregator func(aggregate *U, next *T), initialValue U) U {
	result := initialValue
	for _, elem := range collection {
		aggregator(&result, &elem)
	}

	return result
}

type Transaction struct {
	From, To string
	Sum      float64
}

func NewTransaction(from, to Account, sum float64) Transaction {
	return Transaction{From: from.Name, To: to.Name, Sum: sum}
}

type Account struct {
	Name string
	Balance float64
}

func NewBalanceFor(account Account, transactions []Transaction) Account {
	return Reduce(transactions, applyTransaction, account)
}

func applyTransaction(account *Account, transaction *Transaction) {
	switch {
	case transaction.From == account.Name:
		account.Balance -= transaction.Sum
	case transaction.To == account.Name:
		account.Balance += transaction.Sum
	}
}

func Find[T any](collection []T, predicate func(value *T) bool) (T, bool) {
	var result T
	found := false
	for _, elem := range collection {
		if predicate(&elem) {
			result = elem
			found = true
			break
		}
	}
	
	return result, found
}
