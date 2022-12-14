package main

import (
	"fmt"
	"math"
)

func main() {
	// 숫자 연산(산술, 비교)
	// 타입이 같아야 가능(정수, 실수 x)
	// 다른 타입끼리는 반드시 형 변환 후 연산
	// 형 변환 x -> error 발생
	// +, -, *, %, /, <<, >>, &, ^

	// ex1
	var n1 uint8 = math.MaxUint8
	var n2 uint16 = math.MaxUint16
	var n3 uint32 = math.MaxUint32
	var n4 uint64 = math.MaxUint64

	fmt.Println("ex1 : ", n1)
	fmt.Println("ex1 : ", n2)
	fmt.Println("ex1 : ", n3)
	fmt.Println("ex1 : ", n4)
	fmt.Println("ex1 : ", math.MaxInt8)
	fmt.Println("ex1 : ", math.MaxInt16)
	fmt.Println("ex1 : ", math.MaxInt32)
	fmt.Println("ex1 : ", math.MaxInt64)
	fmt.Println("ex1 : ", math.MaxFloat32)
	fmt.Println("ex1 : ", math.MaxFloat64)

	n5 := 100000       // int
	n6 := int16(10000) // int16
	n7 := uint8(100)

	// fmt.Println("ex2 : ", n5+n6)	// error
	fmt.Println("ex2 : ", n5+int(n6))
	fmt.Println("ex2 : ", n7)
	//fmt.Println("ex2 : ", n6+n7)
	fmt.Println("ex2 : ", n6+int16(n7))
	fmt.Println("ex2 : ", n6 > int16(n7))
	fmt.Println("ex2 : ", n6-int16(n7) > 5000)

}
