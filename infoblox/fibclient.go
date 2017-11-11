package infoblox

import (
	"github.com/fanatic/go-infoblox"
)

type fibClient struct {
	*infoblox.Client 
}

type fibResource struct {
	*infoblox.Resource
}

func (f *fibClient) Network() *fibResource {
	return NewFibResource(f.Client.Network())
}

func NewFibResource(resource *infoblox.Resource) *fibResource {
	fibr := &fibResource{}
	fibr.Resource = resource
	return fibr
}

func NewFibClient(host, username, password string, sslVerify, useCookies bool) *fibClient {
	client := infoblox.NewClient(host, username, password, sslVerify, useCookies)
	fibc := &fibClient{}
	fibc.Client = client
	return fibc
}

