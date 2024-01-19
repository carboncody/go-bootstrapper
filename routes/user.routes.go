package routes

import (
	"github.com/carboncody/go-bootstrapper/controllers"
	"github.com/carboncody/go-bootstrapper/middleware"
	"github.com/gin-gonic/gin"
)

type UserRouteController struct {
	userController controllers.UserController
}

func NewRouteUserController(userController controllers.UserController) UserRouteController {
	return UserRouteController{userController}
}

func (uc *UserRouteController) UserRoute(rg *gin.RouterGroup) {

	router := rg.Group("users")
	router.GET("/me", middleware.DeserializeUser(), uc.userController.GetMe)
	router.POST("/register", uc.userController.CreateUser)
	router.PATCH("/update", middleware.DeserializeUser(), uc.userController.UpdateUser)
}
