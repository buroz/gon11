package n11

import (
	"github.com/buroz/gon11/src/services"
)

type Services struct {
	CityService services.CityService
}

type Client struct {
	user struct {
		AppKey    string
		AppSecret string
	}
	Services Services
}

func (c *Client) Create(appKey string, appSecret string) struct {
	AppKey    string
	AppSecret string
} {
	c.user.AppKey = appKey
	c.user.AppSecret = appSecret
	c.Services = Services{}
	return c.user
}
