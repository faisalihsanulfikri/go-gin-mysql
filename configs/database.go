package configs

import (
    "fmt"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

func ConnectDB() *gorm.DB  {
    connectionString := EnvDbUsername() + ":" + EnvDbPassword() + "@tcp(" + EnvDbHost() + ":" + EnvDbPort() + ")/" + EnvDbName() + "?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

    if err != nil {
        fmt.Println(err.Error())
		panic("Failed to connect database")
    }

    fmt.Println("Connected to database")

    return db
}

//Client instance
var DB *gorm.DB = ConnectDB()