package n11

import (
	"github.com/buroz/gon11/src/client"
	"github.com/buroz/gon11/src/services"
)

type Services struct {
	CityService     services.CityService
	CategoryService services.CategoryService
}

type Client struct {
	user     client.Auth
	Services Services
}

func (c *Client) Create(appKey string, appSecret string) client.Auth {
	c.user.AppKey = appKey
	c.user.AppSecret = appSecret
	c.Services = Services{
		CityService:     services.CityService{},
		CategoryService: services.CategoryService{},
	}
	return c.user
}
