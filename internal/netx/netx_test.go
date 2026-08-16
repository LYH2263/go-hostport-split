package netx

import "testing"

func TestIPv6Bracket(t *testing.T) {
	h, p, err := Split("[2001:db8::1]:443")
	if err != nil {
		t.Fatal(err)
	}
	if h != "2001:db8::1" || p != "443" {
		t.Fatalf("host=%q port=%q", h, p)
	}
	h4, p4, err := Split("127.0.0.1:80")
	if err != nil || h4 != "127.0.0.1" || p4 != "80" {
		t.Fatalf("ipv4 %q %q err=%v", h4, p4, err)
	}
}
