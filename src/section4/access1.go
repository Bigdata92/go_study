// package 접근제어(1)

package main

import (
	"fmt"
	"section4/lib2"
)

func main() {
	// package 접근 제어
	// 변수, 상수, 함수, 메소드, 구조체 등 식별자
	// 대문자 - pkg 외부에서 접근 가능
	// 소문자 - pkg 외부에서 접근 불가(pkg 내에서만 접근 가능)

	fmt.Println("100보다 큰 수? : ", lib2.CheckNum1(101))
	fmt.Println("1000보다 큰 수? : ", lib2.CheckNum2(999))
}
