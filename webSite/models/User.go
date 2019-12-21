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

func (u *User) CheckLogin() {
	where := " WHERE 1=1"
	if u.Name != "" && u.Password != "" {
		where += fmt.Sprintf(" AND name='%s' AND password='%s'", u.Name, u.Password)
	} else {
		return 
	}
	sql := fmt.Sprintf("SELECT id,name,password,status,role FROM user %s", where)
	err := common.GetMysqlConn().QueryRow(sql).Scan(&u.Id, &u.Name, &u.Password, &u.Status, &u.Role)
	if err != nil{
		fmt.Println(err)
		return
	}
}