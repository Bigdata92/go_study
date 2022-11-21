// file write(1)

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
	// file write
	// Create : new file write and open
	// Close : resource close
	// Write, WriteString, WriteAt : file write
	// 각 OS 권한 주의(error msg check)
	// 예외 처리 중요!

	// file write
	file, err := os.Create("test_write.txt")
	errCheck1(err)

	// resource 해제
	defer file.Close()

	// write ex1(잘 안씀)
	s1 := []byte{48, 49, 50, 51, 52}
	n1, err := file.Write([]byte(s1)) // 문자열 -> byte slice 형으로 변환 후 쓰기
	errCheck2(err)

	fmt.Printf("write 1 complete (%d bytes) \n", n1)

	file.Sync() // write commit(Stable)!

	// ex2(자주 씀)
	s2 := "Hello Golang!\n File Write Test! - 1\n"
	n2, err := file.WriteString(s2)
	errCheck2(err)

	fmt.Printf("write 2 complete (%d bytes) \n", n2)
	file.Sync() // write commit(Stable)!

	s3 := "Test WriteAt! - 2\n"
	n3, err := file.WriteAt([]byte(s3), 50) // len(offset) 조절하면서 test
	errCheck1(err)

	fmt.Printf("write 3 complete (%d bytes) \n", n3)
	file.Sync()

	// ex4
	n4, err := file.WriteString("Hello Golang \n File Write Test! - 3\n")
	errCheck1(err)

	fmt.Printf("write 4 complete (%d bytes) \n", n4)
	file.Sync()
}
