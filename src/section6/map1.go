// data type : map(1)

package main

import "fmt"

func main() {

	// map : hashtable, dictionary(python)
	// key - value로 data 저장
	// reference type(참조 값 전달) -> 비교 연산자 사용 불가
	// 특징 : 참조타입(key)로 사용 불가능, 값(value)로 모든 타입 사용 가능
	// make 함수 및 축약(리터럴)로 초기화 가능
	// 순서 x

	// ex1
	var map1 map[string] int = make(map[string]int)	// 정석
	var map2 = make(map[string]int)	// 자료형 생략
	map3 := make(map[string]int)	// 리터럴 형(주로 이렇게 사용)

	fmt.Println("ex1 : ", map1)
	fmt.Println("ex1 : ", map2)
	fmt.Println("ex1 : ", map3)
	fmt.Println()

	// ex2
	map4 := map[string]int{}	// Json 형태 (a = {"a" : 1, "b" : 2})
	map4["apple"] = 25
	map4["banana"] = 40
	map4["orange"] = 33

	map5 := map[string]int{
		"apple" : 15,
		"banana" : 40,
		"orange" : 23,
	}

	map6 := make(map[string]int, 10)
	map6["apple"] = 25
	map6["banana"] = 40
	map6["orange"] = 33

	fmt.Println("ex2 : ", map4)
	fmt.Println("ex2 : ", map5)
	fmt.Println("ex2 : ", map6)
	fmt.Println("ex2 : ", map6["apple"])
	fmt.Println("ex2 : ", map6["orange"])

}