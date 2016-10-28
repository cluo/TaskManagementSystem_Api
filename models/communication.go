package models

import (
	"TaskManagementSystem_Api/models/blls"
	"TaskManagementSystem_Api/models/types"
)

func GetCommunications(id string) (c map[string][]*types.Communication, err error) {
	c, err = (&blls.CommunicationBLL{}).GetCommunications(id)
	return
}

func AddCommunication(c types.Communication_Post) (s map[string]map[string]string, err error) {
	s, err = (&blls.CommunicationBLL{}).AddCommunication(c)
	return
}
