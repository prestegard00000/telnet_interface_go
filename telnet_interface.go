package telnet_interface_go

import (
	"errors"
	"fmt"
	"github.com/reiver/go-telnet"
	//"github.com/reiver/go-oi" //getting a public key required error from github.com
	"bufio"
	"net/url"
	"os"
)

type TelnetProfile struct {
	//TelnetProfile fields to store connection information


	/*  The URL of the site that the profile will attempt to connect to
	    This field should be protected and can be set via the:
	    	SetSite(string) method
	*/
	site		string

	/*  The portNumber the telnet_server is listening on
	    This field should be protected and can be set via the:
		SetPortNumber(uint16) method
	    Restrictions: It will always be positive and shouldn't be more than 65,535
	*/
	portNumber	uint16

	/* The go-telnet library has several Types that could be useful.  However, the
	   The documentaion does not give clear instructions on which to utilize in a particular
	   case."
	*/

	/* The telnet.Caller field appears to contain most of the othere data for the telnet.
	  treating it as a protected field variable.  Not sure if it will need to be exposed later
	*/
	caller       telnet.Caller
	
	/* Storing the most current errors for ease of processing
	   Can be a protected class
	*/

	transmitionError error
	




	/* Connection States to ease in processing
	*/
  	Connected        bool


	/* other types from the telnet library.  Not sure which ones are usful and which ones are
	   execess
	*/
	connection	*telnet.Conn
	//Connect		 *telnet.Conn
	//Context		telnet.Context
	//The Writer and Reader are important
	TransmitTo	telnet.Writer
	//TransmitFrom	telnet.Reader
	TransmitContext telnet.Context
	TClient		telnet.Client
}

/* Error handling function in one place, might need to make it part of a utility and import
*/
func checkError(errorValue error, errorHint string) bool {
	if(errorValue == nil) {
		return true
	} else {
		panic(errorValue)
		fmt.Print(errorHint) //	report error
	}
	return false
}

/* NewTelnetProfile will pass initial values, create a profile, dial to a connection, but won't call
*/
func (s *TelnetProfile) NewTelnetProfile(site string, portNumber int) error {
	//initialize new fields
	//call configureProfile Function
	s.TransmitContext = telnet.NewContext()
	s.ConfigureTelnetProfile(site, portNumber)


	var errSetup error
	return errSetup
}

/* All fields should have been initialized. Call validation methods and set the values
*/
func (s *TelnetProfile) ConfigureTelnetProfile(site string, portNumber int) {
	var configurationError error
	var configuredSite string
	var configuredPortNumber uint16

	
	configuredSite, configurationError = s.validateSiteString(site)

	if(checkError(configurationError, "ValidateSiteString")) {
		s.site = configuredSite
	}
	
	configuredPortNumber, configurationError = s.validatePortNumber(portNumber)
	
	if(checkError(configurationError, "ValidatePortNumber")) {
		s.portNumber = configuredPortNumber
	}
	

	var connection *telnet.Conn
	fmt.Print(s.site)
	connection, configurationError = telnet.DialTo(fmt.Sprintf("%s:%d", s.site, s.portNumber))

	if(checkError(configurationError, fmt.Sprintf("DialTo%s:%d", s.site, s.portNumber))) {
		s.connection = connection
	}
		


}

/* validate site string
*/
func (s *TelnetProfile) validateSiteString(siteStringInput string) (string, error) {
	//@ToDO: The string validation is throwing an error on a site that is valid
	//	 needs fixing
	//@ToDO: Use checkError function
	var err_SiteString error
	u, err_SiteString := url.Parse(siteStringInput)
	if err_SiteString != nil || u.Scheme == "" || u.Host == "" {
		fmt.Print("hello")
	}
//		return "", err_SiteString
//	} 
	fmt.Print(siteStringInput) //"u" + u)
	return siteStringInput, err_SiteString
}

/* validate portNumber and convert to uint1t6
*/
func (s *TelnetProfile) validatePortNumber(portNumberInput int) (uint16, error) {
	var err_PortNumber error = nil
	var validatedPortNumber uint16
	if(portNumberInput<0 || portNumberInput > 65535) {
		err_PortNumber = errors.New("Port Number out of range")
		return validatedPortNumber, err_PortNumber
	}
	validatedPortNumber = uint16(portNumberInput)
	return validatedPortNumber, err_PortNumber
}

/* once the TelnetProfile has been configured, dial into the server
*/

func (s *TelnetProfile) Telnet_connect()(string,error) {
	fmt.Print("******Telnet_connect *****")
	prior := bufio.NewScanner(os.Stdin)
	fmt.Println(prior.Text)
	//s.transmitionError = s.TClient.Call(s.connection)
	s.connection, s.transmitionError = s.TClient.Call(s.connection)

	b := make([]byte,100)
	
	n, _ := s.connection..Write([]byte("x\n"))
	fmt.Print('\n'+n+'\n')

	n, _ = s.connection.Read(b)
	fmt.Println(string(b))

	//defer s.connection.Close()	
	//receive transmition
	message, _ := bufio.NewReader(s.connection).ReadString('\n')
	fmt.Print("************Message from server: ")
	fmt.Println(string(message))

	input := bufio.NewScanner(os.Stdin)
	fmt.Println(input.Text)


	//create a new terminal
	var msg string	
	msg = message
	return msg, s.transmitionError
}

func (s TelnetProfile) Telnet_SendCMD(CMD string) (string, error) {
	//send the cmd
	//expecting string or data
	//receive transmition
	var msg string
	var CMD_error error
	return msg, CMD_error
}

func (s TelnetProfile) Telnet_transmitString(transmition string) (string, error) {

	var incomingTransmition string = ""
	var transmitionError error = nil
	//s.TransmitTo
	s.connection.Write([]byte("x\r\n"))//s.TransmitTo)	
	return incomingTransmition, transmitionError
	//if telnet.Caller is not nil
	//send transmition to	telnet.Caller
	//if the tranmition does not error out, collect return string and return
}

func (s TelnetProfile) Telnet_recieveString() (string, error) {

	var incomingTransmition string = ""
	var p []byte
	var count int
	//count, s.TransmitionError = s.Connect.Read(p)
	
	var transmitionError error = nil
	fmt.Print(count)
	fmt.Print(p)
	return incomingTransmition, transmitionError
	//recieveString
	//if the string does not error, return it
}

func (s TelnetProfile) Telnet_endConnection(exit string) (string, error) {
	var incomingTransmition string = ""
	var transmitionError error = nil
	return incomingTransmition, transmitionError
	//if telnet.Caller is not nil
	//send exit string and check that the connection terminates
	//return string and close terminal
}

func (s TelnetProfile) Telnet_testConneciton() error { // test connection
	var connectionError error = nil
	return connectionError
	//does telnet.Caller exist?
	//if it does log
	//if it doesn't, error
}

func (s TelnetProfile) Telnet_grep_menu() error {
	var menuError error = nil
	return menuError
}
