// +build ignore

package main

import (
	"fmt"

	"github.com/aler9/goroslib"
	"github.com/aler9/goroslib/msg"
)

// define a custom service.
// unlike the standard library, a .srv file is not needed.
// two structure definitions are enough, one for the request
// and the other for the response.

type TestServiceReq struct {
	msg.Package `ros:"my_package"`
	A           float64
	B           string
}

type TestServiceRes struct {
	msg.Package `ros:"my_package"`
	C           float64
}

func onService(req *TestServiceReq) *TestServiceRes {
	fmt.Println("request:", req)
	return &TestServiceRes{
		C: 123,
	}
}

func main() {
	// create a node with given name and linked to given master.
	// master can be reached with an ip or hostname.
	n, err := goroslib.NewNode(goroslib.NodeConf{
		Name:       "/goroslib-sp",
		MasterHost: "127.0.0.1",
	})
	if err != nil {
		panic(err)
	}
	defer n.Close()

	// create a service provider
	sp, err := goroslib.NewServiceProvider(goroslib.ServiceProviderConf{
		Node:     n,
		Service:  "/test_srv",
		Callback: onService,
	})
	if err != nil {
		panic(err)
	}
	defer sp.Close()

	// freeze main loop
	select {}
}
