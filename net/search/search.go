package search

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/exp/maps"
	"os/exec"
	"runtime"
)

var engineUrl = map[string]string{
	"baidu":  "https://www.baidu.com?wq=%s",
	"bing":   "https://bing.com/search?q=%s",
	"google": "https://google.com/search?q=%s",
	"github": "https://github.com/search?q=%s",
}
var (
	eng     string
	Command = &cobra.Command{
		Use:   "search",
		Short: "快速打开搜索引擎",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errors.New("请输入搜索内容")
			}

			key := args[0]
			engName := eng
			if engName == "" {
				engName = viper.GetString("searchEngineName")
			}

			engineUrl, ok := engineUrl[engName]
			if !ok {
				return errors.New(fmt.Sprintf("不支持的搜索引擎:%s", engName))
			}
			return openURl(fmt.Sprintf(engineUrl, key))
		},
	}
)

var commands = map[string]string{
	"windows": "cmd /c start",
	"darwin":  "open",
	"linux":   "xdg-open",
}

func openURl(uri string) error {
	run, ok := commands[runtime.GOOS]
	if !ok {
		return fmt.Errorf("不支持的操作系统：%s", runtime.GOOS)
	}

	cmd := exec.Command(run, uri)
	return cmd.Start()
}

func init() {
	marshal, _ := json.Marshal(maps.Keys(engineUrl))
	Command.Flags().StringVar(&eng, "eng", "", "可选的搜索引擎:"+string(marshal))
}
