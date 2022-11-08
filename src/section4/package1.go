package main

// 선언 방법1
// import "fmt"
// import "os"

// 선언 방법2
import (
	"fmt"
	"os"
)

func main() {
	// package : code 구조화(module화) 및 재사용
	// 응집도, 결합도
	// Go : package 단위의 독립적, 작은 단위로 개발 -> 작은 package를 결합해서 프로그램밍을 작성 할 것을 권고
	// package 이름 = directory 이름
	// 같은 package 내 -> source file들은 directory 명을 package 명으로 사용
	// naming rule : 소문자 - private, 대문자 - public
	// Go : main package는 특별히 인식 -> compiler 공유 lib x -> prog start pt

	// ex1
	var name string

	fmt.Println("이름은? : ")
	fmt.Scanf("%s", &name)
	fmt.Fprintf(os.Stdout, "Hi! %s\n", name)

}
