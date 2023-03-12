package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// the mysql driver lib waasnt correctly imported. the underscore is called blank indentifier in Go
	// you'll need to add it in front of that import in the case of side-effect imports. for side-effect inports, you
	// wouldnt be calling anyfuntion by yourself. you only need to import it. the underscore then would prevent Go formatter from
	// removing it (because naturally go removes all unused imports). Go database driver libraries use this pattern a lot.
	// see: https://stackoverflow.com/questions/21220077/what-does-an-underscore-in-front-of-an-import-statement-mean#:~:text=Underscore%20is%20a%20special%20character,compiler%20will%20simply%20ignore%20it.
	// https://v1.gorm.io/docs/connecting_to_the_database.html#MySQL
)

var (
	// the pointer sign wasnt correctly placed
	db *gorm.DB
)

// Connect:
// helps to open connection to a DB(mysql)
func Connect() {
	d, err := gorm.Open("mysql", "root:thisme@/bookTable?charset=utf8&parseTime=True&loc=Local")

	// if err is not zero it means that there is an error with the connection
	if err != nil {
		panic(err)
	}
	// if not
	db = d
}

// GetDB will be called in other files
func GetDB() *gorm.DB {
	return db
}
