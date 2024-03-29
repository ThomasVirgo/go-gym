package controllers

import (
	"fmt"
	"gym/db"
	"gym/query"
	"net/http"

	sq "github.com/Masterminds/squirrel"
	"github.com/gin-gonic/gin"
)

type Set struct {
	Exercise int     `form:"exercise" binding:"required"`
	Weight   float32 `form:"weight" binding:"required"`
	Reps     int     `form:"reps" binding:"required"`
}

func CreateSet(c *gin.Context) {
	if c.Request.Method == "POST" {
		set := Set{}
		err := c.ShouldBind(&set)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "could not parse request data",
			})
			return
		}
		sql, args, err := sq.
			Insert("sets").Columns("exercise", "weight", "reps").
			Values(set.Exercise, set.Weight, set.Reps).
			ToSql()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("%s", err.Error()),
			})
			return
		}
		_, err = db.DB.Exec(sql, args...)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("%s", err.Error()),
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
		c.HTML(http.StatusCreated, "gym_sessions.html", gin.H{
			"gym_sessions": gym_sessions,
		})
	}
}
