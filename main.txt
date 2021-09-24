package main

import (
	"fmt"
	"net/http"
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

//func main() {
//	dictionary := dict.Dictionary{}
//
//	word := "hello"
//	definition := "Greeting"
//
//	err := dictionary.Add(word, definition)
//	if err != nil {
//		fmt.Println(err)
//	}
//	hello, _ := dictionary.Search(word)
//	fmt.Println(hello)
//
//	err2 := dictionary.Add(word, definition)
//	if err2 != nil {
//		fmt.Println(err2)
//	}

type requestResult struct {
	url    string
	status string
}

func main() {
	c := make(chan requestResult)
	urls := []string{
		"https://www.naver.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
	}

	for _, url := range urls {
		go hitURL(url, c)
	}
	for i := 0; i < len(urls); i++ {
		fmt.Println(<-c)
	}
}

func hitURL(url string, c chan<- requestResult) { // chan<- : send only
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- requestResult{url: url, status: status}
}
