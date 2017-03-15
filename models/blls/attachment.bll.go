package blls

import (
	"TaskManagementSystem_Api/models/dals"
	"mime/multipart"

	"github.com/astaxie/beego/context"
)

// AttachmentBLL 定义
type AttachmentBLL struct {
}

// UploadProductAttachment 定义
func (bll *AttachmentBLL) UploadProductAttachment(productID string, filename string, f multipart.File) (err error) {
	err = (&dals.AttachmentDAL{}).UploadProductAttachment(productID, filename, f)
	return
}

// DownloadAttachment 定义
func (bll *AttachmentBLL) DownloadAttachment(fileID string, writer *context.Response) (errStatusCode int, err error) {
	errStatusCode, err = (&dals.AttachmentDAL{}).DownloadAttachment(fileID, writer)
	return
}
