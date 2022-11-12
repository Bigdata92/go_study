// data type: slice 심화(1)

package main

import "fmt"

func main() {

	// slice 추가 및 병합
	// ex1

	s1 := []int{1, 2, 3, 4, 5} // 길이5, 용량10
	s2 := []int{8, 9, 10, 11, 12}
	s3 := []int{13, 14, 15, 16, 17}

	s1 = append(s1, 6, 7)
	s2 = append(s1, s2...)     // slice에 slice 대입시 ... 사용
	s3 = append(s2, s3[:3]...) // 추출 후 병합

	fmt.Println("ex1 : ", s1)
	fmt.Println("ex1 : ", s2)
	fmt.Println("ex1 : ", s3)

	// ex2
	s4 := make([]int, 0, 5)
	for i := 0; i < 15; i++ {
		s4 = append(s4, i)
		fmt.Println("ex2 -> len : %d, cap : %d, value : %v\n", len(s4), cap(s4), s4) // 길이 및 용량 자동 증가
	}

}
