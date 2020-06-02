package services

import (
	"encoding/xml"

	"github.com/buroz/gon11/src/client"
	"github.com/buroz/gon11/src/constants"
	"github.com/buroz/gon11/src/protos"
)

type ProductService struct{}

func (s *ProductService) GetProductByProductId(user protos.Auth, productId int) *protos.GetProductByProductIdResponse {
	soapClient := &client.Request{}
	request := protos.GetProductByProductIdRequest{}
	request.Auth.AppKey = user.AppKey
	request.Auth.AppSecret = user.AppSecret
	request.ProductId = productId
	response := soapClient.Call(constants.ProductService, &request)
	data := &protos.GetProductByProductIdResponse{}
	xml.Unmarshal(response, &data)
	return data
}

func (s *ProductService) GetProductBySellerCode(user protos.Auth, sellerCode string) *protos.GetProductBySellerCodeResponse {
	soapClient := &client.Request{}
	request := protos.GetProductBySellerCodeRequest{}
	request.Auth.AppKey = user.AppKey
	request.Auth.AppSecret = user.AppSecret
	request.SellerCode = sellerCode
	response := soapClient.Call(constants.ProductService, &request)
	data := &protos.GetProductBySellerCodeResponse{}
	xml.Unmarshal(response, &data)
	return data
}

func (s *ProductService) GetProductList(user protos.Auth, pagingData protos.PagingDataRequest) *protos.GetProductListResponse {
	soapClient := &client.Request{}
	request := protos.GetProductListRequest{}
	request.Auth.AppKey = user.AppKey
	request.Auth.AppSecret = user.AppSecret
	request.PagingData.CurrentPage = pagingData.CurrentPage - 1
	request.PagingData.PageSize = pagingData.PageSize
	response := soapClient.Call(constants.ProductService, &request)
	data := &protos.GetProductListResponse{}
	xml.Unmarshal(response, &data)
	return data
}

func (s *ProductService) SaveProduct(user protos.Auth, product protos.SaveProduct) *protos.SaveProductResponse {
	soapClient := &client.Request{}
	request := protos.SaveProductRequest{}
	request.Auth.AppKey = user.AppKey
	request.Auth.AppSecret = user.AppSecret
	request.Product = product
	response := soapClient.Call(constants.ProductService, &request)
	data := &protos.SaveProductResponse{}
	xml.Unmarshal(response, &data)
	return data
}
