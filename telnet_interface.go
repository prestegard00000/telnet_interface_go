package communications

import (
	"errors"
	"fmt"

	"github.com/reiver/go-telnet"
)

func Telnet_connect(site string, portNumber int) (string, error) {
	//create a new terminal
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

func Telnet_transmitString(transmition string) (string, error) {

	var incomingTransmition string = ""
	var transmitionError error = nil
	return incomingTransmition, transmitionError
	//if telnet.Caller is not nil
	//send transmition to	telnet.Caller
	//if the tranmition does not error out, collect return string and return
}

func Telnet_recieveString() (string, error) {

	var incomingTransmition string = ""
	var transmitionError error = nil
	return incomingTransmition, transmitionError
	//if telnet.Caller is not nil
	//recieveString
	//if the string does not error, return it
}

func Telnet_endConnection(exit string) (string, error) {
	var incomingTransmition string = ""
	var transmitionError error = nil
	return incomingTransmition, transmitionError
	//if telnet.Caller is not nil
	//send exit string and check that the connection terminates
	//return string and close terminal
}

func Telnet_testConneciton() error { // test connection
	var connectionError error = nil
	return connectionError
	//does telnet.Caller exist?
	//if it does log
	//if it doesn't, error
}
