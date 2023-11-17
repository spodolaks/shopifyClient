package repositories

import (
	"context"

	"github.com/spodolaks/shopifyClient/models"
)

type ProductRepo interface {
	FetchProducts(ctx context.Context) ([]models.Product, error)
	UpdateProduct(ctx context.Context, product models.Product) (error)
}