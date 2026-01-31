package nettool

import (
	"fmt"
	"net"
	"path/filepath"
	"strings"
	"sync"

	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
)

const (
	// IP2RegionDBPath ip2region 数据库文件路径
	IP2RegionDBPath = "./data/ip2region_v4.xdb"
)

// IPRegionInfo IP 地址解析结果
type IPRegionInfo struct {
	Country  string // 国家
	Region   string // 省份/地区
	City     string // 城市
	ISP      string // ISP 运营商
	FullInfo string // 完整信息字符串
}

// ip2regionSearcher ip2region 查询器实例（单例模式）
var (
	ip2regionSearcher *xdb.Searcher
	ip2regionOnce     sync.Once
	ip2regionInitErr  error
)

// initIP2RegionSearcher 初始化 ip2region 查询器
func initIP2RegionSearcher() error {
	ip2regionOnce.Do(func() {
		// 获取数据库文件的绝对路径
		dbPath, err := filepath.Abs(IP2RegionDBPath)
		if err != nil {
			ip2regionInitErr = fmt.Errorf("获取数据库文件路径失败: %w", err)
			return
		}

		// 创建查询器实例，使用 IPv4 版本
		searcher, err := xdb.NewWithFileOnly(xdb.IPv4, dbPath)
		if err != nil {
			ip2regionInitErr = fmt.Errorf("初始化 ip2region 查询器失败: %w", err)
			return
		}

		ip2regionSearcher = searcher
	})

	return ip2regionInitErr
}

// IP2Region 通过 ip2region 解析 IP 地址，返回地址信息
// 参数:
//   - ip: 要解析的 IP 地址字符串，例如 "8.8.8.8"
//
// 返回:
//   - *IPRegionInfo: IP 地址解析结果，包含国家、省份、城市、ISP 等信息
//   - error: 如果解析失败返回错误
func IP2Region(ip string) (*IPRegionInfo, error) {
	// 初始化查询器（如果还未初始化）
	if err := initIP2RegionSearcher(); err != nil {
		return nil, err
	}

	// 查询 IP 地址
	region, err := ip2regionSearcher.SearchByStr(ip)
	if err != nil {
		return nil, fmt.Errorf("查询 IP 地址失败: %w", err)
	}

	// 解析返回的字符串，格式通常为: "国家|省份|城市|ISP|..."
	// 例如: "中国|0|北京|北京|联通"
	info := parseRegionString(region)

	return info, nil
}

// parseRegionString 解析 ip2region 返回的字符串
// 格式: "国家|省份|城市|ISP|..."
func parseRegionString(region string) *IPRegionInfo {
	info := &IPRegionInfo{
		FullInfo: region,
	}

	// 按 | 分割字符串
	parts := strings.Split(region, "|")

	// 根据分割结果填充字段
	if len(parts) > 0 {
		info.Country = parts[0]
	}
	if len(parts) > 1 {
		info.Region = parts[1]
	}
	if len(parts) > 2 {
		info.City = parts[2]
	}
	if len(parts) > 3 {
		info.ISP = parts[3]
	}

	return info
}

// IsPublicIP 判断一个 IP 是否为公网 IP
func IsPublicIP(ip net.IP) bool {
	if ip == nil {
		return false
	}

	// 1. 排除非单播地址（如组播地址）
	if ip.IsGlobalUnicast() {
		return false
	}

	// 2. 排除私有地址段 (RFC 1918):
	// 10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16
	// 以及 IPv6 的唯一本地地址 (fc00::/7)
	if ip.IsPrivate() {
		return false
	}

	// 3. 排除回环地址 (127.0.0.0/8, ::1)
	if ip.IsLoopback() {
		return false
	}

	// 4. 排除链路本地地址 (169.254.0.0/16, fe80::/10)
	if ip.IsLinkLocalUnicast() {
		return false
	}

	return true
}
