// file write(2)

package main

import (
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
	// file write
	// csv file write ex1
	// pkg 저장소를 통해서 excel 등 다양한 파일 형식 쓰기, 읽기 가능
	// pkg github 주소 : https://github.com/tealeg/xlsx
	// bufio : file 용량 클 경우, buffer 사용 권장

	// make file
	file, err := os.Create("test_write.csv")
	errCheck1(err)

	// resource 해제
	defer file.Close()

	// csv write 생성
	wr := csv.NewWriter(file)
	// wr := csv.NewWriter(bufio.NewWriter(file))	// big file

	// csv write content
	wr.Write([]string{"Kim", "4.8"})
	wr.Write([]string{"Lee", "4.2"})
	wr.Write([]string{"Park", "4.4"})
	wr.Write([]string{"Cho", "4.45"})
	wr.Write([]string{"Hong", "4.9"})

	wr.Flush() // buffer -> file로 write

	fi, err := file.Stat()
	errCheck1(err)

	fmt.Printf("CSV 쓰기 작업 후 파일 크기(%d bytes)\n", fi.Size())
	fmt.Println("CSV 파일명 : ", fi.Name())
	fmt.Println("OS 파일권한 : ", fi.Mode())

}
