package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"vhelper/decode"
	"vhelper/encode"
	"vhelper/time"
	"vhelper/web"
)
import _ "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "help",
	Short: "",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Hello,World!")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(web.WebCommand)
	rootCmd.AddCommand(time.TimeCommand)
	rootCmd.AddCommand(encode.EncodeCommand)
	rootCmd.AddCommand(decode.DecodeCommand)
}

func main() {
	err := rootCmd.Execute()

	if err != nil {
		os.Exit(-1)
	}
}
