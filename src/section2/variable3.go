package main

import "fmt"

func main() {
	// 짧은 선언 
	// 함수 안에서만 사용(전역 x), 선언 후 할당시 예외 발생
	// 주로 제한된 범위의 함수내에서 사용 할 경우, 코드 가독성 높일 수 있음 -> 추천

	shortVal1 := 3
	shortVal2 := "Test"
	shortVal3 := false

	// shortVal1 := 10 // 예외 발생(재할당 x)

	fmt.Println("shortVal1 : ", shortVal1, "shortVal2 : ", shortVal2, "shortVal3 : ", shortVal3,)

	// ex if문
	if i := 10; i < 11 {
		fmt.Println("Short Variable Test Success!")
	}

}