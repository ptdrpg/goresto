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
}
