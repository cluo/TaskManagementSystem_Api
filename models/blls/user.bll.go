package blls

import (
	"TaskManagementSystem_Api/models/common"
	"TaskManagementSystem_Api/models/dals"
	"TaskManagementSystem_Api/models/types"
	"encoding/json"
)

type UserBLL struct {
}

func (bll *UserBLL) GetToken(uid, password string) (token string, err error) {
	token, err = (&common.TokenStruct{}).ApplyAuthorize(uid, password)
	if err != nil {
		return
	}
	user, err := (&dals.UserDAL{}).GetUserInfo(uid)
	if err != nil {
		token = ""
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
