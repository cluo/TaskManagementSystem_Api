package dals

import (
	"TaskManagementSystem_Api/models/common"
	"TaskManagementSystem_Api/models/types"
	"io"
	"mime/multipart"

	"gopkg.in/mgo.v2/bson"
)

// AttachmentDAL 定义
type AttachmentDAL struct {
	mongo *common.MongoSessionStruct
}

// UploadProductAttachment 定义
func (dal *AttachmentDAL) UploadProductAttachment(productID string, filename string, f multipart.File) (err error) {
	dal.mongo, err = common.GetMongoSession()
	if err != nil {
		return
	}
	defer dal.mongo.CloseSession()
	dal.mongo.UseDB("local")

	file, err1 := dal.mongo.Db.GridFS("fs").Create(filename)
	if err1 != nil {
		err = err1
		return
	}

	fileSize, err1 := io.Copy(file, f)
	if err1 != nil {
		err = err1
		return
	}
	err = file.Close()
	fileOid := file.Id().(bson.ObjectId)

	err = dal.mongo.UseCollection("T_Products")
	if err != nil {
		return
	}
	product := new(types.ObjectID)
	err = dal.mongo.Collection.Find(bson.M{"id": productID}).One(product)
	if err != nil {
		return
	}

	err = dal.mongo.UseCollection("T_ProductAttachments")
	if err != nil {
		return
	}

	productAttachment := new(types.ProductAttachment)
	productAttachment.OID = bson.NewObjectId()
	productAttachment.FileLength = &fileSize
	productAttachment.FileName = &filename
	productAttachment.FileObjectID = &fileOid
	if product.Oid == nil {
		productAttachment.ProductID = nil
		productAttachment.ProductObjectID = nil
	} else {
		productAttachment.ProductID = &productID
		productAttachment.ProductObjectID = product.Oid
	}

	err = dal.mongo.Collection.Insert(productAttachment)
	if err != nil {
		return
	}

	return
}

// DownloadAttachment 定义
func (dal *AttachmentDAL) DownloadAttachment(fileID string) (err error) {
	return
}
