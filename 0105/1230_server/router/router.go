package router

import (
	"server/controller"
	"server/middleware"

	"github.com/gin-gonic/gin"
)

func Router() {
	r := gin.Default()
	r.Use(middleware.Cors())
	r.POST("/Register", controller.Register)
	r.POST("/login", controller.Login)

	r.GET("/middlewareAuth", controller.MiddlewareAuth)
	r.GET("/logout", controller.Logout)

	// ------------------------------------------------
	r.POST("/insert", controller.Insert)
	r.GET("/getpage", controller.Getpage)
	// r.GET("/get", controller.Get)
	r.PUT("/put/:id", controller.Put)
	r.DELETE("/del/:id", controller.Del)
	err := r.Run(":8084")
	if err != nil {
		return
	}
}
