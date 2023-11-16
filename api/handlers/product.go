package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/spodolaks/shopifyClient/logger"
	"github.com/spodolaks/shopifyClient/network"
	"github.com/spodolaks/shopifyClient/repositories"
	"github.com/spodolaks/shopifyClient/repositories/product"
)


func NewProductHandler() *Product {
	return &Product{
		repo: product.NewProductRepo(network.Client),
	}
}

type Product struct {
	repo repositories.ProductRepo
}

func (r *Product) Fetch(c *gin.Context) {
	projects, err := r.repo.FetchProducts(c.Request.Context())
	if err != nil {
		logger.LogHandlerError(c, "Failed to fetch products", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, projects)
}