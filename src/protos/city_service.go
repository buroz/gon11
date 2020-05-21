package protos

import "encoding/xml"

// GetCities
type GetCitiesRequest struct {
	XMLName xml.Name `xml:"sch:GetCitiesRequest"`
}

type GetCitiesResponse struct {
	Root xml.Name `xml:"Envelope"`
	Body struct {
		GetCitiesResponse struct {
			XMLName xml.Name `xml:"GetCitiesResponse"`
			Result  struct {
				Status string `xml:"status"`
			} `xml:"result"`
			Cities struct {
				City []struct {
					CityName string `xml:"cityName"`
					CityCode string `xml:"cityCode"`
					CityId   string `xml:"cityId"`
				} `xml:"city"`
			} `xml:"cities"`
		}
	} `xml:"Body"`
}

// GetCity
type GetCityRequest struct {
	XMLName  xml.Name `xml:"sch:GetCityRequest"`
	CityCode string   `xml:"cityCode"`
}

type GetCityResponse struct {
	Root xml.Name `xml:"Envelope"`
	Body struct {
		GetCityResponse struct {
			Result struct {
				Status string `xml:"status"`
			} `xml:"result"`
			City struct {
				CityName string `xml:"cityName"`
				CityCode string `xml:"cityCode"`
				CityId   string `xml:"cityId"`
			} `xml:"city"`
		} `xml:"GetCityResponse"`
	} `xml:"Body"`
}
