package main

import "fmt"

func main() {
	// ex1
	a := 30 / 15
	switch a {
	case 2, 4, 6: // i가 2, 4, 6인 경우
		fmt.Println("a -> ", a, "는 짝수")
	case 1, 3, 5: // i가 1, 3, 5인 경우
		fmt.Println("a -> ", a, "는 홀수")
	}

	// ex2
	switch e := "go"; e {
	case "java":
		fmt.Println("Java!")
		fallthrough
	case "go":
		fmt.Println("go!")
		fallthrough // 조건문 실행 뒤, case 실행 할 때
	case "python":
		fmt.Println("python!")
	case "ruby":
		fmt.Println("ruby!")
	case "c":
		fmt.Println("c!")
		// fallthrough - 마지막 사용 x
	}

}
