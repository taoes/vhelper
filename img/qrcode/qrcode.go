package qrcode

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/skip2/go-qrcode"
	_ "github.com/skip2/go-qrcode"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

var (
	level         int
	size          int
	filePath      string
	QrCodeCommand = &cobra.Command{
		Use:   "qrcode",
		Short: "生成二维码文件",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errors.New("没有发现需要编码的内容")
			}

			content := args[0]
			encode, err := qrcode.Encode(content, qrcode.Medium, size)
			if err != nil {
				return err
			}

			file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				return err
			}

			absFilePath, err := filepath.Abs(filePath)
			if err != nil {
				return err
			}

			defer file.Close()
			writer := bufio.NewWriter(file)
			writer.Write(encode)
			writer.Flush()

			fmt.Printf("Successfully generated file, And saved in " + absFilePath)
			return nil
		},
	}
)

func init() {
	QrCodeCommand.Flags().IntVar(&level, "level", 1, "二维码等级")
	QrCodeCommand.Flags().IntVar(&size, "size", 256, "图片大小")
	QrCodeCommand.Flags().StringVar(&filePath, "path", "./qrcode.png", "文件保存路径")

}
