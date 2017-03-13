package controllers

import (
	"TaskManagementSystem_Api/models/blls"

	"fmt"

	"github.com/astaxie/beego"
)

// Operations about Attachments
type AttachmentController struct {
	beego.Controller
}

func (u *AttachmentController) UploadProductAttachment() {
	body := &ResponeBodyStruct{}
	// token := u.Ctx.Input.Header("X-Auth-Token")
	// _, err := (&blls.UserBLL{}).ValidateToken(token)
	// if err != nil {
	// 	body.Error = err.Error()
	// 	u.Data["json"] = body
	// 	u.Ctx.Output.SetStatus(401)
	// 	u.ServeJSON()
	// 	return
	// }
	tid := u.GetString(":tid")
	f, h, err := u.GetFile("file")
	if tid != "" && err == nil {
		defer f.Close()
		filename := h.Filename
		err = (&blls.AttachmentBLL{}).UploadProductAttachment(tid, filename, f)
	}

	// reader, err := u.Ctx.Request.MultipartReader()
	// if err == nil {
	// 	for {
	// 		part, err := reader.NextPart()
	// 		if err != nil {
	// 			break
	// 		}

	// 		filename := part.FileName()
	// 		if filename != "" {
	// 			filename = path.Base(filename)
	// 		}
	// 		if filename == "" {
	// 			break
	// 		}
	// 		err = (&blls.AttachmentBLL{}).UploadAttachment(filename, part)
	// 		if err != nil {
	// 			break
	// 		}
	// 	}
	// }
	if err != nil {
		body.Error = err.Error()
	} else {
		body.Data = "文件上传成功"
	}
	u.Data["json"] = body
	u.ServeJSON()
	fmt.Println(body)
}

func (u *AttachmentController) DownloadAttachment() {
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

}
