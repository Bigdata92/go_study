// func(2)
package main

import "fmt"

func sum(i int, f func(int, int)) {
	f(i, 10)
}

func add(a, b int) {	// 매개 변수 둘다 int
	fmt.Println("ex1 : ", a+b)
}

func multi_val(i int) {
	i *= 10
}

func multi_reference(i *int) {
	*i *= 10
}

func main() {
	// 함수(callback) 전달, 참조 전달(call by ref), 값 전달(call by val)
	// ex1
	sum(100, add)	// 함수 전달

	// ex2
	// call by value
	a := 100
	multi_val(a)
	fmt.Println("ex2 : ", a)

	// ex3
	// call by value
	b := 100
	multi_reference(&b)
	fmt.Println("ex3 : ", b)



}