package models

import(
	"github.com/jinzhu/gorm"
	"github.com/obigvee/Go-BookStore/pkg/config"
)

var db *gorm.DB

type Book struct{
	gorm.Model
	Name string `gorm:"json:name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

func init(){
	// initiailize Database connection
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func(book *Book) CreateBook()*Book{
	//function return a type pointer to struct Book
	db.NewRecord(book)
	db.Create(&book)
	return book
}


func GetAllBooks()[]Book{
	//function gets all books
	// and return an array of type Book
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBook(Id int64)(*Book, *gorm.DB){
	// function  gets book by id and 
	// returns type Book, and the DB
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(Id int64) Book{
	// function deletes a book by id
	// and returns a type Book
	var book Book
	db.Where("ID=?", Id).Delete(book)
	return book
}