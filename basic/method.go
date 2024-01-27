package basic

import "fmt"

type account struct {
	balance   int
	firstName string
	lastName  string
}

func (a1 *account) withdrawPointer(amount int) {
	a1.balance -= amount
}

func (a2 account) withdrawValue(amount int) {
	a2.balance -= amount
}

func (a3 account) withdrawReturnValue(amount int) account {
	a3.balance -= amount
	return a3
}

func MethodExample() {
	acc := account{balance: 100, firstName: "Park", lastName: "heesoo"}

	acc.withdrawPointer(30)
	fmt.Println(acc.balance)

	acc.withdrawValue(20)
	fmt.Println(acc.balance)

	acc2 := acc.withdrawReturnValue(20)
	fmt.Println(acc2.balance)
}
