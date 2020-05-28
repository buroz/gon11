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
