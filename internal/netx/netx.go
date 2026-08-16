package netx

import (
	"fmt"
	"strings"
)

func Split(hostport string) (host, port string, err error) {
	i := strings.IndexByte(hostport, ':') // BUG: first colon
	if i < 0 {
		return "", "", fmt.Errorf("missing port")
	}
	return hostport[:i], hostport[i+1:], nil
}
