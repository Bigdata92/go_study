// 함수 defer(3)

package main

import (
	"fmt"
)

func stack() {
	for i := 1; i <= 10; i++ {
		defer fmt.Println("ex1 : ", i)	// stack(LIFO)
	}
}

func main() {
	// ex1
	stack()
}