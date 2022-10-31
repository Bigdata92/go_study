package main

import "fmt"

func main() {
	const (
		Jan = iota + 1
		Feb 
		Mar 
		Apr 
		May 
		Jun 
	)

	fmt.Println(Jan)
	fmt.Println(Feb)
	fmt.Println(Mar)
	fmt.Println(Apr)
	fmt.Println(May)
	fmt.Println(Jun)
}