// lib 접근제어(1)
package lib // fold name

import "fmt"

func init() {
	fmt.Println("lib pkg > init start!")
}

func CheckNum1(c int32) bool {
	return c > 100
}
