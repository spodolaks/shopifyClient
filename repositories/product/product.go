package product

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
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
	shopUrl := os.Getenv("WEBSHOP")
	network.Initialize();
	url := fmt.Sprintf("%s/admin/api/2023-10/products.json", shopUrl)
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
	var products []models.Product
	err = json.Unmarshal([]byte(responseBody), &products)
    if err != nil {
        return nil, err
    }
	return products, nil
}

func (pr *productRepo) UpdateProduct(ctx context.Context, product models.Product) (error) {
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
  
	token := os.Getenv("TOKEN")
	shopUrl := os.Getenv("WEBSHOP")
	network.Initialize();
	if (product.ID == 0) {
		return fmt.Errorf("Product id required but was not provided")
	}
	url := fmt.Sprintf("%s/admin/api/2023-10/products/%d.json", shopUrl,product.ID)
	payloadBytes, err := json.Marshal(product)
	if err != nil {
		return err
	}
	payloadBuffer := bytes.NewReader(payloadBytes)
	req, _ := http.NewRequest(http.MethodPut, url, payloadBuffer)
	req.Header.Add("content-type", "application/json")
	req.Header.Add("X-Shopify-Access-Token", token)
	resp, err := pr.Conn.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)
    // Read the response body
    responseBody, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return err
    }
	var products []models.Product
	err = json.Unmarshal([]byte(responseBody), &products)
    if err != nil {
        return err
    }
	return nil
}

