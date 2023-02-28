package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
)

// 적절한 도메인 탐색 프로그램
var tlds = []string{"com", "net"}

const allowedChars = "abcdefghijklmnopqrstuvwxyz0123456789_"

func main() {
	rand.Seed(time.Now().UTC().UnixNano()) // 랜덤 값을 위한 설정

	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		text := strings.ToLower(s.Text())

		var newText []rune

		for _, r := range text {
			if unicode.IsSpace(r) { // 만약 공백인 경우
				r = '-'
			}

			if !strings.ContainsRune(allowedChars, r) { // a의 문자열 값에 r이라는 char이 있는지
				continue
			}

			newText = append(newText, r)
		}

		fmt.Println(string(newText) + "." + tlds[rand.Intn(len(tlds))])
	}
}
