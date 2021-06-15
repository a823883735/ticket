package model

type User struct {
	Id string `xorm:"'id' varchar(20) unique notnull comment('账号')" json:"userName"`
	Pwd string `xorm:"'pwd' varchar(20) comment('密码')" json:"password"`
}

func (User) TableName() string {
	return "user"
}