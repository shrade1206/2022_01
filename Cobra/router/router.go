package router

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Router() {

}

func G(c *gin.Context) {
	r := gin.Default()
	r.GET("/")
	r.GET("/")
	r.GET("/")

	err := r.Run(":8080")
	if err != nil {
		log.Println(err)
		return
	}
}
