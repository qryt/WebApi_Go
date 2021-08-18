package Services

import (
	"github.com/jinzhu/gorm"
	_ "gorm.io/driver/mysql"
)

func Init(TableName string) *gorm.DB {
	db,er:=gorm.Open("mysql", "root:ABcd1234@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if er==nil&& len(TableName)>0{
		db=db.Table(TableName)
		return db
	}
	return nil
}
