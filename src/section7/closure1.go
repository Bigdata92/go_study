// 함수 closure(1)
package main

import "fmt"

func main(){

	// closure
	// 익명함수 사용시, 함수를 변수에 할당해서 사용가능
	// 함수 안에서 함수 선언 및 정의 가능 -> 외부 함수에 선언된 변수에 접근 가능한 함수
	// 함수가 선언되는 순간에, 함수가 실행시, 실체의 외부 변수에 접근하기 위한 snapshot(객체)
	// 함수 호출시, 이전에 존재했던 값을 유지하기 위해 -> 비동기, 누적카운트, 무분별한 전역변수 남발 방지
	// 남발 -> 객체들이 메모리에 존재하므로 -> 메모리부족, 오버플로우 현상, 자원을 무분별하게 사용 가능성
	// 클로저 정확하게 이해 및 사용해야

	// ex1
	multiply := func(x, y int) int {// 익명함수
		return x * y
	}

	r1 := multiply(5, 10)

	fmt.Println("ex1 : ",  r1)

	// ex2
	m, n := 5, 10	// 변수가 캡처
	sum := func(c int) int { // 익명함수 변수 할당
		return m + n + c	// 지역변수 소멸x (함수 호출시마다 사용 가능)
	}

	r2 := sum(10)

	fmt.Println("ex2 : ",  r2)

}