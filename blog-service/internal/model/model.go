package model

import (
	"fmt"

	"github.com/wushengyouya/blog-service/global"
	"github.com/wushengyouya/blog-service/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  uint32 `json:"created_on"`
	ModifedOn  uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}

func NewDBEngine(databaseSetting *setting.DataBaseSettingS) (*gorm.DB, error) {
	s := "%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local"
	dsn := fmt.Sprintf(s,
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}))

	if err != nil {
		return nil, err
	}

	if global.ServerSetting.RunMode == "debug" {
		db.Logger.LogMode(logger.Info)
	}
	sqldb, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqldb.SetMaxIdleConns(databaseSetting.MaxIdleConns)
	sqldb.SetMaxOpenConns(databaseSetting.MaxOpenConns)
	return db, nil
}
