package services

import (
	"encoding/xml"

	"github.com/buroz/gon11/src/client"
	"github.com/buroz/gon11/src/constants"
	"github.com/buroz/gon11/src/protos"
)

type CategoryService struct{}

func (s *CategoryService) GetTopLevelCategories(user protos.Auth) *protos.GetTopLevelCategoriesResponse {
	soapClient := &client.Request{}
	request := protos.GetTopLevelCategoriesRequest{}
	request.Auth.AppKey = user.AppKey
	request.Auth.AppSecret = user.AppSecret
	response := soapClient.Call(constants.CategoryService, &request)
	data := &protos.GetTopLevelCategoriesResponse{}
	xml.Unmarshal(response, &data)
	return data
}

func (s *CategoryService) GetSubCategories(user protos.Auth, categoryId string) *protos.GetSubCategoriesResponse {
	soapClient := &client.Request{}
	request := protos.GetSubCategoriesRequest{}
	request.Auth.AppKey = user.AppKey
	request.Auth.AppSecret = user.AppSecret
	request.CategoryId = categoryId
	response := soapClient.Call(constants.CategoryService, &request)
	data := &protos.GetSubCategoriesResponse{}
	xml.Unmarshal(response, &data)
	return data
}

func (s *CategoryService) GetParentCategory(user protos.Auth, categoryId string) *protos.GetParentCategoryResponse {
	soapClient := &client.Request{}
	request := protos.GetParentCategoryRequest{}
	request.Auth.AppKey = user.AppKey
	request.Auth.AppSecret = user.AppSecret
	request.CategoryId = categoryId
	response := soapClient.Call(constants.CategoryService, &request)
	data := &protos.GetParentCategoryResponse{}
	xml.Unmarshal(response, &data)
	return data
}

func (s *CategoryService) GetCategoryAttributes(user protos.Auth, pagingData protos.PagingDataRequest, categoryId string) *protos.GetCategoryAttributesResponse {
	soapClient := &client.Request{}
	request := protos.GetCategoryAttributesRequest{}
	request.Auth.AppKey = user.AppKey
	request.Auth.AppSecret = user.AppSecret

	request.PagingData.CurrentPage = pagingData.CurrentPage - 1
	request.PagingData.PageSize = pagingData.PageSize

	request.CategoryId = categoryId
	response := soapClient.Call(constants.CategoryService, &request)
	data := &protos.GetCategoryAttributesResponse{}
	xml.Unmarshal(response, &data)
	return data
}

func (s *CategoryService) GetCategoryAttributesId(user protos.Auth, categoryId string) *protos.GetCategoryAttributesIdResponse {
	soapClient := &client.Request{}
	request := protos.GetCategoryAttributesIdRequest{}
	request.Auth.AppKey = user.AppKey
	request.Auth.AppSecret = user.AppSecret
	request.CategoryId = categoryId
	response := soapClient.Call(constants.CategoryService, &request)
	data := &protos.GetCategoryAttributesIdResponse{}
	xml.Unmarshal(response, &data)
	return data
}

func (s *CategoryService) GetCategoryAttributeValue(user protos.Auth, pagingData protos.PagingDataRequest, categoryProductAttributeId string) *protos.GetCategoryAttributeValueResponse {
	soapClient := &client.Request{}
	request := protos.GetCategoryAttributeValueRequest{}
	request.Auth.AppKey = user.AppKey
	request.Auth.AppSecret = user.AppSecret

	request.PagingData.CurrentPage = pagingData.CurrentPage - 1
	request.PagingData.PageSize = pagingData.PageSize

	request.CategoryProductAttributeId = categoryProductAttributeId
	response := soapClient.Call(constants.CategoryService, &request)
	data := &protos.GetCategoryAttributeValueResponse{}
	xml.Unmarshal(response, &data)
	return data
}
