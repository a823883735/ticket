package router

import "ticket/service"

func RegisterOrderRouter() {
	orderRouter := ginEngine.Group("/v1/order")
	{
		orderRouter.POST("/addOrder", service.AddOrder)
		orderRouter.GET("/orderList", service.AllOrderList)
		orderRouter.GET("/myOrderList", service.MyOrderList)
		orderRouter.DELETE("/deleteOrderById", service.DeleteOrderById)
	}
}