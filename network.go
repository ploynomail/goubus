package goubus

import (
	"net"
)

type UbusNetwork struct {
	NetworkInterface []NetworkInterface  `json:"network"`
	NetworkDevice    []UbusNetworkDevice `json:"network.device"`
}

func (u *Ubus) NetworkStatus(id int, infName string) UbusNetwork {
	ubusNetwork := UbusNetwork{}
	networkInterface, _ := u.NetworkInterfaceStatus(id, infName)
	networkDevice, _ := u.NetworkDeviceStatus(id, networkInterface.L3Device)
	ubusNetwork.NetworkInterface = append(ubusNetwork.NetworkInterface, networkInterface)
	ubusNetwork.NetworkDevice = append(ubusNetwork.NetworkDevice, networkDevice)
	return ubusNetwork
}

func (u *Ubus) NetworkDevices() ([]string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	var interfaceNames []string
	for _, i := range interfaces {
		interfaceNames = append(interfaceNames, i.Name)
	}
	return interfaceNames, nil
}
