// 사용자 정의 타입(4)

package main

import "fmt"

type shoppingBasket struct{cnt, price int}

func (b shoppingBasket) purchase() int {
	return b.cnt * b.price
}

// 원본 수정(참조 전달 형식)
func (b *shoppingBasket) rePurchaseP(cnt, price int) {
	b.cnt += cnt
	b.price += price
}

// 원본 수정x(값 전달 형식)
func (b shoppingBasket) rePurchaseD(cnt, price int) {
	b.cnt += cnt
	b.price += price
}

func main() {
	// receiver(struct method) 전달(value, ref) 형식
	// 함수는 기본적으로 값 호출 -> 변수의 값이 복사 후 내부 전달(원본 수정x) -> map, slice등은 참조 전달
	// receiver(struct)도 포인터를 활용해서 method 내에서 원본 수정

	// ex1
	bs1 := shoppingBasket{3, 5000}
	fmt.Println("ex1(totPrice) : ", bs1.purchase())

	// 참조 전달(원본 값 수정)
	bs1.rePurchaseP(7, 5000)
	fmt.Println("ex2(totPrice) : ", bs1.purchase())
	// 값 전달(원본 값 수정x)
	bs1.rePurchaseP(10, 0)
	fmt.Println("ex3(totPrice) : ", bs1.purchase())
}