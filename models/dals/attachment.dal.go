package dals

import (
	"TaskManagementSystem_Api/models/common"
	"TaskManagementSystem_Api/models/types"
	"fmt"
	"io"
	"mime/multipart"
	"strconv"
	"strings"

	"github.com/astaxie/beego/context"

	"errors"

	"gopkg.in/mgo.v2/bson"
)

// AttachmentDAL 定义
type AttachmentDAL struct {
	mongo *common.MongoSessionStruct
}

// GetAttachmentList 定义
func (dal *AttachmentDAL) GetAttachmentList(id string) (attachmentsGet []*types.Attachment_Get, err error) {
	dal.mongo, err = common.GetMongoSession()
	if err != nil {
		return
	}
	defer dal.mongo.CloseSession()
	dal.mongo.UseDB("local")
	err = dal.mongo.UseCollection("T_Attachments")
	if err != nil {
		return
	}

	err = dal.mongo.Collection.Find(bson.M{"relevantId": id}).All(&attachmentsGet)
	return
}

// UploadAttachment 定义
func (dal *AttachmentDAL) UploadAttachment(id string, filename string, f multipart.File) (err error) {
	dal.mongo, err = common.GetMongoSession()
	if err != nil {
		return
	}
	defer dal.mongo.CloseSession()
	dal.mongo.UseDB("local")

	objectID := new(types.ObjectID)
	err = dal.mongo.Db.C("T_Tasks").Find(bson.M{"id": id}).One(&objectID)
	err = dal.mongo.Db.C("T_Projects").Find(bson.M{"id": id}).One(&objectID)
	err = dal.mongo.Db.C("T_Products").Find(bson.M{"id": id}).One(&objectID)
	if objectID.Oid == nil || !bson.ObjectId.Valid(*objectID.Oid) {
		err = errors.New("不存在该编号的数据。")
		return
	}

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

	attachment := new(types.Attachment)
	attachment.OID = bson.NewObjectId()
	attachment.RelevantID = &id
	attachment.RelevantObjectID = objectID.Oid
	attachment.FileLength = &fileSize
	attachment.FileName = &filename
	attachment.FileObjectID = &fileOid
	err = dal.mongo.UseCollection("T_Attachments")
	if err != nil {
		return
	}
	err = dal.mongo.Collection.Insert(attachment)
	if err != nil {
		return
	}

	return
}

// DownloadAttachment 定义
func (dal *AttachmentDAL) DownloadAttachment(fileID string, writer *context.Response) (errStatusCode int, err error) {
	dal.mongo, err = common.GetMongoSession()
	if err != nil {
		errStatusCode = 501
		return
	}
	defer dal.mongo.CloseSession()
	dal.mongo.UseDB("local")
	f, err1 := dal.mongo.Db.GridFS("fs").OpenId(bson.ObjectIdHex(fileID))
	if err1 != nil {
		errStatusCode = 404
		err = err1
		return
	}

	size := f.Size()
	writer.Header().Set("Content-Disposition",
		fmt.Sprintf(`inline; filename="%s"`, escapeQuotes(f.Name())))
	writer.Header().Set("Content-Length", strconv.FormatInt(size, 10))
	_, err = io.Copy(writer, f)
	if err != nil {
		errStatusCode = 501
	}
	f.Close()
	return
}

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func escapeQuotes(s string) string {
	return quoteEscaper.Replace(s)
}

// DelAttachment 定义
func (dal *AttachmentDAL) DelAttachment(fileID string) (err error) {
	dal.mongo, err = common.GetMongoSession()
	if err != nil {
		return
	}
	defer dal.mongo.CloseSession()
	dal.mongo.UseDB("local")
	err = dal.mongo.Db.GridFS("fs").RemoveId(bson.ObjectIdHex(fileID))
	if err != nil {
		return
	}

	err = dal.mongo.UseCollection("T_Attachments")
	if err != nil {
		return
	}
	_, err = dal.mongo.Collection.RemoveAll(bson.M{"fileObjectId": bson.ObjectIdHex(fileID)})

	return
}
