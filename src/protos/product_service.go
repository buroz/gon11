package protos

import (
	"encoding/xml"
)

// GetProductByProductId
type GetProductByProductIdRequest struct {
	XMLName   xml.Name `xml:"sch:GetProductByProductIdRequest"`
	Auth      Auth     `xml:"auth"`
	ProductId int      `xml:"productId"`
}

type GetProductByProductIdResponse struct {
	Root xml.Name `xml:"Envelope"`
	Body struct {
		GetProductByProductIdResponse struct {
			XMLName xml.Name `xml:"GetProductByProductIdResponse"`
			Result  Result   `xml:"result"`
			Product Product  `xml:"product"`
		}
	} `xml:"Body"`
}

// GetProductBySellerCode
type GetProductBySellerCodeRequest struct {
	XMLName    xml.Name `xml:"sch:GetProductBySellerCodeRequest"`
	Auth       Auth
	SellerCode string `xml:"sellerCode"`
}

type GetProductBySellerCodeResponse struct {
	Root xml.Name `xml:"Envelope"`
	Body struct {
		GetProductBySellerCodeResponse struct {
			XMLName xml.Name `xml:"GetProductBySellerCodeResponse"`
			Result  Result   `xml:"result"`
			Product Product  `xml:"product"`
		}
	} `xml:"Body"`
}

// GetProductList
type GetProductListRequest struct {
	XMLName    xml.Name          `xml:"sch:GetProductListRequest"`
	Auth       Auth              `xml:"auth"`
	PagingData PagingDataRequest `xml:"pagingData"`
}

type GetProductListResponse struct {
	Root xml.Name `xml:"Envelope"`
	Body struct {
		GetProductListResponse struct {
			XMLName    xml.Name           `xml:"GetProductListResponse"`
			Result     Result             `xml:"result"`
			PagingData PagingDataResponse `xml:"pagingData"`
			Product    Product            `xml:"product"`
		}
	} `xml:"Body"`
}

// SaveProduct
type SaveProductRequest struct {
	XMLName xml.Name    `xml:"sch:SaveProductRequest"`
	Auth    Auth        `xml:"auth"`
	Product SaveProduct `xml:"product"`
}

type SaveProductResponse struct {
	Root xml.Name `xml:"Envelope"`
	Body struct {
		SaveProductResponse struct {
			XMLName    xml.Name           `xml:"SaveProductResponse"`
			Result     Result             `xml:"result"`
			PagingData PagingDataResponse `xml:"pagingData"`
			Product    Product            `xml:"product"`
		}
	} `xml:"Body"`
}
