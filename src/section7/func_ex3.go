// 함수 심화(3)

package main

import "fmt"

func fact(n int) int {

	if n == 0 {
		return 1
	}
	return n * fact(n-1)

}

func prtHello(n int) {
	if n == 0 {
		return
	}
	fmt.Println("ex2 : (", n, ")", "hi!")
	prtHello(n - 1)
}

func main() {

	// 함수 고급
	// 재귀 함수(Recursion)
	// prog 보기 쉽고, 코드 간결, 오류 수정 용이 - 장점
	// 코드 이해 어렵고, 기억공간 많이 사용, 무한 루프 가능성

	// ex1
	x := fact(10)
	fmt.Println("ex1 : ", x)

	// ex2
	prtHello(10)

}
