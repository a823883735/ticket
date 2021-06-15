package model

type Order struct {
	Id string `xorm:"'id' not null pk autoincr comment('主键')" json:"id"`
	TID string `xorm:"'t_id' int" json:"t_id"`
	UID string `xorm:"'u_id' varchar(20)"`
}

func (Order) TableName() string {
	return "order"
}