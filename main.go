package main

import (
	"errors"
	"github.com/spf13/cobra"
	"os"
	"vhelper/decode"
	"vhelper/encode"
	"vhelper/ip"
	"vhelper/qrcode"
	"vhelper/time"
	"vhelper/web"
)
import _ "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "help",
	Short: "Terminal command-line assistant for development engineers",
	Long:  "Terminal command-line assistant for development engineers",
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("No valid command found, please refer to the following documentation for usage")
	},
}

func init() {
	rootCmd.AddCommand(web.WebCommand)
	rootCmd.AddCommand(time.TimeCommand)
	rootCmd.AddCommand(encode.EncodeCommand)
	rootCmd.AddCommand(decode.DecodeCommand)
	rootCmd.AddCommand(qrcode.QrCodeCommand)
	rootCmd.AddCommand(ip.Command)
}

func main() {
	err := rootCmd.Execute()

	if err != nil {
		os.Exit(-1)
	}
}
