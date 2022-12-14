// data type : Bool(1)
package main

import "fmt"

func main() {
	// Bool(Boolean) : True, False
	// 조건부 논리 연산자랑 같이 사용 : !, ||(or), &&(and)
	// 암묵적 형변환 x : 0, Nil -> False 변환 x

	// ex1
	var b1 bool = true
	var b2 bool = false
	b3 := true
	// b4 := 1

	fmt.Println("ex1 : ", b1)
	fmt.Println("ex2 : ", b2)
	fmt.Println("ex3 : ", b3)

	fmt.Println("ex4 : ", b3 == b3)
	fmt.Println("ex5 : ", b1 && b3)
	fmt.Println("ex6 : ", b1 || b3)
	fmt.Println("ex7 : ", !b1)

	/*
		if b4 {
			fmt.Println("ex8 : true")
		}
	*/
}
