package models

import (
	"TaskManagementSystem_Api/models/blls"
	"TaskManagementSystem_Api/models/types"
)

func GetCommunications(id string) (c map[string]*types.Communication, err error) {
	c, err = (&blls.CommunicationBLL{}).GetCommunications(id)
	return
}
