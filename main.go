package main

import (
    "fmt"
    "service-catatin/configs"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    
    //run database
    configs.ConnectDB()

    //routes
    

    fmt.Println("SEVICE CATATIN IS RUNNING")
    router.Run("localhost:6000")
}