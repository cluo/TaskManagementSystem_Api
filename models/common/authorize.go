package common

import (
	"crypto/sha1"
	"fmt"
	"io"
	"math/rand"
	"time"
)

type AuthorizeStruct struct {
}

func (a *AuthorizeStruct) createToken(uid string) (token string, err error) {
	crutime := time.Now().UnixNano()
	rand.Seed(crutime)
	sha := sha1.New()
	io.WriteString(sha, fmt.Sprintf("%d-%d", crutime, rand.Int63()))
	token = fmt.Sprintf("%x", sha.Sum(nil))

	Bunt.Set(token, uid)
	return
}
func (a *AuthorizeStruct) validateToken(token string) (uid string, err error) {
	return Bunt.Get(token)
}

func (a *AuthorizeStruct) ApplyAuthorize(uid, password string) (token string, err error) {
	_, err = (&OAuthServiceStruct{}).GetOAuth2Token(uid, password)
	if err != nil {
		return
	}
	token, err = a.createToken(uid)
	return
}

func (a *AuthorizeStruct) ValidateAuthorize(token string) (uid string, err error) {
	uid, err = a.validateToken(token)
	return
}
