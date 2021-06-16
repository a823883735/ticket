package router

import "ticket/service"

func RegisterTicketRouter() {
	ticketRouter := ginEngine.Group("/v1/ticket")
	{
		ticketRouter.POST("/addTicket", service.AddTicket)
		ticketRouter.GET("/ticketList", service.TicketList)
		ticketRouter.DELETE("/delTicketList", service.DelTicket)
	}
}