package protos

import "encoding/xml"

type Auth struct {
	XMLName   xml.Name `xml:"auth"`
	AppKey    string   `xml:"appKey"`
	AppSecret string   `xml:"appSecret"`
}

type PagingDataRequest struct {
	XMLName     xml.Name `xml:"pagingData"`
	CurrentPage int      `xml:"currentPage"`
	PageSize    int      `xml:"pageSize"`
}

type PagingDataResponse struct {
	XMLName     xml.Name `xml:"pagingData"`
	CurrentPage int      `xml:"currentPage"`
	PageSize    int      `xml:"pageSize"`
	TotalCount  int      `xml:"totalCount"`
	PageCount   int      `xml:"pageCount"`
}

type Result struct {
	Status        string `xml:"status"`
	ErrorCode     string `xml:"errorCode"`
	ErrorMessage  string `xml:"errorMessage"`
	ErrorCategory string `xml:"errorCategory"`
}

type Product struct {
	CurrentAmount     float32 `xml:"currentAmount"`
	Id                int     `xml:"id"`
	Price             float32 `xml:"price"`
	ProductSellerCode string  `xml:"productSellerCode"`
	ApprovalStatus    int     `xml:"approvalStatus"`
	SubTitle          string  `xml:"subTitle"`
	Title             string  `xml:"title"`
	Description       string  `xml:"Description"`
	Domestic          bool    `xml:"domestic"`
	PreparingDay      int     `xml:"preparingDay"`
	ProductCondution  int     `xml:"productCondution"`
	ProductionDate    string  `xml:"productionDate"`
	SaleEndDate       string  `xml:"saleEndDate"`
	SaleStartDate     string  `xml:"saleStartDate"`
	ExpirationDate    string  `xml:"expirationDate"`
	ShipmentTemplate  string  `xml:"shipmentTemplate"`
	// SpecialProductInfoList struct{} `xml:"specialProductInfoList"`

	Images struct {
		Image []struct {
			Order int    `xml:"order"`
			Url   string `xml:"url"`
		} `xml:"image"`
	} `xml:"images"`

	Category struct {
		FullName string `xml:"fullName"`
		Id       int    `xml:"id"`
		Name     int    `xml:"Name"`
	} `xml:"category"`

	Attributes struct {
		Attribute []struct {
			Id    string `xml:"id"`
			Name  string `xml:"name"`
			Value string `xml:"value"`
		} `xml:"attribute"`
	} `xml:"attributes"`

	StockItems struct {
		StockItem []struct {
			CurrentAmount   float32 `xml:"currentAmount"`
			DisplayPrice    float32 `xml:"displayPrice"`
			GTIN            string  `xml:"gtin"`
			OEM             string  `xml:"oem"`
			MPN             string  `xml:"mpn"`
			OptionPrice     float32 `xml:"optionPrice"`
			SellerStockCode string  `xml:"sellerStockCode"`
			Id              int     `xml:"id"`
			Quantity        int     `xml:"quantity"`
			Version         int     `xml:"version"`
			Attributes      struct {
				Attribute []struct {
					Id    string `xml:"id"`
					Name  string `xml:"name"`
					Value string `xml:"value"`
				} `xml:"attribute"`
			} `xml:"attributes"`
		} `xml:"stockItem"`
	} `xml:"stockItems"`
}

type SaveProduct struct {
	ProductSellerCode string  `xml:"productSellerCode"`
	Title             string  `xml:"title"`
	SubTitle          string  `xml:"subTitle"`
	Description       string  `xml:"Description"`
	Price             float32 `xml:"price"`
	CurrentType       string  `xml:"currentType"`

	SaleEndDate    string `xml:"saleEndDate"`
	SaleStartDate  string `xml:"saleStartDate"`
	ProductionDate string `xml:"productionDate"`
	ExpirationDate string `xml:"expirationDate"`

	PreparingDay     int    `xml:"preparingDay"`
	ShipmentTemplate string `xml:"shipmentTemplate"`

	GroupAttribute string `xml:"groupAttribute"`
	GroupItemCode  string `xml:"groupItemCode"`
	ItemName       string `xml:"itemName"`

	StockItems struct {
		StockItem []struct {
			Quantity        int     `xml:"quantity"`
			SellerStockCode string  `xml:"sellerStockCode"`
			OptionPrice     float32 `xml:"optionPrice"`
			// Bundle     ??? `xml:"bundle"`
			// SpecialProductInfoList struct{} `xml:"specialProductInfoList"`

			GTIN string `xml:"gtin"`
			OEM  string `xml:"oem"`
			MPN  string `xml:"mpn"`

			Attributes struct {
				Attribute []struct {
					Name  string `xml:"name"`
					Value string `xml:"value"`
				} `xml:"attribute"`
			} `xml:"attributes"`
		} `xml:"stockItem"`
	} `xml:"stockItems"`

	Images struct {
		Image []struct {
			Order int    `xml:"order"`
			Url   string `xml:"url"`
		} `xml:"image"`
	} `xml:"images"`

	Category struct {
		Id int `xml:"id"`
	} `xml:"category"`

	Attributes struct {
		Attribute []struct {
			Id    string `xml:"id"`
			Name  string `xml:"name"`
			Value string `xml:"value"`
		} `xml:"attribute"`
	} `xml:"attributes"`
}
