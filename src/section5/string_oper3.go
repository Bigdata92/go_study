// data type : string operation(3)

package main

import (
	"fmt"
	"strings"
)

func main() {

	// ex1 (결합 : 일반연산)
	str1 := "Waiting for going home feels like a thousand hours" +
		"Going on the way home feels like a thousand miles" +
		"Maybe, wanna just get along in work" +
		"But it seems like it never works"

	str2 := "I wanna be somewhere like no need to be clear"

	fmt.Println("ex1 : ", str1+str2)

	// ex2 (결합 : join)(추천)
	strSet := []string{} // slice 선언
	strSet = append(strSet, str1)
	strSet = append(strSet, str2)

	fmt.Println("ex2 : ", strings.Join(strSet, ""))
}
