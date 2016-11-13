package blls

import (
	"TaskManagementSystem_Api/models/common"
	"TaskManagementSystem_Api/models/dals"
	"TaskManagementSystem_Api/models/types"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/astaxie/beego"
)

type UserBLL struct {
}

var authorizationMethod = strings.ToLower(beego.AppConfig.String("authorization_method"))

func (bll *UserBLL) SignIn(uid, password string) (u types.UserInfo_Get, err error) {
	var passwordSha1 *string
	if authorizationMethod == "oauth" {
		err = (&common.TokenStruct{}).ApplyAuthorize(uid, password)
		if err != nil {
			return
		}
	} else {
		sha := sha1.New()
		io.WriteString(sha, fmt.Sprintf("%s", password))
		shaString := fmt.Sprintf("%x", sha.Sum(nil))
		passwordSha1 = &shaString
	}
	user, err1 := (&dals.UserDAL{}).GetUserInfo(uid, passwordSha1)
	if err1 != nil {
		err = err1
		return
	}
	token, err1 := (&common.TokenStruct{}).CreateToken(uid)
	if err1 != nil {
		err = err1
		return
	}
	user.Token = &token
	userInfo, _ := json.Marshal(user)
	common.SetRedis(token, string(userInfo))
	u = *user
	return
}
func (bll *UserBLL) GetUserInfo(token string) (u types.UserInfo_Get, err error) {
	userinfo, err := common.GetRedis(token)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(userinfo), &u)
	return
}

func (bll *UserBLL) ValidateToken(token string) (u types.UserInfo_Get, err error) {
	u, err = bll.GetUserInfo(token)
	return
}
