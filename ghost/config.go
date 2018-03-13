package ghost

import (
	"fmt"
	"log"
	"net/url"

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
	if c.Password == "" || c.User == "" || c.URL == "" {
		return nil, fmt.Errorf(`At least 1 ghost parameter is empty: Username: %s,
			 Password, URL: %s`, c.User, c.URL)
	}

	if _, err := url.ParseRequestURI(c.URL); err != nil {
		return nil, fmt.Errorf("Invalid endpoint URL")
	}

	client := ghost.NewClient(c.URL, c.User, c.Password)

	log.Printf("[INFO] Ghost client configured: %s %s", c.User, c.URL)

	return client, nil
}
