package model

import (
	"fmt"
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
	// 数据库操作后的回调函数
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)

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
		createTimeField.Set(db.Statement.Context, db.Statement.ReflectValue, nowTime)
	}

	modifyTimeField := db.Statement.Schema.LookUpField("ModifiedOn")
	if modifyTimeField != nil {
		modifyTimeField.Set(db.Statement.Context, db.Statement.ReflectValue, nowTime)
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
	var extraOption string
	if str, ok := db.Statement.Get("gorm:delete_option"); ok {
		extraOption = fmt.Sprint(str)
	}
	scheme := db.Statement.Schema
	idOnField := scheme.LookUpField("ID")
	deleteOnField := scheme.LookUpField("DeletedOn")
	isDelField := scheme.LookUpField("IsDel")
	if deleteOnField != nil && isDelField != nil {
		now := time.Now().Unix()
		condition := fmt.Sprintf("WHERE id = %v %v", idOnField.DBName, extraOption)
		sql := fmt.Sprintf("UPDATE %v SET %v=%v,%v=%v %v", scheme.Table, deleteOnField.DBName, now, isDelField.DBName, 1, condition)
		err := db.Statement.Raw(sql).Error
		if err != nil {
			global.Logger.Errorf("deleteCallback.db.Statement.Raw errs: %v", err)
		}
	}
}
