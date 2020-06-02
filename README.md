### Hardly typed n11.com API Parser

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

##### City Service

- [x] GetCities
- [x] GetCity
- [x] GetDistrict
- [x] GetNeighborhoods

##### Category Service

- [x] GetCategoryAttributes
- [x] GetCategoryAttributesId
- [x] GetCategoryAttributeValue
- [x] GetParentCategory
- [x] GetSubCategories
- [x] GetTopLevelCategories

##### Product Service

- [x] GetProductByProductId
- [x] GetProductBySellerCode
- [x] GetProductList
- [x] Save Product
- [ ] SearchProducts
- [ ] DeleteProductById
- [ ] DeleteProductBySellerCode
- [ ] UpdateDiscountValueByProductId
- [ ] UpdateDiscountValueBySellerCode
- [ ] UpdateProductPriceById
- [ ] UpdateProductPriceBySellerCode
- [ ] UpdateProductBasic
- [ ] GetProductQuestionList
- [ ] GetProductQuestionDetail
- [ ] SaveProductAnswer
- [ ] ProductAllStatusCountsRequest

##### Product Selling Service

- [ ] ProductSellingService
- [ ] StartSellingProductByProductId
- [ ] StartSellingProductBySellerCode
- [ ] StopSellingProductByProductId
- [ ] StopSellingProductBySellerCode

##### Product Stock Service

- [ ] GetProductStockByProductId
- [ ] GetProductStockBySellerCode
- [ ] DeleteAndUpdateStockByStockAttributes
- [ ] UpdateStockByStockId
- [ ] UpdateStockByStockSellerCode
- [ ] IncreaseStockByStockAttributes
- [ ] IncreaseStockByStockId
- [ ] IncreaseStockByStockSellerCode

##### Order Service

- [ ] DetailedOrderList
- [ ] OrderList
- [ ] OrderDetail
- [ ] OrderItemAccept
- [ ] OrderItemReject
- [ ] MakeOrderItemShipment

##### Shipment Company Service

- [ ] GetShipmentCompanies

##### Shipment Service

- [ ] GetShipmentTemplate
- [ ] CreateOrUpdateShipmentTemplate
- [ ] GetShipmentTemplateList

##### Settlement Service

- [ ] GetSettlementList
- [ ] GetSettlementDetail

##### Ticket Service

- [ ] TicketListingAssignedToSeller
- [ ] TicketListingBelongsToSeller
- [ ] TicketAnswer
- [ ] TicketCreate
