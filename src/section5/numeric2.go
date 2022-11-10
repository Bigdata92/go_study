// data type : Numeric(2)

package main

import "fmt"

func main() {
	// data type : 숫자형
	// 정수형 문자 출력
	// ex1
	// ascii(영문)

	var char1 byte = 72
	var char2 byte = 0110
	var char3 byte = 0x48

	// ex2
	// unicode(한글)
	var char4 rune = 50556   // unicode
	var char5 rune = 0142574 // 44032(8진수)
	var char6 rune = 0xC57C  // 44032(16진수)

	fmt.Printf("%c %c %c\n", char1, char2, char3)	// %c %d %o %x
	fmt.Printf("%d %d %d\n", char1, char2, char3)
	fmt.Printf("%d %o %x\n", char1, char2, char3)

	fmt.Printf("%c %c %c\n", char4, char5, char6)
	fmt.Printf("%d %d %d\n", char4, char5, char6)
	fmt.Printf("%d %o %x\n", char4, char5, char6)
}
