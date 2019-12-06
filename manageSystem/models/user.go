package models

import(
	"longMarch01/manageSystem/common"
	"fmt"
)


type User struct{
	Id int64 `form:"id" json:"id"`
	Name string `form:"name" json:"name"`
	Password string `form:"password" json:"password"`
	Status int32 `form:"status" json:"status"`
	Role int32 `form:"role" json:"role"`
}

func NewUser() *User{
	return new(User)
}

func (u *User) GetUserInfo(){
	var whereSql string = " WHERE 1=1 "
	if u.Name != "" {
		whereSql += fmt.Sprintf(" AND name='%s'", u.Name)
	}
	if u.Password != "" {
		whereSql += fmt.Sprintf(" AND password='%s'", u.Password)
	}
	sql := fmt.Sprintf("SELECT id,name,password,status,role FROM user %s", whereSql)
	rows, err := common.GetDbConn().Query(sql)
	if err != nil{
		fmt.Println(err)
		return
	}
	defer rows.Close()
	for rows.Next(){
		err := rows.Scan(&u.Id, &u.Name, &u.Password, &u.Status, &u.Role)
		if err != nil{
			fmt.Println(err)
			return
		}
	}
}

func (u *User) InsertInto() string{
	sql := "INSERT INTO user(`name`,`password`,`status`,`role`) Value(?,?,?,?)"
	stmt, err := common.GetDbConn().Prepare(sql)
	if err != nil{
		return err.Error()
	}
	defer stmt.Close()
	_, err = stmt.Exec(u.Name, u.Password, u.Status, u.Role)
	if err != nil{
		return err.Error()
	} else {
		return "success"
	}
}