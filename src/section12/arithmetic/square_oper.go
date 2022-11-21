// 두 숫자의 제곱 연산 제공 pkg(2)
package arithmetic

// x, y 제곱의 합을 return
func (o *Numbers) SquarePlus() int {
	return (o.X * o.X) + (o.Y * o.Y)
}

// x, y 제곱의 차를 return

func (o *Numbers) SquareMinus() int {
	return (o.X * o.X) - (o.Y * o.Y)
}
