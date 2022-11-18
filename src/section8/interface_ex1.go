// interface 고급(1)

package main

import "fmt"

func main() {
	// empty interface 활용 : func param, return val, struct field 등 사용 가능 -> all type ok
	// all type을 나타내기 위해 empty interface 사용
	// 동적 타입 으로 생각 하면 쉽다. (타 언어의 obejct type)

	// ex1
	var a interface{}

	printContents(a)

	a = 7.5
	printContents(a)

	a = "Golang!"
	printContents(a)

	a = true
	printContents(a)

	a = nil
	printContents(a)

}

func printContents(v interface{}) {
	fmt.Printf("Type : (%T) ", v) // original type
	fmt.Println("ex1 : ", v)
}
