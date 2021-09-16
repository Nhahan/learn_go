package main

import (
	"fmt"
	"github.com/ksy/learngo/something"
	"strings"
)

func multiply(a int, b int) int { // func 타이핑
	return a * b
}

func lenAndUpper(name string) (int, string) { // return 2개 이상인 func
	defer fmt.Println("I'm done!") // func 끝난 후에 실행되는 코드 (많이 쓴다고 함)
	return len(name), strings.ToUpper(name)
}

func superAdd(numbers ...int) int {
	fmt.Println(numbers)

	total := 0

	for number := range numbers {
		total += number
	}
	return total
}

func main() { // Go는 main 패키지의 main func에서 시작. 컴파일을 위해 반드시 있어야함!
	fmt.Println("Hello world!")

	something.SayHello() // Go는 func 이름이 대문자로 시작하면 public, 소문자로 시작하면 private
	//something.sayBye() <- 소문자로 시작하므로 private이라 호출 불가

	const constName string = "ksy" // const 상수
	var varName string = "ksy"     // var은 변수인데 사용하지 않는 변수는 오류가 뜨고 컴파일이 되지 않음
	name := "ksy"                  // 축약 변수 선언
	fmt.Println(varName)
	fmt.Println(name)

	fmt.Println(multiply(2, 2))
	totalLength, upperName := lenAndUpper("ksy") // 만약 2개 이상을 리턴하는 func에서 1개만 리턴하고 싶다면 리턴 변수명을 _로 해주면 ignore된다
	fmt.Println(totalLength, upperName)

	result := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(result)
}
