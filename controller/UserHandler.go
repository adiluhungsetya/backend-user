package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"user-go-test/model"
	"user-go-test/service"
	"user-go-test/utils"
)

type UserHandler struct {
	userService service.UserService
}

func CreateUserHandler(r *gin.Engine, userService service.UserService)  {
	userHandler := UserHandler{userService: userService}
	r.GET("/user", userHandler.getAllUser)
	r.GET("/user/:id", userHandler.getUserById)
	r.POST("/user", userHandler.insertNewUser)
	r.PUT("/user", userHandler.updateUser)
}

func (u UserHandler) getAllUser(c *gin.Context) {
	user, err := u.userService.GetAllUser()
	if err != nil {
		utils.HandleError(c, http.StatusBadGateway, err.Error())
	}
	utils.HandleSuccess(c, user)
}

func (u UserHandler) insertNewUser(c *gin.Context)  {
	var body model.User

	err := c.Bind(&body)
	if err != nil {
		utils.HandleError(c, http.StatusBadGateway, err.Error())
	}
	if body.Name == "" || body.Ktp == "" {
		utils.HandleError(c, http.StatusBadRequest, "There's empty field")
	}

	user, err := u.userService.InsertNewUser(&body)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, "Server error")
	}
	utils.HandleSuccess(c, user)
}

func (u UserHandler) updateUser(c *gin.Context)  {
	var body model.User
	err := c.Bind(&body)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
	}
	body.Id, err= strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, "Param must integer")
	}
	if body.Name == "" || body.Ktp == "" {
		utils.HandleError(c, http.StatusBadRequest, "There's empty field")
	}

	user, err := u.userService.UpdateUser(body.Id,&body)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, "Server error")
	}
	utils.HandleSuccess(c, user)
}

func (u UserHandler) getUserById(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, "Id must be integer")
	}
	user,err := u.userService.GetUser(id)

	utils.HandleSuccess(c, user)
}