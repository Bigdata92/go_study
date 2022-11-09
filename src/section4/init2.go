// go 초기화 func(2)

package main

import (
	"fmt"
	"section/lib"
)

func init() {
	fmt.Println("Init1 Method start!")
}

func init() {
	fmt.Println("Init2 Method start!")
}

func init() {
	fmt.Println("Init3 Method start!")
}

func init() {
	fmt.Println("Init4 Method start!")
}

func main() {
	fmt.Println("Main Method start!")
}
