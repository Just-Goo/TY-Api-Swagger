package handlers

import (
	"log"
	"net/http"

	"github.com/TY-Api-Swagger/model"
	"github.com/gin-gonic/gin"
)

type userResponse struct {
	Data []model.User `json:"data"`
}

type errorResponse struct {
	Message string
}

// GetUsers return list of all users from the database
// @Summary		 return list of all users
// @Description 	return list of all users from the database
// @Tags Users
// @Success 200 {object} userResponse
// @Router /users [get]
func GetUsers(c *gin.Context) {
	users := model.ListUsers()

	c.JSON(http.StatusOK, userResponse{Data: users})

}

// CreateUser create new usr
// @Summary		 create new user
// @Description 	create new user
// @Security bearerToken
// @Tags Users
// @Accept json
// @Produce json
// @Param user body model.User true "User"
// @Success 201 {object} model.User
// @Failure 400 {object} errorResponse
// @Router /users [post]
func CreateUser(c *gin.Context)  {
	var req model.User

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse{Message: err.Error()})
		return
	}

	err := model.CreateUser(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse{Message: err.Error()})
		return
	}

	auth := c.GetHeader("Authorization")
	log.Println(auth)

	c.JSON(http.StatusOK, req)
}