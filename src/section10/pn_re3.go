// Panic & Recover(3)

package main

import "fmt"

func runFunc() {
	defer func() {
		if s := recover(); s != nil {
			fmt.Println("Error Message : ", s)
		}

	}()

	a := [3]int{1, 2, 3}
	for i := 0; i < 5; i++ {
		fmt.Println("ex1 : ", a[i]) //error 발생(idx 범위)
	}
}

func main() {
	// Go Recover 함수
	// error 복구 가능
	// Panic으로 발생한 에러를 복구 후 코드 재실행(프로그램 종료 되지 않는다.)
	// 즉, 코드 흐름을 정상상태로 복구하는 기능
	// Panic에서 설정한 메시지를 받아올 수 있다.

	// ex1
	runFunc()

	fmt.Println("Hello Golang!")
}
