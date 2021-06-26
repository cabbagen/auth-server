package model

import (
	"go-gateway/schema"
	"go-gateway/database"
	"github.com/jinzhu/gorm"
)

type ApplicationModel struct {
	tableName           string
	databaseHandler     *gorm.DB
}

func NewApplicationModel() ApplicationModel {
	return ApplicationModel {
		tableName: "auth_applications",
		databaseHandler: database.GetDatabaseHandleInstance(),
	}
}

func (am ApplicationModel) GetApplicationDetail(appId int) (schema.ApplicationSchema, error) {
	var detail schema.ApplicationSchema

	if error := am.databaseHandler.Table(am.tableName).Where("id = ?", appId).First(&detail).Error; error != nil {
		return detail, error
	}
	return detail, nil
}

func (am ApplicationModel) GetApplications(name string, pageNo, pageSize int) ([]schema.ApplicationSchema, int, error) {
	var total int
	var applications []schema.ApplicationSchema

	error := am.databaseHandler.Table(am.tableName).Where("app_name like ?", "%" + name + "%").Count(&total).Offset(pageNo * pageSize).Limit(pageSize).Find(&applications).Error

	if error != nil {
		return applications, total, error
	}
	return applications, total, nil
}

func (am ApplicationModel) UpdateApplication(app schema.ApplicationSchema) error {
	if app.ID == 0 {
		return am.databaseHandler.Table(am.tableName).Create(&app).Error
	}
	return am.databaseHandler.Table(am.tableName).Save(&app).Error
}

func (am ApplicationModel) GetAllApplicationKey() ([]string, error) {
	var appKeys []string

	if error := am.databaseHandler.Table(am.tableName).Pluck("app_key", &appKeys).Error; error != nil {
		return appKeys, error
	}
	return appKeys, nil
}
