// interface basic(3)

package main

import "fmt"

type Dog struct {
	name   string
	weight int
}

type Cat struct {
	name   string
	weight int
}

// struct Dog method 구현
func (d Dog) bite() {
	fmt.Println(d.name, " : Dog bites!")
}

func (d Dog) sound() {
	fmt.Println(d.name, " : Dog barks!")
}

func (d Dog) run() {
	fmt.Println(d.name, " : Dog is running!")
}

func (c Cat) bite() {
	fmt.Println(c.name, " : Cat scratch!")
}

func (c Cat) sound() {
	fmt.Println(c.name, " : Cat crys!")
}

func (c Cat) run() {
	fmt.Println(c.name, " : Cat is running!")
}

// 동물의 행동 interface 선언
type Behavior interface {
	bite()
	sound()
	run()
}

// interface를 param 받는다.
func act(animal Behavior) {
	animal.bite()
	animal.sound()
	animal.run()
}

func main() {

	// interface 구현 ex
	// interface 규격화 역할 이해
	// interface에 정의된 method 사용 유도
	// code의 가독성 및 유지보수 증가

	// ducktyping ex
	// ducktyping : 구조체 및 변수의 값이나 타입 상관 x, 구현된 method 로만 판단
	// Go의 ducktyping 특징 : 오리처럼 걷고, 소리, 헤엄 등 행동 같으면 오리라고 볼 수 있다.

	// ex1
	dog := Dog{"poll", 10}
	cat := Cat{"bob", 5}

	// dog 행동 실행
	act(dog)

	// cat 행동 실행
	act(cat)

}
