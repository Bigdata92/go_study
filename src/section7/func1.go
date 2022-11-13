// func(1)
package main

import "fmt"
import "strconv"

// 함수 선언 위치 어느곳이든 가능
func helloGolang(){
	fmt.Println("ex1 : Hello Golang!!")
}

func say_one(m string){
	fmt.Println("ex2 : ", m)
}

func sum(x int, y int) int {
	return x + y
}

func main() {
	// 함수 선언 : func 키워드로 선언
	// func 함수명(매개변수) (return type or return value 변수명) : return value 존재
	// 타 언어와 달리 return value 여러개 가능

	// ex1
	helloGolang()

	// ex2
	say_one("Hello World!")

	// ex3
	result := sum(4, 6)
	fmt.Println("ex3 : ", result)
	fmt.Println("ex3 : ", strconv.Itoa(result))

}