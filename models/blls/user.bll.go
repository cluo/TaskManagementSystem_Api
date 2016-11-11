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

func (bll *UserBLL) GetToken(uid, password string) (token string, err error) {
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
	user, err := (&dals.UserDAL{}).GetUserInfo(uid, passwordSha1)
	if err != nil {
		token = ""
		return
	}
	token, err = (&common.TokenStruct{}).CreateToken(uid)
	if err != nil {
		return
	}
	userInfo, _ := json.Marshal(user)
	common.SetRedis(token, string(userInfo))
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

func (bll *UserBLL) ValidateToken(token string) (err error) {
	_, err = common.GetRedis(token)
	if err != nil {
		return
	}
	return
}
