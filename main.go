package main

import (
	"fmt"
	"log"
	"network"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
  
	token := os.Getenv("TOKEN")
	network.Init();
	fmt.Println(token)
}