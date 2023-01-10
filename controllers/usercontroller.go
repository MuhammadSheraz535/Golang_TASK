package controllers

import (
	"net/http"

	"githhub.com/shery535/models"
	"githhub.com/shery535/services"
	"github.com/gin-gonic/gin"
)

// Define routes in controllers

type Usercontroller struct {
	UserService services.UserService
}

func New(userService services.UserService) Usercontroller {

	return Usercontroller{
		UserService: userService,
	}
}

func (uc *Usercontroller) CreateUser(ctx *gin.Context) {

	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.UserService.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return

	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})

}

func (uc *Usercontroller) GetUser(ctx *gin.Context) {
	username := ctx.Param("name")
	user, err := uc.UserService.GetUser(&username)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, user)

}

func (uc *Usercontroller) GetAll(ctx *gin.Context) {

	users, err := uc.UserService.GetAll()
	if err != nil {

		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, users)

}

func (uc *Usercontroller) UpdateUser(ctx *gin.Context) {

	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})

	}
	err := uc.UserService.UpdateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})

	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})

}

func (uc *Usercontroller) DeleteUser(ctx *gin.Context) {

	username := ctx.Param("name")
	err := uc.UserService.DeleteUser(&username)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return

	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})

}

func (uc *Usercontroller) RegisterUserRoutes(rg *gin.RouterGroup) {
	userroute := rg.Group("/user")
	userroute.GET("/get/:name", uc.GetUser)
	userroute.GET("/getall", uc.GetAll)
	userroute.PATCH("/update", uc.UpdateUser)
	userroute.PATCH("/delete/:name", uc.DeleteUser)
	userroute.POST("/create", uc.CreateUser)

}
