package telnet_interface_go

import (
	"testing"
	"regexp"
)

// Test for empty sites
func test_Telnet_connect(t *testing.T) {
	testConnection := Telnet_profile{
		Site: "",
		PortNumber: 5555,
	}
	expected_string := "welcome"
	expected_result := regexp.MustCompile(expected_string)
	msg, err := testConnection.Telnet_connect()
	if !expected_result.MatchString(msg) || err != nil {
		t.Fatalf("Telnet_connect(%s, %d) = %q, nil", testConnection.Site, testConnection.PortNumber, msg)
	}
}

func test_Telnet_trasmitString(t *testing.T) {
	testConnection := Telnet_profile {
		site: "horizons.jpl.nasa.gov",
		portNumber: 6775,
	}
	expected_string := "nasa"
	expected_result := regexp.MustCompile(expected_string)
	msg,err := testConnection.Telnet_connect()
	if !expected_result.MatchString(msg) || err !=nil {
		t.Fatalf("Telnet_connect(%s, %d) = %q, nil", testConnection.Site, testConnection.PortNumber, msg)
	}
}
