package ghost

import (
	"log"

	"bitbucket.org/morea/go-st"
)

// Config defines the configuration options for the Ghost client
type Config struct {
	User     string
	Password string
	URL      string
}

// Client returns a new Ghost client
func (c *Config) Client() (*ghost.Client, error) {
	client := ghost.NewClient(c.URL, c.User, c.Password)

	log.Printf("[INFO] Ghost client configured: %s %s %s", c.User, c.Password, c.URL)

	return client, nil
}
