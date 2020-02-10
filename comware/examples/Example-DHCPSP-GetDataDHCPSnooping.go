package main

import (
	"log"

	"github.com/davecgh/go-spew/spew"

	"github.com/exsver/netconf/comware"
	"github.com/exsver/netconf/netconf"
)

func main() {
	netconf.LogLevel.Verbose()
	sw, err := comware.NewTargetDevice("10.10.10.10", "netconf", "netconf")
	if err != nil {
		log.Fatalf("%s", err)
	}
	dhcpSnooping, err := sw.GetDataDHCPSP()
	if err != nil {
		log.Fatalf("%s", err)
	}
	spew.Dump(dhcpSnooping)
}