package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func cardConv(x, r int) []string {
	// 정수 x를 r 진수로 변환한 뒤 그 수를 나타내는 문자열을 반환

	var d []string // 변환 후 문자열
	dchar := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	fmt.Println(x, r)
	for x > 0 {
		d = append(d, dchar[x%r]) // 해당하는 문자를 꺼내 결합
		x = int(x / r)
	}

	// 배열의 원소를 역순으로 정렬 합니다.
	n := len(d)
	half := int(n / 2)
	for i := 0; i < half; i++ {
		d[i], d[n-1-i] = d[n-1-i], d[i]
	}

	return d
}

func main() {
	fmt.Println("10진수를 n진수로 변환합니다.")
	var no, cd int
	for {
		// 음이 아닌 정수를 입력받음
		for {
			fmt.Println("변환할 값으로 음이 아닌 정수를 입력하세요.: ")
			reader := bufio.NewReader(os.Stdin)
			input, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}
			input = strings.TrimSpace(input)

			no, err = strconv.Atoi(input)
			if err != nil {
				log.Fatal(err)
			}

			if no > 0 {
				break
			}
		}
		// 2~36진수의 정수값을 입력받음
		for {
			fmt.Println("어떤 진수로 변환할까요?: ")
			reader := bufio.NewReader(os.Stdin)
			input, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}
			input = strings.TrimSpace(input)

			cd, err = strconv.Atoi(input)
			if err != nil {
				log.Fatal(err)
			}

			if 2 <= cd && cd <= 36 {
				break
			}
		}

		fmt.Println(no, "는", cd, "진수로", cardConv(no, cd), "입니다.")

		// 추가 진행여부를 묻습니다.
		fmt.Println("한 번 더 변환할까요?(Y ... 예/N ... 아니오): ")
		reader2 := bufio.NewReader(os.Stdin)
		retry, err := reader2.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		retry = strings.TrimSpace(retry)
		if (retry == "n") || (retry == "N") {
			break
		}
	}
}
