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
