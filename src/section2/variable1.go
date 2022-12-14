/*
go 장점
1. 간결한 문법 및 단순화
2. 병행 프로그래밍 지원(핵심)
3. 정적 타입 및 동적 실행
4. 간편한 협업 지원
5. 컴파일 및 실행속도 빠름
6. 제네릭 및 예외 처리 미지원
7. 컨벤션 통일(fmt)
*/

// 변수1
package main

import "fmt"

func main() {
	// 기본 초기화
	// 정수타입 : 0, 실수(소수점) : 0.0, 문자열 : "", Boolean : True, False
	// 변수명 : 숫자 첫글자 X, 대소문자 구분 O, 문자, 숫자, 밑줄, 특수기호 사용 가능
	// 변수 및 상수 : 함수 내외 사용 가능
	var a int
	var b string
	var c, d, e int
	var f, g, h int = 1, 2, 3
	var i float32 = 11.4
	var j string = "Hi! Golang!"
	var k = 4.74 // 선언 동시 초기화
	var l = "Hi! Seoul!"
	var m = true

	a = 4
	b = "Hello Go!"
	e = 77

	fmt.Println("a : ", a)
	fmt.Println("b : ", b)
	fmt.Println("c : ", c)
	fmt.Println("d : ", d)
	fmt.Println("e : ", e)
	fmt.Println("f : ", f)
	fmt.Println("g : ", g)
	fmt.Println("h : ", h)
	fmt.Println("i : ", i)
	fmt.Println("j : ", j)
	fmt.Println("k : ", k)
	fmt.Println("l : ", l)
	fmt.Println("m : ", m)

}
