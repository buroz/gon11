package n11

import (
	"github.com/buroz/gon11/src/protos"
	"github.com/buroz/gon11/src/services"
)

type Services struct {
	CityService     services.CityService
	CategoryService services.CategoryService
	ProductService  services.ProductService
}

type Client struct {
	user     protos.Auth
	Services Services
}

func (c *Client) Create(appKey string, appSecret string) protos.Auth {
	c.user.AppKey = appKey
	c.user.AppSecret = appSecret
	c.Services = Services{
		CityService:     services.CityService{},
		CategoryService: services.CategoryService{},
		ProductService:  services.ProductService{},
	}
	return c.user
}
