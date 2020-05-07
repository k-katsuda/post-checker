package controller

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/k-katsuda/post-checker/src/model"
)

func TaskPOST(c *gin.Context) {
	db := model.DBConnect()

	title := c.PostForm("title")
	now := time.Now()

	task := &model.Task{
		Title:     title,
		CreatedAt: now,
		UpdatedAt: now,
	}

	err := task.Save(db)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("post sent. title: %s", title)
}
