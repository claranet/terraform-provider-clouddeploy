package ghost

import (
	"bitbucket.org/morea/go-st"
	"log"
)

// Config defines the configuration options for the Ghost client
type Config struct {
	User     string
	Password string
	URL      string
}

// Client returns a new Ghost client
func (c *Config) Client() (*ghost.Client, error) {
	client := ghost.NewClient(c.User, c.Password, c.URL)

	log.Printf("[INFO] Ghost client configured")

	return client, nil
}
