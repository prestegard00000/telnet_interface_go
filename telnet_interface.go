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
	Connect		 *telnet.Conn
	//TransmitConn	telnet.Conn                                                                                                                                                                   	
	TransmitTo	telnet.Writer
	TransmitFrom	telnet.Reader
	TransmitContext telnet.Context
}

func (s Telnet_profile) Telnet_connect()(string,error) { //site string, portNumber int) (string, error)
	//@ToDo: change this method to setupProfile and split out the actual calling method
	//@ToDo: clean up
	//create a new terminal
	//currentProfile := Telnet_profile{connection: telnet.StandardCaller}
	//s.Connection.CallTELNET(s.TransmitContext, s.TransmitTo, s.TransmitFrom)
	
	//if currentProfile.Site == "" {
	if s.Site == "" {
		return "", errors.New("Missing site")
	}

	//if currentProfile.PortNumber == 0 {
	if s.PortNumber == 0 {
		s.PortNumber = 5555
	}

	address := fmt.Sprintf("%s:%d", s.Site, s.PortNumber)
	s.Connect, s.TransmitionError = telnet.DialTo(address) // s.Connection)
	//fmt.Println("msg: ", msg)
	//@TODO: return the message
	return address, s.TransmitionError
	//@TODO:return connected if no error
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

	var p []byte
	var count int
	var incomingTransmition string = ""
	count, s.TransmitionError = s.Connect.Read(p)
	var transmitionError error = nil
	fmt.Print(count)
	fmt.Print(p)
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
