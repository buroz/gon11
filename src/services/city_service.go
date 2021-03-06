package services

import (
	"encoding/xml"

	"github.com/buroz/gon11/src/client"
	"github.com/buroz/gon11/src/constants"
	"github.com/buroz/gon11/src/protos"
)

type CityService struct{}

func (s *CityService) GetCities() *protos.GetCitiesResponse {
	soapClient := &client.Request{}
	response := soapClient.Call(constants.CityService, &protos.GetCitiesRequest{})
	data := &protos.GetCitiesResponse{}
	xml.Unmarshal(response, &data)
	return data
}

func (s *CityService) GetCity(cityCode string) *protos.GetCityResponse {
	soapClient := &client.Request{}
	req := &protos.GetCityRequest{}
	req.CityCode = cityCode
	response := soapClient.Call(constants.CityService, req)
	data := &protos.GetCityResponse{}
	xml.Unmarshal(response, &data)
	return data
}

func (s *CityService) GetDistrict(cityCode string) *protos.GetDistrictResponse {
	soapClient := &client.Request{}
	req := &protos.GetDistrictRequest{}
	req.CityCode = cityCode
	response := soapClient.Call(constants.CityService, req)
	data := &protos.GetDistrictResponse{}
	xml.Unmarshal(response, &data)
	return data
}

func (s *CityService) GetNeighborhoods(districtId string) *protos.GetNeighborhoodsResponse {
	soapClient := &client.Request{}
	req := &protos.GetNeighborhoodsRequest{}
	req.DistrictId = districtId
	response := soapClient.Call(constants.CityService, req)
	data := &protos.GetNeighborhoodsResponse{}
	xml.Unmarshal(response, &data)
	return data
}
