package router

import (
	"token/controller"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func Router() error {

	r := gin.Default()
	// 註冊帳號--------------------------------
	r.POST("/Register", controller.Register)
	// 登錄帳號--------------------------------
	r.POST("/login", controller.Login)
	// 登出-----------------------------------
	r.GET("/logout", controller.Logout)
	r.GET("/middleware", controller.LoginStruct)

	err := r.Run(":8082")
	if err != nil {
		log.Fatal().Err(err).Msg("8082 Error")
		return err
	}
	return nil
}
