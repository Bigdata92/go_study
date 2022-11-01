package main

import "fmt"

func main() {
	// 제어문(조건문)
	// IF 문 : 반드시 Boolean 검사 -> 1, 0 (사용불가 : 자동 형 변환 불가)
	// 소괄호 사용 x

	var a int = 20
	b := 20

	// ex1
	if a >= 15 {
		fmt.Println("15이상")
	}

	if b >= 25 {
		fmt.Println("25이상")
	}

	// error 발생1
	/*
		if b >= 25
		{

		}
	*/

	// error 발생2
	/*
		if b >= 25
			fmt.Println("25이상")
	*/

	// error 발생3
	/*
		if c := 1; c {
			fmt.Println("True")
		}
	*/

	// ex2
	if c := 40; c >= 35 {
		fmt.Println("35이상")
	}

	// c += 20 // err발생
}
