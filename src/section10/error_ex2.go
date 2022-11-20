// Go error 처리 고급(2)

package main

import (
	"fmt"
	"math"
)

// f의 i제곱 구하는 함수
func Power(f float64, i float64) (float64, error) {
	if f == 0 {
		return 0, fmt.Errorf("(&g)는 사용 불가능 합니다.", f)
	}
	return math.Pow(f, i), nil // 제곱, nil return
}

func main() {
	// error 처리 고급
	// Go error pkg New method 사용 -> 사용자 에러 처리 생성

	// ex1
	if f, err := Power(7, 3); err != nil {
		//fmt.Printf("Error Message : %s\n", err)
		fmt.Printf("Error Message : %s\n", err.Error())
	} else {
		fmt.Println("ex1 : ", f)
	}

	// ex2
	if f, err := Power(0, 3); err != nil {
		// fmt.Printf("Error Message : %s\n", err)
		fmt.Printf("Error Message : %s\n", err.Error())
	} else {
		fmt.Println("ex2 : ", f)
	}
}
