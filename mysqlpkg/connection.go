package mysqlpkg

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Open() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:root@/organization?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("ERROR:", err)
		return nil, err
	}
	return db, nil
}
