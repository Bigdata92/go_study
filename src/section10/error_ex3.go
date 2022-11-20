// Go error 처리 고급(3)

package main

import (
	"fmt"
	"log"
	"math"
	"time"
)

// 예외(에러) 처리 구조체
type PowError struct {
	time    time.Time // error 발생 시간
	value   float64   // param
	message string    // error message
}

func (e PowError) Error() string {
	return fmt.Sprintf("[%v]Error - Input Value(value : %g) - %s", e.time, e.value, e.message)
}

func Power(f, i float64) (float64, error) {
	if f == 0 {
		return 0, PowError{time: time.Now(), value: f, message: "0은 사용할 수 없습니다."}
	}
	if math.IsNaN(f) {
		return 0, PowError{time: time.Now(), value: f, message: "숫자가 아닙니다."}
	}
	if math.IsNaN(i) {
		return 0, PowError{time: time.Now(), value: i, message: "숫자가 아닙니다."}
	}
	return math.Pow(f, i), nil
}

func main() {
	// error 처리 고급
	// error 타입 아닌 경우 error 처리 방법
	// Error method를 구현해서 사용자 정의 에러 처리 예제 심화
	// 구조첼르 사용해서 세부적인 정보 출력

	// ex1
	v, err := Power(10, 3) // 정상 처리
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("ex1 : ", v)

	// ex2
	t, err := Power(0, 3) // 예외 발생
	if err != nil {
		// log.Fatal(err)
		log.Fatal(err.Error())
		fmt.Println(err.(PowError).value)
		fmt.Println(err.(PowError).message)
		fmt.Println(err.(PowError).time)
	}

	fmt.Println("ex2 : ", t)
}
