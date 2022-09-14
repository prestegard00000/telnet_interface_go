package telnet_interface_go

import (
	"errors"
	"fmt"
	"github.com/reiver/go-telnet"
)

type Telnet_profile struct {
	Connection       telnet.Caller
	PortNumber       int
	Site             string
	TransmitionError error
	Connected        bool
}

func (s Telnet_profile) Telnet_connect()(string, error) { //site string, portNumber int) (string, error) {
	//create a new terminal
	currentProfile  := Telnet_profile{Connection: telnet.StandardCaller}
	if currentProfile.Site == "" {
		return "", errors.New("Missing site")
	}

	if currentProfile.PortNumber == 0 {
		currentProfile.PortNumber = 5555
	}

	address := fmt.Sprintf("%s:%d", currentProfile.Site, currentProfile.PortNumber)
	//@TODO: replace "example.net:5555" with address you want to connect to.
	var msg string
	telnet.DialToAndCall(address, currentProfile.Connection)
	fmt.Println("msg: ", msg)
	//@TODO: return the message
	return "hello", nil
	//@TODO:return connected if no error
	//return failed connected
}

func (s Telnet_profile) Telnet_transmitString(transmition string) (string, error) {

	var incomingTransmition string = ""
	var transmitionError error = nil
	return incomingTransmition, transmitionError
	//if telnet.Caller is not nil
	//send transmition to	telnet.Caller
	//if the tranmition does not error out, collect return string and return
}

func (s Telnet_profile) Telnet_recieveString() (string, error) {

	var incomingTransmition string = ""
	var transmitionError error = nil
	return incomingTransmition, transmitionError
	//if telnet.Caller is not nil
	//recieveString
	//if the string does not error, return it
}

func (s Telnet_profile) Telnet_endConnection(exit string) (string, error) {
	var incomingTransmition string = ""
	var transmitionError error = nil
	return incomingTransmition, transmitionError
	//if telnet.Caller is not nil
	//send exit string and check that the connection terminates
	//return string and close terminal
}

func (s Telnet_profile) Telnet_testConneciton() error { // test connection
	var connectionError error = nil
	return connectionError
	//does telnet.Caller exist?
	//if it does log
	//if it doesn't, error
}
