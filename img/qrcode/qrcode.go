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
		Short: "Encode the given string text into QR code and save it to file",
		Long:  "Encode the given string text into QR code and save it to file",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errors.New("No found content found")
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
	QrCodeCommand.Flags().IntVar(&level, "level", 1, "picture level")
	QrCodeCommand.Flags().IntVar(&size, "size", 256, "picture size")
	QrCodeCommand.Flags().StringVar(&filePath, "file", "/tmp/qrcode.png", "picture command")

}
