package model

type Ticket struct {
	Id string `xorm:"'id' not null pk autoincr comment('主键')" json:"phone"`
	Title string `xorm:"'title' varchar(20) not null comment('标题')" json:"title"`
	Details string `xorm:"'details' text comment('描述')" json:"details"`
	Price float32 `xorm:"'price' decimal(10,2) default 0.00 comment('价格')" json:"price"`
	IsDel bool `xorm:"'isDel' bit default 0" json:"-"`
}

func (Ticket) TableName() string {
	return "ticket"
}