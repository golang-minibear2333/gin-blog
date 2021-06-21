package upload

import (
	"github.com/golang-minibear2333/gin-blog/global"
	"io/ioutil"
	"mime/multipart"
	"os"
	"strings"
)

// CheckSavePath 检查文件路径是否存在
func CheckSavePath(dst string) bool {
	// 获取文件的描述信息
	_, err := os.Stat(dst)
	// error 值与系统中所定义的 os.error.ErrNotExist 进行判断，以此达到校验效果。
	return os.IsNotExist(err)
}

// CheckPermission 检查文件权限是否足够
func CheckPermission(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsPermission(err)
}

// CheckContainExt 检查上传的文件后缀是否包含在约定的后缀配置项中
func CheckContainExt(t FileType, name string) bool {
	ext := GetFileExt(name)
	ext = strings.ToUpper(ext)
	switch t {
	case TypeImage:
		for _, allowExt := range global.AppSetting.UploadImageAllowExts {
			// 统一转为大写（固定的格式）来进行匹配 这样支持忽略后缀大小写
			if strings.ToUpper(allowExt) == ext {
				return true
			}
		}

	}

	return false
}

// CheckMaxSize 检查文件大小是否超出最大大小限制
func CheckMaxSize(t FileType, f multipart.File) bool {
	content, _ := ioutil.ReadAll(f)
	size := len(content)
	switch t {
	case TypeImage:
		if size >= global.AppSetting.UploadImageMaxSize*1024*1024 {
			return true
		}
	}

	return false
}
