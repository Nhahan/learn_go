package main

import (
	"fmt"
	"github.com/ksy/learngo/something"
)

func main() { // Go는 main 패키지의 main func에서 시작. 컴파일을 위해 반드시 있어야함!
	fmt.Println("Hello world!")

	something.SayHello() // Go는 func 이름이 대문자로 시작하면 public, 소문자로 시작하면 private
	//something.sayBye() <- 소문자로 시작하므로 private이라 호출 불가

	const constName string = "ksy" // const 상수
	var varName string = "ksy"     // var은 변수인데 사용하지 않는 변수는 오류가 뜨고 컴파일이 되지 않음
	name := "ksy"                  // 축약 변수 선언
	fmt.Println(varName)
	fmt.Println(name)
}
