package model

import (
	"github.com/golang-minibear2333/gin-blog/pkg/app"
	"github.com/jinzhu/gorm"
)

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}

// TableName 解析表名，如果不写默认解析结构体名
func (t Tag) TableName() string {
	return "blog_tag"
}

func (t Tag) Count(db *gorm.DB) (int, error) {
	var count int
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	// 传入查询参数 sql 的 where 有防sql注入功能
	db = db.Where("state = ?", t.State)
	// Count 统计行为，用于统计模型的记录数。
	if err := db.Model(&t).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (t Tag) List(db *gorm.DB, pageOffset, pageSize int) ([]*Tag, error) {
	var tags []*Tag
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		// Offset 偏移量，用于指定开始返回记录之前要跳过的记录数 Limit 限制检索的记录数
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("state = ?", t.State)
	// Find 查找符合筛选条件的记录，用来赋值的
	if err = db.Where("is_del = ?", 0).Find(&tags).Error; err != nil {
		return nil, err
	}

	return tags, nil
}

func (t Tag) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

func (t Tag) Update(db *gorm.DB, values interface{}) error {
	// Updates 更新所选字段
	// 不能这样写 db.Model(&Tag{}).Where("id = ? AND is_del = ?", t.ID, 0).Update(t).Error
	// https://github.com/golang-minibear2333/gin-blog/issues/2
	if err := db.Model(t).Where("id = ? AND is_del = ?", t.ID, 0).Updates(values).Error; err != nil {
		return err
	}
	return nil
}

func (t Tag) Delete(db *gorm.DB) error {
	// Delete 删除数据
	// TODO 这里是硬删除，后期考虑换成软删除和增加已删除tag找回功能
	return db.Where("id = ? AND is_del = ?", t.Model.ID, 0).Delete(&t).Error
}
