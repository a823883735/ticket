package db

import "ticket/model"

// 新增订单
func AddOrder(mod model.Order) (bool, error) {

	return true, nil
}

// 查看所有订单
func GetAllOrderList(index, size int) ([]model.Order, error) {
	var list []model.Order
	err := mysqlEngine.Limit(size, index * size).Find(&list)
	return list, err
}

// 查看我的订单
func GetMyOrderList(index, size int) ([]model.Order, error) {
	var list []model.Order
	//err := mysqlEngine.
	return list, nil
}