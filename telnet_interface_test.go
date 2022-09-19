package telnet_interface_go

import (
	"testing"
	"regexp"
)

// Test for empty Sites
func test_Telnet_connect(t *testing.T) {
	var testConnection TelnetProfile
	testConnection.NewTelnetProfile("", 1234)
	expected_string:=""
	expected_result := regexp.MustCompile(expected_string)
	msg, err := testConnection.Telnet_connect()
	if !expected_result.MatchString(msg) || err != nil {
		t.Fatalf("Telnet_connect(%s, %d) = %q, nil", testConnection.site, testConnection.portNumber, msg)
	}
}
