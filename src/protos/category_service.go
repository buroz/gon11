package protos

import (
	"encoding/xml"

	"github.com/buroz/gon11/src/client"
)

// GetTopLevelCategories
type GetTopLevelCategoriesRequest struct {
	XMLName xml.Name `xml:"sch:GetTopLevelCategoriesRequest"`
	Auth    client.Auth
}

type GetTopLevelCategoriesResponse struct {
	Root xml.Name `xml:"Envelope"`
	Body struct {
		GetTopLevelCategoriesResponse struct {
			XMLName xml.Name `xml:"GetTopLevelCategoriesResponse"`
			Result  struct {
				Status        string `xml:"status"`
				ErrorCode     string `xml:"errorCode"`
				ErrorMessage  string `xml:"errorMessage"`
				ErrorCategory string `xml:"errorCategory"`
			} `xml:"result"`
			CategoryList struct {
				Category []struct {
					CategoryName string `xml:"name"`
					CategoryId   string `xml:"id"`
				} `xml:"category"`
			} `xml:"categoryList"`
		}
	} `xml:"Body"`
}

// GetSubCategories
type GetSubCategoriesRequest struct {
	XMLName    xml.Name `xml:"sch:GetSubCategoriesRequest"`
	Auth       client.Auth
	CategoryId string `xml:"categoryId"`
}

type GetSubCategoriesResponse struct {
	Root xml.Name `xml:"Envelope"`
	Body struct {
		GetSubCategoriesResponse struct {
			XMLName xml.Name `xml:"GetSubCategoriesResponse"`
			Result  struct {
				Status        string `xml:"status"`
				ErrorCode     string `xml:"errorCode"`
				ErrorMessage  string `xml:"errorMessage"`
				ErrorCategory string `xml:"errorCategory"`
			} `xml:"result"`
			Category struct {
				CategoryName string `xml:"name"`
				CategoryId   string `xml:"id"`

				SubCategories struct {
					SubCategory []struct {
						SubCategoryName string `xml:"name"`
						SubCategoryId   string `xml:"id"`
					} `xml:"category"`
				} `xml:"subCategoryList"`
			} `xml:"category"`
		}
	} `xml:"Body"`
}
