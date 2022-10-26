package main

import (
    "fmt"
    "service-catatin/configs"
    "service-catatin/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    
    //run database
    configs.ConnectDB()

    //routes    
    routes.Api(router)

    fmt.Println("SEVICE CATATIN IS RUNNING")
    router.Run("localhost:6000")
}