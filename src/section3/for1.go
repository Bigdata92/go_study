package main

import "fmt"

func main() {
	// 반복문 - for문
	// Go에서 유일하게 제공되는 반복문
	// 다양한 사용법 숙지

	// ex1
	for i := 0; i < 5; i++ {
		fmt.Println("ex1 : ", i)
	}

	// error1
	/*
		for i := 0; i < 5; i++
		{
			fmt.Println("ex1 : ", i)
		}
	*/

	// erro2
	/*
		for i := 0; i < 5; i++
		fmt.Println("ex1 : ", i)
	*/

	// ex2 (무한 루프)
	/*
		for {
			fmt.Println("ex2 : Hello, Golang!")
			fmt.Println("ex2 : Infinite, loop!")
		}
	*/

	// ex3 (Range 용법)
	loc := []string{"Seoul", "Busan", "Incheon"}
	for index, name := range loc {
		fmt.Println("ex3 : ", index, name)
	}
}
