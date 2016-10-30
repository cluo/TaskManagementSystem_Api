package blls

import (
	"TaskManagementSystem_Api/models/dals"
	"TaskManagementSystem_Api/models/types"
)

type CommunicationBLL struct {
}

func (bll CommunicationBLL) GetCommunications(id string) (c map[string][]*types.Communication, err error) {
	c, err = (&dals.CommunicationDAL{}).GetCommunications(id)
	return
}

func (bll CommunicationBLL) AddCommunication(c types.Communication_Insert) (s map[string]map[string]string, err error) {
	s, err = (&dals.CommunicationDAL{}).AddCommunication(c)
	return
}
