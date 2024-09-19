package controller

import (
	"github.com/gin-gonic/gin"
)

func DocumentController(route *gin.Engine) {
	route.GET("/document/__all__")
}