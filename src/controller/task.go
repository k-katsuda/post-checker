package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/k-katsuda/post-checker/src/model"
)

func TasksPOST(c *gin.Context) {
	db := model.DBConnect()

	postFlg := c.PostForm("postFlg")
	postCount := c.PostForm("postCount")

	task := &model.Task{
		postFlg:   postFlg,
		postCount: postCount,
	}

	err := task.Save(db)
	if err != nil {
		panic(err.Error())
	}
}
