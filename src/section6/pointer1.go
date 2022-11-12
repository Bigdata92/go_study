// data type : pointer(1)

package main

import "fmt"

func main() {

	// Go : pointer 지원(c)
	// 변수의 지역성, 연속된 메모리 참조, heap, stack
	// python, java(JRE), compiler, interpreter
	// pointer 지원x 언어(python, c#, java 등)
	// 주소의 값은 직접 변경 불가능(잘못된 코딩으로 인한 버그 방지)
	// *(asterisk) 로 사용
	// nil로 초기화(nil != 0)

	// ex1
	var a *int            // 방법1
	var b *int = new(int) // 방법2

	fmt.Println(a) // &
	fmt.Println(b)

	i := 7
	a = &i // & 주소값 전달
	b = &i

	*a = 77

	fmt.Println("ex1 : ", a, &a, *a) // *a : 역참조
	fmt.Println("ex1 : ", b, &b, *b)

	var c = &i
	d := &i

	fmt.Println("ex1 : ", c, &c, *c)
	fmt.Println("ex1 : ", d, &d, *d)

}
