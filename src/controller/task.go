package model

type Task struct {
	ID        uint `json:"id"`
	PostFlg   uint `json:"postFlg"`
	PostCount uint `json:"postCount"`
}
