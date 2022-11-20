// Go error basic(3)

package main

import (
	"errors"
	"fmt"
)

func main() {
	// error 처리
	// errors pkg의 New method 활용한 error 생성
	// Errof 보다 더 많이 사용

	var err1 error = errors.New("Error occurred - 1")
	err2 := errors.New("Error occurred - 2")

	fmt.Println("error1 : ", err1)
	fmt.Println("error1 : ", err1.Error())

	fmt.Println("error2 : ", err2)
	fmt.Println("error2 : ", err2.Error())
}
