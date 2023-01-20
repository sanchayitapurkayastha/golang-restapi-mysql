package config

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
//var urlDSN = "akhil:Axlesharma@12@/simplerest?charset=utf8&parseTime=True&loc=Local"
var urlDSN = "admin:Fanatics0123@tcp(data-interns-2022-database.cclyw00l55b3.us-east-1.rds.amazonaws.com:3306)/mydb?parseTime=true"
//to make connection with the DB
func Connect() {
    d, err := gorm.Open("mysql", urlDSN)
    if err != nil {
        panic("Connection failed!")
    }

    db = d
}

func GetDB() *gorm.DB {
    return db
}