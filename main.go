package main

import (
	"fmt"
	"log"
	"network"
	"os"

	"github.com/joho/godotenv"
	"github.com/spodolaks/shopifyGo/network"
)

func main() {
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
  
	token := os.Getenv("TOKEN")
	network.Initialize();
	fmt.Println(token)
}