package types

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Communication struct {
	RelevantID       *string        `bson:"relevantId" json:"relevantId"`
	RelevantObjectID *bson.ObjectId `bson:"relevantObjectId" json:"relevantObjectId"`
	PersonID         *string        `bson:"personId" json:"personId"`
	PersonObjectID   *bson.ObjectId `bson:"personObjectId" json:"personObjectId"`
	PersonName       *string        `json:"personName"`
	SentTime         *time.Time     `bson:"sentTime" json:"sentTime"`
	Content          *string        `bson:"content" json:"content"`
}

type Communication_Insert struct {
	RelevantID       *string        `bson:"relevantId" json:"relevantId"`
	RelevantObjectID *bson.ObjectId `bson:"relevantObjectId" json:"relevantObjectId"`
	PersonID         *string        `bson:"personId" json:"personId"`
	PersonObjectID   *bson.ObjectId `bson:"personObjectId" json:"personObjectId"`
	SentTime         *time.Time     `bson:"sentTime" json:"sentTime"`
	Content          *string        `bson:"content" json:"content"`
}
type ObjectID struct {
	Oid *bson.ObjectId `bson:"_id"`
}
