package ghost

import (
	"log"
    "bitbucket.org/morea/go-st"
)

// Config defines the configuration options for the Ghost client
type Config struct {
	User string
    Password string
    URL string
}

// Client returns a new PagerDuty client
func (c *Config) Client() (*ghost.Client, error) {
	client := ghost.NewClient(c.User, c.Password, c.URL)

	log.Printf("[INFO] Ghost client configured")

	return client, nil
}