package client

import (
	"errors"
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
	if userId, ok := c.UserData["user_id"]; ok {
		if userIdStr, ok := userId.(string); ok {
			return "//www.gravatar.com/avatar/" + userIdStr, nil
		}
	}
	return "", ErrNoAvatarURL
}
