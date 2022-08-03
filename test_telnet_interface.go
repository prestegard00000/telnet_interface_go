package communications

import (
	"testing"
	"regexp"
)

// Test for empty sites
func test_telnet_connect(t *testing.T) {
	test_site := ""
	test_portNumber := 5555
	expected_string := "welcome"
	expected_result := regexp.MustCompile(expected_string)
	msg, err := telnet_connect(test_site, test_portNumber)
	if !expected_result.MatchString(msg) || err != nil {
		t.Fatalf("telnet_connect(%s, %d) = %q, nil", test_site, test_portNumber, msg)
	}
}
