package main

import (
	"cntp2/server/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	controller.DocumentController(router)
	router.Run()
}