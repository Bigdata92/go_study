package main

import "fmt"

func main() {
	fmt.Println("hello world!")
}

// 실행방법1. go run helloworld.go
// unit test, exe 파일 생성x 

// 실행방법2-1. go build helloworld.go
// 실행방법2-2. helloworld.exe
// exe 파일 생성o

// 실행방법3. go install
// 파일 소속된 프로젝트, 패키지 모아서 bin 폴더에 프로젝트명.exe로 생성(배포용)