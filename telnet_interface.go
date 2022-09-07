package communications

import (
	"errors"
	"fmt"

	"github.com/reiver/go-telnet"
)

func Telnet_connect(site string, portNumber int) (string, error) {
	var caller telnet.Caller = telnet.StandardCaller

	if site == "" {
		return "", errors.New("Missing site")
	}

	if portNumber == 0 {
		portNumber = 5555
	}

	address := fmt.Sprintf("%s:%d", site, portNumber)
	//@TODO: replace "example.net:5555" with address you want to connect to.
	telnet.DialToAndCall(address, caller)
	return "hello", nil
	//return connected if no error
	//return failed connected
}

// if connected, send the string

// if connected, get string

// end connection

// test connection
