package nettool

import (
	"net"
	"os"
	"time"

	"github.com/glstr/futty_golang/context"
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
	Country  string
	Region   string
	City     string
	ISP      string
}

func TraceRouter(logger *context.LogBuffer, target string) ([]*TraceRouterResult, error) {
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
			logger.WriteLog("write to target failed:%s", err.Error())
			continue
		}

		// 5. 等待响应
		reply := make([]byte, 1500)
		c.SetReadDeadline(time.Now().Add(Timeout))
		n, peer, err := c.ReadFrom(reply)
		duration := time.Since(start)

		if err != nil {
			logger.WriteLog("%d: * * * timeout", ttl)
			result.Addr = "*.*.*.*"
			result.Duration = duration
			results = append(results, result)
			continue
		}

		// 6. 解析响应类型
		rm, err := icmp.ParseMessage(1, reply[:n])
		switch rm.Type {
		case ipv4.ICMPTypeTimeExceeded:
			// 路由器的响应
			result.Duration = duration
			result.Network = peer.Network()
			result.Addr = peer.String()
			fillAddRegionInfo(logger, result.Addr, result)
			results = append(results, result)

		case ipv4.ICMPTypeEchoReply:
			// 目标主机的响应
			result.Duration = duration
			result.Network = peer.Network()
			result.Addr = peer.String()
			fillAddRegionInfo(logger, result.Addr, result)
			results = append(results, result)
			return results, nil
		default:
			logger.WriteLog("%d, unknown type:%v from %s", ttl, rm.Type, peer)
		}
	}
	return results, nil
}

func fillAddRegionInfo(logger *context.LogBuffer, addr string, result *TraceRouterResult) {
	if !IsPublicIP(net.IP(addr)) {
		logger.WriteLog("not public addr:%s ", addr)
		return
	}

	regionInfo, err := IP2Region(addr)
	if err != nil {
		logger.WriteLog("ip2region failed:%s, ip:%s ", err.Error(), addr)
		return
	}

	result.Country = regionInfo.Country
	result.Region = regionInfo.Region
	result.City = regionInfo.City
	result.ISP = regionInfo.ISP
}
