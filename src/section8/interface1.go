// interface basic(1)

package main

import "fmt"

type test interface {
}

func main() {

	// interface
	// object의 동작을 표현, 골격
	// 단순히 동작에 대한 방법만 표시
	// 추상화 제공
	// interface의 method를 구현한 type : interface로 사용 가능
	// Go 언어를 유연하게 사용가능
	// ducktyping: Go 언어의 독창적인 특성
	// interface -> java(전략패턴, 템플릿메소드, 팩토리메소드패턴, 어댑터패턴 ...)
	// design pattern 측면에서 client 입장 -> method의 구체적인 class를 몰라도 interface에 정의된 method를 사용하는 object 보장
	// class 간의 결합도 감소 -> 유지보수성 향상,기능 추가의 용이성 -> 독립적인 프로그래밍 가능

	// ex1
	/*
		type interface명 interface {
			method1() 반환 값(타입 형)
			method2() // 반환 값 없을 경우
		}

	*/

	var t test
	fmt.Println("ex1 : ", t) // empty interface의 경우 Nil return

}
