package blls

import "TaskManagementSystem_Api/models/common"

type UserBLL struct {
}

func (bll UserBLL) GetToken(uid, password string) (token string, err error) {
	return (&common.AuthorizeStruct{}).ApplyAuthorize(uid, password)
}
