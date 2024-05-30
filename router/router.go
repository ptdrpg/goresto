package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ptdrpg/resto/controller"
)

type Router struct {
	R *gin.Engine
	C *controller.Controller
}

func NewRouter(r *gin.Engine, c *controller.Controller) *Router {
	return &Router{
		R: r,
		C: c,
	}
}

func (r *Router) RegisterRouter() {
	apiR := r.R.Group("/api")
	v1 := apiR.Group("/v1")

	ub := v1.Group("/customer")
	ub.GET("", r.C.FindAllUsers)
	ub.GET("/:id", r.C.FindUserById)
	ub.POST("/create", r.C.CreateUser)
	ub.PUT("/update/:id", r.C.UpdateUser)
	ub.DELETE("/delete/:id", r.C.DeleteUser)

	sb := ub.Group("/staff")
	sb.GET("/get-all", r.C.FindAllStaff)
	sb.GET("/get-one/:id", r.C.FindStaffById)
	sb.POST("/create", r.C.CreateStaff)
	sb.PUT("/update/:id", r.C.UpdateStaff)
	sb.DELETE("/delete/:id", r.C.DeleteStaff)
	sb.PUT("/update-pass/:id", r.C.UpdatePassword)
	sb.POST("/login", r.C.Login)

	eb := v1.Group("/employee")
	eb.GET("/get-all", r.C.FindAllEmployee)
	eb.GET("/:id", r.C.FindEmployeeById)
	eb.POST("/create", r.C.CreateEmployee)
	eb.PATCH("/update/:id", r.C.UpdateEmployee)
	eb.DELETE("/delete/:id", r.C.DeleteEmployee)

	ib := v1.Group("/items")
	ib.GET("/get-all", r.C.FindAllItems)
	ib.GET("/:id", r.C.FindItemById)
	ib.POST("/create", r.C.CreateItems)
	ib.PATCH("/update/:id", r.C.UpdateItems)
	ib.DELETE("/delete/:id", r.C.DeleteItems)

	tb := v1.Group("/ticket")
	tb.GET("/get-all", r.C.FindAllTicket)
	tb.GET("/:id", r.C.FindTicketById)
	tb.POST("/create", r.C.CreateTickets)
	tb.DELETE("/delete/:id", r.C.DeleteTickets)
}
