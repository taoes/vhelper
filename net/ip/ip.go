package ip

import (
	"fmt"
	"github.com/spf13/cobra"
	"net"
	"strings"
)

var (
	domain  string
	Command = &cobra.Command{
		Use:   "ip",
		Short: "输出当前设备的网络IP信息",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("设备网络:")
			for _, ip := range getIpV6() {
				fmt.Printf("\tip:%-40s\tmask:%s\n", ip.IP.String(), ip.Mask.String())
			}
			fmt.Println("使用中的 IP: ")
			fmt.Println("\t" + getUserIp())
			return nil
		},
	}
)

func init() {
	Command.Flags().StringVar(&domain, "domain", "", "Domain Value")
}

func getUserIp() string {
	dial, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "未发现"
	}
	addr := dial.LocalAddr().String()
	index := strings.LastIndex(addr, ":")
	return addr[:index]
}

func getIpV6() []*net.IPNet {
	ips := make([]*net.IPNet, 0)
	for _, ip := range GetAllIp() {
		if ip.IP.To16() != nil {
			ips = append(ips, ip)
		}
	}
	return ips
}

func GetAllIp() []*net.IPNet {
	ips := make([]*net.IPNet, 0)
	addrs, _ := net.InterfaceAddrs()
	for _, addr := range addrs {
		if ip, ok := addr.(*net.IPNet); ok && !ip.IP.IsLoopback() {
			ips = append(ips, ip)
		}
	}
	return ips
}
