// file I/O(1)

package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func errCheck(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// file read, write -> ioutil pkg 활용
	// 편리하고 직관적으로 file read, write 가능
	// WriteFile(), ReadFile(), ReadAll() 등 사용 가능
	// https://golang.org/io/ioutil

	s := "Hello Golang!\n File Write Test!\n"

	// file mode(chmode, chown, chgrp) -> permisson
	// read : 4, write : 2, 실행 :
	// 소유자, 그룹, 기타 사용자 순서(777)
	// https://golang.org/pkg/os/#FileMode

	err := ioutil.WriteFile("test_write1.txt", []byte(s), os.FileMode(0644))
	errCheck(err)

	data, err := ioutil.ReadFile("sample.txt")
	errCheck(err)

	fmt.Println("===================================")
	fmt.Println(string(data))
	fmt.Println("===================================")
}
