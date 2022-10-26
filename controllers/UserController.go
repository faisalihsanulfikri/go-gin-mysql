package UserController

import (
    "service-catatin/configs"
    "service-catatin/models"
    "net/http"
    
    "github.com/gin-gonic/gin"
    "github.com/go-playground/validator/v10"
)

var validate = validator.New()

func Index() gin.HandlerFunc {
    return func(c *gin.Context) {
        db := configs.GetDB()
        users := []models.User{}
        err := db.Find(&users).Error

        if err != nil {
            c.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "message": "Data not found", "data": err})
            return
        }
        c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "Success", "data": users})
    }
}

func Show() gin.HandlerFunc {
    return func(c *gin.Context) {
        db := configs.GetDB()
        userId := c.Param("userId")
        user := models.User{}
        result := db.First(&user, userId)

        if result.Error != nil {
            c.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "message": "Data not found", "data": nil})
            return
        }
        c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "Success", "data": user})
    }
}

func Store() gin.HandlerFunc {
    return func(c *gin.Context) {
        db := configs.GetDB()
        user := models.User{}
        
        if err := c.BindJSON(&user); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "Invalid data", "data": nil})
            return
        }

        result := db.Create(&user)

        if result.Error != nil {
            c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "Unable to create data", "data": nil})
            return
        }
        c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "Success", "data": user})
    }
}

func Update() gin.HandlerFunc {
    return func(c *gin.Context) {

        type payloadType struct {
            Name            string          `json:"name,omitempty" validate:"required"`
            Email           string          `json:"email,omitempty" validate:"required"`
            UserRoleId      uint            `json:"user_role_id,omitempty" validate:"required"`
        }

        payload := payloadType{}

        db := configs.GetDB()
        userId := c.Param("userId")
        user := models.User{}

        if err := c.BindJSON(&payload); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "Invalid data", "data": nil})
            return
        }

        result := db.First(&user, userId)
        if result.Error != nil {
            c.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "message": "Data not found", "data": nil})
            return
        }

        user.Name = payload.Name
        user.Email = payload.Email
        user.UserRoleId = payload.UserRoleId
        result = db.Save(&user)

        if result.Error != nil {
            c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "Unable to update data", "data": nil})
            return
        }
        c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "Success", "data": user})
    }
}

func Destroy() gin.HandlerFunc {
    return func(c *gin.Context) {
        db := configs.GetDB()
        userId := c.Param("userId")
        user := models.User{}

        result := db.First(&user, userId)
        if result.Error != nil {
            c.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "message": "Data not found", "data": nil})
            return
        }

        result = db.Delete(&user)
        if result.Error != nil {
            c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "Unable to delete data", "data": nil})
            return
        }
        c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "Success", "data": user})
    }
}