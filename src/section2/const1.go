package main

import "fmt"

func main() {
	// 상수
	// const 사용 초기화, 한 번 선언 후 값 변경 금지, 고정 된 값 관리용
	const a string = "Test1"
	const b = "Test2"
	const c int32 = 10 * 10
	// const d = getHeight()	// 함수 값 바뀔 수 o -> const에 함수 사용 x
	const e = 35.6
	const f = false

	/* 
		error 발생(선언과 동시에 초기화 해야 됨)
		cost g string
		g = "Test3"
	*/

	fmt.Println("a : ", a, "b : ", b, "c : ", c, "e : ", e, "f : ", f,)

}