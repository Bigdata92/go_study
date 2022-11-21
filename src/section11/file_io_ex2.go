// file I/O(2)

package main

import (
	"bufio"
	"fmt"
	"os"
)

func errCheck(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// file read, buffer 사용 write -> bufio pkg 활용
	// ioutil, bufio 등은 io.Reader, io.Writer interface를 구현함 -> 즉, Writer, Read method를 사용 가능
	/*
		type Reader interface {
			Read(p []byte) (n int, err error)
		}
		type Writer interface {
			Read(p []byte) (n int, err error)
		}
	*/
	// 즉, bufio의 NewReader, NewWriter를 통해서 객체 반환 후 method 사용

	// bufio(Buffered io) pkg

	// open file
	// 2nd param check

	/*
		상태
		a ------> a
		b ------> ab
		c ------> abc
		d ------> abcd
		e ------> e    ------> abcd
		f ------> ef    ------> abcd
		g ------> efg    ------> abcd
		h ------> efgh    ------> abcd
		i ------> i    -------> abcdefgh

	*/

	file, err := os.OpenFile("test_write2.txt", os.O_CREATE|os.O_RDWR, os.FileMode(0777))

	// bufio file write ex
	wt := bufio.NewWriter(file) // Writer return (buffer 사용)
	wt.WriteString("Hello Golang!\n File Write Test1!\n")
	wt.Write([]byte("Hello Golang!\n File Write Test2!\n"))

	// error check
	errCheck(err)

	// buffer info print
	fmt.Printf("사용한 Buffer Size (%d bytes)\n", wt.Buffered())
	fmt.Printf("남은 Buffer Size (%d bytes)\n", wt.Available())
	fmt.Printf("전체 Buffer Size (%d bytes)\n", wt.Size())

	wt.Flush() // buffer를 비우고 반영(buffer의 내용을 disk에 기록)
	fmt.Println("쓰기 작업 완료\n")
	fmt.Println("====================================")

	rt := bufio.NewReader(file) // Reader return
	fi, err := file.Stat()
	errCheck(err)

	b := make([]byte, fi.Size())

	fmt.Println("파일 정보 출력 : ", fi)
	fmt.Println("파일 이름 : ", fi.Name())
	fmt.Println("파일 크기 : ", fi.Size())
	fmt.Println("파일 수정 시간 : ", fi.ModTime())

	fmt.Println("==============================")

	file.Seek(0, os.SEEK_SET)
	data, _ := rt.Read(b) // Read(ReadLine, ReadByte, ReadBytes 등)
	// rt.Read(b)

	fmt.Printf("전체 Buffer Size : (%d bytes)\n", rt.Size())
	fmt.Printf("읽기 작업 완료 : (%d bytes)\n", data)
	fmt.Println("=====================================")
	fmt.Println(string(b))
	fmt.Println("=====================================")

	defer file.Close()

}
