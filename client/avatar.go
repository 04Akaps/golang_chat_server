package client

import (
	"errors"
	"io/ioutil"
	"path"
)

var ErrNoAvatarURL = errors.New("no avatar url")

type Avatar interface {
	GetAvatarURL(c *Client) (string, error)
}

// -> 인터페이스가 정의한 메서드를 구션하였기 떄문에 문제 x

type AuthAvatar struct{}

var UseAuthAvatar AuthAvatar

func (AuthAvatar) GetAvatarURL(c *Client) (string, error) {
	// userId가 이미 해시화한 값이기 떄문에 바로 사용
	// 사실 이 코드는 html파일을 main.go내부에 사용할 떄 유의미 하다.
	// 단순히 패턴을 익히기 위함
	if userId, ok := c.UserData["user_id"]; ok {
		if userIdStr, ok := userId.(string); ok {

			// file에 업로드 되어 있다면 해당 값을 가져 온다.
			files, err := ioutil.ReadDir("client/avatars")
			if err != nil {
				return "", ErrNoAvatarURL
			}

			for _, file := range files {
				if file.IsDir() { // 폴더인지 확인
					continue
				}

				if match, _ := path.Match(userIdStr+"*", file.Name()); match {
					return "/avatars/" + file.Name(), nil
				}
			}
		}
	}
	return "", ErrNoAvatarURL
}
