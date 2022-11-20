// Go error basic(4)

package main

import (
	"fmt"
	"log"
)

func notZero(n int) (string, error) { // method return val error type 중요
	if n != 0 {
		s := fmt.Sprint("Hello Golang : ", n)
		return s, nil
	}

	return "", fmt.Errorf("0을 입력했습니다. 에러 발생!") // return value로 errors.New 사용
}

func main() {
	// error 처리
	// Errorf를 이용한 error 처리

	a, err := notZero(1)

	if err != nil {
		// log.Fatal(err)
		log.Fatal(err.Error())
	}

	fmt.Println("ex1 : ", a)

	b, err := notZero(0)

	if err != nil {
		// log.Fatal(err)
		log.Fatal(err.Error())
	}

	fmt.Println("ex2 : ", b)

	fmt.Println("End Error Handling")

}
