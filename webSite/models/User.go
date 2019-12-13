package models

import(
	"longMarch01/webSite/common"
	"fmt"
)

type User struct{
	Id int32 `form:"id" json:"id"`
	Name string `form:"name" json:"name"`
	Password string `form:"password" json:"password"`
	Status int32 `form:"status" json:"status"`
	Role int32 `form:"role" json:"role"`
}

func NewUser() *User{
	return new(User)
}

func (u *User) Add() string{
	var data string 
	sql := "INSERT INTO user(`name`,`password`,`status`,`role`) Value(?,?,?,?)"
	stmt, err := common.GetMysqlConn().Prepare(sql)
	if err != nil{
		fmt.Println(err)
		return data
	}
	defer stmt.Close()
	_, err = stmt.Exec(u.Name, u.Password, u.Status, u.Role)
	if err != nil{
		return err.Error()
	} else {
		return "success"
	}
}