package routes

import (
    "service-catatin/controllers"
    "github.com/gin-gonic/gin"
)

func Api(router *gin.Engine) {
    router.GET("/user", UserController.Index())
    router.GET("/user/:userId", UserController.Show())
    router.POST("/user", UserController.Store())
    router.PUT("/user/:userId", UserController.Update())
    router.DELETE("/user/:userId", UserController.Destroy())
}