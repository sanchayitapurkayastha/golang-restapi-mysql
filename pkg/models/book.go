package models

import (
    "github.com/jinzhu/gorm"
    "github.com/sanchayitapurkayastha/golang-rest-api-mysql/pkg/config"
)

var db *gorm.DB

type Book struct {
    gorm.Model
    Name string `gorm:""json:"name"`
    Author string `json:"author"`
    Publication string `json:"publication"`
}

//initialise ur db
func init() {
    config.Connect()
    db = config.GetDB()
    db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
    db.NewRecord(b)
    db.Create(&b)
    return b
}

func GetBooks() []Book {
    var b []Book
    db.Find(&b)
    return b
}

func GetBookById(id int64) (*Book, *gorm.DB) {
    var b Book
    db := db.Where("ID=?", id).Find(&b)
    return &b, db
}

func DeleteBook(id int64) Book {
    var b Book
    db.Where("ID=?", id).Delete(b)
    return b
}

//didnt create update func as we'll be first finding the book with the id then deleting it and then creating the book