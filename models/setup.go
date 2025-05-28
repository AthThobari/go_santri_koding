package models
import (
"gorm.io/driver/postgres"
"gorm.io/gorm"
"os"
"log"
)

var DB *gorm.DB

func ConnectDatabase() {
dsn := os.Getenv("POSTGRES_URL") // Ambil dari environtment
database, err := gorm.Open(postgres.Open(dsn),&gorm.Config{})
if err !=nil {
log.Fatal("failed to connect to database: ", err)
}

database.AutoMigrate(&Post{})

DB = database

}
