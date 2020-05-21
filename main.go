package main

import (
	"fmt"

	"github.com/buroz/gon11/src/n11"
)

func main() {
	var client = &n11.Client{}

	user := client.Create("asd", "asd")
	fmt.Println(user)

	cities := client.Services.CityService.GetCity("07")
	fmt.Println(cities.Body.GetCityResponse.Result.Status)
}
