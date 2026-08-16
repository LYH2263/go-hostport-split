package netx

import "net"

func Split(hostport string) (host, port string, err error) {
	return net.SplitHostPort(hostport)
}
