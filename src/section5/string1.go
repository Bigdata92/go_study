// data type : 문자열(1)

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// 문자열
	// "", ``
	// golang : 문자 char type not exist -> rune(int32)로 문자 코드 값
	// 문자 : ''로 작성
	// 자주 사용하는 escape : \\, \', \", \a(소리), \b(백스페이스), \f(쪽바꿈), \n(줄바꿈), \r(복귀), \t(tab)

	// ex1
	var str1 string = "c:\\go_study\\src\\" // -> c:\go_study\src\
	str2 := `c:\go_study\src\`

	fmt.Println("ex1 : ", str1)
	fmt.Println("ex1 : ", str2)

	// ex2
	var str3 string = "Hello, world!"
	var str4 string = "안녕하세요."
	var str5 string = "\ud55c\uae00" // 4bite 한, 글 표현

	fmt.Println()
	fmt.Println("ex2 : ", str3)
	fmt.Println("ex2 : ", str4)
	fmt.Println("ex2 : ", str5)

	// ex3
	// 길이(byte 수)
	fmt.Println("ex3 : ", len(str3))
	fmt.Println("ex3 : ", len(str4))

	// ex4
	// 길이(실제 길이)
	fmt.Println("ex3 : ", utf8.RuneCountInString(str3))
	fmt.Println("ex3 : ", utf8.RuneCountInString(str4))
	fmt.Println("ex3 : ", len([]rune(str4)))	// len으로 실제 길이 구하는법
}
