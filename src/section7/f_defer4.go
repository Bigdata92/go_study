// 함수 defer(4)

package main

import (
	"fmt"
)

func start(t string) string {
	fmt.Println("start : ", t)
	return t
}

func end(t string) {
	fmt.Println("end : ", t)
}

func a() {
	defer end(start("b"))	// end만 defer 적용(start x) -> defer에는 중첩함수 x(단일 함수만)
	fmt.Println("in a")
}

func main() {
	// ex1
	a()
}