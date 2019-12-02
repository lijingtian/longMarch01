package models

type User struct{
	Id int64 `json:"id"`
	Name string `json:"name"`
	Password string `json:"password"`
	Status int32 `json:"user_status"`
	Role int32 `json:"role"`
}

func (m *User) GetUserInfo(){
	
}