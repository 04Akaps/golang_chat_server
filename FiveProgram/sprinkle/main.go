package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// 이름을 찾는 프로그램
// -> 도메인 이름 찾기

const otherWord = "*"

var transforms = []string{
	otherWord,
	otherWord + "app",
	otherWord + "site",
	otherWord + "time",
	"get" + otherWord,
	"go" + otherWord,
	"lets" + otherWord,
	otherWord + "hq",
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano()) // 랜덤 시드 설정 -> 시작 할떄마다 값이 바뀜

	s := bufio.NewScanner(os.Stdin) // Scanner 객체 생성, os.Stdin에서 입력을 읽도록 지시
	// 이후 표준 출력에 사용

	for s.Scan() { // os.Stdin에서 한줄을 읽어 옵니다.
		// 반환값이 bool이기 떄문에 false가 뜰떄까지 반복
		t := transforms[rand.Intn(len(transforms))] // o 부터 인자 사이에서 랜덤한 숫자를 반환
		fmt.Println(strings.Replace(t, otherWord, s.Text(), -1))
		// replace는 t의 문자열 값에서, otherwork의 값을 s.Text()로 바꾼다는 의미
		// 마지막은 여러개가 있을 떄 몇번 반복하냐의 의미이며, -1은 모든 값을 바꾸겠다는 의미
	}
}
