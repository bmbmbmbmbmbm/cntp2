package controller

import (
	"github.com/gin-gonic/gin"
)

func Document(route *gin.Engine) {
	route.GET("/document/__all__")
}
