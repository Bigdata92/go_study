// data type : 문자열(1)

package main

import (
	"fmt"
)

func main() {
	// 문자열 표현
	// 문자열 : UTF-8 인코딩 (유니코드 문자 집합) -> byte 수 주의

	// ex1
	var str1 string = "Golang"
	var str2 string = "World"
	var str3 string = "고프로그래밍"

	fmt.Println("ex1 : ", str1[0], str1[1], str1[2], str1[3], str1[4], str1[5])
	fmt.Println("ex1 : ", str2[0], str2[1], str2[2], str2[3], str2[4])
	fmt.Println("ex1 : ", str3[0], str3[1], str3[2], str3[3], str3[4], str3[5])

	// ex2
	fmt.Printf("ex2 : %c %c %c %c %c %c\n", str1[0], str1[1], str1[2], str1[3], str1[4], str1[5])
	fmt.Printf("ex2 : %c %c %c %c %c\n", str2[0], str2[1], str2[2], str2[3], str2[4])
	fmt.Printf("ex2 : %c %c %c %c %c %c\n", str3[0], str3[1], str3[2], str3[3], str3[4], str3[5])

	conStr := []rune(str3) // 한글 문자열 깨지지 않게 하는법
	fmt.Printf("ex2 : %c %c %c %c %c %c\n", conStr[0], conStr[1], conStr[2], conStr[3], conStr[4], conStr[5])

	// ex3
	for i, char := range str1 {
		fmt.Printf("ex3 : %c(%d)\t", char, i)
	}
	fmt.Println()

	for i, char := range str2 {
		fmt.Printf("ex4 : %c(%d)\t", char, i)
	}
}
