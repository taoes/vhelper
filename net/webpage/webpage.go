package webpage

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
	"runtime"
)

var (
	Command = &cobra.Command{
		Use:   "url",
		Short: "快速打开网页",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errors.New("请输入网页地址")
			}
			url := args[0]
			err := openURl(url)
			return err
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
		return fmt.Errorf("不支持的操作系统: %s ", runtime.GOOS)
	}

	cmd := exec.Command(run, uri)
	return cmd.Start()
}

func init() {

}
