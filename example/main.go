package main

import (
	"log"

	"github.com/ploynomail/goubus"
)

// This is an example acl of how to use the goubus package,
// you can use the root user to access all the ubus functions,
// and you can increase the access control list to allow other users to access the ubus functions.
// /usr/share/rpcd/acl.d/root.json
// {
//         "root": {
//                 "description": "Super user access role",
//                 "read": {
//                         "ubus": {
//                                 "*": [ "*" ]
//                         },
//                         "uci": [ "*" ],
//                         "file": {
//                                 "*": ["*"]
//                         }
//                 },
//                 "write": {
//                         "ubus": {
//                                 "*": [ "*" ]
//                         },
//                         "uci": [ "*" ],
//                         "file": {
//                                 "*": ["*"]
//                         },
//                         "cgi-io": ["*"]
//                 }
//         }
// }

func main() {
	ubus := goubus.Ubus{
		Username: "turingroot",
		Password: "123456",
		URL:      "http://192.168.23.197/ubus",
		UciTree:  goubus.NewUciTree(),
	}
	_, err := ubus.AuthLogin()
	if err != nil {
		log.Fatal(err)
	}
	if err := ubus.LoginCheck(); err != nil {
		log.Fatal(err)
	}

	// l, err := ubus.LogRead(5, 10, false, false)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("%+v\n", l)
	// var serverRequest = goubus.ServiceListRequest{
	// 	Name:    "uhttpd",
	// 	Verbose: false,
	// }
	// sl, err := ubus.GetServceList(0, serverRequest)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("%+v\n", sl)
	// err = ubus.RcInit(0, goubus.UbusRcInitRequest{
	// 	Name:   "uwsgi",
	// 	Action: "enable",
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// rclist, err := ubus.RcList(0)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("%+v\n", rclist)
	// network := uci.Network{
	// 	Network: map[string]uci.NetworkInterface{
	// 		"lan": {
	// 			Name:    "lan",
	// 			Proto:   "static",
	// 			Ipaddr:  "192.168.22.1",
	// 			Netmask: "255.255.25.0",
	// 			Dns:     "192.168.22.1",
	// 		},
	// 		"wan": {
	// 			Name:   "wan",
	// 			Proto:  "dhcp",
	// 			Dns:    "",
	// 			Device: "eth0",
	// 		},
	// 	},
	// }
	// n := uci.NewNetwork(ubus.UciTree)
	// n.SetInterfaces(network.Network)
	// err = n.SaveConfig()
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
