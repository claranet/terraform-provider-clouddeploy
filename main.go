package main

import (
	provider "bitbucket.org/morea/terraform-provider-ghost/ghost"

	"github.com/hashicorp/terraform/plugin"
)

func main() {
	p := plugin.ServeOpts{
		ProviderFunc: provider.Provider,
	}

	plugin.Serve(&p)
}
