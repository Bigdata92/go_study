// file read(1)

package main

import (
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
	// Open : 기존 file open
	// Close : resource close
	// Read, ReadAt : file read
	// 각 OS 권한 주의(error msg check)
	// 예외 처리 정말 중요!
	// 탐색 seek 중요
	// https://golang.org/pkg/os

	// file read ex
	file, err := os.Open("sample.txt")
	errCheck1(err)

	// read ex 1
	fileInfo, err := file.Stat() // file size check
	errCheck2(err)

	fd1 := make([]byte, fileInfo.Size()) // slice에 읽은 내용 담는다.
	ct1, err := file.Read(fd1)

	fmt.Println("file info print(1) : ", fileInfo, "\n")
	fmt.Println("file name (2) : ", fileInfo.Name(), "\n")
	fmt.Println("file name (3) : ", fileInfo.Size(), "\n")
	fmt.Println("file modify time(4) : ", fileInfo.ModTime(), "\n")
	fmt.Printf("read (1) complete (%d bytes)\n\n", ct1)
	// fmt.Println(fd1)	// type change x
	fmt.Println(string(fd1)) // type change o

	fmt.Println("====================================================")

	// read ex2(seek(offset))
	o1, err := file.Seek(20, 0) // 0: 처음위치, 1: 현재위치, 2: 마지막위치
	errCheck1(err)

	fd2 := make([]byte, 20)
	ct2, err := file.Read(fd2)
	errCheck1(err)

	fmt.Printf("read (2) complete (%d bytes) (%d ret)\n\n", ct2, o1)
	fmt.Println(string(fd2), "\n")
	fmt.Println("====================================================")

	// read ex3
	o2, err := file.Seek(0, 0)
	errCheck1(err)
	fd3 := make([]byte, 50)
	ct3, err := file.ReadAt(fd3, 8) // offset 위치부터 읽어온다.
	errCheck1(err)

	fmt.Printf("read (3) complete (%d bytes) (%d ret)\n\n", ct3, o2)
	fmt.Println(string(fd3), "\n")
	fmt.Println("====================================================")

	defer file.Close()

}
