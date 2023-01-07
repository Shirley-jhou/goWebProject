package database

import (
	_ "fmt"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	_ "gorm.io/plugin/dbresolver"
)

var sqlDB *gorm.DB
func Setup(){

}

func GetDB(names ...string) *gorm.DB{
	if len(names) > 0 {
		return sqlDB.Clauses(dbresolver.Use(names[0]))
	}
	return sqlDB
}