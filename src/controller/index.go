package controller

import (
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
		var postFlg int
		var postCount int

		err = result.Scan(&id, &postFlg, &postCount)
		if err != nil {
			panic(err.Error())
		}

		task.ID = id
		task.postFlg = postFlg
		task.postCount = postCount
	}
	return task
}
