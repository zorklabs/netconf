package comware

import "github.com/exsver/netconf/netconf"

func (targetDevice *TargetDevice) GetDataDHCPServer() (*DHCP, error) {
	request := netconf.RPCMessage{
		InnerXML: []byte(`<get><filter type="subtree"><top xmlns="http://www.hp.com/netconf/data:1.0"><DHCP/></top></filter></get>`),
		Xmlns:    []string{netconf.BaseURI},
	}

	data, err := targetDevice.RetrieveData(request)
	if err != nil {
		return nil, err
	}

	return data.Top.DHCP, nil
}
