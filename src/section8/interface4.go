// interface basic(4)

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

func (d Dog) run() {
	fmt.Println(d.name, " : Dog is running!")
}

func (c Cat) run() {
	fmt.Println(c.name, " : Cat is running!")
}

func act(animal interface{ run() }) { // 익명 interface(type 정의x)
	animal.run()
}

func main() {

	// 익명 interface 사용 예제(즉시 선언 후 사용)

	// ex1
	dog := Dog{"poll", 10}
	cat := Cat{"bob", 5}

	// dog 행동 실행
	act(dog)

	// cat 행동 실행
	act(cat)

}
