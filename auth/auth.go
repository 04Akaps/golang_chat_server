package auth

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/stretchr/gomniauth"
	"github.com/stretchr/objx"
)

type AuthHandler struct {
	next http.Handler
}

func (h *AuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("auth")

	if err == http.ErrNoCookie {
		// 쿠키가 없는 것
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusTemporaryRedirect)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// 성공 -> 다음 핸들러를 호출
	h.next.ServeHTTP(w, r)
}

// 형식은 /auth/{action}/{provider}
func Loginhandler(w http.ResponseWriter, r *http.Request) {
	segs := strings.Split(r.URL.Path, "/")

	action := segs[2]
	provider := segs[3]

	switch action {

	case "login":
		provider, err := gomniauth.Provider(provider)
		if err != nil {
			http.Error(w, fmt.Sprintf("Errror when trying to get Provider %s", provider), http.StatusBadRequest)
			return
		}

		loginUrl, err := provider.GetBeginAuthURL(nil, nil)
		if err != nil {
			http.Error(w, fmt.Sprintf("Errror when trying to getAuthUrl %s", provider), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Location", loginUrl)
		w.WriteHeader(http.StatusTemporaryRedirect)

	case "callback":
		fmt.Println("Login CallBack을 처리")
		provider, err := gomniauth.Provider(provider)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error when trying to get Provider %s", err), http.StatusBadRequest)
			return
		}

		credentials, err := provider.CompleteAuth(objx.MustFromURLQuery(r.URL.RawQuery))
		// r.Url.RawQuery => Query값을 다 가져옴
		if err != nil {
			http.Error(w, fmt.Sprintf("Error when trying to complate auth %s", err), http.StatusInternalServerError)
			return
		}
		// 여기 까지 진행이 되었으면 인증된 사용자 라는 의미

		user, err := provider.GetUser(credentials)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error when trying to get User %s", err), http.StatusInternalServerError)
			return
		}

		m := md5.New()
		io.WriteString(m, strings.ToLower(user.Email()))

		authCookieValue := objx.New(map[string]interface{}{
			"user_id":    m.Sum(nil), // 고유한 userID는 이메일 기반으로 해시값을 사용
			"name":       user.Name(),
			"avatar_url": user.AvatarURL(),
			"email":      user.Email(),
		}).MustBase64()

		http.SetCookie(w, &http.Cookie{
			Name:  "auth",
			Value: authCookieValue,
			Path:  "/",
		})
		// 단순히 이름만 base64로 인코딩 해서 저장
		// 서명되지 않은 쿠키에 중요한 정보를 담을 수는 없기 떄문에

		w.Header().Set("Location", "http://localhost:3000")
		w.WriteHeader(http.StatusTemporaryRedirect)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Auth Action %s not supported", action)
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "auth",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})

	w.Header().Set("Location", "http://localhost:3000/Login")
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func MustAuth(handler http.Handler) http.Handler {
	return &AuthHandler{next: handler}
}
