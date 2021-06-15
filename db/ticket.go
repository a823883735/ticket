package db

import "ticket/model"

// 新增票
func AddTicket(mod model.Ticket) (bool, error) {
	_, err := mysqlEngine.Insert(mod)
	if err != nil {
		return false, err
	}
	return true, nil
}

// 删除票
func DeleteTicket(modId int) (bool, error) {
	rows, err := mysqlEngine.Where("id=?", modId).Delete(model.Ticket{})
	if err != nil {
		return true, err
	}
	if rows == 0 {
		return false, nil
	}
	return true, nil
}

// 查看票
func GetTicket(modId int) (model.Ticket, error) {
	var mod = model.Ticket{}
	_, err := mysqlEngine.Where("id=?", modId).Get(&mod)
	if err != nil {
		return mod, err
	}
	return mod, nil
}