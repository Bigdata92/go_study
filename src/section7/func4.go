// func(4)
package main

import "fmt"

func multiply(x, y int) (r1 int, r2 int) {
	r1 = x * 10
	r2 = y * 20
	return r1, r2
}


func main() {
	// return value 변수 사용

	a, b := multiply(10, 5)
	fmt.Println("ex1 : ", a, b)

}