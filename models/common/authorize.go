package common

import (
	"crypto/sha1"
	"fmt"
	"io"
	"math/rand"
	"time"
)

type TokenStruct struct {
}

func (a *TokenStruct) createToken(uid string) (token string, err error) {
	crutime := time.Now().UnixNano()
	rand.Seed(crutime)
	sha := sha1.New()
	io.WriteString(sha, fmt.Sprintf("%d-%d", crutime, rand.Int63()))
	token = fmt.Sprintf("%x", sha.Sum(nil))
	return
}

func (a *TokenStruct) ApplyAuthorize(uid, password string) (token string, err error) {
	_, err = (&OAuthServiceStruct{}).GetOAuth2Token(uid, password)
	if err != nil {
		return
	}
	token, err = a.createToken(uid)
	return
}
