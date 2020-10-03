package main

import (
	"fmt"

	"github.com/childult-programmer/learngo/accounts"
)

func main() {
	// Owner : "Lee", Balance : 0
	myAccount := accounts.NewAccount("Lee")

	myAccount.Deposit(10000)
	myAccount.ChangeOwner("Kim")

	// when withdraw money from account,
	err := myAccount.Withdraw(50000)
	if err != nil {
		// log.Fatalln(err) // if you want kill the program
		fmt.Println(err)
	}

	fmt.Println(myAccount)
}
