package command

import "fmt"

type BankAccountFunctional struct {
	Balance int
}

func DepositFunctional(ba *BankAccountFunctional, amount int) {
	fmt.Println("Depositing", amount)
	ba.Balance += amount
}

func WithdrawFunctional(ba *BankAccountFunctional, amount int) {
	if ba.Balance >= amount {
		fmt.Println("Withdrawing", amount)
		ba.Balance -= amount
	}
}

func CreateFunctionalCommand() {
	ba := &BankAccountFunctional{0}
	var commands []func()
	commands = append(commands, func() {
		DepositFunctional(ba, 100)
	})
	commands = append(commands, func() {
		WithdrawFunctional(ba, 100)
	})

	for _, cmd := range commands {
		cmd()
	}
}
