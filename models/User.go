package models

type User struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Age        int    `json:"age"`
	Addr       string `json:"addr"`
	CreateTime int    `json:"createTime"`
	updateTime int    `json:"updateTime"`
}
