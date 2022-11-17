// interface basic(5)

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

func printValue(s interface{}) {
	fmt.Println("ex1 : ", s)
}

func main() {

	// interface 활용(빈 interface)
	// 함수내에서 어떠한 type 이라도 유연하게 param으로 받을 수 있다.(만능) -> all type 지정 가능

	dog := Dog{"poll", 10}
	cat := Cat{"bob", 5}

	printValue(dog)
	printValue(cat)
	printValue(15)
	printValue("Animal")
	printValue(25.5)
	printValue([]Dog{})
	printValue([5]Dog{})

}
