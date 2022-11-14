// 사용자 정의 타입(1)

package main

import "fmt"

type Car struct {
	name string
	color string
	price int64
	tax int64
}

// 일반 method
func Price(c Car) int64 {
	return c.price + c.tax
}

// 구조체 <-> method binding
func (c Car) Price() int64 {
	return c.price + c.tax
}

func main() {
	// Go -> 객체 지향 타입을 구조체로 정의 (class, 상속 개념 x)
	// 객체 지향 -> class(속성 : 멤버변수, 기능(상태 : 메소드)) : code 재사용성, 관리 용이, 신뢰성 높은 프로그래밍
	// Go - 객체 지향 언어
	// Go는 전형적인 객체지향 특징x, interface -> 다형성 지원, 구조체를 class 형태로 코딩 가능
	// 객체지향의 기본 개념 -> Go에서 포함 하고 있다. -> 객체 지향 프로그래밍 언어
	// 상태, 메소드 분리해서 정의(결합성x)
	// 사용자 정의 타입 : 구조체, interface, 기본 타입(int, float, string), 함수
	// 구조체와 메소드 연결을 통해 타 언어의 class 처럼 사용 가능(객체 지향)

	// ex1
	bmw := Car{name: "520d", price: 50000000, color: "white", tax: 500}
	benz := Car{name: "220d", price: 60000000, color: "white", tax: 600}

	fmt.Println("ex1 : ", bmw, &bmw)
	fmt.Println("ex1 : ", benz, &benz)

	//fmt.Println("ex2 : ", Price(bmw))
	//fmt.Println("ex2 : ", Price(benz))

	fmt.Println("ex3 : ", bmw.Price())
	fmt.Println("ex3 : ", benz.Price())

	fmt.Println("ex4 : ", &bmw == &benz)
}