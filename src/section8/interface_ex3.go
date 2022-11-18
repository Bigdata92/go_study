// interface 고급(3)

package main

import (
	"fmt"
)

func checkType(arg interface{}) {
	// arg.(type) 을 통해서 현제 data type return
	switch arg.(type) {
	case bool:
		fmt.Println("This is a bool", arg)
	case int, int8, int16, int32, int64:
		fmt.Println("This is a int", arg)
	case float32, float64:
		fmt.Println("This is a float", arg)
	case string:
		fmt.Println("This is a string", arg)
	case nil:
		fmt.Println("This is a nil", arg)
	default:
		fmt.Println("what is this type?", arg)
	}

}

func main() {
	// 실제 type 검사 switch 사용
	// empty interface - all type 전달 ok -> type check를 통해 형 변환 후, 사용 가능

	// ex1
	checkType(true)
	checkType(1)
	checkType(22.542)
	checkType(nil)
	checkType("Hello Golang!")

}
