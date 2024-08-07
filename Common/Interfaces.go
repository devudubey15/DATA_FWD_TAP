package common

import "gorm.io/gorm"

type DBInstanceManager interface {
	GetDBInstance() *gorm.DB
}
