package config

import (
	"github.com/jinzhu/gorm"
	//"github.com/jinzhu/gorm/dialects/mysql"
	//"github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db * gorm.DB
)

// Connect:
// helps to open connection to a DB(mysql)
func Connect(){
	d, err := gorm.Open("mysql","root:thisme@/bookTable?charset=utf8&parseTime=True&loc=Local")
	
	// if err is not zero it means that there is an error with the connection
	if err != nil{
		panic(err)
	}
	// if not 
	db = d
}

// GetDB will be called in other files
func GetDB() *gorm.DB{
	return db
}