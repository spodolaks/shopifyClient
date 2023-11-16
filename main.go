package main

import (
	"context"
	"fmt"

	"github.com/spodolaks/shopifyClient/network"
	"github.com/spodolaks/shopifyClient/repositories/product"
)

func main() {
	network.Initialize();
	ctx := context.Context(context.Background())
	productRepo := product.NewProductRepo(network.Client)
	products, err := productRepo.FetchProducts(ctx)
	if err != nil {
		fmt.Println(err)
	}
    for _, product := range products {
		fmt.Println(product.Title)
	}
}