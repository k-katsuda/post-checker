package main

import (
	"github.com/gin-gonic/gin"
	"github.com/k-katsuda/post-checker/src/controller"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.POST("/tasks", controller.TasksPOST)
	}

	// nginxのreverse proxy設定
	router.Run(":9000")
}
