package icmprw

import (
	"bytes"
	"fmt"
	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
	"golang.org/x/net/ipv6"
	"log"
	"os"
)

func SimplePing(address string, payload string, parameters IcmpParameters) error {

	msgId := os.Getpid() & 0xffff // Make up unique 4 bytes id
	wm := icmp.Message{
		Type: ipv4.ICMPTypeEcho, Code: 0,
		Body: &icmp.Echo{
			ID: msgId, Seq: 1,
			Data: []byte(payload),
		},
	}
	switch parameters.Network {
	case "ip6":
		wm = icmp.Message{
			Type: ipv6.ICMPTypeEchoRequest, Code: 0,
			Body: &icmp.Echo{
				ID: msgId, Seq: 1,
				Data: []byte(payload),
			},
		}
	}
	rm, peer, err := SimpleIcmp(address, wm, parameters)
	if err != nil {
		return err
	}

	body, ok := rm.Body.(*icmp.Echo)
	if !ok {
		return fmt.Errorf("reply is not of icmp.Echo type")
	}
	log.Printf("got reflection from %v  with ID: %v (Code and Seq ignored)", peer, body.ID)
	switch rm.Type {
	case ipv4.ICMPTypeEchoReply, ipv6.ICMPTypeEchoReply:
		if !bytes.Equal([]byte(payload), body.Data) {
			return fmt.Errorf("payload differ: sent %s, received %s", payload, body.Data)
		}
	default:
		return fmt.Errorf("got %+v; want echo reply", rm)
	}
	return nil

}
