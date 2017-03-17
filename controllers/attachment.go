package controllers

import (
	"TaskManagementSystem_Api/models/blls"

	"github.com/astaxie/beego"
)

// Operations about Attachments
type AttachmentController struct {
	beego.Controller
}

func (u *AttachmentController) Get() {
	body := &ResponeBodyStruct{}
	token := u.Ctx.Input.Header("X-Auth-Token")
	_, err := (&blls.UserBLL{}).ValidateToken(token)
	if err != nil {
		body.Error = err.Error()
		u.Data["json"] = body
		u.Ctx.Output.SetStatus(401)
		u.ServeJSON()
		return
	}

	id := u.GetString(":tid")
	if id != "" {
		attachments, err := (&blls.AttachmentBLL{}).GetAttachmentList(id)
		if err != nil {
			body.Error = err.Error()
		} else {
			body.Data = attachments
		}
	}
	u.Data["json"] = body
	u.ServeJSON()
}

func (u *AttachmentController) UploadAttachment() {
	body := &ResponeBodyStruct{}
	token := u.Ctx.Input.Header("X-Auth-Token")
	_, err := (&blls.UserBLL{}).ValidateToken(token)
	if err != nil {
		body.Error = err.Error()
		u.Data["json"] = body
		u.Ctx.Output.SetStatus(401)
		u.ServeJSON()
		return
	}
	tid := u.GetString(":tid")
	f, h, err := u.GetFile("file")
	if tid != "" && err == nil {
		defer f.Close()
		filename := h.Filename
		err = (&blls.AttachmentBLL{}).UploadAttachment(tid, filename, f)
	}
	if err != nil {
		body.Error = err.Error()
	} else {
		body.Data = "文件上传成功"
	}
	u.Data["json"] = body
	u.ServeJSON()
}

func (u *AttachmentController) DownloadAttachment() {
	// body := &ResponeBodyStruct{}
	// token := u.Ctx.Input.Header("X-Auth-Token")
	// _, err := (&blls.UserBLL{}).ValidateToken(token)
	// if err != nil {
	// 	body.Error = err.Error()
	// 	u.Data["json"] = body
	// 	u.Ctx.Output.SetStatus(401)
	// 	return
	// }
	fid := u.GetString(":fid")
	if fid != "" {
		errStatusCode, err := (&blls.AttachmentBLL{}).DownloadAttachment(fid, u.Ctx.ResponseWriter)
		if err != nil {
			u.Ctx.Output.SetStatus(errStatusCode)
		}
	} else {
		u.Ctx.Output.SetStatus(404)
	}
	return
}

func (u *AttachmentController) DeleteAttachment() {
	body := &ResponeBodyStruct{}
	token := u.Ctx.Input.Header("X-Auth-Token")
	_, err := (&blls.UserBLL{}).ValidateToken(token)
	if err != nil {
		body.Error = err.Error()
		u.Data["json"] = body
		u.Ctx.Output.SetStatus(401)
		return
	}
	fid := u.GetString(":fid")
	if fid != "" {
		err = (&blls.AttachmentBLL{}).DelAttachment(fid)
	}
	if err != nil {
		body.Error = err.Error()
	} else {
		body.Data = "文件删除成功"
	}
	u.Data["json"] = body
	u.ServeJSON()
	return
}
