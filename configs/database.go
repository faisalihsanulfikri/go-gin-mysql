package configs

import (
    "fmt"
    "service-catatin/models"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

/*DB is connected database object*/
var DB *gorm.DB

func ConnectDB()  {
    connectionString := EnvDbUsername() + ":" + EnvDbPassword() + "@tcp(" + EnvDbHost() + ":" + EnvDbPort() + ")/" + EnvDbName() + "?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

    if err != nil {
        fmt.Println(err.Error())
		panic("Failed to connect database")
    }

    fmt.Println("Connected to database")

	db.AutoMigrate([]models.User{})
	DB = db
}

// GetDB helps you to get a connection
func GetDB() *gorm.DB {
	return DB
}

