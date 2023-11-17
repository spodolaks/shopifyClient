package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spodolaks/shopifyClient/api"
)

func main() {
	r := gin.Default()
	api.RegisterRoutes(r)
	r.Run()
}