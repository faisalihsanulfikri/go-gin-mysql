package UserController

import (
    // "context"
    // "fmt"
    "service-catatin/configs"
    "service-catatin/models"
    "net/http"
    // "time"
    
    "github.com/gin-gonic/gin"
    "github.com/go-playground/validator/v10"
    "gorm.io/gorm"
)

var db *gorm.DB = configs.GetDB()
var validate = validator.New()

func Show() gin.HandlerFunc {
    return func(c *gin.Context) {
        userId := c.Param("userId")
        user := models.User{}
        err := db.Select("users.*").First(&user, userId).Error

        if err != nil {
            c.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "message": "Data not found", "data": nil})
            return
        }
        c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "Success", "data": user})
    }
}