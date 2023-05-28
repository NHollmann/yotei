package controller

import (
	"net/http"
	"strconv"

	"github.com/NHollmann/yotei/model"
	"github.com/gin-gonic/gin"
)

// Get all events
// @Summary Get all events
// @Description Get all events based on permissions
// @Tags Event
// @Success 200 {string} hello
// @Router /event [get]
func (server *YoteiServer) handleEventList(c *gin.Context) {

	// TODO Check is admin
	// TODO Check if get params are set for user selection if admin

	users := model.UserGetAll(server.db)
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

// Get one event
// @Summary Get one event
// @Description Can be used by anyone who haves the access key
// @Tags Event
// @Param accessKey path string true "Event access key"
// @Success 200 {string} hello
// @Router /event/{accessKey} [get]
func (server *YoteiServer) handleEventGet(c *gin.Context) {

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

// Create a new event
// @Summary Create a new event
// @Description Every user can create a new event
// @Tags Event
// @Success 200 {string} hello
// @Router /event [post]
func (server *YoteiServer) handleEventCreate(c *gin.Context) {

	// TODO Check is admin

	users := model.UserGetAll(server.db)
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

// Update an event
// @Summary Update an event
// @Description Only the creator and all administrators can update an event
// @Tags Event
// @Param accessKey path string true "Event access key"
// @Success 200 {string} hello
// @Router /event/{accessKey} [put]
func (server *YoteiServer) handleEventUpdate(c *gin.Context) {

	// TODO Check is admin

	users := model.UserGetAll(server.db)
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

// Delete an event
// @Summary Delete an event
// @Description Only the creator and all administrators can delete an event
// @Tags Event
// @Param accessKey path string true "Event access key"
// @Success 200 {string} hello
// @Router /event/{accessKey} [delete]
func (server *YoteiServer) handleEventDelete(c *gin.Context) {

	// TODO Check is admin

	users := model.UserGetAll(server.db)
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

// Add a participant to an existing event
// @Summary Add a participant to an existing event
// @Description Everybody can use this route, if the user is logged in, the participant will be linked to the user
// @Tags Event
// @Param accessKey path string true "Event access key"
// @Success 200 {string} hello
// @Router /event/{accessKey} [post]
func (server *YoteiServer) handleEventCreateParticipant(c *gin.Context) {

	// TODO Check is admin

	users := model.UserGetAll(server.db)
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

// Update a participant of the event
// @Summary  Update a participant of the current event
// @Description Everyone can update a participant except there is an user linked to it
// @Tags Event
// @Param accessKey path string true "Event access key"
// @Success 200 {string} hello
// @Router /event/{accessKey} [patch]
func (server *YoteiServer) handleEventUpdateParticipant(c *gin.Context) {

	// TODO Check is admin

	users := model.UserGetAll(server.db)
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}
