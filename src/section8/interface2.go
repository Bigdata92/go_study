// interface basic(2)

package main

import "fmt"

type Dog struct {
	name   string
	weight int
}

// bite method
func (d Dog) bite() {
	fmt.Println(d.name, "bites!")
}

// 동물의 행동 interface 선언
type Behavior interface {
	bite() // 선언만 함
}

func main() {
	// interface 구현 예제
	// ex1
	dog1 := Dog{"poll", 10}

	var inter1 Behavior
	inter1 = dog1
	inter1.bite()
	// dog1.bite()

	// ex2
	dog2 := Dog{"marry", 12}
	inter2 := Behavior(dog2) // 이걸 더 많이 사용
	inter2.bite()

	// ex3
	inters := []Behavior{dog1, dog2}

	// index 형태로 실행
	for idx, _ := range inters {
		inters[idx].bite()
	}

	// value 형태로 실행(interface)
	for _, val := range inters {
		val.bite()
	}

}
