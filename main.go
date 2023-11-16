package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/spodolaks/shopifyClient/models"
	"github.com/spodolaks/shopifyClient/network"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
  
	token := os.Getenv("TOKEN")
	network.Initialize();
	client := network.Client
	url := "https://spodolaks.myshopify.com/admin/api/2023-10/products.json"
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("content-type", "application/json")
	req.Header.Add("X-Shopify-Access-Token", token)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

    // Read the response body
    responseBody, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }
	products := []models.Product
	err := json.Unmarshal([]byte(jsonResponse), &products)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(string(responseBody))
}