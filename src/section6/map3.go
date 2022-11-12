// data type : map(3)

package main

import "fmt"

func main() {

	// map value 변경 및 삭제

	// ex1
	map1 := map[string]string{
		"daum":   "http://daum.net",
		"naver":  "http://naver.com",
		"google": "http://google.com",
		"home1":  "http://test1.com",
	}

	fmt.Println("ex1 : ", map1)

	map1["home2"] = "http://test2.com"	// add
	fmt.Println("ex1 : ", map1)

	map1["home2"] = "http://test2-2.com"	// update
	fmt.Println("ex1 : ", map1)

	// ex2(delete)
	delete(map1, "home2")
	fmt.Println("ex1 : ", map1)

	delete(map1, "google")
	fmt.Println("ex1 : ", map1)
}
