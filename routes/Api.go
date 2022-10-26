package routes

import (
    "service-catatin/controllers/UserController"
    "github.com/gin-gonic/gin"
)

func Api(router *gin.Engine) {
    router.GET("/user/:userId", UserController.Show())
}