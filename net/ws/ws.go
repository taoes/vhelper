package ws

import (
	"errors"
	"github.com/spf13/cobra"
	"strings"
)

var (
	url     string
	Command = &cobra.Command{
		Use:   "ws",
		Short: "建立WebSocket连接，并进行调试",
		RunE: func(cmd *cobra.Command, args []string) error {
			if strings.Compare(url, "") == 0 {
				return errors.New("请输入一个有效的WebSocket地址")
			}
			return nil
		},
	}
)

func init() {
	Command.Flags().StringVar(&url, "url", "", "WebSocket地址, 支持ws/wss 协议")

}
