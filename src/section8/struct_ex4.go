// 구조체 심화(4)

package main

import "fmt"

type Employee struct {
	name   string
	salary float64
	bonus  float64
}

func (e Employee) Calculate() float64 {
	return e.salary + e.bonus
}

type Executives struct {
	Employee     // is a 관계
	specialBonus float64
}

func main() {
	// 구조체 임베디드 패턴
	// 다른 관점으로 method를 재 사용하는 장점 제공
	// 상속을 허용 x Go 언어에서 method 상속 활용을 위한 패턴

	// ex1
	// 직원
	ep1 := Employee{"kim", 2000000, 150000}
	ep2 := Employee{"park", 1500000, 200000}

	// 임원
	ex := Executives{
		Employee{"lee", 5000000, 1000000},
		1000000,
	}

	fmt.Println("ex1 : ", int(ep1.Calculate()))
	fmt.Println("ex1 : ", int(ep2.Calculate()))

	// Employee 구조체 통해서 method 호출
	fmt.Println("ex2 : ", int(ex.Calculate()+ex.specialBonus))

}
