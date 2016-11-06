package ghost

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider represents a resource provider in Terraform
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"user": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("GHOST_USER", nil),
			},
            "password": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("GHOST_PASSWORD", nil),
			},
            "server_url": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("GHOST_SERVER_URL", nil),
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"ghost_app":                resourceGhostApp(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(data *schema.ResourceData) (interface{}, error) {
	config := Config{User: data.Get("user").(string), Password: data.Get("password").(string), URL: data.Get("server_url").(string)}
	log.Println("[INFO] Initializing Ghost client")
	return config.Client()
}