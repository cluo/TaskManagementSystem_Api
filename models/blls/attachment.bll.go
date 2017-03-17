package blls

import (
	"TaskManagementSystem_Api/models/dals"
	"TaskManagementSystem_Api/models/types"
	"mime/multipart"

	"github.com/astaxie/beego/context"
)

// AttachmentBLL 定义
type AttachmentBLL struct {
}

// GetAttachmentList 定义
func (bll *AttachmentBLL) GetAttachmentList(id string) (attachmentsGet []*types.Attachment_Get, err error) {
	attachmentsGet, err = (&dals.AttachmentDAL{}).GetAttachmentList(id)
	return
}

// UploadAttachment 定义
func (bll *AttachmentBLL) UploadAttachment(id string, filename string, f multipart.File) (err error) {
	err = (&dals.AttachmentDAL{}).UploadAttachment(id, filename, f)
	return
}

// DownloadAttachment 定义
func (bll *AttachmentBLL) DownloadAttachment(fileID string, writer *context.Response) (errStatusCode int, err error) {
	errStatusCode, err = (&dals.AttachmentDAL{}).DownloadAttachment(fileID, writer)
	return
}

// DelAttachment 定义
func (bll *AttachmentBLL) DelAttachment(fileID string) (err error) {
	err = (&dals.AttachmentDAL{}).DelAttachment(fileID)
	return
}
