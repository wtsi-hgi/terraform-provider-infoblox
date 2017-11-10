package infoblox

import (
	"log"
	"time"

	"github.com/fanatic/go-infoblox"
)

// Config holds authentication details of Infoblox
type Config struct {
	Host       string
	Password   string
	Username   string
	SSLVerify  bool
	UseCookies bool
}

// Client returns a new client for accessing Infoblox.
func (c *Config) Client() (*infoblox.Client, error) {
	client := infoblox.NewClient(c.Host, c.Username, c.Password, c.SSLVerify, c.UseCookies)
	client.HttpClient.Timeout = 90 * time.Second

	log.Printf("[INFO] Infoblox Client configured for user: %s", client.Username)

	return client, nil
}
