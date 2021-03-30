package main

import (
	"log"

	"github.com/hashicorp/packer-plugin-sdk/plugin"
	"github.com/hashicorp/packer-plugin-sdk/version"
)

var Version = "0.0.1"

func main() {
	pps := plugin.NewSet()
	pps.SetVersion(version.InitializePluginVersion(Version, ""))
	pps.RegisterBuilder(plugin.DEFAULT_NAME, new(Builder))

	err := pps.Run()
	if err != nil {
		log.Fatalf("%v", err)
	}
}
