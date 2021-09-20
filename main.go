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

	word := "hello"
	definition := "Greeting"

	err := dictionary.Add(word, definition)
	if err != nil {
		fmt.Println(err)
	}
	hello, _ := dictionary.Search(word)
	fmt.Println(hello)

	err2 := dictionary.Add(word, definition)
	if err2 != nil {
		fmt.Println(err2)
	}

}
