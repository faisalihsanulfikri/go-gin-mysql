package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
)

// var db *gorm.DB = configs.DB

func main() {
    router := gin.Default()
    
    //run database
    // configs.ConnectDB()
		// defer db.Close()

    //routes
    

    fmt.Println("SEVICE CATATIN IS RUNNING")
    router.Run("localhost:6000")
}