package pboard

import (
	"errors"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var (
	stdout  bool
	file    string
	Command = &cobra.Command{
		Use:   "board",
		Short: "复制内容到剪切板",
		RunE: func(cmd *cobra.Command, args []string) error {
			all, err := clipboard.ReadAll()
			if err != nil {
				return err
			}
			if stdout {
				fmt.Println(all)
			}
			if strings.Compare(file, "") == 0 {
				return errors.New("未指定保存文件的位置,请使用参数 --file=xxx")
			}

			openFile, err := os.Create(file)
			if err != nil {
				return err
			}
			openFile.Write([]byte(all))
			return nil
		},
	}
)

func init() {
	Command.Flags().StringVar(&file, "file", "", "保存的文件路径")
	Command.Flags().BoolVar(&stdout, "stdout", false, "将内容输出")
}
