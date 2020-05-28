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
				Status        string `xml:"status"`
				ErrorCode     string `xml:"errorCode"`
				ErrorMessage  string `xml:"errorMessage"`
				ErrorCategory string `xml:"errorCategory"`
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
				Status        string `xml:"status"`
				ErrorCode     string `xml:"errorCode"`
				ErrorMessage  string `xml:"errorMessage"`
				ErrorCategory string `xml:"errorCategory"`
			} `xml:"result"`
			City struct {
				CityName string `xml:"cityName"`
				CityCode string `xml:"cityCode"`
				CityId   string `xml:"cityId"`
			} `xml:"city"`
		} `xml:"GetCityResponse"`
	} `xml:"Body"`
}

// GetDistrict
type GetDistrictRequest struct {
	XMLName  xml.Name `xml:"sch:GetDistrictRequest"`
	CityCode string   `xml:"cityCode"`
}

type GetDistrictResponse struct {
	Root xml.Name `xml:"Envelope"`
	Body struct {
		GetDistrictResponse struct {
			Result struct {
				Status        string `xml:"status"`
				ErrorCode     string `xml:"errorCode"`
				ErrorMessage  string `xml:"errorMessage"`
				ErrorCategory string `xml:"errorCategory"`
			} `xml:"result"`
			Districts struct {
				District []struct {
					DistrictId   string `xml:"id"`
					DistrictName string `xml:"name"`
				} `xml:"district"`
			} `xml:"districts"`
		} `xml:"GetDistrictResponse"`
	} `xml:"Body"`
}

// GetNeighborhoods
type GetNeighborhoodsRequest struct {
	XMLName    xml.Name `xml:"sch:GetNeighborhoodsRequest"`
	DistrictId string   `xml:"districtId"`
}

type GetNeighborhoodsResponse struct {
	Root xml.Name `xml:"Envelope"`
	Body struct {
		GetNeighborhoodsResponse struct {
			Result struct {
				Status        string `xml:"status"`
				ErrorCode     string `xml:"errorCode"`
				ErrorMessage  string `xml:"errorMessage"`
				ErrorCategory string `xml:"errorCategory"`
			} `xml:"result"`
			Neighborhoods struct {
				Neighborhood []struct {
					NeighborhoodId   string `xml:"id"`
					NeighborhoodName string `xml:"name"`
				} `xml:"neighborhood"`
			} `xml:"neighborhoods"`
		} `xml:"GetNeighborhoodsResponse"`
	} `xml:"Body"`
}
