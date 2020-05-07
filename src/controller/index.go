package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/k-katsuda/post-checker/src/model"
)

// タスク検索
func FindByID(id uint) model.Task {
	db := model.DBConnect()
	result, err := db.Query("SELECT * FROM post WHERE id = ?", id)
	if err != nil {
		panic(err.Error())
	}
	task := model.Task{}
	for result.Next() {
		var postFlg uint
		var postCount uint

		err = result.Scan(&id, &postFlg, &postCount)
		if err != nil {
			panic(err.Error())
		}

		task.ID = id
		task.Postflg = postFlg
		task.Postcount = postCount
	}
	return task
}

func TasksPOST(c *gin.Context) {
	db := model.DBConnect()

	postFlg := c.PostForm("postFlg")
	postCount := c.PostForm("postCount")

	_, err := db.Exec("UPDATE post SET postFlg = ?, postCount=? WHERE id = 1", postFlg, postCount)
	if err != nil {
		panic(err.Error())
	}

	task := FindByID(1)

	fmt.Println(task)
	c.JSON(http.StatusOK, gin.H{"task": task})
}
