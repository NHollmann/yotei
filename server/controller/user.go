package controller

import (
	"net/http"
	"strconv"

	"github.com/NHollmann/yotei/model"
	"github.com/gin-gonic/gin"
)

// Get all users
// @Summary Get all users
// @Description Only works for administrators
// @Tags User
// @Success 200 body {string} hello
// @Router /user [get]
func (server *YoteiServer) handleUserList(c *gin.Context) {

	// TODO Check is admin

	users := model.UserGetAll(server.db)
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

// Get one user
// @Summary Get one user
// @Description Only works for administrators and the user itself
// @Tags User
// @Param userId path uint true "User ID"
// @Success 200 {string} hello
// @Router /user/{userId} [get]
func (server *YoteiServer) handleUserGet(c *gin.Context) {

	// TODO Check is admin or user with same id

	userIdStr := c.Param("userId")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "userId must be an integer",
		})
		return
	}

	user, err := model.UserGetOne(server.db, uint(userId))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "userId not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"hello": user,
	})
}

// Create a new user
// @Summary Create a new user
// @Description Only works for administrators
// @Tags User
// @Param request body controller.handleUserCreate.userDataType true "Data for new user"
// @Success 200 {object} object{userId=integer} "User ID of newly created user"
// @Failure 400 {object} object{error=string} "Error message"
// @Router /user [post]
func (server *YoteiServer) handleUserCreate(c *gin.Context) {

	// TODO Check is admin

	type userDataType struct {
		Name     string `json:"name" binding:"required" example:"Max Mustermann"`
		Username string `json:"username" binding:"required" example:"mmustermann"`
		Password string `json:"password" binding:"required" example:"catsAreAwesome"`
		IsAdmin  bool   `json:"isAdmin" example:"false"`
	}

	var userData userDataType
	if err := c.ShouldBindJSON(&userData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, err := model.UserCreate(
		server.db,
		userData.Name,
		userData.Username,
		userData.Password,
		userData.IsAdmin,
		true,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"userId": userId,
	})
}

// Update an user
// @Summary Update an user
// @Description Only works for administrators and the user itself
// @Tags User
// @Param userId path uint true "User ID"
// @Param request body controller.handleUserUpdate.userDataType true "Data for updated user"
// @Success 200 {object} object{userId=integer} "User ID of updated user"
// @Failure 400 {object} object{error=string} "Error message"
// @Router /user/{userId} [put]
func (server *YoteiServer) handleUserUpdate(c *gin.Context) {

	// TODO Check is admin

	type userDataType struct {
		Name                   string `json:"name" binding:"required" example:"Max Mustermann"`
		Username               string `json:"username" binding:"required" example:"mmustermann"`
		Password               string `json:"password" example:"catsAreAwesome"`
		IsAdmin                bool   `json:"isAdmin" example:"false"`
		PasswordChangeRequired bool   `json:"passworChangeRequired" example:"false"`
	}

	var userData userDataType
	if err := c.ShouldBindJSON(&userData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userIdStr := c.Param("userId")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "userId must be an integer",
		})
		return
	}
	err = model.UserUpdate(
		server.db,
		uint(userId),
		userData.Name,
		userData.Username,
		userData.Password,
		userData.IsAdmin,
		userData.PasswordChangeRequired,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"userId": userId,
	})
}

// Delete an user
// @Summary Delete an user
// @Description Only works for administrators and the user itself
// @Tags User
// @Param userId path uint true "User ID"
// @Success 200 {string} hello
// @Router /user/{userId} [delete]
func (server *YoteiServer) handleUserDelete(c *gin.Context) {

	// TODO Check is admin

	userIdStr := c.Param("userId")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "userId must be an integer",
		})
		return
	}

	err = model.UserDelete(server.db, uint(userId))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"userId": userId,
	})
}
