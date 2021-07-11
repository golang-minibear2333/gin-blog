package version

import "fmt"

var (
	// Version
	// 主版本号 Major 当功能模块有较大的变动，比如增加多个模块或者整体架构发生变化。此版本号由项目决定是否修改
	// 次版本号 Minor 当功能有一定的增加或变化，比如增加了对权限控制、增加自定义视图等功能。此版本号由项目决定是否修改
	// 修订号 Revision 一般是 Bug 修复或是一些小的变动，要经常发布修订版，时间间隔不限，修复一个严重的bug即可发布一个修订版
	Version = "4.1.0"
	// StageSuffix 后缀
	StageSuffix = "dev"
	AppName     = "gin-blog"
)

func GetVersion() string {
	return Version
}
func GetStageSuffix() string {
	return StageSuffix
}

func DetailVersion() string {
	return fmt.Sprintf("%s.%s", Version, StageSuffix)
}
