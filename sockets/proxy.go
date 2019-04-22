package sockets

import (
	"net"
	"os"
	"strings"
)

// GetProxyEnv allows access to the uppercase and the lowercase forms of
// proxy-related variables.  See the Go specification for details on these
// variables. https://golang.org/pkg/net/http/
func GetProxyEnv(key string) string {
	proxyValue := os.Getenv(strings.ToUpper(key))
	if proxyValue == "" {
		return os.Getenv(strings.ToLower(key))
	}
	return proxyValue
}

// Deprecated: Proxies can now be configured using only http.Transport.Proxy
// without configuring http.Transport.Dial.
func DialerFromEnvironment(direct *net.Dialer) (*net.Dialer, error) {
	return direct, nil
}
