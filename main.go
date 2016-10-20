package main

import (
	"github.com/TheWeatherCompany/terraform-provider-swift/swift"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: swift.Provider,
	})
}
