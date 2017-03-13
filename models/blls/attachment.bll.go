package blls

import "TaskManagementSystem_Api/models/dals"

import "mime/multipart"

// AttachmentBLL 定义
type AttachmentBLL struct {
}

// UploadProductAttachment 定义
func (bll *AttachmentBLL) UploadProductAttachment(productID string, filename string, f multipart.File) (err error) {
	return (&dals.AttachmentDAL{}).UploadProductAttachment(productID, filename, f)
}

// DownloadAttachment 定义
func (bll *AttachmentBLL) DownloadAttachment(fileID string) (err error) {
	return (&dals.AttachmentDAL{}).DownloadAttachment(fileID)
}
