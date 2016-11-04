package models

import "TaskManagementSystem_Api/models/blls"

func GetToken(code string) (token string, err error) {
	token, err = (&blls.TokenBLL{}).GetToken(code)
	return
}
