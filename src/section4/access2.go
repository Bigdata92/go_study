// package 접근제어(2)

package main

import (
	"fmt"
	checkUp "section4/lib"
	_ "section4/lib2"
)

func main() {
	// pkg 접근 제어
	// alias 사용
	// 빈 식별자 사용

	fmt.Println("10보다 큰 수? : ", checkUp.CheckNum1(11))
}
