package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	duplicateVoewl bool = true
	removeVowel    bool = false
)

func randBool() bool {
	return rand.Intn(2) == 0
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		word := []byte(s.Text())

		if randBool() { // 0 부터 1 사이에서의 값중 0일 떄 동작
			var VI int = -1
			for i, char := range word {
				switch char {
				case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
					if randBool() {
						VI = i
					}
				}
			}

			if VI >= 0 {
				switch randBool() {
				case duplicateVoewl:
					word = append(word[:VI+1], word[VI:]...)
				case removeVowel:
					word = append(word[:VI], word[VI+1:]...)
				}
			}
		}

		fmt.Println(string(word))
	}
}
