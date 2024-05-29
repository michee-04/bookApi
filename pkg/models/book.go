package models

import (
	"github.com/jinzhu/gorm"
	"github.com/michee-04/dbtest/pkg/config"
)


var db *gorm.DB


type Book struct{
	gorm.Model
	Name string `json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book{
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book{
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB){
	var GetBook Book
	db := db.Where("ID=?", Id).Find(&GetBook)
	return &GetBook, db
}

func DeleteBookId(Id int64) Book{
	var book Book
	db.Where("ID=?", Id).Delete(book)
	return book
}