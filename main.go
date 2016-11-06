package main

import (
	"bitbucket.org/morea/terraform-provider-ghost/provider"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	p := plugin.ServeOpts{
		ProviderFunc: provider.Provider,
	}

	plugin.Serve(&p)
}