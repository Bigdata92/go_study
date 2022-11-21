// 사용자 pkg 작성 및 문서화 ex
package main

import (
	"fmt"
	oper "section12/arithmetic" // alias
)

func main() {
	// 기준은 GOPATH/src
	// pkg fold명(dir명) 명확하게 지정
	// pkg file의 pkg 이름으로 사용 -> 길면 alias
	// pkg main을 제외하고 pkg 문서에 등록
	// 지금까지 우리는 pkg 사용해 왔다.
	// 기본적으로 GOROOT의 pkg 검색 -> GOPATH pkg 검색
	// go install 명령어 실행 -> GOPATH/pkg 등록(문서화)
	// godoc -http:6060(임의의 port) -> pkg 이동 -> 본인 pkg method 및 주석 확인

	// pkg ex(사칙연산)
	nums := oper.Numbers{100, 10}
	fmt.Println("Pkg Used(1) : ", nums.Plus())
	fmt.Println("Pkg Used(2) : ", nums.Minus())
	fmt.Println("Pkg Used(3) : ", nums.multiply())
	fmt.Println("Pkg Used(4) : ", nums.Divide())
	fmt.Println("Pkg Used(5) : ", nums.SquarePlus())
	fmt.Println("Pkg Used(6) : ", nums.SquareMinus())

}
