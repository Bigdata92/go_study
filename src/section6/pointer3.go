// data type : pointer(3)

package main

import "fmt"

func rptc(n *int) {
	*n = 77
}

func vptc(n int) {
	n = 77
}

func main() {

	// pointer 값 전달
	// 함수, 메서드 호출 시 매개변수 값 복사 전달 -> 함수, 메서드 내에서는 원본 값 변경 불가능
	// 원본 값 변경 위해서 포인터로 전달
	// 특히 크기가 큰배열인 경우 값 복사시 시스템 부담 -> 포인터 전달로 해결(slice, map 참조 전달)

	// ex1
	var a int = 10
	var b int = 10

	fmt.Println("ex1 : ", a)
	fmt.Println("ex1 : ", b)
	fmt.Println()

	rptc(&a)
	vptc(b)
	//vptc(&b)	// error 발생

	fmt.Println("ex2 : ", a)
	fmt.Println("ex2 : ", b)
}