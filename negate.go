package main

// 포인터를 매개변수로 받아 함수를 이용해 포인터 주소에 있는 값 변경하기
import (
	"fmt"
)

func negate(b *bool) {
	*b = !*b
}

func main() {
	truth := true
	negate(&truth)
	fmt.Println(truth)
	lies := false
	negate(&lies)
	fmt.Println(lies)
}
