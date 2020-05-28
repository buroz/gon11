### N11.com API Parser

#### (In Progress)

```console
go get github.com/buroz/gon11
```

```go
package main

import (
	"fmt"

	"github.com/buroz/gon11/src/n11"
)

func main() {
	var client = n11.Client{}

	user := client.Create("*appKey*", "*appSecret*")

	cities := client.Services.CityService.GetCity("07")
	fmt.Println(cities.Body.GetCityResponse.Result.Status)


	categories := client.Services.CategoryService.GetTopLevelCategories(user)

	if categories.Body.GetTopLevelCategoriesResponse.Result != "success" {
		panic(categories.Body.GetTopLevelCategoriesResponse.Result.ErrorMessage)
	}

	for _, d := range categories.Body.GetTopLevelCategoriesResponse.CategoryList.Category {
		fmt.Println(d.CategoryId, d.CategoryName)
	}
}
```

---

### Service List

- [x] City Service
- [ ] Category Service (Current)
- [ ] Product Service
- [ ] Product Selling Service
- [ ] Product Stock Service
- [ ] Order Service
- [ ] Shipment Company Service
- [ ] Shipment Service
- [ ] Settlement Service
- [ ] Ticket Service
