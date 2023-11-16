package product

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/spodolaks/shopifyClient/models"
	"github.com/spodolaks/shopifyClient/network"
	"github.com/spodolaks/shopifyClient/repositories"
)
type productRepo struct {
	Conn *http.Client
} 
func NewProductRepo(Conn *http.Client) repositories.ProductRepo {
	return &productRepo{
		Conn: Conn,
	}
}

func (pr *productRepo) FetchProducts(ctx context.Context) ([]models.Product, error) {
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
  
	token := os.Getenv("TOKEN")
	network.Initialize();
	url := "https://spodolaks.myshopify.com/admin/api/2023-10/products.json"
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("content-type", "application/json")
	req.Header.Add("X-Shopify-Access-Token", token)
	resp, err := pr.Conn.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

    // Read the response body
    responseBody, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }
	var products models.Response
	err = json.Unmarshal([]byte(responseBody), &products)
    if err != nil {
        return nil, err
    }
	return products.Products, nil
}

