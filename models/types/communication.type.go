package types

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Communication struct {
	OID              bson.ObjectId  `bson:"_id"`
	RelevantID       *string        `bson:"relevantId"`
	RelevantObjectID *bson.ObjectId `bson:"relevantObjectId"`
	PersonID         *string        `bson:"personId"`
	PersonObjectID   *bson.ObjectId `bson:"personObjectId"`
	SentTime         *time.Time     `bson:"sentTime"`
	Content          *string        `bson:"content"`
}

type Communication_Get struct {
	RelevantID *string    `json:"relevantId"`
	PersonID   *string    `json:"personId"`
	PersonName *string    `json:"personName"`
	SentTime   *time.Time `json:"sentTime"`
	Content    *string    `json:"content"`
}

type Communication_Post struct {
	RelevantID *string    `json:"relevantId"`
	PersonID   *string    `json:"personId"`
	SentTime   *time.Time `json:"sentTime"`
	Content    *string    `json:"content"`
}
type ObjectID struct {
	Oid *bson.ObjectId `bson:"_id"`
}
