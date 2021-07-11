package api

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-minibear2333/gin-blog/global"
	"github.com/golang-minibear2333/gin-blog/internal/service"
	"github.com/golang-minibear2333/gin-blog/pkg/app"
	"github.com/golang-minibear2333/gin-blog/pkg/convert"
	"github.com/golang-minibear2333/gin-blog/pkg/errcode"
	"github.com/golang-minibear2333/gin-blog/pkg/upload"
)

type Upload struct{}

func NewUpload() Upload {
	return Upload{}
}

func (u Upload) UploadFile(c *gin.Context) {
	response := app.NewResponse(c)
	// 读取入参 file 字段的上传文件信息
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}
	// 利用入参 type 字段作为所上传文件类型的确立依据
	// TODO 将通过上传type字段确认文件类型 改为通过解析上传文件后缀来确定文件类型，注意同时更新swagger（注解与调用scripts/build_swagger.sh）
	fileType := convert.StrTo(c.PostForm("type")).MustInt()
	if fileHeader == nil || fileType <= 0 {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}

	svc := service.New(c.Request.Context())
	// 最后通过入参检查后进行 service 的调用，完成上传和文件保存，返回文件的展示地址。
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
	if err != nil {
		global.Logger.Errorf(c, "svc.UploadFile err: %v", err)
		response.ToErrorResponse(errcode.ErrorUploadFileFail.WithDetails(err.Error()))
		return
	}

	response.ToResponse(gin.H{
		"file_access_url": fileInfo.AccessUrl,
	})
}
