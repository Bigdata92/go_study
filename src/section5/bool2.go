package main

import "fmt"

func main() {

	// ex1
	fmt.Println("ex1 : ", true && true)
	fmt.Println("ex1 : ", true && false)
	fmt.Println("ex1 : ", false && false)
	fmt.Println("ex1 : ", true || true)
	fmt.Println("ex1 : ", true || false)
	fmt.Println("ex1 : ", false || false)
	fmt.Println("ex1 : ", !true)
	fmt.Println("ex1 : ", !false)

	// ex2

	num1 := 15
	num2 := 37

	fmt.Println("ex2 : ", num1 < num2)
	fmt.Println("ex2 : ", num1 > num2)
	fmt.Println("ex2 : ", num1 <= num2)
	fmt.Println("ex2 : ", num1 >= num2)
	fmt.Println("ex2 : ", num1 == num2)
	fmt.Println("ex2 : ", num1 != num2)
}
