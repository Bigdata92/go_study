// data type : map(2)

package main

import "fmt"

func main() {

	// map
	// map 조회 및 순회(iterator)

	// ex1
	map1 := map[string]string{
		"daum":   "http://daum.net",
		"naver":  "http://naver.com",
		"google": "http://google.com",
	}

	fmt.Println("ex1 : ", map1["daum"])
	fmt.Println("ex1 : ", map1["google"])
	fmt.Println()

	// ex2(순서 x -> random)
	for k, v := range map1 {
		fmt.Println("ex2 : ", k, v)
	}

	fmt.Println()

	for _, v := range map1 {
		fmt.Println("ex2 : ", v)
	}
}
