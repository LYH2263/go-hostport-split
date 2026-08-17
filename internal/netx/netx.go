package netx

import "net"

// Split separates host and port from a network address of the form
// "host:port" or "[ipv6]:port". It is a thin wrapper around the standard
// library's net.SplitHostPort, which correctly handles bracketed IPv6
// literals (e.g. "[2001:db8::1]:443") instead of splitting on the colons
// inside the address.
func Split(hostport string) (host, port string, err error) {
	return net.SplitHostPort(hostport)
}
