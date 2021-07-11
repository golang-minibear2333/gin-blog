package model

import (
	"fmt"
	"time"

	otgorm "github.com/eddycjy/opentracing-gorm"
	"github.com/golang-minibear2333/gin-blog/global"
	"github.com/golang-minibear2333/gin-blog/pkg/setting"
	"github.com/jinzhu/gorm"

	// 必须以此引入mysql驱动库
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Model 作为公共部分字段被其他model引入
type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  uint32 `json:"created_on"`  // 创建时间
	ModifiedOn uint32 `json:"modified_on"` // 更新时间
	DeletedOn  uint32 `json:"deleted_on"`  // 删除时间
	IsDel      uint8  `json:"is_del"`
}

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	db, err := gorm.Open(databaseSetting.DBType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	))
	if err != nil {
		return nil, err
	}

	if global.ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}
	db.SingularTable(true)
	// 注册回调函数，会在执行相应语句之前回掉执行
	// gorm 把所有的执行都注册为callback，"gorm:xxx" 的字符串注册成不同的方法或者步骤阶段，可以看源码的init()函数
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)
	db.DB().SetMaxIdleConns(databaseSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(databaseSetting.MaxOpenConns)
	// SQL 链路追踪注册回调，逻辑就是先注册上下文，再加这里的回调
	otgorm.AddGormCallbacks(db)
	return db, nil
}

// updateTimeStampForCreateCallback 创建时的回调
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		// 通过调用 scope.FieldByName 方法，获取当前是否包含所需的字段
		if createTimeField, ok := scope.FieldByName("CreatedOn"); ok {
			// 通过判断 Field.IsBlank 的值，可以得知该字段的值是否为空
			if createTimeField.IsBlank {
				// 若为空，则会调用 Field.Set 方法给该字段设置值，入参类型为 interface{}，内部也就是通过反射进行一系列操作赋值。
				_ = createTimeField.Set(nowTime)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
			if modifyTimeField.IsBlank {
				_ = modifyTimeField.Set(nowTime)
			}
		}
	}
}

// updateTimeStampForUpdateCallback 更新时回调
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	// 通过调用 scope.Get("gorm:update_column") 去获取当前设置了标识 gorm:update_column 的字段属性
	if _, ok := scope.Get("gorm:update_column"); !ok {
		//  若不存在，也就是没有自定义设置 update_column，那么将会在更新回调内设置默认字段 ModifiedOn 的值为当前的时间戳
		// TODO 如果是这样的逻辑，就只能更新一次了，假如下一次更新数据怎么控制刷新这个字段呢？
		_ = scope.SetColumn("ModifiedOn", time.Now().Unix())
	}
}

// 这里参考了源码 https://gitea.com/jinzhu/gorm/src/branch/master/callback_delete.go
func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		// 通过调用 scope.Get("gorm:delete_option") 去获取当前设置了标识 gorm:delete_option 的字段属性。
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		// 判断是否存在 DeletedOn 和 IsDel 字段，若存在则调整为执行 UPDATE 操作进行软删除
		//（修改 DeletedOn 和 IsDel 的值），否则执行 DELETE 进行硬删除。
		deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedOn")
		isDelField, hasIsDelField := scope.FieldByName("IsDel")
		if !scope.Search.Unscoped && hasDeletedOnField && hasIsDelField {
			now := time.Now().Unix()
			// 软删除
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v,%v=%v%v%v",
				// 表名
				scope.QuotedTableName(),
				// TODO 这里为什么用的DBName,而不是Name，需要debug
				scope.Quote(deletedOnField.DBName),
				// 删除时间
				scope.AddToVars(now),
				scope.Quote(isDelField.DBName),
				// 设置软删除参数 1 是软删除
				scope.AddToVars(1),
				// 组装sql 判断 deleted_on 和  is_del 是否存在
				// TODO 还未理解此处，待查询资料解释
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			// 调用 scope.QuotedTableName 方法获取当前所引用的表名，并调用一系列方法针对 SQL 语句的组成部分进行处理和转移
			// 最后在完成一些所需参数设置后调用 scope.CombinedConditionSql 方法完成 SQL 语句的组装。
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}

// 为了拼接sql增加额外的空格
func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""

}
