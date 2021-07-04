package global

import (
	"github.com/golang-minibear2333/gin-blog/pkg/logger"
	"github.com/golang-minibear2333/gin-blog/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	JWTSetting      *setting.JWTSettingS
	EmailSetting    *setting.EmailSettingS
)

var (
	Logger *logger.Logger
)
