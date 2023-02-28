package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"go_chat/FiveProgram/theasurus"
)

// 해당 코드를 사용하고자 한다면 적합한 API키가 필요합니다.
// 이후 ~/.bashsrc에 BHT_APIKEY라는 이름으로 export합니다.

func main() {
	// 동의어를 찾는 함수
	apiKey := os.Getenv("BHT_APIKEY")

	theasurus := &theasurus.BigHuge{APIKey: apiKey}

	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		word := s.Text()

		syns, err := theasurus.Synonyms(word)
		if err != nil {
			log.Fatalln("Error  "+word, err)
		}

		if len(syns) == 0 {
			log.Fatalln("not found any syns" + word)
		}

		for _, syn := range syns {
			fmt.Println(syn)
		}
	}
}
