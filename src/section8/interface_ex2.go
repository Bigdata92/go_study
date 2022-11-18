// interface 고급(2)

package main

import (
	"fmt"
	"reflect"
)

func main() {
	// type 변환
	// Type Assertion
	// 실행(런타임) 시에는 interface에 할당한 변수는 실제 타입으로 변환 후 사용해야 하는 경우
	// interface.(type) -> 형 변환
	// interfaceVal.(type)

	// ex1
	var a interface{} = 15
	b := a
	c := a.(int)
	// d := a.(float64)		// error

	fmt.Println("ex1 : ", a)
	fmt.Println("ex1 : ", reflect.TypeOf(a))
	fmt.Println("ex1 : ", b)
	fmt.Println("ex1 : ", reflect.TypeOf(b))
	fmt.Println("ex1 : ", c)
	fmt.Println("ex1 : ", reflect.TypeOf(c))

	fmt.Println()

	// ex2(저장된 실제 type check)
	if v, ok := a.(int); ok {
		fmt.Println("ex2 : ", v, ok)
	}

}
