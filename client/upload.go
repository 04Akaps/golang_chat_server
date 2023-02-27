package client

import (
	"io"
	"io/ioutil"
	"net/http"
	"path"
)

type FileUpload struct{}

// ServeHTTP
// (f *FileUpload)
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
	err = ioutil.WriteFile(filename, data, 0o777)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	io.WriteString(w, "Successful")
}
