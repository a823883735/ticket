package db

import (
	"github.com/go-xorm/xorm"
	"ticket/model"
)

// 新增订单
func AddOrder(mod model.Order) (bool, error) {
	//_, err := mysqlEngine.Insert(&mod)
	//if err != nil {
	//	return false, err
	//}
	_, err := mysqlEngine.Exec("insert into `order`(u_id, t_id, t_num, date) values(?, ?, ?, current_date())", mod.UID, mod.TID, mod.TNum)
	if err != nil {
		return false, err
	}
	return true, nil
}

var (
	allSql = "select `order`.id id, ticket.title title, ticket.details details, `order`.t_num num, ticket.price price, order.u_id u_id, date " +
		"from `order`, ticket " +
		"where `order`.t_id=ticket.id order by id desc "
	mySql = "select `order`.id id, ticket.title title, ticket.details details, `order`.t_num num, ticket.price price, date " +
		"from `order`, ticket " +
		"where `order`.t_id=ticket.id and `order`.isDel=0 and `order`.u_id=? order by id desc "
)

// 查看所有订单, 若uid不为空，则查找我的所有订单
func GetAllOrderList(index, size int, uid string) (int, []model.OrderResponse, error) {
	var sql = "limit ?, ?"
	var rows *xorm.Rows
	var err error
	var total int
	if uid != "" {
		mysqlEngine.SQL("select count(*) total from `order` where isDel=0 and u_id=?", uid).Get(&total)
		rows, err = mysqlEngine.SQL( mySql + sql, uid, (index - 1) * size, size).Rows(new(model.OrderResponse))
	} else {
		mysqlEngine.SQL("select count(*) total from `order` where isDel=0").Get(&total)
		rows, err = mysqlEngine.SQL( allSql + sql, (index - 1) * size, size).Rows(new(model.OrderResponse))
	}
	if err != nil {
		return total, nil, err
	}
	defer rows.Close()
	var list []model.OrderResponse
	for rows.Next() {
		var ticketResponse model.OrderResponse
		err = rows.Scan(&ticketResponse)
		if err != nil {
			return total, nil, err
		}
		list = append(list, ticketResponse)
	}
	return total, list, err
}

// 查看我的订单
func GetMyOrderList(index, size int, uid string) (int, []model.OrderResponse, error) {
	return GetAllOrderList(index, size, uid)
}

// 删除订单
func DeleteOrder(mod model.Order) (bool, error) {
	_, err := mysqlEngine.Cols("isDel").Where("id=?", mod.Id).
		And("u_id=?", mod.UID).Update(&mod)
	if err != nil {
		return false, err
	}
	return true, nil
}