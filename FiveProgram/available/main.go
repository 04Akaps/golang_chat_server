package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

// 나오는 도메인이 사용 가능한지를 체크하는 프로그램

func exists(domain string) (bool, error) {
	const whoisServer string = "com.whois-server.net"
	conn, err := net.Dial("tcp", whoisServer+":43") // 클라이언트가 서버에 통신 할 떄 사용 합니다.
	if err != nil {
		return false, err
	}

	defer conn.Close()

	conn.Write([]byte(domain + "rn")) // 기본적인 스펙을 지키기 위해서 rn을 write
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		if strings.Contains(strings.ToLower(scanner.Text()), "no match") {
			return false, nil
		}
	}

	return true, nil
}

var marks = map[bool]string{true: "O", false: "X"}

func main() {
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		domain := s.Text()
		fmt.Print(domain, " ")

		exist, err := exists(domain)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(marks[!exist])

		time.Sleep(1 * time.Second)
	}
}
