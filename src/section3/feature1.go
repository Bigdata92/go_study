// Go 특징 1

package main

import "fmt"

func main() {
	// GO : 모호하거나, 애매한 문법 금지
	// 후치(후위) 연산자 허용 (i++), 전치(전위) 연산자 비허용 (++i)
	// 증감연산자 반환값 없음 sum := i++
	// Pointer 허용, 연산 비허용

	// ex1
	sum, i := 0, 0
	for i <= 100 {
		sum += i
		i++ // ++i error
	}
	fmt.Println("ex1 : ", sum)
}
