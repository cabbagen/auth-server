package schema

type ApplicationSchema struct {
	BaseSchema
	AppKey          string     `gorm:"column:app_key; type:varchar(256); not null; unique key" json:"appKey"`
	AppName         string     `gorm:"column:app_name; type:varchar(256); not null; default \"\"" json:"appName"`
	AppDomain       string     `gorm:"column:app_domain; type:varchar(256); not null; unique key" json:"appDomain"`
}
