package main

import (
	"fmt"
	"github.com/ksy/learngo/dict"
)

//// accounts main
//func main() {
//	account := accounts.NewAccount("ksy")
//	fmt.Println(account)
//
//	account.Deposit(10)
//	fmt.Println(account.Balance())
//
//	err := account.Withdraw(20)
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(account.Balance(), account.Owner())
//}

func main() {
	dictionary := dict.Dictionary{}
	dictionary["hello"] = "hello"

	definition, err := dictionary.Search("hello")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
}
