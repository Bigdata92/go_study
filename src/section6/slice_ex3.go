// data type: slice 심화(3)

package main

import (
	"fmt"
)

func main() {

	// slice copy
	// copy(복사 대상, 원본)
	// make 로 공간을 할당 후 복사 해야한다.
	// 복사 된 슬라이스 값 변경해도 원본에는 영향 x

	// ex1(복사)
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice2 := make([]int, 5)
	slice3 := []int{}

	copy(slice2, slice1)
	copy(slice3, slice1)	// copy x

	fmt.Println("ex1 : ", slice2)
	fmt.Println("ex1 : ", slice3)

	fmt.Println()

	// ex2
	a := []int{1, 2, 3, 4, 5}
	b := make([]int, 5)

	copy(b, a)

	b[0] = 7
	b[4] = 10

	fmt.Println("ex2 : ", a)
	fmt.Println("ex2 : ", b)

	fmt.Println()

	// ex3
	c := [5]int{1, 2, 3, 4, 5}
	d := c[:3]	// 주의! slice도 추출은 주소 참조 -> 원본 값 변경

	d[1] = 7

	fmt.Println("ex3 : ", c)
	fmt.Println("ex3 : ", d)

	// ex4
	e := []int{1, 2, 3, 4, 5, 6, 7}
	f := e[:5:7]	// 7 : 용량 지정

	fmt.Println("ex4 : ", len(f), cap(f))
	fmt.Println("ex4 : ", f)
}
