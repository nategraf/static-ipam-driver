package main

import (
	"github.com/docker/go-plugins-helpers/ipam"
	"github.com/nategraf/static-ipam-driver/driver"
)

const socketAddress = "/run/docker/plugins/static.sock"

func main() {
	d := &driver.Driver{}
	h := ipam.NewHandler(d)
	h.ServeUnix(socketAddress, 0)
}
