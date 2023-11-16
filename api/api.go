package api

import (
	"github.com/gin-gonic/gin"
	"github.com/spodolaks/shopifyClient/api/handlers"
)

func RegisterRoutes(r *gin.Engine) {
	productHandler := handlers.NewProductHandler()
	r.GET("/products", productHandler.Fetch)
}