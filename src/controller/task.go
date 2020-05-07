package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/k-katsuda/post-checker/src/model"
)

func TasksPOST(c *gin.Context) {
	db := model.DBConnect()
	fmt.Printf("%v", db)
	// postFlg := c.PostForm("postFlg")
	// postCount := c.PostForm("postCount")

	// task := &model.Task{
	// 	Postflg:   postFlg,
	// 	Postcount: postCount,
	// }

	// err := task.Save(db)
	// if err != nil {
	// 	panic(err.Error())
	// }
}
