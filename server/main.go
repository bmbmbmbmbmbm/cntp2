package main

import (
	"cntp2/accounts"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/accounts/login", accounts.accessAccount)
	router.Run()
}
