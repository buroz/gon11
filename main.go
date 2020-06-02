package main

import (
	"fmt"
	"log"
	"os"

	"github.com/buroz/gon11/src/n11"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		os.Exit(1)
	}

	var (
		appKey    = os.Getenv("APP_KEY")
		appSecret = os.Getenv("APP_SECRET")
	)

	var client = n11.Client{}

	user := client.Create(appKey, appSecret)

	// pagination := protos.PagingDataRequest{}

	attr := client.Services.ProductService.GetProductByProductId(user, 369857815)

	fmt.Println(attr.Body.GetProductByProductIdResponse.Product.Title)
	fmt.Println(attr.Body.GetProductByProductIdResponse.Product.SubTitle)
	fmt.Println(attr.Body.GetProductByProductIdResponse.Product.ProductSellerCode)
	fmt.Println(attr.Body.GetProductByProductIdResponse.Product.Price)
	fmt.Println(attr.Body.GetProductByProductIdResponse.Product.CurrentAmount)

}
