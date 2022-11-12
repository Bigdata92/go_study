// data type : array(1)

package main

import "fmt"

func main() {
	// array
	// array는 용량, 길이 항상 같다.
	// (중요) array vs slice 차이점
	// array : 길이 고정, 값 type(복사 전달) vs slice : 길이 가변, 참조 type(참조값 전달)
	// 비교 연산자 사용 가능 vs 비교 연산자 사용 불가능
	// 대부분 slice 사용

	// cap() : array or slice 용량
	// len() : array or slice 개수

	// ex1
	var arr1 [5]int
	var arr2 [5]int = [5]int{1, 2, 3, 4, 5} // array 원소 개수 및 type 모두 맞아야 됨
	var arr3 = [5]int{1, 2, 3, 4, 5}
	arr4 := [5]int{1, 2, 3, 4, 5}
	arr5 := [5]int{1, 2, 3}
	arr6 := [...]int{1, 2, 3, 4, 5} // array 크기, 자동 맞춤
	arr7 := [5][5]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
	}

	arr1[2] = 5

	fmt.Printf("%-5T %d %v\n", arr1, len(arr1), arr1)
	fmt.Printf("%-5T %d %v\n", arr2, len(arr2), arr2)
	fmt.Printf("%-5T %d %v\n", arr3, len(arr3), arr3)
	fmt.Printf("%-5T %d %v\n", arr4, len(arr4), arr4)
	fmt.Printf("%-5T %d %v\n", arr5, len(arr5), arr5)
	fmt.Printf("%-5T %d %v\n", arr6, len(arr6), arr6)
	fmt.Printf("%-5T %d %v\n", arr7, len(arr7), arr7)

	// ex2
	arr8 := [5]int{1, 2, 3, 4, 5}
	arr9 := [5]int{
		1,
		2,
		3,
		4,
		5,
	}
	arr10 := [...]string{"Kim", "Lee", "Park"}
	fmt.Printf("%-5T %d %v\n", arr8, len(arr8), arr8)
	fmt.Printf("%-5T %d %v\n", arr9, len(arr9), arr9)
	fmt.Printf("%-5T %d %v\n", arr10, len(arr10), arr10)
}
