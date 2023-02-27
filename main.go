package main

import (
	"log"
	"net/http"
	"strings"

	"go_chat/auth"
	"go_chat/client"

	"github.com/rs/cors"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/facebook"
	"github.com/stretchr/gomniauth/providers/github"
	"github.com/stretchr/gomniauth/providers/google"
)

const (
	googleId         = "134341176703-okntrpan6f9ivdo1dn3avphkd70pl4nb.apps.googleusercontent.com"
	googlePassword   = "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"
	facebookId       = "3398218997172821"
	facebookPassword = "c4cdc56c0ee1f6d45e694f5b05165e8e"
	githubId         = "3dc3a526e845020106dc"
	githubPassword   = "0cae3f982a7f96cbc8bddd3d9d0e8a998af961e1"
	authKey          = "123423424"
	baseUri          = "http://localhost:8080/auth/callback"
)

func init() {
	gomniauth.SetSecurityKey(authKey)
	gomniauth.WithProviders(
		google.New(googleId, googlePassword, strings.Join([]string{baseUri, "/google"}, "")),
		// https://console.cloud.google.com/apis/credentials/oauthclient/134341176703-okntrpan6f9ivdo1dn3avphkd70pl4nb.apps.googleusercontent.com?hl=ko&project=golang-chat-378904
		facebook.New(facebookId, facebookPassword, strings.Join([]string{baseUri, "/facebook"}, "")),
		// https://developers.facebook.com/apps/3398218997172821/fb-login/settings/
		github.New(githubId, githubPassword, strings.Join([]string{baseUri, "/github"}, "")),
	)
}

func main() {
	r := client.NewRoom(client.UseAuthAvatar)
	mux := http.NewServeMux()

	corHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000/"},
		AllowCredentials: true,
		MaxAge:           0,
		Debug:            false,
	})

	handler := cors.Default().Handler(mux)
	handler = corHandler.Handler(handler)

	go r.Run() // channel을 받아주는 select문 시작
	// &client.FileUpload{}
	mux.Handle("/room", auth.MustAuth(r))
	mux.HandleFunc("/upload", client.UploadFile)
	mux.HandleFunc("/auth/", auth.Loginhandler)
	mux.HandleFunc("/logout", auth.LogoutHandler)

	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal((err))
	}
}
