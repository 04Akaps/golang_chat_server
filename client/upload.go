package client

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"path"
)

func UploadFile(w http.ResponseWriter, req *http.Request) {
	userId := req.FormValue("userid")
	file, header, err := req.FormFile("avatarFile")
	// file은 파일 자체를 가져온다
	// header는 파일에 대한 메타데이더를 가져온다.
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := ioutil.ReadAll(file)
	// file타입은 io.Reader타입이기 떄문에 ReadAll에 입력 가능

	filename := path.Join("client/avatars", userId+path.Ext(header.Filename))
	// ext는 파일 명을 없애 줌
	// 예를 들어서, velog.jpg => .jpg로 변경
	err = ioutil.WriteFile(filename, data, 0o777)
	// Write File은 파일을 생성
	// fileName에 data를 써줌

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("efdjdhsrjsrksrk")

	// w.Header().Set("Location", "http://localhost:3000/chat")
	// w.WriteHeader(http.StatusOK)
	io.WriteString(w, "Successful")
}
