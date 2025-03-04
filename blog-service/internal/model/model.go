package model

import (
	"fmt"
	"log"
	"time"

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
	ModifiedOn uint32 `json:"modified_on"`
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
	// 数据库操作后的回调函数
	db.Callback().Create().Before("gorm:create").Register("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Before("gorm:update").Register("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	db.Callback().Delete().Before("gorm:delete").Register("gorm:update_delete_status", deleteCallback)

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

// 创建全局回调
func updateTimeStampForCreateCallback(db *gorm.DB) {
	if err := db.Error; err != nil {
		return
	}
	nowTime := time.Now().Unix()
	// 第一种写法
	// if createTimeField, ok := db.Statement.Schema.FieldsByDBName["CreateOn"]; ok {
	// 	if createTimeField != nil {
	// 		createTimeField.Set(db.Statement.Context, db.Statement.ReflectValue, nowTime)
	// 	}
	// }
	// 获取结构体并并在create前设置 createOn的值
	// 第二种写法
	// db.Statement.Schema 是一个模型字段的元数据集合，存储了当前操作涉及的数据库模型(即Go结构体)的所有字段的详细信息
	createTimeField := db.Statement.Schema.LookUpField("CreatedOn")
	if createTimeField != nil {
		err := createTimeField.Set(db.Statement.Context, db.Statement.ReflectValue, nowTime)
		if err != nil {
			log.Println(err)
			global.Logger.Errorf("updateTimeStampForCreateCallback.createTimeField.Set err: %v", err)
		}
	}

	modifyTimeField := db.Statement.Schema.LookUpField("ModifiedOn")
	if modifyTimeField != nil {
		err := modifyTimeField.Set(db.Statement.Context, db.Statement.ReflectValue, nowTime)
		if err != nil {
			global.Logger.Errorf("updateTimeStampForCreateCallback.modifyTimeField.Set err: %v", err)
		}
	}

}

// 更新全局回调
func updateTimeStampForUpdateCallback(db *gorm.DB) {
	if _, ok := db.Statement.Get("gorm:update_column"); !ok {
		reflectValue := db.Statement.ReflectValue
		if field := db.Statement.Schema.LookUpField("ModifedOn"); field != nil {
			err := field.Set(db.Statement.Context, reflectValue, time.Now().Unix())
			if err != nil {
				global.Logger.Errorf("updateTimeStampForUpdateCallback errs: %v", err)
			}

		}

	}
}

// 删除全局回调
func deleteCallback(db *gorm.DB) {
	scheme := db.Statement.Schema
	// idOnField := scheme.LookUpField("ID")
	deleteOnField := scheme.LookUpField("DeletedOn")
	now := time.Now().Unix()
	if deleteOnField != nil {
		err := deleteOnField.Set(db.Statement.Context, db.Statement.ReflectValue, now)
		if err != nil {
			global.Logger.Errorf("deleteCallback.deleteOnField.Set errs: %v", err)
		}
	}
	isDelField := scheme.LookUpField("IsDel")
	if isDelField != nil {
		err := isDelField.Set(db.Statement.Context, db.Statement.ReflectValue, 1)
		if err != nil {
			global.Logger.Errorf("deleteCallback.isDelField.Set errs: %v", err)
		}
	}
	// 	condition := fmt.Sprintf("WHERE id = %v %v", idOnField.DBName, extraOption)
	// 	sql := fmt.Sprintf("UPDATE %v SET %v=%v,%v=%v %v", scheme.Table, deleteOnField.DBName, now, isDelField.DBName, 1, condition)
	// 	err := db.Statement.Raw(sql).Error
	// 	if err != nil {
	// 		global.Logger.Errorf("deleteCallback.db.Statement.Raw errs: %v", err)
	// 	}
	// }
}
