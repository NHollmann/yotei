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
// @Success 200 {string} hello
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
// @Success 200 {string} hello
// @Router /user [post]
func (server *YoteiServer) handleUserCreate(c *gin.Context) {

	// TODO Check is admin

	users := model.UserGetAll(server.db)
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

// Update an user
// @Summary Update an user
// @Description Only works for administrators and the user itself
// @Tags User
// @Param userId path uint true "User ID"
// @Success 200 {string} hello
// @Router /user/{userId} [put]
func (server *YoteiServer) handleUserUpdate(c *gin.Context) {

	// TODO Check is admin

	users := model.UserGetAll(server.db)
	c.JSON(http.StatusOK, gin.H{
		"users": users,
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

	users := model.UserGetAll(server.db)
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}
