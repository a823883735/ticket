package model

type Order struct {
	Id int `xorm:"'id' not null pk autoincr comment('主键')" json:"id"`
	TID string `xorm:"'t_id' int" json:"t_id"`
	UID string `xorm:"'u_id' varchar(20)"`
	IsDel bool `xorm:"'isDel' bit default 0" json:"-"`
	Date string `xorm:"'date' date created" json:"date"`
	TNum int `xorm:"'t_num' int default 1" json:"t_num"`
}

func (Order) TableName() string {
	return "order"
}


type OrderResponse struct {
	Id int `xorm:"'id' int" json:"id"`
	Title string `xorm:"'title'" json:"title"`
	Details string `xorm:"'details'" json:"details"`
	Price float32 `xorm:"price" json:"price"`
	Num int `xorm:"num" json:"num"`
	UId string `xorm:"'u_id'" json:"u_id"`
	Date string `xorm:"'date' date created" json:"date"`
}