package nettool

import (
	"fmt"
	"net"
	"os"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

const (
	MaxHops = 30
	Timeout = 2 * time.Second
)

type TraceRouterResult struct {
	TTL      int
	Network  string
	Addr     string
	Duration time.Duration
	Error    error
}

func TraceRouter(target string) ([]*TraceRouterResult, error) {
	destAddr, err := net.ResolveIPAddr("ip4", target)
	if err != nil {
		return nil, err
	}

	// 1. 监听 ICMP 响应
	c, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
	if err != nil {
		return nil, err
	}
	defer c.Close()

	var results []*TraceRouterResult
	for ttl := 1; ttl <= MaxHops; ttl++ {
		result := &TraceRouterResult{
			TTL: ttl,
		}

		// 2. 设置当前包的 TTL
		c.IPv4PacketConn().SetTTL(ttl)

		// 3. 构建 ICMP Echo Request
		msg := icmp.Message{
			Type: ipv4.ICMPTypeEcho,
			Code: 0,
			Body: &icmp.Echo{
				ID:   os.Getpid() & 0xffff,
				Seq:  ttl,
				Data: []byte("HELLO-GOPHER"),
			},
		}

		wb, _ := msg.Marshal(nil)
		start := time.Now()

		// 4. 发送数据包
		if _, err := c.WriteTo(wb, destAddr); err != nil {
			fmt.Printf("%d: Error sending: %v\n", ttl, err)
			continue
		}

		// 5. 等待响应
		reply := make([]byte, 1500)
		c.SetReadDeadline(time.Now().Add(Timeout))
		n, peer, err := c.ReadFrom(reply)
		duration := time.Since(start)

		if err != nil {
			fmt.Printf("%d:  * * * (Timeout)\n", ttl)
			continue
		}

		// 6. 解析响应类型
		rm, err := icmp.ParseMessage(1, reply[:n])
		switch rm.Type {
		case ipv4.ICMPTypeTimeExceeded:
			// 路由器的响应
			fmt.Printf("%d:  %s  %v\n", ttl, peer, duration)
			result.Duration = duration
			result.Network = peer.Network()
			result.Addr = peer.String()
			results = append(results, result)

		case ipv4.ICMPTypeEchoReply:
			// 目标主机的响应
			fmt.Printf("%d:  %s  %v (Reached)\n", ttl, peer, duration)
			result.Duration = duration
			result.Network = peer.Network()
			result.Addr = peer.String()
			results = append(results, result)
			return results, nil
		default:
			fmt.Printf("%d:  Unknown Type: %v from %s\n", ttl, rm.Type, peer)
		}
	}
	return results, nil
}
