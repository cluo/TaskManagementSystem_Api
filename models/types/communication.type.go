package types

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Communication struct {
	RelevantID       string        `bson:"relevantId" json:"relevantId"`
	RelevantObjectID bson.ObjectId `bson:"relevantObjectId" json:"relevantObjectId"`
	PersonID         string        `bson:"personID" json:"personID"`
	PersonObjectID   bson.ObjectId `bson:"personObjectID" json:"personObjectID"`
	Person           string        `bson:"Person" json:"Person"`
	SentTime         time.Time     `bson:"sentTime" json:"sentTime"`
	Content          string        `bson:"content" json:"content"`
}
