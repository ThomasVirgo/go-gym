package controllers

import (
	"net/http"

	"gym/query"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	exercises, err := query.GetAllExercises()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	gym_sessions, err := query.GetGymSessions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"exercises": exercises, "gym_sessions": gym_sessions,
	})
}
