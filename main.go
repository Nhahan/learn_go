package main

import (
	"fmt"
	"github.com/ksy/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("ksy")
	fmt.Println(account)
}
