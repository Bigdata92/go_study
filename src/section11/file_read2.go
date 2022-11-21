// file read(2)

package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

// error check way 1
func errCheck1(e error) {
	if e != nil {
		panic(e)
	}
}

// error check way 2
func errCheck2(e error) {
	if e != nil {
		fmt.Println(e)
		return
	}
}

func main() {
	// file read
	// csv file read ex

	// make file
	file, err := os.Open("sample.csv")
	errCheck1(err)

	// resource 해제
	defer file.Close()

	// CSV Reader 생성
	// rr := csv.NewReader(file)
	rr := csv.NewReader(bufio.NewReader(file)) // 권장

	// csv 내용 읽기(1)

	row1, err1 := rr.Read() // 1개의 Row 단위로 읽기(사용 잘 안함)
	errCheck1(err1)
	row2, err2 := rr.Read()
	errCheck2(err2)
	fmt.Println("CSV Read Example")
	//fmt.Println(row)
	fmt.Println(row1[0], row1[1], row1[7], row1[1:5])
	fmt.Println(row2[0], row2[1], row2[7], row2[1:5])
	fmt.Println("============================")

	// csv 내용 읽기(2)
	rows, err := rr.ReadAll() // 전체 Row read(자주 사용)
	errCheck2(err)
	fmt.Println("CSV ReadAll Example")
	fmt.Println(rows[5][1], " : ", rows[2][1], " : ", rows[6][1])
	fmt.Println(rows)

	// Row 단위로 csv file read

	for i, row := range rows {
		// fmt.Println(i, row)
		for j := range row {
			fmt.Printf("%s ", rows[i][j])
		}
		fmt.Println()
	}
}
