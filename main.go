package main

import (
	provider "terraform-provider-ghost/ghost"

	"github.com/hashicorp/terraform/plugin"
)

func main() {
	p := plugin.ServeOpts{
		ProviderFunc: provider.Provider,
	}

	plugin.Serve(&p)
}
