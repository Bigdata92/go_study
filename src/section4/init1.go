// go 초기화 func(1)

package main

import (
	"fmt"
)

func init() {
	// init : pkg load시 가장 먼저 호출
	// 가장 먼저 초기화 되는 작업 작성시 유용
	fmt.Println("Init Method start!")
}

func main() {
	fmt.Println("Main Method start!")
}
