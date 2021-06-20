package upload

import (
	"github.com/golang-minibear2333/gin-blog/global"
	"github.com/golang-minibear2333/gin-blog/pkg/util"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

type FileType int

const TypeImage FileType = iota + 1

func GetFileName(name string) string {
	ext := GetFileExt(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMD5(fileName)

	return fileName + ext
}

// GetFileExt 获取文件后缀
func GetFileExt(name string) string {
	return path.Ext(name)
}

// GetSavePath 获取保存路径
func GetSavePath() string {
	return global.AppSetting.UploadSavePath
}

// CreateSavePath 创建文件保存目录
func CreateSavePath(dst string, perm os.FileMode) error {
	// 传入的 os.FileMode 权限位去递归创建所需的所有目录结构，若涉及的目录均已存在，则不会进行任何操作，直接返回 nil。
	err := os.MkdirAll(dst, perm)
	if err != nil {
		return err
	}
	return nil
}

// SaveFile 保存文件
func SaveFile(file *multipart.FileHeader, dst string) error {
	// file.Open 方法打开源地址的文件
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	// os.Create 方法创建目标地址的文件
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	// io.Copy 方法实现两者之间的文件内容拷贝
	_, err = io.Copy(out, src)
	return err
}
