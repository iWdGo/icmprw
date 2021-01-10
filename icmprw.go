// Package icmprw implements simple use of the standard net/icmp library.
package icmprw

import (
	"fmt"
	"log"
	"net"
	"time"

	"golang.org/x/net/icmp"
)

// A IcmpParameters contains configuration parameters of an Echo request.
type IcmpParameters struct {
	Network  string // Network name
	NetName  string // Net name appended to Network
	WaitTime int    // Wait interval in seconds to poll for a reply
}

// SimpleIcmp sends an ICMP Echo request and waits for the response.
// parameters the elements of the ICMP request.
// address to target the Echo request
// payload is added to the request
// If an error occurs, it contains an ad hoc message to report the cause of failure.
func SimpleIcmp(address string, wm icmp.Message, parameters IcmpParameters) (rm *icmp.Message, peer net.Addr, err error) {

	protocol := 1 // Default. ICMP protocol number for IPv4
	maxLength := 576
	switch parameters.Network {
	case "ip6":
		protocol = 58
		maxLength = 65575
	}
	netName := parameters.Network + ":" + parameters.NetName
	ipAddr, err := net.ResolveIPAddr(parameters.Network, address)
	if err != nil {
		return nil, peer, fmt.Errorf("%s on %s: %v", address, parameters.Network, err)
	}
	log.Printf("Ping %v using %s", ipAddr, netName)

	c, err := icmp.ListenPacket(netName, ipAddr.String())
	if err != nil {
		return nil, peer, err
	}
	defer c.Close()

	// Send ICMP message
	wb, err := wm.Marshal(nil)
	if err != nil {
		return nil, peer, err
	}
	if _, err := c.WriteTo(wb, ipAddr); err != nil {
		return nil, peer, err
	}

	rb := make([]byte, maxLength)
	var n int
	// Wait WaitTime seconds for the echo
	for i := 0; i < parameters.WaitTime; i++ {
		until := time.Now().Add(1 * time.Second)
		err = c.SetDeadline(until)
		if err != nil {
			return nil, peer, err
		}
		log.Printf("waiting until %v", until)

		n, peer, err = c.ReadFrom(rb)
		if err == nil {
			break
		}
	}
	if err != nil {
		return nil, peer, err
	}

	// Check reflection
	rm, err = icmp.ParseMessage(protocol, rb[:n])
	if err != nil {
		return nil, peer, err
	}
	return rm, peer, nil
}
