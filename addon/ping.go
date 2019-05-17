package addon

import (
	"fmt"
	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
	"golang.org/x/net/ipv6"
	"net"
	"os"
	"strings"
	"time"
)

func ping_(addr string, isIpv6 bool) (int64, error) {
	// 1 for ip4, 58 for ip6
	MsgProtocol := 1

	// ICMP message about to be sent
	m := icmp.Message{
		Type: ipv4.ICMPTypeEcho, Code: 0,
		Body: &icmp.Echo{
			ID: os.Getpid() & 0xffff, Seq: 1,
			Data: []byte(""),
		},
	}

	// ip4
	bindAddr := "0.0.0.0"
	bindProto := "ip4:icmp"
	proto := "ip4"

	// ip6
	if isIpv6 {
		bindAddr = "::"
		proto = "ip6"
		bindProto = "ip6:ipv6-icmp"
		MsgProtocol = 58

		m.Type = ipv6.ICMPTypeEchoRequest
	}

	conn, e := icmp.ListenPacket(bindProto, bindAddr)
	if e != nil {
		return 0, e
	}
	defer conn.Close()

	ipAddr, e := net.ResolveIPAddr(proto, addr)
	if e != nil {
		return 0, e
	}
	bytes, e := m.Marshal(nil)

	if e != nil {
		return 0, e
	}

	start := time.Now()
	n, e := conn.WriteTo(bytes, ipAddr)
	if e != nil {
		return 0, e
	}

	reply := make([]byte, 1500)
	e = conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	if e != nil {
		return 0, e
	}
	n, peer, e := conn.ReadFrom(reply)
	if e != nil {
		return 0, e
	}
	duration := time.Since(start)

	replyMessage, e := icmp.ParseMessage(MsgProtocol, reply[:n])
	if e != nil {
		return 0, e
	}
	switch replyMessage.Type {
	case ipv4.ICMPTypeEchoReply, ipv6.ICMPTypeEchoReply:
		return int64(duration) / 1000 / 1000, nil
	default:
		panic(fmt.Errorf("got %+v from %v; want echo reply", replyMessage, peer))
	}
}

type pinger struct {
	Addr  string
	Proto string
}

func NewPinger(addr string, proto string) *Addon {
	p :=  &pinger{Addr: addr, Proto: proto}
	return &Addon{
		UpdateInterval: 10000 * time.Millisecond,
		Updater:p,
	}
}

func (p *pinger) Update() *Block {
	isIpv6 := strings.ToLower(p.Proto) == "ipv6"
	roundTripMS, e := ping_(p.Addr, isIpv6)
	if e != nil {
		return nil
	}
	msg := fmt.Sprintf(" %s  %dms", IconTime, roundTripMS)
	return &Block{FullText: msg}
}
