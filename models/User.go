package models

import "time"

type User struct {
	Id         int       `form:"id" json:"id"`
	Name       string    `form:"name" json:"name"`
	Age        int       `form:"age" json:"age"`
	Address    string    `form:"address" json:"address"`
	CreateTime time.Time `form:"createTime" json:"createTime"`
	updateTime time.Time `form:"createTime" json:"updateTime"`
}
