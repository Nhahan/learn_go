## 📣 Hello, Go!



```
package main

import "fmt"

func main() {
        fmt.Println("Hello, World!")
}
```
<br>

### ✅ Pointer
& : address

<a>*</a> : see through

<br>

### ✅ Structure, Array, Slice
```
type person struct {
	name string
	age int
	favFood []string // slice
	family [4]string // array
}
```
<br>

### ✅ Method, Error

```
type Account struct {
	owner   string
	balance int
}

// NewAccount Constructor function
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Balance method
func (a Account) Balance() int {
	return a.balance
}

// Deposit method
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Withdraw method
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

var errNoMoney = errors.New("Can't withdraw")
```
```
func main() {
	account := accounts.NewAccount("ksy")
	fmt.Println(account)

	account.Deposit(10)
	fmt.Println(account.Balance())

	err := account.Withdraw(20)
	// Exception occured because withdraw > deposit
	if err != nil {
		log.Fatalln(err) 
	}
	fmt.Println(account.Balance())
}
```
#### result
![image](https://user-images.githubusercontent.com/81916648/134031118-2fe8157b-06e2-4c1d-a321-2fd161fe9ffd.png)

<br>

### ✅ Goroutine, Channel
<b>Channel</b> is the way to communicate with the main function from a <b>Goroutine</b>
```
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
```
