package controllers

import (
	"net/http"

	"github.com/carboncody/go-bootstrapper/initializers"
	"github.com/carboncody/go-bootstrapper/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

func NewUserController(DB *gorm.DB) UserController {
	return UserController{DB}
}

func (uc *UserController) GetMe(ctx *gin.Context) {
    currentUser, ok := ctx.Get("currentUser")
    if !ok {
        ctx.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "User not found"})
        return
    }

    user, ok := currentUser.(models.User)
    if !ok {
        ctx.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "Invalid user data"})
        return
    }

    if err := initializers.DB.Preload("").First(&user, user.ID).Error; err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "User not found"})
        return
    }

    for i := range user.Workspaces {
        if err := initializers.DB.Model(&user.Workspaces[i]).Association("Workspace").Find(&user.Workspaces[i].WorkspaceId); err != nil {
            ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": "Failed to load Workspace"})
            return
        }
    }

    ctx.JSON(http.StatusOK, user)
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
    var payload models.CreateUserPayload

    if err := ctx.ShouldBindJSON(&payload); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
        return
    }

	if err := initializers.DB.Create(&payload).Error; err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
        return
    }

    ctx.JSON(http.StatusCreated, gin.H{"status": "success", "message": "User created successfully"})
}


func (uc *UserController) UpdateUser(ctx *gin.Context) {
    currentUser, ok := ctx.Get("currentUser")
    if !ok {
        ctx.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "User not found"})
        return
    }

    user, ok := currentUser.(models.User)
    if !ok {
        ctx.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "Invalid user data"})
        return
    }

    // Bind the request body to the update payload
    var updatePayload models.UpdateUserPayload
    if err := ctx.ShouldBindJSON(&updatePayload); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
        return
    }

    user.ProfilePicture = &updatePayload.ProfilePicture

    if err := initializers.DB.Save(&user).Error; err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": "Failed to update user"})
        return
    }

    // * Return the updated user
    ctx.JSON(http.StatusOK, user)
}
