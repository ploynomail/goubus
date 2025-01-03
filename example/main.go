package main

import (
	"fmt"
	"log"
	"time"

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
		Username: "root",
		Password: "",
		URL:      "http://127.0.0.1/ubus",
	}
	fmt.Printf("%+v\n", ubus)
	_, err := ubus.AuthLogin()
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(1 * time.Minute)

	if err := ubus.LoginCheck(); err != nil {
		log.Fatal(err)
	}
	// request := goubus.UbusUciRequest{
	// 	UbusUciRequestGeneric: goubus.UbusUciRequestGeneric{
	// 		Config:  "network",
	// 		Section: "lan",
	// 	},
	// }
	// resp, err := ubus.UciGetConfig(2, request)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%+v\n", resp)

	// err = ubus.UciSetConfig(3, goubus.UbusUciRequest{
	// 	UbusUciRequestGeneric: goubus.UbusUciRequestGeneric{
	// 		Config: "testconfig",
	// 		Type:   "interface",
	// 		Name:   "lan",
	// 	},
	// 	Values: map[string]string{
	// 		"type":   "bridge",
	// 		"proto":  "dhcp",
	// 		"ifname": "eth0",
	// 	},
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// err = ubus.UciSetConfig(3, goubus.UbusUciRequest{
	// 	UbusUciRequestGeneric: goubus.UbusUciRequestGeneric{
	// 		Config:  "network",
	// 		Type:    "interface",
	// 		Section: "wan",
	// 	},
	// 	Values: map[string]string{
	// 		"ifname": "eth0",
	// 	},
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = ubus.UciCommit(4, "testconfig")
	// if err != nil {
	// 	log.Fatal(err)
	// }

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

}
