package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/philips-software/terraform-provider-hsdp/hsdp"
)

var commit = "deadbeef"
var release = "v0.0.0"
var date = "0000-00-00"
var buildSource = "unknown"
var buildVersion = release + "-" + commit + "." + date + "." + buildSource

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return hsdp.Provider(buildVersion)
		},
	})
}
