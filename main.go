package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"

	"go_chat/client"

	"github.com/rs/cors"
)

type templateHandler struct {
	once     sync.Once // 얼마나 많은 고루틴이 접근을 하든 상관없이 단 한번만 실행 된다는 것을 보장한다.
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 소스 파일을 로드하고, 컴파일 한 후 실행한뒤
	// 출력에 전달 해 준다.
	// 해당 코드가 사용가능한 이유는 해당 메소드가 http.Handler인터페이스를 만족하기 떄문이다.
	t.once.Do(func() {
		t.templ = template.Must(
			template.ParseFiles(filepath.Join("page", t.filename)))
	})
	t.templ.Execute(w, nil)
}

func main() {
	r := client.NewRoom()
	mux := http.NewServeMux()

	corHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000/"},
		AllowCredentials: true,
		MaxAge:           0,
		Debug:            true,
	})
	handler := cors.Default().Handler(mux)
	handler = corHandler.Handler(handler)

	go r.Run() // channel을 받아주는 select문 시작

	mux.Handle("/room", r)
	mux.Handle("/", &templateHandler{filename: "index.html"})

	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal((err))
	}
}
