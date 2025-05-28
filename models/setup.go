package models
import (
"gorm.io/driver/mysql"
"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
database, err := gorm.Open(mysql.Open("AthThobari:iniaththobari@tcp(127.0.0.1:3306)/go_santrikoding"),&gorm.Config{})
if err !=nil {
panic("failed to connect database")
}

database.AutoMigrate(&Post{})

DB = database

}
