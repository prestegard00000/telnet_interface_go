package main

import (
	"github.com/reiver/go-telnet"
)

func main() {
	var caller telnet.Caller = telnet.StandardCaller

	//@TODO: replace "example.net:5555" with address you want to connect to.
	telnet.DailToAndCall("example.net:5555", caller)
}
