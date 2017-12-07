package main

import (
	"github.com/BSick7/terraform-provider-okta/okta"
	"github.com/hashicorp/terraform/plugin"
)

var Version string

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: okta.Provider,
	})
}
