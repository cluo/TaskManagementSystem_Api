package common

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"time"
)

type TokenStruct struct {
}

func (a *TokenStruct) CreateToken(uid string) (token string, err error) {
	crutime := time.Now().UnixNano()
	rand.Seed(crutime)
	sha := sha1.New()
	io.WriteString(sha, fmt.Sprintf("%d-%d", crutime, rand.Int63()))
	token = fmt.Sprintf("%x", sha.Sum(nil))
	return
}

func (a *TokenStruct) ApplyAuthorize(uid, password string) (err error) {
	_, err = (&OAuthServiceStruct{}).GetOAuth2Token(uid, password)
	if err != nil {
		err = errors.New("该用户不存在，请联系管理员。")
		return
	}
	return
}
