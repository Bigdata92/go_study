// Go error basic(1)

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// error 처리: SW 품질 향상 -> 유형코드 및 에러정보 등을 남기는 것
	// Go 에서는 기본적으로 error pkg에서 error interface 제공
	// type error interface { Error() string }
	// 즉, Error method 구현시 사용자 정의 에러 처리 제작 가능
	// 기본적으로 method마다 return type 2개(return val, error)
	// 주로 Errorf(error type return) method, Fatal(program 종료) method를 통해 error 출력

	// 기본적인 method error 처리 ex
	f, err := os.Open("unnamedfile") // error
	if err != nil {                  // error 발생 경우
		log.Fatal(err.Error()) // 방법1
		// log.Fatal(err)	// 방법2
	}
	fmt.Println("================")
	fmt.Println("ex1 : ", f.Name())

}
