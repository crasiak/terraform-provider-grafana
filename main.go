package main

import (
	"github.com/crasiak/terraform-provider-grafana/grafana"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: grafana.Provider})
}
