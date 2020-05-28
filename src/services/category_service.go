package services

import (
	"encoding/xml"

	"github.com/buroz/gon11/src/client"
	"github.com/buroz/gon11/src/constants"
	"github.com/buroz/gon11/src/protos"
)

type CategoryService struct{}

func (s *CategoryService) GetTopLevelCategories(user client.Auth) *protos.GetTopLevelCategoriesResponse {
	soapClient := &client.Request{}
	request := protos.GetTopLevelCategoriesRequest{}
	request.Auth.AppKey = user.AppKey
	request.Auth.AppSecret = user.AppSecret
	response := soapClient.Call(constants.CategoryService, &request)
	data := &protos.GetTopLevelCategoriesResponse{}
	xml.Unmarshal(response, &data)
	return data
}
