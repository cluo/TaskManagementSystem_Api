package types

import "gopkg.in/mgo.v2/bson"

type Attachment struct {
	OID              bson.ObjectId  `bson:"_id"`
	RelevantObjectID *bson.ObjectId `bson:"relevantObjectId"`
	RelevantID       *string        `bson:"relevantId"`
	FileObjectID     *bson.ObjectId `bson:"fileObjectId"`
	FileName         *string        `bson:"fileName"`
	FileLength       *int64         `bson:"fileLength"`
}

type Attachment_Get struct {
	RelevantID   *string        `json:"relevantId" bson:"relevantId"`
	FileObjectID *bson.ObjectId `json:"fileObjectId" bson:"fileObjectId"`
	FileName     *string        `json:"fileName" bson:"fileName"`
	FileLength   *int64         `json:"fileLength" bson:"fileLength"`
}
