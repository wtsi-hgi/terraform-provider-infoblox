package infoblox

import (
	"log"
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
func (c *Config) Client() (*fibClient, error) {
	client := NewFibClient(c.Host, c.Username, c.Password, c.SSLVerify, c.UseCookies)

	log.Printf("[INFO] Infoblox Client configured for user: %s", client.Username)

	return client, nil
}
