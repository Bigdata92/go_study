// Panic & Recover(1)

package main

import "fmt"

// import "log"

func main() {
	// Go panic 함수
	// 사용자가 에러 발생 기능
	// Panic 함수는 호출 즉시, 해당 method 즉시 중지, defer 함수를 호출하고 자기자신을 호출한 곳으로 리턴
	// Runtime 이외에 user가 code 흐름에 따라 error 발생시킬때 중요!
	// 문법적인 에러 x, 논리적인 코드 흐름에 따른 에러 발생 처리 가능

	fmt.Println("Start Main")
	panic("Error occurred : user stopped!") // 방법1
	//log.panic("Error occurred : user stopped!")	// 방법2
	fmt.Println("End Main") // 실행 불가
}
