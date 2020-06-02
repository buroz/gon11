package protos

import (
	"encoding/xml"
)

// GetTopLevelCategories
type GetTopLevelCategoriesRequest struct {
	XMLName xml.Name `xml:"sch:GetTopLevelCategoriesRequest"`
	Auth    Auth     `xml:"auth"`
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
	Auth       Auth     `xml:"auth"`
	CategoryId string   `xml:"categoryId"`
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

				SubCategoryList struct {
					SubCategory []struct {
						SubCategoryName string `xml:"name"`
						SubCategoryId   string `xml:"id"`
					} `xml:"subCategory"`
				} `xml:"subCategoryList"`
			} `xml:"category"`
		}
	} `xml:"Body"`
}

// GetParentCategory
type GetParentCategoryRequest struct {
	XMLName    xml.Name `xml:"sch:GetParentCategoryRequest"`
	Auth       Auth     `xml:"auth"`
	CategoryId string   `xml:"categoryId"`
}

type GetParentCategoryResponse struct {
	Root xml.Name `xml:"Envelope"`
	Body struct {
		GetParentCategoryResponse struct {
			XMLName xml.Name `xml:"GetParentCategoryResponse"`
			Result  struct {
				Status        string `xml:"status"`
				ErrorCode     string `xml:"errorCode"`
				ErrorMessage  string `xml:"errorMessage"`
				ErrorCategory string `xml:"errorCategory"`
			} `xml:"result"`
			Category struct {
				CategoryName string `xml:"name"`
				CategoryId   string `xml:"id"`

				ParentCategory struct {
					ParentCategoryName string `xml:"name"`
					ParentCategoryId   string `xml:"id"`
				} `xml:"parentCategory"`
			} `xml:"category"`
		}
	} `xml:"Body"`
}

// GetCategoryAttributes
type GetCategoryAttributesRequest struct {
	XMLName    xml.Name          `xml:"sch:GetCategoryAttributesRequest"`
	Auth       Auth              `xml:"auth"`
	PagingData PagingDataRequest `xml:"pagingData"`
	CategoryId string            `xml:"categoryId"`
}

type GetCategoryAttributesResponse struct {
	Root xml.Name `xml:"Envelope"`
	Body struct {
		GetCategoryAttributesResponse struct {
			XMLName xml.Name `xml:"GetCategoryAttributesResponse"`
			Result  struct {
				Status        string `xml:"status"`
				ErrorCode     string `xml:"errorCode"`
				ErrorMessage  string `xml:"errorMessage"`
				ErrorCategory string `xml:"errorCategory"`
			} `xml:"result"`
			PagingData PagingDataResponse `xml:"pagingData"`
			Category   struct {
				CategoryName string `xml:"name"`
				CategoryId   string `xml:"id"`

				ParentCategory struct {
					ParentCategoryName string `xml:"name"`
					ParentCategoryId   string `xml:"id"`
				} `xml:"parentCategory"`

				AttributeList struct {
					Attribute []struct {
						Id             string  `xml:"id"`
						Name           string  `xml:"name"`
						Mandatory      bool    `xml:"mandatory"`
						MultipleSelect bool    `xml:"multipleSelect"`
						Priority       float64 `xml:"priority"`
						ValueList      struct {
							Value []struct {
								Id   string `xml:"id"`
								Name string `xml:"name"`
							} `xml:"value"`
						} `xml:"valueList"`
					} `xml:"attribute"`
				} `xml:"attributeList"`
			} `xml:"category"`
		}
	} `xml:"Body"`
}

// GetCategoryAttributesId
type GetCategoryAttributesIdRequest struct {
	XMLName    xml.Name `xml:"sch:GetCategoryAttributesIdRequest"`
	Auth       Auth     `xml:"auth"`
	CategoryId string   `xml:"categoryId"`
}

type GetCategoryAttributesIdResponse struct {
	Root xml.Name `xml:"Envelope"`
	Body struct {
		GetCategoryAttributesIdResponse struct {
			XMLName xml.Name `xml:"GetCategoryAttributesIdResponse"`
			Result  struct {
				Status        string `xml:"status"`
				ErrorCode     string `xml:"errorCode"`
				ErrorMessage  string `xml:"errorMessage"`
				ErrorCategory string `xml:"errorCategory"`
			} `xml:"result"`
			CategoryProductAttributeList struct {
				CategoryProductAttribute []struct {
					Id             string `xml:"id"`
					Name           string `xml:"name"`
					Mandatory      bool   `xml:"mandatory"`
					MultipleSelect bool   `xml:"multipleSelect"`
				} `xml:"categoryProductAttribute"`
			} `xml:"categoryProductAttributeList"`
		}
	} `xml:"Body"`
}

// GetCategoryAttributeValue
type GetCategoryAttributeValueRequest struct {
	XMLName                    xml.Name          `xml:"sch:GetCategoryAttributeValueRequest"`
	Auth                       Auth              `xml:"auth"`
	PagingData                 PagingDataRequest `xml:"pagingData"`
	CategoryProductAttributeId string            `xml:"categoryProductAttributeId"`
}

type GetCategoryAttributeValueResponse struct {
	Root xml.Name `xml:"Envelope"`
	Body struct {
		GetCategoryAttributeValueResponse struct {
			XMLName xml.Name `xml:"GetCategoryAttributeValueResponse"`
			Result  struct {
				Status        string `xml:"status"`
				ErrorCode     string `xml:"errorCode"`
				ErrorMessage  string `xml:"errorMessage"`
				ErrorCategory string `xml:"errorCategory"`
			} `xml:"result"`
			PagingData                        PagingDataResponse
			CategoryProductAttributeValueList struct {
				CategoryProductAttributeValue []struct {
					Id             string `xml:"id"`
					Name           string `xml:"name"`
					Mandatory      bool   `xml:"mandatory"`
					MultipleSelect bool   `xml:"multipleSelect"`
				} `xml:"categoryProductAttributeValue"`
			} `xml:"categoryProductAttributeValueList"`
		}
	} `xml:"Body"`
}
