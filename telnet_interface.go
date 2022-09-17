package telnet_interface_go

import (
	"errors"
	"fmt"
	"github.com/reiver/go-telnet"
	"bufio"
	//"io/ioutil"
	//"net"
	//"net/http"
	//"os"
)

type Telnet_profile struct {
	Connection       telnet.Caller
	PortNumber       int
	Site             string
	TransmitionError error
	Connected        bool
	Connect		 *telnet.Conn
	Context		telnet.Context
	//TransmitConn	telnet.Conn                                                                                                                                                                   	
	TransmitTo	telnet.Writer
	TransmitFrom	telnet.Reader
	TransmitContext telnet.Context
	TClient		telnet.Client
}

func (s Telnet_profile) Telnet_connect()(string,error) { //site string, portNumber int) (string, error)
	address := fmt.Sprintf("%s:%d", s.Site,s.PortNumber)
	var connection *telnet.Conn
	connection, s.TransmitionError = telnet.DialTo(address)
	s.TransmitContext = telnet.NewContext()
	//s.Connection.CallTELNET(s.TransmitContext, s.TransmitTo, s.TransmitFrom)
	//s.TransmitionError = s.Connection.Call(connection)
	s.TransmitionError = s.TClient.Call(connection)
	
	//type Client -> Caller ->Logger
	//error Client.Call(*Conn)
	//type *Conn (struct)
	//*Conn, error telnet.DialTo(addr string)
	//type Caller interface .CallTELNET(Context, Writer, Reader)	
	//type Context ->Logger
	//Contex NewContext()
	
	//s.TransmitTo.Write(b []byte = "x\n")
	message, _ := bufio.NewReader(connection).ReadString('\n')
	fmt.Print("Message from server: " + message)


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
/*
	//address := fmt.Sprintf("%s:%d", s.Site, s.PortNumber)
	s.Connect, s.TransmitionError = telnet.DialTo(address) // s.Connection)
	//s.Connect.CallTELNET(s.TransmitContext, s.TransmitTo, s.TransmitFrom)
	var msg string
	//telnet.Call()
	msg, s.TransmitionError = s.Telnet_recieveString()
	fmt.Print(msg)*/
	//@TODO: return the message
	return address, s.TransmitionError
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
	var p []byte
	var count int
	count, s.TransmitionError = s.Connect.Read(p)
	
	var transmitionError error = nil
	fmt.Print(count)
	fmt.Print(p)
	return incomingTransmition, transmitionError
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
