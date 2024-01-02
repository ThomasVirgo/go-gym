package controllers

import (
	"fmt"
	"gym/db"
	"net/http"

	sq "github.com/Masterminds/squirrel"
	"github.com/gin-gonic/gin"
)

type Exercise struct {
	Name string `form:"name" binding:"required"`
}

func CreateExercise(c *gin.Context) {
	if c.Request.Method == "POST" {
		exercise := Exercise{}
		err := c.ShouldBind(&exercise)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "could not parse request data",
			})
			return
		}
		fmt.Println(exercise.Name)
		sql, args, err := sq.
			Insert("exercises").Columns("name").
			Values(exercise.Name).
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
	}
}
