package model

import (
	"time"

	"github.com/wushengyouya/blog-service/pkg/app"
	"gorm.io/gorm"
)

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (t Tag) Count(db *gorm.DB) (int, error) {
	var count int64
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("state = ?", t.State)
	if err := db.Model(&t).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}

func (t Tag) List(db *gorm.DB, pageOffset, pageSize int) ([]*Tag, error) {
	var tags []*Tag
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	if err = db.Where("is_del = ?", 0).Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

func (t Tag) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

func (t Tag) Update(db *gorm.DB) error {
	return db.Model(&t).Where("id = ? AND is_del = ?", t.ID, 0).Updates(t).Error
}

// TODO:手动软删除，后续更新到回调软删除
func (t Tag) Delele(db *gorm.DB) error {
	db = db.Where("id = ? AND is_del = ?", t.Model.ID, 0).Find(&t)
	t.IsDel = 1
	t.DeletedOn = uint32(time.Now().Unix())
	return db.Updates(&t).Error
}
func (t Tag) TableName() string {
	return "blog_tag"
}

// swagger
type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}
