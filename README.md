[![Go Reference](https://pkg.go.dev/badge/github.com/iwdgo/icmprw.svg)](https://pkg.go.dev/github.com/iwdgo/icmprw)
[![Go Report Card](https://goreportcard.com/badge/github.com/iwdgo/icmprw)](https://goreportcard.com/report/github.com/iwdgo/icmprw)
[![codecov](https://codecov.io/gh/iWdGo/icmprw/branch/master/graph/badge.svg)](https://codecov.io/gh/iWdGo/icmprw)

[![Build Status](https://travis-ci.com/iWdGo/icmprw.svg?branch=master)](https://travis-ci.com/iWdGo/icmprw)

# Send and receive ICMP packets

An example for an Echo Request (ping) is provided for IPv6 and the IPv4 test can be used as a reference.

## Use

No privilege should be required when using networks `ip4:icmp` and `ip6:ipv6-icmp` as demonstrated.
Any other network name will go through UDP which usually requires privileges for binding.

## Test

```
go get github.com/iwdgo/icmprw
cd <path to repo>
go test

```

Default set up of the OS should allow the use of the module.
Some tests are using invalid values which can trigger firewall rules.
When re-running invalid values many times, some messages may differ and related tests will fail.

### Sample output

```
>go test -run=Example
2021/01/16 14:50:58 Ping ::1 using ip6:ipv6-icmp
2021/01/16 14:50:58 waiting until 2021-01-16 14:50:59.2003359 +0100 CET m=+1.018128401
2021/01/16 14:50:58 got reflection from ::1  with ID: 32164 (Code and Seq ignored)
PASS
ok      github.com/iwdgo/icmprw
```

## Other references

ICMP (Internet Control Message Protocol) basics:
- [IPv4 wiki](https://en.wikipedia.org/wiki/Internet_Control_Message_Protocol).
- [IPv6 wiki](https://en.wikipedia.org/wiki/Internet_Control_Message_Protocol_for_IPv6).

stackoverflow:
- [how to listen for icmp packets](https://stackoverflow.com/questions/33345683/how-to-listen-for-icmp-packets)
- [how to ping an ip address](https://stackoverflow.com/questions/31868639/how-to-ping-an-ip-address-in-golang)
