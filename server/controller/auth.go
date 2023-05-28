package controller

import (
	"net/http"

	"github.com/NHollmann/yotei/model"
	"github.com/gin-gonic/gin"
)

// Login user
// @Summary Login user
// @Description Login an existing user with a username and password
// @Tags Authentication
// @Success 200 {string} hello
// @Router /login [post]
func (server *YoteiServer) handleLogin(c *gin.Context) {

	users := model.UserGetAll(server.db)
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

// Logout user
// @Summary Logout user
// @Description Logout the current user
// @Tags Authentication
// @Success 200 {string} hello
// @Router /logout [post]
func (server *YoteiServer) handleLogout(c *gin.Context) {

	users := model.UserGetAll(server.db)
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}
