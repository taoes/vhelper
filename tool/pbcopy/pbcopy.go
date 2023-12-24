package pbcopy

import (
	"fmt"
	"github.com/atotto/clipboard"
	_ "github.com/atotto/clipboard"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"strings"
)

var (
	stdout  bool
	file    string
	Command = &cobra.Command{
		Use:   "copy",
		Short: "复制内容到剪切板",
		RunE: func(cmd *cobra.Command, args []string) error {
			res := ""

			// 配置文件，尝试冲文件读取
			if strings.Compare(file, "") != 0 {
				f, err := ioutil.ReadFile(file)
				if err != nil {
					return err
				}
				res = string(f)
				if stdout {
					fmt.Println(res)
				}
				return clipboard.WriteAll(res)
			}

			// 有数据，尝试使用数据
			if len(args) > 0 {
				res = args[0]
				if stdout {
					fmt.Println(res)
				}
				return clipboard.WriteAll(res)
			}

			// 无参数，尝试使用重定向流
			for {
				tmpRes, _ := ioutil.ReadAll(os.Stdin)
				if len(tmpRes) == 0 {
					break
				}
				res += string(tmpRes)

			}

			if stdout {
				fmt.Println(res)
			}
			return clipboard.WriteAll(res)
		},
	}
)

func init() {
	Command.Flags().BoolVar(&stdout, "stdout", false, "将内容打印到标准输出")
	Command.Flags().StringVar(&file, "file", "", "需要拷贝的文件路径")
}
