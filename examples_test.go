package icmprw

import (
	"log"
)

const payload = "HELLO-R-U-THERE"

// ExampleSimplePing_ipv6 demonstrates sending and receiving an Echo using IPv6
// Output is ignored as timestamps would fail the test
func ExampleSimplePing_ipv6() {

	ipv6Parameters := IcmpParameters{
		"ip6",
		"ipv6-icmp",
		10, // seconds
	}

	if err := SimplePing( "localhost", payload, ipv6Parameters); err != nil {
		log.Fatalln(err)
	}
	// Output:
}
