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

	categories := client.Services.CategoryService.GetTopLevelCategories(user)

	fmt.Println(categories.Body.GetTopLevelCategoriesResponse.Result.Status)
	fmt.Println(categories.Body.GetTopLevelCategoriesResponse.Result.ErrorMessage)

	for _, d := range categories.Body.GetTopLevelCategoriesResponse.CategoryList.Category {
		fmt.Println(d.CategoryId, d.CategoryName)
	}
}
