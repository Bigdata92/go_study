// data type : string operation(1)

package main

import "fmt"

func main() {
	// string operation
	// 추출, 비교, 조합(결합)

	// ex1
	var str1 string = "Golang"
	var str2 string = "World"

	fmt.Println("ex1 : ", str1[0:2], str1[0], str1[0:1])
	fmt.Println("ex1 : ", str2[3:], str2[0])
	fmt.Println("ex1 : ", str2[:4])
	fmt.Println("ex1 : ", str1[1:3])
}
