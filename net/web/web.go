package web

import (
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
	"path/filepath"
	"strconv"
)

var (
	host    string
	port    int
	path    string
	Command = &cobra.Command{
		Use:   "web",
		Short: "在本地设备运行静态资源",
		RunE: func(cmd *cobra.Command, args []string) error {
			abs, err := filepath.Abs(path)
			if err != nil {
				return err
			}
			fmt.Printf("静态资源服务器启动在[%s:%s],资源文件路径:%s .....", host, strconv.Itoa(port), abs)
			fmt.Println()
			http.Handle("/", http.FileServer(http.Dir(path)))
			err = http.ListenAndServe(host, nil)
			return err
		},
	}
)

func init() {
	Command.Flags().StringVar(&host, "host", "127.0.0.1", "运行主机地址")
	Command.Flags().IntVar(&port, "port", 3364, "运行端口")
	Command.Flags().StringVar(&path, "path", ".", "静态资源路径")
}
