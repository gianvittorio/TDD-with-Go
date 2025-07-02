package arraysandsliceswithgenerics_test

import (
	"strings"
	"testing"

	"example.com/arraysAndSlicesWithGenerics"
)

func TestSum(t *testing.T) {
	t.Run("Collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got, want := arraysandsliceswithgenerics.Reduce(
			numbers,
			func(s, num *int) {
				*s += *num
			}, 0), 15
		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestReduce(t *testing.T) {
	t.Run("multiplication of all elements", func(t *testing.T) {
		multiply := func(x, y *int) {
			*x *= *y
		}

		AssertEqual(t, arraysandsliceswithgenerics.Reduce([]int{1, 2, 3}, multiply, 1), 6)
	})

	t.Run("concatenate strings", func(t *testing.T) {
		concatenate := func(x, y *string) {
			*x += *y
		}

		AssertEqual(t, arraysandsliceswithgenerics.Reduce([]string{"a", "b", "c"}, concatenate, ""), "abc")
	})
}

func TestBadBank(t *testing.T) {
	var (
		riya  = arraysandsliceswithgenerics.Account{Name: "Riya", Balance: 100}
		chris = arraysandsliceswithgenerics.Account{Name: "Chris", Balance: 75}
		adil  = arraysandsliceswithgenerics.Account{Name: "Adil", Balance: 200}

		transactions = []arraysandsliceswithgenerics.Transaction{
			arraysandsliceswithgenerics.NewTransaction(chris, riya, 100),
			arraysandsliceswithgenerics.NewTransaction(adil, chris, 25),
		}
	)

	newBalanceFor := func(account arraysandsliceswithgenerics.Account) float64 {
		return arraysandsliceswithgenerics.NewBalanceFor(account, transactions).Balance
	}

	AssertEqual(t, newBalanceFor(riya), 200)
	AssertEqual(t, newBalanceFor(chris), 0)
	AssertEqual(t, newBalanceFor(adil), 175)
}

func AssertEqual[T comparable](t testing.TB, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func AssertNotEqual[T comparable](t testing.TB, got, want T) {
	t.Helper()
	if got == want {
		t.Errorf("didn't want %v", got)
	}
}

func TestFind(t *testing.T) {
	t.Run("find first even number", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		firstEvenNumber, found := arraysandsliceswithgenerics.Find(numbers, func(x *int) bool {
			return *x%2 == 0
		})
		AssertTrue(t, found)
		AssertEqual(t, firstEvenNumber, 2)
	})

	type Person struct {
		Name string
	}

	t.Run("Find the best programmer", func(t *testing.T) {
		people := []Person{
			Person{Name: "Kent Beck"},
			Person{Name: "Martin Fowler"},
			Person{Name: "Chris James"},
		}

		king, found := arraysandsliceswithgenerics.Find(people, func(p *Person) bool {
			return strings.Contains(p.Name, "Chris")
		})

		AssertTrue(t, found)
		AssertEqual(t, king, Person{Name: "Chris James"})
	})
}

func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got %v, want true", got)
	}
}

func AssertFalse(t *testing.T, got bool) {
	t.Helper()
	if got {
		t.Errorf("got %v, want false", got)
	}
}
