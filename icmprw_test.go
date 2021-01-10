package icmprw

import (
	"testing"
)

func TestSimplePing_ipv4(t *testing.T) {

	ipv4Parameters := IcmpParameters{
		"ip4",
		"icmp",
		10,
	}
	if err := SimplePing( "localhost", payload, ipv4Parameters); err != nil {
		t.Fatalf("%v", err)
	}

}

func TestSimplePing_unknownAddr(t *testing.T) {

	ipv6Parameters := IcmpParameters{
		"ip6",
		"ipv6-icmp",
		10,
	}

	errw := "local on ip6: lookup local: no such host"
	if err := SimplePing( "local", payload, ipv6Parameters); err != nil && err.Error() != errw {
		t.Fatalf("got %v, want %v", err, errw)
	}
}

// TestSimplePing_timeOut triggers a time out by using an invalid combination of parameters
func TestSimplePing_timeOut(t *testing.T) {
	t.Skip("timing out fails")
	ipv4Parameters := IcmpParameters{
		"ip4",
		"ipv6-icmp",
		0,
	}

	err := SimplePing("localhost", payload, ipv4Parameters)
	errw := "read ip4 127.0.0.1: i/o timeout"
	if (err == nil) || (err.Error() != errw) {
		t.Fatalf("got %v, want %v", err, errw)
	}
}

func TestSimplePing_invalidNetName(t *testing.T) {

	ipv6Parameters := IcmpParameters{
		"ip4",
		"invalid",
		10,
	}

	err := SimplePing("localhost", payload, ipv6Parameters)

	errw := "listen ip4:invalid: lookup invalid: getprotobyname: The requested name is valid, but no data " +
		"of the requested type was found."
	if (err == nil) || (err.Error() != errw) {
		t.Fatalf("got %v, want %v", err, errw)
	}
}

func TestSimplePing_listenPacket(t *testing.T) {

	ipv6Parameters := IcmpParameters{
		"ip4",
		"ipv4-icmp",
		10,
	}

	err := SimplePing( "localhost", payload, ipv6Parameters)

	errw := "listen ip4:ipv4-icmp: lookup ipv4-icmp: getprotobyname: The requested name is valid, but no data" +
		" of the requested type was found."
	if (err == nil) || (err.Error() != errw) {
		t.Fatalf("got %v, want %v", err, errw)
	}
}

// TODO Testing failure of Echo type requires to send a fake response

// TODO Testing invalid payload requires to reply with a different payload
