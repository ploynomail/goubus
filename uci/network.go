package uci

import (
	"github.com/digineo/go-uci/v2"
)

type NetworkInterface struct {
	Name    string `json:"name"`
	Type    string `json:"type"`  // "bridge" or "ethernet" or "wireless"
	Proto   string `json:"proto"` // "static" or "dhcp"
	Ipaddr  string `json:"ipaddr"`
	Netmask string `json:"netmask"`
	Gateway string `json:"gateway"`
	Dns     string `json:"dns"`
	Macaddr string
	Device  string `json:"ifname"`
}

type Network struct {
	Network map[string]NetworkInterface
	UciTree uci.Tree
}

func NewNetwork(ut uci.Tree) *Network {
	if ut == nil {
		ut = uci.NewTree("/etc/config")
	}
	return &Network{
		Network: make(map[string]NetworkInterface),
		UciTree: ut,
	}
}

func (n *Network) AddInterface(name string, iface NetworkInterface) {
	n.Network[name] = iface
}

func (n *Network) GetInterface(name string) NetworkInterface {
	return n.Network[name]
}

func (n *Network) DelInterface(name string) {
	delete(n.Network, name)
}

func (n *Network) GetInterfaces() map[string]NetworkInterface {
	return n.Network
}

func (n *Network) SetInterfaces(ifaces map[string]NetworkInterface) {
	n.Network = ifaces
}

func (n *Network) GetInterfaceNames() []string {
	names := []string{}
	for name := range n.Network {
		names = append(names, name)
	}
	return names
}

func (n *Network) SaveConfig() error {
	for _, iface := range n.Network {
		err := n.ReplayInterface(iface)
		if err != nil {
			return err
		}
	}

	return n.UciTree.Commit()
}

func (n *Network) ReplayInterface(inf NetworkInterface) error {
	if inf.Type != "bridge" && inf.Device == "" {
		inf.Device = inf.Name
	}
	infMap, err := StructToMap(inf)
	if err != nil {
		return err
	}

	err = n.UciTree.DelSection("network", inf.Name)
	if err != nil {
		return err
	}
	err = n.UciTree.AddSection("network", inf.Name, "interface")
	if err != nil {
		return err
	}
	for k, v := range infMap {
		n.UciTree.SetType("network", inf.Name, k, uci.TypeOption, v)
	}
	return nil
}
