// data type : slice(2)

package main

import "fmt"

func main() {

	// slice(slice 참조 타입 증명)

	// ex1(array)
	arr1 := [3]int{1, 2, 3}
	var arr2 [3]int

	arr2 = arr1 // 복사
	arr2[0] = 7

	fmt.Println("ex1 : ", arr1)
	fmt.Println("ex1 : ", arr2)

	// ex2(slice)2
	slice1 := []int{1, 2, 3}
	var slice2 []int

	slice2 = slice1
	arr2[0] = 7

	fmt.Println("ex1 : ", slice1)
	fmt.Println("ex1 : ", slice2)

	// ex3(slice 예외 상황)
	slice3 := make([]int, 50, 100)
	fmt.Println("ex3 : ", slice3[4])
	// fmt.Println("ex3 : ", slice3[50])	// idx over 예외(용량x 길이 기준 초기화)

	// ex4(slice 순서)
	for i, v := range slice1 {
		fmt.Println("ex4 : ", i, v)
	}

}
