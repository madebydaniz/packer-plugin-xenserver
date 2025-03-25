package main

import (
	"fmt"
	"os"

	"github.com/madebydaniz/packer-plugin-xenserver/builder/xenserver/iso"
	"github.com/madebydaniz/packer-plugin-xenserver/builder/xenserver/xva"
	"github.com/madebydaniz/packer-plugin-xenserver/version"

	"github.com/hashicorp/packer-plugin-sdk/plugin"
)

func main() {
	pps := plugin.NewSet()
	pps.RegisterBuilder("iso", new(iso.Builder))
	pps.RegisterBuilder("xva", new(xva.Builder))
	pps.SetVersion(version.PluginVersion)
	err := pps.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
