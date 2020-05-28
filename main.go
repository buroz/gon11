package main

import (
	"fmt"
	"log"
	"os"

	"github.com/buroz/gon11/src/n11"
	"github.com/buroz/gon11/src/protos"
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

	pagination := protos.PagingDataRequest{}

	attr := client.Services.CategoryService.GetCategoryAttributeValue(user, "354079908", pagination)

	fmt.Println(attr.Body.GetCategoryAttributeValueResponse.PagingData.CurrentPage)
	fmt.Println(attr.Body.GetCategoryAttributeValueResponse.PagingData.PageCount)
	fmt.Println(attr.Body.GetCategoryAttributeValueResponse.PagingData.PageSize)
	fmt.Println(attr.Body.GetCategoryAttributeValueResponse.PagingData.TotalCount)

	for _, d := range attr.Body.GetCategoryAttributeValueResponse.CategoryProductAttributeValueList.CategoryProductAttributeValue {
		fmt.Println(d.Id)
		fmt.Println(d.Mandatory)
		fmt.Println(d.MultipleSelect)
		fmt.Println(d.Name)
	}

}
