package types

import "gopkg.in/mgo.v2/bson"

type ProductAttachment struct {
	OID             bson.ObjectId  `bson:"_id"`
	ProductObjectID *bson.ObjectId `bson:"productObjectId"`
	ProductID       *string        `bson:"productId"`
	FileObjectID    *bson.ObjectId `bson:"fileObjectId"`
	FileName        *string        `bson:"fileName"`
	FileLength      *int64         `bson:"fileLength"`
}
