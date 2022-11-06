package main

import "fmt"

func main() {
	// ex1
Loop1:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 2 && j == 4 {
				// for문 전체 벗어남
				break Loop1
			}
			fmt.Println("ex1 : ", i, j)
		}
	}

	// ex2
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("ex2 : ", i)
	}

	fmt.Println("======================")

Loop2:
	// Loop 문 안에는 for문 있어야(없으면 error)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 2 {
				continue Loop2
			}
			fmt.Println("ex3 : ", i, j)
		}
	}

}
